package tachibana_grpc_server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *server) LoggingMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		header, _ := metadata.FromIncomingContext(ctx)
		s.logger.request("header:", header, "request:", req)
		return handler(ctx, req)
	}
}
