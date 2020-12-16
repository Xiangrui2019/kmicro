package grpc

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/grpc"
)

// ServerOption is gRPC server option.
type ServerOption func(o *serverOptions)

type serverOptions struct {
	grpcOpts   []grpc.ServerOption
	middleware middleware.Middleware
}

func ServerMiddleware(m middleware.Middleware) ServerOption {
	return func(o *serverOptions) { o.middleware = m }
}

func ServerOptions(s ...grpc.ServerOption) ServerOption {
	return func(o *serverOptions) { o.grpcOpts = append(o.grpcOpts, s...) }
}