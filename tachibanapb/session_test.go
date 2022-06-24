package tachibanapb

import (
	"reflect"
	"testing"
	"time"
)

func Test_LoginRequest_GetKey(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		req   *LoginRequest
		arg1  time.Time
		want1 string
	}{
		{name: "ハッシュ値が得られる",
			req: &LoginRequest{
				UserId:   "user-id",
				Password: "password",
			},
			arg1:  time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local),
			want1: "df8dacf6e2a94335c0fc8b1a1a05601c72152f0f30905feffcb45f74fe20ad00"},
		{name: "同じパラメータなら同じ結果が得られる",
			req: &LoginRequest{
				UserId:   "user-id",
				Password: "password",
			},
			arg1:  time.Date(2022, 6, 24, 0, 0, 0, 0, time.Local),
			want1: "df8dacf6e2a94335c0fc8b1a1a05601c72152f0f30905feffcb45f74fe20ad00"},
		{name: "日付が変われば違う結果が得られる",
			req: &LoginRequest{
				UserId:   "user-id",
				Password: "password",
			},
			arg1:  time.Date(2022, 6, 25, 0, 0, 0, 0, time.Local),
			want1: "b69aa3006a5f7044c34a34d723c56442d8054c8a042fb1555020822dbc74fddd"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.req.GetKey(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}
