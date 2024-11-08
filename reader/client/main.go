package main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"
	   
)

var addr string = "10.46.63.221:30051"

func main() {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	c := pb.NewReaderServiceClient(conn)

	doMigrateVolume(c)
}