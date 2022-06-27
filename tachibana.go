package tachibana_grpc_server

import (
	"context"
	"time"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type iTachibanaApi interface {
	login(ctx context.Context, req *pb.LoginRequest, today time.Time) (*accountSession, error)
	logout(ctx context.Context, session *accountSession, req *pb.LogoutRequest) (*pb.LogoutResponse, error)
}

type tachibanaApi struct {
	client tachibana.Client
}

func (t *tachibanaApi) login(ctx context.Context, req *pb.LoginRequest, today time.Time) (*accountSession, error) {
	res, err := t.client.Login(ctx, *t.toLoginRequest(req))
	if err != nil {
		return nil, err
	}
	session, _ := res.Session()
	return &accountSession{
		date:     today,
		session:  session,
		response: t.fromLoginResponse(res),
	}, nil
}

func (t *tachibanaApi) logout(ctx context.Context, session *accountSession, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	res, err := t.client.Logout(ctx, session.session, *t.toLogoutRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromLogoutResponse(res), nil
}
