package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
)

func TestHealthCheck(t *testing.T) {
	ctx := context.Background()
	req := &imageproxy.HealthCheckRequest{}

	resp, err := HealthCheck()(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, &imageproxy.HealthCheckResponse{}, resp)
}
