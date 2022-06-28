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
			sessions:     map[string]*accountSession{},
			clientTokens: map[string]string{},
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
	clientToken := req.ClientToken(today)
	session := s.sessionStore.getByClientToken(clientToken)
	if session != nil {
		return session.getLoginResponse(clientToken), nil
	}

	// sessionだけあればclientとつないで返す
	sessionToken := req.SessionToken(today)
	session = s.sessionStore.getBySessionToken(sessionToken)
	if session != nil {
		s.sessionStore.addClientToken(sessionToken, clientToken)
		return session.getLoginResponse(clientToken), nil
	}

	// ログインセッションがなければログイン処理を実行
	session, err := s.tachibana.login(ctx, req)
	if err != nil {
		return nil, s.withErrorDetail(err)
	}

	// セッションの取得に成功しなければ
	if session.Session == nil {
		return session.getLoginResponse(""), nil
	}

	session.Date = today
	session.Token = sessionToken
	s.sessionStore.save(sessionToken, clientToken, session)
	// ログイン結果を返す
	return session.getLoginResponse(clientToken), nil
}

func (s *server) withErrorDetail(err error) error {
	switch e := err.(type) {
	default:
		return e
	}
}

func (s *server) getClientToken(ctx context.Context) (string, bool) {
	const SessionHeaderKey = "session-token" // リクエストヘッダに付けられる認証トークン名

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	tokens := md[SessionHeaderKey]
	if len(tokens) == 0 {
		return "", false
	}

	return tokens[0], true
}
