package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	//"time"

	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	//"k8s.io/apimachinery/pkg/api/resource"

	//types "k8s.io/apimachinery/pkg/types"

	//v2 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumesnapshot/v1"
	//"github.com/kubernetes-csi/external-snapshotter/client/v8/clientset/versioned"
	//batchv1 "k8s.io/api/batch/v1"
	//v1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/rest"

	//"k8s.io/client-go/tools/clientcmd"
	pr "github.com/abhishek-kamat-nutanix/read-write-grpc/backup/proto"
)

func(s *Server)MigrateConfig(ctx context.Context,in *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	address := in.ServerAddr

	tFile, err := os.CreateTemp("","configs-*.json")
	if err!=nil{
		log.Printf("error creating file: %v\n",err)
	}
	defer os.Remove(tFile.Name())
	
	ns := in.Namespace
	rs := in.Resources
	labl := in.Labels
	//storageClassName:= in.Storageclassname
	//str := "kubectl get " + rs + " -n " + ns + " -l " + labl + " -o yaml > /yaml/manifests.yaml"
	cmd := exec.Command("sh", "-c", fmt.Sprintf("kubectl get %s -n %s -l %s -o json > %s", rs, ns, labl, tFile.Name()))
	output, err := cmd.CombinedOutput()
	if err != nil{
		log.Printf("kubectl get failed with error %v\n %s",err,string(output))
	} else {log.Printf("kubectl command completed successfully")}

	cmd = exec.Command("sh", "-c", fmt.Sprintf("yq 'del(.items[].metadata.resourceVersion, .items[].metadata.uid, .items[].metadata.creationTimestamp, .items[].status, .items[].spec.clusterIP, .items[].spec.clusterIPs)' -i %s", tFile.Name()))
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("yq eval failed with error: %v\n%s", err, string(output))
	} else {log.Printf("yq cleaning command completed successfully.") }

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	co := pr.NewBackupServiceClient(conn)

	contents, _ := os.ReadFile(tFile.Name())
	if err != nil {
		log.Printf("error reading temp file: %v\n", err)
	}

	res, err := co.SendJSONData(context.Background(),&pr.JSONDataRequest{Jsondata: string(contents)})
	if err != nil {
		log.Printf("error reading temp file: %v\n", err)
	}
	return &pb.ConfigResponse{
		Message: res.Result  ,
	}, nil


	// config, err := rest.InClusterConfig()
	// if err!= nil {
	// 	fmt.Printf("error getting in-cluster config: %v", err)
	// }

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil{
	// 	fmt.Printf("error creating clientset: %v\n", err)
	// }


	// volumeMode := v1.PersistentVolumeFilesystem 
	// persistentVolumeAccessMode := v1.ReadWriteOnce
	// resourceName:= v1.ResourceStorage
	// m := make(v1.ResourceList)
	// volsize := resource.NewQuantity(1*1024*1024*1024, resource.BinarySI)
	// m[resourceName] = *volsize
	// pvclaim := v1.PersistentVolumeClaim{TypeMeta: metav1.TypeMeta{Kind:"PersistentVolumeClaim",APIVersion:"v1"},
	// 									ObjectMeta: metav1.ObjectMeta{Name: "yaml-pv-claim"},
	// 									Spec: v1.PersistentVolumeClaimSpec{StorageClassName: &storageClassName, 
	// 										VolumeMode: &volumeMode, 
	// 										Resources: v1.VolumeResourceRequirements{Limits: v1.ResourceList{},Requests: m}, 
	// 										AccessModes: []v1.PersistentVolumeAccessMode{persistentVolumeAccessMode}},
	// 									Status: v1.PersistentVolumeClaimStatus{}}
										
	
	// create_pvc, err := clientset.CoreV1().PersistentVolumeClaims(ns).Create(context.Background(),&pvclaim,metav1.CreateOptions{})
	// if err != nil{
	// 	fmt.Printf("error while creating diskreader pvc in %v namespace: %v\n",ns,err)
	// }
	// fmt.Printf("pvc created %s\n",create_pvc.UID)

	// var completions int32 = 1
	// var UID int64 = 0
	// labels := make(map[string]string)
	// labels["app"]="app-yamls"
	// command := []string{"/bin/sh","-c"}
	// str := "kubectl get " + rs + " -n " + ns + " -l " + labl + " -o yaml > /yaml/manifests.yaml"
	// //args := []string{"kubectl get svc,secrets,deployments -n wordpress -l app=wordpress -o yaml > /yaml/manifests.yaml"}
	// args := []string{str}
	// // app-config-job job comes here
	// readerjob:= batchv1.Job{TypeMeta: metav1.TypeMeta{Kind: "Job",APIVersion: "batch"},
	// 						ObjectMeta: metav1.ObjectMeta{Name: "app-config-job"},
	// 						Spec: batchv1.JobSpec{Completions: &completions,
	// 							Template: v1.PodTemplateSpec{ObjectMeta:  metav1.ObjectMeta{Labels: labels}, 
	// 							Spec: v1.PodSpec{RestartPolicy: "OnFailure",ImagePullSecrets: []v1.LocalObjectReference{{Name: "my-registry-secret"}}, 
	// 							Volumes: []v1.Volume{{Name: "yamls",VolumeSource: v1.VolumeSource{PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{ClaimName: "yaml-pv-claim"}}}},
	// 							Containers: []v1.Container{{Name: "yaml-getter", Image: "bitnami/kubectl:latest",Command: command,Args: args ,SecurityContext: &v1.SecurityContext{RunAsUser: &UID}, VolumeMounts: []v1.VolumeMount{{Name: "yamls", MountPath: "/yaml"}}}},}}},}

	// reader, err := clientset.BatchV1().Jobs(ns).Create(context.Background(),&readerjob,metav1.CreateOptions{})
	// if err != nil{
	// 	fmt.Printf("error while creating reader job in %v namespace: %v\n",ns,err)
	// }
	// fmt.Printf("app-config-job created %v\n",reader.UID)

	// // clearance of objects once disk reader Job is done
	// job, err := clientset.BatchV1().Jobs(ns).Get(context.Background(),"app-config-job",metav1.GetOptions{})
	// if err != nil{
	// 	fmt.Printf("error while getting job in %v namespace: %v\n",ns,err)
	// }

	// deletePolicy := metav1.DeletePropagationBackground
	// flag := 0
	// 	for {
	// 		time.Sleep(10 * time.Second)
	// 	for _, condition := range job.Status.Conditions {
	// 		if condition.Type == batchv1.JobComplete && condition.Status == v1.ConditionTrue {

	// 			// call migrate volume here
	// 			_, err := s.MigrateVolume(ctx, &pb.VolumeRequest{
	// 				ServerAddr: "10.15.170.49:50051", // writer server address on target cluster
	// 				BackupName: "yaml-pv-claim", // volume on source cluster
	// 				VolumeName: "yaml-pv-claim", // volume name to keep on target cluster
	// 				Namespace: "wordpress",
	// 				Snapclass: "default-snapshotclass", // default-snapshotclass  nutanix-snapshot-class
	// 				Storageclassname: "default-storageclass", // default-storageclass nutanix-volume
	// 			})
			
	// 			if err != nil {
	// 				log.Fatalf("could not MigrateVolume yaml-pv-claim: %v\n", err)
	// 			}

	// 			//delete diskreader job so diskreader pvc is not bound and can be deleted successfully
	// 			err = clientset.BatchV1().Jobs(ns).Delete(context.Background(),"app-config-job",metav1.DeleteOptions{PropagationPolicy: &deletePolicy})	
	// 			if err != nil{
	// 				fmt.Printf("error while deleting job in %v namespace: %v\n",ns,err)
	// 			}
	// 			fmt.Print("app-config-job completed, deleting Job/Pod \n")

	// 			//delete pvc now
	// 			err = clientset.CoreV1().PersistentVolumeClaims(ns).Delete(context.Background(),"yaml-pv-claim",metav1.DeleteOptions{})
	// 			if err != nil{
	// 				fmt.Printf("error while deleting pvc from %v namespace: %v\n",ns ,err)
	// 			}
	// 			fmt.Print("yaml-pv-claim deleted\n")
	// 			flag=1
	// 			break;
	// 		} 
	// 	}
	// 	if flag==1 {
	// 		break
	// 	}
	// 	job, err = clientset.BatchV1().Jobs(ns).Get(context.Background(),"app-config-job",metav1.GetOptions{})	
	// 			if err != nil{
	// 				fmt.Printf("error while getting pvc in %v namespace: %v\n",ns,err)
	// 			}
		
	// 	}


}