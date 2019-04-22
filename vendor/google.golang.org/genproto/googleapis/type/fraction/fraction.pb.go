// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/fraction.proto

package fraction // import "google.golang.org/genproto/googleapis/type/fraction"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents a fraction in terms of a numerator divided by a denominator.
type Fraction struct {
	// The portion of the denominator in the faction, e.g. 2 in 2/3.
	Numerator int64 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	// The value by which the numerator is divided, e.g. 3 in 2/3. Must be
	// positive.
	Denominator          int64    `protobuf:"varint,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fraction) Reset()         { *m = Fraction{} }
func (m *Fraction) String() string { return proto.CompactTextString(m) }
func (*Fraction) ProtoMessage()    {}
func (*Fraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_fraction_b8dcff062bb9c31d, []int{0}
}
func (m *Fraction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fraction.Unmarshal(m, b)
}
func (m *Fraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fraction.Marshal(b, m, deterministic)
}
func (dst *Fraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fraction.Merge(dst, src)
}
func (m *Fraction) XXX_Size() int {
	return xxx_messageInfo_Fraction.Size(m)
}
func (m *Fraction) XXX_DiscardUnknown() {
	xxx_messageInfo_Fraction.DiscardUnknown(m)
}

var xxx_messageInfo_Fraction proto.InternalMessageInfo

func (m *Fraction) GetNumerator() int64 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *Fraction) GetDenominator() int64 {
	if m != nil {
		return m.Denominator
	}
	return 0
}

func init() {
	proto.RegisterType((*Fraction)(nil), "google.type.Fraction")
}

func init() {
	proto.RegisterFile("google/type/fraction.proto", fileDescriptor_fraction_b8dcff062bb9c31d)
}

var fileDescriptor_fraction_b8dcff062bb9c31d = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0x2b, 0x4a, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xc8, 0xe9, 0x81, 0xe4, 0x94, 0xbc, 0xb8,
	0x38, 0xdc, 0xa0, 0xd2, 0x42, 0x32, 0x5c, 0x9c, 0x79, 0xa5, 0xb9, 0xa9, 0x45, 0x89, 0x25, 0xf9,
	0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x08, 0x01, 0x21, 0x05, 0x2e, 0xee, 0x94, 0xd4,
	0xbc, 0xfc, 0xdc, 0xcc, 0x3c, 0xb0, 0x3c, 0x13, 0x58, 0x1e, 0x59, 0xc8, 0x29, 0x8d, 0x8b, 0x3f,
	0x39, 0x3f, 0x57, 0x0f, 0xc9, 0x78, 0x27, 0x5e, 0x98, 0xe1, 0x01, 0x20, 0xab, 0x03, 0x18, 0xa3,
	0x6c, 0xa0, 0xb2, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9,
	0x79, 0x60, 0x87, 0xe9, 0x43, 0xa4, 0x12, 0x0b, 0x32, 0x8b, 0x51, 0xdd, 0x6d, 0x0d, 0x63, 0x2c,
	0x62, 0x62, 0x76, 0x0f, 0x09, 0x48, 0x62, 0x03, 0x2b, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0xdd, 0xa8, 0x56, 0xe5, 0x00, 0x00, 0x00,
}
