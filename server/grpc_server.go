package server

import (
	"time"

	"github.com/hanyoung-banksalad/imageproxy/config"
	"github.com/hanyoung-banksalad/imageproxy/idl/gen/go/apis/v1/imageproxy"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/banksalad/go-banksalad"
)

type ImageproxyServer struct {
	cfg config.Config
}

func NewImageproxyServer(cfg config.Config) (*ImageproxyServer, error) {
	return &ImageproxyServer{cfg: cfg}, nil
}

func (s *ImageproxyServer) Config() config.Config {
	return s.cfg
}

func (s *ImageproxyServer) RegisterServer(srv *grpc.Server) {
	imageproxy.RegisterImageproxyServer(srv, s)
}

func NewGRPCServer(cfg config.Config) (*grpc.Server, error) {
	logrus.ErrorKey = "grpc.error"

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			banksalad.UnaryServerInterceptor(cfg.StatsdClient().CloneWithPrefixExtension(".grpc"), log),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 30 * time.Second,
		}),
	)

	imageproxyServer, err := NewImageproxyServer(cfg)
	if err != nil {
		return nil, err
	}

	imageproxy.RegisterImageproxyServer(grpcServer, imageproxyServer)
	reflection.Register(grpcServer)

	return grpcServer, nil
}
