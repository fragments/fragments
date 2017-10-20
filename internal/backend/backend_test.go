// +build integration

package backend

import (
	"fmt"
	"os"
	"testing"
)

var testETCDEndpoint = fmt.Sprintf("127.0.0.1:%s", os.Getenv("ETCD_TEST_LISTEN_PORT"))
var testVaultEndpoint = fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("VAULT_TEST_PORT"))

type validator interface {
	put(t *testing.T, key, value string)
	get(t *testing.T, key string) (string, bool)
}
