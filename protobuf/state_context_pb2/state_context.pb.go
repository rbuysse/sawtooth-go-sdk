// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/state_context_pb2/state_context.proto

/*
Package state_context_pb2 is a generated protocol buffer package.

It is generated from these files:
	sawtooth_sdk/protobuf/state_context_pb2/state_context.proto

It has these top-level messages:
	Entry
	TpStateGetRequest
	TpStateGetResponse
	TpStateSetRequest
	TpStateSetResponse
	TpStateDeleteRequest
	TpStateDeleteResponse
	TpReceiptAddDataRequest
	TpReceiptAddDataResponse
	TpEventAddRequest
	TpEventAddResponse
*/
package state_context_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import events "sawtooth_sdk/protobuf/events_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TpStateGetResponse_Status int32

const (
	TpStateGetResponse_STATUS_UNSET        TpStateGetResponse_Status = 0
	TpStateGetResponse_OK                  TpStateGetResponse_Status = 1
	TpStateGetResponse_AUTHORIZATION_ERROR TpStateGetResponse_Status = 2
)

var TpStateGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "AUTHORIZATION_ERROR",
}
var TpStateGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"AUTHORIZATION_ERROR": 2,
}

func (x TpStateGetResponse_Status) String() string {
	return proto.EnumName(TpStateGetResponse_Status_name, int32(x))
}
func (TpStateGetResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type TpStateSetResponse_Status int32

const (
	TpStateSetResponse_STATUS_UNSET        TpStateSetResponse_Status = 0
	TpStateSetResponse_OK                  TpStateSetResponse_Status = 1
	TpStateSetResponse_AUTHORIZATION_ERROR TpStateSetResponse_Status = 2
)

var TpStateSetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "AUTHORIZATION_ERROR",
}
var TpStateSetResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"AUTHORIZATION_ERROR": 2,
}

func (x TpStateSetResponse_Status) String() string {
	return proto.EnumName(TpStateSetResponse_Status_name, int32(x))
}
func (TpStateSetResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type TpStateDeleteResponse_Status int32

const (
	TpStateDeleteResponse_STATUS_UNSET        TpStateDeleteResponse_Status = 0
	TpStateDeleteResponse_OK                  TpStateDeleteResponse_Status = 1
	TpStateDeleteResponse_AUTHORIZATION_ERROR TpStateDeleteResponse_Status = 2
)

var TpStateDeleteResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "AUTHORIZATION_ERROR",
}
var TpStateDeleteResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"AUTHORIZATION_ERROR": 2,
}

func (x TpStateDeleteResponse_Status) String() string {
	return proto.EnumName(TpStateDeleteResponse_Status_name, int32(x))
}
func (TpStateDeleteResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type TpReceiptAddDataResponse_Status int32

const (
	TpReceiptAddDataResponse_STATUS_UNSET TpReceiptAddDataResponse_Status = 0
	TpReceiptAddDataResponse_OK           TpReceiptAddDataResponse_Status = 1
	TpReceiptAddDataResponse_ERROR        TpReceiptAddDataResponse_Status = 2
)

var TpReceiptAddDataResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var TpReceiptAddDataResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x TpReceiptAddDataResponse_Status) String() string {
	return proto.EnumName(TpReceiptAddDataResponse_Status_name, int32(x))
}
func (TpReceiptAddDataResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

type TpEventAddResponse_Status int32

const (
	TpEventAddResponse_STATUS_UNSET TpEventAddResponse_Status = 0
	TpEventAddResponse_OK           TpEventAddResponse_Status = 1
	TpEventAddResponse_ERROR        TpEventAddResponse_Status = 2
)

var TpEventAddResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var TpEventAddResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x TpEventAddResponse_Status) String() string {
	return proto.EnumName(TpEventAddResponse_Status_name, int32(x))
}
func (TpEventAddResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

// An entry in the State
type Entry struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// A request from a handler/tp for the values at a series of addresses
type TpStateGetRequest struct {
	// The context id that references a context in the contextmanager
	ContextId string   `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Addresses []string `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *TpStateGetRequest) Reset()                    { *m = TpStateGetRequest{} }
func (m *TpStateGetRequest) String() string            { return proto.CompactTextString(m) }
func (*TpStateGetRequest) ProtoMessage()               {}
func (*TpStateGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TpStateGetRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpStateGetRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// A response from the contextmanager/validator with a series of State entries
type TpStateGetResponse struct {
	Entries []*Entry                  `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	Status  TpStateGetResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpStateGetResponse_Status" json:"status,omitempty"`
}

func (m *TpStateGetResponse) Reset()                    { *m = TpStateGetResponse{} }
func (m *TpStateGetResponse) String() string            { return proto.CompactTextString(m) }
func (*TpStateGetResponse) ProtoMessage()               {}
func (*TpStateGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TpStateGetResponse) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *TpStateGetResponse) GetStatus() TpStateGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpStateGetResponse_STATUS_UNSET
}

