package main

import (
	"context"
	"log"

	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"
)

func doMigrateConfig(c pb.ReaderServiceClient){
	log.Println("doConfigMigrate was invoked")
	
	res, err := c.MigrateConfig(context.Background(), &pb.ConfigRequest{
		Namespace: "wordpress",
		Resources: "deployments",
		Labels: "app=wordpress",
		ServerAddr: "10.15.170.49:50051",
	})

	if err != nil {
		log.Fatalf("could not MigrateConfig: %v\n", err)
	}

	log.Printf("Message recieved from readers server: %v\n", res.Message)
}