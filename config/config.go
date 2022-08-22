package config

import (
	"github.com/hanyoung-banksalad/imageproxy/server/redis"
	"github.com/smira/go-statsd"

	"github.com/banksalad/idl/gen/go/apis/v1/auth"
)

type Config interface {
	Setting() Setting

	StatsdClient() *statsd.Client

	AuthCli() auth.AuthClient

	RedisClient() *redis.Client
}

// verify DefaultConfig implements all Config interface methods
var _ Config = (*DefaultConfig)(nil)

type DefaultConfig struct {
	setting Setting

	statsdClient *statsd.Client

	authCli auth.AuthClient

	redisClient *redis.Client
}

func (c *DefaultConfig) Setting() Setting {
	return c.setting
}

func (c *DefaultConfig) StatsdClient() *statsd.Client {
	return c.statsdClient
}

func (c *DefaultConfig) AuthCli() auth.AuthClient {
	return c.authCli
}

func (c *DefaultConfig) RedisClient() *redis.Client {
	return c.redisClient
}

func NewConfig(
	setting Setting,
	statsdClient *statsd.Client,
	authCli auth.AuthClient,
	redisClient *redis.Client,
) Config {
	return &DefaultConfig{
		setting:      setting,
		statsdClient: statsdClient,
		authCli:      authCli,
		redisClient:  redisClient,
	}
}
