// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc2
// source: client.proto

package services

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

type Phone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country int32 `protobuf:"varint,1,opt,name=country,proto3" json:"country,omitempty"`
	Code    int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Number  int32 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Phone) Reset() {
	*x = Phone{}
	mi := &file_client_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Phone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phone) ProtoMessage() {}

func (x *Phone) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phone.ProtoReflect.Descriptor instead.
func (*Phone) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *Phone) GetCountry() int32 {
	if x != nil {
		return x.Country
	}
	return 0
}

func (x *Phone) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Phone) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CreateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string   `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName   string   `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	MiddleName string   `protobuf:"bytes,3,opt,name=middleName,proto3" json:"middleName,omitempty"`
	Email      string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phones     []*Phone `protobuf:"bytes,5,rep,name=phones,proto3" json:"phones,omitempty"`
	Status     string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateClientRequest) Reset() {
	*x = CreateClientRequest{}
	mi := &file_client_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientRequest) ProtoMessage() {}

func (x *CreateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientRequest.ProtoReflect.Descriptor instead.
func (*CreateClientRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClientRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateClientRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateClientRequest) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *CreateClientRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateClientRequest) GetPhones() []*Phone {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *CreateClientRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventID string `protobuf:"bytes,2,opt,name=eventID,proto3" json:"eventID,omitempty"`
}

func (x *CreateClientResponse) Reset() {
	*x = CreateClientResponse{}
	mi := &file_client_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientResponse) ProtoMessage() {}

func (x *CreateClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientResponse.ProtoReflect.Descriptor instead.
func (*CreateClientResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClientResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateClientResponse) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d,
	0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xbd, 0x01,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x06, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x32,
	0x4e, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_client_proto_goTypes = []any{
	(*Phone)(nil),                // 0: Phone
	(*CreateClientRequest)(nil),  // 1: CreateClientRequest
	(*CreateClientResponse)(nil), // 2: CreateClientResponse
}
var file_client_proto_depIdxs = []int32{
	0, // 0: CreateClientRequest.phones:type_name -> Phone
	1, // 1: ClientService.CreateCustomer:input_type -> CreateClientRequest
	2, // 2: ClientService.CreateCustomer:output_type -> CreateClientResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}