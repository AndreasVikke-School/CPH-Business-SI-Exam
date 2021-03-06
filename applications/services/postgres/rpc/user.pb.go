// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Age                  int64    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserList struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "rpc.User")
	proto.RegisterType((*UserList)(nil), "rpc.UserList")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x52, 0x6b, 0x3b, 0xc1, 0x3f, 0xec, 0x41, 0x42, 0x0a, 0x1a, 0x72, 0x0a, 0x08,
	0x1b, 0xac, 0xc5, 0xbb, 0x8a, 0x14, 0xc1, 0x83, 0x44, 0xf4, 0xe0, 0xc9, 0x6d, 0x32, 0x86, 0x40,
	0xda, 0x2c, 0xb3, 0x1b, 0x8b, 0x5f, 0xc5, 0x8f, 0xe0, 0xa7, 0x94, 0xdd, 0xa5, 0xad, 0xe8, 0x6d,
	0xdf, 0xbe, 0x37, 0xbf, 0x99, 0x07, 0xd0, 0x2b, 0x24, 0x2e, 0xa9, 0xd3, 0x1d, 0x0b, 0x48, 0x96,
	0xf1, 0x69, 0xdd, 0x75, 0x75, 0x8b, 0xb9, 0xfd, 0x5a, 0xf4, 0xef, 0xf9, 0x9a, 0x84, 0x94, 0x48,
	0xca, 0x85, 0xe2, 0xc9, 0x5f, 0x1f, 0x97, 0x52, 0x7f, 0x3a, 0x33, 0x7d, 0x83, 0xc1, 0xb3, 0x42,
	0x62, 0x87, 0xe0, 0x37, 0x55, 0xe4, 0x25, 0x5e, 0x16, 0x14, 0x7e, 0x53, 0xb1, 0x18, 0x46, 0x66,
	0xcf, 0x4a, 0x2c, 0x31, 0xf2, 0x13, 0x2f, 0x1b, 0x17, 0x5b, 0xcd, 0x8e, 0x21, 0x10, 0x35, 0x46,
	0x81, 0x0d, 0x9b, 0xa7, 0x49, 0x4b, 0xa1, 0xd4, 0xba, 0xa3, 0x2a, 0x1a, 0xb8, 0xf4, 0x46, 0xa7,
	0xe7, 0x30, 0x32, 0x1b, 0x1e, 0x1a, 0xa5, 0xd9, 0x19, 0xec, 0x19, 0x8a, 0x8a, 0xbc, 0x24, 0xc8,
	0xc2, 0xe9, 0x98, 0x93, 0x2c, 0xb9, 0x71, 0x0b, 0xf7, 0x3f, 0xfd, 0xf2, 0x20, 0x34, 0xfa, 0x09,
	0xe9, 0xa3, 0x29, 0x91, 0x5d, 0xc0, 0xfe, 0x1c, 0xb5, 0xbd, 0x70, 0xc2, 0x5d, 0x0f, 0xbe, 0xe9,
	0xc1, 0xef, 0x57, 0xfa, 0x6a, 0xf6, 0x22, 0xda, 0x1e, 0xe3, 0x1d, 0x89, 0xcd, 0x20, 0x9c, 0xa3,
	0xbe, 0x6e, 0x5b, 0xa3, 0x14, 0x3b, 0xf9, 0x37, 0x76, 0x67, 0xea, 0xc7, 0x07, 0xdb, 0x09, 0x7b,
	0x59, 0x0a, 0x70, 0x4b, 0x28, 0x34, 0x5a, 0xc6, 0x0e, 0xf7, 0x8b, 0x7c, 0x73, 0xf4, 0x3a, 0xc8,
	0x49, 0x96, 0xdf, 0xfe, 0xf0, 0xd1, 0xb0, 0xd4, 0x62, 0x68, 0x99, 0x97, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xcf, 0x5b, 0xe4, 0x8b, 0x93, 0x01, 0x00, 0x00,
}
