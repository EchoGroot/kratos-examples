// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/user-manage/v1/page.proto

package adminv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page_no与page_token二选一
	PageNo    int32  `protobuf:"varint,1,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// 页面大小
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 排序，符合SQL语法
	OrderBy string `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_manage_v1_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_manage_v1_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_api_user_manage_v1_page_proto_rawDescGZIP(), []int{0}
}

func (x *PageRequest) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *PageRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *PageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type PageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSize     int32  `protobuf:"varint,1,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *PageResponse) Reset() {
	*x = PageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_manage_v1_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResponse) ProtoMessage() {}

func (x *PageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_manage_v1_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResponse.ProtoReflect.Descriptor instead.
func (*PageResponse) Descriptor() ([]byte, []int) {
	return file_api_user_manage_v1_page_proto_rawDescGZIP(), []int{1}
}

func (x *PageResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *PageResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_api_user_manage_v1_page_proto protoreflect.FileDescriptor

var file_api_user_manage_v1_page_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01,
	0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22,
	0x55, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x67, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x63,
	0x72, 0x75, 0x64, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_manage_v1_page_proto_rawDescOnce sync.Once
	file_api_user_manage_v1_page_proto_rawDescData = file_api_user_manage_v1_page_proto_rawDesc
)

func file_api_user_manage_v1_page_proto_rawDescGZIP() []byte {
	file_api_user_manage_v1_page_proto_rawDescOnce.Do(func() {
		file_api_user_manage_v1_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_manage_v1_page_proto_rawDescData)
	})
	return file_api_user_manage_v1_page_proto_rawDescData
}

var file_api_user_manage_v1_page_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_user_manage_v1_page_proto_goTypes = []interface{}{
	(*PageRequest)(nil),  // 0: usermanage.admin.v1.PageRequest
	(*PageResponse)(nil), // 1: usermanage.admin.v1.PageResponse
}
var file_api_user_manage_v1_page_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_manage_v1_page_proto_init() }
func file_api_user_manage_v1_page_proto_init() {
	if File_api_user_manage_v1_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_manage_v1_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_api_user_manage_v1_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResponse); i {
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
			RawDescriptor: file_api_user_manage_v1_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_manage_v1_page_proto_goTypes,
		DependencyIndexes: file_api_user_manage_v1_page_proto_depIdxs,
		MessageInfos:      file_api_user_manage_v1_page_proto_msgTypes,
	}.Build()
	File_api_user_manage_v1_page_proto = out.File
	file_api_user_manage_v1_page_proto_rawDesc = nil
	file_api_user_manage_v1_page_proto_goTypes = nil
	file_api_user_manage_v1_page_proto_depIdxs = nil
}
