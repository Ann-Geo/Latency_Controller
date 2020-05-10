// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cont_test/api/latencypb/latency.proto

package latencypb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ImageData struct {
	Image                []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageData) Reset()         { *m = ImageData{} }
func (m *ImageData) String() string { return proto.CompactTextString(m) }
func (*ImageData) ProtoMessage()    {}
func (*ImageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_83477b7872dcf19d, []int{0}
}

func (m *ImageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageData.Unmarshal(m, b)
}
func (m *ImageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageData.Marshal(b, m, deterministic)
}
func (m *ImageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageData.Merge(m, src)
}
func (m *ImageData) XXX_Size() int {
	return xxx_messageInfo_ImageData.Size(m)
}
func (m *ImageData) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageData.DiscardUnknown(m)
}

var xxx_messageInfo_ImageData proto.InternalMessageInfo

func (m *ImageData) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImageData) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type Response struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_83477b7872dcf19d, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*ImageData)(nil), "latency.ImageData")
	proto.RegisterType((*Response)(nil), "latency.Response")
}

func init() {
	proto.RegisterFile("cont_test/api/latencypb/latency.proto", fileDescriptor_83477b7872dcf19d)
}

var fileDescriptor_83477b7872dcf19d = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0xcf, 0x2b,
	0x89, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x49, 0x2c, 0x49, 0xcd, 0x4b,
	0xae, 0x2c, 0x48, 0x82, 0xb1, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0xa1, 0x5c, 0x25,
	0x7b, 0x2e, 0x4e, 0xcf, 0xdc, 0xc4, 0xf4, 0x54, 0x97, 0xc4, 0x92, 0x44, 0x21, 0x11, 0x2e, 0xd6,
	0x4c, 0x10, 0x47, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc2, 0x11, 0x92, 0xe1, 0xe2, 0x2c,
	0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x42,
	0x08, 0x28, 0x29, 0x71, 0x71, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x71,
	0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x83, 0x0d, 0xe0, 0x08, 0x82, 0xf2, 0x8c, 0x5a, 0x19,
	0xb9, 0xf8, 0x7c, 0x53, 0x13, 0x8b, 0x4b, 0x8b, 0x52, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0xac, 0xb9, 0xf8, 0x82, 0x53, 0xf3, 0x52, 0x1c, 0xf3, 0x52, 0xa0, 0x12, 0x42, 0x42, 0x7a,
	0x30, 0x27, 0xc2, 0x1d, 0x24, 0x25, 0x08, 0x17, 0x83, 0xd9, 0xa1, 0xc4, 0xa0, 0xc1, 0x28, 0x64,
	0xc6, 0xc5, 0x09, 0xd2, 0x0c, 0x56, 0x47, 0x82, 0x3e, 0x27, 0xee, 0x28, 0x4e, 0x78, 0x80, 0x24,
	0xb1, 0x81, 0x43, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x7d, 0xe6, 0xc9, 0x32, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MeasureServiceClient is the client API for MeasureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeasureServiceClient interface {
	SendAndMeasure(ctx context.Context, opts ...grpc.CallOption) (MeasureService_SendAndMeasureClient, error)
	SendImage(ctx context.Context, opts ...grpc.CallOption) (MeasureService_SendImageClient, error)
}

type measureServiceClient struct {
	cc *grpc.ClientConn
}

func NewMeasureServiceClient(cc *grpc.ClientConn) MeasureServiceClient {
	return &measureServiceClient{cc}
}

func (c *measureServiceClient) SendAndMeasure(ctx context.Context, opts ...grpc.CallOption) (MeasureService_SendAndMeasureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MeasureService_serviceDesc.Streams[0], "/latency.MeasureService/SendAndMeasure", opts...)
	if err != nil {
		return nil, err
	}
	x := &measureServiceSendAndMeasureClient{stream}
	return x, nil
}

type MeasureService_SendAndMeasureClient interface {
	Send(*ImageData) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type measureServiceSendAndMeasureClient struct {
	grpc.ClientStream
}

func (x *measureServiceSendAndMeasureClient) Send(m *ImageData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *measureServiceSendAndMeasureClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *measureServiceClient) SendImage(ctx context.Context, opts ...grpc.CallOption) (MeasureService_SendImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MeasureService_serviceDesc.Streams[1], "/latency.MeasureService/SendImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &measureServiceSendImageClient{stream}
	return x, nil
}

type MeasureService_SendImageClient interface {
	Send(*ImageData) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type measureServiceSendImageClient struct {
	grpc.ClientStream
}

func (x *measureServiceSendImageClient) Send(m *ImageData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *measureServiceSendImageClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeasureServiceServer is the server API for MeasureService service.
type MeasureServiceServer interface {
	SendAndMeasure(MeasureService_SendAndMeasureServer) error
	SendImage(MeasureService_SendImageServer) error
}

func RegisterMeasureServiceServer(s *grpc.Server, srv MeasureServiceServer) {
	s.RegisterService(&_MeasureService_serviceDesc, srv)
}

func _MeasureService_SendAndMeasure_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MeasureServiceServer).SendAndMeasure(&measureServiceSendAndMeasureServer{stream})
}

type MeasureService_SendAndMeasureServer interface {
	SendAndClose(*Response) error
	Recv() (*ImageData, error)
	grpc.ServerStream
}

type measureServiceSendAndMeasureServer struct {
	grpc.ServerStream
}

func (x *measureServiceSendAndMeasureServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *measureServiceSendAndMeasureServer) Recv() (*ImageData, error) {
	m := new(ImageData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MeasureService_SendImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MeasureServiceServer).SendImage(&measureServiceSendImageServer{stream})
}

type MeasureService_SendImageServer interface {
	SendAndClose(*Response) error
	Recv() (*ImageData, error)
	grpc.ServerStream
}

type measureServiceSendImageServer struct {
	grpc.ServerStream
}

func (x *measureServiceSendImageServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *measureServiceSendImageServer) Recv() (*ImageData, error) {
	m := new(ImageData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MeasureService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "latency.MeasureService",
	HandlerType: (*MeasureServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendAndMeasure",
			Handler:       _MeasureService_SendAndMeasure_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SendImage",
			Handler:       _MeasureService_SendImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "cont_test/api/latencypb/latency.proto",
}
