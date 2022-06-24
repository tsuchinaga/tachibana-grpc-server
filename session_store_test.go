package tachibana_grpc_server

import (
	"reflect"
	"testing"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type testSessionStore struct {
	iSessionStore
	getSession1       *loginSession
	getSessionHistory []interface{}
	saveHistory       []interface{}
}

func (t *testSessionStore) getSession(key string) *loginSession {
	t.getSessionHistory = append(t.getSessionHistory, key)
	return t.getSession1
}
func (t *testSessionStore) save(key string, session *loginSession) {
	t.saveHistory = append(t.saveHistory, key, session)
}

func Test_sessionStore_getSession(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *sessionStore
		arg1  string
		want1 *loginSession
	}{
		{name: "storeに指定のキーのsessionがなければnilを返す",
			store: &sessionStore{store: map[string]*loginSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
			arg1:  "key000",
			want1: nil},
		{name: "storeに指定のキーのsessionがあればsessionを返す",
			store: &sessionStore{store: map[string]*loginSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
			arg1:  "key002",
			want1: &loginSession{response: &pb.LoginResponse{Token: "token002"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.store.getSession(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_sessionStore_save(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		store     *sessionStore
		arg1      string
		arg2      *loginSession
		wantStore map[string]*loginSession
	}{
		{name: "指定したキーがなければ新たに保存する",
			store: &sessionStore{store: map[string]*loginSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
			arg1: "key000",
			arg2: &loginSession{response: &pb.LoginResponse{Token: "token000"}},
			wantStore: map[string]*loginSession{
				"key000": {response: &pb.LoginResponse{Token: "token000"}},
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
		{name: "指定したキーが存在すれば上書きする",
			store: &sessionStore{store: map[string]*loginSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
			arg1: "key002",
			arg2: &loginSession{response: &pb.LoginResponse{Token: "token222"}},
			wantStore: map[string]*loginSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token222"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.save(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.wantStore, test.store.store) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.wantStore, test.store.store)
			}
		})
	}
}
