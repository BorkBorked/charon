// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: app/protonil/testdata/v1/test.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type M1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	M2         *M2    `protobuf:"bytes,2,opt,name=m2,proto3" json:"m2,omitempty"`
	M2Optional *M2    `protobuf:"bytes,3,opt,name=m2_optional,json=m2Optional,proto3,oneof" json:"m2_optional,omitempty"`
}

func (x *M1) Reset() {
	*x = M1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M1) ProtoMessage() {}

func (x *M1) ProtoReflect() protoreflect.Message {
	mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M1.ProtoReflect.Descriptor instead.
func (*M1) Descriptor() ([]byte, []int) {
	return file_app_protonil_testdata_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *M1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *M1) GetM2() *M2 {
	if x != nil {
		return x.M2
	}
	return nil
}

func (x *M1) GetM2Optional() *M2 {
	if x != nil {
		return x.M2Optional
	}
	return nil
}

type M2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	M3         *M3    `protobuf:"bytes,2,opt,name=m3,proto3" json:"m3,omitempty"`
	M3Optional *M3    `protobuf:"bytes,3,opt,name=m3_optional,json=m3Optional,proto3,oneof" json:"m3_optional,omitempty"`
}

func (x *M2) Reset() {
	*x = M2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2) ProtoMessage() {}

func (x *M2) ProtoReflect() protoreflect.Message {
	mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2.ProtoReflect.Descriptor instead.
func (*M2) Descriptor() ([]byte, []int) {
	return file_app_protonil_testdata_v1_test_proto_rawDescGZIP(), []int{1}
}

func (x *M2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *M2) GetM3() *M3 {
	if x != nil {
		return x.M3
	}
	return nil
}

func (x *M2) GetM3Optional() *M3 {
	if x != nil {
		return x.M3Optional
	}
	return nil
}

type M3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Note it doesn't start at 1
}

func (x *M3) Reset() {
	*x = M3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M3) ProtoMessage() {}

