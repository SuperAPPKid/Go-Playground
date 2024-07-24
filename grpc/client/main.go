package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userpb "github.com/SuperAPPKid/Go-Playground/grpc/pbgo/user"
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("localhost:8787", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	usercli := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := 0
	flagId := flag.Arg(0)
	if flagId == "" {
		fmt.Println("please append userid")
		return
	}
	id, err = strconv.Atoi(flagId)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := usercli.Get(ctx, &userpb.GetRequest{Id: int64(id)})
	if err != nil {
		fmt.Println(err)
		return
	}

	user := r.GetUser()
	if name := user.GetName(); name != "" {
		fmt.Println("name:", name)
	}
	if gender := user.GetGender(); gender != userpb.User_GENDER_UNSPECIFIED {
		var genderStr string
		switch gender {
		case userpb.User_GENDER_MALE:
			genderStr = "male"
		case userpb.User_GENDER_FEMALE:
			genderStr = "female"
		}
		fmt.Println("gender:", genderStr)
	}
	if favorites := user.Favorites; len(favorites) > 0 {
		fmt.Println("favorite:")
		for _, favorite := range favorites {
			fmt.Println(favorite)
		}
	}
}
