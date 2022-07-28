package examples

import "log"

var (
	Target         = "localhost:8900"
	UserId         = "sdy04168"
	Password       = "27cy55a7"
	SecondPassword = "second-password"
	ClientName     = "example-client"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
