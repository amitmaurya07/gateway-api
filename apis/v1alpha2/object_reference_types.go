/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha2

// LocalObjectReference identifies an API object within the namespace of the
// referrer.
type LocalObjectReference struct {
	// Group is the group of the referent.
	//
	// +kubebuilder:validation:MaxLength=253
	Group string `json:"group"`

	// Kind is kind of the referent.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Kind string `json:"kind"`

	// Name is the name of the referent.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

// ObjectReference identifies an API object including its namespace.
type ObjectReference struct {
	// Group is the group of the referent.
	// When unspecified (empty string), core API group is inferred.
	//
	// +optional
	// +kubebuilder:default=""
	// +kubebuilder:validation:MaxLength=253
	Group *string `json:"group"`

	// Kind is kind of the referent.
	//
	// +optional
	// +kubebuilder:default=Service
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Kind *string `json:"kind"`

	// Name is the name of the referent.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`

	// Namespace is the namespace of the backend. When unspecified, the local
	// namespace is inferred.
	//
	// Note that when a namespace is specified, a ReferencePolicy object
	// is required in the referent namespace to allow that namespace's
	// owner to accept the reference. See the ReferencePolicy documentation
	// for details.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +optional
	Namespace *string `json:"namespace,omitempty"`
}

// BackendObjectReference defines how an ObjectReference that is
// specific to BackendRef. It includes a few additional fields and features
// than a regular ObjectReference.
//
// Note that when a namespace is specified, a ReferencePolicy object
// is required in the referent namespace to allow that namespace's
// owner to accept the reference. See the ReferencePolicy documentation
// for details.
type BackendObjectReference struct {
	// Group is the group of the referent.
	// When unspecified (empty string), core API group is inferred.
	//
	// +optional
	// +kubebuilder:default=""
	// +kubebuilder:validation:MaxLength=253
	Group *string `json:"group,omitempty"`

	// Kind is kind of the referent.
	//
	// +optional
	// +kubebuilder:default=Service
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Kind *string `json:"kind,omitempty"`

	// Name is the name of the referent.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`

	// Namespace is the namespace of the backend. When unspecified, the local
	// namespace is inferred.
	//
	// Note that when a namespace is specified, a ReferencePolicy object
	// is required in the referent namespace to allow that namespace's
	// owner to accept the reference. See the ReferencePolicy documentation
	// for details.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +optional
	Namespace *string `json:"namespace,omitempty"`

	// Port specifies the destination port number to use for this resource.
	// Port is required when the referent is a Kubernetes Service.
	// For other resources, destination port can be derived from the referent
	// resource or this field.
	//
	// +optional
	Port *PortNumber `json:"port,omitempty"`
}