func (x *M3) ProtoReflect() protoreflect.Message {
	mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M3.ProtoReflect.Descriptor instead.
func (*M3) Descriptor() ([]byte, []int) {
	return file_app_protonil_testdata_v1_test_proto_rawDescGZIP(), []int{2}
}

func (x *M3) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type M4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	M3Map  map[string]*M3 `protobuf:"bytes,2,rep,name=m3_map,json=m3Map,proto3" json:"m3_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	M3List []*M3          `protobuf:"bytes,3,rep,name=m3_list,json=m3List,proto3" json:"m3_list,omitempty"`
}

func (x *M4) Reset() {
	*x = M4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M4) ProtoMessage() {}

func (x *M4) ProtoReflect() protoreflect.Message {
	mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M4.ProtoReflect.Descriptor instead.
func (*M4) Descriptor() ([]byte, []int) {
	return file_app_protonil_testdata_v1_test_proto_rawDescGZIP(), []int{3}
}

func (x *M4) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *M4) GetM3Map() map[string]*M3 {
	if x != nil {
		return x.M3Map
	}
	return nil
}

func (x *M4) GetM3List() []*M3 {
	if x != nil {
		return x.M3List
	}
	return nil
}

// MaxIndex is used to test the max index
type MaxIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,65,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MaxIndex) Reset() {
	*x = MaxIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxIndex) ProtoMessage() {}

func (x *MaxIndex) ProtoReflect() protoreflect.Message {
	mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxIndex.ProtoReflect.Descriptor instead.
func (*MaxIndex) Descriptor() ([]byte, []int) {
	return file_app_protonil_testdata_v1_test_proto_rawDescGZIP(), []int{4}
}

func (x *MaxIndex) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Attack is a weird message that extends M1
// and adds a unknown field and a invalid numbered field.
// Both should be ignored.
type Attack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	M2         *M2    `protobuf:"bytes,2,opt,name=m2,proto3" json:"m2,omitempty"`
	M2Optional *M2    `protobuf:"bytes,3,opt,name=m2_optional,json=m2Optional,proto3,oneof" json:"m2_optional,omitempty"`
	M3Unknown  *M3    `protobuf:"bytes,4,opt,name=m3_unknown,json=m3Unknown,proto3" json:"m3_unknown,omitempty"`
	M3Attack   *M3    `protobuf:"bytes,99999999,opt,name=m3_attack,json=m3Attack,proto3" json:"m3_attack,omitempty"`
}

func (x *Attack) Reset() {
	*x = Attack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attack) ProtoMessage() {}

func (x *Attack) ProtoReflect() protoreflect.Message {
	mi := &file_app_protonil_testdata_v1_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attack.ProtoReflect.Descriptor instead.
func (*Attack) Descriptor() ([]byte, []int) {
	return file_app_protonil_testdata_v1_test_proto_rawDescGZIP(), []int{5}
}

func (x *Attack) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attack) GetM2() *M2 {
	if x != nil {
		return x.M2
	}
	return nil
}

func (x *Attack) GetM2Optional() *M2 {
	if x != nil {
		return x.M2Optional
	}
	return nil
}

func (x *Attack) GetM3Unknown() *M3 {
	if x != nil {
		return x.M3Unknown
	}
	return nil
}

func (x *Attack) GetM3Attack() *M3 {
	if x != nil {
		return x.M3Attack
	}
	return nil
}

var File_app_protonil_testdata_v1_test_proto protoreflect.FileDescriptor

var file_app_protonil_testdata_v1_test_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x22,
	0x9a, 0x01, 0x0a, 0x02, 0x4d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x02, 0x6d, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x32, 0x52, 0x02, 0x6d, 0x32, 0x12, 0x42, 0x0a, 0x0b, 0x6d, 0x32, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x32, 0x48, 0x00, 0x52, 0x0a, 0x6d,
	0x32, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x6d, 0x32, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x9a, 0x01, 0x0a,
	0x02, 0x4d, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x02, 0x6d, 0x33, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e,
	0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x33, 0x52, 0x02, 0x6d, 0x33, 0x12, 0x42, 0x0a, 0x0b, 0x6d, 0x33, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x33, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x33, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x33,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x18, 0x0a, 0x02, 0x4d, 0x33, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xed, 0x01, 0x0a, 0x02, 0x4d, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e,
	0x0a, 0x06, 0x6d, 0x33, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x34, 0x2e, 0x4d, 0x33, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x33, 0x4d, 0x61, 0x70, 0x12, 0x35,
	0x0a, 0x07, 0x6d, 0x33, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x33, 0x52, 0x06, 0x6d,
	0x33, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x56, 0x0a, 0x0a, 0x4d, 0x33, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x33, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08,
	0x04, 0x10, 0x05, 0x22, 0x1e, 0x0a, 0x08, 0x4d, 0x61, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x41, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x99, 0x02, 0x0a, 0x06, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x02, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x32, 0x52, 0x02, 0x6d, 0x32,
	0x12, 0x42, 0x0a, 0x0b, 0x6d, 0x32, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x32, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x32, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0a, 0x6d, 0x33, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x33, 0x52, 0x09, 0x6d, 0x33, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x12, 0x3c, 0x0a, 0x09, 0x6d, 0x33, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0xff,
	0xc1, 0xd7, 0x2f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x33, 0x52, 0x08, 0x6d, 0x33, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x32, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x62,
	0x6f, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x6c, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_app_protonil_testdata_v1_test_proto_rawDescOnce sync.Once
	file_app_protonil_testdata_v1_test_proto_rawDescData = file_app_protonil_testdata_v1_test_proto_rawDesc
)

func file_app_protonil_testdata_v1_test_proto_rawDescGZIP() []byte {
	file_app_protonil_testdata_v1_test_proto_rawDescOnce.Do(func() {
		file_app_protonil_testdata_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_protonil_testdata_v1_test_proto_rawDescData)
	})
	return file_app_protonil_testdata_v1_test_proto_rawDescData
}

var file_app_protonil_testdata_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_app_protonil_testdata_v1_test_proto_goTypes = []interface{}{
	(*M1)(nil),       // 0: app.protonil.testdata.v1.M1
	(*M2)(nil),       // 1: app.protonil.testdata.v1.M2
	(*M3)(nil),       // 2: app.protonil.testdata.v1.M3
	(*M4)(nil),       // 3: app.protonil.testdata.v1.M4
	(*MaxIndex)(nil), // 4: app.protonil.testdata.v1.MaxIndex
	(*Attack)(nil),   // 5: app.protonil.testdata.v1.Attack
	nil,              // 6: app.protonil.testdata.v1.M4.M3MapEntry
}
var file_app_protonil_testdata_v1_test_proto_depIdxs = []int32{
	1,  // 0: app.protonil.testdata.v1.M1.m2:type_name -> app.protonil.testdata.v1.M2
	1,  // 1: app.protonil.testdata.v1.M1.m2_optional:type_name -> app.protonil.testdata.v1.M2
	2,  // 2: app.protonil.testdata.v1.M2.m3:type_name -> app.protonil.testdata.v1.M3
	2,  // 3: app.protonil.testdata.v1.M2.m3_optional:type_name -> app.protonil.testdata.v1.M3
	6,  // 4: app.protonil.testdata.v1.M4.m3_map:type_name -> app.protonil.testdata.v1.M4.M3MapEntry
	2,  // 5: app.protonil.testdata.v1.M4.m3_list:type_name -> app.protonil.testdata.v1.M3
	1,  // 6: app.protonil.testdata.v1.Attack.m2:type_name -> app.protonil.testdata.v1.M2
	1,  // 7: app.protonil.testdata.v1.Attack.m2_optional:type_name -> app.protonil.testdata.v1.M2
	2,  // 8: app.protonil.testdata.v1.Attack.m3_unknown:type_name -> app.protonil.testdata.v1.M3
	2,  // 9: app.protonil.testdata.v1.Attack.m3_attack:type_name -> app.protonil.testdata.v1.M3
	2,  // 10: app.protonil.testdata.v1.M4.M3MapEntry.value:type_name -> app.protonil.testdata.v1.M3
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_app_protonil_testdata_v1_test_proto_init() }
func file_app_protonil_testdata_v1_test_proto_init() {
	if File_app_protonil_testdata_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_protonil_testdata_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M1); i {
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
		file_app_protonil_testdata_v1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2); i {
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
		file_app_protonil_testdata_v1_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M3); i {
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
		file_app_protonil_testdata_v1_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M4); i {
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
		file_app_protonil_testdata_v1_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxIndex); i {
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
		file_app_protonil_testdata_v1_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attack); i {
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
	file_app_protonil_testdata_v1_test_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_app_protonil_testdata_v1_test_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_app_protonil_testdata_v1_test_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_protonil_testdata_v1_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_protonil_testdata_v1_test_proto_goTypes,
		DependencyIndexes: file_app_protonil_testdata_v1_test_proto_depIdxs,
		MessageInfos:      file_app_protonil_testdata_v1_test_proto_msgTypes,
	}.Build()
	File_app_protonil_testdata_v1_test_proto = out.File
	file_app_protonil_testdata_v1_test_proto_rawDesc = nil
	file_app_protonil_testdata_v1_test_proto_goTypes = nil
	file_app_protonil_testdata_v1_test_proto_depIdxs = nil
}
