package tachibana_grpc_server

import (
	"context"
	"google.golang.org/grpc/metadata"
	"time"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

func NewServer() Server {
	return &server{
		tachibana: &tachibanaApi{
			client:         tachibana.NewClient(tachibana.EnvironmentProduction, tachibana.ApiVersionLatest),
			requestTimeout: 3 * time.Second,
		},
		clock: &clock{},
		sessionStore: &sessionStore{
			sessions:     map[string]*accountSession{},
			clientTokens: map[string]string{},
		},
	}
}

type Server interface {
	pb.TachibanaServiceServer
	StartScheduler()
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

func (s *server) NewOrder(ctx context.Context, req *pb.NewOrderRequest) (*pb.NewOrderResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.newOrder(ctx, session.Session, req)
}

func (s *server) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.cancelOrder(ctx, session.Session, req)
}

func (s *server) OrderList(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.orderList(ctx, session.Session, req)
}

func (s *server) OrderDetail(ctx context.Context, req *pb.OrderDetailRequest) (*pb.OrderDetailResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.orderDetail(ctx, session.Session, req)
}

func (s *server) StockMaster(ctx context.Context, req *pb.StockMasterRequest) (*pb.StockMasterResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.stockMaster(ctx, session.Session, req)
}

func (s *server) StockExchangeMaster(ctx context.Context, req *pb.StockExchangeMasterRequest) (*pb.StockExchangeMasterResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.stockExchangeMaster(ctx, session.Session, req)
}

func (s *server) MarketPrice(ctx context.Context, req *pb.MarketPriceRequest) (*pb.MarketPriceResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.marketPrice(ctx, session.Session, req)
}

func (s *server) BusinessDay(ctx context.Context, req *pb.BusinessDayRequest) (*pb.BusinessDayResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.businessDay(ctx, session.Session, req)
}

func (s *server) TickGroup(ctx context.Context, req *pb.TickGroupRequest) (*pb.TickGroupResponse, error) {
	session, err := s.getSession(ctx)
	if err != nil {
		return nil, err
	}
	return s.tachibana.tickGroup(ctx, session.Session, req)
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

func (s *server) getSession(ctx context.Context) (*accountSession, error) {
	clientToken, ok := s.getClientToken(ctx)
	if !ok {
		return nil, notLoggedInErr
	}

	session := s.sessionStore.getByClientToken(clientToken)
	if session == nil {
		return nil, notLoggedInErr
	}
	return session, nil
}
