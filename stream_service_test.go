package tachibana_grpc_server

import (
	"context"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
	"reflect"
	"testing"
	"time"
)

type testStreamServer struct {
	send1       error
	sendHistory []interface{}
	context1    context.Context
	pb.TachibanaService_StreamServer
}

func (t *testStreamServer) Send(response *pb.StreamResponse) error {
	t.sendHistory = append(t.sendHistory, response)
	return t.send1
}

func (t *testStreamServer) Context() context.Context {
	return t.context1
}

type testStreamService struct {
	iStreamService
	connect1       error
	connectCount   int
	connectHistory []interface{}
	clearCount     int
}

func (t *testStreamService) connect(ctx context.Context, session *accountSession, clientToken string, req *pb.StreamRequest, stream pb.TachibanaService_StreamServer) error {
	t.connectCount++
	t.connectHistory = append(t.connectHistory, ctx, session, clientToken, req, stream)
	return t.connect1
}
func (t *testStreamService) clear() {
	t.clearCount++
}

type testClientStream struct {
	iClientStream
	requestCount      int
	request1          *pb.StreamRequest
	startCount        int
	sendCount         int
	sendHistory       []interface{}
	disconnectCount   int
	disconnectHistory []interface{}
}

func (t *testClientStream) getRequest() *pb.StreamRequest {
	t.requestCount++
	return t.request1
}
func (t *testClientStream) start() {
	t.startCount++
}
func (t *testClientStream) send(res *pb.StreamResponse) {
	t.sendCount++
	t.sendHistory = append(t.sendHistory, res)
}
func (t *testClientStream) disconnect(err error) {
	t.disconnectCount++
	t.disconnectHistory = append(t.disconnectHistory, err)
}

func Test_sessionStream_addClient(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	errCh := make(chan error)
	tests := []struct {
		name                string
		oldClientStream     iClientStream
		arg1                context.Context
		arg2                string
		arg3                *pb.StreamRequest
		arg4                pb.TachibanaService_StreamServer
		arg5                chan error
		want1               *pb.StreamRequest
		wantOldClientStream iClientStream
		wantStreams         map[string]iClientStream
	}{
		{name: "すでにクライアント接続があったら切断を叩く",
			oldClientStream: &testClientStream{},
			arg1:            ctx,
			arg2:            "client-token",
			arg3:            &pb.StreamRequest{EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_ERROR_STATUS}},
			arg4:            &testStreamServer{},
			arg5:            errCh,
			want1: &pb.StreamRequest{
				EventTypes:   []pb.EventType{pb.EventType_EVENT_TYPE_ERROR_STATUS, pb.EventType_EVENT_TYPE_KEEPALIVE, pb.EventType_EVENT_TYPE_CONTRACT},
				StreamIssues: []*pb.StreamIssue{},
			},
			wantOldClientStream: &testClientStream{disconnectCount: 1, disconnectHistory: []interface{}{newStreamRequestErr}},
			wantStreams: map[string]iClientStream{"client-token": &clientStream{
				clientToken: "client-token",
				request: &pb.StreamRequest{
					EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_ERROR_STATUS}},
				stream:      &testStreamServer{},
				ctx:         ctx,
				errCh:       errCh,
				isConnected: true,
			}}},
		{name: "クライアント接続がなければ切断は叩かない",
			arg1: ctx,
			arg2: "client-token",
			arg3: &pb.StreamRequest{EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_ERROR_STATUS}},
			arg4: &testStreamServer{},
			arg5: errCh,
			want1: &pb.StreamRequest{
				EventTypes:    []pb.EventType{pb.EventType_EVENT_TYPE_ERROR_STATUS, pb.EventType_EVENT_TYPE_KEEPALIVE, pb.EventType_EVENT_TYPE_CONTRACT},
				ReceiveResend: false,
				StreamIssues:  []*pb.StreamIssue{},
			},
			wantOldClientStream: nil,
			wantStreams: map[string]iClientStream{"client-token": &clientStream{
				clientToken: "client-token",
				request:     &pb.StreamRequest{EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_ERROR_STATUS}},
				stream:      &testStreamServer{},
				ctx:         ctx,
				errCh:       errCh,
				isConnected: true,
			}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			sessionStream := &sessionStream{}
			if test.oldClientStream != nil {
				sessionStream.streams = map[string]iClientStream{"client-token": test.oldClientStream}
			}
			got1 := sessionStream.addClient(test.arg1, test.arg2, test.arg3, test.arg4, test.arg5)
			if !reflect.DeepEqual(test.want1, got1) ||
				!reflect.DeepEqual(test.wantOldClientStream, test.oldClientStream) ||
				!reflect.DeepEqual(test.wantStreams, sessionStream.streams) {
				t.Errorf("%s error\nresult: %+v, %+v, %+v\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(),
					!reflect.DeepEqual(test.want1, got1),
					!reflect.DeepEqual(test.wantOldClientStream, test.oldClientStream),
					!reflect.DeepEqual(test.wantStreams, sessionStream.streams),
					test.want1, test.wantOldClientStream, test.wantStreams,
					got1, test.oldClientStream, sessionStream.streams)
			}
		})
	}
}

