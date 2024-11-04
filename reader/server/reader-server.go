package main

import (
	"context"
	"fmt"
	"time"
	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"

	v2 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumesnapshot/v1"
	"github.com/kubernetes-csi/external-snapshotter/client/v8/clientset/versioned"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	//"k8s.io/client-go/tools/clientcmd"
)

func (s *Server)MigrateVolume(ctx context.Context,in *pb.VolumeRequest) (*pb.VolumeResponse, error) {

	//kubeconfig := flag.String("kubeconfig","/home/nutanix/nke-source.cfg", "location to your kubeconfig file")
	
	// writer servers IP address
	address := in.ServerAddr
	// address := os.Getenv("GRPC_SERVER_ADDR")
    // if address == "" {
	// 	address = "10.15.170.49:50051"
    //     //log.Fatalf("GRPC_SERVER_ADDR environment variable is not set")
    // }

	// name to keep target pvc
	volumeName := in.VolumeName
	// volumeName := os.Getenv("VOLUME_NAME")
	// if volumeName == "" {
	// 	volumeName = "wp-pv-claim3"
	// 	//log.Fatalf("VOLUME_NAME environment variable is not set") 
	// }

	// source volume to migrate
	backupVolume := in.BackupName
	// backupVolume := os.Getenv("BACKUP_VOLUME")
	// if backupVolume == "" {
	// 	backupVolume = "wp-pv-claim"
	// 	//log.Fatalf("BACKUP_VOLUME environment variable is not set") 
	// }

	//config, err := clientcmd.BuildConfigFromFlags("",*kubeconfig)
	
	
	//fmt.Printf("error building config from flags: %s\n",err.Error())
	config, err := rest.InClusterConfig()
	if err!= nil {
		fmt.Printf("error getting in-cluster config: %v", err)
	}
	

	
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil{
		fmt.Printf("error creating clientset: %v\n", err)
	}

	// get source snapshot details
	pvc, err := clientset.CoreV1().PersistentVolumeClaims("default").Get(context.Background(),backupVolume,metav1.GetOptions{})
	if err != nil{
		fmt.Printf("error while getting pvc %v from default namespace: %v\n", backupVolume,err)
	}

	size := pvc.Spec.Resources.Requests.Storage()

	//fmt.Printf("size of pvc is %v \n",size)

	clientset2, err := versioned.NewForConfig(config)
	if err !=nil {
		fmt.Printf("error creating clientset: %v\n", err)
	}


	// take a snapshot of source pvc
	snapClass := "lvms-vg1"
	snap:= v2.VolumeSnapshot{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Name: "source-snap"},
		Spec: v2.VolumeSnapshotSpec{Source: v2.VolumeSnapshotSource{PersistentVolumeClaimName: &backupVolume}, VolumeSnapshotClassName: &snapClass},
		Status: &v2.VolumeSnapshotStatus{},
	}

	ss, err := clientset2.SnapshotV1().VolumeSnapshots("default").Create(context.Background(),&snap,metav1.CreateOptions{})
	if err != nil{
		fmt.Printf("error while listing pvc from default namespace: %v\n", err)
	}
	fmt.Printf("ss created %s \n",ss.UID)

	// create pvc for diskreader
	storageClassName:=  "lvms-vg1"
	volumeMode := v1.PersistentVolumeBlock 
	persistentVolumeAccessMode := v1.ReadWriteOnce
	resourceName:= v1.ResourceStorage
	m := make(v1.ResourceList)
	m[resourceName] = *size
	apiGroup := "snapshot.storage.k8s.io"
	pvclaim := v1.PersistentVolumeClaim{TypeMeta: metav1.TypeMeta{Kind:"PersistentVolumeClaim",APIVersion:"v1"},
										ObjectMeta: metav1.ObjectMeta{Name: "diskreader-pvc"},
										Spec: v1.PersistentVolumeClaimSpec{StorageClassName: &storageClassName, 
											VolumeMode: &volumeMode, 
											Resources: v1.VolumeResourceRequirements{Limits: v1.ResourceList{},Requests: m}, 
											DataSource: &v1.TypedLocalObjectReference{APIGroup: &apiGroup  ,Kind: "VolumeSnapshot" , Name:"source-snap"},
											AccessModes: []v1.PersistentVolumeAccessMode{persistentVolumeAccessMode}},
										Status: v1.PersistentVolumeClaimStatus{}}
										
	
	create_pvc, err := clientset.CoreV1().PersistentVolumeClaims("default").Create(context.Background(),&pvclaim,metav1.CreateOptions{})
	if err != nil{
		fmt.Printf("error while creating diskreader pvc in default namespace: %v\n",err)
	}
	fmt.Printf("pvc created %s\n",create_pvc.UID)
	

	var completions int32 = 1
	labels := make(map[string]string)
	labels["app"]="reader"
	command := []string{"./client"}
	priv := true
	// diskereader job comes here
	readerjob:= batchv1.Job{TypeMeta: metav1.TypeMeta{Kind: "Job",APIVersion: "batch"},
							ObjectMeta: metav1.ObjectMeta{Name: "reader-job"},
							Spec: batchv1.JobSpec{Completions: &completions,
								Template: v1.PodTemplateSpec{ObjectMeta:  metav1.ObjectMeta{Labels: labels}, 
								Spec: v1.PodSpec{RestartPolicy: "OnFailure",ImagePullSecrets: []v1.LocalObjectReference{{Name: "my-registry-secret"}}, 
								Volumes: []v1.Volume{{Name: "diskreader-pvc",VolumeSource: v1.VolumeSource{PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{ClaimName: "diskreader-pvc"}}}},
								Containers: []v1.Container{{Name: "reader", Image: "abhishekkamat27/grpc_reader:volume",Command: command, Env: []v1.EnvVar{{Name: "GRPC_SERVER_ADDR", Value: address},{Name: "VOLUME_NAME", Value: volumeName}}, SecurityContext: &v1.SecurityContext{Privileged: &priv}, VolumeDevices: []v1.VolumeDevice{{Name: "diskreader-pvc", DevicePath: "/dev/xvda"}}}},}}},}
	reader, err := clientset.BatchV1().Jobs("default").Create(context.Background(),&readerjob,metav1.CreateOptions{})
	if err != nil{
		fmt.Printf("error while creating reader job in default namespace: %v\n",err)
	}
	fmt.Printf("disk reader job created; %v\n",reader.UID)

	// delete snapshot once used
	for create_pvc.Status.Phase!= v1.ClaimBound {
		create_pvc, err = clientset.CoreV1().PersistentVolumeClaims("default").Get(context.Background(),"diskreader-pvc",metav1.GetOptions{})
		if err != nil{
			fmt.Printf("error while getting pvc in default namespace: %v\n",err)
		}
	}
    
	 err = clientset2.SnapshotV1().VolumeSnapshots("default").Delete(context.Background(),"source-snap",metav1.DeleteOptions{})
	 if err != nil{
	 fmt.Printf("error while listing pvc from default namespace: %v\n", err)
	 }
	 fmt.Print("volumesnapshot source-snap deleted\n")


	// clearance of objects once disk reader Job is done
	job, err := clientset.BatchV1().Jobs("default").Get(context.Background(),"reader-job",metav1.GetOptions{})
	if err != nil{
		fmt.Printf("error while getting job in default namespace: %v\n",err)
	}

	deletePolicy := metav1.DeletePropagationBackground
	flag := 0
		for {
			time.Sleep(10 * time.Second)
		for _, condition := range job.Status.Conditions {
			if condition.Type == batchv1.JobComplete && condition.Status == v1.ConditionTrue {
				//delete diskreader job so diskreader pvc is not bound and can be deleted successfully
				err = clientset.BatchV1().Jobs("default").Delete(context.Background(),"reader-job",metav1.DeleteOptions{PropagationPolicy: &deletePolicy})	
				if err != nil{
					fmt.Printf("error while getting job in default namespace: %v\n",err)
				}
				fmt.Print("reader job completed, deleting Job/Pod \n")

				//delete pvc now
				err := clientset.CoreV1().PersistentVolumeClaims("default").Delete(context.Background(),"diskreader-pvc",metav1.DeleteOptions{})
				if err != nil{
					fmt.Printf("error while deleting pvc from default namespace: %v\n", err)
				}
				fmt.Print("diskreader-pvc deleted\n")
				flag=1
				break;
			} 
		}
		if flag==1 {
			break
		}
		job, err = clientset.BatchV1().Jobs("default").Get(context.Background(),"reader-job",metav1.GetOptions{})	
				if err != nil{
					fmt.Printf("error while getting pvc in default namespace: %v\n",err)
				}
		
		}

		// readers rpc Migrate volume finished
		currentTime := time.Now().Format(time.DateTime)
		return &pb.VolumeResponse{
			Message: "Backup Volume has Completed at: " + currentTime ,
		}, nil
		 

	
}