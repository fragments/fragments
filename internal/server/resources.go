package server

// Meta contains metadata for a resource
type Meta struct {
	// Name is the name for a resource. It must be unique among other resource of the same type
	Name string
	// Labels are used to identify a resource
	Labels map[string]string
}

// Function represents a function specification
type Function struct {
	// Meta contains function metadata
	Meta Meta
	// Runtime is the function runtime
	Runtime string
	// AWS is the Amazon Web Services specific configuration for the function
	AWS *FunctionAWS
}

// FunctionAWSSpec contains AWS function (Lambda) specific configuration info
type FunctionAWS struct {
	// Timeout is the timeout in seconds for the function
	Timeout int64
	// Memory is the memory in mb for the function
	Memory int64
}
