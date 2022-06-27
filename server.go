package tachibana_grpc_server

import (
	"context"
	"google.golang.org/grpc/metadata"

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
			store: map[string]*accountSession{},
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
		session.token = sessionKey
		session.response.Token = sessionKey
		s.sessionStore.save(sessionKey, session)
	}

	// ログイン結果を返す
	return session.response, nil
}

func (s *server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	session, ok := s.getSession(ctx)
	if !ok {
		return nil, s.withErrorDetail(notLoggedInErr)
	}

	res, err := s.tachibana.logout(ctx, session, req)
	if err != nil {
		return nil, s.withErrorDetail(err)
	}

	s.sessionStore.remove(session.token)
	return res, nil
}

func (s *server) withErrorDetail(err error) error {
	switch e := err.(type) {
	default:
		return e
	}
}

func (s *server) getSession(ctx context.Context) (*accountSession, bool) {
	const SessionHeaderKey = "session-token" // リクエストヘッダに付けられる認証トークン名

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}

	tokens := md[SessionHeaderKey]
	if len(tokens) == 0 {
		return nil, false
	}

	session := s.sessionStore.getSession(tokens[0])
	if session == nil {
		return nil, false
	}
	return session, true
}
