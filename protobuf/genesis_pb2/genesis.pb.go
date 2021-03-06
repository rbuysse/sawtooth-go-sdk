// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/genesis_pb2/genesis.proto

/*
Package genesis_pb2 is a generated protocol buffer package.

It is generated from these files:
	sawtooth_sdk/protobuf/genesis_pb2/genesis.proto

It has these top-level messages:
	GenesisData
*/
package genesis_pb2

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

type GenesisData struct {
	// The list of batches that will be applied during the genesis process
	Batches []*batch.Batch `protobuf:"bytes,1,rep,name=batches" json:"batches,omitempty"`
}

func (m *GenesisData) Reset()                    { *m = GenesisData{} }
func (m *GenesisData) String() string            { return proto.CompactTextString(m) }
func (*GenesisData) ProtoMessage()               {}
func (*GenesisData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GenesisData) GetBatches() []*batch.Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisData)(nil), "GenesisData")
}

func init() { proto.RegisterFile("sawtooth_sdk/protobuf/genesis_pb2/genesis.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x4e, 0x2c, 0x2f,
	0xc9, 0xcf, 0x2f, 0xc9, 0x88, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a,
	0x4d, 0xd3, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0x8e, 0x2f, 0x48, 0x32, 0x82, 0xb1, 0xf5,
	0xc0, 0x92, 0x52, 0xda, 0xd8, 0x35, 0x24, 0x25, 0x96, 0x24, 0x67, 0x80, 0x95, 0x83, 0x59, 0x10,
	0xc5, 0x4a, 0xfa, 0x5c, 0xdc, 0xee, 0x10, 0xdd, 0x2e, 0x89, 0x25, 0x89, 0x42, 0x0a, 0x5c, 0xec,
	0x60, 0xd9, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x36, 0x3d, 0x27, 0x10, 0x3f,
	0x08, 0x26, 0xec, 0xa4, 0xc6, 0x25, 0x0a, 0x33, 0x5f, 0x0f, 0x68, 0xbe, 0x1e, 0xcc, 0xfc, 0x00,
	0xc6, 0x28, 0x6e, 0x24, 0x37, 0x25, 0xb1, 0x81, 0x25, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x52, 0x92, 0x8a, 0x21, 0xbf, 0x00, 0x00, 0x00,
}
