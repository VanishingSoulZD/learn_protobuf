// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: apps/pbstock/create_stock.proto

package pbstock

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

type CreateStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stock Code
	StockCode string `protobuf:"bytes,1,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
	// Stock Name
	StockName string `protobuf:"bytes,2,opt,name=stock_name,json=stockName,proto3" json:"stock_name,omitempty"`
	// Price
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	// Stock Type, Enum: StockType
	StockType int64 `protobuf:"varint,4,opt,name=stock_type,json=stockType,proto3" json:"stock_type,omitempty"`
	// One Hand
	OneHand int64 `protobuf:"varint,5,opt,name=one_hand,json=oneHand,proto3" json:"one_hand,omitempty"`
}

func (x *CreateStockRequest) Reset() {
	*x = CreateStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pbstock_create_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStockRequest) ProtoMessage() {}

func (x *CreateStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pbstock_create_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStockRequest.ProtoReflect.Descriptor instead.
func (*CreateStockRequest) Descriptor() ([]byte, []int) {
	return file_apps_pbstock_create_stock_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStockRequest) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

func (x *CreateStockRequest) GetStockName() string {
	if x != nil {
		return x.StockName
	}
	return ""
}

func (x *CreateStockRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateStockRequest) GetStockType() int64 {
	if x != nil {
		return x.StockType
	}
	return 0
}

func (x *CreateStockRequest) GetOneHand() int64 {
	if x != nil {
		return x.OneHand
	}
	return 0
}

type CreateStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of stock_tab
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateStockResponse) Reset() {
	*x = CreateStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pbstock_create_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStockResponse) ProtoMessage() {}

func (x *CreateStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pbstock_create_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStockResponse.ProtoReflect.Descriptor instead.
func (*CreateStockResponse) Descriptor() ([]byte, []int) {
	return file_apps_pbstock_create_stock_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStockResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_apps_pbstock_create_stock_proto protoreflect.FileDescriptor

var file_apps_pbstock_create_stock_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x62, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x70, 0x62, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0xa2, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6e, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x6e, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x22,
	0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x70, 0x62, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_pbstock_create_stock_proto_rawDescOnce sync.Once
	file_apps_pbstock_create_stock_proto_rawDescData = file_apps_pbstock_create_stock_proto_rawDesc
)

func file_apps_pbstock_create_stock_proto_rawDescGZIP() []byte {
	file_apps_pbstock_create_stock_proto_rawDescOnce.Do(func() {
		file_apps_pbstock_create_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_pbstock_create_stock_proto_rawDescData)
	})
	return file_apps_pbstock_create_stock_proto_rawDescData
}

var file_apps_pbstock_create_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_pbstock_create_stock_proto_goTypes = []interface{}{
	(*CreateStockRequest)(nil),  // 0: pbstock.CreateStockRequest
	(*CreateStockResponse)(nil), // 1: pbstock.CreateStockResponse
}
var file_apps_pbstock_create_stock_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_pbstock_create_stock_proto_init() }
func file_apps_pbstock_create_stock_proto_init() {
	if File_apps_pbstock_create_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_pbstock_create_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStockRequest); i {
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
		file_apps_pbstock_create_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStockResponse); i {
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
			RawDescriptor: file_apps_pbstock_create_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_pbstock_create_stock_proto_goTypes,
		DependencyIndexes: file_apps_pbstock_create_stock_proto_depIdxs,
		MessageInfos:      file_apps_pbstock_create_stock_proto_msgTypes,
	}.Build()
	File_apps_pbstock_create_stock_proto = out.File
	file_apps_pbstock_create_stock_proto_rawDesc = nil
	file_apps_pbstock_create_stock_proto_goTypes = nil
	file_apps_pbstock_create_stock_proto_depIdxs = nil
}
