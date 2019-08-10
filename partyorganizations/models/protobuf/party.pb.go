// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/party.proto

package models

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

type Order_Status int32

const (
	Order_CANCELED  Order_Status = 0
	Order_CONFIRMED Order_Status = 1
	Order_ON_HOLD   Order_Status = 2
	Order_DONE      Order_Status = 3
)

var Order_Status_name = map[int32]string{
	0: "CANCELED",
	1: "CONFIRMED",
	2: "ON_HOLD",
	3: "DONE",
}

var Order_Status_value = map[string]int32{
	"CANCELED":  0,
	"CONFIRMED": 1,
	"ON_HOLD":   2,
	"DONE":      3,
}

func (x Order_Status) String() string {
	return proto.EnumName(Order_Status_name, int32(x))
}

func (Order_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{2, 0}
}

type Alcohol_Alcohol int32

const (
	Alcohol_WHISKEY Alcohol_Alcohol = 0
	Alcohol_VODKA   Alcohol_Alcohol = 1
	Alcohol_VINO    Alcohol_Alcohol = 2
)

var Alcohol_Alcohol_name = map[int32]string{
	0: "WHISKEY",
	1: "VODKA",
	2: "VINO",
}

var Alcohol_Alcohol_value = map[string]int32{
	"WHISKEY": 0,
	"VODKA":   1,
	"VINO":    2,
}

func (x Alcohol_Alcohol) String() string {
	return proto.EnumName(Alcohol_Alcohol_name, int32(x))
}

func (Alcohol_Alcohol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{4, 0}
}

type Snack_Snacks int32

const (
	Snack_CHIPS  Snack_Snacks = 0
	Snack_CRISPS Snack_Snacks = 1
)

var Snack_Snacks_name = map[int32]string{
	0: "CHIPS",
	1: "CRISPS",
}

var Snack_Snacks_value = map[string]int32{
	"CHIPS":  0,
	"CRISPS": 1,
}

func (x Snack_Snacks) String() string {
	return proto.EnumName(Snack_Snacks_name, int32(x))
}

func (Snack_Snacks) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{5, 0}
}

type Party struct {
	Id                   int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 int64     `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Location             *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Order                *Order    `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
	AmountOfHostess      int32     `protobuf:"varint,5,opt,name=amountOfHostess,proto3" json:"amountOfHostess,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Party) Reset()         { *m = Party{} }
func (m *Party) String() string { return proto.CompactTextString(m) }
func (*Party) ProtoMessage()    {}
func (*Party) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{0}
}

func (m *Party) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Party.Unmarshal(m, b)
}
func (m *Party) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Party.Marshal(b, m, deterministic)
}
func (m *Party) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Party.Merge(m, src)
}
func (m *Party) XXX_Size() int {
	return xxx_messageInfo_Party.Size(m)
}
func (m *Party) XXX_DiscardUnknown() {
	xxx_messageInfo_Party.DiscardUnknown(m)
}

var xxx_messageInfo_Party proto.InternalMessageInfo

func (m *Party) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Party) GetUser() int64 {
	if m != nil {
		return m.User
	}
	return 0
}

func (m *Party) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Party) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Party) GetAmountOfHostess() int32 {
	if m != nil {
		return m.AmountOfHostess
	}
	return 0
}

type Location struct {
	City                 string   `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Street               string   `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
	HouseNumber          string   `protobuf:"bytes,3,opt,name=houseNumber,proto3" json:"houseNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Location) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Location) GetHouseNumber() string {
	if m != nil {
		return m.HouseNumber
	}
	return ""
}

type Order struct {
	// id from user's dataset
	Deliver              int32        `protobuf:"varint,1,opt,name=deliver,proto3" json:"deliver,omitempty"`
	Status               Order_Status `protobuf:"varint,2,opt,name=status,proto3,enum=tutorial.Order_Status" json:"status,omitempty"`
	Content              *Content     `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{2}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetDeliver() int32 {
	if m != nil {
		return m.Deliver
	}
	return 0
}

func (m *Order) GetStatus() Order_Status {
	if m != nil {
		return m.Status
	}
	return Order_CANCELED
}

func (m *Order) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

type Content struct {
	Alcohols             []*Alcohol `protobuf:"bytes,1,rep,name=alcohols,proto3" json:"alcohols,omitempty"`
	Snacks               []*Snack   `protobuf:"bytes,2,rep,name=snacks,proto3" json:"snacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{3}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetAlcohols() []*Alcohol {
	if m != nil {
		return m.Alcohols
	}
	return nil
}

func (m *Content) GetSnacks() []*Snack {
	if m != nil {
		return m.Snacks
	}
	return nil
}

