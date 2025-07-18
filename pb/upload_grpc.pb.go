// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: upload.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StreamUpload_Upload_FullMethodName = "/StreamUpload/Upload"
)

// StreamUploadClient is the client API for StreamUpload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamUploadClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error)
}

type streamUploadClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamUploadClient(cc grpc.ClientConnInterface) StreamUploadClient {
	return &streamUploadClient{cc}
}

func (c *streamUploadClient) Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamUpload_ServiceDesc.Streams[0], StreamUpload_Upload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadRequest, UploadResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamUpload_UploadClient = grpc.ClientStreamingClient[UploadRequest, UploadResponse]

// StreamUploadServer is the server API for StreamUpload service.
// All implementations must embed UnimplementedStreamUploadServer
// for forward compatibility.
type StreamUploadServer interface {
	Upload(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error
	mustEmbedUnimplementedStreamUploadServer()
}

// UnimplementedStreamUploadServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStreamUploadServer struct{}

func (UnimplementedStreamUploadServer) Upload(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedStreamUploadServer) mustEmbedUnimplementedStreamUploadServer() {}
func (UnimplementedStreamUploadServer) testEmbeddedByValue()                      {}

// UnsafeStreamUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamUploadServer will
// result in compilation errors.
type UnsafeStreamUploadServer interface {
	mustEmbedUnimplementedStreamUploadServer()
}

func RegisterStreamUploadServer(s grpc.ServiceRegistrar, srv StreamUploadServer) {
	// If the following call pancis, it indicates UnimplementedStreamUploadServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StreamUpload_ServiceDesc, srv)
}

func _StreamUpload_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamUploadServer).Upload(&grpc.GenericServerStream[UploadRequest, UploadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamUpload_UploadServer = grpc.ClientStreamingServer[UploadRequest, UploadResponse]

// StreamUpload_ServiceDesc is the grpc.ServiceDesc for StreamUpload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamUpload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StreamUpload",
	HandlerType: (*StreamUploadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _StreamUpload_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "upload.proto",
}
