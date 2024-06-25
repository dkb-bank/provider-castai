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

type AffinityInitParameters struct {

	// Key of the node affinity selector.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Operator of the node affinity selector. Allowed values: In, NotIn, Exists, DoesNotExist, Gt, Lt.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Values of the node affinity selector.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AffinityObservation struct {

	// Key of the node affinity selector.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Operator of the node affinity selector. Allowed values: In, NotIn, Exists, DoesNotExist, Gt, Lt.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Values of the node affinity selector.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AffinityParameters struct {

	// Key of the node affinity selector.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Operator of the node affinity selector. Allowed values: In, NotIn, Exists, DoesNotExist, Gt, Lt.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Values of the node affinity selector.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConstraintsInitParameters struct {

	// List of acceptable instance CPU architectures, the default is amd64. Allowed values: amd64, arm64.
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// The list of AZ names to consider for the node template, if empty or not set all AZs are considered.
	Azs []*string `json:"azs,omitempty" tf:"azs,omitempty"`

	// Will only include burstable instances when enabled otherwise they will be excluded. Supported values: `enabled`, `disabled` or “.
	BurstableInstances *string `json:"burstableInstances,omitempty" tf:"burstable_instances,omitempty"`

	// Compute optimized instance constraint (deprecated).
	ComputeOptimized *bool `json:"computeOptimized,omitempty" tf:"compute_optimized,omitempty"`

	// Will only include compute optimized nodes when enabled and exclude compute optimized nodes when disabled. Empty value won't have effect on instances filter. Supported values: `enabled`, `disabled` or empty string.
	ComputeOptimizedState *string `json:"computeOptimizedState,omitempty" tf:"compute_optimized_state,omitempty"`

	CustomPriority []CustomPriorityInitParameters `json:"customPriority,omitempty" tf:"custom_priority,omitempty"`

	// Dedicated node affinity - creates preference for instances to be created on sole tenancy or dedicated nodes. This
	// feature is only available for GCP clusters and sole tenancy nodes with local
	// SSDs or GPUs are not supported. If the sole tenancy or dedicated nodes don't have capacity for selected instance
	// type, the Autoscaler will fall back to multi-tenant instance types available for this Node Template.
	// Other instance constraints are applied when the Autoscaler picks available instance types that can be created on
	// the sole tenancy or dedicated node (example: setting min CPU to 16).
	DedicatedNodeAffinity []DedicatedNodeAffinityInitParameters `json:"dedicatedNodeAffinity,omitempty" tf:"dedicated_node_affinity,omitempty"`

	// Enable/disable spot diversity policy. When enabled, autoscaler will try to balance between diverse and cost optimal instance types.
	EnableSpotDiversity *bool `json:"enableSpotDiversity,omitempty" tf:"enable_spot_diversity,omitempty"`

	// Fallback restore rate in seconds: defines how much time should pass before spot fallback should be attempted to be restored to real spot.
	FallbackRestoreRateSeconds *float64 `json:"fallbackRestoreRateSeconds,omitempty" tf:"fallback_restore_rate_seconds,omitempty"`

	Gpu []GpuInitParameters `json:"gpu,omitempty" tf:"gpu,omitempty"`

	InstanceFamilies []InstanceFamiliesInitParameters `json:"instanceFamilies,omitempty" tf:"instance_families,omitempty"`

	// GPU instance constraint - will only pick nodes with GPU if true
	IsGpuOnly *bool `json:"isGpuOnly,omitempty" tf:"is_gpu_only,omitempty"`

	// Max CPU cores per node.
	MaxCPU *float64 `json:"maxCpu,omitempty" tf:"max_cpu,omitempty"`

	// Max Memory (Mib) per node.
	MaxMemory *float64 `json:"maxMemory,omitempty" tf:"max_memory,omitempty"`

	// Min CPU cores per node.
	MinCPU *float64 `json:"minCpu,omitempty" tf:"min_cpu,omitempty"`

	// Min Memory (Mib) per node.
	MinMemory *float64 `json:"minMemory,omitempty" tf:"min_memory,omitempty"`

	// Should include on-demand instances in the considered pool.
	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand,omitempty"`

	// List of acceptable instance Operating Systems, the default is linux. Allowed values: linux, windows.
	Os []*string `json:"os,omitempty" tf:"os,omitempty"`

	// Should include spot instances in the considered pool.
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`

	// Allowed node configuration price increase when diversifying instance types. E.g. if the value is 10%, then the overall price of diversified instance types can be 10% higher than the price of the optimal configuration.
	SpotDiversityPriceIncreaseLimitPercent *float64 `json:"spotDiversityPriceIncreaseLimitPercent,omitempty" tf:"spot_diversity_price_increase_limit_percent,omitempty"`

	// Enable/disable spot interruption predictions.
	SpotInterruptionPredictionsEnabled *bool `json:"spotInterruptionPredictionsEnabled,omitempty" tf:"spot_interruption_predictions_enabled,omitempty"`

	// Spot interruption predictions type. Can be either "aws-rebalance-recommendations" or "interruption-predictions".
	SpotInterruptionPredictionsType *string `json:"spotInterruptionPredictionsType,omitempty" tf:"spot_interruption_predictions_type,omitempty"`

	// Storage optimized instance constraint (deprecated).
	StorageOptimized *bool `json:"storageOptimized,omitempty" tf:"storage_optimized,omitempty"`

	// Storage optimized instance constraint - will only pick storage optimized nodes if enabled and won't pick if disabled. Empty value will have no effect. Supported values: `enabled`, `disabled` or empty string.
	StorageOptimizedState *string `json:"storageOptimizedState,omitempty" tf:"storage_optimized_state,omitempty"`

	// Spot instance fallback constraint - when true, on-demand instances will be created, when spots are unavailable.
	UseSpotFallbacks *bool `json:"useSpotFallbacks,omitempty" tf:"use_spot_fallbacks,omitempty"`
}

type ConstraintsObservation struct {

	// List of acceptable instance CPU architectures, the default is amd64. Allowed values: amd64, arm64.
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// The list of AZ names to consider for the node template, if empty or not set all AZs are considered.
	Azs []*string `json:"azs,omitempty" tf:"azs,omitempty"`

	// Will only include burstable instances when enabled otherwise they will be excluded. Supported values: `enabled`, `disabled` or “.
	BurstableInstances *string `json:"burstableInstances,omitempty" tf:"burstable_instances,omitempty"`

	// Compute optimized instance constraint (deprecated).
	ComputeOptimized *bool `json:"computeOptimized,omitempty" tf:"compute_optimized,omitempty"`

	// Will only include compute optimized nodes when enabled and exclude compute optimized nodes when disabled. Empty value won't have effect on instances filter. Supported values: `enabled`, `disabled` or empty string.
	ComputeOptimizedState *string `json:"computeOptimizedState,omitempty" tf:"compute_optimized_state,omitempty"`

	CustomPriority []CustomPriorityObservation `json:"customPriority,omitempty" tf:"custom_priority,omitempty"`

	// Dedicated node affinity - creates preference for instances to be created on sole tenancy or dedicated nodes. This
	// feature is only available for GCP clusters and sole tenancy nodes with local
	// SSDs or GPUs are not supported. If the sole tenancy or dedicated nodes don't have capacity for selected instance
	// type, the Autoscaler will fall back to multi-tenant instance types available for this Node Template.
	// Other instance constraints are applied when the Autoscaler picks available instance types that can be created on
	// the sole tenancy or dedicated node (example: setting min CPU to 16).
	DedicatedNodeAffinity []DedicatedNodeAffinityObservation `json:"dedicatedNodeAffinity,omitempty" tf:"dedicated_node_affinity,omitempty"`

	// Enable/disable spot diversity policy. When enabled, autoscaler will try to balance between diverse and cost optimal instance types.
	EnableSpotDiversity *bool `json:"enableSpotDiversity,omitempty" tf:"enable_spot_diversity,omitempty"`

	// Fallback restore rate in seconds: defines how much time should pass before spot fallback should be attempted to be restored to real spot.
	FallbackRestoreRateSeconds *float64 `json:"fallbackRestoreRateSeconds,omitempty" tf:"fallback_restore_rate_seconds,omitempty"`

	Gpu []GpuObservation `json:"gpu,omitempty" tf:"gpu,omitempty"`

	InstanceFamilies []InstanceFamiliesObservation `json:"instanceFamilies,omitempty" tf:"instance_families,omitempty"`

	// GPU instance constraint - will only pick nodes with GPU if true
	IsGpuOnly *bool `json:"isGpuOnly,omitempty" tf:"is_gpu_only,omitempty"`

	// Max CPU cores per node.
	MaxCPU *float64 `json:"maxCpu,omitempty" tf:"max_cpu,omitempty"`

	// Max Memory (Mib) per node.
	MaxMemory *float64 `json:"maxMemory,omitempty" tf:"max_memory,omitempty"`

	// Min CPU cores per node.
	MinCPU *float64 `json:"minCpu,omitempty" tf:"min_cpu,omitempty"`

	// Min Memory (Mib) per node.
	MinMemory *float64 `json:"minMemory,omitempty" tf:"min_memory,omitempty"`

	// Should include on-demand instances in the considered pool.
	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand,omitempty"`

	// List of acceptable instance Operating Systems, the default is linux. Allowed values: linux, windows.
	Os []*string `json:"os,omitempty" tf:"os,omitempty"`

	// Should include spot instances in the considered pool.
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`

	// Allowed node configuration price increase when diversifying instance types. E.g. if the value is 10%, then the overall price of diversified instance types can be 10% higher than the price of the optimal configuration.
	SpotDiversityPriceIncreaseLimitPercent *float64 `json:"spotDiversityPriceIncreaseLimitPercent,omitempty" tf:"spot_diversity_price_increase_limit_percent,omitempty"`

	// Enable/disable spot interruption predictions.
	SpotInterruptionPredictionsEnabled *bool `json:"spotInterruptionPredictionsEnabled,omitempty" tf:"spot_interruption_predictions_enabled,omitempty"`

	// Spot interruption predictions type. Can be either "aws-rebalance-recommendations" or "interruption-predictions".
	SpotInterruptionPredictionsType *string `json:"spotInterruptionPredictionsType,omitempty" tf:"spot_interruption_predictions_type,omitempty"`

	// Storage optimized instance constraint (deprecated).
	StorageOptimized *bool `json:"storageOptimized,omitempty" tf:"storage_optimized,omitempty"`

	// Storage optimized instance constraint - will only pick storage optimized nodes if enabled and won't pick if disabled. Empty value will have no effect. Supported values: `enabled`, `disabled` or empty string.
	StorageOptimizedState *string `json:"storageOptimizedState,omitempty" tf:"storage_optimized_state,omitempty"`

	// Spot instance fallback constraint - when true, on-demand instances will be created, when spots are unavailable.
	UseSpotFallbacks *bool `json:"useSpotFallbacks,omitempty" tf:"use_spot_fallbacks,omitempty"`
}

type ConstraintsParameters struct {

	// List of acceptable instance CPU architectures, the default is amd64. Allowed values: amd64, arm64.
	// +kubebuilder:validation:Optional
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// The list of AZ names to consider for the node template, if empty or not set all AZs are considered.
	// +kubebuilder:validation:Optional
	Azs []*string `json:"azs,omitempty" tf:"azs,omitempty"`

	// Will only include burstable instances when enabled otherwise they will be excluded. Supported values: `enabled`, `disabled` or “.
	// +kubebuilder:validation:Optional
	BurstableInstances *string `json:"burstableInstances,omitempty" tf:"burstable_instances,omitempty"`

	// Compute optimized instance constraint (deprecated).
	// +kubebuilder:validation:Optional
	ComputeOptimized *bool `json:"computeOptimized,omitempty" tf:"compute_optimized,omitempty"`

	// Will only include compute optimized nodes when enabled and exclude compute optimized nodes when disabled. Empty value won't have effect on instances filter. Supported values: `enabled`, `disabled` or empty string.
	// +kubebuilder:validation:Optional
	ComputeOptimizedState *string `json:"computeOptimizedState,omitempty" tf:"compute_optimized_state,omitempty"`

	// +kubebuilder:validation:Optional
	CustomPriority []CustomPriorityParameters `json:"customPriority,omitempty" tf:"custom_priority,omitempty"`

	// Dedicated node affinity - creates preference for instances to be created on sole tenancy or dedicated nodes. This
	// feature is only available for GCP clusters and sole tenancy nodes with local
	// SSDs or GPUs are not supported. If the sole tenancy or dedicated nodes don't have capacity for selected instance
	// type, the Autoscaler will fall back to multi-tenant instance types available for this Node Template.
	// Other instance constraints are applied when the Autoscaler picks available instance types that can be created on
	// the sole tenancy or dedicated node (example: setting min CPU to 16).
	// +kubebuilder:validation:Optional
	DedicatedNodeAffinity []DedicatedNodeAffinityParameters `json:"dedicatedNodeAffinity,omitempty" tf:"dedicated_node_affinity,omitempty"`

	// Enable/disable spot diversity policy. When enabled, autoscaler will try to balance between diverse and cost optimal instance types.
	// +kubebuilder:validation:Optional
	EnableSpotDiversity *bool `json:"enableSpotDiversity,omitempty" tf:"enable_spot_diversity,omitempty"`

	// Fallback restore rate in seconds: defines how much time should pass before spot fallback should be attempted to be restored to real spot.
	// +kubebuilder:validation:Optional
	FallbackRestoreRateSeconds *float64 `json:"fallbackRestoreRateSeconds,omitempty" tf:"fallback_restore_rate_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	Gpu []GpuParameters `json:"gpu,omitempty" tf:"gpu,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceFamilies []InstanceFamiliesParameters `json:"instanceFamilies,omitempty" tf:"instance_families,omitempty"`

	// GPU instance constraint - will only pick nodes with GPU if true
	// +kubebuilder:validation:Optional
	IsGpuOnly *bool `json:"isGpuOnly,omitempty" tf:"is_gpu_only,omitempty"`

	// Max CPU cores per node.
	// +kubebuilder:validation:Optional
	MaxCPU *float64 `json:"maxCpu,omitempty" tf:"max_cpu,omitempty"`

	// Max Memory (Mib) per node.
	// +kubebuilder:validation:Optional
	MaxMemory *float64 `json:"maxMemory,omitempty" tf:"max_memory,omitempty"`

	// Min CPU cores per node.
	// +kubebuilder:validation:Optional
	MinCPU *float64 `json:"minCpu,omitempty" tf:"min_cpu,omitempty"`

	// Min Memory (Mib) per node.
	// +kubebuilder:validation:Optional
	MinMemory *float64 `json:"minMemory,omitempty" tf:"min_memory,omitempty"`

	// Should include on-demand instances in the considered pool.
	// +kubebuilder:validation:Optional
	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand,omitempty"`

	// List of acceptable instance Operating Systems, the default is linux. Allowed values: linux, windows.
	// +kubebuilder:validation:Optional
	Os []*string `json:"os,omitempty" tf:"os,omitempty"`

	// Should include spot instances in the considered pool.
	// +kubebuilder:validation:Optional
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`

	// Allowed node configuration price increase when diversifying instance types. E.g. if the value is 10%, then the overall price of diversified instance types can be 10% higher than the price of the optimal configuration.
	// +kubebuilder:validation:Optional
	SpotDiversityPriceIncreaseLimitPercent *float64 `json:"spotDiversityPriceIncreaseLimitPercent,omitempty" tf:"spot_diversity_price_increase_limit_percent,omitempty"`

	// Enable/disable spot interruption predictions.
	// +kubebuilder:validation:Optional
	SpotInterruptionPredictionsEnabled *bool `json:"spotInterruptionPredictionsEnabled,omitempty" tf:"spot_interruption_predictions_enabled,omitempty"`

	// Spot interruption predictions type. Can be either "aws-rebalance-recommendations" or "interruption-predictions".
	// +kubebuilder:validation:Optional
	SpotInterruptionPredictionsType *string `json:"spotInterruptionPredictionsType,omitempty" tf:"spot_interruption_predictions_type,omitempty"`

	// Storage optimized instance constraint (deprecated).
	// +kubebuilder:validation:Optional
	StorageOptimized *bool `json:"storageOptimized,omitempty" tf:"storage_optimized,omitempty"`

	// Storage optimized instance constraint - will only pick storage optimized nodes if enabled and won't pick if disabled. Empty value will have no effect. Supported values: `enabled`, `disabled` or empty string.
	// +kubebuilder:validation:Optional
	StorageOptimizedState *string `json:"storageOptimizedState,omitempty" tf:"storage_optimized_state,omitempty"`

	// Spot instance fallback constraint - when true, on-demand instances will be created, when spots are unavailable.
	// +kubebuilder:validation:Optional
	UseSpotFallbacks *bool `json:"useSpotFallbacks,omitempty" tf:"use_spot_fallbacks,omitempty"`
}

type CustomPriorityInitParameters struct {

	// Instance families to prioritize in this tier.
	InstanceFamilies []*string `json:"instanceFamilies,omitempty" tf:"instance_families,omitempty"`

	// If true, this tier will apply to on-demand instances.
	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand,omitempty"`

	// If true, this tier will apply to spot instances.
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`
}

type CustomPriorityObservation struct {

	// Instance families to prioritize in this tier.
	InstanceFamilies []*string `json:"instanceFamilies,omitempty" tf:"instance_families,omitempty"`

	// If true, this tier will apply to on-demand instances.
	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand,omitempty"`

	// If true, this tier will apply to spot instances.
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`
}

type CustomPriorityParameters struct {

	// Instance families to prioritize in this tier.
	// +kubebuilder:validation:Optional
	InstanceFamilies []*string `json:"instanceFamilies,omitempty" tf:"instance_families,omitempty"`

	// If true, this tier will apply to on-demand instances.
	// +kubebuilder:validation:Optional
	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand,omitempty"`

	// If true, this tier will apply to spot instances.
	// +kubebuilder:validation:Optional
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`
}

type CustomTaintsInitParameters struct {

	// Effect of a taint to be added to nodes created from this template, the default is NoSchedule. Allowed values: NoSchedule, NoExecute.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// Key of a taint to be added to nodes created from this template.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of a taint to be added to nodes created from this template.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomTaintsObservation struct {

	// Effect of a taint to be added to nodes created from this template, the default is NoSchedule. Allowed values: NoSchedule, NoExecute.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// Key of a taint to be added to nodes created from this template.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of a taint to be added to nodes created from this template.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomTaintsParameters struct {

	// Effect of a taint to be added to nodes created from this template, the default is NoSchedule. Allowed values: NoSchedule, NoExecute.
	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// Key of a taint to be added to nodes created from this template.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Value of a taint to be added to nodes created from this template.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DedicatedNodeAffinityInitParameters struct {
	Affinity []AffinityInitParameters `json:"affinity,omitempty" tf:"affinity,omitempty"`

	// Availability zone name.
	AzName *string `json:"azName,omitempty" tf:"az_name,omitempty"`

	// Instance/node types in this node group.
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// Name of node group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DedicatedNodeAffinityObservation struct {
	Affinity []AffinityObservation `json:"affinity,omitempty" tf:"affinity,omitempty"`

	// Availability zone name.
	AzName *string `json:"azName,omitempty" tf:"az_name,omitempty"`

	// Instance/node types in this node group.
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// Name of node group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DedicatedNodeAffinityParameters struct {

	// +kubebuilder:validation:Optional
	Affinity []AffinityParameters `json:"affinity,omitempty" tf:"affinity,omitempty"`

	// Availability zone name.
	// +kubebuilder:validation:Optional
	AzName *string `json:"azName" tf:"az_name,omitempty"`

	// Instance/node types in this node group.
	// +kubebuilder:validation:Optional
	InstanceTypes []*string `json:"instanceTypes" tf:"instance_types,omitempty"`

	// Name of node group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type GpuInitParameters struct {

	// Names of the GPUs to exclude.
	ExcludeNames []*string `json:"excludeNames,omitempty" tf:"exclude_names,omitempty"`

	// Instance families to include when filtering (excludes all other families).
	IncludeNames []*string `json:"includeNames,omitempty" tf:"include_names,omitempty"`

	// Manufacturers of the gpus to select - NVIDIA, AMD.
	Manufacturers []*string `json:"manufacturers,omitempty" tf:"manufacturers,omitempty"`

	// Max GPU count for the instance type to have.
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// Min GPU count for the instance type to have.
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`
}

