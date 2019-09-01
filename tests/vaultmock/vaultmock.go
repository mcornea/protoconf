package vaultmock

import (
	// "bytes"
	"context"
	"crypto/sha256"
	"crypto/x509"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	log "github.com/hashicorp/go-hclog"
	vaultAPI "github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/audit"
	credAppId "github.com/hashicorp/vault/builtin/credential/app-id"
	vaultHTTP "github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/sdk/helper/salt"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hashicorp/vault/sdk/physical/inmem"
	"github.com/hashicorp/vault/vault"
)

func GetServer(core *vault.Core) *httptest.Server {
	mux := http.NewServeMux()
	mux.Handle("/", vaultHTTP.Handler(
		&vault.HandlerProperties{Core: core}))

	s := httptest.NewTLSServer(mux)

	return s
}

func RunMockVault(t *testing.T) (*httptest.Server, *vaultAPI.Client) {
	core, err := vault.NewCore(GetMockVaultConfig())
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	keys, rootToken := vault.TestCoreInit(t, core)
	for _, key := range keys {
		if _, err = core.Unseal(key); err != nil {
			t.Fatalf("unseal err: %s", err)
		}
	}
	if core.Sealed() {
		t.Fatalf("err checking seal status: %s", err)
	}
	s := GetServer(core)
	addr := s.URL

	// Create client to vault for configuration
	cfg := vaultAPI.DefaultConfig()
	cfg.Address = addr
	certBytes := s.TLS.Certificates[0].Certificate[0]
	cert, _ := x509.ParseCertificate(certBytes)
	certPool := x509.NewCertPool()
	certPool.AddCert(cert)
	cfg.HttpClient.Transport.(*http.Transport).TLSClientConfig.ClientCAs = certPool
	// Turn off certificate check (don't do this in production)
	cfg.ConfigureTLS(&vaultAPI.TLSConfig{Insecure: true})
	c, err := vaultAPI.NewClient(cfg)
	if err != nil {
		t.Fatalf("Error creating client in mock vault setup: %v\n", err)
	}
	c.SetToken(rootToken)

	// Set policy to allow use of anything /secrets/*
	rules := `path "secret/*" {
  policy = "write"
}`
	err = c.Sys().PutPolicy("allSecrets_Test", rules)
	if err != nil {
		t.Fatalf("Error applying policy: %v", err)
	}

	// Mount a KV store on `secret/`
	mountConfig := &vaultAPI.MountInput{Type: "kv", Description: "secret"}
	err = c.Sys().Mount("secret", mountConfig)
	if err != nil {
		t.Fatalf("Error mounting secret/: %v", err)
	}

	return s, c

}

func GetMockVaultConfig() *vault.CoreConfig {
	logger := log.L()
	inm, _ := inmem.NewInmem(nil, logger)

	noopAudits := map[string]audit.Factory{
		"noop": func(ctx context.Context, config *audit.BackendConfig) (audit.Backend, error) {
			view := &logical.InmemStorage{}
			view.Put(ctx, &logical.StorageEntry{
				Key:   "salt",
				Value: []byte("foo"),
			})
			config.SaltConfig = &salt.Config{
				HMAC:     sha256.New,
				HMACType: "hmac-sha256",
			}
			config.SaltView = view
			return &NoopAudit{
				Config: config,
			}, nil
		},
	}

	conf := &vault.CoreConfig{
		Physical:      inm,
		AuditBackends: noopAudits,
		LogicalBackends: map[string]logical.Factory{
			"generic": vault.LeasedPassthroughBackendFactory,
		},
		CredentialBackends: map[string]logical.Factory{
			"app-id": credAppId.Factory,
		},
		HAPhysical:   nil,
		DisableMlock: true,
		Logger:       logger,
	}

	return conf
}

type NoopAudit struct {
	Config    *audit.BackendConfig
	salt      *salt.Salt
	saltMutex sync.RWMutex
}

func (n *NoopAudit) GetHash(ctx context.Context, data string) (string, error) {
	salt, err := n.Salt()
	if err != nil {
		return "", err
	}
	return salt.GetIdentifiedHMAC(data), nil
}

func (n *NoopAudit) LogRequest(ctx context.Context, l *logical.LogInput) error {
	return nil
}

func (n *NoopAudit) LogResponse(ctx context.Context, l *logical.LogInput) error {
	return nil
}

func (n *NoopAudit) Reload(ctx context.Context) error {
	return nil
}

func (n *NoopAudit) Invalidate(ctx context.Context) {
	n.saltMutex.Lock()
	defer n.saltMutex.Unlock()
	n.salt = nil
}

func (n *NoopAudit) Salt() (*salt.Salt, error) {
	n.saltMutex.RLock()
	if n.salt != nil {
		defer n.saltMutex.RUnlock()
		return n.salt, nil
	}
	n.saltMutex.RUnlock()
	n.saltMutex.Lock()
	defer n.saltMutex.Unlock()
	if n.salt != nil {
		return n.salt, nil
	}
	salt, err := salt.NewSalt(context.Background(), n.Config.SaltView, n.Config.SaltConfig)
	if err != nil {
		return nil, err
	}
	n.salt = salt
	return salt, nil
}
