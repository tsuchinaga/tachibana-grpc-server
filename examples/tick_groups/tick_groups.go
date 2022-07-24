package main

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"

	"gitlab.com/tsuchinaga/tachibana-grpc-server/examples"

	"gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), examples.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	cli := tachibanapb.NewTachibanaServiceClient(conn)
	var header metadata.MD

	// ログイン
	{
		res, err := cli.Login(context.Background(), &tachibanapb.LoginRequest{
			UserId:     examples.UserId,
			Password:   examples.Password,
			ClientName: examples.ClientName,
		})
		log.Printf("%+v, %+v\n", res, err)
		if res.Token == "" {
			log.Fatalln("can not get token")
		}
		header = metadata.Pairs("session-token", res.Token)
	}

	// 呼値グループ
	{
		res, err := cli.TickGroup(metadata.NewOutgoingContext(context.Background(), header), &tachibanapb.TickGroupRequest{})
		log.Printf("%+v, %+v\n", res, err)
	}
}
