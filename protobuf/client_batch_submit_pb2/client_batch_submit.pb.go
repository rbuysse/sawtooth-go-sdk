// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/client_batch_submit_pb2/client_batch_submit.proto

/*
Package client_batch_submit_pb2 is a generated protocol buffer package.

It is generated from these files:
	sawtooth_sdk/protobuf/client_batch_submit_pb2/client_batch_submit.proto

It has these top-level messages:
	ClientBatchStatus
	ClientBatchSubmitRequest
	ClientBatchSubmitResponse
	ClientBatchStatusRequest
	ClientBatchStatusResponse
*/
package client_batch_submit_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import batch "sawtooth_sdk/protobuf/batch_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientBatchStatus_Status int32

const (
	ClientBatchStatus_STATUS_UNSET ClientBatchStatus_Status = 0
	ClientBatchStatus_COMMITTED    ClientBatchStatus_Status = 1
	ClientBatchStatus_INVALID      ClientBatchStatus_Status = 2
	ClientBatchStatus_PENDING      ClientBatchStatus_Status = 3
	ClientBatchStatus_UNKNOWN      ClientBatchStatus_Status = 4
)

var ClientBatchStatus_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "COMMITTED",
	2: "INVALID",
	3: "PENDING",
	4: "UNKNOWN",
}
var ClientBatchStatus_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"COMMITTED":    1,
	"INVALID":      2,
	"PENDING":      3,
	"UNKNOWN":      4,
}

func (x ClientBatchStatus_Status) String() string {
	return proto.EnumName(ClientBatchStatus_Status_name, int32(x))
}
func (ClientBatchStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ClientBatchSubmitResponse_Status int32

const (
	ClientBatchSubmitResponse_STATUS_UNSET   ClientBatchSubmitResponse_Status = 0
	ClientBatchSubmitResponse_OK             ClientBatchSubmitResponse_Status = 1
	ClientBatchSubmitResponse_INTERNAL_ERROR ClientBatchSubmitResponse_Status = 2
	ClientBatchSubmitResponse_INVALID_BATCH  ClientBatchSubmitResponse_Status = 3
)

var ClientBatchSubmitResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "INVALID_BATCH",
}
var ClientBatchSubmitResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"INVALID_BATCH":  3,
}

func (x ClientBatchSubmitResponse_Status) String() string {
	return proto.EnumName(ClientBatchSubmitResponse_Status_name, int32(x))
}
func (ClientBatchSubmitResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

type ClientBatchStatusResponse_Status int32

const (
	ClientBatchStatusResponse_STATUS_UNSET   ClientBatchStatusResponse_Status = 0
	ClientBatchStatusResponse_OK             ClientBatchStatusResponse_Status = 1
	ClientBatchStatusResponse_INTERNAL_ERROR ClientBatchStatusResponse_Status = 2
	ClientBatchStatusResponse_NO_RESOURCE    ClientBatchStatusResponse_Status = 5
)

var ClientBatchStatusResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
}
var ClientBatchStatusResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
}

func (x ClientBatchStatusResponse_Status) String() string {
	return proto.EnumName(ClientBatchStatusResponse_Status_name, int32(x))
}
func (ClientBatchStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

// Information about the status of a batch submitted to the validator.
//
// Attributes:
//     batch_id: The id (header_signature) of the batch
//     status: The committed status of the batch
//     invalid_transactions: Info for transactions that failed, if any
//
// Statuses:
//     COMMITTED - the batch was accepted and has been committed to the chain
//     INVALID - the batch failed validation, it should be resubmitted
//     PENDING - the batch is still being processed
//     UNKNOWN - no status for the batch could be found (possibly invalid)
type ClientBatchStatus struct {
	BatchId             string                                  `protobuf:"bytes,1,opt,name=batch_id,json=batchId" json:"batch_id,omitempty"`
	Status              ClientBatchStatus_Status                `protobuf:"varint,2,opt,name=status,enum=ClientBatchStatus_Status" json:"status,omitempty"`
	InvalidTransactions []*ClientBatchStatus_InvalidTransaction `protobuf:"bytes,3,rep,name=invalid_transactions,json=invalidTransactions" json:"invalid_transactions,omitempty"`
}

func (m *ClientBatchStatus) Reset()                    { *m = ClientBatchStatus{} }
func (m *ClientBatchStatus) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchStatus) ProtoMessage()               {}
func (*ClientBatchStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientBatchStatus) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

