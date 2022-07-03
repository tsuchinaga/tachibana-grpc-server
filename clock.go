package tachibana_grpc_server

import "time"

type iClock interface {
	now() time.Time
	today() time.Time
	nextDateTimeDuration(t time.Time, now time.Time) time.Duration
}

type clock struct{}

func (c *clock) now() time.Time {
	return time.Now()
}

func (c *clock) today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// nextDateTime - 次の指定時刻の日時を返す
func (c *clock) nextDateTime(t time.Time, now time.Time) time.Time {
	tt := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), now.Location())
	if !now.Before(tt) {
		tt = tt.Add(24 * time.Hour)
	}
	return tt
}

// nextDateTimeDuration - 次の指定時刻の日時までの時間を返す
func (c *clock) nextDateTimeDuration(t time.Time, now time.Time) time.Duration {
	return c.nextDateTime(t, now).Sub(now)
}
