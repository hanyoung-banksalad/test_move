package config

import (
	"log"
	"os"
	"strconv"

	"github.com/banksalad/go-banksalad"
)

type Setting struct {
	ServiceName        string
	GRPCServerEndpoint string
	GRPCServerPort     string
	HTTPServerPort     string
	StatsdUDPPort      string

	Env       string
	SubEnvID  string
	Namespace string

	GracefulShutdownTimeoutMs int

	AuthGRPCEndpoint string

	RedisHost             string
	RedisPoolSize         int
	RedisMinIdleConns     int
	RedisExpiresInMinutes int
	RedisEncryptionKey    string
	RedisEncryptionNonce  string
}

func NewSetting() Setting {
	return Setting{
		ServiceName:        "imageproxy",
		GRPCServerEndpoint: banksalad.MustGetEnv("GRPC_SERVER_ENDPOINT", "localhost:18081"),
		GRPCServerPort:     banksalad.MustGetEnv("GRPC_SERVER_PORT", "18081"),
		HTTPServerPort:     banksalad.MustGetEnv("HTTP_SERVER_PORT", "18082"),
		StatsdUDPPort:      banksalad.MustGetEnv("STATSD_UDP_PORT", "8125"),

		Env:       banksalad.MustGetEnv("ENV", "development"),
		SubEnvID:  banksalad.MustGetEnv("SUB_ENV_ID", "local"),
		Namespace: banksalad.MustGetEnv("NAMESPACE", "development-local"),

		GracefulShutdownTimeoutMs: banksalad.MustAtoi(banksalad.MustGetEnv("GRACEFUL_SHUTDOWN_TIMEOUT_MS", "10000")),

		AuthGRPCEndpoint: banksalad.MustGetEnv("AUTH_GRPC_ENDPOINT", "dns:///auth-headless:18081"),

		RedisHost:             getEnv("REDIS_HOST", "localhost:6379"),
		RedisPoolSize:         mustAtoi(getEnv("REDIS_POOL_SIZE", "100")),
		RedisMinIdleConns:     mustAtoi(getEnv("REDIS_MIN_IDLE_CONNS", "30")),
		RedisExpiresInMinutes: mustAtoi(getEnv("REDIS_EXPIRES_IN_MINUTES", "10")),
		//RedisEncryptionKey:    getEnv("REDIS_ENCRYPTION_KEY", ""),
		//RedisEncryptionNonce:  getEnv("REDIS_ENCRYPTION_NONCE", ""),
	}
}

func getEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		if defaultValue != "" {
			value = defaultValue
		} else {
			log.Fatalf("missing required environment variable: %v", key)
		}
	}
	return value
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
