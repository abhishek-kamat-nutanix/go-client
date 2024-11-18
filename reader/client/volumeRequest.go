package main

import (
	"context"
	"log"

	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"
)

func doMigrateVolume(c pb.ReaderServiceClient){
	log.Println("doVolumeMigrate was invoked")
	
	res, err := c.MigrateVolume(context.Background(), &pb.VolumeRequest{
		ServerAddr: "10.15.170.49:50051", // writer server address on target cluster
		BackupName: "mysql-pv-claim", // volume on source cluster
		VolumeName: "mysql-pv-claim", // volume name to keep on target cluster
		Namespace: "wordpress",
		Snapclass: "default-snapshotclass", // default-snapshotclass  nutanix-snapshot-class
		Storageclassname: "default-storageclass", // default-storageclass nutanix-volume
	})

	if err != nil {
		log.Fatalf("could not MigrateVolume: %v\n", err)
	}

	log.Printf("Message recieved from readers server: %v\n", res.Message)
}