// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zap.proto

package zappb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_Redact = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         847847,
	Name:          "zap.redact",
	Tag:           "varint,847847,opt,name=redact",
	Filename:      "zap.proto",
}

func init() {
	proto.RegisterExtension(E_Redact)
}

func init() { proto.RegisterFile("zap.proto", fileDescriptor_zap_6aeaad0210499773) }

var fileDescriptor_zap_6aeaad0210499773 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xac, 0x4a, 0x2c, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xae, 0x4a, 0x2c, 0x90, 0x52, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16,
	0x94, 0xe4, 0x17, 0x41, 0x94, 0x59, 0x99, 0x73, 0xb1, 0x15, 0xa5, 0xa6, 0x24, 0x26, 0x97, 0x08,
	0xc9, 0xea, 0x41, 0x14, 0xeb, 0xc1, 0x14, 0xeb, 0xb9, 0x65, 0xa6, 0xe6, 0xa4, 0xf8, 0x17, 0x94,
	0x64, 0xe6, 0xe7, 0x15, 0x4b, 0x3c, 0xbf, 0x6f, 0xac, 0xc0, 0xa8, 0xc1, 0x11, 0x04, 0x55, 0xee,
	0xc4, 0x1e, 0xc5, 0x5a, 0x95, 0x58, 0x50, 0x90, 0x94, 0xc4, 0x06, 0x56, 0x6f, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x82, 0x93, 0x5c, 0xa6, 0x7c, 0x00, 0x00, 0x00,
}
