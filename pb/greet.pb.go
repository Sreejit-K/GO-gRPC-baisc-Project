// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/greet.proto

package pb

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

type NoParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoParams) Reset()         { *m = NoParams{} }
func (m *NoParams) String() string { return proto.CompactTextString(m) }
func (*NoParams) ProtoMessage()    {}
func (*NoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{0}
}

func (m *NoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoParams.Unmarshal(m, b)
}
func (m *NoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoParams.Marshal(b, m, deterministic)
}
func (m *NoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoParams.Merge(m, src)
}
func (m *NoParams) XXX_Size() int {
	return xxx_messageInfo_NoParams.Size(m)
}
func (m *NoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_NoParams.DiscardUnknown(m)
}

var xxx_messageInfo_NoParams proto.InternalMessageInfo

type ReqList struct {
	Requests             []string `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqList) Reset()         { *m = ReqList{} }
func (m *ReqList) String() string { return proto.CompactTextString(m) }
func (*ReqList) ProtoMessage()    {}
func (*ReqList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{1}
}

func (m *ReqList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqList.Unmarshal(m, b)
}
func (m *ReqList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqList.Marshal(b, m, deterministic)
}
func (m *ReqList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqList.Merge(m, src)
}
func (m *ReqList) XXX_Size() int {
	return xxx_messageInfo_ReqList.Size(m)
}
func (m *ReqList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqList proto.InternalMessageInfo

func (m *ReqList) GetRequests() []string {
	if m != nil {
		return m.Requests
	}
	return nil
}

type ResList struct {
	Responses            []string `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResList) Reset()         { *m = ResList{} }
func (m *ResList) String() string { return proto.CompactTextString(m) }
func (*ResList) ProtoMessage()    {}
func (*ResList) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{2}
}

func (m *ResList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResList.Unmarshal(m, b)
}
func (m *ResList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResList.Marshal(b, m, deterministic)
}
func (m *ResList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResList.Merge(m, src)
}
func (m *ResList) XXX_Size() int {
	return xxx_messageInfo_ResList.Size(m)
}
func (m *ResList) XXX_DiscardUnknown() {
	xxx_messageInfo_ResList.DiscardUnknown(m)
}

var xxx_messageInfo_ResList proto.InternalMessageInfo

func (m *ResList) GetResponses() []string {
	if m != nil {
		return m.Responses
	}
	return nil
}

type GreetingRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingRequest) Reset()         { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()    {}
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{3}
}

func (m *GreetingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingRequest.Unmarshal(m, b)
}
func (m *GreetingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingRequest.Marshal(b, m, deterministic)
}
func (m *GreetingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingRequest.Merge(m, src)
}
func (m *GreetingRequest) XXX_Size() int {
	return xxx_messageInfo_GreetingRequest.Size(m)
}
func (m *GreetingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingRequest proto.InternalMessageInfo

func (m *GreetingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GreetingResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingResponse) Reset()         { *m = GreetingResponse{} }
func (m *GreetingResponse) String() string { return proto.CompactTextString(m) }
func (*GreetingResponse) ProtoMessage()    {}
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{4}
}

func (m *GreetingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingResponse.Unmarshal(m, b)
}
func (m *GreetingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingResponse.Marshal(b, m, deterministic)
}
func (m *GreetingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingResponse.Merge(m, src)
}
func (m *GreetingResponse) XXX_Size() int {
	return xxx_messageInfo_GreetingResponse.Size(m)
}
func (m *GreetingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingResponse proto.InternalMessageInfo

func (m *GreetingResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*NoParams)(nil), "greet_service.NoParams")
	proto.RegisterType((*ReqList)(nil), "greet_service.reqList")
	proto.RegisterType((*ResList)(nil), "greet_service.resList")
	proto.RegisterType((*GreetingRequest)(nil), "greet_service.greetingRequest")
	proto.RegisterType((*GreetingResponse)(nil), "greet_service.greetingResponse")
}

