package backend

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	etcdapi "github.com/coreos/etcd/clientv3"
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testETCDEndpoint = fmt.Sprintf("127.0.0.1:%s", os.Getenv("ETCD_TEST_LISTEN_PORT"))
var testVaultEndpoint = fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("VAULT_TEST_PORT"))

type testTarget struct {
	TargetName   string
	TestClient   func(t *testing.T) interface{}
	New          func(testClient interface{}) KV
	AddTestValue func(t *testing.T, testClient interface{}, key, value string)
	GetTestValue func(t *testing.T, testClient interface{}, key string) string
}

func getTestTargets(t *testing.T) []testTarget {
	return []testTarget{
		{
			TargetName: "ETCD",
			TestClient: func(t *testing.T) interface{} {
				cli, err := etcdapi.New(etcdapi.Config{
					Endpoints:   []string{testETCDEndpoint},
					DialTimeout: 3 * time.Second,
				})
				if err != nil {
					t.Fatal(errors.Wrap(err, "could not create test ETCD client"))
				}
				return cli
			},
			New: func(testClient interface{}) KV {
				cli := testClient.(*etcdapi.Client)
				if _, err := cli.Delete(context.Background(), "", etcdapi.WithPrefix()); err != nil {
					t.Fatal(errors.Wrap(err, "could not clean test ETCD db"))
				}
				return &ETCD{cli}
			},
			AddTestValue: func(t *testing.T, testClient interface{}, key, value string) {
				cli := testClient.(*etcdapi.Client)
				if _, err := cli.Put(context.Background(), key, value); err != nil {
					t.Fatal(errors.Wrapf(err, "could not put test ETCD value: %s", key))
				}
			},
			GetTestValue: func(t *testing.T, testClient interface{}, key string) string {
				cli := testClient.(*etcdapi.Client)
				res, err := cli.Get(context.Background(), key)
				if err != nil {
					t.Fatal(errors.Wrapf(err, "could not get test ETCD value: %s", key))
				}
				if res.Count < 1 {
					t.Fatalf("test ETCD get did not return any values: %s", key)
				}
				return string(res.Kvs[0].Value)
			},
		},
		{
			TargetName: "Vault",
			TestClient: func(t *testing.T) interface{} {
				cli, err := vaultapi.NewClient(&vaultapi.Config{
					Address: testVaultEndpoint,
				})
				cli.SetToken(os.Getenv("VAULT_TEST_ROOT_TOKEN"))
				if err != nil {
					t.Fatal(errors.Wrap(err, "could not create test Vault client"))
				}
				return cli
			},
			New: func(testClient interface{}) KV {
				cli := testClient.(*vaultapi.Client)
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
				return &Vault{
					client: cli,
				}
			},
			AddTestValue: func(t *testing.T, testClient interface{}, key, value string) {
				cli := testClient.(*vaultapi.Client)
				data := wrapVaultData(value)
				_, err := cli.Logical().Write(fmt.Sprintf("secret/%s", key), data)
				if err != nil {
					t.Fatal(errors.Wrapf(err, "%s: could not write test Vault data", key))
				}
			},
			GetTestValue: func(t *testing.T, testClient interface{}, key string) string {
				cli := testClient.(*vaultapi.Client)
				s, err := cli.Logical().Read(fmt.Sprintf("secret/%s", key))
				if err != nil {
					t.Fatal(errors.Wrapf(err, "%s: could not read test Vault data", key))
				}
				if s == nil {
					t.Fatalf("test Vault get did not return any values: %s", key)
				}
				value, err := unwrapVaultData(s.Data)
				if err != nil {
					t.Fatalf("vault: could not read %s: data does not contain %s", key, vaultDataKey)
				}
				return value
			},
		},
	}
}

func TestPut(t *testing.T) {
	targets := getTestTargets(t)

	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	tests := []struct {
		TestName string
		Context  context.Context
		Key      string
		Value    string
		Error    bool
	}{
		{
			TestName: "Empty key",
			Context:  ctx,
			Key:      "",
			Value:    "empty",
			Error:    true,
		},
		{
			TestName: "Context canceled",
			Context:  ctxCanceled,
			Key:      "canceled",
			Value:    "context",
			Error:    true,
		},
		{
			TestName: "OK",
			Context:  ctx,
			Key:      "foo",
			Value:    "foo",
		},
		{
			TestName: "Overwrite",
			Context:  ctx,
			Key:      "bar",
			Value:    "foobar",
		},
	}

	for _, target := range targets {
		t.Run(target.TargetName, func(t *testing.T) {
			testClient := target.TestClient(t)
			client := target.New(testClient)
			target.AddTestValue(t, testClient, "bar", "bar")

			for _, test := range tests {
				t.Run(test.TestName, func(t *testing.T) {
					err := client.Put(test.Context, test.Key, test.Value)
					if test.Error {
						require.Error(t, err)
						return
					}
					require.NoError(t, err)

					actual := target.GetTestValue(t, testClient, test.Key)
					assert.Equal(t, test.Value, actual)
				})
			}
			if closer, ok := client.(io.Closer); ok {
				_ = closer.Close()
			}
		})
	}
}

func TestGet(t *testing.T) {
	targets := getTestTargets(t)

	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	tests := []struct {
		TestName string
		Context  context.Context
		Key      string
		Expected string
		Error    bool
		NotFound bool
	}{
		{
			TestName: "Context canceled",
			Context:  ctxCanceled,
			Key:      "foo",
			Error:    true,
		},
		{
			TestName: "Not found",
			Context:  ctx,
			Key:      "baz",
			NotFound: true,
		},
		{
			TestName: "Found",
			Context:  ctx,
			Key:      "foo",
			Expected: "bar",
		},
	}

	for _, target := range targets {
		t.Run(target.TargetName, func(t *testing.T) {
			testClient := target.TestClient(t)
			client := target.New(testClient)
			target.AddTestValue(t, testClient, "foo", "bar")

			for _, test := range tests {
				t.Run(test.TestName, func(t *testing.T) {
					actual, err := client.Get(test.Context, test.Key)
					if test.Error || test.NotFound {
						require.Error(t, err)

						switch errors.Cause(err).(type) {
						case *ErrNotFound:
							assert.True(t, test.NotFound)
							assert.Contains(t, err.Error(), test.Key)
						default:
							assert.False(t, test.NotFound)
						}

						return
					}
					require.NoError(t, err)

					assert.EqualValues(t, test.Expected, actual)
				})
			}
			if closer, ok := client.(io.Closer); ok {
				_ = closer.Close()
			}
		})
	}
}
