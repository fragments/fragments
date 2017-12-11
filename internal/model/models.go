package model

// Function represents a function specification.
type Function struct {
	// Name is the unique name for a function.
	Name string `json:"name,omitempty"`
	// Labels are labels used to identify a function.
	Labels map[string]string `json:"labels,omitempty"`
	// Runtime is the function runtime.
	Runtime string `json:"runtime,omitempty"`
	// Checksum is the checksum calculated by the client of the source files for.
	// the function
	Checksum string `json:"checksum,omitempty"`
	// SourceFilename is the name of the source file. For functions being created
	// this is ignored, it is set when the source has been confirmed.
	SourceFilename string `json:"source_filename,omitempty"`
	// AWS is the Amazon Web Services specific configuration for the function.
	AWS *FunctionAWS `json:"aws,omitempty"`
}

// FunctionAWS contains AWS function (Lambda) specific configuration info.
type FunctionAWS struct {
	// Timeout is the timeout in seconds for the function.
	Timeout int64 `json:"timeout,omitempty"`
	// Memory is the memory in mb for the function.,
	Memory int64 `json:"memory,omitempty"`
}

// PendingUpload is a source code request that has been returned to the client.
// When source is confirmed it is used to fetch the source and apply function
// changes.
type PendingUpload struct {
	// Token is the unique token to identify a pending upload.
	Token string `json:"token,omitempty"`
	// Filename is the filename to retrieve the source by from the filestore.
	Filename string `json:"filename,omitempty"`
	// PreviousFilename is the filename of the previous source in case the
	// function code has been updated. It is blank if the function did not exist
	// before.
	PreviousFilename string `json:"previous_filename,omitempty"`
	// Function is the function configuration for the new function. Once the
	// source code upload has been confirmed the function is created with this
	// configuration.
	Function *Function `json:"function,omitempty"`
}

// InfraType is a target infrastructure to deploy to
type InfraType string

const (
	// InfrastructureTypeAWS is Amazon Web Services
	InfrastructureTypeAWS InfraType = "aws"
)

// Environment is a target deployment environment.
type Environment struct {
	// Name is the unique name for an environment.
	Name string `json:"name,omitempty"`
	// Labels are labels used to identify an environment.
	Labels map[string]string `json:"labels,omitempty"`
	// Infrastructure defines what type the infrastructure type is for the environment.
	Infrastructure InfraType `json:"infrastructure,omitempty"`
	// AWS specifies AWS specific deployment information
	AWS *InfrastructureAWS `json:"aws,omitempty"`
}

// InfrastructureAWS contains information for an AWS deployment
type InfrastructureAWS struct {
	Region string `json:"region,omitempty"`
}

// Deployment represents a connection between functions to environments.
type Deployment struct {
	// Name is the unique name for a deployment.
	Name string `json:"name,omitempty"`
	// EnvironmentLabels is the label seletor for which environment(s) should be
	// the destination of the deployment. The environment must have every label
	// assigned to be included.
	EnvironmentLabels map[string]string `json:"environment_labels,omitempty"`
	// FunctionLabels is the label selector for which function(s) should be part
	// of the deployment. Every label must match for the function to be included.
	FunctionLabels map[string]string `json:"function_labels,omitempty"`
}
