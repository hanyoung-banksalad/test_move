// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.3
// source: imageproxy.proto

package imageproxy

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imageproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imageproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_imageproxy_proto_rawDescGZIP(), []int{0}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imageproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imageproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_imageproxy_proto_rawDescGZIP(), []int{1}
}

type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Bg       string `protobuf:"bytes,2,opt,name=bg,proto3" json:"bg,omitempty"`
	Shape    string `protobuf:"bytes,3,opt,name=shape,proto3" json:"shape,omitempty"`
	Size     int32  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imageproxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imageproxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_imageproxy_proto_rawDescGZIP(), []int{2}
}

func (x *GetImageRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *GetImageRequest) GetBg() string {
	if x != nil {
		return x.Bg
	}
	return ""
}

func (x *GetImageRequest) GetShape() string {
	if x != nil {
		return x.Shape
	}
	return ""
}

func (x *GetImageRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type GetImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *GetImageResponse) Reset() {
	*x = GetImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imageproxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageResponse) ProtoMessage() {}

func (x *GetImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imageproxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageResponse.ProtoReflect.Descriptor instead.
func (*GetImageResponse) Descriptor() ([]byte, []int) {
	return file_imageproxy_proto_rawDescGZIP(), []int{3}
}

func (x *GetImageResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

var File_imageproxy_proto protoreflect.FileDescriptor

var file_imageproxy_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x14, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x62,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x62, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x68, 0x61, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x32, 0xac, 0x02, 0x0a,
	0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x65, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0xb6, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x69, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x63, 0x12, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2f, 0x7b, 0x73, 0x69, 0x7a, 0x65, 0x7d, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x5a, 0x35, 0x12, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x65, 0x74, 0x2d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x7b, 0x73, 0x69, 0x7a, 0x65, 0x7d,
	0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x86, 0x01, 0x0a, 0x2b,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x73,
	0x61, 0x6c, 0x61, 0x64, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x42, 0x0f, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6e, 0x79, 0x6f, 0x75, 0x6e,
	0x67, 0x2d, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x61, 0x6c, 0x61, 0x64, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imageproxy_proto_rawDescOnce sync.Once
	file_imageproxy_proto_rawDescData = file_imageproxy_proto_rawDesc
)

func file_imageproxy_proto_rawDescGZIP() []byte {
	file_imageproxy_proto_rawDescOnce.Do(func() {
		file_imageproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_imageproxy_proto_rawDescData)
	})
	return file_imageproxy_proto_rawDescData
}

var file_imageproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_imageproxy_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),  // 0: v1.imageproxy.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 1: v1.imageproxy.HealthCheckResponse
	(*GetImageRequest)(nil),     // 2: v1.imageproxy.GetImageRequest
	(*GetImageResponse)(nil),    // 3: v1.imageproxy.GetImageResponse
}
var file_imageproxy_proto_depIdxs = []int32{
	0, // 0: v1.imageproxy.Imageproxy.HealthCheck:input_type -> v1.imageproxy.HealthCheckRequest
	2, // 1: v1.imageproxy.Imageproxy.GetImage:input_type -> v1.imageproxy.GetImageRequest
	1, // 2: v1.imageproxy.Imageproxy.HealthCheck:output_type -> v1.imageproxy.HealthCheckResponse
	3, // 3: v1.imageproxy.Imageproxy.GetImage:output_type -> v1.imageproxy.GetImageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_imageproxy_proto_init() }
func file_imageproxy_proto_init() {
	if File_imageproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imageproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_imageproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_imageproxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageRequest); i {
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
		file_imageproxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageResponse); i {
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
			RawDescriptor: file_imageproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imageproxy_proto_goTypes,
		DependencyIndexes: file_imageproxy_proto_depIdxs,
		MessageInfos:      file_imageproxy_proto_msgTypes,
	}.Build()
	File_imageproxy_proto = out.File
	file_imageproxy_proto_rawDesc = nil
	file_imageproxy_proto_goTypes = nil
	file_imageproxy_proto_depIdxs = nil
}
