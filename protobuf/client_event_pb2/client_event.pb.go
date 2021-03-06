// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/client_event_pb2/client_event.proto

/*
Package client_event_pb2 is a generated protocol buffer package.

It is generated from these files:
	sawtooth_sdk/protobuf/client_event_pb2/client_event.proto

It has these top-level messages:
	ClientEventsSubscribeRequest
	ClientEventsSubscribeResponse
	ClientEventsUnsubscribeRequest
	ClientEventsUnsubscribeResponse
	ClientEventsGetRequest
	ClientEventsGetResponse
*/
package client_event_pb2

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

type ClientEventsSubscribeResponse_Status int32

const (
	ClientEventsSubscribeResponse_STATUS_UNSET   ClientEventsSubscribeResponse_Status = 0
	ClientEventsSubscribeResponse_OK             ClientEventsSubscribeResponse_Status = 1
	ClientEventsSubscribeResponse_INVALID_FILTER ClientEventsSubscribeResponse_Status = 2
	ClientEventsSubscribeResponse_UNKNOWN_BLOCK  ClientEventsSubscribeResponse_Status = 3
)

var ClientEventsSubscribeResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INVALID_FILTER",
	3: "UNKNOWN_BLOCK",
}
var ClientEventsSubscribeResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INVALID_FILTER": 2,
	"UNKNOWN_BLOCK":  3,
}

func (x ClientEventsSubscribeResponse_Status) String() string {
	return proto.EnumName(ClientEventsSubscribeResponse_Status_name, int32(x))
}
func (ClientEventsSubscribeResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ClientEventsUnsubscribeResponse_Status int32

const (
	ClientEventsUnsubscribeResponse_STATUS_UNSET   ClientEventsUnsubscribeResponse_Status = 0
	ClientEventsUnsubscribeResponse_OK             ClientEventsUnsubscribeResponse_Status = 1
	ClientEventsUnsubscribeResponse_INTERNAL_ERROR ClientEventsUnsubscribeResponse_Status = 2
)

var ClientEventsUnsubscribeResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
}
var ClientEventsUnsubscribeResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
}

func (x ClientEventsUnsubscribeResponse_Status) String() string {
	return proto.EnumName(ClientEventsUnsubscribeResponse_Status_name, int32(x))
}
func (ClientEventsUnsubscribeResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type ClientEventsGetResponse_Status int32

const (
	ClientEventsGetResponse_STATUS_UNSET   ClientEventsGetResponse_Status = 0
	ClientEventsGetResponse_OK             ClientEventsGetResponse_Status = 1
	ClientEventsGetResponse_INTERNAL_ERROR ClientEventsGetResponse_Status = 2
	ClientEventsGetResponse_INVALID_FILTER ClientEventsGetResponse_Status = 3
	ClientEventsGetResponse_UNKNOWN_BLOCK  ClientEventsGetResponse_Status = 4
)

var ClientEventsGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "INVALID_FILTER",
	4: "UNKNOWN_BLOCK",
}
var ClientEventsGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"INVALID_FILTER": 3,
	"UNKNOWN_BLOCK":  4,
}

func (x ClientEventsGetResponse_Status) String() string {
	return proto.EnumName(ClientEventsGetResponse_Status_name, int32(x))
}
func (ClientEventsGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type ClientEventsSubscribeRequest struct {
	Subscriptions []*events.EventSubscription `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// The block id (or ids, if trying to walk back a fork) the subscriber last
	// received events on. It can be set to empty if it has not yet received the
	// genesis block.
	LastKnownBlockIds []string `protobuf:"bytes,2,rep,name=last_known_block_ids,json=lastKnownBlockIds" json:"last_known_block_ids,omitempty"`
}

func (m *ClientEventsSubscribeRequest) Reset()                    { *m = ClientEventsSubscribeRequest{} }
func (m *ClientEventsSubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientEventsSubscribeRequest) ProtoMessage()               {}
func (*ClientEventsSubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientEventsSubscribeRequest) GetSubscriptions() []*events.EventSubscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *ClientEventsSubscribeRequest) GetLastKnownBlockIds() []string {
	if m != nil {
		return m.LastKnownBlockIds
	}
	return nil
}

type ClientEventsSubscribeResponse struct {
	Status ClientEventsSubscribeResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientEventsSubscribeResponse_Status" json:"status,omitempty"`
	// Additional information about the response status
	ResponseMessage string `protobuf:"bytes,2,opt,name=response_message,json=responseMessage" json:"response_message,omitempty"`
}

func (m *ClientEventsSubscribeResponse) Reset()                    { *m = ClientEventsSubscribeResponse{} }
func (m *ClientEventsSubscribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientEventsSubscribeResponse) ProtoMessage()               {}
func (*ClientEventsSubscribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientEventsSubscribeResponse) GetStatus() ClientEventsSubscribeResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientEventsSubscribeResponse_STATUS_UNSET
}

