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

type EksUserArnInitParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksCluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksCluster in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksCluster in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`
}

type EksUserArnObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EksUserArnParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksCluster in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksCluster in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`
}

// EksUserArnSpec defines the desired state of EksUserArn
type EksUserArnSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EksUserArnParameters `json:"forProvider"`
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
	InitProvider EksUserArnInitParameters `json:"initProvider,omitempty"`
}

// EksUserArnStatus defines the observed state of EksUserArn.
type EksUserArnStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EksUserArnObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EksUserArn is the Schema for the EksUserArns API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type EksUserArn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksUserArnSpec   `json:"spec"`
	Status            EksUserArnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksUserArnList contains a list of EksUserArns
type EksUserArnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksUserArn `json:"items"`
}

// Repository type metadata.
var (
	EksUserArn_Kind             = "EksUserArn"
	EksUserArn_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EksUserArn_Kind}.String()
	EksUserArn_KindAPIVersion   = EksUserArn_Kind + "." + CRDGroupVersion.String()
	EksUserArn_GroupVersionKind = CRDGroupVersion.WithKind(EksUserArn_Kind)
)

func init() {
	SchemeBuilder.Register(&EksUserArn{}, &EksUserArnList{})
}
