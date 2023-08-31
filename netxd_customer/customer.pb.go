// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: netxd_customer/customer.proto

package netxd_customer

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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankID        int32  `protobuf:"varint,1,opt,name=BankID,proto3" json:"BankID,omitempty"`
	Customer_Name string `protobuf:"bytes,2,opt,name=Customer_Name,json=CustomerName,proto3" json:"Customer_Name,omitempty"`
	Customer_ID   int32  `protobuf:"varint,3,opt,name=Customer_ID,json=CustomerID,proto3" json:"Customer_ID,omitempty"`
	Balance       int32  `protobuf:"varint,4,opt,name=Balance,proto3" json:"Balance,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netxd_customer_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_netxd_customer_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_netxd_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetBankID() int32 {
	if x != nil {
		return x.BankID
	}
	return 0
}

func (x *Customer) GetCustomer_Name() string {
	if x != nil {
		return x.Customer_Name
	}
	return ""
}

func (x *Customer) GetCustomer_ID() int32 {
	if x != nil {
		return x.Customer_ID
	}
	return 0
}

func (x *Customer) GetBalance() int32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type CustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance int32 `protobuf:"varint,1,opt,name=Balance,proto3" json:"Balance,omitempty"`
}

func (x *CustomerResponse) Reset() {
	*x = CustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netxd_customer_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerResponse) ProtoMessage() {}

func (x *CustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_netxd_customer_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerResponse.ProtoReflect.Descriptor instead.
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return file_netxd_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerResponse) GetBalance() int32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type CustomerDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From_ID int32 `protobuf:"varint,1,opt,name=From_ID,json=FromID,proto3" json:"From_ID,omitempty"`
	TO_ID   int32 `protobuf:"varint,2,opt,name=TO_ID,json=TOID,proto3" json:"TO_ID,omitempty"`
	Amount  int32 `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *CustomerDetails) Reset() {
	*x = CustomerDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netxd_customer_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDetails) ProtoMessage() {}

func (x *CustomerDetails) ProtoReflect() protoreflect.Message {
	mi := &file_netxd_customer_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDetails.ProtoReflect.Descriptor instead.
func (*CustomerDetails) Descriptor() ([]byte, []int) {
	return file_netxd_customer_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerDetails) GetFrom_ID() int32 {
	if x != nil {
		return x.From_ID
	}
	return 0
}

func (x *CustomerDetails) GetTO_ID() int32 {
	if x != nil {
		return x.TO_ID
	}
	return 0
}

func (x *CustomerDetails) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CustomerDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From_ID int32 `protobuf:"varint,1,opt,name=From_ID,json=FromID,proto3" json:"From_ID,omitempty"`
}

func (x *CustomerDetailsResponse) Reset() {
	*x = CustomerDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netxd_customer_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDetailsResponse) ProtoMessage() {}

func (x *CustomerDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_netxd_customer_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDetailsResponse.ProtoReflect.Descriptor instead.
func (*CustomerDetailsResponse) Descriptor() ([]byte, []int) {
	return file_netxd_customer_customer_proto_rawDescGZIP(), []int{3}
}

func (x *CustomerDetailsResponse) GetFrom_ID() int32 {
	if x != nil {
		return x.From_ID
	}
	return 0
}

var File_netxd_customer_customer_proto protoreflect.FileDescriptor

var file_netxd_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22,
	0x82, 0x01, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x42, 0x61, 0x6e, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x42, 0x61,
	0x6e, 0x6b, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x46, 0x72, 0x6f, 0x6d, 0x5f, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x13,
	0x0a, 0x05, 0x54, 0x4f, 0x5f, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54,
	0x4f, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x17, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x46, 0x72, 0x6f, 0x6d, 0x5f, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44, 0x32,
	0xbb, 0x01, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a,
	0x20, 0x2e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x1a, 0x27, 0x2e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x41, 0x56, 0x41,
	0x4e, 0x45, 0x45, 0x53, 0x48, 0x4b, 0x4d, 0x2f, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_netxd_customer_customer_proto_rawDescOnce sync.Once
	file_netxd_customer_customer_proto_rawDescData = file_netxd_customer_customer_proto_rawDesc
)

func file_netxd_customer_customer_proto_rawDescGZIP() []byte {
	file_netxd_customer_customer_proto_rawDescOnce.Do(func() {
		file_netxd_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_netxd_customer_customer_proto_rawDescData)
	})
	return file_netxd_customer_customer_proto_rawDescData
}

var file_netxd_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_netxd_customer_customer_proto_goTypes = []interface{}{
	(*Customer)(nil),                // 0: netxd_customer.Customer
	(*CustomerResponse)(nil),        // 1: netxd_customer.CustomerResponse
	(*CustomerDetails)(nil),         // 2: netxd_customer.CustomerDetails
	(*CustomerDetailsResponse)(nil), // 3: netxd_customer.CustomerDetailsResponse
}
var file_netxd_customer_customer_proto_depIdxs = []int32{
	0, // 0: netxd_customer.CustomerService.CreateCustomer:input_type -> netxd_customer.Customer
	2, // 1: netxd_customer.CustomerService.UpdateCustomer:input_type -> netxd_customer.CustomerDetails
	1, // 2: netxd_customer.CustomerService.CreateCustomer:output_type -> netxd_customer.CustomerResponse
	3, // 3: netxd_customer.CustomerService.UpdateCustomer:output_type -> netxd_customer.CustomerDetailsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_netxd_customer_customer_proto_init() }
func file_netxd_customer_customer_proto_init() {
	if File_netxd_customer_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_netxd_customer_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_netxd_customer_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerResponse); i {
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
		file_netxd_customer_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDetails); i {
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
		file_netxd_customer_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDetailsResponse); i {
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
			RawDescriptor: file_netxd_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_netxd_customer_customer_proto_goTypes,
		DependencyIndexes: file_netxd_customer_customer_proto_depIdxs,
		MessageInfos:      file_netxd_customer_customer_proto_msgTypes,
	}.Build()
	File_netxd_customer_customer_proto = out.File
	file_netxd_customer_customer_proto_rawDesc = nil
	file_netxd_customer_customer_proto_goTypes = nil
	file_netxd_customer_customer_proto_depIdxs = nil
}