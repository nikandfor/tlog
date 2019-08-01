// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tlogpb/tlog.proto

package tlogpb

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

type Record struct {
	Labels               []string    `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Location             *Location   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Msg                  *Message    `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	SpanStart            *SpanStart  `protobuf:"bytes,4,opt,name=span_start,json=spanStart,proto3" json:"span_start,omitempty"`
	SpanFinish           *SpanFinish `protobuf:"bytes,5,opt,name=span_finish,json=spanFinish,proto3" json:"span_finish,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Record) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Record) GetMsg() *Message {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Record) GetSpanStart() *SpanStart {
	if m != nil {
		return m.SpanStart
	}
	return nil
}

func (m *Record) GetSpanFinish() *SpanFinish {
	if m != nil {
		return m.SpanFinish
	}
	return nil
}

type Location struct {
	Pc                   int64    `protobuf:"varint,1,opt,name=pc,proto3" json:"pc,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	File                 string   `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Line                 int32    `protobuf:"varint,4,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{1}
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

func (m *Location) GetPc() int64 {
	if m != nil {
		return m.Pc
	}
	return 0
}

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Location) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

type Message struct {
	Span                 int64    `protobuf:"varint,1,opt,name=span,proto3" json:"span,omitempty"`
	Location             int64    `protobuf:"varint,2,opt,name=location,proto3" json:"location,omitempty"`
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSpan() int64 {
	if m != nil {
		return m.Span
	}
	return 0
}

func (m *Message) GetLocation() int64 {
	if m != nil {
		return m.Location
	}
	return 0
}

func (m *Message) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SpanStart struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               int64    `protobuf:"varint,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Location             int64    `protobuf:"varint,3,opt,name=location,proto3" json:"location,omitempty"`
	Started              int64    `protobuf:"varint,4,opt,name=started,proto3" json:"started,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanStart) Reset()         { *m = SpanStart{} }
func (m *SpanStart) String() string { return proto.CompactTextString(m) }
func (*SpanStart) ProtoMessage()    {}
func (*SpanStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{3}
}

func (m *SpanStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanStart.Unmarshal(m, b)
}
func (m *SpanStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanStart.Marshal(b, m, deterministic)
}
func (m *SpanStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanStart.Merge(m, src)
}
func (m *SpanStart) XXX_Size() int {
	return xxx_messageInfo_SpanStart.Size(m)
}
func (m *SpanStart) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanStart.DiscardUnknown(m)
}

var xxx_messageInfo_SpanStart proto.InternalMessageInfo

func (m *SpanStart) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpanStart) GetParent() int64 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *SpanStart) GetLocation() int64 {
	if m != nil {
		return m.Location
	}
	return 0
}

func (m *SpanStart) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

type SpanFinish struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Elapsed              int64    `protobuf:"varint,2,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	Flags                int64    `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanFinish) Reset()         { *m = SpanFinish{} }
func (m *SpanFinish) String() string { return proto.CompactTextString(m) }
func (*SpanFinish) ProtoMessage()    {}
func (*SpanFinish) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec43dc181f2a6e80, []int{4}
}

func (m *SpanFinish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanFinish.Unmarshal(m, b)
}
func (m *SpanFinish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanFinish.Marshal(b, m, deterministic)
}
func (m *SpanFinish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanFinish.Merge(m, src)
}
func (m *SpanFinish) XXX_Size() int {
	return xxx_messageInfo_SpanFinish.Size(m)
}
func (m *SpanFinish) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanFinish.DiscardUnknown(m)
}

var xxx_messageInfo_SpanFinish proto.InternalMessageInfo

func (m *SpanFinish) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpanFinish) GetElapsed() int64 {
	if m != nil {
		return m.Elapsed
	}
	return 0
}

func (m *SpanFinish) GetFlags() int64 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "tlogpb.Record")
	proto.RegisterType((*Location)(nil), "tlogpb.Location")
	proto.RegisterType((*Message)(nil), "tlogpb.Message")
	proto.RegisterType((*SpanStart)(nil), "tlogpb.SpanStart")
	proto.RegisterType((*SpanFinish)(nil), "tlogpb.SpanFinish")
}

func init() { proto.RegisterFile("tlogpb/tlog.proto", fileDescriptor_ec43dc181f2a6e80) }

