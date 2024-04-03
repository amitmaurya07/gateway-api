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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
	internal "sigs.k8s.io/gateway-api/apis/applyconfiguration/internal"
	apisv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// GRPCRouteApplyConfiguration represents an declarative configuration of the GRPCRoute type for use
// with apply.
type GRPCRouteApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *GRPCRouteSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *GRPCRouteStatusApplyConfiguration `json:"status,omitempty"`
}

// GRPCRoute constructs an declarative configuration of the GRPCRoute type for use with
// apply.
func GRPCRoute(name, namespace string) *GRPCRouteApplyConfiguration {
	b := &GRPCRouteApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("GRPCRoute")
	b.WithAPIVersion("gateway.networking.k8s.io/v1alpha2")
	return b
}

// ExtractGRPCRoute extracts the applied configuration owned by fieldManager from
// gRPCRoute. If no managedFields are found in gRPCRoute for fieldManager, a
// GRPCRouteApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// gRPCRoute must be a unmodified GRPCRoute API object that was retrieved from the Kubernetes API.
// ExtractGRPCRoute provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractGRPCRoute(gRPCRoute *apisv1alpha2.GRPCRoute, fieldManager string) (*GRPCRouteApplyConfiguration, error) {
	return extractGRPCRoute(gRPCRoute, fieldManager, "")
}

// ExtractGRPCRouteStatus is the same as ExtractGRPCRoute except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractGRPCRouteStatus(gRPCRoute *apisv1alpha2.GRPCRoute, fieldManager string) (*GRPCRouteApplyConfiguration, error) {
	return extractGRPCRoute(gRPCRoute, fieldManager, "status")
}

func extractGRPCRoute(gRPCRoute *apisv1alpha2.GRPCRoute, fieldManager string, subresource string) (*GRPCRouteApplyConfiguration, error) {
	b := &GRPCRouteApplyConfiguration{}
	err := managedfields.ExtractInto(gRPCRoute, internal.Parser().Type("io.k8s.sigs.gateway-api.apis.v1alpha2.GRPCRoute"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(gRPCRoute.Name)
	b.WithNamespace(gRPCRoute.Namespace)

	b.WithKind("GRPCRoute")
	b.WithAPIVersion("gateway.networking.k8s.io/v1alpha2")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithKind(value string) *GRPCRouteApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithAPIVersion(value string) *GRPCRouteApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithName(value string) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithGenerateName(value string) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithNamespace(value string) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithUID(value types.UID) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithResourceVersion(value string) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithGeneration(value int64) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithCreationTimestamp(value metav1.Time) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *GRPCRouteApplyConfiguration) WithLabels(entries map[string]string) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *GRPCRouteApplyConfiguration) WithAnnotations(entries map[string]string) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *GRPCRouteApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *GRPCRouteApplyConfiguration) WithFinalizers(values ...string) *GRPCRouteApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *GRPCRouteApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithSpec(value *GRPCRouteSpecApplyConfiguration) *GRPCRouteApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *GRPCRouteApplyConfiguration) WithStatus(value *GRPCRouteStatusApplyConfiguration) *GRPCRouteApplyConfiguration {
	b.Status = value
	return b
}
