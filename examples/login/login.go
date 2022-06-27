package main

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

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
	var header metadata.MD

	// ログイン
	{
		res, err := cli.Login(context.Background(), &tachibanapb.LoginRequest{
			UserId:   examples.UserId,
			Password: examples.Password,
		})
		log.Printf("%+v, %+v\n", res, err)

		if err == nil {
			header = metadata.Pairs("session-token", res.Token)
		}
	}

	// ログアウト
	{
		res, err := cli.Logout(metadata.NewOutgoingContext(context.Background(), header), &tachibanapb.LogoutRequest{})
		log.Printf("%+v, %+v\n", res, err)
	}
}
