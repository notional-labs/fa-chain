// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fachain/interchainquery/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Query struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ChainId      string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	QueryType    string `protobuf:"bytes,4,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	Request      []byte `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	CallbackId   string `protobuf:"bytes,8,opt,name=callback_id,json=callbackId,proto3" json:"callback_id,omitempty"`
	Ttl          uint64 `protobuf:"varint,9,opt,name=ttl,proto3" json:"ttl,omitempty"`
	RequestSent  bool   `protobuf:"varint,11,opt,name=request_sent,json=requestSent,proto3" json:"request_sent,omitempty"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a9dcb51083ce376, []int{0}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Query.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return m.Size()
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Query) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *Query) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *Query) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Query) GetCallbackId() string {
	if m != nil {
		return m.CallbackId
	}
	return ""
}

func (m *Query) GetTtl() uint64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Query) GetRequestSent() bool {
	if m != nil {
		return m.RequestSent
	}
	return false
}

type DataPoint struct {
	Id           string                                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RemoteHeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=remote_height,json=remoteHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"remote_height"`
	LocalHeight  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=local_height,json=localHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"local_height"`
	Value        []byte                                 `protobuf:"bytes,4,opt,name=value,proto3" json:"result,omitempty"`
}

func (m *DataPoint) Reset()         { *m = DataPoint{} }
func (m *DataPoint) String() string { return proto.CompactTextString(m) }
func (*DataPoint) ProtoMessage()    {}
func (*DataPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a9dcb51083ce376, []int{1}
}
func (m *DataPoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataPoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPoint.Merge(m, src)
}
func (m *DataPoint) XXX_Size() int {
	return m.Size()
}
func (m *DataPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPoint.DiscardUnknown(m)
}

var xxx_messageInfo_DataPoint proto.InternalMessageInfo

func (m *DataPoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataPoint) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// GenesisState defines the epochs module's genesis state.
type GenesisState struct {
	Queries []Query `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a9dcb51083ce376, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetQueries() []Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func init() {
	proto.RegisterType((*Query)(nil), "fachain.interchainquery.v1.Query")
	proto.RegisterType((*DataPoint)(nil), "fachain.interchainquery.v1.DataPoint")
	proto.RegisterType((*GenesisState)(nil), "fachain.interchainquery.v1.GenesisState")
}

func init() {
	proto.RegisterFile("fachain/interchainquery/v1/genesis.proto", fileDescriptor_0a9dcb51083ce376)
}

