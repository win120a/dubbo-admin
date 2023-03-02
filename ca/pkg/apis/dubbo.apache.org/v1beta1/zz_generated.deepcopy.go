//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendConfig) DeepCopyInto(out *ExtendConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendConfig.
func (in *ExtendConfig) DeepCopy() *ExtendConfig {
	if in == nil {
		return nil
	}
	out := new(ExtendConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerAuthentication) DeepCopyInto(out *PeerAuthentication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthentication.
func (in *PeerAuthentication) DeepCopy() *PeerAuthentication {
	if in == nil {
		return nil
	}
	out := new(PeerAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeerAuthentication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerAuthenticationList) DeepCopyInto(out *PeerAuthenticationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PeerAuthentication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthenticationList.
func (in *PeerAuthenticationList) DeepCopy() *PeerAuthenticationList {
	if in == nil {
		return nil
	}
	out := new(PeerAuthenticationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeerAuthenticationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerAuthenticationSpec) DeepCopyInto(out *PeerAuthenticationSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthenticationSpec.
func (in *PeerAuthenticationSpec) DeepCopy() *PeerAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(PeerAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	in.From.DeepCopyInto(&out.From)
	in.To.DeepCopyInto(&out.To)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotNamespaces != nil {
		in, out := &in.NotNamespaces, &out.NotNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IpBlocks != nil {
		in, out := &in.IpBlocks, &out.IpBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotIpBlocks != nil {
		in, out := &in.NotIpBlocks, &out.NotIpBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Principals != nil {
		in, out := &in.Principals, &out.Principals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotPrincipals != nil {
		in, out := &in.NotPrincipals, &out.NotPrincipals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extends != nil {
		in, out := &in.Extends, &out.Extends
		*out = make([]ExtendConfig, len(*in))
		copy(*out, *in)
	}
	if in.NotExtends != nil {
		in, out := &in.NotExtends, &out.NotExtends
		*out = make([]ExtendConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	if in.IpBlocks != nil {
		in, out := &in.IpBlocks, &out.IpBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotIpBlocks != nil {
		in, out := &in.NotIpBlocks, &out.NotIpBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Principals != nil {
		in, out := &in.Principals, &out.Principals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotPrincipals != nil {
		in, out := &in.NotPrincipals, &out.NotPrincipals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extends != nil {
		in, out := &in.Extends, &out.Extends
		*out = make([]ExtendConfig, len(*in))
		copy(*out, *in)
	}
	if in.NotExtends != nil {
		in, out := &in.NotExtends, &out.NotExtends
		*out = make([]ExtendConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}
