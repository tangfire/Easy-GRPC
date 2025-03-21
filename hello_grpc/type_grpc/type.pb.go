// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: hello_grpc/type.proto

package type_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A1  float64 `protobuf:"fixed64,1,opt,name=a1,proto3" json:"a1,omitempty"`
	A2  float32 `protobuf:"fixed32,2,opt,name=a2,proto3" json:"a2,omitempty"`
	A3  int32   `protobuf:"varint,3,opt,name=a3,proto3" json:"a3,omitempty"`
	A4  uint32  `protobuf:"varint,4,opt,name=a4,proto3" json:"a4,omitempty"`
	A5  uint64  `protobuf:"varint,5,opt,name=a5,proto3" json:"a5,omitempty"`
	A6  int32   `protobuf:"zigzag32,6,opt,name=a6,proto3" json:"a6,omitempty"`
	A7  int64   `protobuf:"zigzag64,7,opt,name=a7,proto3" json:"a7,omitempty"`
	A8  uint32  `protobuf:"fixed32,8,opt,name=a8,proto3" json:"a8,omitempty"`
	A9  uint64  `protobuf:"fixed64,9,opt,name=a9,proto3" json:"a9,omitempty"`
	A10 int32   `protobuf:"fixed32,10,opt,name=a10,proto3" json:"a10,omitempty"`
	A11 int64   `protobuf:"fixed64,11,opt,name=a11,proto3" json:"a11,omitempty"`
	A12 bool    `protobuf:"varint,12,opt,name=a12,proto3" json:"a12,omitempty"`
	A13 string  `protobuf:"bytes,13,opt,name=a13,proto3" json:"a13,omitempty"`
	A14 []byte  `protobuf:"bytes,14,opt,name=a14,proto3" json:"a14,omitempty"` // []bytes
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_grpc_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_hello_grpc_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_hello_grpc_type_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetA1() float64 {
	if x != nil {
		return x.A1
	}
	return 0
}

func (x *Request) GetA2() float32 {
	if x != nil {
		return x.A2
	}
	return 0
}

func (x *Request) GetA3() int32 {
	if x != nil {
		return x.A3
	}
	return 0
}

func (x *Request) GetA4() uint32 {
	if x != nil {
		return x.A4
	}
	return 0
}

func (x *Request) GetA5() uint64 {
	if x != nil {
		return x.A5
	}
	return 0
}

func (x *Request) GetA6() int32 {
	if x != nil {
		return x.A6
	}
	return 0
}

func (x *Request) GetA7() int64 {
	if x != nil {
		return x.A7
	}
	return 0
}

func (x *Request) GetA8() uint32 {
	if x != nil {
		return x.A8
	}
	return 0
}

func (x *Request) GetA9() uint64 {
	if x != nil {
		return x.A9
	}
	return 0
}

func (x *Request) GetA10() int32 {
	if x != nil {
		return x.A10
	}
	return 0
}

func (x *Request) GetA11() int64 {
	if x != nil {
		return x.A11
	}
	return 0
}

func (x *Request) GetA12() bool {
	if x != nil {
		return x.A12
	}
	return false
}

func (x *Request) GetA13() string {
	if x != nil {
		return x.A13
	}
	return ""
}

