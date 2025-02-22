/*
Copyright 2022 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GlobalaccountTrustConfigurationInitParameters struct {

	// (String) Description of the trust configuration.
	// Description of the trust configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The tenant's domain which should be used for user logon.
	// The tenant's domain which should be used for user logon.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) The name of the Identity Authentication tenant that you want to connect to the global account.
	// The name of the Identity Authentication tenant that you want to connect to the global account.
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// (String) The display name of the trust configuration.
	// The display name of the trust configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The origin of the identity provider.
	// The origin of the identity provider.
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`
}

type GlobalaccountTrustConfigurationObservation struct {

	// (String) Description of the trust configuration.
	// Description of the trust configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The tenant's domain which should be used for user logon.
	// The tenant's domain which should be used for user logon.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String, Deprecated) The origin of the identity provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the Identity Authentication tenant that you want to connect to the global account.
	// The name of the Identity Authentication tenant that you want to connect to the global account.
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// (String) The display name of the trust configuration.
	// The display name of the trust configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The origin of the identity provider.
	// The origin of the identity provider.
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// (String) The protocol used to establish trust with the identity provider.
	// The protocol used to establish trust with the identity provider.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Boolean) Shows whether the trust configuration can be modified.
	// Shows whether the trust configuration can be modified.
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// (String) Determines whether the identity provider is currently 'active' or 'inactive'.
	// Determines whether the identity provider is currently 'active' or 'inactive'.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String) The trust type.
	// The trust type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GlobalaccountTrustConfigurationParameters struct {

	// (String) Description of the trust configuration.
	// Description of the trust configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The tenant's domain which should be used for user logon.
	// The tenant's domain which should be used for user logon.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) The name of the Identity Authentication tenant that you want to connect to the global account.
	// The name of the Identity Authentication tenant that you want to connect to the global account.
	// +kubebuilder:validation:Optional
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// (String) The display name of the trust configuration.
	// The display name of the trust configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The origin of the identity provider.
	// The origin of the identity provider.
	// +kubebuilder:validation:Optional
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`
}

// GlobalaccountTrustConfigurationSpec defines the desired state of GlobalaccountTrustConfiguration
type GlobalaccountTrustConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalaccountTrustConfigurationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider GlobalaccountTrustConfigurationInitParameters `json:"initProvider,omitempty"`
}

// GlobalaccountTrustConfigurationStatus defines the observed state of GlobalaccountTrustConfiguration.
type GlobalaccountTrustConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalaccountTrustConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GlobalaccountTrustConfiguration is the Schema for the GlobalaccountTrustConfigurations API. Establishes trust from a global account to an Identity Authentication tenant. Tip: You must be assigned to the admin role of the global account. Further documentation: https://help.sap.com/docs/btp/sap-business-technology-platform/trust-and-federation-with-identity-providers
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,account}
type GlobalaccountTrustConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityProvider) || (has(self.initProvider) && has(self.initProvider.identityProvider))",message="spec.forProvider.identityProvider is a required parameter"
	Spec   GlobalaccountTrustConfigurationSpec   `json:"spec"`
	Status GlobalaccountTrustConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalaccountTrustConfigurationList contains a list of GlobalaccountTrustConfigurations
type GlobalaccountTrustConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalaccountTrustConfiguration `json:"items"`
}

// Repository type metadata.
var (
	GlobalaccountTrustConfiguration_Kind             = "GlobalaccountTrustConfiguration"
	GlobalaccountTrustConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalaccountTrustConfiguration_Kind}.String()
	GlobalaccountTrustConfiguration_KindAPIVersion   = GlobalaccountTrustConfiguration_Kind + "." + CRDGroupVersion.String()
	GlobalaccountTrustConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(GlobalaccountTrustConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalaccountTrustConfiguration{}, &GlobalaccountTrustConfigurationList{})
}
