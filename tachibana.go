package tachibana_grpc_server

import (
	"context"
	"time"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type iTachibanaApi interface {
	login(ctx context.Context, req *pb.LoginRequest, today time.Time) (*loginSession, error)
}

type tachibanaApi struct {
	client tachibana.Client
}

func (t *tachibanaApi) login(ctx context.Context, req *pb.LoginRequest, today time.Time) (*loginSession, error) {
	res, err := t.client.Login(ctx, *t.toLoginRequest(req))
	if err != nil {
		return nil, err
	}
	session, _ := res.Session()
	return &loginSession{
		date:     today,
		session:  session,
		response: t.fromLoginResponse(res),
	}, nil
}