func (x *Request) GetA14() []byte {
	if x != nil {
		return x.A14
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code uint32 `protobuf:"fixed32,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_grpc_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_hello_grpc_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_hello_grpc_type_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type ArrayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I6List   []int64  `protobuf:"varint,1,rep,packed,name=i6_list,json=i6List,proto3" json:"i6_list,omitempty"`
	SList    []string `protobuf:"bytes,2,rep,name=s_list,json=sList,proto3" json:"s_list,omitempty"`
	ItemList []*Item  `protobuf:"bytes,3,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *ArrayRequest) Reset() {
	*x = ArrayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_grpc_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayRequest) ProtoMessage() {}

func (x *ArrayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_grpc_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayRequest.ProtoReflect.Descriptor instead.
func (*ArrayRequest) Descriptor() ([]byte, []int) {
	return file_hello_grpc_type_proto_rawDescGZIP(), []int{2}
}

func (x *ArrayRequest) GetI6List() []int64 {
	if x != nil {
		return x.I6List
	}
	return nil
}

func (x *ArrayRequest) GetSList() []string {
	if x != nil {
		return x.SList
	}
	return nil
}

func (x *ArrayRequest) GetItemList() []*Item {
	if x != nil {
		return x.ItemList
	}
	return nil
}

type MapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IS    map[int32]string  `protobuf:"bytes,1,rep,name=i_s,json=iS,proto3" json:"i_s,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SS    map[string]string `protobuf:"bytes,2,rep,name=s_s,json=sS,proto3" json:"s_s,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SItem map[string]*Item  `protobuf:"bytes,3,rep,name=s_item,json=sItem,proto3" json:"s_item,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapRequest) Reset() {
	*x = MapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_grpc_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRequest) ProtoMessage() {}

func (x *MapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_grpc_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRequest.ProtoReflect.Descriptor instead.
func (*MapRequest) Descriptor() ([]byte, []int) {
	return file_hello_grpc_type_proto_rawDescGZIP(), []int{3}
}

func (x *MapRequest) GetIS() map[int32]string {
	if x != nil {
		return x.IS
	}
	return nil
}

func (x *MapRequest) GetSS() map[string]string {
	if x != nil {
		return x.SS
	}
	return nil
}

func (x *MapRequest) GetSItem() map[string]*Item {
	if x != nil {
		return x.SItem
	}
	return nil
}

type Q1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Q2   *Q1_Q2 `protobuf:"bytes,2,opt,name=q2,proto3" json:"q2,omitempty"`
}

func (x *Q1) Reset() {
	*x = Q1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_grpc_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Q1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Q1) ProtoMessage() {}

func (x *Q1) ProtoReflect() protoreflect.Message {
	mi := &file_hello_grpc_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Q1.ProtoReflect.Descriptor instead.
func (*Q1) Descriptor() ([]byte, []int) {
	return file_hello_grpc_type_proto_rawDescGZIP(), []int{4}
}

func (x *Q1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Q1) GetQ2() *Q1_Q2 {
	if x != nil {
		return x.Q2
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_grpc_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_hello_grpc_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_hello_grpc_type_proto_rawDescGZIP(), []int{5}
}

type Q1_Q2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Q1_Q2) Reset() {
	*x = Q1_Q2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_grpc_type_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Q1_Q2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Q1_Q2) ProtoMessage() {}

func (x *Q1_Q2) ProtoReflect() protoreflect.Message {
	mi := &file_hello_grpc_type_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Q1_Q2.ProtoReflect.Descriptor instead.
func (*Q1_Q2) Descriptor() ([]byte, []int) {
	return file_hello_grpc_type_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Q1_Q2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_hello_grpc_type_proto protoreflect.FileDescriptor

var file_hello_grpc_type_proto_rawDesc = []byte{
	0x0a, 0x15, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x02, 0x61, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x02, 0x61, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x61, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x61, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x61, 0x35, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x02, 0x61, 0x36, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x02, 0x61, 0x37, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x38, 0x18, 0x08, 0x20, 0x01, 0x28, 0x07, 0x52,
	0x02, 0x61, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x39, 0x18, 0x09, 0x20, 0x01, 0x28, 0x06, 0x52,
	0x02, 0x61, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x31, 0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0f,
	0x52, 0x03, 0x61, 0x31, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x10, 0x52, 0x03, 0x61, 0x31, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x31, 0x32, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x31, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x31, 0x33,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x31, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x31, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x61, 0x31, 0x34, 0x22, 0x2e, 0x0a,
	0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x62, 0x0a,
	0x0c, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x36, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06,
	0x69, 0x36, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0xb6, 0x02, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x03, 0x69, 0x5f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x53, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x02, 0x69, 0x53, 0x12, 0x24, 0x0a, 0x03, 0x73, 0x5f, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x73, 0x53, 0x12, 0x2d, 0x0a, 0x06,
	0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x49, 0x74, 0x65, 0x6d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x35, 0x0a, 0x07, 0x49,
	0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x35, 0x0a, 0x07, 0x53, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x0a, 0x53, 0x49, 0x74,
	0x65, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x02, 0x51, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x02, 0x71, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x51, 0x31, 0x2e, 0x51, 0x32, 0x52, 0x02, 0x71, 0x32, 0x1a, 0x18, 0x0a, 0x02,
	0x51, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x2b, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_grpc_type_proto_rawDescOnce sync.Once
	file_hello_grpc_type_proto_rawDescData = file_hello_grpc_type_proto_rawDesc
)

func file_hello_grpc_type_proto_rawDescGZIP() []byte {
	file_hello_grpc_type_proto_rawDescOnce.Do(func() {
		file_hello_grpc_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_grpc_type_proto_rawDescData)
	})
	return file_hello_grpc_type_proto_rawDescData
}

var file_hello_grpc_type_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_hello_grpc_type_proto_goTypes = []interface{}{
	(*Request)(nil),      // 0: Request
	(*Item)(nil),         // 1: Item
	(*ArrayRequest)(nil), // 2: ArrayRequest
	(*MapRequest)(nil),   // 3: MapRequest
	(*Q1)(nil),           // 4: Q1
	(*Response)(nil),     // 5: Response
	nil,                  // 6: MapRequest.ISEntry
	nil,                  // 7: MapRequest.SSEntry
	nil,                  // 8: MapRequest.SItemEntry
	(*Q1_Q2)(nil),        // 9: Q1.Q2
}
var file_hello_grpc_type_proto_depIdxs = []int32{
	1, // 0: ArrayRequest.item_list:type_name -> Item
	6, // 1: MapRequest.i_s:type_name -> MapRequest.ISEntry
	7, // 2: MapRequest.s_s:type_name -> MapRequest.SSEntry
	8, // 3: MapRequest.s_item:type_name -> MapRequest.SItemEntry
	9, // 4: Q1.q2:type_name -> Q1.Q2
	1, // 5: MapRequest.SItemEntry.value:type_name -> Item
	0, // 6: TypeService.Say:input_type -> Request
	5, // 7: TypeService.Say:output_type -> Response
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_hello_grpc_type_proto_init() }
func file_hello_grpc_type_proto_init() {
	if File_hello_grpc_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_grpc_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_hello_grpc_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_hello_grpc_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayRequest); i {
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
		file_hello_grpc_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRequest); i {
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
		file_hello_grpc_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Q1); i {
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
		file_hello_grpc_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_hello_grpc_type_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Q1_Q2); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_grpc_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_grpc_type_proto_goTypes,
		DependencyIndexes: file_hello_grpc_type_proto_depIdxs,
		MessageInfos:      file_hello_grpc_type_proto_msgTypes,
	}.Build()
	File_hello_grpc_type_proto = out.File
	file_hello_grpc_type_proto_rawDesc = nil
	file_hello_grpc_type_proto_goTypes = nil
	file_hello_grpc_type_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TypeServiceClient is the client API for TypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TypeServiceClient interface {
	Say(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type typeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTypeServiceClient(cc grpc.ClientConnInterface) TypeServiceClient {
	return &typeServiceClient{cc}
}

func (c *typeServiceClient) Say(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/TypeService/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TypeServiceServer is the server API for TypeService service.
type TypeServiceServer interface {
	Say(context.Context, *Request) (*Response, error)
}

// UnimplementedTypeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTypeServiceServer struct {
}

func (*UnimplementedTypeServiceServer) Say(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}

func RegisterTypeServiceServer(s *grpc.Server, srv TypeServiceServer) {
	s.RegisterService(&_TypeService_serviceDesc, srv)
}

func _TypeService_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeServiceServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TypeService/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeServiceServer).Say(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _TypeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TypeService",
	HandlerType: (*TypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _TypeService_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_grpc/type.proto",
}
