package tachibana_grpc_server

import (
	"context"
	"time"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type testTachibanaApi struct {
	iTachibanaApi
	login1        *accountSession
	login2        error
	loginHistory  []interface{}
	logout1       *pb.LogoutResponse
	logout2       error
	logoutHistory []interface{}
}

func (t *testTachibanaApi) login(_ context.Context, req *pb.LoginRequest, today time.Time) (*accountSession, error) {
	t.loginHistory = append(t.loginHistory, req, today)
	return t.login1, t.login2
}
func (t *testTachibanaApi) logout(_ context.Context, session *accountSession, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	t.logoutHistory = append(t.logoutHistory, session, req)
	return t.logout1, t.logout2
}
