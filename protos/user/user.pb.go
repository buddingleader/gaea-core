// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/gaeanetwork/gaea-core/protos/user/user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PublicKey            string   `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	SecretPrivKey        string   `protobuf:"bytes,5,opt,name=secret_priv_key,json=secretPrivKey,proto3" json:"secret_priv_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_881f7609b656d340, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *User) GetSecretPrivKey() string {
	if m != nil {
		return m.SecretPrivKey
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
}

func init() {
	proto.RegisterFile("github.com/gaeanetwork/gaea-core/protos/user/user.proto", fileDescriptor_user_881f7609b656d340)
}

var fileDescriptor_user_881f7609b656d340 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xc1, 0x6a, 0x83, 0x40,
	0x10, 0x86, 0xd1, 0xda, 0xa2, 0x03, 0x6d, 0x61, 0x4f, 0x4b, 0x4b, 0xa1, 0xf4, 0x50, 0x7a, 0xa9,
	0x1e, 0x72, 0xc8, 0x43, 0x08, 0x21, 0x04, 0x72, 0x96, 0x75, 0x1d, 0xcc, 0x62, 0x74, 0x97, 0xd9,
	0x55, 0xf1, 0x45, 0xf2, 0xbc, 0xc1, 0x91, 0x5c, 0x86, 0x7f, 0xbe, 0xff, 0x3b, 0xfc, 0xb0, 0x6f,
	0x4d, 0xb8, 0x8c, 0x75, 0xae, 0x6d, 0x5f, 0xb4, 0x0a, 0xd5, 0x80, 0x61, 0xb6, 0xd4, 0x71, 0xfe,
	0xd7, 0x96, 0xb0, 0x70, 0x64, 0x83, 0xf5, 0xc5, 0xe8, 0x91, 0xf8, 0xe4, 0x0c, 0x44, 0xb2, 0xe6,
	0x9f, 0x5b, 0x04, 0xc9, 0xd9, 0x23, 0x89, 0x37, 0x88, 0x4d, 0x23, 0xa3, 0xef, 0xe8, 0x2f, 0x3b,
	0xc5, 0xa6, 0x11, 0x9f, 0x90, 0xad, 0x42, 0x35, 0xa8, 0x1e, 0x65, 0xcc, 0x38, 0x5d, 0xc1, 0x41,
	0xf5, 0x28, 0x3e, 0x20, 0x75, 0xca, 0xfb, 0xd9, 0x52, 0x23, 0x9f, 0xb6, 0xee, 0xf1, 0x8b, 0x2f,
	0x00, 0x37, 0xd6, 0x57, 0xa3, 0xab, 0x0e, 0x17, 0x99, 0x70, 0x9b, 0x6d, 0xa4, 0xc4, 0x45, 0xfc,
	0xc2, 0xbb, 0x47, 0x4d, 0x18, 0x2a, 0x47, 0x66, 0x62, 0xe7, 0x99, 0x9d, 0xd7, 0x0d, 0x1f, 0xc9,
	0x4c, 0x25, 0x2e, 0xf5, 0x0b, 0xaf, 0xdc, 0xdd, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xc9, 0x28,
	0x8a, 0xe0, 0x00, 0x00, 0x00,
}