func Test_sessionStream_getRequest(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		sessionStream *sessionStream
		want1         *pb.StreamRequest
	}{
		{name: "クライアントのリクエストが空っぽなら空っぽのリクエストが出される",
			sessionStream: &sessionStream{},
			want1:         &pb.StreamRequest{EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_KEEPALIVE}, StreamIssues: []*pb.StreamIssue{}}},
		{name: "クライアントのリクエストが単一のリクエストなら、同じリクエストが返される",
			sessionStream: &sessionStream{
				streams: map[string]iClientStream{"client-token": &clientStream{
					request: &pb.StreamRequest{
						EventTypes:   []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_ERROR_STATUS},
						StreamIssues: []*pb.StreamIssue{{IssueCode: "1475", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1476", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}},
					},
				}},
			},
			want1: &pb.StreamRequest{
				EventTypes:   []pb.EventType{pb.EventType_EVENT_TYPE_ERROR_STATUS, pb.EventType_EVENT_TYPE_KEEPALIVE, pb.EventType_EVENT_TYPE_CONTRACT},
				StreamIssues: []*pb.StreamIssue{{IssueCode: "1475", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1476", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}},
			}},
		{name: "クライアントのリクエストが複数あるなら、結合したリクエストを作る",
			sessionStream: &sessionStream{
				streams: map[string]iClientStream{
					"client-token001": &clientStream{
						request: &pb.StreamRequest{
							EventTypes:   []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_SYSTEM_STATUS, pb.EventType_EVENT_TYPE_MARKET_PRICE},
							StreamIssues: []*pb.StreamIssue{{IssueCode: "1475", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1476", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}},
						},
					},
					"client-token002": &clientStream{
						request: &pb.StreamRequest{
							EventTypes:   []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_ERROR_STATUS, pb.EventType_EVENT_TYPE_MARKET_PRICE},
							StreamIssues: []*pb.StreamIssue{{IssueCode: "1475", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1477", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}},
						},
					}},
			},
			want1: &pb.StreamRequest{
				EventTypes:   []pb.EventType{pb.EventType_EVENT_TYPE_ERROR_STATUS, pb.EventType_EVENT_TYPE_KEEPALIVE, pb.EventType_EVENT_TYPE_MARKET_PRICE, pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_SYSTEM_STATUS},
				StreamIssues: []*pb.StreamIssue{{IssueCode: "1475", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1476", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1477", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.sessionStream.getRequest()
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_sessionStream_send(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		sessionStream *sessionStream
		arg1          *pb.StreamResponse
		wantStreams   map[string]iClientStream
	}{
		{name: "すべてのクライアントに通知する",
			sessionStream: &sessionStream{streams: map[string]iClientStream{
				"client-token001": &testClientStream{},
				"client-token002": &testClientStream{},
				"client-token003": &testClientStream{},
			}},
			arg1: &pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_CONTRACT},
			wantStreams: map[string]iClientStream{
				"client-token001": &testClientStream{sendCount: 1, sendHistory: []interface{}{&pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_CONTRACT}}},
				"client-token002": &testClientStream{sendCount: 1, sendHistory: []interface{}{&pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_CONTRACT}}},
				"client-token003": &testClientStream{sendCount: 1, sendHistory: []interface{}{&pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_CONTRACT}}},
			}},
		{name: "時価情報なら、銘柄情報を追加する",
			sessionStream: &sessionStream{
				request: &pb.StreamRequest{StreamIssues: []*pb.StreamIssue{
					{IssueCode: "1111", Exchange: pb.Exchange_EXCHANGE_TOUSHOU},
					{IssueCode: "2222", Exchange: pb.Exchange_EXCHANGE_TOUSHOU},
					{IssueCode: "3333", Exchange: pb.Exchange_EXCHANGE_TOUSHOU},
				}},
				streams: map[string]iClientStream{
					"client-token001": &testClientStream{},
					"client-token002": &testClientStream{},
					"client-token003": &testClientStream{},
				}},
			arg1: &pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, MarketPriceStreamResponse: &pb.MarketPriceStreamResponse{ColumnNumber: 2}},
			wantStreams: map[string]iClientStream{
				"client-token001": &testClientStream{sendCount: 1, sendHistory: []interface{}{&pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, MarketPriceStreamResponse: &pb.MarketPriceStreamResponse{ColumnNumber: 2, IssueCode: "2222", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}}}},
				"client-token002": &testClientStream{sendCount: 1, sendHistory: []interface{}{&pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, MarketPriceStreamResponse: &pb.MarketPriceStreamResponse{ColumnNumber: 2, IssueCode: "2222", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}}}},
				"client-token003": &testClientStream{sendCount: 1, sendHistory: []interface{}{&pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, MarketPriceStreamResponse: &pb.MarketPriceStreamResponse{ColumnNumber: 2, IssueCode: "2222", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}}}},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.sessionStream.send(test.arg1)
			time.Sleep(time.Second) // 非同期なので少し待機
			if !reflect.DeepEqual(test.wantStreams, test.sessionStream.streams) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.wantStreams, test.sessionStream.streams)
			}
		})
	}
}

