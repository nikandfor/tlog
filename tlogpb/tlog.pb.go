// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: tlogpb/tlog.proto

package tlogpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels     *Labels     `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels,omitempty"`
	Location   *Location   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Message    *Message    `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	SpanStart  *SpanStart  `protobuf:"bytes,4,opt,name=span_start,json=spanStart,proto3" json:"span_start,omitempty"`
	SpanFinish *SpanFinish `protobuf:"bytes,5,opt,name=span_finish,json=spanFinish,proto3" json:"span_finish,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlogpb_tlog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_tlogpb_tlog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_tlogpb_tlog_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetLabels() *Labels {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Record) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Record) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Record) GetSpanStart() *SpanStart {
	if x != nil {
		return x.SpanStart
	}
	return nil
}

func (x *Record) GetSpanFinish() *SpanFinish {
	if x != nil {
		return x.SpanFinish
	}
	return nil
}

type Labels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Span   []byte   `protobuf:"bytes,1,opt,name=span,proto3" json:"span,omitempty"`
	Labels []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Labels) Reset() {
	*x = Labels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlogpb_tlog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labels) ProtoMessage() {}

func (x *Labels) ProtoReflect() protoreflect.Message {
	mi := &file_tlogpb_tlog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labels.ProtoReflect.Descriptor instead.
func (*Labels) Descriptor() ([]byte, []int) {
	return file_tlogpb_tlog_proto_rawDescGZIP(), []int{1}
}

func (x *Labels) GetSpan() []byte {
	if x != nil {
		return x.Span
	}
	return nil
}

func (x *Labels) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pc    int64  `protobuf:"varint,1,opt,name=pc,proto3" json:"pc,omitempty"`
	Entry int64  `protobuf:"varint,2,opt,name=entry,proto3" json:"entry,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	File  string `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Line  int32  `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlogpb_tlog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_tlogpb_tlog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_tlogpb_tlog_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetPc() int64 {
	if x != nil {
		return x.Pc
	}
	return 0
}

func (x *Location) GetEntry() int64 {
	if x != nil {
		return x.Entry
	}
	return 0
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Location) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Span     []byte `protobuf:"bytes,1,opt,name=span,proto3" json:"span,omitempty"`
	Location int64  `protobuf:"varint,2,opt,name=location,proto3" json:"location,omitempty"`
	Time     int64  `protobuf:"fixed64,3,opt,name=time,proto3" json:"time,omitempty"`
	Text     string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlogpb_tlog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_tlogpb_tlog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_tlogpb_tlog_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetSpan() []byte {
	if x != nil {
		return x.Span
	}
	return nil
}

func (x *Message) GetLocation() int64 {
	if x != nil {
		return x.Location
	}
	return 0
}

func (x *Message) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SpanStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent   []byte `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Location int64  `protobuf:"varint,3,opt,name=location,proto3" json:"location,omitempty"`
	Started  int64  `protobuf:"fixed64,4,opt,name=started,proto3" json:"started,omitempty"`
}

func (x *SpanStart) Reset() {
	*x = SpanStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlogpb_tlog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpanStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpanStart) ProtoMessage() {}

func (x *SpanStart) ProtoReflect() protoreflect.Message {
	mi := &file_tlogpb_tlog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpanStart.ProtoReflect.Descriptor instead.
func (*SpanStart) Descriptor() ([]byte, []int) {
	return file_tlogpb_tlog_proto_rawDescGZIP(), []int{4}
}

func (x *SpanStart) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SpanStart) GetParent() []byte {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *SpanStart) GetLocation() int64 {
	if x != nil {
		return x.Location
	}
	return 0
}

func (x *SpanStart) GetStarted() int64 {
	if x != nil {
		return x.Started
	}
	return 0
}

type SpanFinish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Elapsed int64  `protobuf:"varint,2,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
}

func (x *SpanFinish) Reset() {
	*x = SpanFinish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlogpb_tlog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpanFinish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpanFinish) ProtoMessage() {}

func (x *SpanFinish) ProtoReflect() protoreflect.Message {
	mi := &file_tlogpb_tlog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpanFinish.ProtoReflect.Descriptor instead.
func (*SpanFinish) Descriptor() ([]byte, []int) {
	return file_tlogpb_tlog_proto_rawDescGZIP(), []int{5}
}

func (x *SpanFinish) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SpanFinish) GetElapsed() int64 {
	if x != nil {
		return x.Elapsed
	}
	return 0
}

var File_tlogpb_tlog_proto protoreflect.FileDescriptor

var file_tlogpb_tlog_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2f, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x22, 0xf0, 0x01, 0x0a, 0x06,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2c,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x74, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x70, 0x61, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x6c,
	0x6f, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x09,
	0x73, 0x70, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x73, 0x70, 0x61,
	0x6e, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x52, 0x0a, 0x73, 0x70, 0x61, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x22, 0x34,
	0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x61, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x70, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x22, 0x6c, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x70, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x70, 0x63,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x65, 0x22, 0x61, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x70, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x70, 0x61,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x10, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x69, 0x0a, 0x09, 0x53, 0x70, 0x61, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x10, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x22, 0x36, 0x0a, 0x0a, 0x53, 0x70, 0x61, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x74, 0x6c,
	0x6f, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tlogpb_tlog_proto_rawDescOnce sync.Once
	file_tlogpb_tlog_proto_rawDescData = file_tlogpb_tlog_proto_rawDesc
)

func file_tlogpb_tlog_proto_rawDescGZIP() []byte {
	file_tlogpb_tlog_proto_rawDescOnce.Do(func() {
		file_tlogpb_tlog_proto_rawDescData = protoimpl.X.CompressGZIP(file_tlogpb_tlog_proto_rawDescData)
	})
	return file_tlogpb_tlog_proto_rawDescData
}

var file_tlogpb_tlog_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tlogpb_tlog_proto_goTypes = []interface{}{
	(*Record)(nil),     // 0: tlogpb.Record
	(*Labels)(nil),     // 1: tlogpb.Labels
	(*Location)(nil),   // 2: tlogpb.Location
	(*Message)(nil),    // 3: tlogpb.Message
	(*SpanStart)(nil),  // 4: tlogpb.SpanStart
	(*SpanFinish)(nil), // 5: tlogpb.SpanFinish
}
var file_tlogpb_tlog_proto_depIdxs = []int32{
	1, // 0: tlogpb.Record.labels:type_name -> tlogpb.Labels
	2, // 1: tlogpb.Record.location:type_name -> tlogpb.Location
	3, // 2: tlogpb.Record.message:type_name -> tlogpb.Message
	4, // 3: tlogpb.Record.span_start:type_name -> tlogpb.SpanStart
	5, // 4: tlogpb.Record.span_finish:type_name -> tlogpb.SpanFinish
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tlogpb_tlog_proto_init() }
func file_tlogpb_tlog_proto_init() {
	if File_tlogpb_tlog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tlogpb_tlog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tlogpb_tlog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Labels); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tlogpb_tlog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tlogpb_tlog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tlogpb_tlog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpanStart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tlogpb_tlog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpanFinish); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tlogpb_tlog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tlogpb_tlog_proto_goTypes,
		DependencyIndexes: file_tlogpb_tlog_proto_depIdxs,
		MessageInfos:      file_tlogpb_tlog_proto_msgTypes,
	}.Build()
	File_tlogpb_tlog_proto = out.File
	file_tlogpb_tlog_proto_rawDesc = nil
	file_tlogpb_tlog_proto_goTypes = nil
	file_tlogpb_tlog_proto_depIdxs = nil
}
