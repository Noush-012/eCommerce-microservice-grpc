// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pkg/proto/user.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserDataRequest) Reset() {
	*x = UserDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataRequest) ProtoMessage() {}

func (x *UserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataRequest.ProtoReflect.Descriptor instead.
func (*UserDataRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserDataRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDataRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserDataRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserDataRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UserDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName   string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age         uint32 `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone       string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	BlockStatus bool   `protobuf:"varint,8,opt,name=block_status,json=blockStatus,proto3" json:"block_status,omitempty"`
	Password    string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserDataResponse) Reset() {
	*x = UserDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataResponse) ProtoMessage() {}

func (x *UserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataResponse.ProtoReflect.Descriptor instead.
func (*UserDataResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserDataResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDataResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserDataResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserDataResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserDataResponse) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserDataResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserDataResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserDataResponse) GetBlockStatus() bool {
	if x != nil {
		return x.BlockStatus
	}
	return false
}

func (x *UserDataResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_pkg_proto_user_proto protoreflect.FileDescriptor

var file_pkg_proto_user_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x22, 0xf8, 0x01, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x44,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_user_proto_rawDescOnce sync.Once
	file_pkg_proto_user_proto_rawDescData = file_pkg_proto_user_proto_rawDesc
)

func file_pkg_proto_user_proto_rawDescGZIP() []byte {
	file_pkg_proto_user_proto_rawDescOnce.Do(func() {
		file_pkg_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_user_proto_rawDescData)
	})
	return file_pkg_proto_user_proto_rawDescData
}

var file_pkg_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_user_proto_goTypes = []interface{}{
	(*UserDataRequest)(nil),  // 0: pb.userDataRequest
	(*UserDataResponse)(nil), // 1: pb.userDataResponse
}
var file_pkg_proto_user_proto_depIdxs = []int32{
	0, // 0: pb.UserService.FindUser:input_type -> pb.userDataRequest
	1, // 1: pb.UserService.FindUser:output_type -> pb.userDataResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_user_proto_init() }
func file_pkg_proto_user_proto_init() {
	if File_pkg_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataRequest); i {
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
		file_pkg_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataResponse); i {
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
			RawDescriptor: file_pkg_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_user_proto_goTypes,
		DependencyIndexes: file_pkg_proto_user_proto_depIdxs,
		MessageInfos:      file_pkg_proto_user_proto_msgTypes,
	}.Build()
	File_pkg_proto_user_proto = out.File
	file_pkg_proto_user_proto_rawDesc = nil
	file_pkg_proto_user_proto_goTypes = nil
	file_pkg_proto_user_proto_depIdxs = nil
}
