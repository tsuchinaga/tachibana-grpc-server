package tachibana_grpc_server

import (
	"context"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type testTachibanaApi struct {
	iTachibanaApi
	login1       *accountSession
	login2       error
	loginHistory []interface{}
}

func (t *testTachibanaApi) login(_ context.Context, req *pb.LoginRequest) (*accountSession, error) {
	t.loginHistory = append(t.loginHistory, req)
	return t.login1, t.login2
}
