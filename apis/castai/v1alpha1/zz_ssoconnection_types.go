/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AadInitParameters struct {

	// (String) Azure AD domain
	// Azure AD domain
	AdDomain *string `json:"adDomain,omitempty" tf:"ad_domain,omitempty"`

	// (String) Azure AD client ID
	// Azure AD client ID
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`
}

type AadObservation struct {

	// (String) Azure AD domain
	// Azure AD domain
	AdDomain *string `json:"adDomain,omitempty" tf:"ad_domain,omitempty"`

	// (String) Azure AD client ID
	// Azure AD client ID
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`
}

type AadParameters struct {

	// (String) Azure AD domain
	// Azure AD domain
	// +kubebuilder:validation:Optional
	AdDomain *string `json:"adDomain" tf:"ad_domain,omitempty"`

	// (String) Azure AD client ID
	// Azure AD client ID
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// (String, Sensitive) Azure AD client secret
	// Azure AD client secret
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`
}

type OktaInitParameters struct {

	// (String) Azure AD client ID
	// Okta client ID
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (String) Okta domain
	// Okta domain
	OktaDomain *string `json:"oktaDomain,omitempty" tf:"okta_domain,omitempty"`
}

type OktaObservation struct {

	// (String) Azure AD client ID
	// Okta client ID
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (String) Okta domain
	// Okta domain
	OktaDomain *string `json:"oktaDomain,omitempty" tf:"okta_domain,omitempty"`
}

type OktaParameters struct {

	// (String) Azure AD client ID
	// Okta client ID
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// (String, Sensitive) Azure AD client secret
	// Okta client secret
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// (String) Okta domain
	// Okta domain
	// +kubebuilder:validation:Optional
	OktaDomain *string `json:"oktaDomain" tf:"okta_domain,omitempty"`
}

type SSOConnectionInitParameters struct {

	// (Block List, Max: 1) Azure AD connector (see below for nested schema)
	// Azure AD connector
	Aad []AadInitParameters `json:"aad,omitempty" tf:"aad,omitempty"`

	// (String) Email domain of the connection
	// Email domain of the connection
	EmailDomain *string `json:"emailDomain,omitempty" tf:"email_domain,omitempty"`

	// (String) Connection name
	// Connection name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Okta connector (see below for nested schema)
	// Okta connector
	Okta []OktaInitParameters `json:"okta,omitempty" tf:"okta,omitempty"`
}

type SSOConnectionObservation struct {

	// (Block List, Max: 1) Azure AD connector (see below for nested schema)
	// Azure AD connector
	Aad []AadObservation `json:"aad,omitempty" tf:"aad,omitempty"`

	// (String) Email domain of the connection
	// Email domain of the connection
	EmailDomain *string `json:"emailDomain,omitempty" tf:"email_domain,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Connection name
	// Connection name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Okta connector (see below for nested schema)
	// Okta connector
	Okta []OktaObservation `json:"okta,omitempty" tf:"okta,omitempty"`
}

type SSOConnectionParameters struct {

	// (Block List, Max: 1) Azure AD connector (see below for nested schema)
	// Azure AD connector
	// +kubebuilder:validation:Optional
	Aad []AadParameters `json:"aad,omitempty" tf:"aad,omitempty"`

	// (String) Email domain of the connection
	// Email domain of the connection
	// +kubebuilder:validation:Optional
	EmailDomain *string `json:"emailDomain,omitempty" tf:"email_domain,omitempty"`

	// (String) Connection name
	// Connection name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Okta connector (see below for nested schema)
	// Okta connector
	// +kubebuilder:validation:Optional
	Okta []OktaParameters `json:"okta,omitempty" tf:"okta,omitempty"`
}

// SSOConnectionSpec defines the desired state of SSOConnection
type SSOConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSOConnectionParameters `json:"forProvider"`
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
	InitProvider SSOConnectionInitParameters `json:"initProvider,omitempty"`
}

// SSOConnectionStatus defines the observed state of SSOConnection.
type SSOConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSOConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SSOConnection is the Schema for the SSOConnections API. SSO Connection resource allows creating SSO trust relationship with CAST AI.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type SSOConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.emailDomain) || (has(self.initProvider) && has(self.initProvider.emailDomain))",message="spec.forProvider.emailDomain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SSOConnectionSpec   `json:"spec"`
	Status SSOConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSOConnectionList contains a list of SSOConnections
type SSOConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSOConnection `json:"items"`
}

// Repository type metadata.
var (
	SSOConnection_Kind             = "SSOConnection"
	SSOConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSOConnection_Kind}.String()
	SSOConnection_KindAPIVersion   = SSOConnection_Kind + "." + CRDGroupVersion.String()
	SSOConnection_GroupVersionKind = CRDGroupVersion.WithKind(SSOConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&SSOConnection{}, &SSOConnectionList{})
}
