// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package session

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Cookie struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Expiration           string   `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cookie) Reset()         { *m = Cookie{} }
func (m *Cookie) String() string { return proto.CompactTextString(m) }
func (*Cookie) ProtoMessage()    {}
func (*Cookie) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Cookie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookie.Unmarshal(m, b)
}
func (m *Cookie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookie.Marshal(b, m, deterministic)
}
func (m *Cookie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookie.Merge(m, src)
}
func (m *Cookie) XXX_Size() int {
	return xxx_messageInfo_Cookie.Size(m)
}
func (m *Cookie) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookie.DiscardUnknown(m)
}

var xxx_messageInfo_Cookie proto.InternalMessageInfo

func (m *Cookie) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Cookie) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

type UserID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

type User struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	HashPassword         string   `protobuf:"bytes,4,opt,name=hashPassword,proto3" json:"hashPassword,omitempty"`
	Score                int32    `protobuf:"varint,5,opt,name=score,proto3" json:"score,omitempty"`
	AvatarLink           string   `protobuf:"bytes,6,opt,name=avatarLink,proto3" json:"avatarLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
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

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetHashPassword() string {
	if m != nil {
		return m.HashPassword
	}
	return ""
}

func (m *User) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *User) GetAvatarLink() string {
	if m != nil {
		return m.AvatarLink
	}
	return ""
}

type UserNew struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserNew) Reset()         { *m = UserNew{} }
func (m *UserNew) String() string { return proto.CompactTextString(m) }
func (*UserNew) ProtoMessage()    {}
func (*UserNew) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *UserNew) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserNew.Unmarshal(m, b)
}
func (m *UserNew) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserNew.Marshal(b, m, deterministic)
}
func (m *UserNew) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserNew.Merge(m, src)
}
func (m *UserNew) XXX_Size() int {
	return xxx_messageInfo_UserNew.Size(m)
}
func (m *UserNew) XXX_DiscardUnknown() {
	xxx_messageInfo_UserNew.DiscardUnknown(m)
}

var xxx_messageInfo_UserNew proto.InternalMessageInfo

func (m *UserNew) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserNew) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserNew) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UpdateUserReq struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	NewAvatar            string   `protobuf:"bytes,2,opt,name=newAvatar,proto3" json:"newAvatar,omitempty"`
	OldPassword          string   `protobuf:"bytes,3,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword          string   `protobuf:"bytes,4,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserReq) Reset()         { *m = UpdateUserReq{} }
func (m *UpdateUserReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq) ProtoMessage()    {}
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *UpdateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq.Unmarshal(m, b)
}
func (m *UpdateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq.Merge(m, src)
}
func (m *UpdateUserReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq.Size(m)
}
func (m *UpdateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq proto.InternalMessageInfo

func (m *UpdateUserReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UpdateUserReq) GetNewAvatar() string {
	if m != nil {
		return m.NewAvatar
	}
	return ""
}

func (m *UpdateUserReq) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *UpdateUserReq) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ScoresParam struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScoresParam) Reset()         { *m = ScoresParam{} }
func (m *ScoresParam) String() string { return proto.CompactTextString(m) }
func (*ScoresParam) ProtoMessage()    {}
func (*ScoresParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *ScoresParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScoresParam.Unmarshal(m, b)
}
func (m *ScoresParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScoresParam.Marshal(b, m, deterministic)
}
func (m *ScoresParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScoresParam.Merge(m, src)
}
func (m *ScoresParam) XXX_Size() int {
	return xxx_messageInfo_ScoresParam.Size(m)
}
func (m *ScoresParam) XXX_DiscardUnknown() {
	xxx_messageInfo_ScoresParam.DiscardUnknown(m)
}

var xxx_messageInfo_ScoresParam proto.InternalMessageInfo

func (m *ScoresParam) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ScoresParam) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Score struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Score                int32    `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Score) Reset()         { *m = Score{} }
func (m *Score) String() string { return proto.CompactTextString(m) }
func (*Score) ProtoMessage()    {}
func (*Score) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *Score) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Score.Unmarshal(m, b)
}
func (m *Score) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Score.Marshal(b, m, deterministic)
}
func (m *Score) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Score.Merge(m, src)
}
func (m *Score) XXX_Size() int {
	return xxx_messageInfo_Score.Size(m)
}
func (m *Score) XXX_DiscardUnknown() {
	xxx_messageInfo_Score.DiscardUnknown(m)
}

var xxx_messageInfo_Score proto.InternalMessageInfo

func (m *Score) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Score) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type Scores struct {
	Scores               []*Score `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scores) Reset()         { *m = Scores{} }
