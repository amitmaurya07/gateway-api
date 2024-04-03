/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// HTTPHeaderFilterApplyConfiguration represents an declarative configuration of the HTTPHeaderFilter type for use
// with apply.
type HTTPHeaderFilterApplyConfiguration struct {
	Set    []HTTPHeaderApplyConfiguration `json:"set,omitempty"`
	Add    []HTTPHeaderApplyConfiguration `json:"add,omitempty"`
	Remove []string                       `json:"remove,omitempty"`
}

// HTTPHeaderFilterApplyConfiguration constructs an declarative configuration of the HTTPHeaderFilter type for use with
// apply.
func HTTPHeaderFilter() *HTTPHeaderFilterApplyConfiguration {
	return &HTTPHeaderFilterApplyConfiguration{}
}

// WithSet adds the given value to the Set field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Set field.
func (b *HTTPHeaderFilterApplyConfiguration) WithSet(values ...*HTTPHeaderApplyConfiguration) *HTTPHeaderFilterApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSet")
		}
		b.Set = append(b.Set, *values[i])
	}
	return b
}

// WithAdd adds the given value to the Add field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Add field.
func (b *HTTPHeaderFilterApplyConfiguration) WithAdd(values ...*HTTPHeaderApplyConfiguration) *HTTPHeaderFilterApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAdd")
		}
		b.Add = append(b.Add, *values[i])
	}
	return b
}

// WithRemove adds the given value to the Remove field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Remove field.
func (b *HTTPHeaderFilterApplyConfiguration) WithRemove(values ...string) *HTTPHeaderFilterApplyConfiguration {
	for i := range values {
		b.Remove = append(b.Remove, values[i])
	}
	return b
}
