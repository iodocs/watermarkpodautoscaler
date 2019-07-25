// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossVersionObjectReference) DeepCopyInto(out *CrossVersionObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossVersionObjectReference.
func (in *CrossVersionObjectReference) DeepCopy() *CrossVersionObjectReference {
	if in == nil {
		return nil
	}
	out := new(CrossVersionObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetricSource) DeepCopyInto(out *ExternalMetricSource) {
	*out = *in
	if in.MetricSelector != nil {
		in, out := &in.MetricSelector, &out.MetricSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.HighWatermark != nil {
		in, out := &in.HighWatermark, &out.HighWatermark
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.LowWatermark != nil {
		in, out := &in.LowWatermark, &out.LowWatermark
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricSource.
func (in *ExternalMetricSource) DeepCopy() *ExternalMetricSource {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalMetricSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkPodAutoscaler) DeepCopyInto(out *WatermarkPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkPodAutoscaler.
func (in *WatermarkPodAutoscaler) DeepCopy() *WatermarkPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(WatermarkPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatermarkPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkPodAutoscalerList) DeepCopyInto(out *WatermarkPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WatermarkPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkPodAutoscalerList.
func (in *WatermarkPodAutoscalerList) DeepCopy() *WatermarkPodAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(WatermarkPodAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatermarkPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkPodAutoscalerSpec) DeepCopyInto(out *WatermarkPodAutoscalerSpec) {
	*out = *in
	out.ScaleTargetRef = in.ScaleTargetRef
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkPodAutoscalerSpec.
func (in *WatermarkPodAutoscalerSpec) DeepCopy() *WatermarkPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(WatermarkPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkPodAutoscalerStatus) DeepCopyInto(out *WatermarkPodAutoscalerStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.LastScaleTime != nil {
		in, out := &in.LastScaleTime, &out.LastScaleTime
		*out = (*in).DeepCopy()
	}
	if in.CurrentMetrics != nil {
		in, out := &in.CurrentMetrics, &out.CurrentMetrics
		*out = make([]v2beta1.MetricStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v2beta1.HorizontalPodAutoscalerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkPodAutoscalerStatus.
func (in *WatermarkPodAutoscalerStatus) DeepCopy() *WatermarkPodAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(WatermarkPodAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}