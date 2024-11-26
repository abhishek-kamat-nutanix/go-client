package main

import (
	"context"
	"fmt"
	"log"

	//"time"

	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"
	//types "k8s.io/apimachinery/pkg/types"

	//v2 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumesnapshot/v1"
	//"github.com/kubernetes-csi/external-snapshotter/client/v8/clientset/versioned"
	//batchv1 "k8s.io/api/batch/v1"
	//v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	//"k8s.io/client-go/rest"
	//"k8s.io/client-go/tools/clientcmd"
)

func (s *Server) MigrateApp(ctx context.Context, in *pb.AppRequest) (*pb.AppResponse, error) {
    // kubeconfig := flag.String("kubeconfig", "/home/nutanix/nke-source.cfg", "location to your kubeconfig file")
    // flag.Parse() // Ensure flags are parsed before use

    // Build the Kubernetes config
    // config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    // if err != nil {
    //     return nil, fmt.Errorf("error getting kubeconfig: %v", err)
    // }
    config, err := rest.InClusterConfig()
	if err!= nil {
		fmt.Printf("error getting in-cluster config: %v", err)
	}

    // Create Kubernetes clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, fmt.Errorf("error creating Kubernetes client: %v", err)
    }

    namespace := in.Namespace
    labl := in.Labels
    serverip := in.ServerAddr
    rsc := in.Resources

    // List PersistentVolumeClaims in the namespace
    pvc, err := clientset.CoreV1().PersistentVolumeClaims(namespace).List(context.Background(), metav1.ListOptions{LabelSelector: labl})
    if err != nil {
        return nil, fmt.Errorf("error listing PVCs in namespace %s: %v", namespace, err)
    }

    // Debug information
    for _,item := range pvc.Items {
    fmt.Printf("PVCs in namespace %s: %v\n", namespace, item.Name)

    _, err := s.MigrateVolume(ctx, &pb.VolumeRequest{
        				ServerAddr: serverip, // writer server address on target cluster
        				BackupName: item.Name, // volume on source cluster
        				Namespace: namespace,
        			})
        if err != nil {
            log.Printf("error migrating pv %v error: %v\n",item.Name,err)
        }

    }

    _, err = s.MigrateConfig(ctx, &pb.ConfigRequest{
        Namespace: namespace,
        Resources: rsc,
        Labels: labl,
        ServerAddr: serverip,
    })
    if err != nil {
        log.Printf("error migrating configs error: %v\n",err)
    }

    return &pb.AppResponse{
        Message: "Backup Volume has Completed",
    }, nil
}
