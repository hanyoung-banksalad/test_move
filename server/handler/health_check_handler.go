package handler

import (
	"context"

	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
)

type HealthCheckHandlerFunc func(ctx context.Context, req *imageproxy.HealthCheckRequest) (*imageproxy.HealthCheckResponse, error)

func HealthCheck() HealthCheckHandlerFunc {
	return func(ctx context.Context, req *imageproxy.HealthCheckRequest) (*imageproxy.HealthCheckResponse, error) {
		return &imageproxy.HealthCheckResponse{}, nil
	}
}
