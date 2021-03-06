// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MySmallMessage struct {
	Tiny                 string   `protobuf:"bytes,1,opt,name=tiny,proto3" json:"tiny,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySmallMessage) Reset()         { *m = MySmallMessage{} }
func (m *MySmallMessage) String() string { return proto.CompactTextString(m) }
func (*MySmallMessage) ProtoMessage()    {}
func (*MySmallMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *MySmallMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySmallMessage.Unmarshal(m, b)
}
func (m *MySmallMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySmallMessage.Marshal(b, m, deterministic)
}
func (m *MySmallMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySmallMessage.Merge(m, src)
}
func (m *MySmallMessage) XXX_Size() int {
	return xxx_messageInfo_MySmallMessage.Size(m)
}
func (m *MySmallMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MySmallMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MySmallMessage proto.InternalMessageInfo

func (m *MySmallMessage) GetTiny() string {
	if m != nil {
		return m.Tiny
	}
	return ""
}

func init() {
	proto.RegisterType((*MySmallMessage)(nil), "MySmallMessage")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 89 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe1, 0xe2, 0xf3, 0xad, 0x0c,
	0xce, 0x4d, 0xcc, 0xc9, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x29,
	0xc9, 0xcc, 0xab, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0x38, 0xa2, 0xd8,
	0xf4, 0xac, 0x73, 0x13, 0x33, 0xf3, 0x92, 0xd8, 0xc0, 0xda, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0x19, 0x53, 0x9e, 0x47, 0x00, 0x00, 0x00,
}
