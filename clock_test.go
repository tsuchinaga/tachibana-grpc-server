package tachibana_grpc_server

import (
	"reflect"
	"testing"
	"time"
)

type testClock struct {
	iClock
	today1 time.Time
}

func (t *testClock) today() time.Time {
	return t.today1
}

func Test_clock_today(t *testing.T) {
	t.Parallel()
	now := time.Now()
	tests := []struct {
		name  string
		want1 time.Time
	}{
		{name: "当日が取得できる", want1: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			clock := &clock{}
			got1 := clock.today()
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}