func Test_sessionStream_sendError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		sessionStream *sessionStream
		arg1          error
		wantStreams   map[string]iClientStream
	}{
		{name: "すべてのクライアントに通知する",
			sessionStream: &sessionStream{streams: map[string]iClientStream{
				"client-token001": &testClientStream{},
				"client-token002": &testClientStream{},
				"client-token003": &testClientStream{},
			}},
			arg1: newStreamRequestErr,
			wantStreams: map[string]iClientStream{
				"client-token001": &testClientStream{disconnectCount: 1, disconnectHistory: []interface{}{newStreamRequestErr}},
				"client-token002": &testClientStream{disconnectCount: 1, disconnectHistory: []interface{}{newStreamRequestErr}},
				"client-token003": &testClientStream{disconnectCount: 1, disconnectHistory: []interface{}{newStreamRequestErr}},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.sessionStream.sendError(test.arg1)
			time.Sleep(time.Second) // 非同期なので少し待機
			if !reflect.DeepEqual(test.wantStreams, test.sessionStream.streams) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.wantStreams, test.sessionStream.streams)
			}
		})
	}
}

func Test_sessionStream_disconnect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		sessionStream *sessionStream
		wantIsConnect bool
	}{
		{name: "未接続なら何もしない",
			sessionStream: &sessionStream{isConnected: false},
			wantIsConnect: false},
		{name: "接続済みならcontextのキャンセルを叩いて未接続に戻す",
			sessionStream: &sessionStream{isConnected: true},
			wantIsConnect: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctx, cf := context.WithCancel(context.Background())
			test.sessionStream.ctx = ctx
			test.sessionStream.cf = cf
			test.sessionStream.disconnect()
			if !reflect.DeepEqual(test.wantIsConnect, test.sessionStream.isConnected) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.wantIsConnect, test.sessionStream.isConnected)
			}
		})
	}
}

