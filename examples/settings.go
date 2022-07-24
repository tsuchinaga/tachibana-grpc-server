package examples

import "log"

var (
	Target         = "localhost:8900"
	UserId         = "user-id"
	Password       = "password"
	SecondPassword = "second-password"
	ClientName     = "example-client"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
