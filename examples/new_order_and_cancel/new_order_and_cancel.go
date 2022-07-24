package main

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"

	"gitlab.com/tsuchinaga/tachibana-grpc-server/examples"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), examples.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	cli := pb.NewTachibanaServiceClient(conn)
	var header metadata.MD

	// ログイン
	{
		res, err := cli.Login(context.Background(), &pb.LoginRequest{
			UserId:     examples.UserId,
			Password:   examples.Password,
			ClientName: examples.ClientName,
		})
		log.Printf("%+v, %+v\n", res, err)
		if err != nil {
			log.Fatalln(err)
		}
		header = metadata.Pairs("session-token", res.Token)
	}

	// 新規注文
	var orderNumber string
	var executionDate time.Time
	{
		res, err := cli.NewOrder(metadata.NewOutgoingContext(context.Background(), header), &pb.NewOrderRequest{
			AccountType:         pb.AccountType_ACCOUNT_TYPE_SPECIFIC,
			DeliveryAccountType: pb.DeliveryAccountType_DELIVERY_ACCOUNT_TYPE_UNUSED,
			IssueCode:           "1475",
			Exchange:            pb.Exchange_EXCHANGE_TOUSHOU,
			Side:                pb.Side_SIDE_BUY,
			ExecutionTiming:     pb.ExecutionTiming_EXECUTION_TIMING_NORMAL,
			OrderPrice:          0,
			OrderQuantity:       1,
			TradeType:           pb.TradeType_TRADE_TYPE_STANDARD_ENTRY,
			ExpireDate:          nil,
			ExpireDateIsToday:   true,
			StopOrderType:       pb.StopOrderType_STOP_ORDER_TYPE_NORMAL,
			TriggerPrice:        0,
			StopOrderPrice:      0,
			ExitPositionType:    pb.ExitPositionType_EXIT_POSITION_TYPE_UNUSED,
			SecondPassword:      examples.SecondPassword,
			ExitPositions:       nil,
		})
		log.Printf("%+v, %+v\n", res, err)
		if err != nil {
			log.Fatalln(err)
		}
		if res.ResultCode != "0" {
			log.Fatalln(res.ResultCode, res.ResultText)
		}
		orderNumber = res.OrderNumber
		executionDate = res.ExecutionDate.AsTime().In(time.Local)
	}

	// 取消
	{
		res, err := cli.CancelOrder(metadata.NewOutgoingContext(context.Background(), header), &pb.CancelOrderRequest{
			OrderNumber:    orderNumber,
			ExecutionDate:  timestamppb.New(executionDate),
			SecondPassword: examples.SecondPassword,
		})
		if res.ResultCode != "0" {
			log.Fatalln(res.ResultCode, res.ResultText)
		}
		log.Printf("%+v, %+v\n", res, err)
	}
}
