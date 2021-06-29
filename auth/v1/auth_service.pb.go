// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.2
// source: auth/v1/auth_service.proto

package auth_grpc

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

type GetPayloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetPayloadRequest) Reset() {
	*x = GetPayloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPayloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPayloadRequest) ProtoMessage() {}

func (x *GetPayloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPayloadRequest.ProtoReflect.Descriptor instead.
func (*GetPayloadRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPayloadRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetPayloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PhoneNumber string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Identity    int32  `protobuf:"varint,4,opt,name=identity,proto3" json:"identity,omitempty"`
	Role        string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Platform    string `protobuf:"bytes,6,opt,name=platform,proto3" json:"platform,omitempty"`
	Iat         int64  `protobuf:"varint,7,opt,name=iat,proto3" json:"iat,omitempty"`
	Exp         int64  `protobuf:"varint,8,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (x *GetPayloadResponse) Reset() {
	*x = GetPayloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPayloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPayloadResponse) ProtoMessage() {}

func (x *GetPayloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPayloadResponse.ProtoReflect.Descriptor instead.
func (*GetPayloadResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPayloadResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetPayloadResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *GetPayloadResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetPayloadResponse) GetIdentity() int32 {
	if x != nil {
		return x.Identity
	}
	return 0
}

func (x *GetPayloadResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetPayloadResponse) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *GetPayloadResponse) GetIat() int64 {
	if x != nil {
		return x.Iat
	}
	return 0
}

func (x *GetPayloadResponse) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

var File_auth_v1_auth_service_proto protoreflect.FileDescriptor

var file_auth_v1_auth_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xcd, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x61, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70,
	0x32, 0x54, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x76, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x65, 0x61,
	0x6c, 0x69, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_auth_service_proto_rawDescOnce sync.Once
	file_auth_v1_auth_service_proto_rawDescData = file_auth_v1_auth_service_proto_rawDesc
)

func file_auth_v1_auth_service_proto_rawDescGZIP() []byte {
	file_auth_v1_auth_service_proto_rawDescOnce.Do(func() {
		file_auth_v1_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_auth_service_proto_rawDescData)
	})
	return file_auth_v1_auth_service_proto_rawDescData
}

var file_auth_v1_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_v1_auth_service_proto_goTypes = []interface{}{
	(*GetPayloadRequest)(nil),  // 0: auth.v1.GetPayloadRequest
	(*GetPayloadResponse)(nil), // 1: auth.v1.GetPayloadResponse
}
var file_auth_v1_auth_service_proto_depIdxs = []int32{
	0, // 0: auth.v1.AuthService.GetPayload:input_type -> auth.v1.GetPayloadRequest
	1, // 1: auth.v1.AuthService.GetPayload:output_type -> auth.v1.GetPayloadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_v1_auth_service_proto_init() }
func file_auth_v1_auth_service_proto_init() {
	if File_auth_v1_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_v1_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPayloadRequest); i {
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
		file_auth_v1_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPayloadResponse); i {
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
			RawDescriptor: file_auth_v1_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_v1_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_v1_auth_service_proto_depIdxs,
		MessageInfos:      file_auth_v1_auth_service_proto_msgTypes,
	}.Build()
	File_auth_v1_auth_service_proto = out.File
	file_auth_v1_auth_service_proto_rawDesc = nil
	file_auth_v1_auth_service_proto_goTypes = nil
	file_auth_v1_auth_service_proto_depIdxs = nil
}
