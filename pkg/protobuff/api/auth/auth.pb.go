// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/auth/auth.proto

package auth

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

// Message for new Authentication Request
type AuthRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Message for retrieving Authentication Response
type AuthResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "authservice.auth.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "authservice.auth.AuthResponse")
}

func init() { proto.RegisterFile("api/auth/auth.proto", fileDescriptor_aa33bda803119f75) }

var fileDescriptor_aa33bda803119f75 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2c, 0x2d, 0xc9, 0x00, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x02, 0x20, 0x76,
	0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x1e, 0x88, 0xad, 0x64, 0xcf, 0xc5, 0xed, 0x58, 0x5a,
	0x92, 0x11, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc2, 0xc5, 0x9a, 0x9a, 0x9b, 0x98,
	0x99, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x49, 0x71, 0x71, 0x14, 0x24,
	0x16, 0x17, 0x97, 0xe7, 0x17, 0xa5, 0x48, 0x30, 0x81, 0x25, 0xe0, 0x7c, 0x25, 0x2f, 0x2e, 0x1e,
	0x88, 0x01, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x60, 0xed,
	0xac, 0x41, 0x4c, 0x99, 0x29, 0x20, 0x13, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0xa0, 0x1a, 0x21, 0x1c,
	0x84, 0x3d, 0xcc, 0x48, 0xf6, 0x38, 0xa9, 0x44, 0x29, 0x21, 0x39, 0x50, 0xbf, 0x20, 0x3b, 0x5d,
	0x1f, 0xec, 0xee, 0xa4, 0xd2, 0xb4, 0x34, 0x7d, 0x98, 0x7f, 0x92, 0xd8, 0xc0, 0x62, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xb9, 0x17, 0x79, 0xe2, 0x00, 0x00, 0x00,
}
