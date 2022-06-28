package tachibana_grpc_server

import (
	"context"
	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type iTachibanaApi interface {
	login(ctx context.Context, req *pb.LoginRequest) (*accountSession, error)
}

type tachibanaApi struct {
	client tachibana.Client
}

func (t *tachibanaApi) login(ctx context.Context, req *pb.LoginRequest) (*accountSession, error) {
	res, err := t.client.Login(ctx, *t.toLoginRequest(req))
	if err != nil {
		return nil, err
	}
	session, _ := res.Session()
	return &accountSession{
		Session:      session,
		BaseResponse: t.fromLoginResponse(res),
	}, nil
}
