package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/Assigment4"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// AddUser
	addUserResponse, err := c.AddUser(context.Background(), &pb.User{
		Id:    1,
		Name:  "Dauletkhan Zhambyl",
		Email: "210107089@stu.sdu.edu.kz",
	})
	if err != nil {
		log.Fatalf("AddUser failed: %v", err)
	}
	log.Printf("User added with ID: %d", addUserResponse.Id)

	// GetUser
	getUserResponse, err := c.GetUser(context.Background(), &pb.UserID{Id: 1})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}
	log.Printf("User retrieved: %+v", getUserResponse)

	// ListUsers
	stream, err := c.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("ListUsers failed: %v", err)
	}
	for {
		user, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving user: %v", err)
		}
		log.Printf("User: %+v", user)
	}
}