func (m *ClientBatchStatus) GetStatus() ClientBatchStatus_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchStatus_STATUS_UNSET
}

func (m *ClientBatchStatus) GetInvalidTransactions() []*ClientBatchStatus_InvalidTransaction {
	if m != nil {
		return m.InvalidTransactions
	}
	return nil
}

type ClientBatchStatus_InvalidTransaction struct {
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	Message       string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	ExtendedData  []byte `protobuf:"bytes,3,opt,name=extended_data,json=extendedData,proto3" json:"extended_data,omitempty"`
}

func (m *ClientBatchStatus_InvalidTransaction) Reset()         { *m = ClientBatchStatus_InvalidTransaction{} }
func (m *ClientBatchStatus_InvalidTransaction) String() string { return proto.CompactTextString(m) }
func (*ClientBatchStatus_InvalidTransaction) ProtoMessage()    {}
func (*ClientBatchStatus_InvalidTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *ClientBatchStatus_InvalidTransaction) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *ClientBatchStatus_InvalidTransaction) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ClientBatchStatus_InvalidTransaction) GetExtendedData() []byte {
	if m != nil {
		return m.ExtendedData
	}
	return nil
}

// Submits a list of Batches to be added to the blockchain.
type ClientBatchSubmitRequest struct {
	Batches []*batch.Batch `protobuf:"bytes,1,rep,name=batches" json:"batches,omitempty"`
}

func (m *ClientBatchSubmitRequest) Reset()                    { *m = ClientBatchSubmitRequest{} }
func (m *ClientBatchSubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchSubmitRequest) ProtoMessage()               {}
func (*ClientBatchSubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientBatchSubmitRequest) GetBatches() []*batch.Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

// This is a response to a submission of one or more Batches.
// Statuses:
//   * OK - everything with the request worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * INVALID_BATCH - the batch failed validation, likely due to a bad signature
type ClientBatchSubmitResponse struct {
	Status ClientBatchSubmitResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientBatchSubmitResponse_Status" json:"status,omitempty"`
}

func (m *ClientBatchSubmitResponse) Reset()                    { *m = ClientBatchSubmitResponse{} }
func (m *ClientBatchSubmitResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchSubmitResponse) ProtoMessage()               {}
func (*ClientBatchSubmitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientBatchSubmitResponse) GetStatus() ClientBatchSubmitResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchSubmitResponse_STATUS_UNSET
}

// A request for the status of one or more batches, specified by id.
// If `wait` is set to true, the validator will wait to respond until all
// batches are committed, or until the specified `timeout` in seconds has
// elapsed. Defaults to 300.
type ClientBatchStatusRequest struct {
	BatchIds []string `protobuf:"bytes,1,rep,name=batch_ids,json=batchIds" json:"batch_ids,omitempty"`
	Wait     bool     `protobuf:"varint,2,opt,name=wait" json:"wait,omitempty"`
	Timeout  uint32   `protobuf:"varint,3,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ClientBatchStatusRequest) Reset()                    { *m = ClientBatchStatusRequest{} }
func (m *ClientBatchStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchStatusRequest) ProtoMessage()               {}
func (*ClientBatchStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientBatchStatusRequest) GetBatchIds() []string {
	if m != nil {
		return m.BatchIds
	}
	return nil
}

func (m *ClientBatchStatusRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

func (m *ClientBatchStatusRequest) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// This is a response to a request for the status of specific batches.
// Statuses:
//   * OK - everything with the request worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - the response contains no data, likely because
//     no ids were specified in the request
type ClientBatchStatusResponse struct {
	Status        ClientBatchStatusResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientBatchStatusResponse_Status" json:"status,omitempty"`
	BatchStatuses []*ClientBatchStatus             `protobuf:"bytes,2,rep,name=batch_statuses,json=batchStatuses" json:"batch_statuses,omitempty"`
}

func (m *ClientBatchStatusResponse) Reset()                    { *m = ClientBatchStatusResponse{} }
func (m *ClientBatchStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientBatchStatusResponse) ProtoMessage()               {}
func (*ClientBatchStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClientBatchStatusResponse) GetStatus() ClientBatchStatusResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchStatusResponse_STATUS_UNSET
}

