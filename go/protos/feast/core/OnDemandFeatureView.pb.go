//
// Copyright 2020 The Feast Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: feast/core/OnDemandFeatureView.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OnDemandFeatureView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User-specified specifications of this feature view.
	Spec *OnDemandFeatureViewSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Meta *OnDemandFeatureViewMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *OnDemandFeatureView) Reset() {
	*x = OnDemandFeatureView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnDemandFeatureView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnDemandFeatureView) ProtoMessage() {}

func (x *OnDemandFeatureView) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnDemandFeatureView.ProtoReflect.Descriptor instead.
func (*OnDemandFeatureView) Descriptor() ([]byte, []int) {
	return file_feast_core_OnDemandFeatureView_proto_rawDescGZIP(), []int{0}
}

func (x *OnDemandFeatureView) GetSpec() *OnDemandFeatureViewSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *OnDemandFeatureView) GetMeta() *OnDemandFeatureViewMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type OnDemandFeatureViewSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the feature view. Must be unique. Not updated.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of Feast project that this feature view belongs to.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// List of features specifications for each feature defined with this feature view.
	Features []*FeatureSpecV2 `protobuf:"bytes,3,rep,name=features,proto3" json:"features,omitempty"`
	// Map of inputs for this feature view.
	Inputs              map[string]*OnDemandInput `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserDefinedFunction *UserDefinedFunction      `protobuf:"bytes,5,opt,name=user_defined_function,json=userDefinedFunction,proto3" json:"user_defined_function,omitempty"`
}

func (x *OnDemandFeatureViewSpec) Reset() {
	*x = OnDemandFeatureViewSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnDemandFeatureViewSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnDemandFeatureViewSpec) ProtoMessage() {}

func (x *OnDemandFeatureViewSpec) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnDemandFeatureViewSpec.ProtoReflect.Descriptor instead.
func (*OnDemandFeatureViewSpec) Descriptor() ([]byte, []int) {
	return file_feast_core_OnDemandFeatureView_proto_rawDescGZIP(), []int{1}
}

func (x *OnDemandFeatureViewSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OnDemandFeatureViewSpec) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *OnDemandFeatureViewSpec) GetFeatures() []*FeatureSpecV2 {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *OnDemandFeatureViewSpec) GetInputs() map[string]*OnDemandInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *OnDemandFeatureViewSpec) GetUserDefinedFunction() *UserDefinedFunction {
	if x != nil {
		return x.UserDefinedFunction
	}
	return nil
}

type OnDemandFeatureViewMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time where this Feature View is created
	CreatedTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_timestamp,json=createdTimestamp,proto3" json:"created_timestamp,omitempty"`
	// Time where this Feature View is last updated
	LastUpdatedTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_updated_timestamp,json=lastUpdatedTimestamp,proto3" json:"last_updated_timestamp,omitempty"`
}

func (x *OnDemandFeatureViewMeta) Reset() {
	*x = OnDemandFeatureViewMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnDemandFeatureViewMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnDemandFeatureViewMeta) ProtoMessage() {}

func (x *OnDemandFeatureViewMeta) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnDemandFeatureViewMeta.ProtoReflect.Descriptor instead.
func (*OnDemandFeatureViewMeta) Descriptor() ([]byte, []int) {
	return file_feast_core_OnDemandFeatureView_proto_rawDescGZIP(), []int{2}
}

func (x *OnDemandFeatureViewMeta) GetCreatedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTimestamp
	}
	return nil
}

func (x *OnDemandFeatureViewMeta) GetLastUpdatedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedTimestamp
	}
	return nil
}

type OnDemandInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Input:
	//	*OnDemandInput_FeatureView
	//	*OnDemandInput_FeatureViewProjection
	//	*OnDemandInput_RequestDataSource
	Input isOnDemandInput_Input `protobuf_oneof:"input"`
}

func (x *OnDemandInput) Reset() {
	*x = OnDemandInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnDemandInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnDemandInput) ProtoMessage() {}

