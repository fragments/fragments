package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// Load loads a resource from a file. The file can be a yaml or json and can
// contain one or more resources.
func Load(file string) ([]Resource, error) {
	// Open file and read it to memory
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.Wrap(err, "could not load file")
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// split possible multiple resources defined in file
	docs, err := split(data)
	if err != nil {
		return nil, err
	}

	// parse resources
	out := []Resource{}
	for _, d := range docs {
		r, err := parse(d, file)
		if err != nil {
			return nil, err
		}
		if r != nil {
			out = append(out, r)
		}
	}

	return out, nil
}

// split extracts individual resources from a byte array and returns the json
// representation of each resource.
// Supports loading a single json object, a json array, single yaml object and
// multiple documents in a yaml file. For yaml documents the result is
// converted to json.
func split(data []byte) ([][]byte, error) {
	out := [][]byte{}
	trimmed := bytes.TrimFunc(data, unicode.IsSpace)

	if bytes.HasPrefix(trimmed, []byte("{")) {
		// Input is a simple json object, only one resrouce returned
		out = [][]byte{data}
		return out, nil
	}

	if bytes.HasPrefix(trimmed, []byte("[")) {
		// Input is a json array, split it and return multiple values
		var arr []json.RawMessage
		if err := json.Unmarshal(data, &arr); err != nil {
			return nil, err
		}

		for _, j := range arr {
			out = append(out, j)
		}
		return out, nil
	}

	// Input is probably yaml, possibly containing multiple documents
	split := bytes.Split(trimmed, []byte("\n---\n"))
	for _, data := range split {
		jsonData, err := yaml.YAMLToJSON(data)
		if err != nil {
			return nil, err
		}
		out = append(out, jsonData)
	}

	return out, nil
}

// parse parses a generic resource to a specific type. Returns nil if the
// target doesn't look like a valid resource (missing type)
func parse(data []byte, filepath string) (Resource, error) {
	raw, err := parseRaw(data)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse resource")
	}
	if raw.Type == "" {
		// No type defined, don't consider this as a valid resource
		return nil, nil
	}
	if raw.Meta == nil {
		return nil, errors.New("resource meta not set")
	}
	if raw.Meta.Name == "" {
		return nil, errors.New("resource name not set")
	}

	switch strings.ToLower(raw.Type) {
	case "function":
		return parseFunction(raw, filepath)
	default:
		return nil, errors.Errorf("unknown resource type %s", raw.Type)
	}
}

// rawResource represents a generic resource
type rawResource struct {
	Type string
	Meta *Meta
	Spec json.RawMessage
}

// parseRaw parses a raw repsentation of a resource definition. This is used to
// further process the resource, depending on its type
func parseRaw(data []byte) (*rawResource, error) {
	var raw *rawResource
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

// parseFunction parses a function spec
func parseFunction(raw *rawResource, filepath string) (*functionResource, error) {
	f := &functionResource{
		file: filepath,
		meta: raw.Meta,
		spec: &FunctionSpec{},
	}

	if err := json.Unmarshal(raw.Spec, f.spec); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal function resource")
	}

	return f, nil
}
