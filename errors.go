package tachibana_grpc_server

import "errors"

var (
	unknownErr          = errors.New("unknown")
	notLoggedInErr      = errors.New("not logged in")
	newStreamRequestErr = errors.New("new stream request")
	stopStreamErr       = errors.New("stop stream")
)
