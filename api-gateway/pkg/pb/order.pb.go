// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pkg/proto/order.proto

package pb

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

type GetOrderHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	UserId     uint32      `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetOrderHistoryRequest) Reset() {
	*x = GetOrderHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryRequest) ProtoMessage() {}

func (x *GetOrderHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrderHistoryRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetOrderHistoryRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count      uint32 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	PageNumber uint32 `protobuf:"varint,2,opt,name=PageNumber,proto3" json:"PageNumber,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_pkg_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *Pagination) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Pagination) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

type GetOrderHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderHistoryList []*OrderHistory `protobuf:"bytes,1,rep,name=OrderHistoryList,proto3" json:"OrderHistoryList,omitempty"`
}

func (x *GetOrderHistoryResponse) Reset() {
	*x = GetOrderHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryResponse) ProtoMessage() {}

func (x *GetOrderHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderHistoryResponse) GetOrderHistoryList() []*OrderHistory {
	if x != nil {
		return x.OrderHistoryList
	}
	return nil
}

type OrderHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             uint32                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OrderDate      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=OrderDate,proto3" json:"OrderDate,omitempty"`
	OrderStatus    string                 `protobuf:"bytes,3,opt,name=OrderStatus,proto3" json:"OrderStatus,omitempty"`
	DeliveryStatus string                 `protobuf:"bytes,4,opt,name=DeliveryStatus,proto3" json:"DeliveryStatus,omitempty"`
	OrderTotal     float32                `protobuf:"fixed32,5,opt,name=OrderTotal,proto3" json:"OrderTotal,omitempty"`
	PaymentMethod  string                 `protobuf:"bytes,6,opt,name=PaymentMethod,proto3" json:"PaymentMethod,omitempty"`
	PaymentStatus  string                 `protobuf:"bytes,7,opt,name=PaymentStatus,proto3" json:"PaymentStatus,omitempty"`
}

func (x *OrderHistory) Reset() {
	*x = OrderHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderHistory) ProtoMessage() {}

func (x *OrderHistory) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderHistory.ProtoReflect.Descriptor instead.
func (*OrderHistory) Descriptor() ([]byte, []int) {
	return file_pkg_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderHistory) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *OrderHistory) GetOrderDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OrderDate
	}
	return nil
}

func (x *OrderHistory) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

func (x *OrderHistory) GetDeliveryStatus() string {
	if x != nil {
		return x.DeliveryStatus
	}
	return ""
}

func (x *OrderHistory) GetOrderTotal() float32 {
	if x != nil {
		return x.OrderTotal
	}
	return 0
}

func (x *OrderHistory) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *OrderHistory) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

var File_pkg_proto_order_proto protoreflect.FileDescriptor

var file_pkg_proto_order_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42,
	0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x57, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x8e, 0x02, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x5a, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_order_proto_rawDescOnce sync.Once
	file_pkg_proto_order_proto_rawDescData = file_pkg_proto_order_proto_rawDesc
)

func file_pkg_proto_order_proto_rawDescGZIP() []byte {
	file_pkg_proto_order_proto_rawDescOnce.Do(func() {
		file_pkg_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_order_proto_rawDescData)
	})
	return file_pkg_proto_order_proto_rawDescData
}

var file_pkg_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_order_proto_goTypes = []interface{}{
	(*GetOrderHistoryRequest)(nil),  // 0: pb.GetOrderHistoryRequest
	(*Pagination)(nil),              // 1: pb.Pagination
	(*GetOrderHistoryResponse)(nil), // 2: pb.GetOrderHistoryResponse
	(*OrderHistory)(nil),            // 3: pb.OrderHistory
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
}
var file_pkg_proto_order_proto_depIdxs = []int32{
	1, // 0: pb.GetOrderHistoryRequest.pagination:type_name -> pb.Pagination
	3, // 1: pb.GetOrderHistoryResponse.OrderHistoryList:type_name -> pb.OrderHistory
	4, // 2: pb.OrderHistory.OrderDate:type_name -> google.protobuf.Timestamp
	0, // 3: pb.OrderService.GetOrderHistory:input_type -> pb.GetOrderHistoryRequest
	2, // 4: pb.OrderService.GetOrderHistory:output_type -> pb.GetOrderHistoryResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_order_proto_init() }
func file_pkg_proto_order_proto_init() {
	if File_pkg_proto_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryRequest); i {
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
		file_pkg_proto_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_pkg_proto_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryResponse); i {
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
		file_pkg_proto_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderHistory); i {
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
			RawDescriptor: file_pkg_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_order_proto_goTypes,
		DependencyIndexes: file_pkg_proto_order_proto_depIdxs,
		MessageInfos:      file_pkg_proto_order_proto_msgTypes,
	}.Build()
	File_pkg_proto_order_proto = out.File
	file_pkg_proto_order_proto_rawDesc = nil
	file_pkg_proto_order_proto_goTypes = nil
	file_pkg_proto_order_proto_depIdxs = nil
}