// A request from the handler/tp to put entries in the state of a context
type TpStateSetRequest struct {
	ContextId string   `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Entries   []*Entry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *TpStateSetRequest) Reset()                    { *m = TpStateSetRequest{} }
func (m *TpStateSetRequest) String() string            { return proto.CompactTextString(m) }
func (*TpStateSetRequest) ProtoMessage()               {}
func (*TpStateSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TpStateSetRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpStateSetRequest) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// A response from the contextmanager/validator with the addresses that were set
type TpStateSetResponse struct {
	Addresses []string                  `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Status    TpStateSetResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpStateSetResponse_Status" json:"status,omitempty"`
}

func (m *TpStateSetResponse) Reset()                    { *m = TpStateSetResponse{} }
func (m *TpStateSetResponse) String() string            { return proto.CompactTextString(m) }
func (*TpStateSetResponse) ProtoMessage()               {}
func (*TpStateSetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TpStateSetResponse) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *TpStateSetResponse) GetStatus() TpStateSetResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpStateSetResponse_STATUS_UNSET
}

// A request from the handler/tp to delete state entries at an collection of addresses
type TpStateDeleteRequest struct {
	ContextId string   `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Addresses []string `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *TpStateDeleteRequest) Reset()                    { *m = TpStateDeleteRequest{} }
func (m *TpStateDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*TpStateDeleteRequest) ProtoMessage()               {}
func (*TpStateDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TpStateDeleteRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpStateDeleteRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// A response form the contextmanager/validator with the addresses that were deleted
type TpStateDeleteResponse struct {
	Addresses []string                     `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Status    TpStateDeleteResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpStateDeleteResponse_Status" json:"status,omitempty"`
}

func (m *TpStateDeleteResponse) Reset()                    { *m = TpStateDeleteResponse{} }
func (m *TpStateDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*TpStateDeleteResponse) ProtoMessage()               {}
func (*TpStateDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TpStateDeleteResponse) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *TpStateDeleteResponse) GetStatus() TpStateDeleteResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpStateDeleteResponse_STATUS_UNSET
}

// The request from the transaction processor to the validator append data
// to a transaction receipt
type TpReceiptAddDataRequest struct {
	// The context id that references a context in the context manager
	ContextId string `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TpReceiptAddDataRequest) Reset()                    { *m = TpReceiptAddDataRequest{} }
func (m *TpReceiptAddDataRequest) String() string            { return proto.CompactTextString(m) }
func (*TpReceiptAddDataRequest) ProtoMessage()               {}
func (*TpReceiptAddDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TpReceiptAddDataRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpReceiptAddDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The response from the validator to the transaction processor to verify that
// data has been appended to a transaction receipt
type TpReceiptAddDataResponse struct {
	Status TpReceiptAddDataResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpReceiptAddDataResponse_Status" json:"status,omitempty"`
}

func (m *TpReceiptAddDataResponse) Reset()                    { *m = TpReceiptAddDataResponse{} }
func (m *TpReceiptAddDataResponse) String() string            { return proto.CompactTextString(m) }
func (*TpReceiptAddDataResponse) ProtoMessage()               {}
func (*TpReceiptAddDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TpReceiptAddDataResponse) GetStatus() TpReceiptAddDataResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpReceiptAddDataResponse_STATUS_UNSET
}

type TpEventAddRequest struct {
	ContextId string        `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Event     *events.Event `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
}

func (m *TpEventAddRequest) Reset()                    { *m = TpEventAddRequest{} }
func (m *TpEventAddRequest) String() string            { return proto.CompactTextString(m) }
func (*TpEventAddRequest) ProtoMessage()               {}
func (*TpEventAddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TpEventAddRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpEventAddRequest) GetEvent() *events.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type TpEventAddResponse struct {
	Status TpEventAddResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpEventAddResponse_Status" json:"status,omitempty"`
}

