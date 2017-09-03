package client

import (
	"encoding/json"
	"fmt"
)

// ResourceType defines the resource type. The type is read from the spec file
// and determines how the resource is interpreted.
type ResourceType string

const (
	// ResourceTypeFunction is the type for a function.
	ResourceTypeFunction ResourceType = "function"
	// ResourceTypeDeployment is the type for a deployment.
	ResourceTypeDeployment ResourceType = "deployment"
)

// Resource is a generic resource on disk. It can represent any resource type.
type Resource interface {
	// File is the configuration file the resource definition was parsed from.
	File() string
	// Meta contains metadata for a resource.
	Meta() *Meta
	// Type defines what type of resource this is.
	Type() ResourceType
	testString() string
}

// Meta contains metadata for a resource.
type Meta struct {
	// Name is the name for a resource. It must be unique among other resource of the same type.
	Name string `json:"name"`
	// Labels are used to identify a resource.
	Labels map[string]string `json:"labels"`
}

// Function is the configuration for a function on disk.
type Function interface {
	Function() *FunctionSpec
}

type functionResource struct {
	file string
	meta *Meta
	spec *FunctionSpec
}

func (f *functionResource) File() string            { return f.file }
func (f *functionResource) Meta() *Meta             { return f.meta }
func (f *functionResource) Type() ResourceType      { return ResourceTypeFunction }
func (f *functionResource) Function() *FunctionSpec { return f.spec }
func (f *functionResource) testString() string {
	meta, _ := json.MarshalIndent(f.Meta(), "", "    ")
	spec, _ := json.MarshalIndent(f.Function(), "", "    ")
	return fmt.Sprintf("function\nfile: %s\nmeta: %s\nspec: %s", f.File(), meta, spec)
}

// FunctionSpec represents a function specification.
type FunctionSpec struct {
	// Runtime is the function runtime.
	Runtime string `json:"runtime"`
	// AWS is the Amazon Web Services specific configuration for the function.
	AWS *FunctionAWSSpec `json:"aws,omitempty"`
}

// FunctionAWSSpec contains AWS function (Lambda) specific configuration info.
type FunctionAWSSpec struct {
	// Timeout is the timeout in seconds for the function
	Timeout int64 `json:"timeout,omitempty"`
	// Memory is the memory in mb for the function
	Memory int64 `json:"memory,omitempty"`
}

// Deployment is the configuration for a deployment on disk.
type Deployment interface {
	Deployment() *DeploymentSpec
}

type deploymentResource struct {
	file string
	meta *Meta
	spec *DeploymentSpec
}

func (d *deploymentResource) File() string                { return d.file }
func (d *deploymentResource) Meta() *Meta                 { return d.meta }
func (d *deploymentResource) Type() ResourceType          { return ResourceTypeDeployment }
func (d *deploymentResource) Deployment() *DeploymentSpec { return d.spec }
func (d *deploymentResource) testString() string {
	meta, _ := json.MarshalIndent(d.Meta(), "", "    ")
	spec, _ := json.MarshalIndent(d.Deployment(), "", "    ")
	return fmt.Sprintf("deployment\nfile: %s\nmeta: %s\nspec: %s", d.File(), meta, spec)
}

// DeploymentSpec represents a deployment.
type DeploymentSpec struct {
	// EnvironmentLabels is the label seletor for which environment(s) should be
	// the destination of the deployment. The environment must have every label
	// assigned to be included.
	EnvironmentLabels map[string]string `json:"environment"`
	// FunctionLabels is the label selector for which function(s) should be part
	// of the deployment. Every label must match for the function to be included.
	FunctionLabels map[string]string `json:"function"`
}
