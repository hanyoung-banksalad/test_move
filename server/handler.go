package server

import (
	"context"

	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
	"github.com/hanyoung-banksalad/imageproxy/server/handler"
)

// verify ImageproxyServer implements all interface methods
var _ imageproxy.ImageproxyServer = (*ImageproxyServer)(nil)

func (s *ImageproxyServer) HealthCheck(ctx context.Context, req *imageproxy.HealthCheckRequest) (*imageproxy.HealthCheckResponse, error) {
	return handler.HealthCheck()(ctx, req)
}

func (s *ImageproxyServer) GetImage(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error) {
	return handler.GetImage()(ctx, req)
}