type Alcohol struct {
	Alcohol              Alcohol_Alcohol `protobuf:"varint,1,opt,name=alcohol,proto3,enum=tutorial.Alcohol_Alcohol" json:"alcohol,omitempty"`
	Amount               int32           `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Alcohol) Reset()         { *m = Alcohol{} }
func (m *Alcohol) String() string { return proto.CompactTextString(m) }
func (*Alcohol) ProtoMessage()    {}
func (*Alcohol) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{4}
}

func (m *Alcohol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alcohol.Unmarshal(m, b)
}
func (m *Alcohol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alcohol.Marshal(b, m, deterministic)
}
func (m *Alcohol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alcohol.Merge(m, src)
}
func (m *Alcohol) XXX_Size() int {
	return xxx_messageInfo_Alcohol.Size(m)
}
func (m *Alcohol) XXX_DiscardUnknown() {
	xxx_messageInfo_Alcohol.DiscardUnknown(m)
}

var xxx_messageInfo_Alcohol proto.InternalMessageInfo

func (m *Alcohol) GetAlcohol() Alcohol_Alcohol {
	if m != nil {
		return m.Alcohol
	}
	return Alcohol_WHISKEY
}

func (m *Alcohol) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Snack struct {
	Snack                Snack_Snacks `protobuf:"varint,1,opt,name=snack,proto3,enum=tutorial.Snack_Snacks" json:"snack,omitempty"`
	Amount               int32        `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Snack) Reset()         { *m = Snack{} }
func (m *Snack) String() string { return proto.CompactTextString(m) }
func (*Snack) ProtoMessage()    {}
func (*Snack) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3cd75d7e545703e, []int{5}
}

func (m *Snack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snack.Unmarshal(m, b)
}
func (m *Snack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snack.Marshal(b, m, deterministic)
}
func (m *Snack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snack.Merge(m, src)
}
func (m *Snack) XXX_Size() int {
	return xxx_messageInfo_Snack.Size(m)
}
func (m *Snack) XXX_DiscardUnknown() {
	xxx_messageInfo_Snack.DiscardUnknown(m)
}

var xxx_messageInfo_Snack proto.InternalMessageInfo

func (m *Snack) GetSnack() Snack_Snacks {
	if m != nil {
		return m.Snack
	}
	return Snack_CHIPS
}

func (m *Snack) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterEnum("tutorial.Order_Status", Order_Status_name, Order_Status_value)
	proto.RegisterEnum("tutorial.Alcohol_Alcohol", Alcohol_Alcohol_name, Alcohol_Alcohol_value)
	proto.RegisterEnum("tutorial.Snack_Snacks", Snack_Snacks_name, Snack_Snacks_value)
	proto.RegisterType((*Party)(nil), "tutorial.Party")
	proto.RegisterType((*Location)(nil), "tutorial.Location")
	proto.RegisterType((*Order)(nil), "tutorial.Order")
	proto.RegisterType((*Content)(nil), "tutorial.Content")
	proto.RegisterType((*Alcohol)(nil), "tutorial.Alcohol")
	proto.RegisterType((*Snack)(nil), "tutorial.Snack")
}

func init() { proto.RegisterFile("protobuf/party.proto", fileDescriptor_b3cd75d7e545703e) }

