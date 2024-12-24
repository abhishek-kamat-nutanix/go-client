package main

import (
	//"context"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"
	   
)

var addr string = "10.15.174.11:30051"
// 10.15.170.150:30051 nke
// 10.46.63.221:30051 ocp
// 10.15.168.215:30051 new ocp 
// 10.15.174.11:30051 ocp2

func main() {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	c := pb.NewReaderServiceClient(conn)

	//doMigrateVolume(c)
	doMigrateConfig(c)
	//doMigrateApp(c)

}