// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/staking/v1beta1/hooks.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type AfterValidatorCreatedRequest struct {
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *AfterValidatorCreatedRequest) Reset()         { *m = AfterValidatorCreatedRequest{} }
func (m *AfterValidatorCreatedRequest) String() string { return proto.CompactTextString(m) }
func (*AfterValidatorCreatedRequest) ProtoMessage()    {}
func (*AfterValidatorCreatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5837010b9e1c0afc, []int{0}
}
func (m *AfterValidatorCreatedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AfterValidatorCreatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AfterValidatorCreatedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AfterValidatorCreatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AfterValidatorCreatedRequest.Merge(m, src)
}
func (m *AfterValidatorCreatedRequest) XXX_Size() int {
	return m.Size()
}
func (m *AfterValidatorCreatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AfterValidatorCreatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AfterValidatorCreatedRequest proto.InternalMessageInfo

func (m *AfterValidatorCreatedRequest) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

type AfterValidatorCreatedResponse struct {
}

func (m *AfterValidatorCreatedResponse) Reset()         { *m = AfterValidatorCreatedResponse{} }
func (m *AfterValidatorCreatedResponse) String() string { return proto.CompactTextString(m) }
func (*AfterValidatorCreatedResponse) ProtoMessage()    {}
func (*AfterValidatorCreatedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5837010b9e1c0afc, []int{1}
}
func (m *AfterValidatorCreatedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AfterValidatorCreatedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AfterValidatorCreatedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AfterValidatorCreatedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AfterValidatorCreatedResponse.Merge(m, src)
}
func (m *AfterValidatorCreatedResponse) XXX_Size() int {
	return m.Size()
}
func (m *AfterValidatorCreatedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AfterValidatorCreatedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AfterValidatorCreatedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AfterValidatorCreatedRequest)(nil), "cosmos.staking.v1beta1.AfterValidatorCreatedRequest")
	proto.RegisterType((*AfterValidatorCreatedResponse)(nil), "cosmos.staking.v1beta1.AfterValidatorCreatedResponse")
}

func init() {
	proto.RegisterFile("cosmos/staking/v1beta1/hooks.proto", fileDescriptor_5837010b9e1c0afc)
}

var fileDescriptor_5837010b9e1c0afc = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0xcf, 0xc8, 0xcf, 0xcf, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x83, 0xa8, 0xd1, 0x83, 0xaa, 0xd1, 0x83, 0xaa, 0x51, 0xf2, 0xe6, 0x92, 0x71, 0x4c, 0x2b, 0x49,
	0x2d, 0x0a, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x72, 0x2e, 0x4a, 0x4d, 0x2c, 0x49,
	0x4d, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xd2, 0xe6, 0x12, 0x2c, 0x83, 0x49, 0xc5,
	0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x09, 0xc0,
	0x25, 0x1c, 0x21, 0xe2, 0x4a, 0xf2, 0x5c, 0xb2, 0x38, 0x0c, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e,
	0x35, 0xea, 0x63, 0xe4, 0x62, 0xf5, 0x00, 0xb9, 0x4a, 0xa8, 0x85, 0x91, 0x4b, 0x14, 0xab, 0x5a,
	0x21, 0x13, 0x3d, 0xec, 0x4e, 0xd5, 0xc3, 0xe7, 0x4e, 0x29, 0x53, 0x12, 0x75, 0x41, 0x1c, 0xe4,
	0xe4, 0x76, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x3a, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xd0, 0xf0, 0x85, 0x50, 0xba, 0xc5, 0x29, 0xd9,
	0xfa, 0x15, 0xf0, 0xc0, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x87, 0xb2, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xf4, 0x30, 0x98, 0xe2, 0x8b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HooksClient is the client API for Hooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HooksClient interface {
	AfterValidatorCreated(ctx context.Context, in *AfterValidatorCreatedRequest, opts ...grpc.CallOption) (*AfterValidatorCreatedResponse, error)
}

type hooksClient struct {
	cc grpc1.ClientConn
}

func NewHooksClient(cc grpc1.ClientConn) HooksClient {
	return &hooksClient{cc}
}

func (c *hooksClient) AfterValidatorCreated(ctx context.Context, in *AfterValidatorCreatedRequest, opts ...grpc.CallOption) (*AfterValidatorCreatedResponse, error) {
	out := new(AfterValidatorCreatedResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Hooks/AfterValidatorCreated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HooksServer is the server API for Hooks service.
type HooksServer interface {
	AfterValidatorCreated(context.Context, *AfterValidatorCreatedRequest) (*AfterValidatorCreatedResponse, error)
}

// UnimplementedHooksServer can be embedded to have forward compatible implementations.
type UnimplementedHooksServer struct {
}

func (*UnimplementedHooksServer) AfterValidatorCreated(ctx context.Context, req *AfterValidatorCreatedRequest) (*AfterValidatorCreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AfterValidatorCreated not implemented")
}

func RegisterHooksServer(s grpc1.Server, srv HooksServer) {
	s.RegisterService(&_Hooks_serviceDesc, srv)
}

func _Hooks_AfterValidatorCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterValidatorCreatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksServer).AfterValidatorCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Hooks/AfterValidatorCreated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksServer).AfterValidatorCreated(ctx, req.(*AfterValidatorCreatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hooks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.staking.v1beta1.Hooks",
	HandlerType: (*HooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AfterValidatorCreated",
			Handler:    _Hooks_AfterValidatorCreated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/staking/v1beta1/hooks.proto",
}

func (m *AfterValidatorCreatedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AfterValidatorCreatedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AfterValidatorCreatedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintHooks(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AfterValidatorCreatedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AfterValidatorCreatedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AfterValidatorCreatedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintHooks(dAtA []byte, offset int, v uint64) int {
	offset -= sovHooks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AfterValidatorCreatedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovHooks(uint64(l))
	}
	return n
}

func (m *AfterValidatorCreatedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovHooks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHooks(x uint64) (n int) {
	return sovHooks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AfterValidatorCreatedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHooks
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
			return fmt.Errorf("proto: AfterValidatorCreatedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AfterValidatorCreatedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHooks
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
				return ErrInvalidLengthHooks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHooks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHooks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHooks
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHooks
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
func (m *AfterValidatorCreatedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHooks
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
			return fmt.Errorf("proto: AfterValidatorCreatedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AfterValidatorCreatedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHooks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHooks
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHooks
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
func skipHooks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHooks
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
					return 0, ErrIntOverflowHooks
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
					return 0, ErrIntOverflowHooks
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
				return 0, ErrInvalidLengthHooks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHooks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHooks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHooks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHooks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHooks = fmt.Errorf("proto: unexpected end of group")
)