var fileDescriptor_ec43dc181f2a6e80 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0xcd, 0x52, 0x58, 0xd8, 0x21, 0x51, 0x69, 0x8c, 0x69, 0x3c, 0xe1, 0x9e, 0x38, 0x18, 0x34,
	0xf2, 0x0f, 0x9e, 0xf0, 0x32, 0x24, 0x5e, 0x4d, 0xd9, 0x2d, 0x6b, 0x93, 0xb2, 0xdb, 0x6c, 0x7b,
	0xf0, 0x53, 0xfd, 0x1c, 0xd3, 0xd9, 0x16, 0x11, 0x4f, 0xbc, 0x99, 0x79, 0x33, 0xef, 0xf5, 0xb1,
	0xb0, 0xf0, 0xa6, 0x6b, 0xec, 0xfe, 0x29, 0xfc, 0xac, 0x6d, 0xdf, 0xf9, 0x8e, 0xe7, 0x43, 0xab,
	0xfc, 0xce, 0x20, 0x47, 0x55, 0x75, 0x7d, 0xcd, 0xef, 0x20, 0x37, 0x72, 0xaf, 0x8c, 0x13, 0xd9,
	0x92, 0xad, 0x0a, 0x8c, 0x15, 0x7f, 0x84, 0x99, 0xe9, 0x2a, 0xe9, 0x75, 0xd7, 0x8a, 0xd1, 0x32,
	0x5b, 0xcd, 0x5f, 0x6e, 0xd6, 0xc3, 0xf6, 0x7a, 0x1b, 0xfb, 0x78, 0x62, 0xf0, 0x07, 0x60, 0x47,
	0xd7, 0x08, 0x46, 0xc4, 0xeb, 0x44, 0x7c, 0x53, 0xce, 0xc9, 0x46, 0x61, 0x98, 0xf1, 0x67, 0x00,
	0x67, 0x65, 0xfb, 0xe1, 0xbc, 0xec, 0xbd, 0x18, 0x13, 0x73, 0x91, 0x98, 0x3b, 0x2b, 0xdb, 0x5d,
	0x18, 0x60, 0xe1, 0x12, 0xe4, 0x1b, 0x98, 0xd3, 0xc6, 0x41, 0xb7, 0xda, 0x7d, 0x8a, 0x09, 0xad,
	0xf0, 0xf3, 0x95, 0x57, 0x9a, 0x20, 0x1d, 0x1e, 0x70, 0xf9, 0x0e, 0xb3, 0xe4, 0x8f, 0x5f, 0xc1,
	0xc8, 0x56, 0x22, 0x5b, 0x66, 0x2b, 0x86, 0x23, 0x5b, 0x71, 0x0e, 0xe3, 0x56, 0x1e, 0x15, 0xbd,
	0xa7, 0x40, 0xc2, 0xa1, 0x77, 0xd0, 0x46, 0x91, 0xf5, 0x02, 0x09, 0x87, 0x9e, 0xd1, 0xad, 0x22,
	0x93, 0x13, 0x24, 0x5c, 0x4a, 0x98, 0xc6, 0xe7, 0x84, 0x71, 0x10, 0x8c, 0x87, 0x09, 0xf3, 0xfb,
	0x8b, 0xb8, 0xd8, 0x59, 0x38, 0x1c, 0xc6, 0x5e, 0x1f, 0x07, 0x09, 0x86, 0x84, 0xa9, 0xa7, 0xbe,
	0x86, 0x1c, 0x0a, 0x24, 0x5c, 0x6a, 0x28, 0x4e, 0x39, 0x04, 0xef, 0xba, 0x4e, 0xde, 0x35, 0xfd,
	0x4f, 0x56, 0xf6, 0xaa, 0xf5, 0xf1, 0x7c, 0xac, 0xfe, 0x08, 0xb3, 0x0b, 0x61, 0x01, 0x53, 0x4a,
	0x5b, 0xd5, 0xa4, 0xc3, 0x30, 0x95, 0xe5, 0x16, 0xe0, 0x37, 0xbf, 0x7f, 0x5a, 0x02, 0xa6, 0xca,
	0x48, 0xeb, 0x54, 0x1d, 0xc5, 0x52, 0xc9, 0x6f, 0x61, 0x72, 0x30, 0xb2, 0x71, 0x51, 0x6a, 0x28,
	0xf6, 0x39, 0x7d, 0x5d, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x16, 0x67, 0xf6, 0x72,
	0x02, 0x00, 0x00,
}