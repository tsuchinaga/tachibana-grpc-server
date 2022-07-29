package tachibana_grpc_server

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc/metadata"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

func Test_server_Login(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name               string
		clock              *testClock
		sessionStore       *testSessionStore
		tachibanaApi       *testTachibanaApi
		arg1               context.Context
		arg2               *pb.LoginRequest
		want1              *pb.LoginResponse
		want2              error
		wantSave           []interface{}
		wantAddClientToken []interface{}
	}{
		{name: "clientTokenでsessionStoreから取得出来たら、取得した情報を返す",
			clock:              &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore:       &testSessionStore{getByClientToken1: &accountSession{BaseResponse: &pb.LoginResponse{}}},
			tachibanaApi:       &testTachibanaApi{},
			arg1:               context.Background(),
			arg2:               &pb.LoginRequest{UserId: "user-id", Password: "password", ClientName: "client-name"},
			want1:              &pb.LoginResponse{Token: "fba13b4714c61a91e9be7a2844d281c43bb889cca7b5a432dffb1ef935e4b26a"},
			want2:              nil,
			wantSave:           nil,
			wantAddClientToken: nil},
		{name: "clientTokenで取れなくても、sessionTokenでsessionStoreから取得出来たら、取得した情報を返す",
			clock:              &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore:       &testSessionStore{getBySessionToken1: &accountSession{BaseResponse: &pb.LoginResponse{}}},
			tachibanaApi:       &testTachibanaApi{},
			arg1:               context.Background(),
			arg2:               &pb.LoginRequest{UserId: "user-id", Password: "password", ClientName: "client-name"},
			want1:              &pb.LoginResponse{Token: "fba13b4714c61a91e9be7a2844d281c43bb889cca7b5a432dffb1ef935e4b26a"},
			want2:              nil,
			wantSave:           nil,
			wantAddClientToken: []interface{}{"df8dacf6e2a94335c0fc8b1a1a05601c72152f0f30905feffcb45f74fe20ad00", "fba13b4714c61a91e9be7a2844d281c43bb889cca7b5a432dffb1ef935e4b26a"}},
		{name: "ログイン処理でエラーが返されたらエラーを返す",
			clock:              &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore:       &testSessionStore{},
			tachibanaApi:       &testTachibanaApi{login2: unknownErr},
			arg1:               context.Background(),
			arg2:               &pb.LoginRequest{UserId: "user-id", Password: "password", ClientName: "client-name"},
			want1:              nil,
			want2:              unknownErr,
			wantSave:           nil,
			wantAddClientToken: nil},
		{name: "ログイン処理でエラーでなくても、ログインに失敗していたら保存せず結果を返す",
			clock:              &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore:       &testSessionStore{},
			tachibanaApi:       &testTachibanaApi{login1: &accountSession{Session: nil, BaseResponse: &pb.LoginResponse{ResultCode: "0"}}},
			arg1:               context.Background(),
			arg2:               &pb.LoginRequest{UserId: "user-id", Password: "password", ClientName: "client-name"},
			want1:              &pb.LoginResponse{ResultCode: "0", Token: ""},
			want2:              nil,
			wantSave:           nil,
			wantAddClientToken: nil},
		{name: "ログイン処理で成功していたらsessionを保存して結果を返す",
			clock:        &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{login1: &accountSession{Session: &tachibana.Session{}, BaseResponse: &pb.LoginResponse{ResultCode: "0"}}},
			arg1:         context.Background(),
			arg2:         &pb.LoginRequest{UserId: "user-id", Password: "password", ClientName: "client-name"},
			want1:        &pb.LoginResponse{ResultCode: "0", Token: "fba13b4714c61a91e9be7a2844d281c43bb889cca7b5a432dffb1ef935e4b26a"},
			want2:        nil,
			wantSave: []interface{}{
				"df8dacf6e2a94335c0fc8b1a1a05601c72152f0f30905feffcb45f74fe20ad00",
				"fba13b4714c61a91e9be7a2844d281c43bb889cca7b5a432dffb1ef935e4b26a",
				&accountSession{Date: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local), Token: "df8dacf6e2a94335c0fc8b1a1a05601c72152f0f30905feffcb45f74fe20ad00", Session: &tachibana.Session{}, BaseResponse: &pb.LoginResponse{ResultCode: "0", Token: ""}},
			},
			wantAddClientToken: nil},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{clock: test.clock, sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.Login(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) ||
				!errors.Is(got2, test.want2) ||
				!reflect.DeepEqual(test.wantSave, test.sessionStore.saveHistory) ||
				!reflect.DeepEqual(test.wantAddClientToken, test.sessionStore.addClientTokenHistory) {
				_wantSave, _ := json.Marshal(test.wantSave)
				_saveHistory, _ := json.Marshal(test.sessionStore.saveHistory)

				t.Errorf("%s error\nresult: %+v, %+v, %+v, %+v\nwant: %+v, %+v, %+v, %+v\ngot: %+v, %+v, %+v, %+v\n", t.Name(),
					!reflect.DeepEqual(test.want1, got1),
					!errors.Is(got2, test.want2),
					!reflect.DeepEqual(test.wantSave, test.sessionStore.saveHistory),
					!reflect.DeepEqual(test.wantAddClientToken, test.sessionStore.addClientTokenHistory),
					test.want1, test.want2, string(_wantSave), test.wantAddClientToken,
					got1, got2, string(_saveHistory), test.sessionStore.addClientTokenHistory)
			}
		})
	}
}