type GpuObservation struct {

	// Names of the GPUs to exclude.
	ExcludeNames []*string `json:"excludeNames,omitempty" tf:"exclude_names,omitempty"`

	// Instance families to include when filtering (excludes all other families).
	IncludeNames []*string `json:"includeNames,omitempty" tf:"include_names,omitempty"`

	// Manufacturers of the gpus to select - NVIDIA, AMD.
	Manufacturers []*string `json:"manufacturers,omitempty" tf:"manufacturers,omitempty"`

	// Max GPU count for the instance type to have.
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// Min GPU count for the instance type to have.
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`
}

type GpuParameters struct {

	// Names of the GPUs to exclude.
	// +kubebuilder:validation:Optional
	ExcludeNames []*string `json:"excludeNames,omitempty" tf:"exclude_names,omitempty"`

	// Instance families to include when filtering (excludes all other families).
	// +kubebuilder:validation:Optional
	IncludeNames []*string `json:"includeNames,omitempty" tf:"include_names,omitempty"`

	// Manufacturers of the gpus to select - NVIDIA, AMD.
	// +kubebuilder:validation:Optional
	Manufacturers []*string `json:"manufacturers,omitempty" tf:"manufacturers,omitempty"`

	// Max GPU count for the instance type to have.
	// +kubebuilder:validation:Optional
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// Min GPU count for the instance type to have.
	// +kubebuilder:validation:Optional
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`
}