func (m *TpEventAddResponse) Reset()                    { *m = TpEventAddResponse{} }
func (m *TpEventAddResponse) String() string            { return proto.CompactTextString(m) }
func (*TpEventAddResponse) ProtoMessage()               {}
func (*TpEventAddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TpEventAddResponse) GetStatus() TpEventAddResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpEventAddResponse_STATUS_UNSET
}

func init() {
	proto.RegisterType((*Entry)(nil), "Entry")
	proto.RegisterType((*TpStateGetRequest)(nil), "TpStateGetRequest")
	proto.RegisterType((*TpStateGetResponse)(nil), "TpStateGetResponse")
	proto.RegisterType((*TpStateSetRequest)(nil), "TpStateSetRequest")
	proto.RegisterType((*TpStateSetResponse)(nil), "TpStateSetResponse")
	proto.RegisterType((*TpStateDeleteRequest)(nil), "TpStateDeleteRequest")
	proto.RegisterType((*TpStateDeleteResponse)(nil), "TpStateDeleteResponse")
	proto.RegisterType((*TpReceiptAddDataRequest)(nil), "TpReceiptAddDataRequest")
	proto.RegisterType((*TpReceiptAddDataResponse)(nil), "TpReceiptAddDataResponse")
	proto.RegisterType((*TpEventAddRequest)(nil), "TpEventAddRequest")
	proto.RegisterType((*TpEventAddResponse)(nil), "TpEventAddResponse")
	proto.RegisterEnum("TpStateGetResponse_Status", TpStateGetResponse_Status_name, TpStateGetResponse_Status_value)
	proto.RegisterEnum("TpStateSetResponse_Status", TpStateSetResponse_Status_name, TpStateSetResponse_Status_value)
	proto.RegisterEnum("TpStateDeleteResponse_Status", TpStateDeleteResponse_Status_name, TpStateDeleteResponse_Status_value)
	proto.RegisterEnum("TpReceiptAddDataResponse_Status", TpReceiptAddDataResponse_Status_name, TpReceiptAddDataResponse_Status_value)
	proto.RegisterEnum("TpEventAddResponse_Status", TpEventAddResponse_Status_name, TpEventAddResponse_Status_value)
}

