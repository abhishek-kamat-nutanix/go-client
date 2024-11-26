package main

import (
	"context"
	"log"

	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"
)

func doMigrateApp(c pb.ReaderServiceClient){
	log.Println("doAppMigrate was invoked")
	
	res, err := c.MigrateApp(context.Background(), &pb.AppRequest{
		Namespace: "wordpress",
		Resources: "deployments,secrets,svc",
		Labels: "app=wordpress",
		ServerAddr: "10.15.170.49:50051",
	})

	if err != nil {
		log.Fatalf("could not MigrateVolume: %v\n", err)
	}

	log.Printf("Message recieved from readers server: %v\n", res.Message)
}