func Test_clientStream_request(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		clientStream *clientStream
		want1        *pb.StreamRequest
	}{
		{name: "requestの内容が返される",
			clientStream: &clientStream{request: &pb.StreamRequest{EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT}}},
			want1:        &pb.StreamRequest{EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT}}},
		{name: "requestがnilでも返す",
			clientStream: &clientStream{request: nil},
			want1:        nil},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.clientStream.getRequest()
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_clientStream_send(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		clientStream *clientStream
		arg1         *pb.StreamResponse
		wantStream   *testStreamServer
		wantErrCount int
	}{
		{name: "streamが接続されていなければ何もしない",
			clientStream: &clientStream{
				isConnected: false,
				request:     &pb.StreamRequest{EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_MARKET_PRICE}},
				stream:      &testStreamServer{}},
			arg1:         &pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, IsFirstTime: true},
			wantStream:   &testStreamServer{},
			wantErrCount: 0},
		{name: "リクエストとレスポンスが送信不要な関係であれば何も送らない",
			clientStream: &clientStream{
				isConnected: true,
				request:     &pb.StreamRequest{EventTypes: []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_SYSTEM_STATUS}},
				stream:      &testStreamServer{}},
			arg1:         &pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, IsFirstTime: true},
			wantStream:   &testStreamServer{},
			wantErrCount: 0},
		{name: "リクエストとレスポンスが送信可能な関係なら送信する",
			clientStream: &clientStream{
				isConnected: true,
				request: &pb.StreamRequest{
					EventTypes:   []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_MARKET_PRICE},
					StreamIssues: []*pb.StreamIssue{{IssueCode: "1111", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}}},
				stream: &testStreamServer{}},
			arg1:         &pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, IsFirstTime: true, MarketPriceStreamResponse: &pb.MarketPriceStreamResponse{IssueCode: "1111", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}},
			wantStream:   &testStreamServer{sendHistory: []interface{}{&pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, IsFirstTime: true, MarketPriceStreamResponse: &pb.MarketPriceStreamResponse{IssueCode: "1111", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}}}},
			wantErrCount: 0},
		{name: "リクエストとレスポンスが送信可能でもエラーになったら切断する",
			clientStream: &clientStream{
				isConnected: true,
				request: &pb.StreamRequest{
					EventTypes:   []pb.EventType{pb.EventType_EVENT_TYPE_CONTRACT, pb.EventType_EVENT_TYPE_MARKET_PRICE},
					StreamIssues: []*pb.StreamIssue{{IssueCode: "1111", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}}},
				stream: &testStreamServer{send1: unknownErr}},
			arg1:         &pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, IsFirstTime: true, MarketPriceStreamResponse: &pb.MarketPriceStreamResponse{IssueCode: "1111", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}},
			wantStream:   &testStreamServer{send1: unknownErr, sendHistory: []interface{}{&pb.StreamResponse{EventType: pb.EventType_EVENT_TYPE_MARKET_PRICE, IsFirstTime: true, MarketPriceStreamResponse: &pb.MarketPriceStreamResponse{IssueCode: "1111", Exchange: pb.Exchange_EXCHANGE_TOUSHOU}}}},
			wantErrCount: 1},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ch := make(chan error)
			var errCnt int
			go func() {
				for {
					<-ch
					errCnt++
				}
			}()
			test.clientStream.errCh = ch

			test.clientStream.send(test.arg1)
			if !reflect.DeepEqual(test.wantStream, test.clientStream.stream) ||
				!reflect.DeepEqual(test.wantErrCount, errCnt) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(), test.wantStream, test.wantErrCount, test.clientStream.stream, errCnt)
			}
		})
	}
}

