// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/user/proto/user.proto

package userpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/fzerorubigd/protobuf/extra"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UserStatus int32

const (
	UserStatus_USER_STATUS_INVALID    UserStatus = 0
	UserStatus_USER_STATUS_REGISTERED UserStatus = 1
	UserStatus_USER_STATUS_ACTIVE     UserStatus = 2
	UserStatus_USER_STATUS_BANNED     UserStatus = 3
)

var UserStatus_name = map[int32]string{
	0: "USER_STATUS_INVALID",
	1: "USER_STATUS_REGISTERED",
	2: "USER_STATUS_ACTIVE",
	3: "USER_STATUS_BANNED",
}
var UserStatus_value = map[string]int32{
	"USER_STATUS_INVALID":    0,
	"USER_STATUS_REGISTERED": 1,
	"USER_STATUS_ACTIVE":     2,
	"USER_STATUS_BANNED":     3,
}

func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{0}
}

type User struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" `
	Email                string           `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" db:"email" `
	Password             string           `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" db:"password" `
	Status               UserStatus       `protobuf:"varint,4,opt,name=status,proto3,enum=user.UserStatus" json:"status,omitempty" db:"status" `
	CreatedAt            *types.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at" `
	UpdatedAt            *types.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at" `
	LastLogin            *types.Timestamp `protobuf:"bytes,7,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty" db:"last_login" `
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{0}
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

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_USER_STATUS_INVALID
}

func (m *User) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetLastLogin() *types.Timestamp {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

type UserResponse struct {
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string     `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Token                string     `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Status               UserStatus `protobuf:"varint,4,opt,name=status,proto3,enum=user.UserStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserResponse) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_USER_STATUS_INVALID
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" validate:"email,required" `
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"gte=6,required" `
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{2}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogoutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{3}
}
func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (dst *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(dst, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{4}
}
func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (dst *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(dst, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

type RegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" validate:"email,required" `
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"gte=6,required" `
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{5}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{6}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type ChangePasswordRequest struct {
	OldPassword          string   `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty" validate:"gte=6,required" `
	NewPassword          string   `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty" validate:"gte=6,required" `
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordRequest) Reset()         { *m = ChangePasswordRequest{} }
func (m *ChangePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordRequest) ProtoMessage()    {}
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{7}
}
func (m *ChangePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordRequest.Unmarshal(m, b)
}
func (m *ChangePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordRequest.Marshal(b, m, deterministic)
}
func (dst *ChangePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordRequest.Merge(dst, src)
}
func (m *ChangePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordRequest.Size(m)
}
func (m *ChangePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordRequest proto.InternalMessageInfo

func (m *ChangePasswordRequest) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *ChangePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ChangePasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordResponse) Reset()         { *m = ChangePasswordResponse{} }
func (m *ChangePasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordResponse) ProtoMessage()    {}
func (*ChangePasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d177c44034e7d891, []int{8}
}
func (m *ChangePasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordResponse.Unmarshal(m, b)
}
func (m *ChangePasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordResponse.Marshal(b, m, deterministic)
}
func (dst *ChangePasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordResponse.Merge(dst, src)
}
func (m *ChangePasswordResponse) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordResponse.Size(m)
}
func (m *ChangePasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LogoutRequest)(nil), "user.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "user.LogoutResponse")
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*PingRequest)(nil), "user.PingRequest")
	proto.RegisterType((*ChangePasswordRequest)(nil), "user.ChangePasswordRequest")
	proto.RegisterType((*ChangePasswordResponse)(nil), "user.ChangePasswordResponse")
	proto.RegisterEnum("user.UserStatus", UserStatus_name, UserStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserSystemClient is the client API for UserSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSystemClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
}

type userSystemClient struct {
	cc *grpc.ClientConn
}

func NewUserSystemClient(cc *grpc.ClientConn) UserSystemClient {
	return &userSystemClient{cc}
}

