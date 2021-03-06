// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/client_batch_pb2/client_batch.proto

/*
Package client_batch_pb2 is a generated protocol buffer package.

It is generated from these files:
	sawtooth_sdk/protobuf/client_batch_pb2/client_batch.proto

It has these top-level messages:
	ClientBatchListRequest
	ClientBatchListResponse
	ClientBatchGetRequest
	ClientBatchGetResponse
*/
package client_batch_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import batch "sawtooth_sdk/protobuf/batch_pb2"
import client_list_control "sawtooth_sdk/protobuf/client_list_control_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientBatchListResponse_Status int32

const (
	ClientBatchListResponse_STATUS_UNSET   ClientBatchListResponse_Status = 0
	ClientBatchListResponse_OK             ClientBatchListResponse_Status = 1
	ClientBatchListResponse_INTERNAL_ERROR ClientBatchListResponse_Status = 2
	ClientBatchListResponse_NOT_READY      ClientBatchListResponse_Status = 3
	ClientBatchListResponse_NO_ROOT        ClientBatchListResponse_Status = 4
	ClientBatchListResponse_NO_RESOURCE    ClientBatchListResponse_Status = 5
	ClientBatchListResponse_INVALID_PAGING ClientBatchListResponse_Status = 6
	ClientBatchListResponse_INVALID_SORT   ClientBatchListResponse_Status = 7
)

var ClientBatchListResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_PAGING",
	7: "INVALID_SORT",
}
var ClientBatchListResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NOT_READY":      3,
	"NO_ROOT":        4,
	"NO_RESOURCE":    5,
	"INVALID_PAGING": 6,
	"INVALID_SORT":   7,
}

func (x ClientBatchListResponse_Status) String() string {
	return proto.EnumName(ClientBatchListResponse_Status_name, int32(x))
}
func (ClientBatchListResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ClientBatchGetResponse_Status int32

const (
	ClientBatchGetResponse_STATUS_UNSET   ClientBatchGetResponse_Status = 0
	ClientBatchGetResponse_OK             ClientBatchGetResponse_Status = 1
	ClientBatchGetResponse_INTERNAL_ERROR ClientBatchGetResponse_Status = 2
	ClientBatchGetResponse_NO_RESOURCE    ClientBatchGetResponse_Status = 5
)

var ClientBatchGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
}
var ClientBatchGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
}

func (x ClientBatchGetResponse_Status) String() string {
	return proto.EnumName(ClientBatchGetResponse_Status_name, int32(x))
}
func (ClientBatchGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// A request to return a list of batches from the validator. May include the id
// of a particular block to be the `head` of the chain being requested. In that
// case the list will include the batches from that block, and all batches
// previous to that block on the chain. Filter with specific `batch_ids`.
type ClientBatchListRequest struct {
	HeadId   string                                    `protobuf:"bytes,1,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	BatchIds []string                                  `protobuf:"bytes,2,rep,name=batch_ids,json=batchIds" json:"batch_ids,omitempty"`
	Paging   *client_list_control.ClientPagingControls `protobuf:"bytes,3,opt,name=paging" json:"paging,omitempty"`
	Sorting  []*client_list_control.ClientSortControls `protobuf:"bytes,4,rep,name=sorting" json:"sorting,omitempty"`
}

func (m *ClientBatchListRequest) Reset()                    { *m = ClientBatchListRequest{} }
func (m *ClientBatchListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchListRequest) ProtoMessage()               {}
func (*ClientBatchListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientBatchListRequest) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientBatchListRequest) GetBatchIds() []string {
	if m != nil {
		return m.BatchIds
	}
	return nil
}

func (m *ClientBatchListRequest) GetPaging() *client_list_control.ClientPagingControls {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ClientBatchListRequest) GetSorting() []*client_list_control.ClientSortControls {
	if m != nil {
		return m.Sorting
	}
	return nil
}

// A response that lists batches from newest to oldest.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NOT_READY - the validator does not yet have a genesis block
//   * NO_ROOT - the head block specified was not found
//   * NO_RESOURCE - no batches were found with the parameters specified
//   * INVALID_PAGING - the paging controls were malformed or out of range
//   * INVALID_SORT - the sorting controls were malformed or invalid
type ClientBatchListResponse struct {
	Status  ClientBatchListResponse_Status            `protobuf:"varint,1,opt,name=status,enum=ClientBatchListResponse_Status" json:"status,omitempty"`
	Batches []*batch.Batch                            `protobuf:"bytes,2,rep,name=batches" json:"batches,omitempty"`
	HeadId  string                                    `protobuf:"bytes,3,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	Paging  *client_list_control.ClientPagingResponse `protobuf:"bytes,4,opt,name=paging" json:"paging,omitempty"`
}

func (m *ClientBatchListResponse) Reset()                    { *m = ClientBatchListResponse{} }
func (m *ClientBatchListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchListResponse) ProtoMessage()               {}
func (*ClientBatchListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientBatchListResponse) GetStatus() ClientBatchListResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchListResponse_STATUS_UNSET
}

