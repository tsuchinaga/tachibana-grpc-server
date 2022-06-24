package tachibana_grpc_server

import (
	"context"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

func NewServer() pb.TachibanaServiceServer {
	return &server{
		tachibana: &tachibanaApi{
			client: tachibana.NewClient(tachibana.EnvironmentProduction, tachibana.ApiVersionLatest),
		},
		clock: &clock{},
		sessionStore: &sessionStore{
			store: map[string]*loginSession{},
		},
	}
}

type server struct {
	pb.UnimplementedTachibanaServiceServer
	tachibana    iTachibanaApi
	clock        iClock
	sessionStore iSessionStore
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// ログイン済みならすでに取得したレスポンスを返す
	today := s.clock.today()
	sessionKey := req.GetKey(today)
	session := s.sessionStore.getSession(sessionKey)
	if session != nil {
		return session.response, nil
	}

	// 未ログインならログイン処理を実行
	session, err := s.tachibana.login(ctx, req, today)
	if err != nil {
		return nil, s.withErrorDetail(err)
	}

	// セッションの取得に成功したら
	if session.session != nil {
		session.response.Token = sessionKey
		s.sessionStore.save(sessionKey, session)
	}

	// ログイン結果を返す
	return session.response, nil
}

func (s *server) withErrorDetail(err error) error {
	switch e := err.(type) {
	default:
		return e
	}
}
