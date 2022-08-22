package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type Client struct {
	ttl time.Duration

	cli *redis.Client

	// AES256 GCM 암복호화를 위함
	key, nonce []byte
}

const (
	aes256KeyBytesSize      = 32
	aes256GCMNonceBytesSize = 12
)

func NewClient(
	redisHost string,
	poolSize int,
	minIdleConns int,
	expiresInMinutes int,
//keyBase64Encoded string,
//nonceBase64Encoded string,
) (*Client, error) {
	/*
		key, err := base64.StdEncoding.DecodeString(keyBase64Encoded)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(key) != aes256KeyBytesSize {
			return nil, errors.Errorf("key size should be %d", aes256KeyBytesSize)
		}

		nonce, err := base64.StdEncoding.DecodeString(nonceBase64Encoded)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(nonce) != aes256GCMNonceBytesSize {
			return nil, errors.Errorf("nonce size should be %d", aes256GCMNonceBytesSize)
		}
	*/

	cli := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Client{
		ttl: time.Duration(expiresInMinutes) * time.Minute,
		cli: cli,
		//key:   key,
		//nonce: nonce,
	}, nil
}

func (c *Client) Set(ctx context.Context, key string, value string) error {
	/*
		// proto.Message 자체가 아닌 json 값으로서 캐시하기 위해 의도적으로 protojson 패키지를 사용하지 않음
		bb, err := json.Marshal(p)
		if err != nil {
			return errors.Wrapf(err, "marshal to json value: %s", p.ProtoReflect().Descriptor().Name())
		}

		s, err := banksalad.EncryptByAESGCM(c.key, c.nonce, string(bb), nil)
		if err != nil {
			return errors.Wrap(err, "encrypt")
		}
	*/

	if err := c.cli.SetEX(ctx, key, value, c.ttl).Err(); err != nil {
		return errors.Wrap(err, "redis.SetEX")
	}

	return nil
}

//func (c *Client) Get(ctx context.Context, key string, p proto.Message) (bool, error) {
func (c *Client) Get(ctx context.Context, key string) (bool, error) {
	s, err := c.cli.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, errors.Wrap(err, "redis.Get")
	}
	if s == "" {
		return false, nil
	}
	/*
		decrypted, err := banksalad.DecryptByAESGCM(c.key, c.nonce, s, nil)
		if err != nil {
			return false, errors.Wrap(err, "decrypt")
		}

		if err := json.Unmarshal([]byte(decrypted), p); err != nil {
			return false, errors.Wrapf(err, "unmarshal from json value: %s", p.ProtoReflect().Descriptor().Name())
		}
	*/
	return true, nil
}

/*
func (c *Client) Delete(ctx context.Context, keys ...string) error {
	if err := c.cli.Del(ctx, keys...).Err(); err != nil {
		return errors.Wrap(err, "redis.Del")
	}
	return nil
}
*/

func (c *Client) Close() error {
	return c.cli.Close()
}

// NewMockClient 함수는 테스트 코드에서 사용할 수 있는 in-memory redis 클라이언트를 반환합니다.
// e.g.
//   mockRedisCli, closer := redis.NewMockClient()
//   defer closer()
/*
func NewMockClient() (*Client, func()) {
	// 32 bytes zero filled
	//key := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
	// 12 bytes zero filled
	//nonce := "AAAAAAAAAAAAAAAA"

	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	cli, err := NewClient(s.Addr(), 1, 1, 1, key, nonce)
	if err != nil {
		panic(err)
	}
	return cli, func() {
		err := cli.Close()
		if err != nil {
			panic(err)
		}
		s.Close()
	}
}
*/
