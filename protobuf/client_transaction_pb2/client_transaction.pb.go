// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/client_transaction_pb2/client_transaction.proto

/*
Package client_transaction_pb2 is a generated protocol buffer package.

It is generated from these files:
	sawtooth_sdk/protobuf/client_transaction_pb2/client_transaction.proto

It has these top-level messages:
	ClientTransactionListRequest
	ClientTransactionListResponse
	ClientTransactionGetRequest
	ClientTransactionGetResponse
*/
package client_transaction_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import transaction "sawtooth_sdk/protobuf/transaction_pb2"
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

type ClientTransactionListResponse_Status int32

const (
	ClientTransactionListResponse_STATUS_UNSET   ClientTransactionListResponse_Status = 0
	ClientTransactionListResponse_OK             ClientTransactionListResponse_Status = 1
	ClientTransactionListResponse_INTERNAL_ERROR ClientTransactionListResponse_Status = 2
	ClientTransactionListResponse_NOT_READY      ClientTransactionListResponse_Status = 3
	ClientTransactionListResponse_NO_ROOT        ClientTransactionListResponse_Status = 4
	ClientTransactionListResponse_NO_RESOURCE    ClientTransactionListResponse_Status = 5
	ClientTransactionListResponse_INVALID_PAGING ClientTransactionListResponse_Status = 6
	ClientTransactionListResponse_INVALID_SORT   ClientTransactionListResponse_Status = 7
)

var ClientTransactionListResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_PAGING",
	7: "INVALID_SORT",
}
var ClientTransactionListResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NOT_READY":      3,
	"NO_ROOT":        4,
	"NO_RESOURCE":    5,
	"INVALID_PAGING": 6,
	"INVALID_SORT":   7,
}

func (x ClientTransactionListResponse_Status) String() string {
	return proto.EnumName(ClientTransactionListResponse_Status_name, int32(x))
}
func (ClientTransactionListResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ClientTransactionGetResponse_Status int32

const (
	ClientTransactionGetResponse_STATUS_UNSET   ClientTransactionGetResponse_Status = 0
	ClientTransactionGetResponse_OK             ClientTransactionGetResponse_Status = 1
	ClientTransactionGetResponse_INTERNAL_ERROR ClientTransactionGetResponse_Status = 2
	ClientTransactionGetResponse_NO_RESOURCE    ClientTransactionGetResponse_Status = 5
)

var ClientTransactionGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
}
var ClientTransactionGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
}

func (x ClientTransactionGetResponse_Status) String() string {
	return proto.EnumName(ClientTransactionGetResponse_Status_name, int32(x))
}
func (ClientTransactionGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// A request to return a list of txns from the validator. May include the id
// of a particular block to be the `head` of the chain being requested. In that
// case the list will include the txns from that block, and all txns
// previous to that block on the chain. Filter with specific `transaction_ids`.
type ClientTransactionListRequest struct {
	HeadId         string                                    `protobuf:"bytes,1,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	TransactionIds []string                                  `protobuf:"bytes,2,rep,name=transaction_ids,json=transactionIds" json:"transaction_ids,omitempty"`
	Paging         *client_list_control.ClientPagingControls `protobuf:"bytes,3,opt,name=paging" json:"paging,omitempty"`
	Sorting        []*client_list_control.ClientSortControls `protobuf:"bytes,4,rep,name=sorting" json:"sorting,omitempty"`
}

func (m *ClientTransactionListRequest) Reset()                    { *m = ClientTransactionListRequest{} }
func (m *ClientTransactionListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientTransactionListRequest) ProtoMessage()               {}
func (*ClientTransactionListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientTransactionListRequest) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientTransactionListRequest) GetTransactionIds() []string {
	if m != nil {
		return m.TransactionIds
	}
	return nil
}

func (m *ClientTransactionListRequest) GetPaging() *client_list_control.ClientPagingControls {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ClientTransactionListRequest) GetSorting() []*client_list_control.ClientSortControls {
	if m != nil {
		return m.Sorting
	}
	return nil
}

// A response that lists transactions from newest to oldest.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NOT_READY - the validator does not yet have a genesis block
//   * NO_ROOT - the head block specified was not found
//   * NO_RESOURCE - no txns were found with the parameters specified
//   * INVALID_PAGING - the paging controls were malformed or out of range
//   * INVALID_SORT - the sorting controls were malformed or invalid
type ClientTransactionListResponse struct {
	Status       ClientTransactionListResponse_Status      `protobuf:"varint,1,opt,name=status,enum=ClientTransactionListResponse_Status" json:"status,omitempty"`
	Transactions []*transaction.Transaction                `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	HeadId       string                                    `protobuf:"bytes,3,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	Paging       *client_list_control.ClientPagingResponse `protobuf:"bytes,4,opt,name=paging" json:"paging,omitempty"`
}

func (m *ClientTransactionListResponse) Reset()                    { *m = ClientTransactionListResponse{} }
func (m *ClientTransactionListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientTransactionListResponse) ProtoMessage()               {}
func (*ClientTransactionListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientTransactionListResponse) GetStatus() ClientTransactionListResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientTransactionListResponse_STATUS_UNSET
}

func (m *ClientTransactionListResponse) GetTransactions() []*transaction.Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *ClientTransactionListResponse) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientTransactionListResponse) GetPaging() *client_list_control.ClientPagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

// Fetches a specific txn by its id (header_signature) from the blockchain.
type ClientTransactionGetRequest struct {
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
}

func (m *ClientTransactionGetRequest) Reset()                    { *m = ClientTransactionGetRequest{} }
func (m *ClientTransactionGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientTransactionGetRequest) ProtoMessage()               {}
func (*ClientTransactionGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientTransactionGetRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

// A response that returns the txn specified by a ClientTransactionGetRequest.
//
// Statuses:
//   * OK - everything worked as expected, txn has been fetched
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - no txn with the specified id exists
type ClientTransactionGetResponse struct {
	Status      ClientTransactionGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientTransactionGetResponse_Status" json:"status,omitempty"`
	Transaction *transaction.Transaction            `protobuf:"bytes,2,opt,name=transaction" json:"transaction,omitempty"`
	// Identifier of the block this transaction is in. Not set if the
	// transaction is pending.
	BlockId string `protobuf:"bytes,3,opt,name=block_id,json=blockId" json:"block_id,omitempty"`
}

func (m *ClientTransactionGetResponse) Reset()                    { *m = ClientTransactionGetResponse{} }
func (m *ClientTransactionGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientTransactionGetResponse) ProtoMessage()               {}
func (*ClientTransactionGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientTransactionGetResponse) GetStatus() ClientTransactionGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientTransactionGetResponse_STATUS_UNSET
}