type InstanceFamiliesInitParameters struct {

	// Instance families to include when filtering (excludes all other families).
	Exclude []*string `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// Instance families to exclude when filtering (includes all other families).
	Include []*string `json:"include,omitempty" tf:"include,omitempty"`
}

type InstanceFamiliesObservation struct {

	// Instance families to include when filtering (excludes all other families).
	Exclude []*string `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// Instance families to exclude when filtering (includes all other families).
	Include []*string `json:"include,omitempty" tf:"include,omitempty"`
}

type InstanceFamiliesParameters struct {

	// Instance families to include when filtering (excludes all other families).
	// +kubebuilder:validation:Optional
	Exclude []*string `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// Instance families to exclude when filtering (includes all other families).
	// +kubebuilder:validation:Optional
	Include []*string `json:"include,omitempty" tf:"include,omitempty"`
}

type NodeTemplateInitParameters struct {

	// CAST AI cluster id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// CAST AI node configuration id to be used for node template.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.NodeConfiguration
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	// Reference to a NodeConfiguration in castai to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDRef *v1.Reference `json:"configurationIdRef,omitempty" tf:"-"`

	// Selector for a NodeConfiguration in castai to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDSelector *v1.Selector `json:"configurationIdSelector,omitempty" tf:"-"`

	Constraints []ConstraintsInitParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Marks whether custom instances should be used when deciding which parts of inventory are available. Custom instances are only supported in GCP.
	CustomInstancesEnabled *bool `json:"customInstancesEnabled,omitempty" tf:"custom_instances_enabled,omitempty"`

	// Marks whether custom instances with extended memory should be used when deciding which parts of inventory are available. Custom instances are only supported in GCP.
	CustomInstancesWithExtendedMemoryEnabled *bool `json:"customInstancesWithExtendedMemoryEnabled,omitempty" tf:"custom_instances_with_extended_memory_enabled,omitempty"`

	// Custom labels to be added to nodes created from this template.
	// +mapType=granular
	CustomLabels map[string]*string `json:"customLabels,omitempty" tf:"custom_labels,omitempty"`

	// Custom taints to be added to the nodes created from this template. `shouldTaint` has to be `true` in order to create/update the node template with custom taints. If `shouldTaint` is `true`, but no custom taints are provided, the nodes will be tainted with the default node template taint.
	CustomTaints []CustomTaintsInitParameters `json:"customTaints,omitempty" tf:"custom_taints,omitempty"`

	// Flag whether the node template is default.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Flag whether the node template is enabled and considered for autoscaling.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// Name of the node template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Minimum nodes that will be kept when rebalancing nodes using this node template.
	RebalancingConfigMinNodes *float64 `json:"rebalancingConfigMinNodes,omitempty" tf:"rebalancing_config_min_nodes,omitempty"`

	// Marks whether the templated nodes will have a taint.
	ShouldTaint *bool `json:"shouldTaint,omitempty" tf:"should_taint,omitempty"`
}

type NodeTemplateObservation struct {

	// CAST AI cluster id.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// CAST AI node configuration id to be used for node template.
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	Constraints []ConstraintsObservation `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Marks whether custom instances should be used when deciding which parts of inventory are available. Custom instances are only supported in GCP.
	CustomInstancesEnabled *bool `json:"customInstancesEnabled,omitempty" tf:"custom_instances_enabled,omitempty"`

	// Marks whether custom instances with extended memory should be used when deciding which parts of inventory are available. Custom instances are only supported in GCP.
	CustomInstancesWithExtendedMemoryEnabled *bool `json:"customInstancesWithExtendedMemoryEnabled,omitempty" tf:"custom_instances_with_extended_memory_enabled,omitempty"`

	// Custom labels to be added to nodes created from this template.
	// +mapType=granular
	CustomLabels map[string]*string `json:"customLabels,omitempty" tf:"custom_labels,omitempty"`

	// Custom taints to be added to the nodes created from this template. `shouldTaint` has to be `true` in order to create/update the node template with custom taints. If `shouldTaint` is `true`, but no custom taints are provided, the nodes will be tainted with the default node template taint.
	CustomTaints []CustomTaintsObservation `json:"customTaints,omitempty" tf:"custom_taints,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Flag whether the node template is default.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Flag whether the node template is enabled and considered for autoscaling.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// Name of the node template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Minimum nodes that will be kept when rebalancing nodes using this node template.
	RebalancingConfigMinNodes *float64 `json:"rebalancingConfigMinNodes,omitempty" tf:"rebalancing_config_min_nodes,omitempty"`

	// Marks whether the templated nodes will have a taint.
	ShouldTaint *bool `json:"shouldTaint,omitempty" tf:"should_taint,omitempty"`
}

type NodeTemplateParameters struct {

	// CAST AI cluster id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// CAST AI node configuration id to be used for node template.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.NodeConfiguration
	// +kubebuilder:validation:Optional
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	// Reference to a NodeConfiguration in castai to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDRef *v1.Reference `json:"configurationIdRef,omitempty" tf:"-"`

	// Selector for a NodeConfiguration in castai to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDSelector *v1.Selector `json:"configurationIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Constraints []ConstraintsParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Marks whether custom instances should be used when deciding which parts of inventory are available. Custom instances are only supported in GCP.
	// +kubebuilder:validation:Optional
	CustomInstancesEnabled *bool `json:"customInstancesEnabled,omitempty" tf:"custom_instances_enabled,omitempty"`

	// Marks whether custom instances with extended memory should be used when deciding which parts of inventory are available. Custom instances are only supported in GCP.
	// +kubebuilder:validation:Optional
	CustomInstancesWithExtendedMemoryEnabled *bool `json:"customInstancesWithExtendedMemoryEnabled,omitempty" tf:"custom_instances_with_extended_memory_enabled,omitempty"`

	// Custom labels to be added to nodes created from this template.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomLabels map[string]*string `json:"customLabels,omitempty" tf:"custom_labels,omitempty"`

	// Custom taints to be added to the nodes created from this template. `shouldTaint` has to be `true` in order to create/update the node template with custom taints. If `shouldTaint` is `true`, but no custom taints are provided, the nodes will be tainted with the default node template taint.
	// +kubebuilder:validation:Optional
	CustomTaints []CustomTaintsParameters `json:"customTaints,omitempty" tf:"custom_taints,omitempty"`

	// Flag whether the node template is default.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Flag whether the node template is enabled and considered for autoscaling.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// Name of the node template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Minimum nodes that will be kept when rebalancing nodes using this node template.
	// +kubebuilder:validation:Optional
	RebalancingConfigMinNodes *float64 `json:"rebalancingConfigMinNodes,omitempty" tf:"rebalancing_config_min_nodes,omitempty"`

	// Marks whether the templated nodes will have a taint.
	// +kubebuilder:validation:Optional
	ShouldTaint *bool `json:"shouldTaint,omitempty" tf:"should_taint,omitempty"`
}

// NodeTemplateSpec defines the desired state of NodeTemplate
type NodeTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeTemplateParameters `json:"forProvider"`
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
	InitProvider NodeTemplateInitParameters `json:"initProvider,omitempty"`
}

// NodeTemplateStatus defines the observed state of NodeTemplate.
type NodeTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NodeTemplate is the Schema for the NodeTemplates API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type NodeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   NodeTemplateSpec   `json:"spec"`
	Status NodeTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeTemplateList contains a list of NodeTemplates
type NodeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeTemplate `json:"items"`
}

// Repository type metadata.
var (
	NodeTemplate_Kind             = "NodeTemplate"
	NodeTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeTemplate_Kind}.String()
	NodeTemplate_KindAPIVersion   = NodeTemplate_Kind + "." + CRDGroupVersion.String()
	NodeTemplate_GroupVersionKind = CRDGroupVersion.WithKind(NodeTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeTemplate{}, &NodeTemplateList{})
}