func (x *OnDemandInput) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnDemandInput.ProtoReflect.Descriptor instead.
func (*OnDemandInput) Descriptor() ([]byte, []int) {
	return file_feast_core_OnDemandFeatureView_proto_rawDescGZIP(), []int{3}
}

func (m *OnDemandInput) GetInput() isOnDemandInput_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *OnDemandInput) GetFeatureView() *FeatureView {
	if x, ok := x.GetInput().(*OnDemandInput_FeatureView); ok {
		return x.FeatureView
	}
	return nil
}

func (x *OnDemandInput) GetFeatureViewProjection() *FeatureViewProjection {
	if x, ok := x.GetInput().(*OnDemandInput_FeatureViewProjection); ok {
		return x.FeatureViewProjection
	}
	return nil
}

func (x *OnDemandInput) GetRequestDataSource() *DataSource {
	if x, ok := x.GetInput().(*OnDemandInput_RequestDataSource); ok {
		return x.RequestDataSource
	}
	return nil
}

type isOnDemandInput_Input interface {
	isOnDemandInput_Input()
}

type OnDemandInput_FeatureView struct {
	FeatureView *FeatureView `protobuf:"bytes,1,opt,name=feature_view,json=featureView,proto3,oneof"`
}

type OnDemandInput_FeatureViewProjection struct {
	FeatureViewProjection *FeatureViewProjection `protobuf:"bytes,3,opt,name=feature_view_projection,json=featureViewProjection,proto3,oneof"`
}

type OnDemandInput_RequestDataSource struct {
	RequestDataSource *DataSource `protobuf:"bytes,2,opt,name=request_data_source,json=requestDataSource,proto3,oneof"`
}

func (*OnDemandInput_FeatureView) isOnDemandInput_Input() {}

func (*OnDemandInput_FeatureViewProjection) isOnDemandInput_Input() {}

func (*OnDemandInput_RequestDataSource) isOnDemandInput_Input() {}

// Serialized representation of python function.
type UserDefinedFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The function name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The python-syntax function body (serialized by dill)
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UserDefinedFunction) Reset() {
	*x = UserDefinedFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDefinedFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDefinedFunction) ProtoMessage() {}

func (x *UserDefinedFunction) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_OnDemandFeatureView_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDefinedFunction.ProtoReflect.Descriptor instead.
func (*UserDefinedFunction) Descriptor() ([]byte, []int) {
	return file_feast_core_OnDemandFeatureView_proto_rawDescGZIP(), []int{4}
}

func (x *UserDefinedFunction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDefinedFunction) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_feast_core_OnDemandFeatureView_proto protoreflect.FileDescriptor

var file_feast_core_OnDemandFeatureView_proto_rawDesc = []byte{
	0x0a, 0x24, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x4f, 0x6e, 0x44,
	0x65, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x01, 0x0a, 0x13, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x6e, 0x44,
	0x65, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0xf2, 0x02, 0x0a, 0x17, 0x4f,
	0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69,
	0x65, 0x77, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x70, 0x65, 0x63, 0x56,
	0x32, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x06, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x66, 0x65,
	0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e,
	0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x12, 0x53, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x54, 0x0a, 0x0b, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x65, 0x61, 0x73,
	0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xb4, 0x01, 0x0a, 0x17, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x11, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x50, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xfd, 0x01, 0x0a, 0x0d, 0x4f, 0x6e, 0x44, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x5b, 0x0a, 0x17, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x15, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x3d, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x5d, 0x0a, 0x10, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x18, 0x4f, 0x6e, 0x44, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feast_core_OnDemandFeatureView_proto_rawDescOnce sync.Once
	file_feast_core_OnDemandFeatureView_proto_rawDescData = file_feast_core_OnDemandFeatureView_proto_rawDesc
)

func file_feast_core_OnDemandFeatureView_proto_rawDescGZIP() []byte {
	file_feast_core_OnDemandFeatureView_proto_rawDescOnce.Do(func() {
		file_feast_core_OnDemandFeatureView_proto_rawDescData = protoimpl.X.CompressGZIP(file_feast_core_OnDemandFeatureView_proto_rawDescData)
	})
	return file_feast_core_OnDemandFeatureView_proto_rawDescData
}