func (m *ClientBatchStatusResponse) GetBatchStatuses() []*ClientBatchStatus {
	if m != nil {
		return m.BatchStatuses
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientBatchStatus)(nil), "ClientBatchStatus")
	proto.RegisterType((*ClientBatchStatus_InvalidTransaction)(nil), "ClientBatchStatus.InvalidTransaction")
	proto.RegisterType((*ClientBatchSubmitRequest)(nil), "ClientBatchSubmitRequest")
	proto.RegisterType((*ClientBatchSubmitResponse)(nil), "ClientBatchSubmitResponse")
	proto.RegisterType((*ClientBatchStatusRequest)(nil), "ClientBatchStatusRequest")
	proto.RegisterType((*ClientBatchStatusResponse)(nil), "ClientBatchStatusResponse")
	proto.RegisterEnum("ClientBatchStatus_Status", ClientBatchStatus_Status_name, ClientBatchStatus_Status_value)
	proto.RegisterEnum("ClientBatchSubmitResponse_Status", ClientBatchSubmitResponse_Status_name, ClientBatchSubmitResponse_Status_value)
	proto.RegisterEnum("ClientBatchStatusResponse_Status", ClientBatchStatusResponse_Status_name, ClientBatchStatusResponse_Status_value)
}

func init() {
	proto.RegisterFile("sawtooth_sdk/protobuf/client_batch_submit_pb2/client_batch_submit.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0xb5, 0xb0, 0xf2, 0x71, 0xa1, 0xd8, 0x1d, 0x35, 0xc2, 0xfa, 0x82, 0x35, 0x9b, 0x90, 0x98,
	0x74, 0x23, 0x3e, 0x6d, 0xe2, 0x0b, 0x1f, 0x0d, 0x36, 0xbb, 0x3b, 0x25, 0x43, 0x51, 0xe3, 0x4b,
	0x33, 0xa5, 0xa3, 0x36, 0x2e, 0x14, 0x77, 0x06, 0xd7, 0xf8, 0x0f, 0xfc, 0x19, 0xfe, 0x2d, 0x7f,
	0x8d, 0xd3, 0x69, 0x2b, 0x2c, 0xb0, 0xc6, 0xec, 0x13, 0xbd, 0xa7, 0xe7, 0xde, 0xe1, 0x9c, 0x33,
	0xb7, 0x30, 0xe2, 0xf4, 0x5a, 0xc4, 0xb1, 0xf8, 0xec, 0xf3, 0xf0, 0xcb, 0xc9, 0xf2, 0x2a, 0x16,
	0x71, 0xb0, 0xfa, 0x78, 0x32, 0xbb, 0x8c, 0xd8, 0x42, 0xf8, 0x01, 0x15, 0x33, 0xf9, 0x66, 0x15,
	0xcc, 0x23, 0xe1, 0x2f, 0x83, 0xee, 0x3e, 0xdc, 0x52, 0x4d, 0x47, 0x2f, 0xf6, 0x0f, 0x4a, 0x99,
	0x49, 0xab, 0x7a, 0x4a, 0xc9, 0xe6, 0xcf, 0x22, 0x1c, 0x0e, 0xd4, 0xa8, 0x7e, 0x82, 0x4e, 0x04,
	0x15, 0x2b, 0x8e, 0x5a, 0x50, 0x49, 0xe9, 0x51, 0xd8, 0xd4, 0xda, 0x5a, 0xa7, 0x4a, 0xca, 0xaa,
	0x76, 0x42, 0xf4, 0x12, 0x4a, 0x5c, 0x91, 0x9a, 0x05, 0xf9, 0xa2, 0xd1, 0x6d, 0x59, 0x3b, 0xed,
	0x56, 0xfa, 0x43, 0x32, 0x22, 0x7a, 0x0f, 0x8f, 0xa2, 0xc5, 0x37, 0x7a, 0x19, 0x85, 0xbe, 0xb8,
	0xa2, 0x0b, 0x4e, 0x67, 0x22, 0x8a, 0x17, 0xbc, 0x59, 0x6c, 0x17, 0x3b, 0xb5, 0xee, 0xf1, 0x9e,
	0x01, 0x4e, 0x4a, 0xf7, 0xd6, 0x6c, 0xf2, 0x30, 0xda, 0xc1, 0xf8, 0xd1, 0x0f, 0x40, 0xbb, 0x54,
	0x74, 0x0c, 0x8d, 0x8d, 0x73, 0xd6, 0x1a, 0xf4, 0x0d, 0x54, 0x2a, 0x69, 0x42, 0x79, 0xce, 0x38,
	0xa7, 0x9f, 0x98, 0x92, 0x22, 0x35, 0x66, 0x25, 0x7a, 0x0e, 0x3a, 0xfb, 0x2e, 0xd8, 0x22, 0x64,
	0xa1, 0x1f, 0x52, 0x41, 0xe5, 0x3f, 0xd5, 0x3a, 0x75, 0x52, 0xcf, 0xc1, 0xa1, 0xc4, 0xcc, 0x31,
	0x94, 0x32, 0xb7, 0x0c, 0xa8, 0x4f, 0xbc, 0x9e, 0x37, 0x9d, 0xf8, 0x53, 0x3c, 0xb1, 0x3d, 0xe3,
	0x1e, 0xd2, 0xa1, 0x3a, 0x70, 0x2f, 0x2e, 0x1c, 0xcf, 0xb3, 0x87, 0x86, 0x86, 0x6a, 0x50, 0x76,
	0xf0, 0xdb, 0xde, 0xb9, 0x33, 0x34, 0x0a, 0x49, 0x31, 0xb6, 0xf1, 0xd0, 0xc1, 0x23, 0xa3, 0x98,
	0x14, 0x53, 0x7c, 0x86, 0xdd, 0x77, 0xd8, 0x38, 0x30, 0x5f, 0x43, 0x73, 0xd3, 0x0a, 0x95, 0x29,
	0x61, 0x5f, 0x57, 0x8c, 0x0b, 0xd4, 0x86, 0x34, 0x01, 0xc6, 0xa5, 0x98, 0xc4, 0xb6, 0x92, 0xa5,
	0x58, 0x24, 0x87, 0xcd, 0x5f, 0x1a, 0xb4, 0xf6, 0xb4, 0xf3, 0xa5, 0x34, 0x8a, 0xa1, 0xd3, 0xbf,
	0xb1, 0x69, 0x2a, 0xb6, 0x67, 0xd6, 0xad, 0xdc, 0xad, 0xf8, 0x4c, 0xe7, 0x1f, 0x42, 0x4b, 0x50,
	0x70, 0xcf, 0xa4, 0x42, 0x04, 0x0d, 0x07, 0x7b, 0x36, 0xc1, 0xbd, 0x73, 0xdf, 0x26, 0xc4, 0x25,
	0x52, 0xe8, 0x21, 0xe8, 0x99, 0x6a, 0xbf, 0xdf, 0xf3, 0x06, 0x6f, 0x8c, 0xa2, 0xc9, 0x6e, 0x2a,
	0x4c, 0xcf, 0xc9, 0x14, 0x3e, 0x85, 0x6a, 0x7e, 0xe7, 0x52, 0x8d, 0x55, 0x52, 0xc9, 0x2e, 0x1d,
	0x97, 0xf3, 0x0f, 0xae, 0x69, 0x24, 0x54, 0x50, 0x15, 0xa2, 0x9e, 0x93, 0xfc, 0x44, 0x34, 0x67,
	0xf1, 0x4a, 0xa8, 0x7c, 0x74, 0x92, 0x97, 0xe6, 0xef, 0x2d, 0x2b, 0xb2, 0x73, 0xfe, 0xcb, 0x8a,
	0x1b, 0xdc, 0xed, 0x9b, 0x7c, 0x0a, 0x8d, 0x6c, 0xe1, 0x54, 0xcd, 0x92, 0x25, 0x48, 0xc2, 0x40,
	0x7b, 0x46, 0xe8, 0xc1, 0xba, 0x90, 0xf1, 0x8c, 0xee, 0xe8, 0xe2, 0x03, 0xa8, 0x61, 0xd7, 0x27,
	0xf6, 0xc4, 0x9d, 0x92, 0x81, 0x6d, 0xdc, 0xef, 0x77, 0xe1, 0x71, 0xbe, 0xe0, 0x96, 0x5c, 0x70,
	0x2b, 0x5f, 0xf0, 0xb1, 0xf6, 0xe1, 0xc9, 0x2d, 0x1f, 0x8b, 0xa0, 0xa4, 0x48, 0xaf, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xd4, 0xa2, 0x33, 0xa4, 0x64, 0x04, 0x00, 0x00,
}
