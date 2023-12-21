// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rollup/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryL1BlockInfoRequest is the request type for the Query/L1BlockInfo RPC
type QueryL1BlockInfoRequest struct {
	// L2 block height; use 0 for latest block height
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *QueryL1BlockInfoRequest) Reset()         { *m = QueryL1BlockInfoRequest{} }
func (m *QueryL1BlockInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryL1BlockInfoRequest) ProtoMessage()    {}
func (*QueryL1BlockInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_262fcfe2d8831f2e, []int{0}
}
func (m *QueryL1BlockInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryL1BlockInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryL1BlockInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryL1BlockInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryL1BlockInfoRequest.Merge(m, src)
}
func (m *QueryL1BlockInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryL1BlockInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryL1BlockInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryL1BlockInfoRequest proto.InternalMessageInfo

// QueryL1BlockInfoResponse is the stored L1 block info
type QueryL1BlockInfoResponse struct {
	// Block number
	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// Block timestamp
	Time uint64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	// Base fee for the block
	BaseFee []byte `protobuf:"bytes,3,opt,name=baseFee,proto3" json:"baseFee,omitempty"`
	// Hash of the blocK; bytes32
	BlockHash []byte `protobuf:"bytes,4,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	// Number of L2 blocks since the start of the epoch
	// Not strictly a piece of L1 information. Represents the number of L2 blocks since the start of the epoch,
	// i.e. when the actual L1 info was first introduced.
	SequenceNumber uint64 `protobuf:"varint,5,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`
	// Fields 6,7,8 are SystemConfig
	// Address of the batcher; bytes20
	BatcherAddr []byte `protobuf:"bytes,6,opt,name=batcherAddr,proto3" json:"batcherAddr,omitempty"`
	// Overhead fee for L1; bytes32
	L1FeeOverhead []byte `protobuf:"bytes,7,opt,name=l1FeeOverhead,proto3" json:"l1FeeOverhead,omitempty"`
	// Scalar fee for L1; bytes32
	L1FeeScalar []byte `protobuf:"bytes,8,opt,name=l1FeeScalar,proto3" json:"l1FeeScalar,omitempty"`
}

func (m *QueryL1BlockInfoResponse) Reset()         { *m = QueryL1BlockInfoResponse{} }
func (m *QueryL1BlockInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryL1BlockInfoResponse) ProtoMessage()    {}
func (*QueryL1BlockInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_262fcfe2d8831f2e, []int{1}
}
func (m *QueryL1BlockInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryL1BlockInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryL1BlockInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryL1BlockInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryL1BlockInfoResponse.Merge(m, src)
}
func (m *QueryL1BlockInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryL1BlockInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryL1BlockInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryL1BlockInfoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryL1BlockInfoRequest)(nil), "rollup.QueryL1BlockInfoRequest")
	proto.RegisterType((*QueryL1BlockInfoResponse)(nil), "rollup.QueryL1BlockInfoResponse")
}

func init() { proto.RegisterFile("rollup/query.proto", fileDescriptor_262fcfe2d8831f2e) }

