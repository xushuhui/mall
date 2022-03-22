// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: api/mall/admin/admin.proto

package admin

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mall_admin_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_mall_admin_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_api_mall_admin_admin_proto_rawDescGZIP(), []int{0}
}

func (x *IdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Banner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Title       string        `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Img         string        `protobuf:"bytes,4,opt,name=img,proto3" json:"img,omitempty"`
	Description string        `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Items       []*BannerItem `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Banner) Reset() {
	*x = Banner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mall_admin_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Banner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Banner) ProtoMessage() {}

func (x *Banner) ProtoReflect() protoreflect.Message {
	mi := &file_api_mall_admin_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Banner.ProtoReflect.Descriptor instead.
func (*Banner) Descriptor() ([]byte, []int) {
	return file_api_mall_admin_admin_proto_rawDescGZIP(), []int{1}
}

func (x *Banner) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Banner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Banner) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Banner) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *Banner) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Banner) GetItems() []*BannerItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type BannerItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Img      string `protobuf:"bytes,2,opt,name=img,proto3" json:"img,omitempty"`
	Keyword  string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Type     int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Name     string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	BannerId int64  `protobuf:"varint,6,opt,name=bannerId,proto3" json:"bannerId,omitempty"`
}

func (x *BannerItem) Reset() {
	*x = BannerItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mall_admin_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerItem) ProtoMessage() {}

func (x *BannerItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_mall_admin_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerItem.ProtoReflect.Descriptor instead.
func (*BannerItem) Descriptor() ([]byte, []int) {
	return file_api_mall_admin_admin_proto_rawDescGZIP(), []int{2}
}

func (x *BannerItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BannerItem) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *BannerItem) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *BannerItem) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *BannerItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BannerItem) GetBannerId() int64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

var File_api_mall_admin_admin_proto protoreflect.FileDescriptor

var file_api_mall_admin_admin_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61,
	0x6c, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xa4, 0x01, 0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x32, 0x5c, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x53, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x15, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x69, 0x64, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0x18, 0x5a, 0x16, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_mall_admin_admin_proto_rawDescOnce sync.Once
	file_api_mall_admin_admin_proto_rawDescData = file_api_mall_admin_admin_proto_rawDesc
)

func file_api_mall_admin_admin_proto_rawDescGZIP() []byte {
	file_api_mall_admin_admin_proto_rawDescOnce.Do(func() {
		file_api_mall_admin_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mall_admin_admin_proto_rawDescData)
	})
	return file_api_mall_admin_admin_proto_rawDescData
}

var file_api_mall_admin_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_mall_admin_admin_proto_goTypes = []interface{}{
	(*IdRequest)(nil),  // 0: mall.admin.IdRequest
	(*Banner)(nil),     // 1: mall.admin.Banner
	(*BannerItem)(nil), // 2: mall.admin.BannerItem
}
var file_api_mall_admin_admin_proto_depIdxs = []int32{
	2, // 0: mall.admin.Banner.items:type_name -> mall.admin.BannerItem
	0, // 1: mall.admin.Admin.GetBannerById:input_type -> mall.admin.IdRequest
	1, // 2: mall.admin.Admin.GetBannerById:output_type -> mall.admin.Banner
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_mall_admin_admin_proto_init() }
func file_api_mall_admin_admin_proto_init() {
	if File_api_mall_admin_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_mall_admin_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_api_mall_admin_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Banner); i {
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
		file_api_mall_admin_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerItem); i {
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
			RawDescriptor: file_api_mall_admin_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_mall_admin_admin_proto_goTypes,
		DependencyIndexes: file_api_mall_admin_admin_proto_depIdxs,
		MessageInfos:      file_api_mall_admin_admin_proto_msgTypes,
	}.Build()
	File_api_mall_admin_admin_proto = out.File
	file_api_mall_admin_admin_proto_rawDesc = nil
	file_api_mall_admin_admin_proto_goTypes = nil
	file_api_mall_admin_admin_proto_depIdxs = nil
}