func (m *Scores) String() string { return proto.CompactTextString(m) }
func (*Scores) ProtoMessage()    {}
func (*Scores) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
}

func (m *Scores) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scores.Unmarshal(m, b)
}
func (m *Scores) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scores.Marshal(b, m, deterministic)
}
func (m *Scores) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scores.Merge(m, src)
}
func (m *Scores) XXX_Size() int {
	return xxx_messageInfo_Scores.Size(m)
}
func (m *Scores) XXX_DiscardUnknown() {
	xxx_messageInfo_Scores.DiscardUnknown(m)
}

var xxx_messageInfo_Scores proto.InternalMessageInfo

func (m *Scores) GetScores() []*Score {
	if m != nil {
		return m.Scores
	}
	return nil
}

type Num struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Num) Reset()         { *m = Num{} }
func (m *Num) String() string { return proto.CompactTextString(m) }
func (*Num) ProtoMessage()    {}
func (*Num) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{9}
}

func (m *Num) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Num.Unmarshal(m, b)
}
func (m *Num) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Num.Marshal(b, m, deterministic)
}
func (m *Num) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Num.Merge(m, src)
}
func (m *Num) XXX_Size() int {
	return xxx_messageInfo_Num.Size(m)
}
func (m *Num) XXX_DiscardUnknown() {
	xxx_messageInfo_Num.DiscardUnknown(m)
}

var xxx_messageInfo_Num proto.InternalMessageInfo