func init() {
	proto.RegisterFile("proto/greet.proto", fileDescriptor_95ca2ad3f55a58e3)
}

var fileDescriptor_95ca2ad3f55a58e3 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xc9, 0xdb, 0x97, 0xfe, 0x19, 0x14, 0x6b, 0x0e, 0x6d, 0x29, 0x52, 0xcb, 0xa2, 0xb8,
	0x20, 0x6c, 0x8b, 0x7e, 0x02, 0xab, 0x60, 0x0f, 0x22, 0x65, 0xeb, 0xc9, 0x8b, 0xa4, 0x65, 0x58,
	0x02, 0xbb, 0x49, 0x3b, 0x13, 0x05, 0xbf, 0xb3, 0x1f, 0x42, 0x36, 0x6d, 0x5c, 0x6c, 0x91, 0xde,
	0x66, 0x26, 0xcf, 0xfc, 0x66, 0xe6, 0x21, 0x70, 0xba, 0x22, 0xeb, 0xec, 0x28, 0x23, 0x44, 0x97,
	0xf8, 0x58, 0x1e, 0xfb, 0xe4, 0x8d, 0x91, 0x3e, 0xf4, 0x12, 0x23, 0x80, 0xe6, 0xb3, 0x9d, 0x29,
	0x52, 0x05, 0x47, 0x97, 0xd0, 0x20, 0x5c, 0x3f, 0x69, 0x76, 0xb2, 0x0f, 0x4d, 0xc2, 0xf5, 0x3b,
	0xb2, 0xe3, 0x9e, 0x18, 0xd6, 0xe2, 0x56, 0xfa, 0x93, 0x47, 0x57, 0xa5, 0x8c, 0xbd, 0xec, 0x0c,
	0x5a, 0x84, 0xbc, 0xb2, 0x86, 0x31, 0xe8, 0xaa, 0x42, 0x74, 0x0d, 0x27, 0x7e, 0x98, 0x36, 0x59,
	0xba, 0x69, 0x96, 0x3d, 0x68, 0x14, 0xc8, 0xac, 0x32, 0xec, 0x89, 0xa1, 0x88, 0x5b, 0x69, 0x48,
	0xa3, 0x0b, 0x68, 0x57, 0xe2, 0x0d, 0x41, 0xb6, 0xa1, 0x56, 0x70, 0xb6, 0x55, 0x96, 0xe1, 0xcd,
	0xd7, 0xbf, 0x8a, 0x39, 0xdf, 0x9c, 0x20, 0xa7, 0x70, 0xf4, 0x58, 0x96, 0xee, 0x66, 0x48, 0x6c,
	0x8d, 0xec, 0x26, 0xbf, 0x4e, 0x4c, 0xc2, 0x7d, 0xfd, 0xf3, 0x9d, 0x87, 0xbd, 0x79, 0x2f, 0xd0,
	0x9d, 0xab, 0xcf, 0x29, 0xe6, 0xb9, 0x2d, 0xe1, 0x48, 0x73, 0x47, 0xa8, 0x0a, 0x6d, 0x32, 0xd9,
	0xd9, 0xe9, 0xdd, 0x1a, 0x75, 0x90, 0x39, 0x16, 0x32, 0x85, 0x4e, 0xa0, 0xde, 0xe7, 0x1a, 0x8d,
	0x0b, 0x54, 0x39, 0xf8, 0xb3, 0xd9, 0xbb, 0xd5, 0xdf, 0x1f, 0xea, 0x6d, 0x8f, 0x85, 0x44, 0x18,
	0x04, 0xe6, 0x44, 0x3f, 0x68, 0xc2, 0x25, 0x69, 0x6b, 0x54, 0x5e, 0x2d, 0x7c, 0x88, 0x7d, 0x68,
	0xf1, 0x58, 0x8c, 0xc5, 0xa4, 0xfe, 0xfa, 0x3f, 0x19, 0xad, 0x16, 0x8b, 0xba, 0xff, 0x3b, 0xb7,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x0a, 0xe5, 0x6c, 0x50, 0x02, 0x00, 0x00,
}
