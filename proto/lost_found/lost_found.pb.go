// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/lost_found/lost_found.proto

package lostfound

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LostAndFoundType int32

const (
	LostAndFoundType_Lost  LostAndFoundType = 0
	LostAndFoundType_Found LostAndFoundType = 1
)

var LostAndFoundType_name = map[int32]string{
	0: "Lost",
	1: "Found",
}

var LostAndFoundType_value = map[string]int32{
	"Lost":  0,
	"Found": 1,
}

func (x LostAndFoundType) String() string {
	return proto.EnumName(LostAndFoundType_name, int32(x))
}

func (LostAndFoundType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1a49343fa726016, []int{0}
}

type GetBriefsReq struct {
	Before               *timestamp.Timestamp `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetBriefsReq) Reset()         { *m = GetBriefsReq{} }
func (m *GetBriefsReq) String() string { return proto.CompactTextString(m) }
func (*GetBriefsReq) ProtoMessage()    {}
func (*GetBriefsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1a49343fa726016, []int{0}
}

func (m *GetBriefsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBriefsReq.Unmarshal(m, b)
}
func (m *GetBriefsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBriefsReq.Marshal(b, m, deterministic)
}
func (m *GetBriefsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBriefsReq.Merge(m, src)
}
func (m *GetBriefsReq) XXX_Size() int {
	return xxx_messageInfo_GetBriefsReq.Size(m)
}
func (m *GetBriefsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBriefsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetBriefsReq proto.InternalMessageInfo

func (m *GetBriefsReq) GetBefore() *timestamp.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

type GetBriefsResp struct {
	Briefs               []*LostAndFoundBrief `protobuf:"bytes,1,rep,name=briefs,proto3" json:"briefs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetBriefsResp) Reset()         { *m = GetBriefsResp{} }
func (m *GetBriefsResp) String() string { return proto.CompactTextString(m) }
func (*GetBriefsResp) ProtoMessage()    {}
func (*GetBriefsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1a49343fa726016, []int{1}
}

func (m *GetBriefsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBriefsResp.Unmarshal(m, b)
}
func (m *GetBriefsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBriefsResp.Marshal(b, m, deterministic)
}
func (m *GetBriefsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBriefsResp.Merge(m, src)
}
func (m *GetBriefsResp) XXX_Size() int {
	return xxx_messageInfo_GetBriefsResp.Size(m)
}
func (m *GetBriefsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBriefsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetBriefsResp proto.InternalMessageInfo

func (m *GetBriefsResp) GetBriefs() []*LostAndFoundBrief {
	if m != nil {
		return m.Briefs
	}
	return nil
}

type LostAndFoundBrief struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string               `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Type                 LostAndFoundType     `protobuf:"varint,3,opt,name=type,proto3,enum=xmux.lost_found.v4.LostAndFoundType" json:"type,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Location             string               `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LostAndFoundBrief) Reset()         { *m = LostAndFoundBrief{} }
func (m *LostAndFoundBrief) String() string { return proto.CompactTextString(m) }
func (*LostAndFoundBrief) ProtoMessage()    {}
func (*LostAndFoundBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1a49343fa726016, []int{2}
}

func (m *LostAndFoundBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LostAndFoundBrief.Unmarshal(m, b)
}
func (m *LostAndFoundBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LostAndFoundBrief.Marshal(b, m, deterministic)
}
func (m *LostAndFoundBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LostAndFoundBrief.Merge(m, src)
}
func (m *LostAndFoundBrief) XXX_Size() int {
	return xxx_messageInfo_LostAndFoundBrief.Size(m)
}
func (m *LostAndFoundBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_LostAndFoundBrief.DiscardUnknown(m)
}

var xxx_messageInfo_LostAndFoundBrief proto.InternalMessageInfo

func (m *LostAndFoundBrief) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LostAndFoundBrief) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LostAndFoundBrief) GetType() LostAndFoundType {
	if m != nil {
		return m.Type
	}
	return LostAndFoundType_Lost
}

func (m *LostAndFoundBrief) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LostAndFoundBrief) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *LostAndFoundBrief) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type GetDetailReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailReq) Reset()         { *m = GetDetailReq{} }
func (m *GetDetailReq) String() string { return proto.CompactTextString(m) }
func (*GetDetailReq) ProtoMessage()    {}
func (*GetDetailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1a49343fa726016, []int{3}
}

