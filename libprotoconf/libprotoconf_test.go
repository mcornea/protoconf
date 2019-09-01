package libprotoconf

import (
	"protoconf.com/tests/vaultmock"
	"testing"
)

func TestVaultSecret(t *testing.T) {
	// httpserver, addr, certPool, cert, test_app_id, test_user_id := vaultmock.RunMockVault(t)
	_, _, _, _, _, test_user_id := vaultmock.RunMockVault(t)
	if test_user_id == "---" {
		t.Fail()
	}
}
