package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/abhishek-kamat-nutanix/go-client/reader/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.ReaderServiceServer
}

func main() {
	lis, err := net.Listen("tcp",addr)

	if err!=nil {
		log.Fatalf("Failed to listen on: %v\n",err)
	}

	log.Printf("Listening on %s\n", addr)

	s:= grpc.NewServer()

	reflection.Register(s)

	pb.RegisterReaderServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen on: %v\n",err)
	}
}