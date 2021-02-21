package main

import (
	"context"
	"log"
	"os"

	"github.com/thesepehrm/simple-grpc/pb/update"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port        = ":3000"
	defaultName = "test-client"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := update.NewUpdateClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch os.Args[1] {
	case "login":

		r, err := c.Login(ctx, &update.LoginRequest{
			Username: defaultName,
			Password: "12345", // token: 3132333435e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
		})
		if err != nil {
			log.Fatalf("could not login: %v", err)
		}

		log.Printf("Login Token: %s", r.GetToken())

		break
	case "logout":

		_, err = c.Logout(ctx, &update.LogoutRequest{
			Token: "3132333435e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Success!")

		break

	case "update":
		stream, err := c.ServerPromotions(ctx, new(emptypb.Empty))
		if err != nil {
			log.Fatal(err)
		}

		for {
			r, err := stream.Recv()
			if r == nil {
				continue
			}
			if err != nil {
				log.Print(err)
			}

			log.Printf("%s: %s\n", r.Update.Type, r.Update.Text)
		}

	}

}
