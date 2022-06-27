package tachibana_grpc_server

import "errors"

var (
	unknownErr     = errors.New("unknown")
	notLoggedInErr = errors.New("not logged in")
)