func (m *GetDetailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailReq.Unmarshal(m, b)
}
func (m *GetDetailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailReq.Marshal(b, m, deterministic)
}
func (m *GetDetailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailReq.Merge(m, src)
}
func (m *GetDetailReq) XXX_Size() int {
	return xxx_messageInfo_GetDetailReq.Size(m)
}
func (m *GetDetailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailReq proto.InternalMessageInfo

func (m *GetDetailReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type LostAndFoundDetail struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type                 LostAndFoundType     `protobuf:"varint,2,opt,name=type,proto3,enum=xmux.lost_found.v4.LostAndFoundType" json:"type,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Location             string               `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Description          string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Contacts             map[string]string    `protobuf:"bytes,7,rep,name=contacts,proto3" json:"contacts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Pictures             []string             `protobuf:"bytes,8,rep,name=pictures,proto3" json:"pictures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LostAndFoundDetail) Reset()         { *m = LostAndFoundDetail{} }
func (m *LostAndFoundDetail) String() string { return proto.CompactTextString(m) }
func (*LostAndFoundDetail) ProtoMessage()    {}
func (*LostAndFoundDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1a49343fa726016, []int{4}
}

func (m *LostAndFoundDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LostAndFoundDetail.Unmarshal(m, b)
}
func (m *LostAndFoundDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LostAndFoundDetail.Marshal(b, m, deterministic)
}
func (m *LostAndFoundDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LostAndFoundDetail.Merge(m, src)
}
func (m *LostAndFoundDetail) XXX_Size() int {
	return xxx_messageInfo_LostAndFoundDetail.Size(m)
}
func (m *LostAndFoundDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_LostAndFoundDetail.DiscardUnknown(m)
}

var xxx_messageInfo_LostAndFoundDetail proto.InternalMessageInfo

func (m *LostAndFoundDetail) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LostAndFoundDetail) GetType() LostAndFoundType {
	if m != nil {
		return m.Type
	}
	return LostAndFoundType_Lost
}

func (m *LostAndFoundDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LostAndFoundDetail) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *LostAndFoundDetail) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *LostAndFoundDetail) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LostAndFoundDetail) GetContacts() map[string]string {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *LostAndFoundDetail) GetPictures() []string {
	if m != nil {
		return m.Pictures
	}
	return nil
}

type AddItemReq struct {
	Type                 LostAndFoundType     `protobuf:"varint,1,opt,name=type,proto3,enum=xmux.lost_found.v4.LostAndFoundType" json:"type,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Location             string               `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Contacts             map[string]string    `protobuf:"bytes,6,rep,name=contacts,proto3" json:"contacts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Pictures             []string             `protobuf:"bytes,7,rep,name=pictures,proto3" json:"pictures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AddItemReq) Reset()         { *m = AddItemReq{} }
func (m *AddItemReq) String() string { return proto.CompactTextString(m) }
func (*AddItemReq) ProtoMessage()    {}
func (*AddItemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1a49343fa726016, []int{5}
}

func (m *AddItemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddItemReq.Unmarshal(m, b)
}
func (m *AddItemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddItemReq.Marshal(b, m, deterministic)
}
func (m *AddItemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddItemReq.Merge(m, src)
}
func (m *AddItemReq) XXX_Size() int {
	return xxx_messageInfo_AddItemReq.Size(m)
}
func (m *AddItemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddItemReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddItemReq proto.InternalMessageInfo

func (m *AddItemReq) GetType() LostAndFoundType {
	if m != nil {
		return m.Type
	}
	return LostAndFoundType_Lost
}

func (m *AddItemReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddItemReq) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *AddItemReq) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *AddItemReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddItemReq) GetContacts() map[string]string {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *AddItemReq) GetPictures() []string {
	if m != nil {
		return m.Pictures
	}
	return nil
}

type DeleteItemReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteItemReq) Reset()         { *m = DeleteItemReq{} }
func (m *DeleteItemReq) String() string { return proto.CompactTextString(m) }
func (*DeleteItemReq) ProtoMessage()    {}
func (*DeleteItemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1a49343fa726016, []int{6}
}

