// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/client_block_pb2/client_block.proto

/*
Package client_block_pb2 is a generated protocol buffer package.

It is generated from these files:
	sawtooth_sdk/protobuf/client_block_pb2/client_block.proto

It has these top-level messages:
	ClientBlockListRequest
	ClientBlockListResponse
	ClientBlockGetRequest
	ClientBlockGetResponse
*/
package client_block_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import block "sawtooth_sdk/protobuf/block_pb2"
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

type ClientBlockListResponse_Status int32

const (
	ClientBlockListResponse_STATUS_UNSET   ClientBlockListResponse_Status = 0
	ClientBlockListResponse_OK             ClientBlockListResponse_Status = 1
	ClientBlockListResponse_INTERNAL_ERROR ClientBlockListResponse_Status = 2
	ClientBlockListResponse_NOT_READY      ClientBlockListResponse_Status = 3
	ClientBlockListResponse_NO_ROOT        ClientBlockListResponse_Status = 4
	ClientBlockListResponse_NO_RESOURCE    ClientBlockListResponse_Status = 5
	ClientBlockListResponse_INVALID_PAGING ClientBlockListResponse_Status = 6
	ClientBlockListResponse_INVALID_SORT   ClientBlockListResponse_Status = 7
)

var ClientBlockListResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_PAGING",
	7: "INVALID_SORT",
}
var ClientBlockListResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NOT_READY":      3,
	"NO_ROOT":        4,
	"NO_RESOURCE":    5,
	"INVALID_PAGING": 6,
	"INVALID_SORT":   7,
}

func (x ClientBlockListResponse_Status) String() string {
	return proto.EnumName(ClientBlockListResponse_Status_name, int32(x))
}
func (ClientBlockListResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ClientBlockGetResponse_Status int32

const (
	ClientBlockGetResponse_STATUS_UNSET   ClientBlockGetResponse_Status = 0
	ClientBlockGetResponse_OK             ClientBlockGetResponse_Status = 1
	ClientBlockGetResponse_INTERNAL_ERROR ClientBlockGetResponse_Status = 2
	ClientBlockGetResponse_NO_RESOURCE    ClientBlockGetResponse_Status = 5
)

var ClientBlockGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
}
var ClientBlockGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
}

func (x ClientBlockGetResponse_Status) String() string {
	return proto.EnumName(ClientBlockGetResponse_Status_name, int32(x))
}
func (ClientBlockGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// A request to return a list of blocks from the validator. May include the id
// of a particular block to be the `head` of the chain being requested. In that
// case the list will include that block (if found), and all blocks previous
// to it on the chain. Can be filtered using specific `block_ids`.
type ClientBlockListRequest struct {
	HeadId   string                                    `protobuf:"bytes,1,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	BlockIds []string                                  `protobuf:"bytes,2,rep,name=block_ids,json=blockIds" json:"block_ids,omitempty"`
	Paging   *client_list_control.ClientPagingControls `protobuf:"bytes,3,opt,name=paging" json:"paging,omitempty"`
	Sorting  []*client_list_control.ClientSortControls `protobuf:"bytes,4,rep,name=sorting" json:"sorting,omitempty"`
}

func (m *ClientBlockListRequest) Reset()                    { *m = ClientBlockListRequest{} }
func (m *ClientBlockListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientBlockListRequest) ProtoMessage()               {}
func (*ClientBlockListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientBlockListRequest) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientBlockListRequest) GetBlockIds() []string {
	if m != nil {
		return m.BlockIds
	}
	return nil
}

func (m *ClientBlockListRequest) GetPaging() *client_list_control.ClientPagingControls {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ClientBlockListRequest) GetSorting() []*client_list_control.ClientSortControls {
	if m != nil {
		return m.Sorting
	}
	return nil
}

// A response that lists a chain of blocks with the newest at the beginning,
// and the oldest (genesis) block at the end.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NOT_READY - the validator does not yet have a genesis block
//   * NO_ROOT - the head block specified was not found
//   * NO_RESOURCE - no blocks were found with the parameters specified
//   * INVALID_PAGING - the paging controls were malformed or out of range
//   * INVALID_SORT - the sorting controls were malformed or invalid
type ClientBlockListResponse struct {
	Status ClientBlockListResponse_Status            `protobuf:"varint,1,opt,name=status,enum=ClientBlockListResponse_Status" json:"status,omitempty"`
	Blocks []*block.Block                            `protobuf:"bytes,2,rep,name=blocks" json:"blocks,omitempty"`
	HeadId string                                    `protobuf:"bytes,3,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	Paging *client_list_control.ClientPagingResponse `protobuf:"bytes,4,opt,name=paging" json:"paging,omitempty"`
}

func (m *ClientBlockListResponse) Reset()                    { *m = ClientBlockListResponse{} }
func (m *ClientBlockListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientBlockListResponse) ProtoMessage()               {}
func (*ClientBlockListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientBlockListResponse) GetStatus() ClientBlockListResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBlockListResponse_STATUS_UNSET
}

