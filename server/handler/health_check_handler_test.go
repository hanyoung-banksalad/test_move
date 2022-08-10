package handler

import (
	"context"
	"testing"

	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	ctx := context.Background()
	req := &imageproxy.HealthCheckRequest{}

	resp, err := HealthCheck()(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, &imageproxy.HealthCheckResponse{}, resp)
}