func Test_clientStream_disconnect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		clientStream     *clientStream
		arg1             error
		wantErrCount     int
		wantClientStream *clientStream
	}{
		{name: "接続されていなければ何もしない",
			clientStream:     &clientStream{isConnected: false},
			arg1:             unknownErr,
			wantErrCount:     0,
			wantClientStream: &clientStream{isConnected: false}},
		{name: "接続されていれば、未接続に戻してエラーを送信する",
			clientStream:     &clientStream{isConnected: true},
			arg1:             unknownErr,
			wantErrCount:     1,
			wantClientStream: &clientStream{isConnected: false}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ch := make(chan error)
			var errCnt int
			go func() {
				for {
					<-ch
					errCnt++
				}
			}()
			test.clientStream.errCh = ch
			test.wantClientStream.errCh = ch

			test.clientStream.disconnect(test.arg1)

			if !reflect.DeepEqual(test.wantClientStream, test.clientStream) ||
				!reflect.DeepEqual(test.wantErrCount, errCnt) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(),
					test.wantClientStream, test.wantErrCount,
					test.clientStream, errCnt)
			}
		})
	}
}

func Test_streamService_clear(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                      string
		sessionStream             map[string]*sessionStream
		wantOriginalSessionStream map[string]*sessionStream
		wantAfterSessionStreams   map[string]*sessionStream
	}{
		{name: "server, clientのsessionがなければ何もしない",
			sessionStream:             map[string]*sessionStream{},
			wantOriginalSessionStream: map[string]*sessionStream{},
			wantAfterSessionStreams:   map[string]*sessionStream{}},
		{name: "server, clientのsessionがあれば切断してから消し込む",
			sessionStream: map[string]*sessionStream{
				"session1": {
					sessionToken: "session1",
					request:      nil,
					streams:      map[string]iClientStream{"client1": &testClientStream{}, "client2": &testClientStream{}},
					isConnected:  false,
					resCh:        nil,
					errCh:        nil,
					ctx:          nil,
					cf:           nil,
				},
				"session2": {
					sessionToken: "session2",
					request:      nil,
					streams:      map[string]iClientStream{"client1": &testClientStream{}, "client2": &testClientStream{}},
					isConnected:  false,
					resCh:        nil,
					errCh:        nil,
					ctx:          nil,
					cf:           nil,
				},
			},
			wantOriginalSessionStream: map[string]*sessionStream{
				"session1": {
					sessionToken: "session1",
					request:      nil,
					streams:      map[string]iClientStream{"client1": &testClientStream{disconnectCount: 1, disconnectHistory: []interface{}{stopStreamErr}}, "client2": &testClientStream{disconnectCount: 1, disconnectHistory: []interface{}{stopStreamErr}}},
					isConnected:  false,
					resCh:        nil,
					errCh:        nil,
					ctx:          nil,
					cf:           nil,
				},
				"session2": {
					sessionToken: "session2",
					request:      nil,
					streams:      map[string]iClientStream{"client1": &testClientStream{disconnectCount: 1, disconnectHistory: []interface{}{stopStreamErr}}, "client2": &testClientStream{disconnectCount: 1, disconnectHistory: []interface{}{stopStreamErr}}},
					isConnected:  false,
					resCh:        nil,
					errCh:        nil,
					ctx:          nil,
					cf:           nil,
				},
			},
			wantAfterSessionStreams: map[string]*sessionStream{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			streamService := &streamService{sessionStreams: test.sessionStream}
			streamService.clear()
			if !reflect.DeepEqual(test.wantOriginalSessionStream, test.sessionStream) || !reflect.DeepEqual(test.wantAfterSessionStreams, streamService.sessionStreams) {
				t.Errorf("%s error\nresult: %+v, %+v\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(),
					!reflect.DeepEqual(test.wantOriginalSessionStream, test.sessionStream),
					!reflect.DeepEqual(test.wantAfterSessionStreams, streamService.sessionStreams),
					test.wantOriginalSessionStream, test.wantAfterSessionStreams,
					test.sessionStream, streamService.sessionStreams)
			}
		})
	}
}
