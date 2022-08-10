package handler

import (
	"context"
	"fmt"

	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
)

type GetImageHandlerFunc func(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error)

func GetImage() GetImageHandlerFunc {
	return func(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error) {
		res := fmt.Sprintf("%d size %s", req.Size, req.Filename)
		return &imageproxy.GetImageResponse{
			Path: res,
		}, nil
	}
}
