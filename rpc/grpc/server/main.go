package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	userpb "grpc/pbgo/user"
)

type userServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

var userData = map[int64]*userpb.User{
	1: {
		Uid:       1,
		Name:      "user1",
		Gender:    userpb.User_GENDER_MALE,
		Favorites: []string{"Apple", "Banana", "Orange"},
	},
	2: {
		Uid:    2,
		Name:   "user2",
		Gender: userpb.User_GENDER_MALE,
	},
	3: {
		Uid:       3,
		Name:      "user3",
		Gender:    userpb.User_GENDER_FEMALE,
		Favorites: []string{"Google", "Apple", "Microsoft"},
	},
}

func (s *userServiceServer) Get(ctx context.Context, in *userpb.GetRequest) (*userpb.GetResponse, error) {
	uid := in.GetId()
	log.Printf("Received: %v\n", uid)

	user, found := userData[uid]
	if !found {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", uid)
	}

	resp := &userpb.GetResponse{
		User: user,
	}
	return resp, nil
}

func main() {
	l, err := net.Listen("tcp", ":8787")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("listen at :8787")

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &userServiceServer{})
	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}
