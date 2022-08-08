package handler

import (
	"context"

	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
)

type GetImageHandlerFunc func(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error)

func GetImage() GetImgeHandlerFunc {
	return func(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error) {
		return &imageproxy.GetImageResponse{
			Path: "AA",
		}, nil
	}
}
