package aws

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

type mockAWS struct {
	iamMock    *mockIAM
	lambdaMock *mockLambda
}

func newMockAWS() *mockAWS {
	return &mockAWS{
		iamMock:    newMockIAM(),
		lambdaMock: newMockLambda(),
	}
}

func (m *mockAWS) iam() (iamService, error) {
	return m.iamMock, nil
}

func (m *mockAWS) lambda() (lambdaService, error) {
	return m.lambdaMock, nil
}

func snapshotService(svc interface{}) string {
	data, err := yaml.Marshal(svc)
	if err != nil {
		log.Panic(err)
	}
	return string(data)
}
