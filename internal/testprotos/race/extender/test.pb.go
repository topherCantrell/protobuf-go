// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/race/extender/test.proto

package extender

import (
	message "google.golang.org/protobuf/internal/testprotos/race/message"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type OtherMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	I32           *int32                 `protobuf:"varint,1,opt,name=i32" json:"i32,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OtherMessage) Reset() {
	*x = OtherMessage{}
	mi := &file_internal_testprotos_race_extender_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OtherMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherMessage) ProtoMessage() {}

func (x *OtherMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_race_extender_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherMessage.ProtoReflect.Descriptor instead.
func (*OtherMessage) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_race_extender_test_proto_rawDescGZIP(), []int{0}
}

func (x *OtherMessage) GetI32() int32 {
	if x != nil && x.I32 != nil {
		return *x.I32
	}
	return 0
}

var file_internal_testprotos_race_extender_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*message.MyMessage)(nil),
		ExtensionType: (*string)(nil),
		Field:         2,
		Name:          "goproto.proto.test.s",
		Tag:           "bytes,2,opt,name=s",
		Filename:      "internal/testprotos/race/extender/test.proto",
	},
}

// Extension fields to message.MyMessage.
var (
	// optional string s = 2;
	E_S = &file_internal_testprotos_race_extender_test_proto_extTypes[0]
)

var File_internal_testprotos_race_extender_test_proto protoreflect.FileDescriptor

const file_internal_testprotos_race_extender_test_proto_rawDesc = "" +
	"\n" +
	",internal/testprotos/race/extender/test.proto\x12\x12goproto.proto.test\x1a+internal/testprotos/race/message/test.proto\" \n" +
	"\fOtherMessage\x12\x10\n" +
	"\x03i32\x18\x01 \x01(\x05R\x03i32:+\n" +
	"\x01s\x12\x1d.goproto.proto.test.MyMessage\x18\x02 \x01(\tR\x01sB>Z<google.golang.org/protobuf/internal/testprotos/race/extenderb\beditionsp\xe8\a"

var (
	file_internal_testprotos_race_extender_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_race_extender_test_proto_rawDescData []byte
)

func file_internal_testprotos_race_extender_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_race_extender_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_race_extender_test_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_testprotos_race_extender_test_proto_rawDesc), len(file_internal_testprotos_race_extender_test_proto_rawDesc)))
	})
	return file_internal_testprotos_race_extender_test_proto_rawDescData
}

var file_internal_testprotos_race_extender_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_race_extender_test_proto_goTypes = []any{
	(*OtherMessage)(nil),      // 0: goproto.proto.test.OtherMessage
	(*message.MyMessage)(nil), // 1: goproto.proto.test.MyMessage
}
var file_internal_testprotos_race_extender_test_proto_depIdxs = []int32{
	1, // 0: goproto.proto.test.s:extendee -> goproto.proto.test.MyMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_race_extender_test_proto_init() }
func file_internal_testprotos_race_extender_test_proto_init() {
	if File_internal_testprotos_race_extender_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_race_extender_test_proto_rawDesc), len(file_internal_testprotos_race_extender_test_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_race_extender_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_race_extender_test_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_race_extender_test_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_race_extender_test_proto_extTypes,
	}.Build()
	File_internal_testprotos_race_extender_test_proto = out.File
	file_internal_testprotos_race_extender_test_proto_goTypes = nil
	file_internal_testprotos_race_extender_test_proto_depIdxs = nil
}