var fileDescriptor_262fcfe2d8831f2e = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0x6b, 0x9a, 0xea, 0x54, 0x3d, 0x0c, 0xa2, 0xc3, 0xb2, 0xa4, 0x61, 0x29, 0xd2,
	0x53, 0x86, 0x55, 0x44, 0xf0, 0x66, 0x0f, 0x45, 0x41, 0x2c, 0xae, 0x37, 0x6f, 0x93, 0xe4, 0x35,
	0x09, 0x4e, 0xe6, 0xa5, 0x33, 0x93, 0xe2, 0x22, 0x22, 0x78, 0xf2, 0x28, 0x08, 0x9e, 0xfd, 0x38,
	0x1e, 0x0b, 0x5e, 0x3c, 0xca, 0xae, 0x1f, 0x44, 0x32, 0xd3, 0xe0, 0xaa, 0xec, 0xed, 0xbd, 0xdf,
	0x7b, 0xef, 0x3f, 0xc3, 0xff, 0x3d, 0x42, 0x35, 0x4a, 0xd9, 0xb5, 0xfc, 0xac, 0x03, 0xbd, 0x4c,
	0x5b, 0x8d, 0x16, 0x69, 0xe4, 0xd9, 0xe4, 0x56, 0x89, 0x25, 0x3a, 0xc4, 0xfb, 0xc8, 0x57, 0x27,
	0xd3, 0x12, 0xb1, 0x94, 0xc0, 0x45, 0x5b, 0x73, 0xa1, 0x14, 0x5a, 0x61, 0x6b, 0x54, 0xc6, 0x57,
	0x67, 0x0f, 0xc9, 0x9d, 0x17, 0xbd, 0xd4, 0xb3, 0xf9, 0x91, 0xc4, 0xfc, 0xf5, 0x53, 0x75, 0x8a,
	0x0b, 0x38, 0xeb, 0xc0, 0x58, 0x7a, 0x9b, 0x44, 0x15, 0xd4, 0x65, 0x65, 0x59, 0x90, 0x04, 0x87,
	0xe1, 0xe2, 0x32, 0x7b, 0x14, 0x7e, 0xfc, 0xba, 0x3f, 0x9a, 0x7d, 0x19, 0x13, 0xf6, 0xff, 0xa4,
	0x69, 0x51, 0x19, 0xe8, 0x47, 0x55, 0xd7, 0x64, 0xa0, 0x87, 0x51, 0x9f, 0x51, 0x4a, 0x42, 0x5b,
	0x37, 0xc0, 0xc6, 0x8e, 0xba, 0x98, 0x32, 0xb2, 0x9b, 0x09, 0x03, 0xc7, 0x00, 0xec, 0x4a, 0x12,
	0x1c, 0x5e, 0x5f, 0x0c, 0x29, 0x9d, 0x92, 0x6b, 0x59, 0x2f, 0xfd, 0x44, 0x98, 0x8a, 0x85, 0xae,
	0xf6, 0x07, 0xd0, 0xbb, 0xe4, 0xa6, 0xe9, 0x7f, 0xaa, 0x72, 0x78, 0xee, 0xdf, 0xda, 0x71, 0xaa,
	0xff, 0x50, 0x9a, 0x90, 0xbd, 0x4c, 0xd8, 0xbc, 0x02, 0xfd, 0xb8, 0x28, 0x34, 0x8b, 0x9c, 0xce,
	0x26, 0xa2, 0x07, 0xe4, 0x86, 0x9c, 0x1f, 0x03, 0x9c, 0x9c, 0x83, 0xae, 0x40, 0x14, 0x6c, 0xd7,
	0xf5, 0xfc, 0x0d, 0x7b, 0x1d, 0x07, 0x5e, 0xe6, 0x42, 0x0a, 0xcd, 0xae, 0x7a, 0x9d, 0x0d, 0xe4,
	0x8d, 0xb9, 0xf7, 0x9e, 0xec, 0x38, 0x5f, 0xe8, 0x39, 0xd9, 0xdb, 0xf0, 0x86, 0xee, 0xa7, 0x7e,
	0x4d, 0xe9, 0x16, 0xbf, 0x27, 0xc9, 0xf6, 0x06, 0x6f, 0xeb, 0xec, 0xe0, 0xc3, 0xf7, 0x5f, 0x9f,
	0xc7, 0x31, 0x9d, 0xf2, 0xcb, 0x2b, 0x90, 0x73, 0xe7, 0x47, 0xad, 0x4e, 0x91, 0xbf, 0xf5, 0xeb,
	0x79, 0x77, 0x74, 0xf2, 0x6d, 0x15, 0x07, 0x17, 0xab, 0x38, 0xf8, 0xb9, 0x8a, 0x83, 0x4f, 0xeb,
	0x78, 0x74, 0xb1, 0x8e, 0x47, 0x3f, 0xd6, 0xf1, 0xe8, 0xd5, 0x83, 0xb2, 0xb6, 0x55, 0x97, 0xa5,
	0x39, 0x36, 0xbc, 0x45, 0xb9, 0x6c, 0x40, 0x17, 0x02, 0x87, 0x50, 0x18, 0xe0, 0x79, 0x25, 0x6a,
	0xc5, 0xdf, 0x0c, 0xfa, 0x76, 0xd9, 0x82, 0xc9, 0x22, 0x77, 0x2a, 0xf7, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xb2, 0x40, 0x0a, 0xf6, 0x7c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	L1BlockInfo(ctx context.Context, in *QueryL1BlockInfoRequest, opts ...grpc.CallOption) (*QueryL1BlockInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) L1BlockInfo(ctx context.Context, in *QueryL1BlockInfoRequest, opts ...grpc.CallOption) (*QueryL1BlockInfoResponse, error) {
	out := new(QueryL1BlockInfoResponse)
	err := c.cc.Invoke(ctx, "/rollup.Query/L1BlockInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	L1BlockInfo(context.Context, *QueryL1BlockInfoRequest) (*QueryL1BlockInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) L1BlockInfo(ctx context.Context, req *QueryL1BlockInfoRequest) (*QueryL1BlockInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method L1BlockInfo not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_L1BlockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryL1BlockInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).L1BlockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rollup.Query/L1BlockInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).L1BlockInfo(ctx, req.(*QueryL1BlockInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rollup.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "L1BlockInfo",
			Handler:    _Query_L1BlockInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rollup/query.proto",
}

func (m *QueryL1BlockInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryL1BlockInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryL1BlockInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryL1BlockInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryL1BlockInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryL1BlockInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.L1FeeScalar) > 0 {
		i -= len(m.L1FeeScalar)
		copy(dAtA[i:], m.L1FeeScalar)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.L1FeeScalar)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.L1FeeOverhead) > 0 {
		i -= len(m.L1FeeOverhead)
		copy(dAtA[i:], m.L1FeeOverhead)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.L1FeeOverhead)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BatcherAddr) > 0 {
		i -= len(m.BatcherAddr)
		copy(dAtA[i:], m.BatcherAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BatcherAddr)))
		i--
		dAtA[i] = 0x32
	}
	if m.SequenceNumber != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.SequenceNumber))
		i--
		dAtA[i] = 0x28
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BaseFee) > 0 {
		i -= len(m.BaseFee)
		copy(dAtA[i:], m.BaseFee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BaseFee)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Time != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x10
	}
	if m.Number != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryL1BlockInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	return n
}

func (m *QueryL1BlockInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovQuery(uint64(m.Number))
	}
	if m.Time != 0 {
		n += 1 + sovQuery(uint64(m.Time))
	}
	l = len(m.BaseFee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.SequenceNumber != 0 {
		n += 1 + sovQuery(uint64(m.SequenceNumber))
	}
	l = len(m.BatcherAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.L1FeeOverhead)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.L1FeeScalar)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryL1BlockInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryL1BlockInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryL1BlockInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryL1BlockInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryL1BlockInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryL1BlockInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseFee = append(m.BaseFee[:0], dAtA[iNdEx:postIndex]...)
			if m.BaseFee == nil {
				m.BaseFee = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatcherAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatcherAddr = append(m.BatcherAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.BatcherAddr == nil {
				m.BatcherAddr = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field L1FeeOverhead", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.L1FeeOverhead = append(m.L1FeeOverhead[:0], dAtA[iNdEx:postIndex]...)
			if m.L1FeeOverhead == nil {
				m.L1FeeOverhead = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field L1FeeScalar", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.L1FeeScalar = append(m.L1FeeScalar[:0], dAtA[iNdEx:postIndex]...)
			if m.L1FeeScalar == nil {
				m.L1FeeScalar = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
