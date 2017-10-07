// +build integration

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
					return "notfound"
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
					return "notfound"
				}
				value, err := unwrapVaultData(s.Data)
				if err != nil {
					t.Fatalf("vault: could not read %s: data does not contain %s", key, vaultDataKey)
				}
				return value
			},
		},
		{
			TargetName: "Test",
			TestClient: func(t *testing.T) interface{} {
				// return the map that's used for storing data
				return map[string]string{}
			},
			New: func(data interface{}) KV {
				d := data.(map[string]string)
				return &TestKV{Data: d}
			},
			AddTestValue: func(t *testing.T, data interface{}, key, value string) {
				d := data.(map[string]string)
				d[key] = value
			},
			GetTestValue: func(t *testing.T, data interface{}, key string) string {
				d := data.(map[string]string)
				v, ok := d[key]
				if !ok {
					return "notfound"
				}
				return v
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

						if test.NotFound {
							assert.True(t, test.NotFound, IsNotFound(err))
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

func TestDelete(t *testing.T) {
	targets := getTestTargets(t)

	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	tests := []struct {
		TestName string
		Context  context.Context
		Key      string
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
		},
	}

	for _, target := range targets {
		t.Run(target.TargetName, func(t *testing.T) {
			testClient := target.TestClient(t)
			client := target.New(testClient)
			target.AddTestValue(t, testClient, "foo", "bar")

			for _, test := range tests {
				t.Run(test.TestName, func(t *testing.T) {
					err := client.Delete(test.Context, test.Key)
					if test.Error || test.NotFound {
						require.Error(t, err)
						if test.NotFound {
							assert.True(t, test.NotFound, IsNotFound(err))
						}
						return
					}
					require.NoError(t, err)

					actual := target.GetTestValue(t, testClient, test.Key)
					assert.Equal(t, "notfound", actual)
				})
			}
			if closer, ok := client.(io.Closer); ok {
				_ = closer.Close()
			}
		})
	}
}

func TestList(t *testing.T) {
	targets := getTestTargets(t)

	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	existing := map[string]string{
		"/foo":         "foo",
		"/foo/bar":     "bar",
		"/foo/bar/baz": "baz",
	}

	tests := []struct {
		TestName string
		Context  context.Context
		Root     string
		Expected map[string]string
		Error    bool
	}{
		{
			TestName: "Context canceled",
			Context:  ctxCanceled,
			Root:     "foo",
			Error:    true,
		},
		{
			TestName: "No root match",
			Context:  ctx,
			Root:     "/bar",
			Expected: map[string]string{},
		},
		{
			TestName: "Match",
			Context:  ctx,
			Root:     "/foo/bar",
			Expected: map[string]string{
				"baz": "baz",
			},
		},
		{
			TestName: "Nested",
			Context:  ctx,
			Root:     "/foo",
			Expected: map[string]string{
				"bar":     "bar",
				"bar/baz": "baz",
			},
		},
	}

	for _, target := range targets {
		t.Run(target.TargetName, func(t *testing.T) {
			testClient := target.TestClient(t)
			client := target.New(testClient)
			lister, ok := client.(Lister)
			if !ok {
				t.Logf("%s does not implement Lister", target.TargetName)
				return
			}

			for k, v := range existing {
				target.AddTestValue(t, testClient, k, v)
			}

			for _, test := range tests {
				t.Run(test.TestName, func(t *testing.T) {
					actual, err := lister.List(test.Context, test.Root)
					if test.Error {
						require.Error(t, err)
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