func (m *Num) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type UpdateScoreReq struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Score                int32    `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateScoreReq) Reset()         { *m = UpdateScoreReq{} }
func (m *UpdateScoreReq) String() string { return proto.CompactTextString(m) }
func (*UpdateScoreReq) ProtoMessage()    {}
func (*UpdateScoreReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{10}
}

func (m *UpdateScoreReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateScoreReq.Unmarshal(m, b)
}
func (m *UpdateScoreReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateScoreReq.Marshal(b, m, deterministic)
}
func (m *UpdateScoreReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateScoreReq.Merge(m, src)
}
func (m *UpdateScoreReq) XXX_Size() int {
	return xxx_messageInfo_UpdateScoreReq.Size(m)
}
func (m *UpdateScoreReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateScoreReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateScoreReq proto.InternalMessageInfo

func (m *UpdateScoreReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UpdateScoreReq) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type DataAuth struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataAuth) Reset()         { *m = DataAuth{} }
func (m *DataAuth) String() string { return proto.CompactTextString(m) }
func (*DataAuth) ProtoMessage()    {}
func (*DataAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{11}
}

func (m *DataAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataAuth.Unmarshal(m, b)
}
func (m *DataAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataAuth.Marshal(b, m, deterministic)
}
func (m *DataAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataAuth.Merge(m, src)
}
func (m *DataAuth) XXX_Size() int {
	return xxx_messageInfo_DataAuth.Size(m)
}
func (m *DataAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_DataAuth.DiscardUnknown(m)
}

var xxx_messageInfo_DataAuth proto.InternalMessageInfo

func (m *DataAuth) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *DataAuth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Cookie)(nil), "session.Cookie")
	proto.RegisterType((*UserID)(nil), "session.UserID")
	proto.RegisterType((*Nothing)(nil), "session.Nothing")
	proto.RegisterType((*User)(nil), "session.User")
	proto.RegisterType((*UserNew)(nil), "session.UserNew")
	proto.RegisterType((*UpdateUserReq)(nil), "session.UpdateUserReq")
	proto.RegisterType((*ScoresParam)(nil), "session.ScoresParam")
	proto.RegisterType((*Score)(nil), "session.Score")
	proto.RegisterType((*Scores)(nil), "session.Scores")
	proto.RegisterType((*Num)(nil), "session.Num")
	proto.RegisterType((*UpdateScoreReq)(nil), "session.UpdateScoreReq")
	proto.RegisterType((*DataAuth)(nil), "session.DataAuth")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0xd3, 0xda, 0x6d, 0x27, 0x4d, 0x7e, 0xfd, 0xad, 0xaa, 0x62, 0x05, 0x84, 0xa2, 0x3d,
	0xa0, 0x1e, 0xa0, 0x40, 0x8b, 0x2a, 0x0a, 0x08, 0xa9, 0xc4, 0xa2, 0xb2, 0x84, 0xa2, 0xc8, 0xa8,
	0xe2, 0xbc, 0x24, 0xdb, 0x7a, 0x95, 0x78, 0x37, 0x78, 0x37, 0x84, 0x1e, 0x39, 0xf3, 0x2d, 0xf8,
	0xa4, 0x68, 0x77, 0xfd, 0x37, 0x71, 0x05, 0xc7, 0x79, 0x3b, 0x6f, 0xe6, 0xcd, 0xf3, 0x8c, 0x01,
	0xc8, 0x52, 0xc5, 0x27, 0x8b, 0x54, 0x28, 0x81, 0x76, 0x24, 0x95, 0x92, 0x09, 0x8e, 0xdf, 0x83,
	0x37, 0x14, 0x62, 0xc6, 0x28, 0x3a, 0x04, 0x57, 0x89, 0x19, 0xe5, 0xbe, 0x33, 0x70, 0x8e, 0xf7,
	0x22, 0x1b, 0xa0, 0xc7, 0x00, 0xf4, 0xc7, 0x82, 0xa5, 0x44, 0x31, 0xc1, 0xfd, 0xb6, 0x79, 0xaa,
	0x20, 0xd8, 0x07, 0xef, 0x5a, 0xd2, 0x34, 0x0c, 0x50, 0x0f, 0xda, 0x61, 0x90, 0x91, 0xdb, 0x61,
	0x80, 0xf7, 0x60, 0x67, 0x24, 0x54, 0xcc, 0xf8, 0x2d, 0xfe, 0xed, 0xc0, 0xb6, 0xce, 0x5a, 0xcf,
	0xd1, 0x3d, 0x69, 0x42, 0xd8, 0x3c, 0x2b, 0x6c, 0x03, 0xd4, 0x87, 0x5d, 0xce, 0x26, 0x33, 0x4e,
	0x12, 0xea, 0x6f, 0x99, 0x87, 0x22, 0x46, 0x18, 0xf6, 0x63, 0x22, 0xe3, 0x31, 0x91, 0x72, 0x25,
	0xd2, 0xa9, 0xbf, 0x6d, 0xde, 0x6b, 0x98, 0xae, 0x2a, 0x27, 0x22, 0xa5, 0xbe, 0x3b, 0x70, 0x8e,
	0xdd, 0xc8, 0x06, 0x7a, 0x12, 0xf2, 0x9d, 0x28, 0x92, 0x7e, 0x62, 0x7c, 0xe6, 0x7b, 0x76, 0x92,
	0x12, 0xc1, 0x5f, 0x60, 0x47, 0x6b, 0x1c, 0xd1, 0x55, 0x4d, 0x80, 0xb3, 0x26, 0xe0, 0x5e, 0xc9,
	0x8b, 0x5c, 0x52, 0x26, 0x39, 0x8f, 0xf1, 0x4f, 0x07, 0xba, 0xd7, 0x8b, 0x29, 0x51, 0x54, 0xd7,
	0x8f, 0xe8, 0xb7, 0x0d, 0x1b, 0x1e, 0xc1, 0x1e, 0xa7, 0xab, 0x4b, 0xa3, 0x25, 0xab, 0x5b, 0x02,
	0x68, 0x00, 0x1d, 0x31, 0x9f, 0x8e, 0xeb, 0xe5, 0xab, 0x90, 0xce, 0xe0, 0x74, 0xb5, 0xe6, 0x49,
	0x15, 0xc2, 0x6f, 0xa1, 0xf3, 0x59, 0xbb, 0x20, 0xc7, 0x24, 0x25, 0x89, 0x1e, 0x62, 0xce, 0x12,
	0xa6, 0x8c, 0x06, 0x37, 0xb2, 0x01, 0x3a, 0x02, 0x4f, 0xdc, 0xdc, 0x48, 0xaa, 0x8c, 0x06, 0x37,
	0xca, 0x22, 0x7c, 0x01, 0xae, 0x21, 0xff, 0xcd, 0x17, 0x6b, 0x7a, 0xbb, 0x62, 0x3a, 0x7e, 0x01,
	0x9e, 0xed, 0x8b, 0x9e, 0x80, 0x67, 0x20, 0xe9, 0x3b, 0x83, 0xad, 0xe3, 0xce, 0x69, 0xef, 0x24,
	0x5b, 0xc1, 0x13, 0x93, 0x10, 0x65, 0xaf, 0xf8, 0x21, 0x6c, 0x8d, 0x96, 0x46, 0xe1, 0x44, 0x2c,
	0x79, 0xa1, 0xd0, 0x04, 0xf8, 0x1c, 0x7a, 0xd6, 0x49, 0xcb, 0x69, 0xb0, 0xb2, 0x59, 0xc6, 0x3b,
	0xd8, 0x0d, 0x88, 0x22, 0x97, 0x4b, 0x15, 0x9b, 0xd9, 0xc5, 0x2d, 0x2b, 0xf6, 0xdc, 0x04, 0xb5,
	0x0f, 0xd8, 0xae, 0x7f, 0xc0, 0xd3, 0x5f, 0x2e, 0x74, 0x34, 0x75, 0x18, 0xd3, 0xc9, 0x8c, 0xa6,
	0xe8, 0x0c, 0xba, 0xc3, 0x94, 0x6a, 0x15, 0x76, 0x02, 0xf4, 0x5f, 0x31, 0x8b, 0xbd, 0x85, 0x7e,
	0x09, 0xd8, 0xe3, 0xc2, 0x2d, 0xf4, 0x0a, 0xba, 0x01, 0x9d, 0xd3, 0x26, 0x92, 0xcd, 0xe9, 0x1f,
	0x14, 0x40, 0x7e, 0x37, 0x2d, 0xdd, 0x2a, 0x1b, 0xf8, 0x3e, 0x56, 0x43, 0xab, 0x73, 0x38, 0xb8,
	0xa2, 0x2a, 0x0c, 0x3e, 0xa6, 0x22, 0xf9, 0x07, 0x9e, 0xd5, 0x8c, 0x5b, 0xe8, 0x39, 0x80, 0x9d,
	0xcb, 0xdc, 0xea, 0x41, 0x2d, 0x61, 0x44, 0x57, 0xfd, 0x6e, 0x0d, 0xc1, 0x2d, 0x74, 0x0a, 0xfb,
	0xe1, 0x94, 0x72, 0xc5, 0x6e, 0xee, 0x0c, 0xe5, 0xff, 0x22, 0x21, 0x77, 0x7b, 0x93, 0xf3, 0x0c,
	0x3a, 0x57, 0x54, 0xe9, 0xe0, 0xc3, 0x5d, 0x18, 0xa0, 0xfa, 0xfb, 0x66, 0xfa, 0x6b, 0x80, 0xf2,
	0x76, 0xd0, 0x51, 0xf9, 0x5c, 0x3d, 0xa8, 0x46, 0xeb, 0x2e, 0xa0, 0x97, 0x35, 0x92, 0xd9, 0x0a,
	0x1e, 0xd6, 0x57, 0xce, 0xde, 0x42, 0xc5, 0x08, 0x8b, 0xe2, 0x16, 0x7a, 0x09, 0xdd, 0x9c, 0x3a,
	0xd4, 0x7b, 0x87, 0x36, 0xea, 0xf7, 0xf7, 0x4b, 0x64, 0x99, 0xe0, 0x16, 0x7a, 0x03, 0x9d, 0xca,
	0x66, 0xa2, 0x07, 0x6b, 0x42, 0xf3, 0x7d, 0x6d, 0x54, 0xfa, 0x14, 0xb6, 0xc7, 0x8c, 0xdf, 0x36,
	0x74, 0x69, 0xc8, 0xfe, 0xea, 0x99, 0x3f, 0xf8, 0xd9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38,
	0xb9, 0x63, 0x8d, 0xcf, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthCheckerClient is the client API for AuthChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthCheckerClient interface {
	CreateSession(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Cookie, error)
	DeleteSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Nothing, error)
	UpdateSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Cookie, error)
	GetIDFromSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*UserID, error)
	CreateUser(ctx context.Context, in *UserNew, opts ...grpc.CallOption) (*User, error)
	IdentifyUser(ctx context.Context, in *DataAuth, opts ...grpc.CallOption) (*User, error)
	GetUserByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*Nothing, error)
	GetUsersScores(ctx context.Context, in *ScoresParam, opts ...grpc.CallOption) (*Scores, error)
	GetUsersCount(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Num, error)
	UpdateScore(ctx context.Context, in *UpdateScoreReq, opts ...grpc.CallOption) (*Nothing, error)
	Ping(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Nothing, error)
}

type authCheckerClient struct {
	cc *grpc.ClientConn
}

func NewAuthCheckerClient(cc *grpc.ClientConn) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) CreateSession(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Cookie, error) {
	out := new(Cookie)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) DeleteSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) UpdateSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Cookie, error) {
	out := new(Cookie)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/UpdateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetIDFromSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/GetIDFromSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) CreateUser(ctx context.Context, in *UserNew, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) IdentifyUser(ctx context.Context, in *DataAuth, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/IdentifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetUserByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetUsersScores(ctx context.Context, in *ScoresParam, opts ...grpc.CallOption) (*Scores, error) {
	out := new(Scores)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/GetUsersScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetUsersCount(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Num, error) {
	out := new(Num)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/GetUsersCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) UpdateScore(ctx context.Context, in *UpdateScoreReq, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/UpdateScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) Ping(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
type AuthCheckerServer interface {
	CreateSession(context.Context, *UserID) (*Cookie, error)
	DeleteSession(context.Context, *Cookie) (*Nothing, error)
	UpdateSession(context.Context, *Cookie) (*Cookie, error)
	GetIDFromSession(context.Context, *Cookie) (*UserID, error)
	CreateUser(context.Context, *UserNew) (*User, error)
	IdentifyUser(context.Context, *DataAuth) (*User, error)
	GetUserByID(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *UpdateUserReq) (*Nothing, error)
	GetUsersScores(context.Context, *ScoresParam) (*Scores, error)
	GetUsersCount(context.Context, *Nothing) (*Num, error)
	UpdateScore(context.Context, *UpdateScoreReq) (*Nothing, error)
	Ping(context.Context, *Nothing) (*Nothing, error)
}

func RegisterAuthCheckerServer(s *grpc.Server, srv AuthCheckerServer) {
	s.RegisterService(&_AuthChecker_serviceDesc, srv)
}

func _AuthChecker_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).CreateSession(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).DeleteSession(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_UpdateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).UpdateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/UpdateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).UpdateSession(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetIDFromSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetIDFromSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/GetIDFromSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetIDFromSession(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).CreateUser(ctx, req.(*UserNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_IdentifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).IdentifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/IdentifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).IdentifyUser(ctx, req.(*DataAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetUserByID(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetUsersScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoresParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetUsersScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/GetUsersScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetUsersScores(ctx, req.(*ScoresParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetUsersCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetUsersCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/GetUsersCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetUsersCount(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_UpdateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).UpdateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/UpdateScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).UpdateScore(ctx, req.(*UpdateScoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Ping(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _AuthChecker_CreateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _AuthChecker_DeleteSession_Handler,
		},
		{
			MethodName: "UpdateSession",
			Handler:    _AuthChecker_UpdateSession_Handler,
		},
		{
			MethodName: "GetIDFromSession",
			Handler:    _AuthChecker_GetIDFromSession_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AuthChecker_CreateUser_Handler,
		},
		{
			MethodName: "IdentifyUser",
			Handler:    _AuthChecker_IdentifyUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _AuthChecker_GetUserByID_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthChecker_UpdateUser_Handler,
		},
		{
			MethodName: "GetUsersScores",
			Handler:    _AuthChecker_GetUsersScores_Handler,
		},
		{
			MethodName: "GetUsersCount",
			Handler:    _AuthChecker_GetUsersCount_Handler,
		},
		{
			MethodName: "UpdateScore",
			Handler:    _AuthChecker_UpdateScore_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _AuthChecker_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}