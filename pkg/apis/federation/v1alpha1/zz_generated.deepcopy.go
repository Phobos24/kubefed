// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedCluster) DeepCopyInto(out *FederatedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedCluster.
func (in *FederatedCluster) DeepCopy() *FederatedCluster {
	if in == nil {
		return nil
	}
	out := new(FederatedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedClusterList) DeepCopyInto(out *FederatedClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedClusterList.
func (in *FederatedClusterList) DeepCopy() *FederatedClusterList {
	if in == nil {
		return nil
	}
	out := new(FederatedClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedClusterSpec) DeepCopyInto(out *FederatedClusterSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LocalObjectReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedClusterSpec.
func (in *FederatedClusterSpec) DeepCopy() *FederatedClusterSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedClusterStatus) DeepCopyInto(out *FederatedClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedClusterStatus.
func (in *FederatedClusterStatus) DeepCopy() *FederatedClusterStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedConfigMap) DeepCopyInto(out *FederatedConfigMap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedConfigMap.
func (in *FederatedConfigMap) DeepCopy() *FederatedConfigMap {
	if in == nil {
		return nil
	}
	out := new(FederatedConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedConfigMap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedConfigMapList) DeepCopyInto(out *FederatedConfigMapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedConfigMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedConfigMapList.
func (in *FederatedConfigMapList) DeepCopy() *FederatedConfigMapList {
	if in == nil {
		return nil
	}
	out := new(FederatedConfigMapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedConfigMapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedConfigMapPlacement) DeepCopyInto(out *FederatedConfigMapPlacement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedConfigMapPlacement.
func (in *FederatedConfigMapPlacement) DeepCopy() *FederatedConfigMapPlacement {
	if in == nil {
		return nil
	}
	out := new(FederatedConfigMapPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedConfigMapPlacement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedConfigMapPlacementList) DeepCopyInto(out *FederatedConfigMapPlacementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedConfigMapPlacement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedConfigMapPlacementList.
func (in *FederatedConfigMapPlacementList) DeepCopy() *FederatedConfigMapPlacementList {
	if in == nil {
		return nil
	}
	out := new(FederatedConfigMapPlacementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedConfigMapPlacementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedConfigMapPlacementSpec) DeepCopyInto(out *FederatedConfigMapPlacementSpec) {
	*out = *in
	if in.ClusterNames != nil {
		in, out := &in.ClusterNames, &out.ClusterNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedConfigMapPlacementSpec.
func (in *FederatedConfigMapPlacementSpec) DeepCopy() *FederatedConfigMapPlacementSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedConfigMapPlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedConfigMapPlacementStatus) DeepCopyInto(out *FederatedConfigMapPlacementStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedConfigMapPlacementStatus.
func (in *FederatedConfigMapPlacementStatus) DeepCopy() *FederatedConfigMapPlacementStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedConfigMapPlacementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedConfigMapSpec) DeepCopyInto(out *FederatedConfigMapSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedConfigMapSpec.
func (in *FederatedConfigMapSpec) DeepCopy() *FederatedConfigMapSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedConfigMapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedConfigMapStatus) DeepCopyInto(out *FederatedConfigMapStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedConfigMapStatus.
func (in *FederatedConfigMapStatus) DeepCopy() *FederatedConfigMapStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedConfigMapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSet) DeepCopyInto(out *FederatedReplicaSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSet.
func (in *FederatedReplicaSet) DeepCopy() *FederatedReplicaSet {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedReplicaSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSetClusterOverride) DeepCopyInto(out *FederatedReplicaSetClusterOverride) {
	*out = *in
	in.Override.DeepCopyInto(&out.Override)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSetClusterOverride.
func (in *FederatedReplicaSetClusterOverride) DeepCopy() *FederatedReplicaSetClusterOverride {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSetClusterOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSetList) DeepCopyInto(out *FederatedReplicaSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedReplicaSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSetList.
func (in *FederatedReplicaSetList) DeepCopy() *FederatedReplicaSetList {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedReplicaSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSetOverride) DeepCopyInto(out *FederatedReplicaSetOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSetOverride.
func (in *FederatedReplicaSetOverride) DeepCopy() *FederatedReplicaSetOverride {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSetOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedReplicaSetOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSetOverrideList) DeepCopyInto(out *FederatedReplicaSetOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedReplicaSetOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSetOverrideList.
func (in *FederatedReplicaSetOverrideList) DeepCopy() *FederatedReplicaSetOverrideList {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSetOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedReplicaSetOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSetOverrideSpec) DeepCopyInto(out *FederatedReplicaSetOverrideSpec) {
	*out = *in
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]FederatedReplicaSetClusterOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSetOverrideSpec.
func (in *FederatedReplicaSetOverrideSpec) DeepCopy() *FederatedReplicaSetOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSetOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSetOverrideStatus) DeepCopyInto(out *FederatedReplicaSetOverrideStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSetOverrideStatus.
func (in *FederatedReplicaSetOverrideStatus) DeepCopy() *FederatedReplicaSetOverrideStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSetOverrideStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSetSpec) DeepCopyInto(out *FederatedReplicaSetSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSetSpec.
func (in *FederatedReplicaSetSpec) DeepCopy() *FederatedReplicaSetSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedReplicaSetStatus) DeepCopyInto(out *FederatedReplicaSetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedReplicaSetStatus.
func (in *FederatedReplicaSetStatus) DeepCopy() *FederatedReplicaSetStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedReplicaSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecret) DeepCopyInto(out *FederatedSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecret.
func (in *FederatedSecret) DeepCopy() *FederatedSecret {
	if in == nil {
		return nil
	}
	out := new(FederatedSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretClusterOverride) DeepCopyInto(out *FederatedSecretClusterOverride) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]byte, len(val))
				copy((*out)[key], val)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretClusterOverride.
func (in *FederatedSecretClusterOverride) DeepCopy() *FederatedSecretClusterOverride {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretClusterOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretList) DeepCopyInto(out *FederatedSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretList.
func (in *FederatedSecretList) DeepCopy() *FederatedSecretList {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretOverride) DeepCopyInto(out *FederatedSecretOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretOverride.
func (in *FederatedSecretOverride) DeepCopy() *FederatedSecretOverride {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecretOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretOverrideList) DeepCopyInto(out *FederatedSecretOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedSecretOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretOverrideList.
func (in *FederatedSecretOverrideList) DeepCopy() *FederatedSecretOverrideList {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecretOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretOverrideSpec) DeepCopyInto(out *FederatedSecretOverrideSpec) {
	*out = *in
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]FederatedSecretClusterOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretOverrideSpec.
func (in *FederatedSecretOverrideSpec) DeepCopy() *FederatedSecretOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretOverrideStatus) DeepCopyInto(out *FederatedSecretOverrideStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretOverrideStatus.
func (in *FederatedSecretOverrideStatus) DeepCopy() *FederatedSecretOverrideStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretOverrideStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretPlacement) DeepCopyInto(out *FederatedSecretPlacement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretPlacement.
func (in *FederatedSecretPlacement) DeepCopy() *FederatedSecretPlacement {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecretPlacement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretPlacementList) DeepCopyInto(out *FederatedSecretPlacementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedSecretPlacement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretPlacementList.
func (in *FederatedSecretPlacementList) DeepCopy() *FederatedSecretPlacementList {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretPlacementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecretPlacementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretPlacementSpec) DeepCopyInto(out *FederatedSecretPlacementSpec) {
	*out = *in
	if in.ClusterNames != nil {
		in, out := &in.ClusterNames, &out.ClusterNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretPlacementSpec.
func (in *FederatedSecretPlacementSpec) DeepCopy() *FederatedSecretPlacementSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretPlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretPlacementStatus) DeepCopyInto(out *FederatedSecretPlacementStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretPlacementStatus.
func (in *FederatedSecretPlacementStatus) DeepCopy() *FederatedSecretPlacementStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretPlacementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretSpec) DeepCopyInto(out *FederatedSecretSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretSpec.
func (in *FederatedSecretSpec) DeepCopy() *FederatedSecretSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretStatus) DeepCopyInto(out *FederatedSecretStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretStatus.
func (in *FederatedSecretStatus) DeepCopy() *FederatedSecretStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretStatus)
	in.DeepCopyInto(out)
	return out
}