var fileDescriptor_b3cd75d7e545703e = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0xed, 0xb4, 0xcd, 0x9f, 0xde, 0xfe, 0x7e, 0xdd, 0x78, 0x91, 0x12, 0x9f, 0x2c, 0x01, 0xb1,
	0xb0, 0x1a, 0xa1, 0xfb, 0xea, 0x4b, 0x49, 0x2b, 0x0d, 0x5b, 0x93, 0x32, 0x81, 0x55, 0x9f, 0x24,
	0x4d, 0x67, 0xd9, 0x60, 0x36, 0xb3, 0x64, 0x26, 0xc2, 0x3e, 0xfb, 0x75, 0xfc, 0x08, 0x7e, 0x38,
	0xc9, 0x64, 0xd2, 0xba, 0x15, 0x5f, 0xc2, 0xdc, 0x7b, 0xce, 0x9c, 0x73, 0xe6, 0x5e, 0x02, 0xcf,
	0x1f, 0x2a, 0x2e, 0xf9, 0xbe, 0xbe, 0x7d, 0xf7, 0x90, 0x56, 0xf2, 0xd1, 0x57, 0x25, 0xda, 0xb2,
	0x96, 0xbc, 0xca, 0xd3, 0xc2, 0xfb, 0x49, 0xc0, 0xd8, 0x35, 0x08, 0x4e, 0xa0, 0x9f, 0x1f, 0x5c,
	0x32, 0x23, 0x73, 0x83, 0xf6, 0xf3, 0x03, 0x22, 0x0c, 0x6b, 0xc1, 0x2a, 0xb7, 0x3f, 0x23, 0xf3,
	0x01, 0x55, 0x67, 0xf4, 0xc1, 0x2e, 0x78, 0x96, 0xca, 0x9c, 0x97, 0xee, 0x60, 0x46, 0xe6, 0xe3,
	0x05, 0xfa, 0x9d, 0x94, 0xbf, 0xd5, 0x08, 0x3d, 0x72, 0xf0, 0x15, 0x18, 0xbc, 0x3a, 0xb0, 0xca,
	0x1d, 0x2a, 0xf2, 0xc5, 0x89, 0x1c, 0x37, 0x6d, 0xda, 0xa2, 0x38, 0x87, 0x8b, 0xf4, 0x9e, 0xd7,
	0xa5, 0x8c, 0x6f, 0x37, 0x5c, 0x48, 0x26, 0x84, 0x6b, 0xa8, 0x1c, 0xe7, 0x6d, 0xef, 0x33, 0xd8,
	0x9d, 0x4d, 0x13, 0x30, 0xcb, 0xe5, 0xa3, 0x8a, 0x3c, 0xa2, 0xea, 0x8c, 0x53, 0x30, 0x85, 0xac,
	0x18, 0x93, 0x2a, 0xf6, 0x88, 0xea, 0x0a, 0x67, 0x30, 0xbe, 0xe3, 0xb5, 0x60, 0x51, 0x7d, 0xbf,
	0x67, 0x95, 0xca, 0x3e, 0xa2, 0x7f, 0xb6, 0xbc, 0x5f, 0x04, 0x0c, 0x15, 0x0a, 0x5d, 0xb0, 0x0e,
	0xac, 0xc8, 0xbf, 0xb3, 0x4a, 0x4f, 0xa3, 0x2b, 0xd1, 0x6f, 0xd4, 0x53, 0x59, 0x0b, 0xa5, 0x3e,
	0x59, 0x4c, 0xcf, 0xde, 0xe3, 0x27, 0x0a, 0xa5, 0x9a, 0x85, 0x97, 0x60, 0x65, 0xbc, 0x94, 0xac,
	0x94, 0x7a, 0x5a, 0xcf, 0x4e, 0x17, 0x82, 0x16, 0xa0, 0x1d, 0xc3, 0x7b, 0x0f, 0x66, 0x7b, 0x1d,
	0xff, 0x03, 0x3b, 0x58, 0x46, 0xc1, 0x7a, 0xbb, 0x5e, 0x39, 0x3d, 0xfc, 0x1f, 0x46, 0x41, 0x1c,
	0x7d, 0x08, 0xe9, 0xc7, 0xf5, 0xca, 0x21, 0x38, 0x06, 0x2b, 0x8e, 0xbe, 0x6e, 0xe2, 0xed, 0xca,
	0xe9, 0xa3, 0x0d, 0xc3, 0x55, 0x1c, 0xad, 0x9d, 0x81, 0x97, 0x82, 0xa5, 0x15, 0xf1, 0x2d, 0xd8,
	0x69, 0x91, 0xf1, 0x3b, 0x5e, 0x08, 0x97, 0xcc, 0x06, 0x4f, 0x6d, 0x97, 0x2d, 0x42, 0x8f, 0x14,
	0x7c, 0x0d, 0xa6, 0x28, 0xd3, 0xec, 0x5b, 0xf3, 0xa8, 0xc1, 0xd3, 0x25, 0x25, 0x4d, 0x9f, 0x6a,
	0xd8, 0xfb, 0x41, 0xc0, 0xd2, 0xd7, 0xf1, 0x0a, 0x2c, 0x2d, 0xa0, 0x66, 0x34, 0x59, 0xbc, 0xf8,
	0xcb, 0xe2, 0x68, 0xd5, 0x31, 0x9b, 0xe5, 0xb4, 0xfb, 0x54, 0xe3, 0x33, 0xa8, 0xae, 0xbc, 0xcb,
	0x93, 0xee, 0x18, 0xac, 0x4f, 0x9b, 0x30, 0xb9, 0x5e, 0x7f, 0x71, 0x7a, 0x38, 0x02, 0xe3, 0x26,
	0x5e, 0x5d, 0x2f, 0x1d, 0xd2, 0x3c, 0xf4, 0x26, 0x8c, 0x62, 0xa7, 0xef, 0x95, 0x60, 0xa8, 0x58,
	0xf8, 0x06, 0x0c, 0x15, 0x4c, 0x07, 0x98, 0x9e, 0xc5, 0x6e, 0xbf, 0x82, 0xb6, 0xa4, 0x7f, 0x7a,
	0xbf, 0x04, 0xb3, 0x25, 0x36, 0x6e, 0xc1, 0x26, 0xdc, 0x25, 0x4e, 0x0f, 0x01, 0xcc, 0x80, 0x86,
	0xc9, 0x2e, 0x71, 0xc8, 0xde, 0x54, 0x7f, 0xcc, 0xd5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06,
	0x98, 0x5e, 0xaf, 0x49, 0x03, 0x00, 0x00,
}
