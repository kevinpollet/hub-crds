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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/traefik/hub-crds/hub/v1alpha1/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// APIs returns a APIInformer.
	APIs() APIInformer
	// APIAccesses returns a APIAccessInformer.
	APIAccesses() APIAccessInformer
	// APICollections returns a APICollectionInformer.
	APICollections() APICollectionInformer
	// APIGateways returns a APIGatewayInformer.
	APIGateways() APIGatewayInformer
	// APIPortals returns a APIPortalInformer.
	APIPortals() APIPortalInformer
	// APIRateLimits returns a APIRateLimitInformer.
	APIRateLimits() APIRateLimitInformer
	// APIVersions returns a APIVersionInformer.
	APIVersions() APIVersionInformer
	// AccessControlPolicies returns a AccessControlPolicyInformer.
	AccessControlPolicies() AccessControlPolicyInformer
	// EdgeIngresses returns a EdgeIngressInformer.
	EdgeIngresses() EdgeIngressInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// APIs returns a APIInformer.
func (v *version) APIs() APIInformer {
	return &aPIInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIAccesses returns a APIAccessInformer.
func (v *version) APIAccesses() APIAccessInformer {
	return &aPIAccessInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// APICollections returns a APICollectionInformer.
func (v *version) APICollections() APICollectionInformer {
	return &aPICollectionInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// APIGateways returns a APIGatewayInformer.
func (v *version) APIGateways() APIGatewayInformer {
	return &aPIGatewayInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// APIPortals returns a APIPortalInformer.
func (v *version) APIPortals() APIPortalInformer {
	return &aPIPortalInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// APIRateLimits returns a APIRateLimitInformer.
func (v *version) APIRateLimits() APIRateLimitInformer {
	return &aPIRateLimitInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// APIVersions returns a APIVersionInformer.
func (v *version) APIVersions() APIVersionInformer {
	return &aPIVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AccessControlPolicies returns a AccessControlPolicyInformer.
func (v *version) AccessControlPolicies() AccessControlPolicyInformer {
	return &accessControlPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// EdgeIngresses returns a EdgeIngressInformer.
func (v *version) EdgeIngresses() EdgeIngressInformer {
	return &edgeIngressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
