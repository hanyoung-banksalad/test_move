package server

import (
	"context"

	"github.com/banksalad/imageproxy/server/handler"
	"github.com/banksalad/idl/gen/go/apis/v1/imageproxy"
)

// verify ImageproxyServer implements all interface methods
var _ imageproxy.ImageproxyServer = (*ImageproxyServer)(nil)

func (s *ImageproxyServer) HealthCheck(ctx context.Context, req *imageproxy.HealthCheckRequest) (*imageproxy.HealthCheckResponse, error) {
	return handler.HealthCheck()(ctx, req)
}
