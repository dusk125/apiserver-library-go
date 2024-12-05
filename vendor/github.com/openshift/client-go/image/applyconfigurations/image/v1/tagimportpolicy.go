// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	imagev1 "github.com/openshift/api/image/v1"
)

// TagImportPolicyApplyConfiguration represents a declarative configuration of the TagImportPolicy type for use
// with apply.
type TagImportPolicyApplyConfiguration struct {
	Insecure   *bool                   `json:"insecure,omitempty"`
	Scheduled  *bool                   `json:"scheduled,omitempty"`
	ImportMode *imagev1.ImportModeType `json:"importMode,omitempty"`
}

// TagImportPolicyApplyConfiguration constructs a declarative configuration of the TagImportPolicy type for use with
// apply.
func TagImportPolicy() *TagImportPolicyApplyConfiguration {
	return &TagImportPolicyApplyConfiguration{}
}

// WithInsecure sets the Insecure field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Insecure field is set to the value of the last call.
func (b *TagImportPolicyApplyConfiguration) WithInsecure(value bool) *TagImportPolicyApplyConfiguration {
	b.Insecure = &value
	return b
}

// WithScheduled sets the Scheduled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scheduled field is set to the value of the last call.
func (b *TagImportPolicyApplyConfiguration) WithScheduled(value bool) *TagImportPolicyApplyConfiguration {
	b.Scheduled = &value
	return b
}

// WithImportMode sets the ImportMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImportMode field is set to the value of the last call.
func (b *TagImportPolicyApplyConfiguration) WithImportMode(value imagev1.ImportModeType) *TagImportPolicyApplyConfiguration {
	b.ImportMode = &value
	return b
}