func Test_server_getClientToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		arg1         context.Context
		want1        string
		want2        bool
	}{
		{name: "metadataのないcontextを渡したらfalse",
			sessionStore: &testSessionStore{},
			arg1:         context.Background(),
			want1:        "",
			want2:        false},
		{name: "Session-tokenがメタデータになければfalse",
			sessionStore: &testSessionStore{},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("foo", "token001")),
			want1:        "",
			want2:        false},
		{name: "session-tokenがあれば文字列を返す",
			sessionStore: &testSessionStore{getBySessionToken1: &accountSession{BaseResponse: &pb.LoginResponse{Token: "token001"}}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			want1:        "token001",
			want2:        true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore}
			got1, got2 := server.getClientToken(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_server_getSession(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		arg1         context.Context
		want1        *accountSession
		hasError     bool
	}{
		{name: "metadataのないcontextを渡したらエラー",
			sessionStore: &testSessionStore{},
			arg1:         context.Background(),
			want1:        nil,
			hasError:     true},
		{name: "session-tokenがメタデータになければエラー",
			sessionStore: &testSessionStore{},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("foo", "token001")),
			want1:        nil,
			hasError:     true},
		{name: "session-tokenがあっても、storeにデータがなければエラー",
			sessionStore: &testSessionStore{getByClientToken1: nil},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			want1:        nil,
			hasError:     true},
		{name: "session-tokenがあっても、storeにデータがあればそれを返す",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{BaseResponse: &pb.LoginResponse{Token: "token001"}}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			want1:        &accountSession{BaseResponse: &pb.LoginResponse{Token: "token001"}},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore}
			got1, got2 := server.getSession(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_NewOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.NewOrderRequest
		want1        *pb.NewOrderResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.NewOrderRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{newOrder2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.NewOrderRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{newOrder1: &pb.NewOrderResponse{OrderNumber: "order-number001"}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.NewOrderRequest{},
			want1:        &pb.NewOrderResponse{OrderNumber: "order-number001"},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.NewOrder(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_CancelOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.CancelOrderRequest
		want1        *pb.CancelOrderResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.CancelOrderRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{cancelOrder2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.CancelOrderRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{cancelOrder1: &pb.CancelOrderResponse{OrderNumber: "order-number001"}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.CancelOrderRequest{},
			want1:        &pb.CancelOrderResponse{OrderNumber: "order-number001"},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.CancelOrder(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_OrderList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.OrderListRequest
		want1        *pb.OrderListResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.OrderListRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{orderList2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.OrderListRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{orderList1: &pb.OrderListResponse{ResultCode: "0"}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.OrderListRequest{},
			want1:        &pb.OrderListResponse{ResultCode: "0"},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.OrderList(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_OrderDetail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.OrderDetailRequest
		want1        *pb.OrderDetailResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.OrderDetailRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{orderDetail2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.OrderDetailRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{orderDetail1: &pb.OrderDetailResponse{ResultCode: "0"}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.OrderDetailRequest{},
			want1:        &pb.OrderDetailResponse{ResultCode: "0"},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.OrderDetail(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_StockMaster(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.StockMasterRequest
		want1        *pb.StockMasterResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.StockMasterRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{stockMaster2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.StockMasterRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{stockMaster1: &pb.StockMasterResponse{}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.StockMasterRequest{},
			want1:        &pb.StockMasterResponse{},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.StockMaster(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_StockExchangeMaster(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.StockExchangeMasterRequest
		want1        *pb.StockExchangeMasterResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.StockExchangeMasterRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{stockExchangeMaster2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.StockExchangeMasterRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{stockExchangeMaster1: &pb.StockExchangeMasterResponse{}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.StockExchangeMasterRequest{},
			want1:        &pb.StockExchangeMasterResponse{},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.StockExchangeMaster(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_MarketPrice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.MarketPriceRequest
		want1        *pb.MarketPriceResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.MarketPriceRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{marketPrice2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.MarketPriceRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{marketPrice1: &pb.MarketPriceResponse{}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.MarketPriceRequest{},
			want1:        &pb.MarketPriceResponse{},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.MarketPrice(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_BusinessDay(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.BusinessDayRequest
		want1        *pb.BusinessDayResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.BusinessDayRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{businessDay2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.BusinessDayRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{businessDay1: &pb.BusinessDayResponse{}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.BusinessDayRequest{},
			want1:        &pb.BusinessDayResponse{},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.BusinessDay(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_TickGroup(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.TickGroupRequest
		want1        *pb.TickGroupResponse
		hasError     bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore: &testSessionStore{},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.TickGroupRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社からエラーがあったらエラーを返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{tickGroup2: unknownErr},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.TickGroupRequest{},
			want1:        nil,
			hasError:     true},
		{name: "証券会社のレスポンスが返される",
			sessionStore: &testSessionStore{getByClientToken1: &accountSession{}},
			tachibanaApi: &testTachibanaApi{tickGroup1: &pb.TickGroupResponse{}},
			arg1:         metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001")),
			arg2:         &pb.TickGroupRequest{},
			want1:        &pb.TickGroupResponse{},
			hasError:     false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.TickGroup(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || (got2 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.want1, test.hasError, got1, got2)
			}
		})
	}
}

func Test_server_Stream(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		sessionStore  *testSessionStore
		streamService *testStreamService
		arg1          *pb.StreamRequest
		arg2          pb.TachibanaService_StreamServer
		hasError      bool
	}{
		{name: "ログインに失敗したらエラーを返される",
			sessionStore:  &testSessionStore{},
			streamService: &testStreamService{},
			arg1:          &pb.StreamRequest{},
			arg2:          &testStreamServer{context1: context.Background()},
			hasError:      true},
		{name: "streamでエラーが返されたらエラー",
			sessionStore:  &testSessionStore{getByClientToken1: &accountSession{}},
			streamService: &testStreamService{connect1: unknownErr},
			arg1:          &pb.StreamRequest{},
			arg2:          &testStreamServer{context1: metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001"))},
			hasError:      true},
		{name: "streamから返されたエラーがnilならnilをそのまま返す",
			sessionStore:  &testSessionStore{getByClientToken1: &accountSession{}},
			streamService: &testStreamService{connect1: nil},
			arg1:          &pb.StreamRequest{},
			arg2:          &testStreamServer{context1: metadata.NewIncomingContext(context.Background(), metadata.Pairs("session-token", "token001"))},
			hasError:      false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{sessionStore: test.sessionStore, streamService: test.streamService}
			got1 := server.Stream(test.arg1, test.arg2)
			if (got1 != nil) != test.hasError {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.hasError, got1)
			}
		})
	}
}