func (m *ClientBlockListResponse) GetBlocks() []*block.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *ClientBlockListResponse) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientBlockListResponse) GetPaging() *client_list_control.ClientPagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

// A request to return a specific block from the validator. The block must be
// specified by its unique id, in this case the block's header signature
type ClientBlockGetRequest struct {
	BlockId  string `protobuf:"bytes,1,opt,name=block_id,json=blockId" json:"block_id,omitempty"`
	BlockNum uint64 `protobuf:"varint,2,opt,name=block_num,json=blockNum" json:"block_num,omitempty"`
}

func (m *ClientBlockGetRequest) Reset()                    { *m = ClientBlockGetRequest{} }
func (m *ClientBlockGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientBlockGetRequest) ProtoMessage()               {}
func (*ClientBlockGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientBlockGetRequest) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func (m *ClientBlockGetRequest) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

// A response that returns the block specified by a ClientBlockGetRequest.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - no block with the specified id exists
type ClientBlockGetResponse struct {
	Status ClientBlockGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientBlockGetResponse_Status" json:"status,omitempty"`
	Block  *block.Block                  `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *ClientBlockGetResponse) Reset()                    { *m = ClientBlockGetResponse{} }
func (m *ClientBlockGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientBlockGetResponse) ProtoMessage()               {}
func (*ClientBlockGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientBlockGetResponse) GetStatus() ClientBlockGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBlockGetResponse_STATUS_UNSET
}

func (m *ClientBlockGetResponse) GetBlock() *block.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientBlockListRequest)(nil), "ClientBlockListRequest")
	proto.RegisterType((*ClientBlockListResponse)(nil), "ClientBlockListResponse")
	proto.RegisterType((*ClientBlockGetRequest)(nil), "ClientBlockGetRequest")
	proto.RegisterType((*ClientBlockGetResponse)(nil), "ClientBlockGetResponse")
	proto.RegisterEnum("ClientBlockListResponse_Status", ClientBlockListResponse_Status_name, ClientBlockListResponse_Status_value)
	proto.RegisterEnum("ClientBlockGetResponse_Status", ClientBlockGetResponse_Status_name, ClientBlockGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("sawtooth_sdk/protobuf/client_block_pb2/client_block.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x8f, 0x93, 0x40,
	0x14, 0x95, 0xb6, 0x0b, 0xdb, 0x8b, 0xae, 0x93, 0x31, 0x75, 0xeb, 0x47, 0xd6, 0x0d, 0x4f, 0x9b,
	0x6c, 0x64, 0x13, 0x4c, 0x34, 0x3e, 0xb2, 0x5d, 0x42, 0x88, 0x0d, 0x34, 0x03, 0x35, 0xd1, 0x97,
	0x49, 0x5b, 0x70, 0x97, 0x6c, 0x97, 0xa9, 0x9d, 0x21, 0xfe, 0x06, 0xff, 0x8b, 0xef, 0xfe, 0x03,
	0x7f, 0x97, 0x30, 0xc0, 0x96, 0xd6, 0xea, 0x83, 0x8f, 0x9c, 0x39, 0xf7, 0x9c, 0x7b, 0xef, 0xe1,
	0xc2, 0x7b, 0x3e, 0xfb, 0x26, 0x18, 0x13, 0x37, 0x94, 0xc7, 0xb7, 0x17, 0xab, 0x35, 0x13, 0x6c,
	0x9e, 0x7f, 0xb9, 0x58, 0x2c, 0xd3, 0x24, 0x13, 0x74, 0xbe, 0x64, 0x8b, 0x5b, 0xba, 0x9a, 0x5b,
	0x5b, 0x80, 0x29, 0x69, 0xcf, 0xcf, 0xf7, 0x97, 0x6e, 0x6a, 0xda, 0x64, 0xf7, 0x9f, 0x3e, 0xcb,
	0x94, 0x0b, 0xba, 0x60, 0x99, 0x58, 0xb3, 0x65, 0xdb, 0xae, 0x8d, 0x57, 0x42, 0xc6, 0x0f, 0x05,
	0x9e, 0x8e, 0xe4, 0xeb, 0x65, 0x29, 0x3f, 0x2e, 0x18, 0x24, 0xf9, 0x9a, 0x27, 0x5c, 0xe0, 0x63,
	0xd0, 0x6e, 0x92, 0x59, 0x4c, 0xd3, 0x78, 0xa8, 0x9c, 0x2a, 0x67, 0x7d, 0xa2, 0x96, 0x9f, 0x5e,
	0x8c, 0x5f, 0x40, 0xbf, 0xea, 0x2a, 0x8d, 0xf9, 0xb0, 0x73, 0xda, 0x2d, 0x9e, 0x0e, 0x25, 0xe0,
	0xc5, 0x1c, 0xbf, 0x06, 0x75, 0x35, 0xbb, 0x4e, 0xb3, 0xeb, 0x61, 0xb7, 0x28, 0xd2, 0xad, 0x81,
	0x59, 0xc9, 0x4f, 0x24, 0x38, 0xaa, 0xcc, 0x39, 0xa9, 0x49, 0x05, 0x5d, 0xe3, 0x6c, 0x2d, 0x4a,
	0x7e, 0xaf, 0x50, 0xd2, 0xad, 0x27, 0x35, 0x3f, 0x2c, 0xd0, 0x7b, 0x76, 0xc3, 0x31, 0x7e, 0x75,
	0xe0, 0xf8, 0x8f, 0x76, 0xf9, 0x8a, 0x65, 0x3c, 0xc1, 0xef, 0x40, 0xe5, 0x62, 0x26, 0x72, 0x2e,
	0xdb, 0x3d, 0xb2, 0x5e, 0x99, 0x7f, 0x61, 0x9a, 0xa1, 0xa4, 0x91, 0x9a, 0x8e, 0x4f, 0x40, 0x95,
	0xed, 0x57, 0xc3, 0xe8, 0x96, 0x6a, 0xca, 0x12, 0x52, 0xa3, 0xed, 0x45, 0x74, 0xb7, 0x16, 0xb1,
	0x99, 0xb5, 0xb7, 0x67, 0xd6, 0xc6, 0xae, 0x99, 0xd5, 0xf8, 0xae, 0x80, 0x5a, 0x59, 0x63, 0x04,
	0x0f, 0xc3, 0xc8, 0x8e, 0xa6, 0x21, 0x9d, 0xfa, 0xa1, 0x13, 0xa1, 0x07, 0x58, 0x85, 0x4e, 0xf0,
	0x01, 0x29, 0x18, 0xc3, 0x91, 0xe7, 0x47, 0x0e, 0xf1, 0xed, 0x31, 0x75, 0x08, 0x09, 0x08, 0xea,
	0xe0, 0x47, 0xd0, 0xf7, 0x83, 0x88, 0x12, 0xc7, 0xbe, 0xfa, 0x84, 0xba, 0x58, 0x07, 0xcd, 0x0f,
	0x28, 0x09, 0x82, 0x08, 0xf5, 0xf0, 0x63, 0xd0, 0xcb, 0x0f, 0x27, 0x0c, 0xa6, 0x64, 0xe4, 0xa0,
	0x83, 0x4a, 0xe0, 0xa3, 0x3d, 0xf6, 0xae, 0xe8, 0xc4, 0x76, 0x3d, 0xdf, 0x45, 0x6a, 0x69, 0xd7,
	0x60, 0x61, 0x40, 0x22, 0xa4, 0x19, 0x01, 0x0c, 0x5a, 0xdb, 0x71, 0x93, 0xfb, 0xd4, 0x9f, 0xc1,
	0x61, 0x13, 0x6e, 0x1d, 0xbb, 0x56, 0x67, 0xbb, 0xc9, 0x3d, 0xcb, 0xef, 0x8a, 0x55, 0x29, 0x67,
	0xbd, 0x3a, 0x77, 0x3f, 0xbf, 0x33, 0x7e, 0x6e, 0xff, 0x48, 0x52, 0xb1, 0x0e, 0xe6, 0xed, 0x4e,
	0x30, 0x27, 0xe6, 0x7e, 0xe2, 0x6e, 0x2e, 0x2f, 0xe1, 0x40, 0xca, 0x4b, 0xaf, 0x4d, 0x2c, 0x15,
	0x68, 0xb8, 0xff, 0xb9, 0xcc, 0xdd, 0x85, 0x5d, 0x9e, 0xc3, 0xa0, 0xb9, 0x26, 0xb3, 0xb8, 0x26,
	0xb3, 0xb9, 0xa6, 0x89, 0xf2, 0x19, 0xed, 0x1e, 0xee, 0x5c, 0x95, 0xaf, 0x6f, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x17, 0x3c, 0x3b, 0x40, 0xe9, 0x03, 0x00, 0x00,
}
