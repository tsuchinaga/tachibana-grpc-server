package tachibana_grpc_server

import (
	"time"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type accountSession struct {
	date     time.Time
	token    string
	session  *tachibana.Session
	response *pb.LoginResponse
}
