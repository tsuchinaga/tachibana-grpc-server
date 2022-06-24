package tachibana_grpc_server

import "time"

type iClock interface {
	today() time.Time
}

type clock struct{}

func (c *clock) today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
