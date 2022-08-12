package handler

import (
	"context"
	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
)

type GetImageHandlerFunc func(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error)

func GetImage() GetImageHandlerFunc {
	return func(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error) {
		res := "https://cdn.banksalad.com/graphic/color/logo/circle/kbank.png"
		return &imageproxy.GetImageResponse{
			Path: res,
		}, nil
	}
}
