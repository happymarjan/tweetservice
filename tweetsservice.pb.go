// Code generated by protoc-gen-go.
// source: tweetsservice.proto
// DO NOT EDIT!

/*
Package tweetservice is a generated protocol buffer package.

It is generated from these files:
	tweetsservice.proto

It has these top-level messages:
	EmptyParam
	Tweet
	Tweetlist
	Status
	Statuslist
*/
package tweetservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmptyParam struct {
}

func (m *EmptyParam) Reset()                    { *m = EmptyParam{} }
func (m *EmptyParam) String() string            { return proto.CompactTextString(m) }
func (*EmptyParam) ProtoMessage()               {}
func (*EmptyParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Tweet struct {
	Author string `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	Date   int32  `protobuf:"varint,2,opt,name=date" json:"date,omitempty"`
	Text   string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
}

func (m *Tweet) Reset()                    { *m = Tweet{} }
func (m *Tweet) String() string            { return proto.CompactTextString(m) }
func (*Tweet) ProtoMessage()               {}
func (*Tweet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Tweet) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Tweet) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Tweet) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Tweetlist struct {
	Tweetsingle []*Tweet `protobuf:"bytes,1,rep,name=tweetsingle" json:"tweetsingle,omitempty"`
}

func (m *Tweetlist) Reset()                    { *m = Tweetlist{} }
func (m *Tweetlist) String() string            { return proto.CompactTextString(m) }
func (*Tweetlist) ProtoMessage()               {}
func (*Tweetlist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tweetlist) GetTweetsingle() []*Tweet {
	if m != nil {
		return m.Tweetsingle
	}
	return nil
}

type Status struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Status) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Statuslist struct {
	StatusArray []*Status `protobuf:"bytes,1,rep,name=StatusArray" json:"StatusArray,omitempty"`
}

func (m *Statuslist) Reset()                    { *m = Statuslist{} }
func (m *Statuslist) String() string            { return proto.CompactTextString(m) }
func (*Statuslist) ProtoMessage()               {}
func (*Statuslist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Statuslist) GetStatusArray() []*Status {
	if m != nil {
		return m.StatusArray
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyParam)(nil), "tweetservice.EmptyParam")
	proto.RegisterType((*Tweet)(nil), "tweetservice.Tweet")
	proto.RegisterType((*Tweetlist)(nil), "tweetservice.Tweetlist")
	proto.RegisterType((*Status)(nil), "tweetservice.Status")
	proto.RegisterType((*Statuslist)(nil), "tweetservice.Statuslist")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TweetsService service

type TweetsServiceClient interface {
	GetTweet(ctx context.Context, in *EmptyParam, opts ...grpc.CallOption) (TweetsService_GetTweetClient, error)
}

type tweetsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTweetsServiceClient(cc *grpc.ClientConn) TweetsServiceClient {
	return &tweetsServiceClient{cc}
}

func (c *tweetsServiceClient) GetTweet(ctx context.Context, in *EmptyParam, opts ...grpc.CallOption) (TweetsService_GetTweetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TweetsService_serviceDesc.Streams[0], c.cc, "/tweetservice.TweetsService/GetTweet", opts...)
	if err != nil {
		return nil, err
	}
	x := &tweetsServiceGetTweetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TweetsService_GetTweetClient interface {
	Recv() (*Statuslist, error)
	grpc.ClientStream
}

type tweetsServiceGetTweetClient struct {
	grpc.ClientStream
}

func (x *tweetsServiceGetTweetClient) Recv() (*Statuslist, error) {
	m := new(Statuslist)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TweetsService service

type TweetsServiceServer interface {
	GetTweet(*EmptyParam, TweetsService_GetTweetServer) error
}

func RegisterTweetsServiceServer(s *grpc.Server, srv TweetsServiceServer) {
	s.RegisterService(&_TweetsService_serviceDesc, srv)
}

func _TweetsService_GetTweet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TweetsServiceServer).GetTweet(m, &tweetsServiceGetTweetServer{stream})
}

type TweetsService_GetTweetServer interface {
	Send(*Statuslist) error
	grpc.ServerStream
}

type tweetsServiceGetTweetServer struct {
	grpc.ServerStream
}

func (x *tweetsServiceGetTweetServer) Send(m *Statuslist) error {
	return x.ServerStream.SendMsg(m)
}

var _TweetsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tweetservice.TweetsService",
	HandlerType: (*TweetsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTweet",
			Handler:       _TweetsService_GetTweet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tweetsservice.proto",
}

func init() { proto.RegisterFile("tweetsservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0x5d, 0x6b, 0x83, 0x9d, 0xd4, 0xcb, 0x46, 0x74, 0x11, 0x0f, 0x61, 0x4f, 0x39, 0x05,
	0xa9, 0xe8, 0xdd, 0xa0, 0xf4, 0x1a, 0x12, 0x5f, 0x60, 0xd5, 0x41, 0x03, 0xad, 0x29, 0xbb, 0x53,
	0xb5, 0x6f, 0x2f, 0x99, 0x29, 0xba, 0x81, 0xde, 0xe6, 0xe7, 0xdb, 0xf9, 0x86, 0x59, 0xc8, 0xe8,
	0x1b, 0x91, 0x42, 0x40, 0xff, 0xd5, 0xbd, 0x62, 0xb9, 0xf1, 0x3d, 0xf5, 0x7a, 0x2e, 0x45, 0xa9,
	0xd9, 0x39, 0xc0, 0xd3, 0x7a, 0x43, 0xbb, 0xda, 0x79, 0xb7, 0xb6, 0x4b, 0x98, 0x3e, 0x0f, 0x5d,
	0x7d, 0x01, 0x89, 0xdb, 0xd2, 0x47, 0xef, 0x8d, 0xca, 0x55, 0x31, 0x6b, 0xf6, 0x99, 0xd6, 0x70,
	0xf2, 0xe6, 0x08, 0xcd, 0x71, 0xae, 0x8a, 0x69, 0xc3, 0xf1, 0x50, 0x23, 0xfc, 0x21, 0x33, 0x61,
	0x92, 0x63, 0x5b, 0xc1, 0x8c, 0x07, 0xad, 0xba, 0x40, 0xfa, 0x0e, 0x52, 0x71, 0x76, 0x9f, 0xef,
	0x2b, 0x34, 0x2a, 0x9f, 0x14, 0xe9, 0x22, 0x2b, 0xe3, 0x3d, 0x4a, 0xa6, 0x9b, 0x98, 0xb3, 0xd7,
	0x90, 0xb4, 0xe4, 0x68, 0x1b, 0xfe, 0x0c, 0x2a, 0x32, 0x3c, 0x02, 0x48, 0x97, 0x15, 0xf7, 0x90,
	0x4a, 0xf6, 0xe0, 0xbd, 0xdb, 0xed, 0x15, 0xe7, 0x63, 0x85, 0x00, 0x4d, 0x0c, 0x2e, 0x5a, 0x38,
	0x63, 0x73, 0x68, 0x05, 0xd2, 0x15, 0x9c, 0x2e, 0x91, 0xe4, 0x08, 0x66, 0xfc, 0xfe, 0xff, 0x4e,
	0x57, 0xe6, 0xd0, 0xe4, 0x61, 0x11, 0x7b, 0x74, 0xa3, 0xaa, 0xcb, 0x2a, 0x1b, 0x0d, 0xcd, 0xeb,
	0xe1, 0xf0, 0xb5, 0x7a, 0x49, 0xf8, 0x07, 0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x77, 0xa1,
	0x69, 0xce, 0x98, 0x01, 0x00, 0x00,
}