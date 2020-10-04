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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GatewayAllowType specifies which Gateways should be allowed to use a Route.
type GatewayAllowType string

const (
	// GatewayAllowAll indicates that all Gateways will be able to use this
	// route.
	GatewayAllowAll GatewayAllowType = "All"
	// GatewayAllowFromList indicates that only Gateways that have been
	// specified in GatewayRefs will be able to use this route.
	GatewayAllowFromList GatewayAllowType = "FromList"
	// GatewayAllowSameNamespace indicates that only Gateways within the same
	// namespace will be able to use this route.
	GatewayAllowSameNamespace GatewayAllowType = "SameNamespace"
)

// RouteGateways defines which Gateways will be able to use a route. If this
// field results in preventing the selection of a Route by a Gateway, an
// "Admitted" condition with a status of false must be set for the Gateway on
// that Route.
type RouteGateways struct {
	// Allow indicates which Gateways will be allowed to use this route.
	// Possible values are:
	// * All: Gateways in any namespace can use this route.
	// * FromList: Only Gateways specified in GatewayRefs may use this route.
	// * SameNamespace: Only Gateways in the same namespace may use this route.
	// +kubebuilder:validation:Enum=All;FromList;SameNamespace
	// +kubebuilder:default=SameNamespace
	Allow GatewayAllowType `json:"allow,omitempty"`
	// GatewayRefs must be specified when Allow is set to "FromList". In that
	// case, only Gateways referenced in this list will be allowed to use this
	// route. This field is ignored for other values of "Allow".
	// +optional
	GatewayRefs []GatewayReference `json:"gatewayRefs,omitempty"`
}

// GatewayReference identifies a Gateway in a specified namespace.
type GatewayReference struct {
	// Name is the name of the referent.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
	// Namespace is the namespace of the referent.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Namespace string `json:"namespace"`
}

// RouteForwardTo defines how a Route should forward a request.
type RouteForwardTo struct {
	// ServiceName refers to the name of the Service to forward matched requests
	// to. When specified, this takes the place of BackendRef. If both
	// BackendRef and ServiceName are specified, ServiceName will be given
	// precedence. If the referent cannot be found, controllers must set the
	// "InvalidRoutes" status condition on any Gateway that includes this
	// Route to true.
	//
	// Support: Core
	//
	// +optional
	// +kubebuilder:validation:MaxLength=253
	ServiceName *string `json:"serviceName,omitempty"`

	// BackendRef is a reference to a backend to forward matched requests to. If
	// both BackendRef and ServiceName are specified, ServiceName will be given
	// precedence. If the referent cannot be found, controllers must set the
	// "InvalidRoutes" status condition on any Gateway that includes this Route
	// to true.
	//
	// Support: Custom
	//
	// +optional
	BackendRef *LocalObjectReference `json:"backendRef,omitempty"`

	// Port specifies the destination port number to use for the backend
	// referenced by the ServiceName or BackendRef field. If unspecified and a
	// Service object consisting of a single port definition is the backend,
	// that port will be used. If unspecified and the backend is a Service
	// object consisting of multiple port definitions, controllers must set the
	// "InvalidRoutes" status condition on any Gateway that includes this Route
	// to true.
	//
	// Support: Core
	//
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port *int32 `json:"port,omitempty"`

	// Weight specifies the proportion of traffic forwarded to the backend
	// referenced by the ServiceName or BackendRef field. computed as
	// weight/(sum of all weights in this ForwardTo list). Weight is not a
	// percentage and the sum of weights does not need to equal 100. If only one
	// backend is specified, 100% of the traffic is forwarded to that backend.
	// If unspecified, weight defaults to 1.
	//
	// Support: Extended
	//
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10000
	Weight int32 `json:"weight,omitempty"`
}

// RouteConditionType is a type of condition for a route.
type RouteConditionType string

const (
	// ConditionRouteAdmitted indicates whether the route has been admitted
	// or rejected by a Gateway, and why.
	ConditionRouteAdmitted RouteConditionType = "Admitted"
)

// RouteGatewayStatus describes the status of a route with respect to an
// associated Gateway.
type RouteGatewayStatus struct {
	// GatewayRef is a reference to a Gateway object that is associated with
	// the route.
	GatewayRef GatewayReference `json:"gatewayRef"`
	// Conditions describes the status of the route with respect to the
	// Gateway.  For example, the "Admitted" condition indicates whether the
	// route has been admitted or rejected by the Gateway, and why.  Note
	// that the route's availability is also subject to the Gateway's own
	// status conditions and listener status.
	//
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// RouteStatus defines the observed state that is required across
// all route types.
type RouteStatus struct {
	// Gateways is a list of the Gateways that are associated with the
	// route, and the status of the route with respect to each of these
	// Gateways. When a Gateway selects this route, the controller that
	// manages the Gateway should add an entry to this list when the
	// controller first sees the route and should update the entry as
	// appropriate when the route is modified.
	Gateways []RouteGatewayStatus `json:"gateways"`
}
