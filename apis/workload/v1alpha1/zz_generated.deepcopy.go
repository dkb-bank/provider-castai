//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntiAffinityInitParameters) DeepCopyInto(out *AntiAffinityInitParameters) {
	*out = *in
	if in.ConsiderAntiAffinity != nil {
		in, out := &in.ConsiderAntiAffinity, &out.ConsiderAntiAffinity
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntiAffinityInitParameters.
func (in *AntiAffinityInitParameters) DeepCopy() *AntiAffinityInitParameters {
	if in == nil {
		return nil
	}
	out := new(AntiAffinityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntiAffinityObservation) DeepCopyInto(out *AntiAffinityObservation) {
	*out = *in
	if in.ConsiderAntiAffinity != nil {
		in, out := &in.ConsiderAntiAffinity, &out.ConsiderAntiAffinity
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntiAffinityObservation.
func (in *AntiAffinityObservation) DeepCopy() *AntiAffinityObservation {
	if in == nil {
		return nil
	}
	out := new(AntiAffinityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntiAffinityParameters) DeepCopyInto(out *AntiAffinityParameters) {
	*out = *in
	if in.ConsiderAntiAffinity != nil {
		in, out := &in.ConsiderAntiAffinity, &out.ConsiderAntiAffinity
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntiAffinityParameters.
func (in *AntiAffinityParameters) DeepCopy() *AntiAffinityParameters {
	if in == nil {
		return nil
	}
	out := new(AntiAffinityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyThresholdStrategyInitParameters) DeepCopyInto(out *ApplyThresholdStrategyInitParameters) {
	*out = *in
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyThresholdStrategyInitParameters.
func (in *ApplyThresholdStrategyInitParameters) DeepCopy() *ApplyThresholdStrategyInitParameters {
	if in == nil {
		return nil
	}
	out := new(ApplyThresholdStrategyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyThresholdStrategyObservation) DeepCopyInto(out *ApplyThresholdStrategyObservation) {
	*out = *in
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyThresholdStrategyObservation.
func (in *ApplyThresholdStrategyObservation) DeepCopy() *ApplyThresholdStrategyObservation {
	if in == nil {
		return nil
	}
	out := new(ApplyThresholdStrategyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyThresholdStrategyParameters) DeepCopyInto(out *ApplyThresholdStrategyParameters) {
	*out = *in
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyThresholdStrategyParameters.
func (in *ApplyThresholdStrategyParameters) DeepCopy() *ApplyThresholdStrategyParameters {
	if in == nil {
		return nil
	}
	out := new(ApplyThresholdStrategyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUInitParameters) DeepCopyInto(out *CPUInitParameters) {
	*out = *in
	if in.ApplyThreshold != nil {
		in, out := &in.ApplyThreshold, &out.ApplyThreshold
		*out = new(float64)
		**out = **in
	}
	if in.ApplyThresholdStrategy != nil {
		in, out := &in.ApplyThresholdStrategy, &out.ApplyThresholdStrategy
		*out = make([]ApplyThresholdStrategyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = make([]LimitInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LookBackPeriodSeconds != nil {
		in, out := &in.LookBackPeriodSeconds, &out.LookBackPeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUInitParameters.
func (in *CPUInitParameters) DeepCopy() *CPUInitParameters {
	if in == nil {
		return nil
	}
	out := new(CPUInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUObservation) DeepCopyInto(out *CPUObservation) {
	*out = *in
	if in.ApplyThreshold != nil {
		in, out := &in.ApplyThreshold, &out.ApplyThreshold
		*out = new(float64)
		**out = **in
	}
	if in.ApplyThresholdStrategy != nil {
		in, out := &in.ApplyThresholdStrategy, &out.ApplyThresholdStrategy
		*out = make([]ApplyThresholdStrategyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = make([]LimitObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LookBackPeriodSeconds != nil {
		in, out := &in.LookBackPeriodSeconds, &out.LookBackPeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUObservation.
func (in *CPUObservation) DeepCopy() *CPUObservation {
	if in == nil {
		return nil
	}
	out := new(CPUObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUParameters) DeepCopyInto(out *CPUParameters) {
	*out = *in
	if in.ApplyThreshold != nil {
		in, out := &in.ApplyThreshold, &out.ApplyThreshold
		*out = new(float64)
		**out = **in
	}
	if in.ApplyThresholdStrategy != nil {
		in, out := &in.ApplyThresholdStrategy, &out.ApplyThresholdStrategy
		*out = make([]ApplyThresholdStrategyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = make([]LimitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LookBackPeriodSeconds != nil {
		in, out := &in.LookBackPeriodSeconds, &out.LookBackPeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUParameters.
func (in *CPUParameters) DeepCopy() *CPUParameters {
	if in == nil {
		return nil
	}
	out := new(CPUParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownscalingInitParameters) DeepCopyInto(out *DownscalingInitParameters) {
	*out = *in
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownscalingInitParameters.
func (in *DownscalingInitParameters) DeepCopy() *DownscalingInitParameters {
	if in == nil {
		return nil
	}
	out := new(DownscalingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownscalingObservation) DeepCopyInto(out *DownscalingObservation) {
	*out = *in
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownscalingObservation.
func (in *DownscalingObservation) DeepCopy() *DownscalingObservation {
	if in == nil {
		return nil
	}
	out := new(DownscalingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownscalingParameters) DeepCopyInto(out *DownscalingParameters) {
	*out = *in
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownscalingParameters.
func (in *DownscalingParameters) DeepCopy() *DownscalingParameters {
	if in == nil {
		return nil
	}
	out := new(DownscalingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitInitParameters) DeepCopyInto(out *LimitInitParameters) {
	*out = *in
	if in.Multiplier != nil {
		in, out := &in.Multiplier, &out.Multiplier
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitInitParameters.
func (in *LimitInitParameters) DeepCopy() *LimitInitParameters {
	if in == nil {
		return nil
	}
	out := new(LimitInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitObservation) DeepCopyInto(out *LimitObservation) {
	*out = *in
	if in.Multiplier != nil {
		in, out := &in.Multiplier, &out.Multiplier
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitObservation.
func (in *LimitObservation) DeepCopy() *LimitObservation {
	if in == nil {
		return nil
	}
	out := new(LimitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitParameters) DeepCopyInto(out *LimitParameters) {
	*out = *in
	if in.Multiplier != nil {
		in, out := &in.Multiplier, &out.Multiplier
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitParameters.
func (in *LimitParameters) DeepCopy() *LimitParameters {
	if in == nil {
		return nil
	}
	out := new(LimitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryApplyThresholdStrategyInitParameters) DeepCopyInto(out *MemoryApplyThresholdStrategyInitParameters) {
	*out = *in
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryApplyThresholdStrategyInitParameters.
func (in *MemoryApplyThresholdStrategyInitParameters) DeepCopy() *MemoryApplyThresholdStrategyInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemoryApplyThresholdStrategyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryApplyThresholdStrategyObservation) DeepCopyInto(out *MemoryApplyThresholdStrategyObservation) {
	*out = *in
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryApplyThresholdStrategyObservation.
func (in *MemoryApplyThresholdStrategyObservation) DeepCopy() *MemoryApplyThresholdStrategyObservation {
	if in == nil {
		return nil
	}
	out := new(MemoryApplyThresholdStrategyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryApplyThresholdStrategyParameters) DeepCopyInto(out *MemoryApplyThresholdStrategyParameters) {
	*out = *in
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryApplyThresholdStrategyParameters.
func (in *MemoryApplyThresholdStrategyParameters) DeepCopy() *MemoryApplyThresholdStrategyParameters {
	if in == nil {
		return nil
	}
	out := new(MemoryApplyThresholdStrategyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryEventInitParameters) DeepCopyInto(out *MemoryEventInitParameters) {
	*out = *in
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryEventInitParameters.
func (in *MemoryEventInitParameters) DeepCopy() *MemoryEventInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemoryEventInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryEventObservation) DeepCopyInto(out *MemoryEventObservation) {
	*out = *in
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryEventObservation.
func (in *MemoryEventObservation) DeepCopy() *MemoryEventObservation {
	if in == nil {
		return nil
	}
	out := new(MemoryEventObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryEventParameters) DeepCopyInto(out *MemoryEventParameters) {
	*out = *in
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryEventParameters.
func (in *MemoryEventParameters) DeepCopy() *MemoryEventParameters {
	if in == nil {
		return nil
	}
	out := new(MemoryEventParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryInitParameters) DeepCopyInto(out *MemoryInitParameters) {
	*out = *in
	if in.ApplyThreshold != nil {
		in, out := &in.ApplyThreshold, &out.ApplyThreshold
		*out = new(float64)
		**out = **in
	}
	if in.ApplyThresholdStrategy != nil {
		in, out := &in.ApplyThresholdStrategy, &out.ApplyThresholdStrategy
		*out = make([]MemoryApplyThresholdStrategyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = make([]MemoryLimitInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LookBackPeriodSeconds != nil {
		in, out := &in.LookBackPeriodSeconds, &out.LookBackPeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryInitParameters.
func (in *MemoryInitParameters) DeepCopy() *MemoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryLimitInitParameters) DeepCopyInto(out *MemoryLimitInitParameters) {
	*out = *in
	if in.Multiplier != nil {
		in, out := &in.Multiplier, &out.Multiplier
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryLimitInitParameters.
func (in *MemoryLimitInitParameters) DeepCopy() *MemoryLimitInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemoryLimitInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryLimitObservation) DeepCopyInto(out *MemoryLimitObservation) {
	*out = *in
	if in.Multiplier != nil {
		in, out := &in.Multiplier, &out.Multiplier
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryLimitObservation.
func (in *MemoryLimitObservation) DeepCopy() *MemoryLimitObservation {
	if in == nil {
		return nil
	}
	out := new(MemoryLimitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryLimitParameters) DeepCopyInto(out *MemoryLimitParameters) {
	*out = *in
	if in.Multiplier != nil {
		in, out := &in.Multiplier, &out.Multiplier
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryLimitParameters.
func (in *MemoryLimitParameters) DeepCopy() *MemoryLimitParameters {
	if in == nil {
		return nil
	}
	out := new(MemoryLimitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryObservation) DeepCopyInto(out *MemoryObservation) {
	*out = *in
	if in.ApplyThreshold != nil {
		in, out := &in.ApplyThreshold, &out.ApplyThreshold
		*out = new(float64)
		**out = **in
	}
	if in.ApplyThresholdStrategy != nil {
		in, out := &in.ApplyThresholdStrategy, &out.ApplyThresholdStrategy
		*out = make([]MemoryApplyThresholdStrategyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = make([]MemoryLimitObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LookBackPeriodSeconds != nil {
		in, out := &in.LookBackPeriodSeconds, &out.LookBackPeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryObservation.
func (in *MemoryObservation) DeepCopy() *MemoryObservation {
	if in == nil {
		return nil
	}
	out := new(MemoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryParameters) DeepCopyInto(out *MemoryParameters) {
	*out = *in
	if in.ApplyThreshold != nil {
		in, out := &in.ApplyThreshold, &out.ApplyThreshold
		*out = new(float64)
		**out = **in
	}
	if in.ApplyThresholdStrategy != nil {
		in, out := &in.ApplyThresholdStrategy, &out.ApplyThresholdStrategy
		*out = make([]MemoryApplyThresholdStrategyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = make([]MemoryLimitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LookBackPeriodSeconds != nil {
		in, out := &in.LookBackPeriodSeconds, &out.LookBackPeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryParameters.
func (in *MemoryParameters) DeepCopy() *MemoryParameters {
	if in == nil {
		return nil
	}
	out := new(MemoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicy) DeepCopyInto(out *ScalingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicy.
func (in *ScalingPolicy) DeepCopy() *ScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyInitParameters) DeepCopyInto(out *ScalingPolicyInitParameters) {
	*out = *in
	if in.AntiAffinity != nil {
		in, out := &in.AntiAffinity, &out.AntiAffinity
		*out = make([]AntiAffinityInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = make([]CPUInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.Downscaling != nil {
		in, out := &in.Downscaling, &out.Downscaling
		*out = make([]DownscalingInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = make([]MemoryInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemoryEvent != nil {
		in, out := &in.MemoryEvent, &out.MemoryEvent
		*out = make([]MemoryEventInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = make([]StartupInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyInitParameters.
func (in *ScalingPolicyInitParameters) DeepCopy() *ScalingPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyList) DeepCopyInto(out *ScalingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyList.
func (in *ScalingPolicyList) DeepCopy() *ScalingPolicyList {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyObservation) DeepCopyInto(out *ScalingPolicyObservation) {
	*out = *in
	if in.AntiAffinity != nil {
		in, out := &in.AntiAffinity, &out.AntiAffinity
		*out = make([]AntiAffinityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = make([]CPUObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.Downscaling != nil {
		in, out := &in.Downscaling, &out.Downscaling
		*out = make([]DownscalingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = make([]MemoryObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemoryEvent != nil {
		in, out := &in.MemoryEvent, &out.MemoryEvent
		*out = make([]MemoryEventObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = make([]StartupObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyObservation.
func (in *ScalingPolicyObservation) DeepCopy() *ScalingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyParameters) DeepCopyInto(out *ScalingPolicyParameters) {
	*out = *in
	if in.AntiAffinity != nil {
		in, out := &in.AntiAffinity, &out.AntiAffinity
		*out = make([]AntiAffinityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ApplyType != nil {
		in, out := &in.ApplyType, &out.ApplyType
		*out = new(string)
		**out = **in
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = make([]CPUParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.Downscaling != nil {
		in, out := &in.Downscaling, &out.Downscaling
		*out = make([]DownscalingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagementOption != nil {
		in, out := &in.ManagementOption, &out.ManagementOption
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = make([]MemoryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemoryEvent != nil {
		in, out := &in.MemoryEvent, &out.MemoryEvent
		*out = make([]MemoryEventParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = make([]StartupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyParameters.
func (in *ScalingPolicyParameters) DeepCopy() *ScalingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpec) DeepCopyInto(out *ScalingPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpec.
func (in *ScalingPolicySpec) DeepCopy() *ScalingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyStatus) DeepCopyInto(out *ScalingPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyStatus.
func (in *ScalingPolicyStatus) DeepCopy() *ScalingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupInitParameters) DeepCopyInto(out *StartupInitParameters) {
	*out = *in
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupInitParameters.
func (in *StartupInitParameters) DeepCopy() *StartupInitParameters {
	if in == nil {
		return nil
	}
	out := new(StartupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupObservation) DeepCopyInto(out *StartupObservation) {
	*out = *in
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupObservation.
func (in *StartupObservation) DeepCopy() *StartupObservation {
	if in == nil {
		return nil
	}
	out := new(StartupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupParameters) DeepCopyInto(out *StartupParameters) {
	*out = *in
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupParameters.
func (in *StartupParameters) DeepCopy() *StartupParameters {
	if in == nil {
		return nil
	}
	out := new(StartupParameters)
	in.DeepCopyInto(out)
	return out
}
