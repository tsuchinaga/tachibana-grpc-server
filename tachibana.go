package tachibana_grpc_server

import (
	"context"
	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"

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

func (t *tachibanaApi) newOrder(ctx context.Context, session *tachibana.Session, req *pb.NewOrderRequest) (*pb.NewOrderResponse, error) {
	res, err := t.client.NewOrder(ctx, session, *t.toNewOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromNewOrderResponse(res), nil
}

func (t *tachibanaApi) cancelOrder(ctx context.Context, session *tachibana.Session, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	res, err := t.client.CancelOrder(ctx, session, *t.toCancelOrderRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromCancelOrderResponse(res), nil
}

func (t *tachibanaApi) orderList(ctx context.Context, session *tachibana.Session, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	res, err := t.client.OrderList(ctx, session, *t.toOrderListRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromOrderListResponse(res), nil
}

func (t *tachibanaApi) orderDetail(ctx context.Context, session *tachibana.Session, req *pb.OrderDetailRequest) (*pb.OrderDetailResponse, error) {
	res, err := t.client.OrderDetail(ctx, session, *t.toOrderDetailRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromOrderDetailResponse(res), nil
}

func (t *tachibanaApi) stockMaster(ctx context.Context, session *tachibana.Session, req *pb.StockMasterRequest) (*pb.StockMasterResponse, error) {
	res, err := t.client.StockMaster(ctx, session, *t.toStockMasterRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromStockMasterResponse(res), nil
}

func (t *tachibanaApi) stockExchangeMaster(ctx context.Context, session *tachibana.Session, req *pb.StockExchangeMasterRequest) (*pb.StockExchangeMasterResponse, error) {
	res, err := t.client.StockExchangeMaster(ctx, session, *t.toStockExchangeMasterRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromStockExchangeMasterResponse(res), nil
}

func (t *tachibanaApi) marketPrice(ctx context.Context, session *tachibana.Session, req *pb.MarketPriceRequest) (*pb.MarketPriceResponse, error) {
	res, err := t.client.MarketPrice(ctx, session, *t.toMarketPriceRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromMarketPriceResponse(res), nil
}

func (t *tachibanaApi) businessDay(ctx context.Context, session *tachibana.Session, req *pb.BusinessDayRequest) (*pb.BusinessDayResponse, error) {
	res, err := t.client.BusinessDay(ctx, session, *t.toBusinessDayRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromBusinessDayResponse(res), nil
}

func (t *tachibanaApi) tickGroup(ctx context.Context, session *tachibana.Session, req *pb.TickGroupRequest) (*pb.TickGroupResponse, error) {
	res, err := t.client.TickGroup(ctx, session, *t.toTickGroupRequest(req))
	if err != nil {
		return nil, err
	}
	return t.fromTickGroupResponse(res), nil
}
