package tachibana_grpc_server

import (
	"reflect"
	"testing"

	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
)

type testSessionStore struct {
	iSessionStore
	getBySessionToken1       *accountSession
	getBySessionTokenHistory []interface{}
	getByClientToken1        *accountSession
	getByClientTokenHistory  []interface{}
	saveHistory              []interface{}
	addClientTokenHistory    []interface{}
	removeClientHistory      []interface{}
}

func (t *testSessionStore) getBySessionToken(token string) *accountSession {
	t.getBySessionTokenHistory = append(t.getBySessionTokenHistory, token)
	return t.getBySessionToken1
}
func (t *testSessionStore) getByClientToken(token string) *accountSession {
	t.getByClientTokenHistory = append(t.getByClientTokenHistory, token)
	return t.getByClientToken1
}
func (t *testSessionStore) save(sessionToken string, clientToken string, session *accountSession) {
	t.saveHistory = append(t.saveHistory, sessionToken, clientToken, session)
}
func (t *testSessionStore) addClientToken(sessionToken string, clientToken string) {
	t.addClientTokenHistory = append(t.addClientTokenHistory, sessionToken, clientToken)
}
func (t *testSessionStore) removeClient(token string) {
	t.removeClientHistory = append(t.removeClientHistory, token)
}

func Test_sessionStore_getBySessionToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *sessionStore
		arg1  string
		want1 *accountSession
	}{
		{name: "storeに指定のキーのsessionがなければnilを返す",
			store: &sessionStore{sessions: map[string]*accountSession{
				"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
				"session002": {BaseResponse: &pb.LoginResponse{Token: "session002"}},
				"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}}},
			arg1:  "session000",
			want1: nil},
		{name: "storeに指定のキーのsessionがあればsessionを返す",
			store: &sessionStore{sessions: map[string]*accountSession{
				"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
				"session002": {BaseResponse: &pb.LoginResponse{Token: "session002"}},
				"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}}},
			arg1:  "session002",
			want1: &accountSession{BaseResponse: &pb.LoginResponse{Token: "session002"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.store.getBySessionToken(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_sessionStore_getByClientToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		store *sessionStore
		arg1  string
		want1 *accountSession
	}{
		{name: "storeに指定のキーのsessionがなければnilを返す",
			store: &sessionStore{
				sessions: map[string]*accountSession{
					"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
					"session002": {BaseResponse: &pb.LoginResponse{Token: "session002"}},
					"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}},
				clientTokens: map[string]string{
					"client001": "session001",
					"client002": "session001",
					"client101": "session002",
					"client301": "session003"}},
			arg1:  "client000",
			want1: nil},
		{name: "storeに指定のキーがあってもsessionとつながっていなければnilを返す",
			store: &sessionStore{
				sessions: map[string]*accountSession{
					"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
					"session002": {BaseResponse: &pb.LoginResponse{Token: "session002"}},
					"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}},
				clientTokens: map[string]string{
					"client000": "session000",
					"client001": "session001",
					"client002": "session001",
					"client101": "session002",
					"client301": "session003"}},
			arg1:  "client000",
			want1: nil},
		{name: "storeに指定のキーのsessionがあればsessionを返す",
			store: &sessionStore{
				sessions: map[string]*accountSession{
					"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
					"session002": {BaseResponse: &pb.LoginResponse{Token: "session002"}},
					"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}},
				clientTokens: map[string]string{
					"client001": "session001",
					"client002": "session001",
					"client101": "session002",
					"client301": "session003"}},
			arg1:  "client002",
			want1: &accountSession{BaseResponse: &pb.LoginResponse{Token: "session001"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.store.getByClientToken(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_sessionStore_save(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		store            *sessionStore
		arg1             string
		arg2             string
		arg3             *accountSession
		wantSessions     map[string]*accountSession
		wantClientTokens map[string]string
	}{
		{name: "指定したキーがなければ新たに保存する",
			store: &sessionStore{
				sessions: map[string]*accountSession{
					"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
					"session002": {BaseResponse: &pb.LoginResponse{Token: "session002"}},
					"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}},
				clientTokens: map[string]string{
					"client101": "session001",
					"client102": "session001",
					"client201": "session002",
					"client301": "session003"}},
			arg1: "session000",
			arg2: "client000",
			arg3: &accountSession{BaseResponse: &pb.LoginResponse{Token: "token000"}},
			wantSessions: map[string]*accountSession{
				"session000": {BaseResponse: &pb.LoginResponse{Token: "token000"}},
				"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
				"session002": {BaseResponse: &pb.LoginResponse{Token: "session002"}},
				"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}},
			wantClientTokens: map[string]string{
				"client000": "session000",
				"client101": "session001",
				"client102": "session001",
				"client201": "session002",
				"client301": "session003"}},
		{name: "指定したキーが存在すれば上書きする",
			store: &sessionStore{
				sessions: map[string]*accountSession{
					"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
					"session002": {BaseResponse: &pb.LoginResponse{Token: "session002"}},
					"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}},
				clientTokens: map[string]string{
					"client101": "session001",
					"client102": "session001",
					"client201": "session002",
					"client301": "session003"}},
			arg1: "session002",
			arg2: "client201",
			arg3: &accountSession{BaseResponse: &pb.LoginResponse{Token: "token222"}},
			wantSessions: map[string]*accountSession{
				"session001": {BaseResponse: &pb.LoginResponse{Token: "session001"}},
				"session002": {BaseResponse: &pb.LoginResponse{Token: "token222"}},
				"session003": {BaseResponse: &pb.LoginResponse{Token: "session003"}}},
			wantClientTokens: map[string]string{
				"client101": "session001",
				"client102": "session001",
				"client201": "session002",
				"client301": "session003"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.save(test.arg1, test.arg2, test.arg3)
			if !reflect.DeepEqual(test.wantSessions, test.store.sessions) || !reflect.DeepEqual(test.wantClientTokens, test.store.clientTokens) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(),
					test.wantSessions, test.wantClientTokens,
					test.store.sessions, test.store.clientTokens)
			}
		})
	}
}

func Test_sessionStore_addClientToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		store            *sessionStore
		arg1             string
		arg2             string
		wantClientTokens map[string]string
	}{
		{name: "storeになければ追加される",
			store: &sessionStore{clientTokens: map[string]string{
				"client101": "session001",
				"client102": "session001",
				"client201": "session002",
				"client301": "session003"}},
			arg1: "session000",
			arg2: "client000",
			wantClientTokens: map[string]string{
				"client101": "session001",
				"client102": "session001",
				"client201": "session002",
				"client301": "session003",
				"client000": "session000"}},
		{name: "storeにあれば上書きされる(基本的には発生しないはずやけど)",
			store: &sessionStore{clientTokens: map[string]string{
				"client101": "session001",
				"client102": "session001",
				"client201": "session002",
				"client301": "session003"}},
			arg1: "session000",
			arg2: "client102",
			wantClientTokens: map[string]string{
				"client101": "session001",
				"client102": "session000",
				"client201": "session002",
				"client301": "session003"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.addClientToken(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.wantClientTokens, test.store.clientTokens) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.wantClientTokens, test.store.clientTokens)
			}
		})
	}
}

func Test_sessionStore_removeClient(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		store            *sessionStore
		arg1             string
		wantSessions     map[string]*accountSession
		wantClientTokens map[string]string
	}{
		{name: "指定したキーがなければ何もしない",
			store: &sessionStore{
				sessions: map[string]*accountSession{
					"session001": {Token: "session001"},
					"session002": {Token: "session002"},
					"session003": {Token: "session003"}},
				clientTokens: map[string]string{
					"client101": "session001",
					"client102": "session001",
					"client201": "session002",
					"client301": "session003"}},
			arg1: "client000",
			wantSessions: map[string]*accountSession{
				"session001": {Token: "session001"},
				"session002": {Token: "session002"},
				"session003": {Token: "session003"}},
			wantClientTokens: map[string]string{
				"client101": "session001",
				"client102": "session001",
				"client201": "session002",
				"client301": "session003"}},
		{name: "指定したキーがあれば該当するデータを削除する",
			store: &sessionStore{
				sessions: map[string]*accountSession{
					"session001": {Token: "session001"},
					"session002": {Token: "session002"},
					"session003": {Token: "session003"}},
				clientTokens: map[string]string{
					"client101": "session001",
					"client102": "session001",
					"client201": "session002",
					"client301": "session003"}},
			arg1: "client102",
			wantSessions: map[string]*accountSession{
				"session001": {Token: "session001"},
				"session002": {Token: "session002"},
				"session003": {Token: "session003"}},
			wantClientTokens: map[string]string{
				"client101": "session001",
				"client201": "session002",
				"client301": "session003"}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.removeClient(test.arg1)
			if !reflect.DeepEqual(test.wantSessions, test.store.sessions) || !reflect.DeepEqual(test.wantClientTokens, test.store.clientTokens) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(),
					test.wantSessions, test.wantClientTokens,
					test.store.sessions, test.store.clientTokens)
			}
		})
	}
}

func Test_sessionStore_clear(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		store            *sessionStore
		wantSessions     map[string]*accountSession
		wantClientTokens map[string]string
	}{
		{name: "元からマップが空なら何もしない",
			store: &sessionStore{
				sessions:     map[string]*accountSession{},
				clientTokens: map[string]string{},
			},
			wantSessions:     map[string]*accountSession{},
			wantClientTokens: map[string]string{}},
		{name: "マップに値があってもすべて削除する",
			store: &sessionStore{
				sessions: map[string]*accountSession{
					"token001": {},
					"token002": {},
				},
				clientTokens: map[string]string{
					"client-token101": "token001",
					"client-token102": "token001",
					"client-token201": "token002",
				},
			},
			wantSessions:     map[string]*accountSession{},
			wantClientTokens: map[string]string{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.store.clear()
			if !reflect.DeepEqual(test.wantSessions, test.store.sessions) || !reflect.DeepEqual(test.wantClientTokens, test.store.clientTokens) {
				t.Errorf("%s error\nwant: %+v, %+v\ngot: %+v, %+v\n", t.Name(),
					test.wantSessions, test.wantClientTokens,
					test.store.sessions, test.store.clientTokens)
			}
		})
	}
}
