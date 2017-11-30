// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sssp_message.proto

/*
	Package protobuf is a generated protocol buffer package.

	It is generated from these files:
		sssp_message.proto

	It has these top-level messages:
		PutRequest
		ResponseHeader
		PutResponse
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PutRequest struct {
	NodeIndex            []byte `protobuf:"bytes,1,opt,name=nodeIndex,proto3" json:"nodeIndex,omitempty"`
	DistanceTosourceNode []byte `protobuf:"bytes,2,opt,name=distanceTosourceNode,proto3" json:"distanceTosourceNode,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptorSsspMessage, []int{0} }

type ResponseHeader struct {
	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptorSsspMessage, []int{1} }

type PutResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptorSsspMessage, []int{2} }

func init() {
	proto.RegisterType((*PutRequest)(nil), "protobuf.PutRequest")
	proto.RegisterType((*ResponseHeader)(nil), "protobuf.ResponseHeader")
	proto.RegisterType((*PutResponse)(nil), "protobuf.PutResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DistKV service

type DistKVClient interface {
	// Put puts the given dist(s,v) into the destination's buffer(message chan).
	// and generates one event in the event history.
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
}

type distKVClient struct {
	cc *grpc.ClientConn
}

func NewDistKVClient(cc *grpc.ClientConn) DistKVClient {
	return &distKVClient{cc}
}

func (c *distKVClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/protobuf.DistKV/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DistKV service

type DistKVServer interface {
	// Put puts the given dist(s,v) into the destination's buffer(message chan).
	// and generates one event in the event history.
	Put(context.Context, *PutRequest) (*PutResponse, error)
}

func RegisterDistKVServer(s *grpc.Server, srv DistKVServer) {
	s.RegisterService(&_DistKV_serviceDesc, srv)
}

func _DistKV_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistKVServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.DistKV/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistKVServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistKV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.DistKV",
	HandlerType: (*DistKVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _DistKV_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sssp_message.proto",
}

func (m *PutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NodeIndex) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSsspMessage(dAtA, i, uint64(len(m.NodeIndex)))
		i += copy(dAtA[i:], m.NodeIndex)
	}
	if len(m.DistanceTosourceNode) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSsspMessage(dAtA, i, uint64(len(m.DistanceTosourceNode)))
		i += copy(dAtA[i:], m.DistanceTosourceNode)
	}
	return i, nil
}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ok {
		dAtA[i] = 0x8
		i++
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *PutResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSsspMessage(dAtA, i, uint64(m.Header.Size()))
		n1, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintSsspMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PutRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.NodeIndex)
	if l > 0 {
		n += 1 + l + sovSsspMessage(uint64(l))
	}
	l = len(m.DistanceTosourceNode)
	if l > 0 {
		n += 1 + l + sovSsspMessage(uint64(l))
	}
	return n
}

func (m *ResponseHeader) Size() (n int) {
	var l int
	_ = l
	if m.Ok {
		n += 2
	}
	return n
}

func (m *PutResponse) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovSsspMessage(uint64(l))
	}
	return n
}

func sovSsspMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSsspMessage(x uint64) (n int) {
	return sovSsspMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSsspMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeIndex", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSsspMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSsspMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeIndex = append(m.NodeIndex[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeIndex == nil {
				m.NodeIndex = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistanceTosourceNode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSsspMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSsspMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistanceTosourceNode = append(m.DistanceTosourceNode[:0], dAtA[iNdEx:postIndex]...)
			if m.DistanceTosourceNode == nil {
				m.DistanceTosourceNode = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSsspMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSsspMessage
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
func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSsspMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSsspMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSsspMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSsspMessage
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
func (m *PutResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSsspMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSsspMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSsspMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSsspMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSsspMessage
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
func skipSsspMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSsspMessage
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
					return 0, ErrIntOverflowSsspMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSsspMessage
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSsspMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSsspMessage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSsspMessage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSsspMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSsspMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sssp_message.proto", fileDescriptorSsspMessage) }

var fileDescriptorSsspMessage = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xcd, 0x4a, 0xc3, 0x50,
	0x10, 0x85, 0x93, 0x08, 0xa1, 0x4e, 0xa5, 0x8b, 0x4b, 0x84, 0x50, 0x24, 0x94, 0xac, 0xdc, 0x98,
	0x4a, 0x74, 0xad, 0x20, 0x2e, 0x14, 0x41, 0x4a, 0x10, 0x97, 0x4a, 0x7e, 0xc6, 0x34, 0x94, 0x66,
	0x62, 0xe6, 0x5e, 0xf0, 0x51, 0x7c, 0xa4, 0x2e, 0x7d, 0x04, 0x8d, 0x2f, 0x22, 0x4c, 0x2c, 0x41,
	0xe9, 0xea, 0xde, 0x99, 0xf3, 0x9d, 0xc3, 0x1c, 0x50, 0xcc, 0xdc, 0x3c, 0xaf, 0x91, 0x39, 0x2d,
	0x31, 0x6a, 0x5a, 0xd2, 0xa4, 0x46, 0xf2, 0x64, 0xe6, 0x65, 0x7a, 0x52, 0x56, 0x7a, 0x69, 0xb2,
	0x28, 0xa7, 0xf5, 0xbc, 0xa4, 0x92, 0xe6, 0x5b, 0x45, 0x26, 0x19, 0xe4, 0xd7, 0x1b, 0xc3, 0x27,
	0x80, 0x85, 0xd1, 0x09, 0xbe, 0x1a, 0x64, 0xad, 0x8e, 0x60, 0xbf, 0xa6, 0x02, 0x6f, 0xeb, 0x02,
	0xdf, 0x7c, 0x7b, 0x66, 0x1f, 0x1f, 0x24, 0xc3, 0x42, 0xc5, 0xe0, 0x15, 0x15, 0xeb, 0xb4, 0xce,
	0xf1, 0x81, 0x98, 0x4c, 0x9b, 0xe3, 0x3d, 0x15, 0xe8, 0x3b, 0x02, 0xee, 0xd4, 0xc2, 0x19, 0x4c,
	0x12, 0xe4, 0x86, 0x6a, 0xc6, 0x1b, 0x4c, 0x0b, 0x6c, 0xd5, 0x04, 0x1c, 0x5a, 0x49, 0xf8, 0x28,
	0x71, 0x68, 0x15, 0x5e, 0xc2, 0x58, 0x2e, 0xe8, 0x21, 0x75, 0x0a, 0xee, 0x52, 0x40, 0x41, 0xc6,
	0xb1, 0x1f, 0x6d, 0x0b, 0x44, 0x7f, 0x83, 0x92, 0x5f, 0x2e, 0xbe, 0x00, 0xf7, 0xba, 0x62, 0x7d,
	0xf7, 0xa8, 0xce, 0x61, 0x6f, 0x61, 0xb4, 0xf2, 0x06, 0xcb, 0xd0, 0x6d, 0x7a, 0xf8, 0x6f, 0xdb,
	0x67, 0x85, 0xd6, 0x95, 0xb7, 0xf9, 0x0a, 0xac, 0x4d, 0x17, 0xd8, 0x1f, 0x5d, 0x60, 0x7f, 0x76,
	0x81, 0xfd, 0xfe, 0x1d, 0x58, 0x99, 0x2b, 0xf4, 0xd9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc,
	0x78, 0x72, 0xb2, 0x6e, 0x01, 0x00, 0x00,
}
