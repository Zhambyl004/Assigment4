package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Assignment4" // Import the generated protobuf package
)

type server struct {
	users []*pb.User // Slice to store users
}

// AddUser adds a new user to the server
func (s *server) AddUser(ctx context.Context, user *pb.User) (*pb.UserID, error) {
	s.users = append(s.users, user)     // Append the new user to the slice
	return &pb.UserID{Id: user.Id}, nil // Return the user ID as response
}

// GetUser retrieves a user by ID
func (s *server) GetUser(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
	for _, user := range s.users {
		if user.Id == userID.Id {
			return user, nil // Return the user if found
		}
	}
	return nil, grpc.Errorf(grpc.Code(grpc.NotFound), "User not found") // Return an error if user not found
}

// ListUsers streams all users to the client
func (s *server) ListUsers(empty *pb.Empty, stream pb.UserService_ListUsersServer) error {
	for _, user := range s.users {
		if err := stream.Send(user); err != nil {
			return err // Return error if sending fails
		}
	}
	return nil // Return nil if streaming is successful
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // Listen for incoming connections on port 50051
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()                                  // Create a new gRPC server instance
	pb.RegisterUserServiceServer(s, &server{})             // Register the UserService server
	log.Println("Server started listening on port :50051") // Log that the server has started
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
