// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pkg/proto/auth.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// User SignUp
type SignupDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName  string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age       uint32 `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Password  string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignupDataRequest) Reset() {
	*x = SignupDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupDataRequest) ProtoMessage() {}

func (x *SignupDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupDataRequest.ProtoReflect.Descriptor instead.
func (*SignupDataRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignupDataRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *SignupDataRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SignupDataRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SignupDataRequest) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *SignupDataRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupDataRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignupDataRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// User Login
type LoginDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginDataRequest) Reset() {
	*x = LoginDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginDataRequest) ProtoMessage() {}

func (x *LoginDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginDataRequest.ProtoReflect.Descriptor instead.
func (*LoginDataRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginDataRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginDataRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginDataRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LoginDataResponse) Reset() {
	*x = LoginDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginDataResponse) ProtoMessage() {}

func (x *LoginDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginDataResponse.ProtoReflect.Descriptor instead.
func (*LoginDataResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginDataResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// OTP Login
type OTPLoginDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Otp int32 `protobuf:"varint,1,opt,name=otp,proto3" json:"otp,omitempty"` // One Time Password generated by the system or sent to mobile
}

func (x *OTPLoginDataRequest) Reset() {
	*x = OTPLoginDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPLoginDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPLoginDataRequest) ProtoMessage() {}

func (x *OTPLoginDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPLoginDataRequest.ProtoReflect.Descriptor instead.
func (*OTPLoginDataRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *OTPLoginDataRequest) GetOtp() int32 {
	if x != nil {
		return x.Otp
	}
	return 0
}

type OTPLoginDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age       uint32 `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *OTPLoginDataResponse) Reset() {
	*x = OTPLoginDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPLoginDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPLoginDataResponse) ProtoMessage() {}

func (x *OTPLoginDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPLoginDataResponse.ProtoReflect.Descriptor instead.
func (*OTPLoginDataResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *OTPLoginDataResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OTPLoginDataResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *OTPLoginDataResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *OTPLoginDataResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *OTPLoginDataResponse) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *OTPLoginDataResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OTPLoginDataResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_pkg_proto_auth_proto protoreflect.FileDescriptor

var file_pkg_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x60, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x13, 0x4f, 0x54, 0x50, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x74, 0x70,
	0x22, 0xbd, 0x01, 0x0a, 0x14, 0x4f, 0x54, 0x50, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x32, 0xbb, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x08, 0x4f, 0x54, 0x50, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x4f, 0x54, 0x50, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x54, 0x50, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_proto_auth_proto_rawDescOnce sync.Once
	file_pkg_proto_auth_proto_rawDescData = file_pkg_proto_auth_proto_rawDesc
)

func file_pkg_proto_auth_proto_rawDescGZIP() []byte {
	file_pkg_proto_auth_proto_rawDescOnce.Do(func() {
		file_pkg_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_auth_proto_rawDescData)
	})
	return file_pkg_proto_auth_proto_rawDescData
}

var file_pkg_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_auth_proto_goTypes = []interface{}{
	(*SignupDataRequest)(nil),    // 0: pb.SignupDataRequest
	(*LoginDataRequest)(nil),     // 1: pb.LoginDataRequest
	(*LoginDataResponse)(nil),    // 2: pb.LoginDataResponse
	(*OTPLoginDataRequest)(nil),  // 3: pb.OTPLoginDataRequest
	(*OTPLoginDataResponse)(nil), // 4: pb.OTPLoginDataResponse
	(*emptypb.Empty)(nil),        // 5: google.protobuf.Empty
}
var file_pkg_proto_auth_proto_depIdxs = []int32{
	0, // 0: pb.AuthService.SignUp:input_type -> pb.SignupDataRequest
	1, // 1: pb.AuthService.Login:input_type -> pb.LoginDataRequest
	3, // 2: pb.AuthService.OTPLogin:input_type -> pb.OTPLoginDataRequest
	5, // 3: pb.AuthService.SignUp:output_type -> google.protobuf.Empty
	2, // 4: pb.AuthService.Login:output_type -> pb.LoginDataResponse
	4, // 5: pb.AuthService.OTPLogin:output_type -> pb.OTPLoginDataResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_auth_proto_init() }
func file_pkg_proto_auth_proto_init() {
	if File_pkg_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupDataRequest); i {
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
		file_pkg_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginDataRequest); i {
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
		file_pkg_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginDataResponse); i {
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
		file_pkg_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPLoginDataRequest); i {
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
		file_pkg_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPLoginDataResponse); i {
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
			RawDescriptor: file_pkg_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_auth_proto_goTypes,
		DependencyIndexes: file_pkg_proto_auth_proto_depIdxs,
		MessageInfos:      file_pkg_proto_auth_proto_msgTypes,
	}.Build()
	File_pkg_proto_auth_proto = out.File
	file_pkg_proto_auth_proto_rawDesc = nil
	file_pkg_proto_auth_proto_goTypes = nil
	file_pkg_proto_auth_proto_depIdxs = nil
}