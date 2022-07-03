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
	conn, err := grpc.DialContext(context.Background(), "localhost:8900", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	// 銘柄一覧
	{
		res, err := cli.StockMaster(metadata.NewOutgoingContext(context.Background(), header), &tachibanapb.StockMasterRequest{
			Columns: []tachibanapb.StockMasterColumn{
				tachibanapb.StockMasterColumn_STOCK_MASTER_COLUMN_ISSUE_CODE,
				tachibanapb.StockMasterColumn_STOCK_MASTER_COLUMN_NAME,
				tachibanapb.StockMasterColumn_STOCK_MASTER_COLUMN_SHORT_NAME,
				tachibanapb.StockMasterColumn_STOCK_MASTER_COLUMN_TRADING_UNIT,
			},
		})
		log.Printf("%+v, %+v\n", res, err)
	}

	// 銘柄市場一覧
	{
		res, err := cli.StockExchangeMaster(metadata.NewOutgoingContext(context.Background(), header), &tachibanapb.StockExchangeMasterRequest{
			Columns: []tachibanapb.StockExchangeMasterColumn{
				tachibanapb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_ISSUE_CODE,
				tachibanapb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_TICK_GROUP_TYPE,
			},
		})
		log.Printf("%+v, %+v\n", res, err)
	}
}
