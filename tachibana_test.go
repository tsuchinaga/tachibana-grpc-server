package tachibana_grpc_server

import (
	"context"
	"time"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type testTachibanaApi struct {
	iTachibanaApi
	login1       *loginSession
	login2       error
	loginHistory []interface{}
}

func (t *testTachibanaApi) login(_ context.Context, req *pb.LoginRequest, today time.Time) (*loginSession, error) {
	t.loginHistory = append(t.loginHistory, req, today)
	return t.login1, t.login2
}
