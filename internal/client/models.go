package client

import (
	"encoding/json"
	"fmt"
)

// ModelType defines the model type. The type is read from the spec file
// and determines how the model is interpreted.
type ModelType string

const (
	// ModelTypeFunction is the type for a function.
	ModelTypeFunction ModelType = "function"
	// ModelTypeDeployment is the type for a deployment.
	ModelTypeDeployment ModelType = "deployment"
)

// Model is a generic model on disk. It can represent any model type.
type Model interface {
	// File is the configuration file the model definition was parsed from.
	File() string
	// Meta contains metadata for a model.
	Meta() *Meta
	// Type defines what type of model this is.
	Type() ModelType
	testString() string
}

// Meta contains metadata for a model.
type Meta struct {
	// Name is the name for a model. It must be unique among other model of the same type.
	Name string `json:"name"`
	// Labels are used to identify a model.
	Labels map[string]string `json:"labels"`
}

// Function is the configuration for a function on disk.
type Function interface {
	Function() *FunctionSpec
}

type functionModel struct {
	file string
	meta *Meta
	spec *FunctionSpec
}

func (f *functionModel) File() string            { return f.file }
func (f *functionModel) Meta() *Meta             { return f.meta }
func (f *functionModel) Type() ModelType         { return ModelTypeFunction }
func (f *functionModel) Function() *FunctionSpec { return f.spec }
func (f *functionModel) testString() string {
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

type deploymentModel struct {
	file string
	meta *Meta
	spec *DeploymentSpec
}

func (d *deploymentModel) File() string                { return d.file }
func (d *deploymentModel) Meta() *Meta                 { return d.meta }
func (d *deploymentModel) Type() ModelType             { return ModelTypeDeployment }
func (d *deploymentModel) Deployment() *DeploymentSpec { return d.spec }
func (d *deploymentModel) testString() string {
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
