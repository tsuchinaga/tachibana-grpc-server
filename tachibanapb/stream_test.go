package tachibanapb

import (
	"reflect"
	"testing"
)

func Test_StreamRequest_Sendable(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		req   *StreamRequest
		arg1  *StreamResponse
		want1 bool
	}{
		{name: "再送を受けない場合に、初回配信出ないイベントは通知されない",
			req:   &StreamRequest{EventTypes: []EventType{EventType_EVENT_TYPE_CONTRACT, EventType_EVENT_TYPE_ERROR_STATUS}, ReceiveResend: false},
			arg1:  &StreamResponse{EventType: EventType_EVENT_TYPE_CONTRACT, IsFirstTime: false},
			want1: false},
		{name: "reqに含まれないEventならfalse",
			req:   &StreamRequest{EventTypes: []EventType{EventType_EVENT_TYPE_CONTRACT, EventType_EVENT_TYPE_ERROR_STATUS}, ReceiveResend: true},
			arg1:  &StreamResponse{EventType: EventType_EVENT_TYPE_MARKET_PRICE},
			want1: false},
		{name: "reqに含まれるEventならtrue",
			req:   &StreamRequest{EventTypes: []EventType{EventType_EVENT_TYPE_CONTRACT, EventType_EVENT_TYPE_ERROR_STATUS}, ReceiveResend: true},
			arg1:  &StreamResponse{EventType: EventType_EVENT_TYPE_CONTRACT},
			want1: true},
		{name: "reqに含まれるEventでも、MarketPriceなら銘柄一覧に含まれていなければfalse",
			req: &StreamRequest{
				EventTypes:    []EventType{EventType_EVENT_TYPE_MARKET_PRICE, EventType_EVENT_TYPE_ERROR_STATUS},
				ReceiveResend: true,
				StreamIssues: []*StreamIssue{
					{IssueCode: "1111", Exchange: Exchange_EXCHANGE_MEISHOU},
					{IssueCode: "2222", Exchange: Exchange_EXCHANGE_TOUSHOU},
					{IssueCode: "3333", Exchange: Exchange_EXCHANGE_TOUSHOU}}},
			arg1: &StreamResponse{EventType: EventType_EVENT_TYPE_MARKET_PRICE,
				MarketPriceStreamResponse: &MarketPriceStreamResponse{IssueCode: "1111", Exchange: Exchange_EXCHANGE_TOUSHOU}},
			want1: false},
		{name: "reqに含まれるEventでかつMarketPriceであり、銘柄一覧に含まれてればtrue",
			req: &StreamRequest{
				EventTypes:    []EventType{EventType_EVENT_TYPE_MARKET_PRICE, EventType_EVENT_TYPE_ERROR_STATUS},
				ReceiveResend: true,
				StreamIssues: []*StreamIssue{
					{IssueCode: "1111", Exchange: Exchange_EXCHANGE_MEISHOU},
					{IssueCode: "2222", Exchange: Exchange_EXCHANGE_TOUSHOU},
					{IssueCode: "3333", Exchange: Exchange_EXCHANGE_TOUSHOU}}},
			arg1: &StreamResponse{EventType: EventType_EVENT_TYPE_MARKET_PRICE,
				MarketPriceStreamResponse: &MarketPriceStreamResponse{IssueCode: "2222", Exchange: Exchange_EXCHANGE_TOUSHOU}},
			want1: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.req.Sendable(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_StreamRequest_Union(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		req   *StreamRequest
		arg1  *StreamRequest
		want1 *StreamRequest
	}{
		{name: "元の配列がnilなら空配列にして返す",
			req:   &StreamRequest{},
			arg1:  &StreamRequest{},
			want1: &StreamRequest{EventTypes: []EventType{}, StreamIssues: []*StreamIssue{}}},
		{name: "元の配列がnilでも新しいリクエストがあればその値が入る",
			req: &StreamRequest{},
			arg1: &StreamRequest{
				EventTypes:   []EventType{EventType_EVENT_TYPE_CONTRACT, EventType_EVENT_TYPE_ERROR_STATUS},
				StreamIssues: []*StreamIssue{{IssueCode: "1475", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1476", Exchange: Exchange_EXCHANGE_TOUSHOU}}},
			want1: &StreamRequest{EventTypes: []EventType{EventType_EVENT_TYPE_CONTRACT, EventType_EVENT_TYPE_ERROR_STATUS},
				StreamIssues: []*StreamIssue{{IssueCode: "1475", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1476", Exchange: Exchange_EXCHANGE_TOUSHOU}}}},
		{name: "2つのリクエストから重複しないリクエストを作る",
			req: &StreamRequest{
				EventTypes:   []EventType{EventType_EVENT_TYPE_CONTRACT, EventType_EVENT_TYPE_ERROR_STATUS},
				StreamIssues: []*StreamIssue{{IssueCode: "1475", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1476", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1478", Exchange: Exchange_EXCHANGE_TOUSHOU}}},
			arg1: &StreamRequest{
				EventTypes:   []EventType{EventType_EVENT_TYPE_CONTRACT, EventType_EVENT_TYPE_OPERATION_STATUS},
				StreamIssues: []*StreamIssue{{IssueCode: "1475", Exchange: Exchange_EXCHANGE_MEISHOU}, {IssueCode: "1477", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1478", Exchange: Exchange_EXCHANGE_TOUSHOU}}},
			want1: &StreamRequest{EventTypes: []EventType{EventType_EVENT_TYPE_CONTRACT, EventType_EVENT_TYPE_ERROR_STATUS, EventType_EVENT_TYPE_OPERATION_STATUS},
				StreamIssues: []*StreamIssue{{IssueCode: "1475", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1476", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1478", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1475", Exchange: Exchange_EXCHANGE_MEISHOU}, {IssueCode: "1477", Exchange: Exchange_EXCHANGE_TOUSHOU}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.req.Union(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_StreamRequest_hasIssue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		req   *StreamRequest
		arg1  string
		arg2  Exchange
		want1 bool
	}{
		{name: "issuesがnilならfalse",
			req:   &StreamRequest{StreamIssues: nil},
			arg1:  "1111",
			arg2:  Exchange_EXCHANGE_TOUSHOU,
			want1: false},
		{name: "issuesに引数の銘柄がなければfalse",
			req:   &StreamRequest{StreamIssues: []*StreamIssue{{IssueCode: "2222", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1111", Exchange: Exchange_EXCHANGE_MEISHOU}}},
			arg1:  "1111",
			arg2:  Exchange_EXCHANGE_TOUSHOU,
			want1: false},
		{name: "issuesに引数の銘柄があればtrue",
			req:   &StreamRequest{StreamIssues: []*StreamIssue{{IssueCode: "2222", Exchange: Exchange_EXCHANGE_TOUSHOU}, {IssueCode: "1111", Exchange: Exchange_EXCHANGE_TOUSHOU}}},
			arg1:  "1111",
			arg2:  Exchange_EXCHANGE_TOUSHOU,
			want1: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.req.hasIssue(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}