func (c *userSystemClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSystemServer is the server API for UserSystem service.
type UserSystemServer interface {
	Login(context.Context, *LoginRequest) (*UserResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Register(context.Context, *RegisterRequest) (*UserResponse, error)
	Ping(context.Context, *PingRequest) (*UserResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
}

func RegisterUserSystemServer(s *grpc.Server, srv UserSystemServer) {
	s.RegisterService(&_UserSystem_serviceDesc, srv)
}

func _UserSystem_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSystem_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSystem_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSystem_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSystem_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserSystem",
	HandlerType: (*UserSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserSystem_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserSystem_Logout_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserSystem_Register_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _UserSystem_Ping_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserSystem_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/user/proto/user.proto",
}

func init() { proto.RegisterFile("modules/user/proto/user.proto", fileDescriptor_user_d177c44034e7d891) }

var fileDescriptor_user_d177c44034e7d891 = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x3d, 0x6f, 0x23, 0x45,
	0x18, 0xce, 0xfa, 0x8b, 0x64, 0xec, 0xd8, 0xce, 0xc4, 0x67, 0x9c, 0x25, 0x87, 0x57, 0x2b, 0x81,
	0xac, 0xe8, 0xe2, 0xd5, 0x85, 0x13, 0x85, 0x25, 0x24, 0xec, 0x8b, 0x85, 0x82, 0xa2, 0x10, 0xad,
	0x9d, 0x2b, 0x68, 0xac, 0xb1, 0x77, 0x6e, 0x33, 0xc2, 0xde, 0xd9, 0xec, 0xcc, 0x9e, 0x81, 0x0a,
	0x41, 0x77, 0x88, 0x06, 0x6a, 0x7e, 0x03, 0x12, 0x1d, 0xff, 0x80, 0x9e, 0x8e, 0xc2, 0x15, 0x3f,
	0xe0, 0xe4, 0x5f, 0x80, 0x66, 0x66, 0xbf, 0x12, 0x2c, 0x25, 0x54, 0xd7, 0xd8, 0x33, 0xef, 0xc7,
	0xf3, 0x3e, 0x7e, 0xde, 0xc7, 0x03, 0x1e, 0x2f, 0xa8, 0x13, 0xce, 0x31, 0xb3, 0x42, 0x86, 0x03,
	0xcb, 0x0f, 0x28, 0xa7, 0xf2, 0xd8, 0x95, 0x47, 0x58, 0x10, 0x67, 0xfd, 0xa9, 0x4b, 0xf8, 0x75,
	0x38, 0xed, 0xce, 0xe8, 0xc2, 0x7a, 0xf9, 0x2d, 0x0e, 0x68, 0x10, 0x4e, 0x89, 0xeb, 0xa8, 0xf2,
	0x69, 0xf8, 0xd2, 0xc2, 0x5f, 0xf3, 0x00, 0xa9, 0x4f, 0xd5, 0xa8, 0x1f, 0x67, 0x5a, 0x5c, 0xea,
	0xd2, 0xb4, 0x56, 0xdc, 0xd4, 0x1c, 0x71, 0x8a, 0xca, 0x0f, 0x5d, 0x4a, 0xdd, 0x39, 0xb6, 0x90,
	0x4f, 0x2c, 0xe4, 0x79, 0x94, 0x23, 0x4e, 0xa8, 0xc7, 0xa2, 0x6c, 0x3b, 0xca, 0x26, 0x18, 0x9c,
	0x2c, 0x30, 0xe3, 0x68, 0xe1, 0x47, 0x05, 0x4f, 0xe4, 0xd7, 0xec, 0xd8, 0xc5, 0xde, 0x31, 0x5b,
	0x22, 0xd7, 0xc5, 0x81, 0x45, 0x7d, 0x09, 0xf1, 0x5f, 0x38, 0xf3, 0x8f, 0x3c, 0x28, 0x5c, 0x31,
	0x1c, 0xc0, 0x43, 0x90, 0x23, 0x4e, 0x4b, 0x33, 0xb4, 0x4e, 0x7e, 0x50, 0x59, 0xaf, 0xda, 0xdb,
	0xce, 0xb4, 0x67, 0x12, 0xc7, 0x34, 0xec, 0x1c, 0x71, 0xe0, 0x07, 0xa0, 0x88, 0x17, 0x88, 0xcc,
	0x5b, 0x39, 0x43, 0xeb, 0xec, 0x0c, 0x6a, 0xeb, 0x55, 0xbb, 0x2c, 0x0a, 0x64, 0xd0, 0x34, 0x6c,
	0x95, 0x85, 0x5d, 0xb0, 0xed, 0x23, 0xc6, 0x96, 0x34, 0x70, 0x5a, 0x79, 0x59, 0x09, 0xd7, 0xab,
	0x76, 0x55, 0x54, 0xc6, 0x71, 0xd3, 0xb0, 0x93, 0x1a, 0xd8, 0x03, 0x25, 0xc6, 0x11, 0x0f, 0x59,
	0xab, 0x60, 0x68, 0x9d, 0xea, 0x49, 0xbd, 0x2b, 0xf5, 0x16, 0x84, 0x46, 0x32, 0x3e, 0xa8, 0xaf,
	0x57, 0xed, 0x8a, 0xe8, 0x57, 0x75, 0xa6, 0x61, 0x47, 0x1d, 0xf0, 0x0b, 0x00, 0x66, 0x01, 0x46,
	0x1c, 0x3b, 0x13, 0xc4, 0x5b, 0x45, 0x43, 0xeb, 0x94, 0x4f, 0xf4, 0xae, 0x52, 0xa7, 0x1b, 0xab,
	0xd3, 0x1d, 0xc7, 0xea, 0x0c, 0x1a, 0xeb, 0x55, 0xbb, 0x2e, 0x90, 0xd2, 0x2e, 0xd3, 0xb0, 0x77,
	0xa2, 0x5b, 0x9f, 0x0b, 0xc0, 0xd0, 0x77, 0x62, 0xc0, 0xd2, 0xc3, 0x01, 0xd3, 0x2e, 0x01, 0x18,
	0xdd, 0x14, 0xe0, 0x1c, 0x31, 0x3e, 0x99, 0x53, 0x97, 0x78, 0xad, 0x77, 0x1e, 0x0e, 0x98, 0x76,
	0x09, 0x40, 0x71, 0x3b, 0x17, 0x97, 0xde, 0xee, 0x9f, 0x6f, 0x0e, 0xb4, 0xbf, 0xdf, 0x1c, 0x14,
	0x85, 0x4a, 0xcc, 0xe4, 0xa0, 0x22, 0x94, 0xb2, 0x31, 0xf3, 0xa9, 0xc7, 0x30, 0xac, 0xa6, 0x2b,
	0x94, 0x4b, 0x6b, 0xdc, 0x5a, 0x5a, 0xbc, 0xa3, 0x06, 0x28, 0x72, 0xfa, 0x15, 0xf6, 0xd4, 0x82,
	0x6c, 0x75, 0x81, 0x9d, 0xfb, 0x36, 0x11, 0xeb, 0x6e, 0x7e, 0xa7, 0x81, 0x8a, 0xa4, 0x63, 0xe3,
	0x9b, 0x10, 0x33, 0x0e, 0x9f, 0xc5, 0x63, 0x34, 0xb9, 0xf1, 0xf7, 0xd7, 0xab, 0xb6, 0xfe, 0x0a,
	0xcd, 0x89, 0x90, 0x21, 0x72, 0xc8, 0x93, 0x00, 0xdf, 0x84, 0x24, 0xc0, 0x4e, 0x6a, 0x95, 0x5e,
	0xc6, 0x2a, 0xb9, 0x4d, 0x8d, 0x2e, 0xc7, 0x9f, 0x7c, 0x9c, 0x6d, 0x4c, 0xea, 0xcd, 0x1a, 0xd8,
	0x3d, 0xa7, 0x2e, 0x0d, 0x79, 0x44, 0xc1, 0xac, 0x83, 0x6a, 0x1c, 0x50, 0x5a, 0x98, 0x3f, 0x68,
	0xa0, 0x66, 0x63, 0x97, 0x30, 0x2e, 0x04, 0x7a, 0x5b, 0x44, 0x77, 0x41, 0xf9, 0x92, 0x78, 0x6e,
	0x4c, 0xf3, 0x57, 0x0d, 0x3c, 0x7a, 0x7e, 0x8d, 0x3c, 0x17, 0x5f, 0x46, 0x15, 0x31, 0xb5, 0x3e,
	0xa8, 0xd0, 0xb9, 0x33, 0x49, 0x06, 0x69, 0x0f, 0x1a, 0x54, 0xa6, 0x73, 0x27, 0x46, 0x12, 0x10,
	0x1e, 0x5e, 0x4e, 0xfe, 0x27, 0xd7, 0xb2, 0x87, 0x97, 0x31, 0x84, 0xd9, 0x02, 0xcd, 0xbb, 0xf4,
	0x94, 0x9c, 0x47, 0x37, 0x00, 0xa4, 0x56, 0x80, 0xef, 0x82, 0xfd, 0xab, 0xd1, 0xd0, 0x9e, 0x8c,
	0xc6, 0xfd, 0xf1, 0xd5, 0x68, 0x72, 0x76, 0xf1, 0xa2, 0x7f, 0x7e, 0x76, 0x5a, 0xdf, 0x82, 0x3a,
	0x68, 0x66, 0x13, 0xf6, 0xf0, 0xb3, 0xb3, 0xd1, 0x78, 0x68, 0x0f, 0x4f, 0xeb, 0x1a, 0x6c, 0x02,
	0x98, 0xcd, 0xf5, 0x9f, 0x8f, 0xcf, 0x5e, 0x0c, 0xeb, 0xb9, 0xbb, 0xf1, 0x41, 0xff, 0xe2, 0x62,
	0x78, 0x5a, 0xcf, 0x9f, 0xfc, 0x58, 0x88, 0x66, 0x7e, 0xc3, 0x38, 0x5e, 0xc0, 0xcf, 0x41, 0x51,
	0xba, 0x0e, 0x42, 0xe5, 0xcc, 0xac, 0x05, 0x75, 0x98, 0xba, 0x35, 0x71, 0xc0, 0xc1, 0xf7, 0x7f,
	0xfd, 0xf3, 0x4b, 0x6e, 0xdf, 0xac, 0x5a, 0xaf, 0x9e, 0xaa, 0x17, 0x5d, 0xfe, 0xa3, 0x7a, 0xda,
	0x11, 0x24, 0xa0, 0xa4, 0xec, 0x02, 0xf7, 0x13, 0xb0, 0xd4, 0x4d, 0x7a, 0xe3, 0x76, 0x30, 0xc2,
	0x7b, 0xf6, 0x73, 0xbf, 0x39, 0x6d, 0x00, 0x08, 0xaa, 0xfd, 0x90, 0x5f, 0x63, 0x8f, 0x93, 0x99,
	0x7c, 0x49, 0xe1, 0xd6, 0xeb, 0xdf, 0xf4, 0x2d, 0x39, 0x6c, 0x0f, 0xd6, 0xb2, 0xc3, 0xc4, 0x80,
	0x11, 0xd8, 0x8e, 0x6d, 0x08, 0x1f, 0x29, 0xdc, 0x3b, 0xb6, 0xdc, 0x48, 0xfe, 0x50, 0xe2, 0x35,
	0xcd, 0xbd, 0x04, 0x2f, 0x88, 0xba, 0x04, 0x7f, 0x04, 0x0a, 0xc2, 0x56, 0x70, 0x4f, 0x75, 0x66,
	0x2c, 0xb6, 0x11, 0xec, 0xe4, 0x7e, 0xe6, 0x35, 0xb8, 0x9b, 0x4c, 0xf2, 0x05, 0xf4, 0x4f, 0x1a,
	0xa8, 0xde, 0xf6, 0x02, 0x7c, 0x4f, 0x41, 0x6f, 0x34, 0xb0, 0x7e, 0xb8, 0x39, 0x19, 0x31, 0xf8,
	0xf4, 0x7e, 0x06, 0x8f, 0xcd, 0x56, 0xc2, 0x60, 0x26, 0x51, 0xac, 0xd8, 0xca, 0x3d, 0xed, 0x68,
	0xf0, 0xe1, 0xeb, 0xdf, 0x0f, 0xf2, 0x08, 0x21, 0xb0, 0x3d, 0xa3, 0x0b, 0x39, 0x6b, 0xb0, 0x23,
	0x7e, 0xe4, 0xa5, 0x78, 0x46, 0x2f, 0xb5, 0x2f, 0x4b, 0x22, 0xe4, 0x4f, 0xa7, 0x25, 0xf9, 0xae,
	0x7e, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x65, 0x18, 0x82, 0xcc, 0x07, 0x00, 0x00,
}