var file_feast_core_OnDemandFeatureView_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_feast_core_OnDemandFeatureView_proto_goTypes = []interface{}{
	(*OnDemandFeatureView)(nil),     // 0: feast.core.OnDemandFeatureView
	(*OnDemandFeatureViewSpec)(nil), // 1: feast.core.OnDemandFeatureViewSpec
	(*OnDemandFeatureViewMeta)(nil), // 2: feast.core.OnDemandFeatureViewMeta
	(*OnDemandInput)(nil),           // 3: feast.core.OnDemandInput
	(*UserDefinedFunction)(nil),     // 4: feast.core.UserDefinedFunction
	nil,                             // 5: feast.core.OnDemandFeatureViewSpec.InputsEntry
	(*FeatureSpecV2)(nil),           // 6: feast.core.FeatureSpecV2
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
	(*FeatureView)(nil),             // 8: feast.core.FeatureView
	(*FeatureViewProjection)(nil),   // 9: feast.core.FeatureViewProjection
	(*DataSource)(nil),              // 10: feast.core.DataSource
}
var file_feast_core_OnDemandFeatureView_proto_depIdxs = []int32{
	1,  // 0: feast.core.OnDemandFeatureView.spec:type_name -> feast.core.OnDemandFeatureViewSpec
	2,  // 1: feast.core.OnDemandFeatureView.meta:type_name -> feast.core.OnDemandFeatureViewMeta
	6,  // 2: feast.core.OnDemandFeatureViewSpec.features:type_name -> feast.core.FeatureSpecV2
	5,  // 3: feast.core.OnDemandFeatureViewSpec.inputs:type_name -> feast.core.OnDemandFeatureViewSpec.InputsEntry
	4,  // 4: feast.core.OnDemandFeatureViewSpec.user_defined_function:type_name -> feast.core.UserDefinedFunction
	7,  // 5: feast.core.OnDemandFeatureViewMeta.created_timestamp:type_name -> google.protobuf.Timestamp
	7,  // 6: feast.core.OnDemandFeatureViewMeta.last_updated_timestamp:type_name -> google.protobuf.Timestamp
	8,  // 7: feast.core.OnDemandInput.feature_view:type_name -> feast.core.FeatureView
	9,  // 8: feast.core.OnDemandInput.feature_view_projection:type_name -> feast.core.FeatureViewProjection
	10, // 9: feast.core.OnDemandInput.request_data_source:type_name -> feast.core.DataSource
	3,  // 10: feast.core.OnDemandFeatureViewSpec.InputsEntry.value:type_name -> feast.core.OnDemandInput
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_feast_core_OnDemandFeatureView_proto_init() }
func file_feast_core_OnDemandFeatureView_proto_init() {
	if File_feast_core_OnDemandFeatureView_proto != nil {
		return
	}
	file_feast_core_FeatureView_proto_init()
	file_feast_core_FeatureViewProjection_proto_init()
	file_feast_core_Feature_proto_init()
	file_feast_core_DataSource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_feast_core_OnDemandFeatureView_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnDemandFeatureView); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_OnDemandFeatureView_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnDemandFeatureViewSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_OnDemandFeatureView_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnDemandFeatureViewMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_OnDemandFeatureView_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnDemandInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_OnDemandFeatureView_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDefinedFunction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_feast_core_OnDemandFeatureView_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*OnDemandInput_FeatureView)(nil),
		(*OnDemandInput_FeatureViewProjection)(nil),
		(*OnDemandInput_RequestDataSource)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feast_core_OnDemandFeatureView_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feast_core_OnDemandFeatureView_proto_goTypes,
		DependencyIndexes: file_feast_core_OnDemandFeatureView_proto_depIdxs,
		MessageInfos:      file_feast_core_OnDemandFeatureView_proto_msgTypes,
	}.Build()
	File_feast_core_OnDemandFeatureView_proto = out.File
	file_feast_core_OnDemandFeatureView_proto_rawDesc = nil
	file_feast_core_OnDemandFeatureView_proto_goTypes = nil
	file_feast_core_OnDemandFeatureView_proto_depIdxs = nil
}
