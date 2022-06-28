package main

import (
	"context"
	"log"

	"gitlab.com/tsuchinaga/tachibana-grpc-server/examples"

	"gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:8900", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	cli := tachibanapb.NewTachibanaServiceClient(conn)

	// ログイン
	{
		res, err := cli.Login(context.Background(), &tachibanapb.LoginRequest{
			UserId:     examples.UserId,
			Password:   examples.Password,
			ClientName: examples.ClientName,
		})
		log.Printf("%+v, %+v\n", res, err)
	}
}
