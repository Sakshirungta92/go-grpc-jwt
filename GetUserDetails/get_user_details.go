package main

import (
	"log"
	"context"
	"os"
	"fmt"
	"google.golang.org/grpc"
	pb "authservice/pkg/protobuff/service"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:7777"
)

func main() {
	args := os.Args
	token := args[1]
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewAuthClient(conn)

	request := &pb.Empty{}

	md := metadata.Pairs("Authorization", "Bearer " + token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	response, err := client.GetUserDetails(ctx, request)
    if err != nil {
        log.Fatalf("fail to dial: %v", err)
    }
   fmt.Println("User Id: ", response.Id)
   fmt.Println("User email: ", response.Email)
   fmt.Println("User name: ", response.Name)
}
	
