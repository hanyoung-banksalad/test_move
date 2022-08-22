package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
	"github.com/hanyoung-banksalad/imageproxy/server/redis"
	"github.com/sirupsen/logrus"
)

type GetImageHandlerFunc func(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error)

func GetImage(redisCli *redis.Client) GetImageHandlerFunc {
	return func(ctx context.Context, req *imageproxy.GetImageRequest) (*imageproxy.GetImageResponse, error) {
		//"https://cdn.banksalad.com/graphic/color/logo/circle/kbank.png"
		requests_strings := []string{req.Filename, req.Bg, req.Shape}
		base_url := "https://cdn.banksalad.com/graphic/%s/logo/%s/%s"
		rd_key := strings.Join(requests_strings, "_")
		rewrite_string := fmt.Sprintf(base_url, req.Bg, req.Shape, req.Filename)

		//no_image_url := "https://cdn.banksalad.com/graphic/color/noimage.png"
		res := ""
		if existCache(rd_key) {
			res = rewrite_string
		} else {
			if StatusCode(rewrite_string) {
				res = rewrite_string
				saveCache(ctx, redisCli, rd_key, rewrite_string)
			}
		}
		return &imageproxy.GetImageResponse{
			RedirectUrl: res,
		}, nil
	}
}

func StatusCode(url string) bool {
	// Setup the request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	//req.Header.Set("Authorization", AUTH)

	// Execute the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	// Close response body as required.
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusFound || resp.StatusCode == http.StatusNotModified {
		return true
	}

	return false
	// or fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}

// check if exists in s3
//func existS3(url string) bool {
//	return true
//}

func existCache(url string) bool {
	return false
}

func saveCache(ctx context.Context, redisCli *redis.Client, rd_key string, url string) bool {
	redisCli.Set(ctx, rd_key, url )
	return true
}

func scaleDown(url string) string {
	new_url := url
	return new_url
}
