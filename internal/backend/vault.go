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
		return errors.Wrap(err, "could not write data")
	}
	return nil
}

// Get gets a generic secret from Vault.
// Returns NotFoundError if the secret does not exist.
func (v *Vault) Get(ctx context.Context, key string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	s, err := v.client.Logical().Read(fmt.Sprintf("secret/%s", key))
	if err != nil {
		return "", errors.Wrap(err, "could not read value")
	}
	if s == nil {
		return "", &NotFoundError{key}
	}
	value, err := unwrapVaultData(s.Data)
	if err != nil {
		return "", err
	}
	return value, nil
}

// Delete deletes a generic secret from Vault.
// Returns NotFoundError if the secret does not exist.
//
// Since Vault doesn't return if the secret was actually deleted a check is
// done first to read the key.
func (v *Vault) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	// Check if the value exists, Vault does not say if the value was actually
	// deleted.
	path := fmt.Sprintf("secret/%s", key)
	s, err := v.client.Logical().Read(path)
	if err != nil {
		return errors.Wrap(err, "could not check if value exists")
	}
	if s == nil {
		return &NotFoundError{key}
	}

	_, err = v.client.Logical().Delete(path)
	if err != nil {
		return errors.Wrap(err, "could not delete key")
	}
	return nil
}