func (m *ClientTransactionGetResponse) GetTransaction() *transaction.Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *ClientTransactionGetResponse) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientTransactionListRequest)(nil), "ClientTransactionListRequest")
	proto.RegisterType((*ClientTransactionListResponse)(nil), "ClientTransactionListResponse")
	proto.RegisterType((*ClientTransactionGetRequest)(nil), "ClientTransactionGetRequest")
	proto.RegisterType((*ClientTransactionGetResponse)(nil), "ClientTransactionGetResponse")
	proto.RegisterEnum("ClientTransactionListResponse_Status", ClientTransactionListResponse_Status_name, ClientTransactionListResponse_Status_value)
	proto.RegisterEnum("ClientTransactionGetResponse_Status", ClientTransactionGetResponse_Status_name, ClientTransactionGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("sawtooth_sdk/protobuf/client_transaction_pb2/client_transaction.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x8e, 0x93, 0x40,
	0x14, 0xc7, 0xa5, 0xad, 0xe0, 0x1e, 0xba, 0x5d, 0x32, 0x66, 0xb5, 0x7e, 0x25, 0x1b, 0xe2, 0x46,
	0x6f, 0x64, 0x15, 0x2f, 0xbc, 0xd1, 0x0b, 0x6c, 0x09, 0x21, 0x36, 0xd0, 0x0c, 0xd4, 0x44, 0x6f,
	0x26, 0xb4, 0xe0, 0x2e, 0xd9, 0x86, 0xa9, 0x9d, 0x69, 0x7c, 0x06, 0x5f, 0xca, 0xa7, 0xf1, 0x01,
	0x7c, 0x04, 0x29, 0x1f, 0xed, 0xb4, 0xc5, 0x5e, 0x78, 0x39, 0x87, 0xdf, 0xf9, 0xfa, 0x9f, 0x73,
	0x00, 0x9b, 0x45, 0x3f, 0x38, 0xa5, 0xfc, 0x86, 0xb0, 0xf8, 0xf6, 0x6a, 0xb1, 0xa4, 0x9c, 0x4e,
	0x57, 0xdf, 0xae, 0x66, 0xf3, 0x34, 0xc9, 0x38, 0xe1, 0xcb, 0x28, 0x63, 0xd1, 0x8c, 0xa7, 0x34,
	0x23, 0x8b, 0xa9, 0xd9, 0x60, 0x36, 0x0a, 0x97, 0xc7, 0xef, 0x9a, 0xc3, 0xec, 0xfb, 0x1f, 0x3a,
	0x3a, 0x47, 0xf3, 0xcf, 0x53, 0xc6, 0xc9, 0x8c, 0x66, 0x7c, 0x49, 0xe7, 0x62, 0x01, 0xa2, 0xbd,
	0x0c, 0xa4, 0xff, 0x92, 0xe0, 0xe9, 0xa0, 0xf8, 0x1a, 0x6e, 0x93, 0x8c, 0x72, 0x0e, 0x27, 0xdf,
	0x57, 0x09, 0xe3, 0xe8, 0x21, 0x28, 0x37, 0x49, 0x14, 0x93, 0x34, 0xee, 0x4b, 0x17, 0xd2, 0xcb,
	0x13, 0x2c, 0xaf, 0x9f, 0x6e, 0x8c, 0x5e, 0xc0, 0x99, 0x58, 0x67, 0x1a, 0xb3, 0x7e, 0xeb, 0xa2,
	0x9d, 0x03, 0x3d, 0xc1, 0xec, 0xc6, 0x0c, 0xbd, 0x02, 0x79, 0x11, 0x5d, 0xa7, 0xd9, 0x75, 0xbf,
	0x9d, 0x07, 0x50, 0xcd, 0x73, 0xa3, 0x4c, 0x38, 0x2e, 0x8c, 0x83, 0xb2, 0x1c, 0x86, 0x2b, 0x28,
	0xc7, 0x15, 0x46, 0x97, 0x7c, 0xcd, 0x77, 0xf2, 0x78, 0xaa, 0x79, 0xbf, 0xe2, 0x83, 0xdc, 0xba,
	0xa1, 0x6b, 0x46, 0xff, 0xdd, 0x82, 0x67, 0xff, 0x68, 0x80, 0x2d, 0x68, 0xc6, 0x12, 0xf4, 0x01,
	0x64, 0xc6, 0x23, 0xbe, 0x62, 0x45, 0x03, 0x3d, 0xf3, 0xd2, 0x38, 0xca, 0x1b, 0x41, 0x01, 0xe3,
	0xca, 0x09, 0xbd, 0x86, 0xae, 0xd0, 0x50, 0xd9, 0xa4, 0x6a, 0x76, 0x0d, 0xc1, 0x1d, 0xef, 0x10,
	0xa2, 0x64, 0xed, 0x1d, 0xc9, 0xb6, 0x4a, 0x74, 0x1a, 0x94, 0xa8, 0x0b, 0xa8, 0x95, 0xd0, 0x7f,
	0x4a, 0x20, 0x97, 0xc5, 0x20, 0x0d, 0xba, 0x41, 0x68, 0x85, 0x93, 0x80, 0x4c, 0xbc, 0xc0, 0x0e,
	0xb5, 0x3b, 0x48, 0x86, 0x96, 0xff, 0x49, 0x93, 0x10, 0x82, 0x9e, 0xeb, 0x85, 0x36, 0xf6, 0xac,
	0x11, 0xb1, 0x31, 0xf6, 0xb1, 0xd6, 0x42, 0xa7, 0x70, 0xe2, 0xf9, 0x21, 0xc1, 0xb6, 0x35, 0xfc,
	0xa2, 0xb5, 0x91, 0x0a, 0x8a, 0xe7, 0x13, 0xec, 0xfb, 0xa1, 0xd6, 0x41, 0x67, 0xa0, 0xae, 0x1f,
	0x76, 0xe0, 0x4f, 0xf0, 0xc0, 0xd6, 0xee, 0x96, 0x01, 0x3e, 0x5b, 0x23, 0x77, 0x48, 0xc6, 0x96,
	0xe3, 0x7a, 0x8e, 0x26, 0xaf, 0xd3, 0xd5, 0xb6, 0xc0, 0xc7, 0xa1, 0xa6, 0xe8, 0x43, 0x78, 0x72,
	0xa0, 0x9a, 0x93, 0x6c, 0xb6, 0xe4, 0x12, 0x7a, 0xbb, 0xcb, 0x50, 0x2d, 0xcb, 0xe9, 0xce, 0x2e,
	0xe8, 0x7f, 0x9a, 0xb6, 0xad, 0x08, 0x53, 0xcd, 0xea, 0xfd, 0xde, 0xac, 0x9e, 0x1b, 0xc7, 0xf0,
	0xfd, 0x51, 0x19, 0xa0, 0x0a, 0xf9, 0xf2, 0x49, 0x49, 0x07, 0x93, 0x12, 0x01, 0xf4, 0x08, 0xee,
	0x4d, 0xe7, 0x74, 0x76, 0xbb, 0x9d, 0x94, 0x52, 0xbc, 0xf3, 0x4a, 0x9d, 0xff, 0x94, 0x7e, 0x5f,
	0xde, 0x8f, 0x6f, 0xe0, 0xbc, 0xbe, 0x55, 0x23, 0xbf, 0x55, 0xa3, 0xbe, 0xd5, 0xb1, 0xf4, 0xf5,
	0x41, 0xf3, 0xef, 0x62, 0x2a, 0x17, 0xcc, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0xef,
	0x46, 0xe7, 0x65, 0x04, 0x00, 0x00,
}
