package main

import (
	"flag"
	"log"
	"net"

	tachibana_grpc_server "gitlab.com/tsuchinaga/tachibana-grpc-server"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"

	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "8900", "port")
	flag.Parse()

	// サーバーの起動
	ln, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalln(err)
	}

	tachibana := tachibana_grpc_server.NewServer()
	go tachibana.StartScheduler()

	s := grpc.NewServer(grpc.UnaryInterceptor(tachibana.LoggingMiddleware()))
	pb.RegisterTachibanaServiceServer(s, tachibana)
	if err := s.Serve(ln); err != nil {
		log.Fatalln(err)
	}
}
