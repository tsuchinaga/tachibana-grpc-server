package tachibana_grpc_server

import (
	"time"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type loginSession struct {
	date     time.Time
	session  *tachibana.Session
	response *pb.LoginResponse
}
