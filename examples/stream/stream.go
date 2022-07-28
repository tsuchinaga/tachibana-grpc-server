package main

import (
	"context"
	"gitlab.com/tsuchinaga/tachibana-grpc-server/examples"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"sync"

	"gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	if file, err := os.OpenFile("stream.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
		log.SetOutput(file)
	}

	conn, err := grpc.DialContext(context.Background(), examples.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	streaming := func(clientName string) {
		cli := tachibanapb.NewTachibanaServiceClient(conn)
		var header metadata.MD

		// ログイン
		{
			res, err := cli.Login(context.Background(), &tachibanapb.LoginRequest{
				UserId:     examples.UserId,
				Password:   examples.Password,
				ClientName: clientName,
			})
			log.Printf("%s: %+v, %+v\n", clientName, res, err)
			if res.Token == "" {
				log.Println(clientName, "can not get token")
			}
			header = metadata.Pairs("session-token", res.Token)
		}

		// streaming
		{
			//timeout := time.Duration(rand.Float64()*15)*time.Second + 5*time.Second
			//log.Println(clientName, "timeout", timeout)
			//ctx, cf := context.WithTimeout(context.Background(), timeout)
			ctx, cf := context.WithCancel(context.Background())
			stream, err := cli.Stream(metadata.NewOutgoingContext(ctx, header), &tachibanapb.StreamRequest{
				EventTypes: []tachibanapb.EventType{
					tachibanapb.EventType_EVENT_TYPE_ERROR_STATUS,
					tachibanapb.EventType_EVENT_TYPE_KEEPALIVE,
					tachibanapb.EventType_EVENT_TYPE_MARKET_PRICE,
					tachibanapb.EventType_EVENT_TYPE_CONTRACT,
					tachibanapb.EventType_EVENT_TYPE_NEWS,
					tachibanapb.EventType_EVENT_TYPE_SYSTEM_STATUS,
					tachibanapb.EventType_EVENT_TYPE_OPERATION_STATUS,
				},
				ReceiveResend: true,
				StreamIssues: []*tachibanapb.StreamIssue{
					{IssueCode: "1459", Exchange: tachibanapb.Exchange_EXCHANGE_TOUSHOU},
					{IssueCode: "1475", Exchange: tachibanapb.Exchange_EXCHANGE_TOUSHOU},
					{IssueCode: "1476", Exchange: tachibanapb.Exchange_EXCHANGE_TOUSHOU},
					{IssueCode: "1552", Exchange: tachibanapb.Exchange_EXCHANGE_TOUSHOU},
					{IssueCode: "1699", Exchange: tachibanapb.Exchange_EXCHANGE_TOUSHOU},
				},
			})
			if err != nil {
				log.Println(clientName, "can not get stream", err)
				cf()
				return
			}
			for {
				res, err := stream.Recv()
				if err != nil {
					log.Println(clientName, err)
					cf()
					break
				}
				log.Println(clientName, res)
			}
		}
	}

	clients := map[int]string{
		1: "client-name-001",
		//2: "client-name-002",
		//3: "client-name-003",
	}
	var wg sync.WaitGroup
	for _, c := range clients {
		wg.Add(1)
		clientName := c
		go func() {
			streaming(clientName)
			wg.Done()
		}()
	}
	wg.Wait()
}
