// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/user/user.proto

package user

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

type UserDetails struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDetails) Reset()         { *m = UserDetails{} }
func (m *UserDetails) String() string { return proto.CompactTextString(m) }
func (*UserDetails) ProtoMessage()    {}
func (*UserDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{0}
}

func (m *UserDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetails.Unmarshal(m, b)
}
func (m *UserDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetails.Marshal(b, m, deterministic)
}
func (m *UserDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetails.Merge(m, src)
}
func (m *UserDetails) XXX_Size() int {
	return xxx_messageInfo_UserDetails.Size(m)
}
func (m *UserDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetails.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetails proto.InternalMessageInfo

func (m *UserDetails) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserDetails) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UpdateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{2}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDetails)(nil), "authservice.user.UserDetails")
	proto.RegisterType((*CreateUserRequest)(nil), "authservice.user.CreateUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "authservice.user.UpdateUserRequest")
}

func init() { proto.RegisterFile("api/user/user.proto", fileDescriptor_3bc452c1a8bd0aa8) }

var fileDescriptor_3bc452c1a8bd0aa8 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2c, 0xc8, 0xd4,
	0x2f, 0x2d, 0x4e, 0x2d, 0x02, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x02, 0x89, 0xa5,
	0x25, 0x19, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x20, 0x71, 0x25, 0x77, 0x2e, 0xee,
	0xd0, 0xe2, 0xd4, 0x22, 0x97, 0xd4, 0x92, 0xc4, 0xcc, 0x9c, 0x62, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x84, 0x8b, 0x35,
	0x35, 0x37, 0x31, 0x33, 0x47, 0x82, 0x19, 0x2c, 0x08, 0xe1, 0x28, 0x45, 0x72, 0x09, 0x3a, 0x17,
	0xa5, 0x26, 0x96, 0xa4, 0x82, 0x8c, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x81, 0x6b, 0x67,
	0xc4, 0xa6, 0x9d, 0x09, 0x49, 0xbb, 0x90, 0x14, 0x17, 0x47, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e,
	0x51, 0x0a, 0xd4, 0x5c, 0x38, 0x5f, 0xc9, 0x99, 0x4b, 0x30, 0xb4, 0x20, 0x85, 0x08, 0xa3, 0xf1,
	0x18, 0xe2, 0xa4, 0x12, 0xa5, 0x84, 0xe4, 0x79, 0xfd, 0x82, 0xec, 0x74, 0x7d, 0x70, 0x98, 0x24,
	0x95, 0xa6, 0xa5, 0xe9, 0xc3, 0xc2, 0x2a, 0x89, 0x0d, 0x2c, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x1c, 0x9f, 0x8d, 0x3e, 0x01, 0x00, 0x00,
}
