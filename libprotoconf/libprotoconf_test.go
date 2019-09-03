package libprotoconf

import (
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/golang/protobuf/ptypes"
	vault "github.com/hashicorp/vault/api"
	assert "github.com/stretchr/testify/require"
	protoconfvalue "protoconf.com/datatypes/proto/v1/protoconfvalue"
	crawlerpb "protoconf.com/examples/crawler"
	"protoconf.com/protostdlib/secret"
	"protoconf.com/tests/vaultmock"
)

var (
	client     *vault.Client
	httpserver *httptest.Server
)

func setUpVaultTest(t *testing.T) func(*testing.T) {
	httpserver, client = vaultmock.RunMockVault(t)

	_, err := client.Sys().Health()
	assert.NoError(t, err)

	mounts, err := client.Sys().ListMounts()
	assert.NoError(t, err)
	assert.NotEqual(t, mounts, "")
	data := map[string]interface{}{
		"HELLO": "world",
	}
	_, err = client.Logical().Write("secret/test", data)
	assert.NoError(t, err)
	secret, err := client.Logical().Read("secret/test")
	assert.NoError(t, err)
	assert.Equal(t, secret.Data, data)

	// tearDown
	return func(*testing.T) {
		httpserver.Close()
	}
}

func TestVaultSecret(t *testing.T) {
	tearDown := setUpVaultTest(t)
	defer tearDown(t)
	var wg sync.WaitGroup

	secret_ := &secret.Secret{
		Secret: &secret.Secret_VaultSecret{
			VaultSecret: &secret.VaultSecret{Path: "secret/test", Key: "HELLO"},
		},
	}
	assert.Equal(t, secret_, secret_)
	// TODO: Fix the way we initialize the Vault client for libprotoconf to all the test client usage.
	// config := &crawlerpb.CrawlerService{DbPassword: secret_}
	config := &crawlerpb.CrawlerService{}
	any, err := ptypes.MarshalAny(config)
	assert.NoError(t, err)
	payload := &protoconfvalue.ProtoconfValue{Value: any}

	valuesCh := make(chan *protoconfvalue.ProtoconfValue)
	w, _ := NewTestWatcher(valuesCh)
	stopCh := make(chan struct{})
	resultsCh, err := w.Watch("test", stopCh)
	assert.NoError(t, err)
	wg.Add(1)
	go func() {
		res := <-resultsCh

		assert.NoError(t, res.Error)
		c1 := &crawlerpb.CrawlerService{}
		err = ptypes.UnmarshalAny(res.Value, c1)
		assert.Equal(t, config, c1)
		wg.Done()
	}()
	valuesCh <- payload
	wg.Wait()
}
