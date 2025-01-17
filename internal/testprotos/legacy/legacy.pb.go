// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/legacy/legacy.proto

package legacy

import (
	proto2_20160225_2fc053c5 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20160225_2fc053c5"
	proto2_20160519_a4ab9ec5 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20160519_a4ab9ec5"
	proto2_20180125_92554152 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20180125_92554152"
	proto2_20180430_b4deda09 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20180430_b4deda09"
	proto2_20180814_aa810b61 "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20180814_aa810b61"
	proto2_20190205_c823c79e "google.golang.org/protobuf/internal/testprotos/legacy/proto2_20190205_c823c79e"
	proto3_20160225_2fc053c5 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20160225_2fc053c5"
	proto3_20160519_a4ab9ec5 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20160519_a4ab9ec5"
	proto3_20180125_92554152 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20180125_92554152"
	proto3_20180430_b4deda09 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20180430_b4deda09"
	proto3_20180814_aa810b61 "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20180814_aa810b61"
	proto3_20190205_c823c79e "google.golang.org/protobuf/internal/testprotos/legacy/proto3_20190205_c823c79e"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type Legacy struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	F1            *proto2_20160225_2fc053c5.Message `protobuf:"bytes,1,opt,name=f1,proto3" json:"f1,omitempty"`
	F2            *proto3_20160225_2fc053c5.Message `protobuf:"bytes,2,opt,name=f2,proto3" json:"f2,omitempty"`
	F3            *proto2_20160519_a4ab9ec5.Message `protobuf:"bytes,3,opt,name=f3,proto3" json:"f3,omitempty"`
	F4            *proto3_20160519_a4ab9ec5.Message `protobuf:"bytes,4,opt,name=f4,proto3" json:"f4,omitempty"`
	F5            *proto2_20180125_92554152.Message `protobuf:"bytes,5,opt,name=f5,proto3" json:"f5,omitempty"`
	F6            *proto3_20180125_92554152.Message `protobuf:"bytes,6,opt,name=f6,proto3" json:"f6,omitempty"`
	F7            *proto2_20180430_b4deda09.Message `protobuf:"bytes,7,opt,name=f7,proto3" json:"f7,omitempty"`
	F8            *proto3_20180430_b4deda09.Message `protobuf:"bytes,8,opt,name=f8,proto3" json:"f8,omitempty"`
	F9            *proto2_20180814_aa810b61.Message `protobuf:"bytes,9,opt,name=f9,proto3" json:"f9,omitempty"`
	F10           *proto3_20180814_aa810b61.Message `protobuf:"bytes,10,opt,name=f10,proto3" json:"f10,omitempty"`
	F11           *proto2_20190205_c823c79e.Message `protobuf:"bytes,11,opt,name=f11,proto3" json:"f11,omitempty"`
	F12           *proto3_20190205_c823c79e.Message `protobuf:"bytes,12,opt,name=f12,proto3" json:"f12,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Legacy) Reset() {
	*x = Legacy{}
	mi := &file_internal_testprotos_legacy_legacy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Legacy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Legacy) ProtoMessage() {}

func (x *Legacy) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_legacy_legacy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Legacy.ProtoReflect.Descriptor instead.
func (*Legacy) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_legacy_legacy_proto_rawDescGZIP(), []int{0}
}

func (x *Legacy) GetF1() *proto2_20160225_2fc053c5.Message {
	if x != nil {
		return x.F1
	}
	return nil
}

func (x *Legacy) GetF2() *proto3_20160225_2fc053c5.Message {
	if x != nil {
		return x.F2
	}
	return nil
}

func (x *Legacy) GetF3() *proto2_20160519_a4ab9ec5.Message {
	if x != nil {
		return x.F3
	}
	return nil
}

func (x *Legacy) GetF4() *proto3_20160519_a4ab9ec5.Message {
	if x != nil {
		return x.F4
	}
	return nil
}

func (x *Legacy) GetF5() *proto2_20180125_92554152.Message {
	if x != nil {
		return x.F5
	}
	return nil
}

func (x *Legacy) GetF6() *proto3_20180125_92554152.Message {
	if x != nil {
		return x.F6
	}
	return nil
}

func (x *Legacy) GetF7() *proto2_20180430_b4deda09.Message {
	if x != nil {
		return x.F7
	}
	return nil
}

func (x *Legacy) GetF8() *proto3_20180430_b4deda09.Message {
	if x != nil {
		return x.F8
	}
	return nil
}

func (x *Legacy) GetF9() *proto2_20180814_aa810b61.Message {
	if x != nil {
		return x.F9
	}
	return nil
}

func (x *Legacy) GetF10() *proto3_20180814_aa810b61.Message {
	if x != nil {
		return x.F10
	}
	return nil
}

func (x *Legacy) GetF11() *proto2_20190205_c823c79e.Message {
	if x != nil {
		return x.F11
	}
	return nil
}

func (x *Legacy) GetF12() *proto3_20190205_c823c79e.Message {
	if x != nil {
		return x.F12
	}
	return nil
}

var File_internal_testprotos_legacy_legacy_proto protoreflect.FileDescriptor

const file_internal_testprotos_legacy_legacy_proto_rawDesc = "" +
	"\n" +
	"'internal/testprotos/legacy/legacy.proto\x12\x11google.golang.org\x1a>internal/testprotos/legacy/proto2_20160225_2fc053c5/test.proto\x1a>internal/testprotos/legacy/proto3_20160225_2fc053c5/test.proto\x1a>internal/testprotos/legacy/proto2_20160519_a4ab9ec5/test.proto\x1a>internal/testprotos/legacy/proto3_20160519_a4ab9ec5/test.proto\x1a>internal/testprotos/legacy/proto2_20180125_92554152/test.proto\x1a>internal/testprotos/legacy/proto3_20180125_92554152/test.proto\x1a>internal/testprotos/legacy/proto2_20180430_b4deda09/test.proto\x1a>internal/testprotos/legacy/proto3_20180430_b4deda09/test.proto\x1a>internal/testprotos/legacy/proto2_20180814_aa810b61/test.proto\x1a>internal/testprotos/legacy/proto3_20180814_aa810b61/test.proto\x1a>internal/testprotos/legacy/proto2_20190205_c823c79e/test.proto\x1a>internal/testprotos/legacy/proto3_20190205_c823c79e/test.proto\"\xde\x05\n" +
	"\x06Legacy\x12:\n" +
	"\x02f1\x18\x01 \x01(\v2*.google.golang.org.proto2_20160225.MessageR\x02f1\x12:\n" +
	"\x02f2\x18\x02 \x01(\v2*.google.golang.org.proto3_20160225.MessageR\x02f2\x12:\n" +
	"\x02f3\x18\x03 \x01(\v2*.google.golang.org.proto2_20160519.MessageR\x02f3\x12:\n" +
	"\x02f4\x18\x04 \x01(\v2*.google.golang.org.proto3_20160519.MessageR\x02f4\x12:\n" +
	"\x02f5\x18\x05 \x01(\v2*.google.golang.org.proto2_20180125.MessageR\x02f5\x12:\n" +
	"\x02f6\x18\x06 \x01(\v2*.google.golang.org.proto3_20180125.MessageR\x02f6\x12:\n" +
	"\x02f7\x18\a \x01(\v2*.google.golang.org.proto2_20180430.MessageR\x02f7\x12:\n" +
	"\x02f8\x18\b \x01(\v2*.google.golang.org.proto3_20180430.MessageR\x02f8\x12:\n" +
	"\x02f9\x18\t \x01(\v2*.google.golang.org.proto2_20180814.MessageR\x02f9\x12<\n" +
	"\x03f10\x18\n" +
	" \x01(\v2*.google.golang.org.proto3_20180814.MessageR\x03f10\x12<\n" +
	"\x03f11\x18\v \x01(\v2*.google.golang.org.proto2_20190205.MessageR\x03f11\x12<\n" +
	"\x03f12\x18\f \x01(\v2*.google.golang.org.proto3_20190205.MessageR\x03f12B7Z5google.golang.org/protobuf/internal/testprotos/legacyb\x06proto3"

var (
	file_internal_testprotos_legacy_legacy_proto_rawDescOnce sync.Once
	file_internal_testprotos_legacy_legacy_proto_rawDescData []byte
)

func file_internal_testprotos_legacy_legacy_proto_rawDescGZIP() []byte {
	file_internal_testprotos_legacy_legacy_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_legacy_legacy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_testprotos_legacy_legacy_proto_rawDesc), len(file_internal_testprotos_legacy_legacy_proto_rawDesc)))
	})
	return file_internal_testprotos_legacy_legacy_proto_rawDescData
}

var file_internal_testprotos_legacy_legacy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_legacy_legacy_proto_goTypes = []any{
	(*Legacy)(nil),                           // 0: google.golang.org.Legacy
	(*proto2_20160225_2fc053c5.Message)(nil), // 1: google.golang.org.proto2_20160225.Message
	(*proto3_20160225_2fc053c5.Message)(nil), // 2: google.golang.org.proto3_20160225.Message
	(*proto2_20160519_a4ab9ec5.Message)(nil), // 3: google.golang.org.proto2_20160519.Message
	(*proto3_20160519_a4ab9ec5.Message)(nil), // 4: google.golang.org.proto3_20160519.Message
	(*proto2_20180125_92554152.Message)(nil), // 5: google.golang.org.proto2_20180125.Message
	(*proto3_20180125_92554152.Message)(nil), // 6: google.golang.org.proto3_20180125.Message
	(*proto2_20180430_b4deda09.Message)(nil), // 7: google.golang.org.proto2_20180430.Message
	(*proto3_20180430_b4deda09.Message)(nil), // 8: google.golang.org.proto3_20180430.Message
	(*proto2_20180814_aa810b61.Message)(nil), // 9: google.golang.org.proto2_20180814.Message
	(*proto3_20180814_aa810b61.Message)(nil), // 10: google.golang.org.proto3_20180814.Message
	(*proto2_20190205_c823c79e.Message)(nil), // 11: google.golang.org.proto2_20190205.Message
	(*proto3_20190205_c823c79e.Message)(nil), // 12: google.golang.org.proto3_20190205.Message
}
var file_internal_testprotos_legacy_legacy_proto_depIdxs = []int32{
	1,  // 0: google.golang.org.Legacy.f1:type_name -> google.golang.org.proto2_20160225.Message
	2,  // 1: google.golang.org.Legacy.f2:type_name -> google.golang.org.proto3_20160225.Message
	3,  // 2: google.golang.org.Legacy.f3:type_name -> google.golang.org.proto2_20160519.Message
	4,  // 3: google.golang.org.Legacy.f4:type_name -> google.golang.org.proto3_20160519.Message
	5,  // 4: google.golang.org.Legacy.f5:type_name -> google.golang.org.proto2_20180125.Message
	6,  // 5: google.golang.org.Legacy.f6:type_name -> google.golang.org.proto3_20180125.Message
	7,  // 6: google.golang.org.Legacy.f7:type_name -> google.golang.org.proto2_20180430.Message
	8,  // 7: google.golang.org.Legacy.f8:type_name -> google.golang.org.proto3_20180430.Message
	9,  // 8: google.golang.org.Legacy.f9:type_name -> google.golang.org.proto2_20180814.Message
	10, // 9: google.golang.org.Legacy.f10:type_name -> google.golang.org.proto3_20180814.Message
	11, // 10: google.golang.org.Legacy.f11:type_name -> google.golang.org.proto2_20190205.Message
	12, // 11: google.golang.org.Legacy.f12:type_name -> google.golang.org.proto3_20190205.Message
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_internal_testprotos_legacy_legacy_proto_init() }
func file_internal_testprotos_legacy_legacy_proto_init() {
	if File_internal_testprotos_legacy_legacy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_legacy_legacy_proto_rawDesc), len(file_internal_testprotos_legacy_legacy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_legacy_legacy_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_legacy_legacy_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_legacy_legacy_proto_msgTypes,
	}.Build()
	File_internal_testprotos_legacy_legacy_proto = out.File
	file_internal_testprotos_legacy_legacy_proto_goTypes = nil
	file_internal_testprotos_legacy_legacy_proto_depIdxs = nil
}