func (m *ClientBatchListResponse) GetBatches() []*batch.Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *ClientBatchListResponse) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientBatchListResponse) GetPaging() *client_list_control.ClientPagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

// Fetches a specific batch by its id (header_signature) from the blockchain.
type ClientBatchGetRequest struct {
	BatchId string `protobuf:"bytes,1,opt,name=batch_id,json=batchId" json:"batch_id,omitempty"`
}

func (m *ClientBatchGetRequest) Reset()                    { *m = ClientBatchGetRequest{} }
func (m *ClientBatchGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchGetRequest) ProtoMessage()               {}
func (*ClientBatchGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientBatchGetRequest) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

// A response that returns the batch specified by a ClientBatchGetRequest.
//
// Statuses:
//   * OK - everything worked as expected, batch has been fetched
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - no batch with the specified id exists
type ClientBatchGetResponse struct {
	Status ClientBatchGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientBatchGetResponse_Status" json:"status,omitempty"`
	Batch  *batch.Batch                  `protobuf:"bytes,2,opt,name=batch" json:"batch,omitempty"`
}

func (m *ClientBatchGetResponse) Reset()                    { *m = ClientBatchGetResponse{} }
func (m *ClientBatchGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchGetResponse) ProtoMessage()               {}
func (*ClientBatchGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientBatchGetResponse) GetStatus() ClientBatchGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchGetResponse_STATUS_UNSET
}

