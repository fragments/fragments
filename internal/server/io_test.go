package server

import (
	"testing"

	"github.com/fragments/fragments/pkg/testutils"
)

func TestPaths(t *testing.T) {
	paths := map[string]string{
		"function":      functionPath("function-name"),
		"deployment":    deploymentPath("deployment-name"),
		"environment":   environmentPath("environment-name"),
		"pendingupload": pendingUploadPath("pending-upload-token"),
		"user-username": userSecretName("user-secret-username"),
		"user-password": userSecretPass("user-secret-password"),
	}
	testutils.AssertGolden(t, testutils.SnapshotStringMap(paths), "testdata/paths.yaml")
}