func (m *ClientEventsSubscribeResponse) GetResponseMessage() string {
	if m != nil {
		return m.ResponseMessage
	}
	return ""
}

type ClientEventsUnsubscribeRequest struct {
}

func (m *ClientEventsUnsubscribeRequest) Reset()                    { *m = ClientEventsUnsubscribeRequest{} }
func (m *ClientEventsUnsubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientEventsUnsubscribeRequest) ProtoMessage()               {}
func (*ClientEventsUnsubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ClientEventsUnsubscribeResponse struct {
	Status ClientEventsUnsubscribeResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientEventsUnsubscribeResponse_Status" json:"status,omitempty"`
}

func (m *ClientEventsUnsubscribeResponse) Reset()                    { *m = ClientEventsUnsubscribeResponse{} }
func (m *ClientEventsUnsubscribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientEventsUnsubscribeResponse) ProtoMessage()               {}
func (*ClientEventsUnsubscribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientEventsUnsubscribeResponse) GetStatus() ClientEventsUnsubscribeResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientEventsUnsubscribeResponse_STATUS_UNSET
}

type ClientEventsGetRequest struct {
	Subscriptions []*events.EventSubscription `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	BlockIds      []string                    `protobuf:"bytes,2,rep,name=block_ids,json=blockIds" json:"block_ids,omitempty"`
}

func (m *ClientEventsGetRequest) Reset()                    { *m = ClientEventsGetRequest{} }
func (m *ClientEventsGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientEventsGetRequest) ProtoMessage()               {}
func (*ClientEventsGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClientEventsGetRequest) GetSubscriptions() []*events.EventSubscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *ClientEventsGetRequest) GetBlockIds() []string {
	if m != nil {
		return m.BlockIds
	}
	return nil
}

type ClientEventsGetResponse struct {
	Status ClientEventsGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientEventsGetResponse_Status" json:"status,omitempty"`
	Events []*events.Event                `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
}

func (m *ClientEventsGetResponse) Reset()                    { *m = ClientEventsGetResponse{} }
func (m *ClientEventsGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientEventsGetResponse) ProtoMessage()               {}
func (*ClientEventsGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClientEventsGetResponse) GetStatus() ClientEventsGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientEventsGetResponse_STATUS_UNSET
}

func (m *ClientEventsGetResponse) GetEvents() []*events.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientEventsSubscribeRequest)(nil), "ClientEventsSubscribeRequest")
	proto.RegisterType((*ClientEventsSubscribeResponse)(nil), "ClientEventsSubscribeResponse")
	proto.RegisterType((*ClientEventsUnsubscribeRequest)(nil), "ClientEventsUnsubscribeRequest")
	proto.RegisterType((*ClientEventsUnsubscribeResponse)(nil), "ClientEventsUnsubscribeResponse")
	proto.RegisterType((*ClientEventsGetRequest)(nil), "ClientEventsGetRequest")
	proto.RegisterType((*ClientEventsGetResponse)(nil), "ClientEventsGetResponse")
	proto.RegisterEnum("ClientEventsSubscribeResponse_Status", ClientEventsSubscribeResponse_Status_name, ClientEventsSubscribeResponse_Status_value)
	proto.RegisterEnum("ClientEventsUnsubscribeResponse_Status", ClientEventsUnsubscribeResponse_Status_name, ClientEventsUnsubscribeResponse_Status_value)
	proto.RegisterEnum("ClientEventsGetResponse_Status", ClientEventsGetResponse_Status_name, ClientEventsGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("sawtooth_sdk/protobuf/client_event_pb2/client_event.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0xed, 0x8a, 0xd3, 0x40,
	0x14, 0x75, 0x5a, 0x09, 0xee, 0x5d, 0x77, 0xcd, 0x0e, 0x7e, 0x2c, 0x7e, 0xec, 0x2e, 0x03, 0xe2,
	0x8a, 0x38, 0x85, 0x0a, 0x7e, 0xfc, 0x10, 0x69, 0xd7, 0x2a, 0xa1, 0x35, 0x91, 0x49, 0xa2, 0x20,
	0xc8, 0x90, 0xb4, 0xa3, 0x86, 0xd6, 0x4c, 0xec, 0x4c, 0xdc, 0x67, 0xf0, 0x15, 0x7c, 0x2b, 0xc1,
	0x07, 0x32, 0x99, 0x24, 0xd0, 0xcd, 0x6e, 0x55, 0xf4, 0x4f, 0x29, 0xf7, 0x9c, 0x3b, 0xf7, 0x9c,
	0x9b, 0x73, 0xe1, 0x89, 0x8a, 0x8e, 0xb5, 0x94, 0xfa, 0x13, 0x57, 0xb3, 0x79, 0x2f, 0x5b, 0x4a,
	0x2d, 0xe3, 0xfc, 0x43, 0x6f, 0xba, 0x48, 0x44, 0xaa, 0xb9, 0xf8, 0x5a, 0xfe, 0x66, 0x71, 0xff,
	0x44, 0x81, 0x1a, 0xda, 0xf5, 0xfb, 0x67, 0xb7, 0x1a, 0x8a, 0x32, 0x4d, 0xd5, 0xdf, 0x8a, 0x4e,
	0xbe, 0x21, 0xb8, 0x79, 0x64, 0x5e, 0x19, 0x99, 0xb2, 0x9f, 0xc7, 0x6a, 0xba, 0x4c, 0x62, 0xc1,
	0xc4, 0x97, 0x5c, 0x28, 0x8d, 0x1f, 0xc3, 0x96, 0xaa, 0x6a, 0x99, 0x4e, 0x64, 0xaa, 0x76, 0xd1,
	0x41, 0xf7, 0x70, 0xb3, 0x8f, 0xa9, 0xe1, 0xfb, 0x2b, 0x10, 0x3b, 0x49, 0xc4, 0x3d, 0xb8, 0xbc,
	0x88, 0x94, 0xe6, 0xf3, 0x54, 0x1e, 0xa7, 0x3c, 0x5e, 0xc8, 0xe9, 0x9c, 0x27, 0x33, 0xb5, 0xdb,
	0x29, 0x1e, 0xd8, 0x60, 0x3b, 0x25, 0x36, 0x2e, 0xa1, 0x61, 0x89, 0x38, 0x33, 0x45, 0x7e, 0x22,
	0xb8, 0xb5, 0x46, 0x8b, 0xca, 0x8a, 0x17, 0x05, 0x7e, 0x0a, 0x96, 0xd2, 0x91, 0xce, 0x4b, 0x15,
	0xe8, 0x70, 0xbb, 0x7f, 0x9b, 0xfe, 0x96, 0x4f, 0x7d, 0x43, 0x66, 0x75, 0x13, 0xbe, 0x0b, 0xf6,
	0xb2, 0x86, 0xf8, 0x67, 0xa1, 0x54, 0xf4, 0x51, 0x14, 0x6a, 0x50, 0xa1, 0xe6, 0x52, 0x53, 0x7f,
	0x55, 0x95, 0x89, 0x03, 0x56, 0xd5, 0x8c, 0x6d, 0xb8, 0xe8, 0x07, 0x83, 0x20, 0xf4, 0x79, 0xe8,
	0xfa, 0xa3, 0xc0, 0x3e, 0x87, 0x2d, 0xe8, 0x78, 0x63, 0x1b, 0x61, 0x0c, 0xdb, 0x8e, 0xfb, 0x66,
	0x30, 0x71, 0x9e, 0xf3, 0x17, 0xce, 0x24, 0x18, 0x31, 0xbb, 0x83, 0x77, 0x60, 0x2b, 0x74, 0xc7,
	0xae, 0xf7, 0xd6, 0xe5, 0xc3, 0x89, 0x77, 0x34, 0xb6, 0xbb, 0xe4, 0x00, 0xf6, 0x56, 0x55, 0x86,
	0xa9, 0x6a, 0xed, 0x98, 0x7c, 0x47, 0xb0, 0xbf, 0x96, 0x52, 0x5b, 0x7f, 0xd6, 0xb2, 0x7e, 0x87,
	0xfe, 0xa1, 0xa3, 0x65, 0x9e, 0x3c, 0xfc, 0x5b, 0x47, 0x85, 0x11, 0x77, 0x30, 0xe1, 0x23, 0xc6,
	0xbc, 0xc2, 0x11, 0x91, 0x70, 0x75, 0x75, 0xd2, 0x4b, 0xa1, 0xff, 0x3f, 0x1a, 0x37, 0x60, 0xa3,
	0x9d, 0x87, 0x0b, 0x71, 0x13, 0x83, 0x1f, 0x08, 0xae, 0x9d, 0x9a, 0x58, 0x6f, 0xe1, 0x51, 0x6b,
	0x0b, 0xfb, 0x74, 0x0d, 0xb3, 0xfd, 0xe9, 0xf7, 0xc0, 0xaa, 0x72, 0x6f, 0xc6, 0x6d, 0xf6, 0xad,
	0x4a, 0x24, 0xab, 0xab, 0xe4, 0xfd, 0xbf, 0x6d, 0xe7, 0x8c, 0x0c, 0x74, 0x4f, 0x67, 0xe0, 0xfc,
	0xf0, 0x1e, 0x5c, 0x69, 0xee, 0x92, 0x16, 0x77, 0x49, 0x9b, 0xbb, 0x7c, 0x8d, 0xde, 0xd9, 0xed,
	0xab, 0x8e, 0x2d, 0x83, 0x3e, 0xf8, 0x15, 0x00, 0x00, 0xff, 0xff, 0x01, 0x3c, 0x95, 0x77, 0x06,
	0x04, 0x00, 0x00,
}
