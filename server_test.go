package tachibana_grpc_server

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

func Test_server_Login(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		clock        *testClock
		sessionStore *testSessionStore
		tachibanaApi *testTachibanaApi
		arg1         context.Context
		arg2         *pb.LoginRequest
		want1        *pb.LoginResponse
		want2        error
		wantSave     []interface{}
	}{
		{name: "sessionStoreから取得出来たら、取得した情報を返す",
			clock:        &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore: &testSessionStore{getSession1: &loginSession{response: &pb.LoginResponse{Token: "token001"}}},
			tachibanaApi: &testTachibanaApi{},
			arg1:         context.Background(),
			arg2:         &pb.LoginRequest{UserId: "user-id", Password: "password"},
			want1:        &pb.LoginResponse{Token: "token001"},
			want2:        nil,
			wantSave:     nil},
		{name: "ログイン処理でエラーが返されたらエラーを返す",
			clock:        &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore: &testSessionStore{getSession1: nil},
			tachibanaApi: &testTachibanaApi{login2: unknownErr},
			arg1:         context.Background(),
			arg2:         &pb.LoginRequest{UserId: "user-id", Password: "password"},
			want1:        nil,
			want2:        unknownErr,
			wantSave:     nil},
		{name: "ログイン処理でエラーでなくても、ログインに失敗していたら保存せず結果を返す",
			clock:        &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore: &testSessionStore{getSession1: nil},
			tachibanaApi: &testTachibanaApi{login1: &loginSession{session: nil, response: &pb.LoginResponse{ResultCode: "0"}}},
			arg1:         context.Background(),
			arg2:         &pb.LoginRequest{UserId: "user-id", Password: "password"},
			want1:        &pb.LoginResponse{ResultCode: "0", Token: ""},
			want2:        nil,
			wantSave:     nil},
		{name: "ログイン処理で成功していたらsessionを保存して結果を返す",
			clock:        &testClock{today1: time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local)},
			sessionStore: &testSessionStore{getSession1: nil},
			tachibanaApi: &testTachibanaApi{login1: &loginSession{session: &tachibana.Session{}, response: &pb.LoginResponse{ResultCode: "0"}}},
			arg1:         context.Background(),
			arg2:         &pb.LoginRequest{UserId: "user-id", Password: "password"},
			want1:        &pb.LoginResponse{ResultCode: "0", Token: "df8dacf6e2a94335c0fc8b1a1a05601c72152f0f30905feffcb45f74fe20ad00"},
			want2:        nil,
			wantSave: []interface{}{
				"df8dacf6e2a94335c0fc8b1a1a05601c72152f0f30905feffcb45f74fe20ad00",
				&loginSession{session: &tachibana.Session{}, response: &pb.LoginResponse{ResultCode: "0", Token: "df8dacf6e2a94335c0fc8b1a1a05601c72152f0f30905feffcb45f74fe20ad00"}},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := &server{clock: test.clock, sessionStore: test.sessionStore, tachibana: test.tachibanaApi}
			got1, got2 := server.Login(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) || !errors.Is(got2, test.want2) || !reflect.DeepEqual(test.wantSave, test.sessionStore.saveHistory) {
				t.Errorf("%s error\nwant: %+v, %+v, %+v\ngot: %+v, %+v, %+v\n", t.Name(),
					test.want1, test.want2, test.wantSave,
					got1, got2, test.sessionStore.saveHistory)
			}
		})
	}
}
