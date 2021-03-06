// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time.proto

package go_micro_lzx_time

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

// 用户信息
type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CrateTime            int32    `protobuf:"varint,3,opt,name=crate_time,json=crateTime,proto3" json:"crate_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_time_aa2bc4a1b5a7e9ab, []int{0}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCrateTime() int32 {
	if m != nil {
		return m.CrateTime
	}
	return 0
}

// 等待运送的货物
type Username struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Username) Reset()         { *m = Username{} }
func (m *Username) String() string { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()    {}
func (*Username) Descriptor() ([]byte, []int) {
	return fileDescriptor_time_aa2bc4a1b5a7e9ab, []int{1}
}
func (m *Username) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Username.Unmarshal(m, b)
}
func (m *Username) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Username.Marshal(b, m, deterministic)
}
func (dst *Username) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Username.Merge(dst, src)
}
func (m *Username) XXX_Size() int {
	return xxx_messageInfo_Username.Size(m)
}
func (m *Username) XXX_DiscardUnknown() {
	xxx_messageInfo_Username.DiscardUnknown(m)
}

var xxx_messageInfo_Username proto.InternalMessageInfo

func (m *Username) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 返回的用户信息
type GetUserNameResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserNameResponse) Reset()         { *m = GetUserNameResponse{} }
func (m *GetUserNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserNameResponse) ProtoMessage()    {}
func (*GetUserNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_time_aa2bc4a1b5a7e9ab, []int{2}
}
func (m *GetUserNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserNameResponse.Unmarshal(m, b)
}
func (m *GetUserNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserNameResponse.Marshal(b, m, deterministic)
}
func (dst *GetUserNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserNameResponse.Merge(dst, src)
}
func (m *GetUserNameResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserNameResponse.Size(m)
}
func (m *GetUserNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserNameResponse proto.InternalMessageInfo

func (m *GetUserNameResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// 返回的用户信息
type AddUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserResponse) Reset()         { *m = AddUserResponse{} }
func (m *AddUserResponse) String() string { return proto.CompactTextString(m) }
func (*AddUserResponse) ProtoMessage()    {}
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_time_aa2bc4a1b5a7e9ab, []int{3}
}
func (m *AddUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserResponse.Unmarshal(m, b)
}
func (m *AddUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserResponse.Marshal(b, m, deterministic)
}
func (dst *AddUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserResponse.Merge(dst, src)
}
func (m *AddUserResponse) XXX_Size() int {
	return xxx_messageInfo_AddUserResponse.Size(m)
}
func (m *AddUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserResponse proto.InternalMessageInfo

func (m *AddUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.lzx.time.User")
	proto.RegisterType((*Username)(nil), "go.micro.lzx.time.Username")
	proto.RegisterType((*GetUserNameResponse)(nil), "go.micro.lzx.time.GetUserNameResponse")
	proto.RegisterType((*AddUserResponse)(nil), "go.micro.lzx.time.AddUserResponse")
}

func init() { proto.RegisterFile("time.proto", fileDescriptor_time_aa2bc4a1b5a7e9ab) }

var fileDescriptor_time_aa2bc4a1b5a7e9ab = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c, 0x2e, 0xca,
	0xd7, 0xcb, 0xa9, 0xaa, 0xd0, 0x03, 0x49, 0x28, 0x79, 0x72, 0xb1, 0x84, 0x16, 0xa7, 0x16, 0x09,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x45, 0xc0, 0x6c, 0x21, 0x59, 0x2e,
	0xae, 0xe4, 0xa2, 0xc4, 0x92, 0xd4, 0x78, 0x90, 0x4e, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20,
	0x4e, 0xb0, 0x48, 0x08, 0xc8, 0x28, 0x39, 0x2e, 0x0e, 0x90, 0x51, 0x60, 0xa5, 0x30, 0xed, 0x8c,
	0x08, 0xed, 0x4a, 0x4e, 0x5c, 0xc2, 0xee, 0xa9, 0x25, 0x20, 0x25, 0x7e, 0x89, 0xb9, 0xa9, 0x41,
	0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xda, 0x5c, 0x2c, 0xa5, 0xc5, 0xa9, 0x45, 0x60,
	0xa5, 0xdc, 0x46, 0xe2, 0x7a, 0x18, 0x6e, 0xd4, 0x03, 0x69, 0x09, 0x02, 0x2b, 0x52, 0xb2, 0xe3,
	0xe2, 0x77, 0x4c, 0x49, 0x01, 0x0b, 0x90, 0xa3, 0xdf, 0x68, 0x23, 0x23, 0x17, 0x37, 0xc8, 0xb1,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x21, 0x5c, 0xdc, 0x50, 0x37, 0x81, 0x44, 0x85,
	0xa4, 0x71, 0xe8, 0x06, 0xb9, 0x5f, 0x4a, 0x0d, 0x8b, 0x24, 0x16, 0x0f, 0x29, 0x31, 0x08, 0xf9,
	0x70, 0xb1, 0x43, 0x5d, 0x89, 0xdf, 0x44, 0x25, 0x2c, 0x92, 0x68, 0xde, 0x53, 0x62, 0x48, 0x62,
	0x03, 0x47, 0x9e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x87, 0xe1, 0xd8, 0xee, 0xca, 0x01, 0x00,
	0x00,
}
