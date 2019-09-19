package main

import (
	"authservice/pkg/protobuff/api/user"
	pb "authservice/pkg/protobuff/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"strings"
)

const (
	address = "localhost:7777"
)

func main() {

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Creates a new CustomerClient
	client := pb.NewAuthClient(conn)
	execute(client, os.Args)
}

func execute(client pb.AuthClient, args []string) error {
	if len(args) < 3 {
		usage()
		return nil
	}
	cmd := args[1]
	switch strings.ToLower(cmd) {
	case "signup":
		if len(args) != 5 {
			usage()
			return nil
		}
		request := &user.CreateUserRequest{
			Name:     args[2],
			Email:    args[3],
			Password: args[4],
		}

		response, err := client.SignUp(context.Background(), request)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		fmt.Println(response.Token)
		return nil
	case "ping":
		if len(args) != 3 {
			usage()
			return nil
		}
		request := &pb.Empty{}
		token := args[2]
		md := metadata.Pairs("Authorization", "Bearer "+token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		_, err := client.Ping(ctx, request)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		fmt.Println("Ping Successful")
		return nil
	case "get_user_details":
		if len(args) != 3 {
			usage()
			return nil
		}
		request := &pb.Empty{}
		token := args[2]
		md := metadata.Pairs("Authorization", "Bearer "+token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		response, err := client.GetUserDetails(ctx, request)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		fmt.Println("User Id: ", response.Id)
		fmt.Println("User email: ", response.Email)
		fmt.Println("User name: ", response.Name)
		return nil
	case "update_user_details":
		if len(args) != 5 {
			usage()
			return nil
		}
		request := &user.UpdateUserRequest{
			Name:     args[2],
			Password: args[3],
		}
		token := args[4]
		md := metadata.Pairs("Authorization", "Bearer "+token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		response, err := client.UpdateUserDetails(ctx, request)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		fmt.Println("User Id: ", response.Id)
		fmt.Println("User email: ", response.Email)
		fmt.Println("User name: ", response.Name)
		return nil
	default:
		usage()
	}
	return nil
}

func usage() {
	fmt.Println("Usage")
	fmt.Println("Sign Up:             client signup <name> <email> <password>")
	fmt.Println("Ping:                client ping <token>")
	fmt.Println("Get User Details:    client get_user_details <token>")
	fmt.Println("Update User Details: client update_user_details <name> <password> <token>")
	fmt.Println("** token is returned as response of Sign Up")
}
