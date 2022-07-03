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

func Test_clock_now(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		want1 time.Time
	}{
		{name: "先に定義したwantより後の日付が手に入る", want1: time.Now()},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			clock := &clock{}
			got1 := clock.now()
			if test.want1.After(got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_clock_nextDateTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		arg1  time.Time
		arg2  time.Time
		want1 time.Time
	}{
		{name: "指定時刻が未来ならそのまま返す",
			arg1:  time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
			arg2:  time.Date(2022, 5, 20, 7, 0, 0, 0, time.Local),
			want1: time.Date(2022, 5, 20, 8, 0, 0, 0, time.Local)},
		{name: "指定時刻 = 現在時刻なら翌日の指定時刻を返す",
			arg1:  time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
			arg2:  time.Date(2022, 5, 20, 8, 0, 0, 0, time.Local),
			want1: time.Date(2022, 5, 21, 8, 0, 0, 0, time.Local)},
		{name: "指定時刻が過去なら翌日の指定時刻を返す",
			arg1:  time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
			arg2:  time.Date(2022, 5, 20, 9, 0, 0, 0, time.Local),
			want1: time.Date(2022, 5, 21, 8, 0, 0, 0, time.Local)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			clock := &clock{}
			got1 := clock.nextDateTime(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}

func Test_clock_nextDateTimeDuration(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		arg1  time.Time
		arg2  time.Time
		want1 time.Duration
	}{
		{name: "指定時刻が未来ならそのまま返す",
			arg1:  time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
			arg2:  time.Date(2022, 5, 20, 7, 0, 0, 0, time.Local),
			want1: 1 * time.Hour},
		{name: "指定時刻 = 現在時刻なら翌日の指定時刻を返す",
			arg1:  time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
			arg2:  time.Date(2022, 5, 20, 8, 0, 0, 0, time.Local),
			want1: 24 * time.Hour},
		{name: "指定時刻が過去なら翌日の指定時刻を返す",
			arg1:  time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
			arg2:  time.Date(2022, 5, 20, 9, 0, 0, 0, time.Local),
			want1: 23 * time.Hour},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			clock := &clock{}
			got1 := clock.nextDateTimeDuration(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}
