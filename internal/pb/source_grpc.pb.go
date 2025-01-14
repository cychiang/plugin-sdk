// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: internal/pb/source.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SourceClient is the client API for Source service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourceClient interface {
	// Get the current protocol version of the plugin. This helps
	// get the right message about upgrade/downgrade of cli and/or plugin.
	// Also, on the cli side it can try to upgrade/downgrade the protocol if cli supports it.
	GetProtocolVersion(ctx context.Context, in *GetProtocolVersion_Request, opts ...grpc.CallOption) (*GetProtocolVersion_Response, error)
	// Get the name of the plugin
	GetName(ctx context.Context, in *GetName_Request, opts ...grpc.CallOption) (*GetName_Response, error)
	// Get the current version of the plugin
	GetVersion(ctx context.Context, in *GetVersion_Request, opts ...grpc.CallOption) (*GetVersion_Response, error)
	// Get all tables the source plugin supports
	GetTables(ctx context.Context, in *GetTables_Request, opts ...grpc.CallOption) (*GetTables_Response, error)
	// GetSyncSummary returns the latest sync summary of the source plugin. we don't want to send the summary on
	// every sync request.
	GetSyncSummary(ctx context.Context, in *GetSyncSummary_Request, opts ...grpc.CallOption) (*GetSyncSummary_Response, error)
	// Fetch resources
	Sync(ctx context.Context, in *Sync_Request, opts ...grpc.CallOption) (Source_SyncClient, error)
	// Sync2 is a new sync API that supports CQ Types. It is not backward compatible with Sync.
	Sync2(ctx context.Context, in *Sync2_Request, opts ...grpc.CallOption) (Source_Sync2Client, error)
	// Get metrics for the source plugin
	GetMetrics(ctx context.Context, in *GetSourceMetrics_Request, opts ...grpc.CallOption) (*GetSourceMetrics_Response, error)
}

type sourceClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceClient(cc grpc.ClientConnInterface) SourceClient {
	return &sourceClient{cc}
}