func init() {
	proto.RegisterFile("sawtooth_sdk/protobuf/state_context_pb2/state_context.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x71, 0xcb, 0x32, 0xf5, 0x31, 0xa1, 0xce, 0x30, 0xad, 0x9a, 0x98, 0x54, 0xe5, 0xb4,
	0x03, 0xf3, 0xa4, 0xa2, 0x49, 0x48, 0x3b, 0x15, 0xad, 0x82, 0x0a, 0xb4, 0x4e, 0xb6, 0x7b, 0xd9,
	0x25, 0x4a, 0xe7, 0x87, 0xa8, 0x40, 0x49, 0x88, 0x5d, 0x06, 0x67, 0x3e, 0x0a, 0x07, 0xc4, 0xb7,
	0x24, 0x71, 0x9c, 0x2d, 0x4b, 0x33, 0x29, 0x12, 0xe5, 0x96, 0xfc, 0xfd, 0xfc, 0xf7, 0xfb, 0x3d,
	0xfb, 0x3d, 0x38, 0xd3, 0xe1, 0x8d, 0x89, 0x63, 0xf3, 0x29, 0xd0, 0xea, 0xf3, 0x49, 0x92, 0xc6,
	0x26, 0x5e, 0xac, 0x3e, 0x9e, 0x68, 0x13, 0x1a, 0x0c, 0xae, 0xe3, 0xc8, 0xe0, 0x77, 0x13, 0x24,
	0x8b, 0xd1, 0x7d, 0x85, 0xd9, 0xc0, 0x83, 0xe3, 0xe6, 0xcd, 0xf8, 0x0d, 0x23, 0xa3, 0xed, 0xae,
	0xe2, 0xb3, 0x08, 0xf7, 0x4f, 0x61, 0x6b, 0x12, 0x99, 0xf4, 0x07, 0x1d, 0xc0, 0x76, 0xa8, 0x54,
	0x8a, 0x5a, 0x0f, 0xc8, 0x90, 0x1c, 0xf5, 0x78, 0xf9, 0x4b, 0x29, 0x3c, 0x56, 0xa1, 0x09, 0x07,
	0x9d, 0x4c, 0xde, 0xe1, 0xf6, 0xdb, 0xbf, 0x84, 0x5d, 0x99, 0x88, 0xfc, 0xf8, 0xb7, 0x68, 0x38,
	0x7e, 0x5d, 0xa1, 0x36, 0xf4, 0x10, 0xa0, 0xcc, 0x6e, 0xa9, 0x9c, 0x4b, 0xcf, 0x29, 0x53, 0x45,
	0x5f, 0x40, 0xcf, 0x59, 0xa2, 0xce, 0xcc, 0xba, 0xf9, 0xea, 0xad, 0xe0, 0xff, 0x26, 0x40, 0xab,
	0x96, 0x3a, 0x89, 0x23, 0x8d, 0x74, 0x08, 0xdb, 0x59, 0xb6, 0xe9, 0x12, 0xf3, 0xb4, 0xba, 0x47,
	0x4f, 0x46, 0x1e, 0xb3, 0xf9, 0xf2, 0x52, 0xa6, 0x23, 0xf0, 0xf2, 0x3a, 0xac, 0xb4, 0x4d, 0xf0,
	0xe9, 0xe8, 0x80, 0xad, 0xdb, 0x30, 0x61, 0x23, 0xb8, 0x8b, 0xf4, 0xcf, 0xc0, 0x2b, 0x14, 0xda,
	0x87, 0x1d, 0x21, 0xc7, 0x72, 0x2e, 0x82, 0xf9, 0x85, 0x98, 0xc8, 0xfe, 0x23, 0xea, 0x41, 0x67,
	0xf6, 0xbe, 0x4f, 0xe8, 0x3e, 0x3c, 0x1b, 0xcf, 0xe5, 0xbb, 0x19, 0x9f, 0x5e, 0x8d, 0xe5, 0x74,
	0x76, 0x11, 0x4c, 0x38, 0x9f, 0xf1, 0x7e, 0xc7, 0x97, 0xb7, 0xec, 0xa2, 0x35, 0x7b, 0x05, 0xa3,
	0xd3, 0x88, 0xe1, 0xff, 0xba, 0xe3, 0x17, 0x15, 0xfe, 0x7b, 0x45, 0x23, 0xb5, 0xa2, 0x3d, 0xcc,
	0x2e, 0xfe, 0x13, 0xbb, 0x80, 0xe7, 0xee, 0x84, 0x73, 0xfc, 0x82, 0x06, 0x37, 0x72, 0xf5, 0x7f,
	0x08, 0xec, 0xd5, 0x5c, 0x5b, 0xd1, 0x9f, 0xd6, 0xe8, 0x0f, 0x59, 0xa3, 0xcb, 0x46, 0x0b, 0xf0,
	0x01, 0xf6, 0x65, 0xc2, 0xf1, 0x1a, 0x97, 0x89, 0x19, 0x2b, 0x75, 0x9e, 0x35, 0x43, 0xcb, 0x1a,
	0x94, 0x6d, 0xd4, 0xad, 0xb4, 0xd1, 0x4f, 0x02, 0x83, 0x75, 0x3b, 0x07, 0xff, 0xba, 0x86, 0x37,
	0x64, 0x0f, 0x85, 0xd6, 0x09, 0x8f, 0x5b, 0x10, 0xf6, 0xb2, 0xc6, 0x77, 0x4c, 0xb6, 0x99, 0x27,
	0xf9, 0x54, 0xc8, 0x7c, 0x5b, 0xdf, 0xe8, 0x96, 0x9d, 0x23, 0x36, 0x37, 0xfb, 0x9c, 0xf3, 0x3f,
	0x5e, 0x88, 0xfe, 0x4d, 0xfe, 0x96, 0xef, 0x1c, 0x1d, 0x50, 0xd3, 0x6b, 0xad, 0x07, 0xfd, 0x1b,
	0xca, 0x9b, 0x97, 0xb0, 0x57, 0xce, 0x3f, 0x96, 0xcd, 0x3f, 0x56, 0xce, 0xbf, 0x4b, 0x72, 0xb5,
	0xbb, 0x36, 0x3f, 0x17, 0x9e, 0x5d, 0x7e, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x30, 0xde,
	0x33, 0x71, 0x05, 0x00, 0x00,
}
