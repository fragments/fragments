package model

var mockType = &map[string]interface{}{
	"Test": "generic",
}

var mockFunction = &Function{
	Name: "foo",
	Labels: map[string]string{
		"foo": "foo",
	},
	Runtime:        "go",
	Checksum:       "abc",
	SourceFilename: "file.tar.gz",
	AWS: &FunctionAWS{
		Timeout: 3,
		Memory:  512,
	},
}

var mockDeployment = &Deployment{
	Name: "deploy",
	FunctionLabels: map[string]string{
		"func": "foo",
	},
	EnvironmentLabels: map[string]string{
		"deploy": "bar",
	},
}

var mockEnvironment = &Environment{
	Name: "env",
	Labels: map[string]string{
		"foo": "foo",
	},
	Infrastructure: InfrastructureTypeAWS,
	AWS: &InfrastructureAWS{
		Region: "us-west-2",
	},
}

var mockPendingUpload = &PendingUpload{
	Token:    "abc",
	Filename: "file.tar.gz",
	Function: mockFunction,
}
