package tachibana_grpc_server

import (
	"context"
	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
	"time"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type iTachibanaApi interface {
	login(ctx context.Context, req *pb.LoginRequest) (*accountSession, error)
	newOrder(ctx context.Context, session *tachibana.Session, req *pb.NewOrderRequest) (*pb.NewOrderResponse, error)
	cancelOrder(ctx context.Context, session *tachibana.Session, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error)
	orderList(ctx context.Context, session *tachibana.Session, req *pb.OrderListRequest) (*pb.OrderListResponse, error)
	orderDetail(ctx context.Context, session *tachibana.Session, req *pb.OrderDetailRequest) (*pb.OrderDetailResponse, error)
	stockMaster(ctx context.Context, session *tachibana.Session, req *pb.StockMasterRequest) (*pb.StockMasterResponse, error)
	stockExchangeMaster(ctx context.Context, session *tachibana.Session, req *pb.StockExchangeMasterRequest) (*pb.StockExchangeMasterResponse, error)
	marketPrice(ctx context.Context, session *tachibana.Session, req *pb.MarketPriceRequest) (*pb.MarketPriceResponse, error)
	businessDay(ctx context.Context, session *tachibana.Session, req *pb.BusinessDayRequest) (*pb.BusinessDayResponse, error)
	tickGroup(ctx context.Context, session *tachibana.Session, req *pb.TickGroupRequest) (*pb.TickGroupResponse, error)
	stream(ctx context.Context, session *tachibana.Session, req *pb.StreamRequest) (<-chan *pb.StreamResponse, <-chan error)
}

type tachibanaApi struct {
	client         tachibana.Client
	requestTimeout time.Duration
}

func (t *tachibanaApi) login(ctx context.Context, req *pb.LoginRequest) (*accountSession, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.Login(ctx1, *t.toLoginRequest(req))
	if err != nil {
		return nil, err
	}
	session, _ := res.Session()
	return &accountSession{
		Session:      session,
		BaseResponse: t.fromLoginResponse(res),
	}, nil
}

func (t *tachibanaApi) newOrder(ctx context.Context, session *tachibana.Session, req *pb.NewOrderRequest) (*pb.NewOrderResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.NewOrder(ctx1, session, *t.toNewOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromNewOrderResponse(res), nil
}

func (t *tachibanaApi) cancelOrder(ctx context.Context, session *tachibana.Session, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.CancelOrder(ctx1, session, *t.toCancelOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromCancelOrderResponse(res), nil
}

func (t *tachibanaApi) orderList(ctx context.Context, session *tachibana.Session, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.OrderList(ctx1, session, *t.toOrderListRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromOrderListResponse(res), nil
}

func (t *tachibanaApi) orderDetail(ctx context.Context, session *tachibana.Session, req *pb.OrderDetailRequest) (*pb.OrderDetailResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.OrderDetail(ctx1, session, *t.toOrderDetailRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromOrderDetailResponse(res), nil
}

func (t *tachibanaApi) stockMaster(ctx context.Context, session *tachibana.Session, req *pb.StockMasterRequest) (*pb.StockMasterResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.StockMaster(ctx1, session, *t.toStockMasterRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromStockMasterResponse(res), nil
}

func (t *tachibanaApi) stockExchangeMaster(ctx context.Context, session *tachibana.Session, req *pb.StockExchangeMasterRequest) (*pb.StockExchangeMasterResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.StockExchangeMaster(ctx1, session, *t.toStockExchangeMasterRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromStockExchangeMasterResponse(res), nil
}

func (t *tachibanaApi) marketPrice(ctx context.Context, session *tachibana.Session, req *pb.MarketPriceRequest) (*pb.MarketPriceResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.MarketPrice(ctx1, session, *t.toMarketPriceRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromMarketPriceResponse(res), nil
}

func (t *tachibanaApi) businessDay(ctx context.Context, session *tachibana.Session, req *pb.BusinessDayRequest) (*pb.BusinessDayResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.BusinessDay(ctx1, session, *t.toBusinessDayRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromBusinessDayResponse(res), nil
}

func (t *tachibanaApi) tickGroup(ctx context.Context, session *tachibana.Session, req *pb.TickGroupRequest) (*pb.TickGroupResponse, error) {
	ctx1, cf1 := context.WithTimeout(ctx, t.requestTimeout)
	defer cf1()
	res, err := t.client.TickGroup(ctx1, session, *t.toTickGroupRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromTickGroupResponse(res), nil
}

func (t *tachibanaApi) stream(ctx context.Context, session *tachibana.Session, req *pb.StreamRequest) (<-chan *pb.StreamResponse, <-chan error) {
	resCh := make(chan *pb.StreamResponse)
	tResCh, tErrCh := t.client.Stream(ctx, session, *t.toStreamRequest(req))
	go func() {
		defer close(resCh)
		for {
			res, ok := <-tResCh
			if !ok {
				return
			}
			resCh <- t.fromStreamResponse(res)
		}
	}()
	return resCh, tErrCh
}
