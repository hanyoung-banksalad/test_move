package imageproxy

// pr test

import (
	"sync"

	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/banksalad/go-banksalad"
)

//go:generate mockgen -package imageproxy -destination ./mock_client.go -mock_names ImageproxyClient=MockImageproxyClient github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy ImageproxyClient
const serviceConfig = `{"loadBalancingPolicy":"round_robin"}`

var (
	once sync.Once
	cli  imageproxy.ImageproxyClient

	// verify MockImageproxyClient implements all ImageproxyClient interface methods
	_ imageproxy.ImageproxyClient = (*MockImageproxyClient)(nil)
)

func GetImageproxyClient(serviceHost, caller string) imageproxy.ImageproxyClient {
	once.Do(func() {
		conn, _ := grpc.Dial(
			serviceHost,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(serviceConfig),
			grpc.WithUnaryInterceptor(banksalad.UnaryClientInterceptor(caller)),
		)

		cli = imageproxy.NewImageproxyClient(conn)
	})

	return cli
}
