package redis

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/banksalad/idl/gen/go/apis/v1/mydata"
)

func TestNewClient(t *testing.T) {
	_, closer := makeClient(t)
	defer closer()
}

func TestSetAndGetAndDel(t *testing.T) {
	cli, closer := makeClient(t)
	defer closer()

	ctx := context.Background()
	t.Run("proto response", func(t *testing.T) {
		resp := &mydata.ListBankAccountsResponse{
			Accounts: []*mydata.BankAccount{
				{
					AccountNumber:            "123-45-67890",
					SequenceNumber:           wrapperspb.String("001"),
					IsConsent:                true,
					IsForeignCurrencyAccount: false,
					ProductName:              "테스트 계좌",
					IsRevolvingCreditAccount: false,
					AccountTypeEnum:          "ACCOUNT_TYPE_CHECKING",
					BankAccountStatusEnum:    "BANK_ACCOUNT_STATUS_ACTIVATED",
				},
			},
		}
		err := cli.Set(ctx, "key1", resp)
		require.NoError(t, err)

		cached := &mydata.ListBankAccountsResponse{}
		ok, err := cli.Get(ctx, "key1", cached)
		require.NoError(t, err)
		assert.True(t, ok)
		require.NotNil(t, cached)
		assert.True(t, proto.Equal(resp, cached), cmp.Diff(resp, cached, protocmp.Transform()))

		err = cli.Delete(ctx, "key1")
		assert.NoError(t, err)
		ok, err = cli.Get(ctx, "key1", cached)
		require.NoError(t, err)
		assert.False(t, ok)
	})
	t.Run("empty response", func(t *testing.T) {
		resp := &mydata.ListBankAccountsResponse{
			Accounts: nil,
		}
		err := cli.Set(ctx, "key2", resp)
		require.NoError(t, err)

		cached := &mydata.ListBankAccountsResponse{}
		ok, err := cli.Get(ctx, "key2", cached)
		require.NoError(t, err)
		assert.True(t, ok)
		require.NotNil(t, cached)
		assert.True(t, proto.Equal(resp, cached), cmp.Diff(resp, cached, protocmp.Transform()))

		err = cli.Delete(ctx, "key2")
		assert.NoError(t, err)
		ok, err = cli.Get(ctx, "key2", cached)
		require.NoError(t, err)
		assert.False(t, ok)
	})
}

func makeClient(t *testing.T) (*Client, func()) {
	fixedKey := make([]byte, 32)
	key := base64.StdEncoding.EncodeToString(fixedKey)
	fixedNonce := make([]byte, 12)
	nonce := base64.StdEncoding.EncodeToString(fixedNonce)

	s, err := miniredis.Run()
	require.NoError(t, err)

	cli, err := NewClient(s.Addr(), 1, 1, 1, key, nonce)
	require.NoError(t, err)
	require.NotNil(t, cli)

	return cli, func() {
		err := cli.Close()
		require.NoError(t, err)
		s.Close()
	}
}
