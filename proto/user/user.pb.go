// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

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
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	CreatedTime          uint64   `protobuf:"varint,4,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          uint64   `protobuf:"varint,5,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_b65d69083b1bbfbd, []int{0}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *User) GetCreatedTime() uint64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *User) GetUpdatedTime() uint64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_b65d69083b1bbfbd, []int{1}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPwd              string   `protobuf:"bytes,3,opt,name=userPwd,proto3" json:"userPwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_b65d69083b1bbfbd, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_b65d69083b1bbfbd, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.user")
	proto.RegisterType((*Error)(nil), "go.micro.srv.user.Error")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_user_b65d69083b1bbfbd) }

var fileDescriptor_user_b65d69083b1bbfbd = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xcd, 0xd6, 0xee, 0xcf, 0x1d, 0xf8, 0x27, 0xa0, 0x86, 0xf9, 0x52, 0xfa, 0x34, 0x10,
	0x22, 0x6c, 0xdf, 0x40, 0xf4, 0x41, 0x10, 0xd1, 0x30, 0xf1, 0x79, 0x36, 0x17, 0x29, 0xb8, 0xa5,
	0x26, 0xad, 0xb2, 0x37, 0xfd, 0xe6, 0x72, 0x6f, 0xbb, 0x3a, 0xb0, 0x2f, 0xe1, 0x9c, 0x9b, 0x5f,
	0x92, 0x7b, 0x6e, 0x0b, 0xa7, 0x85, 0x77, 0xa5, 0xbb, 0xaa, 0x02, 0x7a, 0x5e, 0x34, 0x7b, 0x79,
	0xf2, 0xe6, 0xf4, 0x3a, 0xcf, 0xbc, 0xd3, 0xc1, 0x7f, 0x6a, 0xda, 0x48, 0xbf, 0x05, 0x44, 0x24,
	0xe4, 0x21, 0xf4, 0x72, 0xab, 0x44, 0x22, 0x66, 0x7d, 0xd3, 0xcb, 0xad, 0x94, 0x10, 0x6d, 0x56,
	0x6b, 0x54, 0xbd, 0x44, 0xcc, 0xc6, 0x86, 0xb5, 0x3c, 0x86, 0x7e, 0xf1, 0x65, 0x55, 0x9f, 0x4b,
	0x24, 0x65, 0x02, 0x93, 0xcc, 0xe3, 0xaa, 0x44, 0xbb, 0xcc, 0xd7, 0xa8, 0xa2, 0x44, 0xcc, 0x22,
	0xb3, 0x5f, 0x22, 0xa2, 0x2a, 0x6c, 0x4b, 0xc4, 0x35, 0xb1, 0x57, 0x4a, 0x17, 0x10, 0xdf, 0x7a,
	0xef, 0x3c, 0x3d, 0x99, 0x39, 0x8b, 0xdc, 0x44, 0x6c, 0x58, 0xcb, 0x33, 0x18, 0x58, 0x2c, 0x57,
	0xf9, 0x7b, 0xd3, 0x48, 0xe3, 0xd2, 0x17, 0x18, 0x1a, 0xfc, 0xa8, 0x30, 0x94, 0x84, 0x50, 0x82,
	0xbb, 0x1b, 0x3e, 0x38, 0x36, 0x8d, 0x93, 0x53, 0x18, 0x91, 0x7a, 0xf8, 0x4b, 0xd1, 0x7a, 0xa9,
	0x60, 0x48, 0xfa, 0xb1, 0x4d, 0xb3, 0xb3, 0xe9, 0x8f, 0x80, 0x91, 0xc1, 0x50, 0xb8, 0x4d, 0x60,
	0x2c, 0x54, 0x59, 0x86, 0x21, 0xf0, 0xdd, 0x23, 0xb3, 0xb3, 0x52, 0x43, 0x8c, 0xd4, 0x34, 0xdf,
	0x3c, 0x99, 0x2b, 0xfd, 0x6f, 0xb4, 0x9a, 0x43, 0x99, 0x1a, 0x93, 0x97, 0xf5, 0x98, 0xf9, 0xb5,
	0xc9, 0xfc, 0xbc, 0x03, 0xa7, 0xc5, 0x30, 0x34, 0x5f, 0x42, 0xf4, 0x4c, 0xdf, 0xe4, 0x1e, 0x8e,
	0x9e, 0x2a, 0xf4, 0x5b, 0x32, 0xd7, 0x5b, 0x6e, 0x7c, 0xda, 0x71, 0xb2, 0x19, 0xc4, 0xf4, 0xa2,
	0x73, 0xaf, 0x8e, 0x92, 0x1e, 0xbc, 0x0e, 0xf8, 0x27, 0x58, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xf7, 0x73, 0xcf, 0x57, 0x1d, 0x02, 0x00, 0x00,
}