func (m *DeleteItemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteItemReq.Unmarshal(m, b)
}
func (m *DeleteItemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteItemReq.Marshal(b, m, deterministic)
}
func (m *DeleteItemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteItemReq.Merge(m, src)
}
func (m *DeleteItemReq) XXX_Size() int {
	return xxx_messageInfo_DeleteItemReq.Size(m)
}
func (m *DeleteItemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteItemReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteItemReq proto.InternalMessageInfo

func (m *DeleteItemReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("xmux.lost_found.v4.LostAndFoundType", LostAndFoundType_name, LostAndFoundType_value)
	proto.RegisterType((*GetBriefsReq)(nil), "xmux.lost_found.v4.GetBriefsReq")
	proto.RegisterType((*GetBriefsResp)(nil), "xmux.lost_found.v4.GetBriefsResp")
	proto.RegisterType((*LostAndFoundBrief)(nil), "xmux.lost_found.v4.LostAndFoundBrief")
	proto.RegisterType((*GetDetailReq)(nil), "xmux.lost_found.v4.GetDetailReq")
	proto.RegisterType((*LostAndFoundDetail)(nil), "xmux.lost_found.v4.LostAndFoundDetail")
	proto.RegisterMapType((map[string]string)(nil), "xmux.lost_found.v4.LostAndFoundDetail.ContactsEntry")
	proto.RegisterType((*AddItemReq)(nil), "xmux.lost_found.v4.AddItemReq")
	proto.RegisterMapType((map[string]string)(nil), "xmux.lost_found.v4.AddItemReq.ContactsEntry")
	proto.RegisterType((*DeleteItemReq)(nil), "xmux.lost_found.v4.DeleteItemReq")
}

func init() {
	proto.RegisterFile("proto/lost_found/lost_found.proto", fileDescriptor_e1a49343fa726016)
}

var fileDescriptor_e1a49343fa726016 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xf1, 0x47, 0xd2, 0x64, 0xd2, 0x54, 0x61, 0x84, 0x90, 0x65, 0xa4, 0x36, 0xb5, 0xf8,
	0x12, 0x42, 0x8e, 0x14, 0xfa, 0x50, 0x81, 0x78, 0x48, 0x68, 0x29, 0x48, 0x08, 0x21, 0xab, 0xbc,
	0xf0, 0x82, 0x1c, 0x7b, 0x53, 0x59, 0xd8, 0x5e, 0x63, 0xaf, 0xab, 0xe6, 0x10, 0x9c, 0x8a, 0x03,
	0x70, 0x10, 0x2e, 0xc1, 0x7a, 0x6d, 0xa7, 0x4e, 0xec, 0x36, 0x55, 0xc4, 0xdb, 0xce, 0xec, 0xcc,
	0xdf, 0xf3, 0xff, 0x79, 0x77, 0xe1, 0x30, 0x8a, 0x29, 0xa3, 0x23, 0x9f, 0x26, 0xec, 0xfb, 0x9c,
	0xa6, 0xa1, 0x5b, 0x59, 0x9a, 0x62, 0x0f, 0xf1, 0x2a, 0x48, 0xaf, 0xcc, 0x4a, 0xfa, 0xf2, 0x48,
	0x7f, 0x74, 0x41, 0xe9, 0x85, 0x4f, 0x46, 0xa2, 0x62, 0x96, 0xce, 0x47, 0x24, 0x88, 0xd8, 0x22,
	0x6f, 0xd0, 0x0f, 0xd6, 0x37, 0x99, 0x17, 0x90, 0x84, 0xd9, 0x41, 0x94, 0x17, 0x18, 0x53, 0xd8,
	0x3d, 0x23, 0x6c, 0x1a, 0x7b, 0x64, 0x9e, 0x58, 0xe4, 0x27, 0x8e, 0xa1, 0x3d, 0x23, 0x73, 0x1a,
	0x13, 0x4d, 0x1a, 0x4a, 0xcf, 0x7b, 0x63, 0xdd, 0xcc, 0x15, 0xcc, 0x52, 0xc1, 0x3c, 0x2f, 0x15,
	0xac, 0xa2, 0xd2, 0xf8, 0x0c, 0xfd, 0x8a, 0x46, 0x12, 0xe1, 0x5b, 0x2e, 0x22, 0x22, 0x2e, 0xa2,
	0x70, 0x91, 0x27, 0x66, 0x7d, 0x6e, 0xf3, 0x13, 0x8f, 0x26, 0xa1, 0xfb, 0x3e, 0x8b, 0x45, 0xaf,
	0x55, 0x34, 0x19, 0x7f, 0x24, 0xb8, 0x5f, 0xdb, 0xc5, 0x3d, 0x90, 0x3d, 0x57, 0x4c, 0xd5, 0xb5,
	0xf8, 0x0a, 0x07, 0xa0, 0xa4, 0x3c, 0x21, 0x8b, 0x44, 0xb6, 0xc4, 0x63, 0x50, 0xd9, 0x22, 0x22,
	0x9a, 0xc2, 0x53, 0x7b, 0xe3, 0xc7, 0x9b, 0x3e, 0x7a, 0xce, 0x6b, 0x2d, 0xd1, 0x81, 0x08, 0x6a,
	0x68, 0x07, 0x44, 0x53, 0x85, 0x98, 0x58, 0xa3, 0xc9, 0xd5, 0xb8, 0x55, 0xad, 0xb5, 0x91, 0x83,
	0xa8, 0x43, 0x1d, 0x3a, 0x3e, 0x75, 0x6c, 0xe6, 0xd1, 0x50, 0x6b, 0x0b, 0x9d, 0x65, 0x6c, 0xec,
	0x0b, 0xca, 0x27, 0x84, 0xd9, 0x9e, 0x9f, 0x51, 0x5e, 0xf3, 0x62, 0xfc, 0x52, 0x00, 0xab, 0xa3,
	0xe5, 0x95, 0xa5, 0x45, 0xa9, 0x6e, 0x51, 0xde, 0xda, 0xa2, 0xd2, 0x60, 0x51, 0xdd, 0xc2, 0x62,
	0x6b, 0xd5, 0x22, 0x0e, 0xa1, 0xe7, 0x92, 0xc4, 0x89, 0xbd, 0xa8, 0x42, 0xa0, 0x9a, 0xc2, 0x2f,
	0xd0, 0x71, 0x68, 0xc8, 0x6c, 0x87, 0x25, 0xda, 0x8e, 0x38, 0x17, 0x47, 0x9b, 0xe6, 0xcf, 0x39,
	0x98, 0xef, 0x8a, 0xb6, 0xd3, 0x90, 0xc5, 0x0b, 0x6b, 0xa9, 0x92, 0xcd, 0x13, 0x79, 0x0e, 0x4b,
	0x63, 0x92, 0x68, 0x1d, 0xae, 0xc8, 0xe7, 0x29, 0x63, 0xfd, 0x0d, 0xf4, 0x57, 0xda, 0x32, 0x98,
	0x3f, 0xc8, 0xa2, 0x84, 0xc9, 0x97, 0xf8, 0x00, 0x5a, 0x97, 0xb6, 0x9f, 0x92, 0xe2, 0x0c, 0xe5,
	0xc1, 0x6b, 0xf9, 0x58, 0x32, 0xfe, 0xca, 0x00, 0x13, 0xd7, 0xfd, 0xc8, 0x48, 0x90, 0xfd, 0xae,
	0x92, 0xba, 0xb4, 0x35, 0x75, 0xb9, 0x81, 0xba, 0xb2, 0x05, 0x75, 0xf5, 0x76, 0xea, 0xad, 0x3a,
	0xf5, 0x0f, 0x15, 0xea, 0x6d, 0x41, 0xfd, 0x65, 0xd3, 0xfc, 0xd7, 0x6e, 0xef, 0x44, 0x7b, 0xe7,
	0x7f, 0xd2, 0x3e, 0x80, 0xfe, 0x09, 0xf1, 0x09, 0x23, 0x25, 0xef, 0xb5, 0xeb, 0xf1, 0xe2, 0x19,
	0x0c, 0xd6, 0xf9, 0x62, 0x07, 0xd4, 0x2c, 0x37, 0xb8, 0x87, 0x5d, 0x68, 0x89, 0xf4, 0x40, 0x1a,
	0xff, 0x96, 0x61, 0xb7, 0x5a, 0xc9, 0xcf, 0x5c, 0x77, 0xf9, 0x34, 0xe1, 0xb0, 0xc9, 0x78, 0xf5,
	0xf5, 0xd3, 0x0f, 0x37, 0x54, 0xf0, 0xb7, 0xed, 0xab, 0x50, 0x2c, 0x2e, 0xe8, 0x4d, 0x8a, 0xcb,
	0x9b, 0xae, 0x3f, 0xbd, 0xdb, 0x11, 0xc7, 0x09, 0xec, 0x14, 0xbf, 0x00, 0xf7, 0x6f, 0xff, 0x3f,
	0xfa, 0xc3, 0xda, 0x89, 0x39, 0xcd, 0x5e, 0x7c, 0x3c, 0x03, 0xb8, 0xc6, 0x88, 0x8d, 0x56, 0x56,
	0x30, 0xdf, 0x24, 0x34, 0xed, 0x7d, 0xeb, 0x66, 0x6d, 0xa2, 0x6b, 0xd6, 0x16, 0x9b, 0xaf, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xef, 0x9a, 0x51, 0x9e, 0x06, 0x00, 0x00,
}
