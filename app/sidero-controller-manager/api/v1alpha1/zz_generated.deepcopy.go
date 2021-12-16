//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1alpha4"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMT) DeepCopyInto(out *AMT) {
	*out = *in
	if in.UserFrom != nil {
		in, out := &in.UserFrom, &out.UserFrom
		*out = new(CredentialSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PassFrom != nil {
		in, out := &in.PassFrom, &out.PassFrom
		*out = new(CredentialSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMT.
func (in *AMT) DeepCopy() *AMT {
	if in == nil {
		return nil
	}
	out := new(AMT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Asset) DeepCopyInto(out *Asset) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Asset.
func (in *Asset) DeepCopy() *Asset {
	if in == nil {
		return nil
	}
	out := new(Asset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetCondition) DeepCopyInto(out *AssetCondition) {
	*out = *in
	out.Asset = in.Asset
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetCondition.
func (in *AssetCondition) DeepCopy() *AssetCondition {
	if in == nil {
		return nil
	}
	out := new(AssetCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BMC) DeepCopyInto(out *BMC) {
	*out = *in
	if in.UserFrom != nil {
		in, out := &in.UserFrom, &out.UserFrom
		*out = new(CredentialSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PassFrom != nil {
		in, out := &in.PassFrom, &out.PassFrom
		*out = new(CredentialSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BMC.
func (in *BMC) DeepCopy() *BMC {
	if in == nil {
		return nil
	}
	out := new(BMC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUInformation) DeepCopyInto(out *CPUInformation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUInformation.
func (in *CPUInformation) DeepCopy() *CPUInformation {
	if in == nil {
		return nil
	}
	out := new(CPUInformation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigPatches) DeepCopyInto(out *ConfigPatches) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigPatches.
func (in *ConfigPatches) DeepCopy() *ConfigPatches {
	if in == nil {
		return nil
	}
	out := new(ConfigPatches)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSource) DeepCopyInto(out *CredentialSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(SecretKeyRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSource.
func (in *CredentialSource) DeepCopy() *CredentialSource {
	if in == nil {
		return nil
	}
	out := new(CredentialSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	in.Kernel.DeepCopyInto(&out.Kernel)
	out.Initrd = in.Initrd
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AssetCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Initrd) DeepCopyInto(out *Initrd) {
	*out = *in
	out.Asset = in.Asset
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Initrd.
func (in *Initrd) DeepCopy() *Initrd {
	if in == nil {
		return nil
	}
	out := new(Initrd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kernel) DeepCopyInto(out *Kernel) {
	*out = *in
	out.Asset = in.Asset
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kernel.
func (in *Kernel) DeepCopy() *Kernel {
	if in == nil {
		return nil
	}
	out := new(Kernel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementAPI) DeepCopyInto(out *ManagementAPI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementAPI.
func (in *ManagementAPI) DeepCopy() *ManagementAPI {
	if in == nil {
		return nil
	}
	out := new(ManagementAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Qualifiers) DeepCopyInto(out *Qualifiers) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = make([]CPUInformation, len(*in))
		copy(*out, *in)
	}
	if in.SystemInformation != nil {
		in, out := &in.SystemInformation, &out.SystemInformation
		*out = make([]SystemInformation, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelectors != nil {
		in, out := &in.LabelSelectors, &out.LabelSelectors
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Qualifiers.
func (in *Qualifiers) DeepCopy() *Qualifiers {
	if in == nil {
		return nil
	}
	out := new(Qualifiers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyRef) DeepCopyInto(out *SecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyRef.
func (in *SecretKeyRef) DeepCopy() *SecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerClass) DeepCopyInto(out *ServerClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerClass.
func (in *ServerClass) DeepCopy() *ServerClass {
	if in == nil {
		return nil
	}
	out := new(ServerClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerClassList) DeepCopyInto(out *ServerClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServerClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerClassList.
func (in *ServerClassList) DeepCopy() *ServerClassList {
	if in == nil {
		return nil
	}
	out := new(ServerClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerClassSpec) DeepCopyInto(out *ServerClassSpec) {
	*out = *in
	if in.EnvironmentRef != nil {
		in, out := &in.EnvironmentRef, &out.EnvironmentRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	in.Qualifiers.DeepCopyInto(&out.Qualifiers)
	in.Selector.DeepCopyInto(&out.Selector)
	if in.ConfigPatches != nil {
		in, out := &in.ConfigPatches, &out.ConfigPatches
		*out = make([]ConfigPatches, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerClassSpec.
func (in *ServerClassSpec) DeepCopy() *ServerClassSpec {
	if in == nil {
		return nil
	}
	out := new(ServerClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerClassStatus) DeepCopyInto(out *ServerClassStatus) {
	*out = *in
	if in.ServersAvailable != nil {
		in, out := &in.ServersAvailable, &out.ServersAvailable
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServersInUse != nil {
		in, out := &in.ServersInUse, &out.ServersInUse
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerClassStatus.
func (in *ServerClassStatus) DeepCopy() *ServerClassStatus {
	if in == nil {
		return nil
	}
	out := new(ServerClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpec) DeepCopyInto(out *ServerSpec) {
	*out = *in
	if in.EnvironmentRef != nil {
		in, out := &in.EnvironmentRef, &out.EnvironmentRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.SystemInformation != nil {
		in, out := &in.SystemInformation, &out.SystemInformation
		*out = new(SystemInformation)
		**out = **in
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPUInformation)
		**out = **in
	}
	if in.BMC != nil {
		in, out := &in.BMC, &out.BMC
		*out = new(BMC)
		(*in).DeepCopyInto(*out)
	}
	if in.AMT != nil {
		in, out := &in.AMT, &out.AMT
		*out = new(AMT)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagementAPI != nil {
		in, out := &in.ManagementAPI, &out.ManagementAPI
		*out = new(ManagementAPI)
		**out = **in
	}
	if in.ConfigPatches != nil {
		in, out := &in.ConfigPatches, &out.ConfigPatches
		*out = make([]ConfigPatches, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpec.
func (in *ServerSpec) DeepCopy() *ServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStatus) DeepCopyInto(out *ServerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha4.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStatus.
func (in *ServerStatus) DeepCopy() *ServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemInformation) DeepCopyInto(out *SystemInformation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemInformation.
func (in *SystemInformation) DeepCopy() *SystemInformation {
	if in == nil {
		return nil
	}
	out := new(SystemInformation)
	in.DeepCopyInto(out)
	return out
}