func (c *sourceClient) GetProtocolVersion(ctx context.Context, in *GetProtocolVersion_Request, opts ...grpc.CallOption) (*GetProtocolVersion_Response, error) {
	out := new(GetProtocolVersion_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/GetProtocolVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) GetName(ctx context.Context, in *GetName_Request, opts ...grpc.CallOption) (*GetName_Response, error) {
	out := new(GetName_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) GetVersion(ctx context.Context, in *GetVersion_Request, opts ...grpc.CallOption) (*GetVersion_Response, error) {
	out := new(GetVersion_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) GetTables(ctx context.Context, in *GetTables_Request, opts ...grpc.CallOption) (*GetTables_Response, error) {
	out := new(GetTables_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/GetTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) GetSyncSummary(ctx context.Context, in *GetSyncSummary_Request, opts ...grpc.CallOption) (*GetSyncSummary_Response, error) {
	out := new(GetSyncSummary_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/GetSyncSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Sync(ctx context.Context, in *Sync_Request, opts ...grpc.CallOption) (Source_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &Source_ServiceDesc.Streams[0], "/proto.Source/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Source_SyncClient interface {
	Recv() (*Sync_Response, error)
	grpc.ClientStream
}

type sourceSyncClient struct {
	grpc.ClientStream
}

func (x *sourceSyncClient) Recv() (*Sync_Response, error) {
	m := new(Sync_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sourceClient) Sync2(ctx context.Context, in *Sync2_Request, opts ...grpc.CallOption) (Source_Sync2Client, error) {
	stream, err := c.cc.NewStream(ctx, &Source_ServiceDesc.Streams[1], "/proto.Source/Sync2", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceSync2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Source_Sync2Client interface {
	Recv() (*Sync2_Response, error)
	grpc.ClientStream
}

type sourceSync2Client struct {
	grpc.ClientStream
}

func (x *sourceSync2Client) Recv() (*Sync2_Response, error) {
	m := new(Sync2_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sourceClient) GetMetrics(ctx context.Context, in *GetSourceMetrics_Request, opts ...grpc.CallOption) (*GetSourceMetrics_Response, error) {
	out := new(GetSourceMetrics_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourceServer is the server API for Source service.
// All implementations must embed UnimplementedSourceServer
// for forward compatibility
type SourceServer interface {
	// Get the current protocol version of the plugin. This helps
	// get the right message about upgrade/downgrade of cli and/or plugin.
	// Also, on the cli side it can try to upgrade/downgrade the protocol if cli supports it.
	GetProtocolVersion(context.Context, *GetProtocolVersion_Request) (*GetProtocolVersion_Response, error)
	// Get the name of the plugin
	GetName(context.Context, *GetName_Request) (*GetName_Response, error)
	// Get the current version of the plugin
	GetVersion(context.Context, *GetVersion_Request) (*GetVersion_Response, error)
	// Get all tables the source plugin supports
	GetTables(context.Context, *GetTables_Request) (*GetTables_Response, error)
	// GetSyncSummary returns the latest sync summary of the source plugin. we don't want to send the summary on
	// every sync request.
	GetSyncSummary(context.Context, *GetSyncSummary_Request) (*GetSyncSummary_Response, error)
	// Fetch resources
	Sync(*Sync_Request, Source_SyncServer) error
	// Sync2 is a new sync API that supports CQ Types. It is not backward compatible with Sync.
	Sync2(*Sync2_Request, Source_Sync2Server) error
	// Get metrics for the source plugin
	GetMetrics(context.Context, *GetSourceMetrics_Request) (*GetSourceMetrics_Response, error)
	mustEmbedUnimplementedSourceServer()
}

// UnimplementedSourceServer must be embedded to have forward compatible implementations.
type UnimplementedSourceServer struct {
}

func (UnimplementedSourceServer) GetProtocolVersion(context.Context, *GetProtocolVersion_Request) (*GetProtocolVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProtocolVersion not implemented")
}
func (UnimplementedSourceServer) GetName(context.Context, *GetName_Request) (*GetName_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}
func (UnimplementedSourceServer) GetVersion(context.Context, *GetVersion_Request) (*GetVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedSourceServer) GetTables(context.Context, *GetTables_Request) (*GetTables_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTables not implemented")
}
func (UnimplementedSourceServer) GetSyncSummary(context.Context, *GetSyncSummary_Request) (*GetSyncSummary_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSyncSummary not implemented")
}
func (UnimplementedSourceServer) Sync(*Sync_Request, Source_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedSourceServer) Sync2(*Sync2_Request, Source_Sync2Server) error {
	return status.Errorf(codes.Unimplemented, "method Sync2 not implemented")
}
func (UnimplementedSourceServer) GetMetrics(context.Context, *GetSourceMetrics_Request) (*GetSourceMetrics_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedSourceServer) mustEmbedUnimplementedSourceServer() {}

// UnsafeSourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourceServer will
// result in compilation errors.
type UnsafeSourceServer interface {
	mustEmbedUnimplementedSourceServer()
}

func RegisterSourceServer(s grpc.ServiceRegistrar, srv SourceServer) {
	s.RegisterService(&Source_ServiceDesc, srv)
}

func _Source_GetProtocolVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProtocolVersion_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).GetProtocolVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/GetProtocolVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).GetProtocolVersion(ctx, req.(*GetProtocolVersion_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetName_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).GetName(ctx, req.(*GetName_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersion_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).GetVersion(ctx, req.(*GetVersion_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_GetTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTables_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).GetTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/GetTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).GetTables(ctx, req.(*GetTables_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_GetSyncSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSyncSummary_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).GetSyncSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/GetSyncSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).GetSyncSummary(ctx, req.(*GetSyncSummary_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Sync_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceServer).Sync(m, &sourceSyncServer{stream})
}

type Source_SyncServer interface {
	Send(*Sync_Response) error
	grpc.ServerStream
}

type sourceSyncServer struct {
	grpc.ServerStream
}

func (x *sourceSyncServer) Send(m *Sync_Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Source_Sync2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Sync2_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceServer).Sync2(m, &sourceSync2Server{stream})
}

type Source_Sync2Server interface {
	Send(*Sync2_Response) error
	grpc.ServerStream
}

type sourceSync2Server struct {
	grpc.ServerStream
}

func (x *sourceSync2Server) Send(m *Sync2_Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Source_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSourceMetrics_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).GetMetrics(ctx, req.(*GetSourceMetrics_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Source_ServiceDesc is the grpc.ServiceDesc for Source service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Source_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Source",
	HandlerType: (*SourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProtocolVersion",
			Handler:    _Source_GetProtocolVersion_Handler,
		},
		{
			MethodName: "GetName",
			Handler:    _Source_GetName_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Source_GetVersion_Handler,
		},
		{
			MethodName: "GetTables",
			Handler:    _Source_GetTables_Handler,
		},
		{
			MethodName: "GetSyncSummary",
			Handler:    _Source_GetSyncSummary_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _Source_GetMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _Source_Sync_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Sync2",
			Handler:       _Source_Sync2_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/pb/source.proto",
}
