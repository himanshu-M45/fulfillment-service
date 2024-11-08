// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/fulfillment_service.proto

package fulfillment_service

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

type AddDERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *AddDERequest) Reset() {
	*x = AddDERequest{}
	mi := &file_proto_fulfillment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddDERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDERequest) ProtoMessage() {}

func (x *AddDERequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fulfillment_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDERequest.ProtoReflect.Descriptor instead.
func (*AddDERequest) Descriptor() ([]byte, []int) {
	return file_proto_fulfillment_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddDERequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type AddDEResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddDEResponse) Reset() {
	*x = AddDEResponse{}
	mi := &file_proto_fulfillment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddDEResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDEResponse) ProtoMessage() {}

func (x *AddDEResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fulfillment_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDEResponse.ProtoReflect.Descriptor instead.
func (*AddDEResponse) Descriptor() ([]byte, []int) {
	return file_proto_fulfillment_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddDEResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AssignDERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int32 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *AssignDERequest) Reset() {
	*x = AssignDERequest{}
	mi := &file_proto_fulfillment_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignDERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignDERequest) ProtoMessage() {}

func (x *AssignDERequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fulfillment_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignDERequest.ProtoReflect.Descriptor instead.
func (*AssignDERequest) Descriptor() ([]byte, []int) {
	return file_proto_fulfillment_service_proto_rawDescGZIP(), []int{2}
}

func (x *AssignDERequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type AssignDEResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AssignDEResponse) Reset() {
	*x = AssignDEResponse{}
	mi := &file_proto_fulfillment_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignDEResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignDEResponse) ProtoMessage() {}

func (x *AssignDEResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fulfillment_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignDEResponse.ProtoReflect.Descriptor instead.
func (*AssignDEResponse) Descriptor() ([]byte, []int) {
	return file_proto_fulfillment_service_proto_rawDescGZIP(), []int{3}
}

func (x *AssignDEResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryExecutiveId int32  `protobuf:"varint,1,opt,name=deliveryExecutiveId,proto3" json:"deliveryExecutiveId,omitempty"`
	Status              string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateStatusRequest) Reset() {
	*x = UpdateStatusRequest{}
	mi := &file_proto_fulfillment_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusRequest) ProtoMessage() {}

func (x *UpdateStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fulfillment_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_fulfillment_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStatusRequest) GetDeliveryExecutiveId() int32 {
	if x != nil {
		return x.DeliveryExecutiveId
	}
	return 0
}

func (x *UpdateStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateStatusResponse) Reset() {
	*x = UpdateStatusResponse{}
	mi := &file_proto_fulfillment_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusResponse) ProtoMessage() {}

func (x *UpdateStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fulfillment_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_fulfillment_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_fulfillment_service_proto protoreflect.FileDescriptor

var file_proto_fulfillment_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x44, 0x45, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x44, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a,
	0x0f, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x44, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x44, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x76, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb6, 0x02, 0x0a, 0x12,
	0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x12, 0x21, 0x2e, 0x66, 0x75, 0x6c,
	0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x44, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x57, 0x0a, 0x08, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x44, 0x45, 0x12, 0x24, 0x2e,
	0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x44, 0x45, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x44, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x28, 0x2e, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x75, 0x6c, 0x66,
	0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x66,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fulfillment_service_proto_rawDescOnce sync.Once
	file_proto_fulfillment_service_proto_rawDescData = file_proto_fulfillment_service_proto_rawDesc
)

func file_proto_fulfillment_service_proto_rawDescGZIP() []byte {
	file_proto_fulfillment_service_proto_rawDescOnce.Do(func() {
		file_proto_fulfillment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fulfillment_service_proto_rawDescData)
	})
	return file_proto_fulfillment_service_proto_rawDescData
}

var file_proto_fulfillment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_fulfillment_service_proto_goTypes = []any{
	(*AddDERequest)(nil),         // 0: fulfillment_service.AddDERequest
	(*AddDEResponse)(nil),        // 1: fulfillment_service.AddDEResponse
	(*AssignDERequest)(nil),      // 2: fulfillment_service.AssignDERequest
	(*AssignDEResponse)(nil),     // 3: fulfillment_service.AssignDEResponse
	(*UpdateStatusRequest)(nil),  // 4: fulfillment_service.UpdateStatusRequest
	(*UpdateStatusResponse)(nil), // 5: fulfillment_service.UpdateStatusResponse
}
var file_proto_fulfillment_service_proto_depIdxs = []int32{
	0, // 0: fulfillment_service.FulfillmentService.AddDeliveryExecutive:input_type -> fulfillment_service.AddDERequest
	2, // 1: fulfillment_service.FulfillmentService.AssignDE:input_type -> fulfillment_service.AssignDERequest
	4, // 2: fulfillment_service.FulfillmentService.UpdateOrderStatus:input_type -> fulfillment_service.UpdateStatusRequest
	1, // 3: fulfillment_service.FulfillmentService.AddDeliveryExecutive:output_type -> fulfillment_service.AddDEResponse
	3, // 4: fulfillment_service.FulfillmentService.AssignDE:output_type -> fulfillment_service.AssignDEResponse
	5, // 5: fulfillment_service.FulfillmentService.UpdateOrderStatus:output_type -> fulfillment_service.UpdateStatusResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_fulfillment_service_proto_init() }
func file_proto_fulfillment_service_proto_init() {
	if File_proto_fulfillment_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_fulfillment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fulfillment_service_proto_goTypes,
		DependencyIndexes: file_proto_fulfillment_service_proto_depIdxs,
		MessageInfos:      file_proto_fulfillment_service_proto_msgTypes,
	}.Build()
	File_proto_fulfillment_service_proto = out.File
	file_proto_fulfillment_service_proto_rawDesc = nil
	file_proto_fulfillment_service_proto_goTypes = nil
	file_proto_fulfillment_service_proto_depIdxs = nil
}
