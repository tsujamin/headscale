// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: headscale/v1/apikey.proto

package v1

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

type ApiKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Prefix     string                 `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Expiration *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastSeen   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_seen,json=lastSeen,proto3" json:"last_seen,omitempty"`
}

func (x *ApiKey) Reset() {
	*x = ApiKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_apikey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKey) ProtoMessage() {}

func (x *ApiKey) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_apikey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKey.ProtoReflect.Descriptor instead.
func (*ApiKey) Descriptor() ([]byte, []int) {
	return file_headscale_v1_apikey_proto_rawDescGZIP(), []int{0}
}

func (x *ApiKey) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApiKey) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ApiKey) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *ApiKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ApiKey) GetLastSeen() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSeen
	}
	return nil
}

type CreateApiKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expiration *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *CreateApiKeyRequest) Reset() {
	*x = CreateApiKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_apikey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiKeyRequest) ProtoMessage() {}

func (x *CreateApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_apikey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_apikey_proto_rawDescGZIP(), []int{1}
}

func (x *CreateApiKeyRequest) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

type CreateApiKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *CreateApiKeyResponse) Reset() {
	*x = CreateApiKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_apikey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiKeyResponse) ProtoMessage() {}

func (x *CreateApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_apikey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_apikey_proto_rawDescGZIP(), []int{2}
}

func (x *CreateApiKeyResponse) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type ExpireApiKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *ExpireApiKeyRequest) Reset() {
	*x = ExpireApiKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_apikey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpireApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpireApiKeyRequest) ProtoMessage() {}

func (x *ExpireApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_apikey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpireApiKeyRequest.ProtoReflect.Descriptor instead.
func (*ExpireApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_apikey_proto_rawDescGZIP(), []int{3}
}

func (x *ExpireApiKeyRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type ExpireApiKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExpireApiKeyResponse) Reset() {
	*x = ExpireApiKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_apikey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpireApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpireApiKeyResponse) ProtoMessage() {}

func (x *ExpireApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_apikey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpireApiKeyResponse.ProtoReflect.Descriptor instead.
func (*ExpireApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_apikey_proto_rawDescGZIP(), []int{4}
}

type ListApiKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListApiKeysRequest) Reset() {
	*x = ListApiKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_apikey_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiKeysRequest) ProtoMessage() {}

func (x *ListApiKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_apikey_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiKeysRequest.ProtoReflect.Descriptor instead.
func (*ListApiKeysRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_apikey_proto_rawDescGZIP(), []int{5}
}

type ListApiKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKeys []*ApiKey `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys,proto3" json:"api_keys,omitempty"`
}

func (x *ListApiKeysResponse) Reset() {
	*x = ListApiKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_apikey_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiKeysResponse) ProtoMessage() {}

func (x *ListApiKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_apikey_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiKeysResponse.ProtoReflect.Descriptor instead.
func (*ListApiKeysResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_apikey_proto_rawDescGZIP(), []int{6}
}

func (x *ListApiKeysResponse) GetApiKeys() []*ApiKey {
	if x != nil {
		return x.ApiKeys
	}
	return nil
}

var File_headscale_v1_apikey_proto protoreflect.FileDescriptor

var file_headscale_v1_apikey_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x06, 0x41,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x3a, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x22, 0x51, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x2f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x22, 0x2d, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x22, 0x16, 0x0a, 0x14, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x61,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x61, 0x6e, 0x66, 0x6f, 0x6e, 0x74, 0x2f, 0x68, 0x65,
	0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_headscale_v1_apikey_proto_rawDescOnce sync.Once
	file_headscale_v1_apikey_proto_rawDescData = file_headscale_v1_apikey_proto_rawDesc
)

func file_headscale_v1_apikey_proto_rawDescGZIP() []byte {
	file_headscale_v1_apikey_proto_rawDescOnce.Do(func() {
		file_headscale_v1_apikey_proto_rawDescData = protoimpl.X.CompressGZIP(file_headscale_v1_apikey_proto_rawDescData)
	})
	return file_headscale_v1_apikey_proto_rawDescData
}

var file_headscale_v1_apikey_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_headscale_v1_apikey_proto_goTypes = []interface{}{
	(*ApiKey)(nil),                // 0: headscale.v1.ApiKey
	(*CreateApiKeyRequest)(nil),   // 1: headscale.v1.CreateApiKeyRequest
	(*CreateApiKeyResponse)(nil),  // 2: headscale.v1.CreateApiKeyResponse
	(*ExpireApiKeyRequest)(nil),   // 3: headscale.v1.ExpireApiKeyRequest
	(*ExpireApiKeyResponse)(nil),  // 4: headscale.v1.ExpireApiKeyResponse
	(*ListApiKeysRequest)(nil),    // 5: headscale.v1.ListApiKeysRequest
	(*ListApiKeysResponse)(nil),   // 6: headscale.v1.ListApiKeysResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_headscale_v1_apikey_proto_depIdxs = []int32{
	7, // 0: headscale.v1.ApiKey.expiration:type_name -> google.protobuf.Timestamp
	7, // 1: headscale.v1.ApiKey.created_at:type_name -> google.protobuf.Timestamp
	7, // 2: headscale.v1.ApiKey.last_seen:type_name -> google.protobuf.Timestamp
	7, // 3: headscale.v1.CreateApiKeyRequest.expiration:type_name -> google.protobuf.Timestamp
	0, // 4: headscale.v1.ListApiKeysResponse.api_keys:type_name -> headscale.v1.ApiKey
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_headscale_v1_apikey_proto_init() }
func file_headscale_v1_apikey_proto_init() {
	if File_headscale_v1_apikey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_headscale_v1_apikey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKey); i {
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
		file_headscale_v1_apikey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiKeyRequest); i {
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
		file_headscale_v1_apikey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiKeyResponse); i {
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
		file_headscale_v1_apikey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpireApiKeyRequest); i {
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
		file_headscale_v1_apikey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpireApiKeyResponse); i {
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
		file_headscale_v1_apikey_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiKeysRequest); i {
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
		file_headscale_v1_apikey_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiKeysResponse); i {
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
			RawDescriptor: file_headscale_v1_apikey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_headscale_v1_apikey_proto_goTypes,
		DependencyIndexes: file_headscale_v1_apikey_proto_depIdxs,
		MessageInfos:      file_headscale_v1_apikey_proto_msgTypes,
	}.Build()
	File_headscale_v1_apikey_proto = out.File
	file_headscale_v1_apikey_proto_rawDesc = nil
	file_headscale_v1_apikey_proto_goTypes = nil
	file_headscale_v1_apikey_proto_depIdxs = nil
}
