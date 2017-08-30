package state

// Meta contains metadata for a resource.
type Meta struct {
	// Name is the name for a resource. It must be unique among other resource of
	// the same type.
	Name string
	// Labels are used to identify a resource.
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
	Infrastructure InfrastructureType
}
