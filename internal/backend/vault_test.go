// +build integration

package backend

import (
	"fmt"
	"os"
	"testing"

	vaultapi "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

// vaultValidator is used for integration tests to validate the actual state in
// vault after the Vault client has interacted with it.
type vaultValidator struct {
	*vaultapi.Client
}

func newVaultValidator(t *testing.T, address string) *vaultValidator {
	t.Helper()
	cli, err := vaultapi.NewClient(&vaultapi.Config{
		Address: testVaultEndpoint,
	})
	cli.SetToken(os.Getenv("VAULT_TEST_ROOT_TOKEN"))
	if err != nil {
		t.Fatal(errors.Wrap(err, "could not create test Vault client"))
	}
	list, err := cli.Logical().List("secret/")
	if err != nil {
		t.Fatal(errors.Wrap(err, "could not list test Vault secrets for clear"))
	}
	if list != nil {
		keys := list.Data["keys"].([]interface{})
		for _, k := range keys {
			if _, err := cli.Logical().Delete(fmt.Sprintf("secret/%s", k)); err != nil {
				t.Fatal(errors.Wrapf(err, "coult not clear test Vault data: failed to delete %s", k))
			}
		}
	}
	return &vaultValidator{cli}
}

func (v *vaultValidator) put(t *testing.T, key, value string) {
	t.Helper()
	data := wrapVaultData(value)
	_, err := v.Logical().Write(fmt.Sprintf("secret/%s", key), data)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "%s: could not write test Vault data", key))
	}
	require.NoError(t, err)
}

func (v *vaultValidator) get(t *testing.T, key string) (string, bool) {
	t.Helper()
	s, err := v.Logical().Read(fmt.Sprintf("secret/%s", key))
	if err != nil {
		t.Fatal(errors.Wrapf(err, "%s: could not read test Vault data", key))
	}
	if s == nil {
		return "", false
	}
	value, err := unwrapVaultData(s.Data)
	if err != nil {
		t.Fatalf("vault: could not read %s: data does not contain %s", key, vaultDataKey)
	}
	return value, true
}

func TestNewVaultClient(t *testing.T) {
	tests := []struct {
		TestName string
		Address  string
		Error    bool
	}{
		{
			TestName: "No address",
			Address:  "",
			Error:    true,
		},
		{
			TestName: "Invalid format (no protocol)",
			Address:  "127.0.0.1:8200",
			Error:    true,
		},
		{
			TestName: "OK",
			Address:  testVaultEndpoint,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			_, err := NewVaultClient(test.Address)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
