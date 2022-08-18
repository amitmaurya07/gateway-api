/*
Copyright 2021 The Kubernetes Authors.

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

package translator

import (
	gatewayv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// HeaderMatchTypePtr translates a string to *PathMatchType
func HeaderMatchTypePtr(s string) *gatewayv1b1.HeaderMatchType {
	result := gatewayv1b1.HeaderMatchType(s)
	return &result
}

// PathMatchTypePtr translates a string to *PathMatchType
func PathMatchTypePtr(s string) *gatewayv1b1.PathMatchType {
	result := gatewayv1b1.PathMatchType(s)
	return &result
}
