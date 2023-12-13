// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: idl/CodePro.proto

package rpc

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

type CodeProRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要编译运行的代码
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// 代码的语言
	Lang string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *CodeProRequest) Reset() {
	*x = CodeProRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_CodePro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeProRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeProRequest) ProtoMessage() {}

func (x *CodeProRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_CodePro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeProRequest.ProtoReflect.Descriptor instead.
func (*CodeProRequest) Descriptor() ([]byte, []int) {
	return file_idl_CodePro_proto_rawDescGZIP(), []int{0}
}

func (x *CodeProRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CodeProRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type CodeProResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 运行的输出结果及编译运行错误信息或报告内部错误信息
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// 运行状态
	// 0 - 编译运行成功  
	// 1 - 内部错误，由服务的bug导致，你可以在 https://github.com/scarletborder/pingyingqi/issues 中上报  
	// 2 - 代码引入了被禁止的模块(库/包)  
	// 3 - 编译运行成功但是运行时间过长
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CodeProResp) Reset() {
	*x = CodeProResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_CodePro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeProResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeProResp) ProtoMessage() {}

func (x *CodeProResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_CodePro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeProResp.ProtoReflect.Descriptor instead.
func (*CodeProResp) Descriptor() ([]byte, []int) {
	return file_idl_CodePro_proto_rawDescGZIP(), []int{1}
}

func (x *CodeProResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CodeProResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type DislikedPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 包的名称
	Pack string `protobuf:"bytes,1,opt,name=pack,proto3" json:"pack,omitempty"`
	// 你想设置的状态
	// True - 可用pass
	// False - 不可用blocked
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DislikedPackage) Reset() {
	*x = DislikedPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_CodePro_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DislikedPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DislikedPackage) ProtoMessage() {}

func (x *DislikedPackage) ProtoReflect() protoreflect.Message {
	mi := &file_idl_CodePro_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DislikedPackage.ProtoReflect.Descriptor instead.
func (*DislikedPackage) Descriptor() ([]byte, []int) {
	return file_idl_CodePro_proto_rawDescGZIP(), []int{2}
}

func (x *DislikedPackage) GetPack() string {
	if x != nil {
		return x.Pack
	}
	return ""
}

func (x *DislikedPackage) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type DislikedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *DislikedResp) Reset() {
	*x = DislikedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_CodePro_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DislikedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DislikedResp) ProtoMessage() {}

func (x *DislikedResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_CodePro_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DislikedResp.ProtoReflect.Descriptor instead.
func (*DislikedResp) Descriptor() ([]byte, []int) {
	return file_idl_CodePro_proto_rawDescGZIP(), []int{3}
}

func (x *DislikedResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *DislikedResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_idl_CodePro_proto protoreflect.FileDescriptor

var file_idl_CodePro_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x64, 0x6c, 0x2f, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x35, 0x0a,
	0x0b, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x64,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x36, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x68, 0x0a, 0x10, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x07, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x50, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x44, 0x69, 0x73,
	0x6c, 0x69, 0x6b, 0x65, 0x12, 0x10, 0x2e, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x79, 0x69, 0x6e,
	0x67, 0x71, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_CodePro_proto_rawDescOnce sync.Once
	file_idl_CodePro_proto_rawDescData = file_idl_CodePro_proto_rawDesc
)

func file_idl_CodePro_proto_rawDescGZIP() []byte {
	file_idl_CodePro_proto_rawDescOnce.Do(func() {
		file_idl_CodePro_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_CodePro_proto_rawDescData)
	})
	return file_idl_CodePro_proto_rawDescData
}

var file_idl_CodePro_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_CodePro_proto_goTypes = []interface{}{
	(*CodeProRequest)(nil),  // 0: CodeProRequest
	(*CodeProResp)(nil),     // 1: CodeProResp
	(*DislikedPackage)(nil), // 2: DislikedPackage
	(*DislikedResp)(nil),    // 3: DislikedResp
}
var file_idl_CodePro_proto_depIdxs = []int32{
	0, // 0: CodeProProgramer.CodePro:input_type -> CodeProRequest
	2, // 1: CodeProProgramer.Dislike:input_type -> DislikedPackage
	1, // 2: CodeProProgramer.CodePro:output_type -> CodeProResp
	3, // 3: CodeProProgramer.Dislike:output_type -> DislikedResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_CodePro_proto_init() }
func file_idl_CodePro_proto_init() {
	if File_idl_CodePro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_CodePro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeProRequest); i {
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
		file_idl_CodePro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeProResp); i {
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
		file_idl_CodePro_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DislikedPackage); i {
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
		file_idl_CodePro_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DislikedResp); i {
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
			RawDescriptor: file_idl_CodePro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_CodePro_proto_goTypes,
		DependencyIndexes: file_idl_CodePro_proto_depIdxs,
		MessageInfos:      file_idl_CodePro_proto_msgTypes,
	}.Build()
	File_idl_CodePro_proto = out.File
	file_idl_CodePro_proto_rawDesc = nil
	file_idl_CodePro_proto_goTypes = nil
	file_idl_CodePro_proto_depIdxs = nil
}