func (m *ClientBatchGetResponse) GetBatch() *batch.Batch {
	if m != nil {
		return m.Batch
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientBatchListRequest)(nil), "ClientBatchListRequest")
	proto.RegisterType((*ClientBatchListResponse)(nil), "ClientBatchListResponse")
	proto.RegisterType((*ClientBatchGetRequest)(nil), "ClientBatchGetRequest")
	proto.RegisterType((*ClientBatchGetResponse)(nil), "ClientBatchGetResponse")
	proto.RegisterEnum("ClientBatchListResponse_Status", ClientBatchListResponse_Status_name, ClientBatchListResponse_Status_value)
	proto.RegisterEnum("ClientBatchGetResponse_Status", ClientBatchGetResponse_Status_name, ClientBatchGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("sawtooth_sdk/protobuf/client_batch_pb2/client_batch.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xdd, 0x8e, 0x93, 0x40,
	0x14, 0xc7, 0xa5, 0xed, 0xc2, 0xf6, 0xa0, 0xeb, 0x64, 0x4c, 0xdd, 0xfa, 0x11, 0x6d, 0xb8, 0x32,
	0xd9, 0xc8, 0x26, 0x98, 0x68, 0xbc, 0x64, 0xbb, 0x84, 0x10, 0x1b, 0x68, 0x06, 0x6a, 0xa2, 0x37,
	0x13, 0x5a, 0x70, 0x4b, 0x6c, 0x4a, 0xed, 0x4c, 0xe3, 0x33, 0xf8, 0x2e, 0xde, 0xfb, 0x0a, 0x3e,
	0x96, 0x30, 0x30, 0x5b, 0x8a, 0xd5, 0x8b, 0xbd, 0xe4, 0x3f, 0xbf, 0xf3, 0xf5, 0x3f, 0x1c, 0x78,
	0xcf, 0xe2, 0xef, 0x3c, 0xcf, 0xf9, 0x92, 0xb2, 0xe4, 0xeb, 0xe5, 0x66, 0x9b, 0xf3, 0x7c, 0xbe,
	0xfb, 0x72, 0xb9, 0x58, 0x65, 0xe9, 0x9a, 0xd3, 0x79, 0xcc, 0x17, 0x4b, 0xba, 0x99, 0x5b, 0x07,
	0x82, 0x29, 0xb0, 0xa7, 0x17, 0xc7, 0x43, 0xf7, 0x31, 0x4d, 0xd8, 0xfd, 0x6f, 0x9d, 0x55, 0xc6,
	0x38, 0x5d, 0xe4, 0x6b, 0xbe, 0xcd, 0x57, 0xcd, 0x72, 0x4d, 0xbd, 0x4a, 0x64, 0xfc, 0x54, 0xe0,
	0xf1, 0x58, 0xbc, 0x5e, 0x95, 0xe9, 0x27, 0x05, 0x41, 0xd2, 0x6f, 0xbb, 0x94, 0x71, 0x7c, 0x0e,
	0xda, 0x32, 0x8d, 0x13, 0x9a, 0x25, 0x43, 0x65, 0xa4, 0xbc, 0xea, 0x13, 0xb5, 0xfc, 0xf4, 0x12,
	0xfc, 0x0c, 0xfa, 0x55, 0x57, 0x59, 0xc2, 0x86, 0x9d, 0x51, 0xb7, 0x78, 0x3a, 0x15, 0x82, 0x97,
	0x30, 0xfc, 0x1a, 0xd4, 0x4d, 0x7c, 0x93, 0xad, 0x6f, 0x86, 0xdd, 0x22, 0x48, 0xb7, 0x06, 0x66,
	0x95, 0x7e, 0x2a, 0xc4, 0x71, 0x55, 0x9c, 0x91, 0x1a, 0x2a, 0x70, 0x8d, 0xe5, 0x5b, 0x5e, 0xf2,
	0xbd, 0x22, 0x93, 0x6e, 0x3d, 0xaa, 0xf9, 0xb0, 0x50, 0x6f, 0x69, 0xc9, 0x18, 0xbf, 0x3b, 0x70,
	0xfe, 0x57, 0xbb, 0x6c, 0x93, 0xaf, 0x59, 0x8a, 0xdf, 0x81, 0xca, 0x78, 0xcc, 0x77, 0x4c, 0xb4,
	0x7b, 0x66, 0xbd, 0x34, 0xff, 0x41, 0x9a, 0xa1, 0xc0, 0x48, 0x8d, 0xe3, 0x11, 0x68, 0xa2, 0xfd,
	0xb4, 0x9a, 0x46, 0xb7, 0x54, 0x53, 0xc4, 0x10, 0x29, 0x37, 0xad, 0xe8, 0x1e, 0x58, 0xb1, 0x9f,
	0xb6, 0x77, 0x64, 0x5a, 0x59, 0x50, 0x4e, 0x6b, 0xfc, 0x50, 0x40, 0xad, 0x8a, 0x63, 0x04, 0xf7,
	0xc3, 0xc8, 0x8e, 0x66, 0x21, 0x9d, 0xf9, 0xa1, 0x13, 0xa1, 0x7b, 0x58, 0x85, 0x4e, 0xf0, 0x01,
	0x29, 0x18, 0xc3, 0x99, 0xe7, 0x47, 0x0e, 0xf1, 0xed, 0x09, 0x75, 0x08, 0x09, 0x08, 0xea, 0xe0,
	0x07, 0xd0, 0xf7, 0x83, 0x88, 0x12, 0xc7, 0xbe, 0xfe, 0x84, 0xba, 0x58, 0x07, 0xcd, 0x0f, 0x28,
	0x09, 0x82, 0x08, 0xf5, 0xf0, 0x43, 0xd0, 0xcb, 0x0f, 0x27, 0x0c, 0x66, 0x64, 0xec, 0xa0, 0x93,
	0x2a, 0xc1, 0x47, 0x7b, 0xe2, 0x5d, 0xd3, 0xa9, 0xed, 0x7a, 0xbe, 0x8b, 0xd4, 0xb2, 0x9c, 0xd4,
	0xc2, 0x80, 0x44, 0x48, 0x33, 0x2c, 0x18, 0x34, 0xfc, 0x71, 0xd3, 0xdb, 0xbd, 0x3f, 0x81, 0x53,
	0xb9, 0xde, 0x7a, 0xf1, 0x5a, 0xbd, 0x5d, 0xe3, 0xd7, 0xe1, 0xdf, 0x22, 0x82, 0x6a, 0xf7, 0xdf,
	0xb6, 0xdc, 0x7f, 0x61, 0x1e, 0x07, 0xdb, 0xe6, 0x3f, 0x87, 0x13, 0x91, 0xbd, 0xb0, 0x5e, 0x69,
	0x58, 0x5f, 0x89, 0x86, 0x7b, 0x47, 0xbf, 0xda, 0x9e, 0x5c, 0x5d, 0xc0, 0x40, 0x9e, 0x8c, 0x59,
	0x9c, 0x8c, 0x29, 0x4f, 0x66, 0xaa, 0x7c, 0x46, 0xed, 0xeb, 0x9c, 0xab, 0xe2, 0xf5, 0xcd, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x7f, 0xa2, 0x76, 0xce, 0x03, 0x00, 0x00,
}
