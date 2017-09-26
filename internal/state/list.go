package state

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// matcher is used for filtering models based on the their metadata.
type matcher interface {
	match(meta *Meta) bool
}

// LabelMatcher is a matcher that matches labels.
type LabelMatcher struct {
	Labels map[string]string
}

// match returns true if the metadata contains all labels in the LabelMatcher.
// The comparison is case-insensitive.
func (l *LabelMatcher) match(meta *Meta) bool {
	matchLabels := meta.Labels
	if len(matchLabels) == 0 {
		return false
	}

	for k, v := range l.Labels {
		k1 := strings.ToLower(k)
		v1 := strings.ToLower(v)
		match := false
		for k2, v2 := range matchLabels {
			if k1 == strings.ToLower(k2) && v1 == strings.ToLower(v2) {
				match = true
			}
		}
		if !match {
			return false
		}
	}

	return true
}

// matchesAll returns true if all matchers match. Returns true if no matchers
// are passed in.
func matchesAll(meta *Meta, matchers []matcher) bool {
	for _, matcher := range matchers {
		if !matcher.match(meta) {
			return false
		}
	}
	return true
}

// ListDeployments lists deployments. Returns all deployments if no matchers
// are provided.
func ListDeployments(ctx context.Context, kv backend.Lister, matchers ...matcher) ([]*Deployment, error) {
	key := modelListPath(ModelTypeDeployment)

	items, err := kv.List(ctx, key)
	if err != nil {
		return nil, errors.Wrap(err, "could not list deployments")
	}

	result := []*Deployment{}
	for _, raw := range items {
		var deployment Deployment
		if err := json.Unmarshal([]byte(raw), &deployment); err != nil {
			return nil, errors.Wrap(err, "could not unmarshal deployment")
		}
		if matchesAll(&deployment.Meta, matchers) {
			result = append(result, &deployment)
		}
	}

	return result, nil
}

// ListEnvironments lists environments. Returns all environments if no matchers
// are provided.
func ListEnvironments(ctx context.Context, kv backend.Lister, matchers ...matcher) ([]*Environment, error) {
	key := modelListPath(ModelTypeEnvironment)

	items, err := kv.List(ctx, key)
	if err != nil {
		return nil, errors.Wrap(err, "could not list environments")
	}

	result := []*Environment{}
	for _, raw := range items {
		var environment Environment
		if err := json.Unmarshal([]byte(raw), &environment); err != nil {
			return nil, errors.Wrap(err, "could not unmarshal environment")
		}
		if matchesAll(&environment.Meta, matchers) {
			result = append(result, &environment)
		}
	}

	return result, nil
}

// ListFunctions lists functions. Returns all functions if no matchers
// are provided.
func ListFunctions(ctx context.Context, kv backend.Lister, matchers ...matcher) ([]*Function, error) {
	key := modelListPath(ModelTypeFunction)

	items, err := kv.List(ctx, key)
	if err != nil {
		return nil, errors.Wrap(err, "could not list functions")
	}

	result := []*Function{}
	for _, raw := range items {
		var function Function
		if err := json.Unmarshal([]byte(raw), &function); err != nil {
			return nil, errors.Wrap(err, "could not unmarshal function")
		}
		if matchesAll(&function.Meta, matchers) {
			result = append(result, &function)
		}
	}

	return result, nil
}
