package tachibana_grpc_server

import (
	"context"
	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type testTachibanaApi struct {
	iTachibanaApi
	login1                     *accountSession
	login2                     error
	loginHistory               []interface{}
	newOrder1                  *pb.NewOrderResponse
	newOrder2                  error
	newOrderHistory            []interface{}
	cancelOrder1               *pb.CancelOrderResponse
	cancelOrder2               error
	cancelOrderHistory         []interface{}
	orderList1                 *pb.OrderListResponse
	orderList2                 error
	orderListHistory           []interface{}
	orderDetail1               *pb.OrderDetailResponse
	orderDetail2               error
	orderDetailHistory         []interface{}
	stockMaster1               *pb.StockMasterResponse
	stockMaster2               error
	stockMasterHistory         []interface{}
	stockExchangeMaster1       *pb.StockExchangeMasterResponse
	stockExchangeMaster2       error
	stockExchangeMasterHistory []interface{}
	marketPrice1               *pb.MarketPriceResponse
	marketPrice2               error
	marketPriceHistory         []interface{}
	businessDay1               *pb.BusinessDayResponse
	businessDay2               error
	businessDayHistory         []interface{}
	tickGroup1                 *pb.TickGroupResponse
	tickGroup2                 error
	tickGroupHistory           []interface{}
}

func (t *testTachibanaApi) login(_ context.Context, req *pb.LoginRequest) (*accountSession, error) {
	t.loginHistory = append(t.loginHistory, req)
	return t.login1, t.login2
}
func (t *testTachibanaApi) newOrder(_ context.Context, session *tachibana.Session, req *pb.NewOrderRequest) (*pb.NewOrderResponse, error) {
	t.newOrderHistory = append(t.newOrderHistory, session, req)
	return t.newOrder1, t.newOrder2
}
func (t *testTachibanaApi) cancelOrder(_ context.Context, session *tachibana.Session, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	t.cancelOrderHistory = append(t.cancelOrderHistory, session, req)
	return t.cancelOrder1, t.cancelOrder2
}
func (t *testTachibanaApi) orderList(_ context.Context, session *tachibana.Session, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	t.orderListHistory = append(t.orderListHistory, session, req)
	return t.orderList1, t.orderList2
}
func (t *testTachibanaApi) orderDetail(_ context.Context, session *tachibana.Session, req *pb.OrderDetailRequest) (*pb.OrderDetailResponse, error) {
	t.orderDetailHistory = append(t.orderDetailHistory, session, req)
	return t.orderDetail1, t.orderDetail2
}
func (t *testTachibanaApi) stockMaster(_ context.Context, session *tachibana.Session, req *pb.StockMasterRequest) (*pb.StockMasterResponse, error) {
	t.stockMasterHistory = append(t.stockMasterHistory, session, req)
	return t.stockMaster1, t.stockMaster2
}
func (t *testTachibanaApi) stockExchangeMaster(_ context.Context, session *tachibana.Session, req *pb.StockExchangeMasterRequest) (*pb.StockExchangeMasterResponse, error) {
	t.stockExchangeMasterHistory = append(t.stockExchangeMasterHistory, session, req)
	return t.stockExchangeMaster1, t.stockExchangeMaster2
}
func (t *testTachibanaApi) marketPrice(_ context.Context, session *tachibana.Session, req *pb.MarketPriceRequest) (*pb.MarketPriceResponse, error) {
	t.marketPriceHistory = append(t.marketPriceHistory, session, req)
	return t.marketPrice1, t.marketPrice2
}
func (t *testTachibanaApi) businessDay(_ context.Context, session *tachibana.Session, req *pb.BusinessDayRequest) (*pb.BusinessDayResponse, error) {
	t.businessDayHistory = append(t.businessDayHistory, session, req)
	return t.businessDay1, t.businessDay2
}
func (t *testTachibanaApi) tickGroup(_ context.Context, session *tachibana.Session, req *pb.TickGroupRequest) (*pb.TickGroupResponse, error) {
	t.tickGroupHistory = append(t.tickGroupHistory, session, req)
	return t.tickGroup1, t.tickGroup2
}
