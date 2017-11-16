package state

// ModelType defines the type of a model. It is used to group the same
// models in the backend.
type ModelType string

const (
	// ModelTypeFunction is a function.
	ModelTypeFunction ModelType = "function"
	// ModelTypeEnvironment is a target deployment environment
	ModelTypeEnvironment ModelType = "environment"
	// ModelTypeDeployment is a target deployment environment
	ModelTypeDeployment ModelType = "deployment"
)

// InfraType is a target infrastructure to deploy to
type InfraType string

const (
	// InfrastructureTypeAWS is Amazon Web Services
	InfrastructureTypeAWS InfraType = "aws"
)

// Model is a a generic model.
type Model interface {
	// Name returns a unique name to identify the model. The name is unique
	// within the model type, not necessarily globally unique.
	Name() string
}

// Meta contains metadata for a model.
type Meta struct {
	// Name is the name for a model. It must be unique among other models of
	// the same type.
	Name string
	// Labels are used to identify a model.
	Labels map[string]string
}

// Function represents a function specification.
type Function struct {
	// Meta contains function metadata.
	Meta Meta
	// Runtime is the function runtime.
	Runtime string
	// Checksum is the checksum calculated by the client of the source files for.
	// the function
	Checksum string
	// SourceFilename is the name of the source file. For functions being created
	// this is ignored, it is set when the source has been confirmed.
	SourceFilename string
	// AWS is the Amazon Web Services specific configuration for the function.
	AWS *FunctionAWS
}

// FunctionAWS contains AWS function (Lambda) specific configuration info.
type FunctionAWS struct {
	// Timeout is the timeout in seconds for the function.
	Timeout int64
	// Memory is the memory in mb for the function.
	Memory int64
}

// Name returns a unique name to identify the function.
func (f *Function) Name() string { return f.Meta.Name }

// PendingUpload is a source code request that has been returned to the client.
// When source is confirmed it is used to fetch the source and apply function
// changes.
type PendingUpload struct {
	// Filename is the filename to retrieve the source by from the filestore.
	Filename string
	// PreviousFilename is the filename of the previous source in case the
	// function code has been updated. It is blank if the function did not exist
	// before.
	PreviousFilename string
	// Function is the function configuration for the new function. Once the
	// source code upload has been confirmed the function is created with this
	// configuration.
	Function *Function
}

// Environment is a target deployment environment.
type Environment struct {
	// Meta contains environment metadata.
	Meta Meta
	// Infrastructure defines what type the infrastructure type is for the environment.
	Infrastructure InfraType
	// AWS specifies AWS specific deployment information
	AWS *InfrastructureAWS
}

// Name returns a unique name to identify the environment.
func (e *Environment) Name() string { return e.Meta.Name }

// InfrastructureAWS contains information for an AWS deployment
type InfrastructureAWS struct {
	Region string
}

// Deployment represents a connection between functions to environments.
type Deployment struct {
	// Meta contains deployment metadata.
	Meta Meta
	// EnvironmentLabels is the label seletor for which environment(s) should be
	// the destination of the deployment. The environment must have every label
	// assigned to be included.
	EnvironmentLabels map[string]string
	// FunctionLabels is the label selector for which function(s) should be part
	// of the deployment. Every label must match for the function to be included.
	FunctionLabels map[string]string
}

// Name returns a unique name to identify the deployment.
func (d *Deployment) Name() string { return d.Meta.Name }
