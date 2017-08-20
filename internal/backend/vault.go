package backend

import (
	"context"
	"fmt"

	vaultapi "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

// Vault implemnts KV by storing values in Vault as generic secrets.
// Only the generic secret backend is supported.
type Vault struct {
	client *vaultapi.Client
}

const vaultDataKey = "data"

// wrapVaultData wraps a string in a map to be stored in vault.
// This is because Vault can only store map[string]interface{} but we want to
// support a straightforward key-value interface
func wrapVaultData(value string) map[string]interface{} {
	return map[string]interface{}{
		vaultDataKey: value,
	}
}

// unwrapVaultData extracts the value from a vault secret's data
func unwrapVaultData(input map[string]interface{}) (string, error) {
	v, ok := input[vaultDataKey].(string)
	if !ok {
		return "", errors.Errorf("data does not contain %s", vaultDataKey)
	}
	return v, nil
}

// NewVaultClient creates a new Vault client and connects to the Vault server
// The environment variable VAULT_TOKEN is read automatically to authenticate
// the client.
// Returns an error if the address is not in a valid url.
func NewVaultClient(address string) (*Vault, error) {
	if address == "" {
		return nil, errors.New("no address supplied")
	}
	cli, err := vaultapi.NewClient(&vaultapi.Config{
		Address: address,
	})
	if err != nil {
		return nil, err
	}
	return &Vault{
		client: cli,
	}, nil
}

// Put stores a generic secret in Vault. If the key already exists it is
// overwritten.
// If the target data is other keys than the vaultDataKey the entire map is
// overwritten.
func (v *Vault) Put(ctx context.Context, key, value string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if key == "" {
		return errors.New("key is empty")
	}
	data := wrapVaultData(value)
	_, err := v.client.Logical().Write(fmt.Sprintf("secret/%s", key), data)
	if err != nil {
		return errors.Wrap(err, "could not write vault data")
	}
	return nil
}

// Get gets a generic secret from Vault.
// Returns ErrNotFound if the secret does not exist.
func (v *Vault) Get(ctx context.Context, key string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	s, err := v.client.Logical().Read(fmt.Sprintf("secret/%s", key))
	if err != nil {
		return "", errors.Wrap(err, "could not read value from vault")
	}
	if s == nil {
		return "", &ErrNotFound{key}
	}
	value, err := unwrapVaultData(s.Data)
	if err != nil {
		return "", err
	}
	return value, nil
}