var fileDescriptor_0a9dcb51083ce376 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xbd, 0x8e, 0xd3, 0x40,
	0x10, 0xce, 0xe6, 0x87, 0x24, 0x1b, 0x1f, 0x3a, 0xad, 0xae, 0xf0, 0x45, 0xc2, 0xc9, 0x05, 0x09,
	0x59, 0x88, 0xd8, 0x3a, 0xa8, 0x90, 0x68, 0x88, 0x90, 0xc0, 0x1d, 0xe7, 0x50, 0xd1, 0x58, 0x1b,
	0x7b, 0xcf, 0x59, 0x9d, 0xbd, 0x9b, 0xf3, 0x8e, 0x23, 0xf2, 0x0c, 0x34, 0x3c, 0x0c, 0x0f, 0x71,
	0xe5, 0x89, 0x0a, 0x51, 0x44, 0x28, 0x91, 0x28, 0x78, 0x0a, 0xe4, 0x5d, 0x47, 0x20, 0x4e, 0x74,
	0x54, 0xde, 0x99, 0x6f, 0xbe, 0x99, 0xf9, 0xc6, 0x1f, 0x76, 0x2f, 0x69, 0xbc, 0xa4, 0x5c, 0xf8,
	0x5c, 0x00, 0x2b, 0xf4, 0xf3, 0xba, 0x64, 0xc5, 0xc6, 0x5f, 0x9f, 0xfb, 0x29, 0x13, 0x4c, 0x71,
	0xe5, 0xad, 0x0a, 0x09, 0x92, 0x0c, 0xeb, 0x4a, 0xef, 0xaf, 0x4a, 0x6f, 0x7d, 0x3e, 0x3c, 0x49,
	0x65, 0x2a, 0x75, 0x99, 0x5f, 0xbd, 0x0c, 0x63, 0x78, 0x1a, 0x4b, 0x95, 0x4b, 0x15, 0x19, 0xc0,
	0x04, 0x06, 0x9a, 0xfc, 0x40, 0xb8, 0x73, 0x51, 0xb1, 0xc9, 0x7d, 0xdc, 0xe4, 0x89, 0x8d, 0xc6,
	0xc8, 0xed, 0x87, 0x4d, 0x9e, 0x90, 0x87, 0xf8, 0x28, 0x96, 0x42, 0xb0, 0x18, 0xb8, 0x14, 0x11,
	0x4f, 0xec, 0xa6, 0x86, 0xac, 0xdf, 0xc9, 0x20, 0x21, 0xa7, 0xb8, 0xa7, 0x17, 0xa8, 0xf0, 0x96,
	0xc6, 0xbb, 0x3a, 0x0e, 0x12, 0xf2, 0x00, 0x63, 0xbd, 0x56, 0x04, 0x9b, 0x15, 0xb3, 0xdb, 0x1a,
	0xec, 0xeb, 0xcc, 0xbb, 0xcd, 0x8a, 0x11, 0x1b, 0x77, 0x0b, 0x76, 0x5d, 0x32, 0x05, 0x76, 0x67,
	0x8c, 0x5c, 0x2b, 0x3c, 0x84, 0x64, 0x84, 0x07, 0x31, 0xcd, 0xb2, 0x05, 0x8d, 0xaf, 0xaa, 0xb6,
	0x3d, 0xcd, 0xc4, 0x87, 0x54, 0x90, 0x90, 0x63, 0xdc, 0x02, 0xc8, 0xec, 0xfe, 0x18, 0xb9, 0xed,
	0xb0, 0x7a, 0x92, 0x33, 0x6c, 0xd5, 0xec, 0x48, 0x31, 0x01, 0xf6, 0x60, 0x8c, 0xdc, 0x5e, 0x38,
	0xa8, 0x73, 0x73, 0x26, 0x60, 0xf2, 0xb1, 0x89, 0xfb, 0xaf, 0x28, 0xd0, 0xb7, 0x92, 0x0b, 0xb8,
	0x23, 0x96, 0xe2, 0xa3, 0x82, 0xe5, 0x12, 0x58, 0xb4, 0x64, 0x3c, 0x5d, 0x82, 0x11, 0x3b, 0x7b,
	0x71, 0xb3, 0x1d, 0x35, 0xbe, 0x6d, 0x47, 0x8f, 0x52, 0x0e, 0xcb, 0x72, 0xe1, 0xc5, 0x32, 0xaf,
	0xcf, 0x57, 0x7f, 0xa6, 0x2a, 0xb9, 0xf2, 0x2b, 0x81, 0xca, 0x0b, 0x04, 0x7c, 0xf9, 0x3c, 0xc5,
	0xf5, 0x75, 0x03, 0x01, 0xa1, 0x65, 0x5a, 0xbe, 0xd1, 0x1d, 0x49, 0x84, 0xad, 0x4c, 0xc6, 0x34,
	0x3b, 0x4c, 0x68, 0xfd, 0x87, 0x09, 0x03, 0xdd, 0xb1, 0x1e, 0xf0, 0x18, 0x77, 0xd6, 0x34, 0x2b,
	0xcd, 0xad, 0xad, 0xd9, 0xc9, 0xcf, 0xed, 0xe8, 0xb8, 0x60, 0xaa, 0xcc, 0xe0, 0x89, 0xcc, 0x39,
	0xb0, 0x7c, 0x05, 0x9b, 0xd0, 0x94, 0x4c, 0x2e, 0xb0, 0xf5, 0xda, 0x98, 0x6a, 0x0e, 0x14, 0x18,
	0x79, 0x89, 0xbb, 0xd5, 0xaf, 0xe1, 0x4c, 0xd9, 0x68, 0xdc, 0x72, 0x07, 0x4f, 0xcf, 0xbc, 0x7f,
	0xbb, 0xcc, 0xd3, 0x86, 0x99, 0xb5, 0xab, 0xd5, 0xc3, 0x03, 0x6f, 0x36, 0xbf, 0xd9, 0x39, 0xe8,
	0x76, 0xe7, 0xa0, 0xef, 0x3b, 0x07, 0x7d, 0xda, 0x3b, 0x8d, 0xdb, 0xbd, 0xd3, 0xf8, 0xba, 0x77,
	0x1a, 0xef, 0x9f, 0xff, 0xa1, 0x4d, 0xc8, 0xca, 0x39, 0x34, 0x9b, 0x66, 0x74, 0xa1, 0xfc, 0x4b,
	0x3a, 0x35, 0xa6, 0xff, 0x70, 0xc7, 0xf6, 0x5a, 0xf2, 0xe2, 0x9e, 0x76, 0xe9, 0xb3, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0xfd, 0x74, 0x00, 0x1e, 0x03, 0x00, 0x00,
}

func (m *Query) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Query) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Query) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestSent {
		i--
		if m.RequestSent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.Ttl != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Ttl))
		i--
		dAtA[i] = 0x48
	}
	if len(m.CallbackId) > 0 {
		i -= len(m.CallbackId)
		copy(dAtA[i:], m.CallbackId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CallbackId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Request) > 0 {
		i -= len(m.Request)
		copy(dAtA[i:], m.Request)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Request)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataPoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataPoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataPoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.LocalHeight.Size()
		i -= size
		if _, err := m.LocalHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RemoteHeight.Size()
		i -= size
		if _, err := m.RemoteHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for iNdEx := len(m.Queries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Queries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Query) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Request)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.CallbackId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Ttl != 0 {
		n += 1 + sovGenesis(uint64(m.Ttl))
	}
	if m.RequestSent {
		n += 2
	}
	return n
}

func (m *DataPoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.RemoteHeight.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.LocalHeight.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Query) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Query: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Query: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request[:0], dAtA[iNdEx:postIndex]...)
			if m.Request == nil {
				m.Request = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestSent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RequestSent = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DataPoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataPoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataPoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemoteHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LocalHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, Query{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
