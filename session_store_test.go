package tachibana_grpc_server

import (
	"reflect"
	"testing"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type testSessionStore struct {
	iSessionStore
	getSession1       *accountSession
	getSessionHistory []interface{}
	saveHistory       []interface{}
	removeHistory     []interface{}
}

func (t *testSessionStore) getSession(key string) *accountSession {
	t.getSessionHistory = append(t.getSessionHistory, key)
	return t.getSession1
}
func (t *testSessionStore) save(key string, session *accountSession) {
	t.saveHistory = append(t.saveHistory, key, session)
}
func (t *testSessionStore) remove(key string) {
	t.removeHistory = append(t.saveHistory, key)
}

func Test_sessionStore_getSession(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *sessionStore
		arg1  string
		want1 *accountSession
	}{
		{name: "storeに指定のキーのsessionがなければnilを返す",
			store: &sessionStore{store: map[string]*accountSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
			arg1:  "key000",
			want1: nil},
		{name: "storeに指定のキーのsessionがあればsessionを返す",
			store: &sessionStore{store: map[string]*accountSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
			arg1:  "key002",
			want1: &accountSession{response: &pb.LoginResponse{Token: "token002"}}},
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
		arg2      *accountSession
		wantStore map[string]*accountSession
	}{
		{name: "指定したキーがなければ新たに保存する",
			store: &sessionStore{store: map[string]*accountSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
			arg1: "key000",
			arg2: &accountSession{response: &pb.LoginResponse{Token: "token000"}},
			wantStore: map[string]*accountSession{
				"key000": {response: &pb.LoginResponse{Token: "token000"}},
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
		{name: "指定したキーが存在すれば上書きする",
			store: &sessionStore{store: map[string]*accountSession{
				"key001": {response: &pb.LoginResponse{Token: "token001"}},
				"key002": {response: &pb.LoginResponse{Token: "token002"}},
				"key003": {response: &pb.LoginResponse{Token: "token003"}}}},
			arg1: "key002",
			arg2: &accountSession{response: &pb.LoginResponse{Token: "token222"}},
			wantStore: map[string]*accountSession{
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

func Test_sessionStore_remove(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *sessionStore
		arg1  string
		want1 map[string]*accountSession
	}{
		{name: "指定したキーがなければ何もしない",
			store: &sessionStore{store: map[string]*accountSession{
				"token001": {token: "token001"},
				"token002": {token: "token002"},
				"token003": {token: "token003"},
			}},
			arg1: "token000",
			want1: map[string]*accountSession{
				"token001": {token: "token001"},
				"token002": {token: "token002"},
				"token003": {token: "token003"},
			}},
		{name: "指定したキーがあれば該当するデータを削除する",
			store: &sessionStore{store: map[string]*accountSession{
				"token001": {token: "token001"},
				"token002": {token: "token002"},
				"token003": {token: "token003"},
			}},
			arg1: "token002",
			want1: map[string]*accountSession{
				"token001": {token: "token001"},
				"token003": {token: "token003"},
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.remove(test.arg1)
			if !reflect.DeepEqual(test.want1, test.store.store) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, test.store.store)
			}
		})
	}
}
