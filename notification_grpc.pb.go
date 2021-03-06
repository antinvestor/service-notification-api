// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package notification_v1

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

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	//Out method for queueing massages as requested
	Out(ctx context.Context, in *MessageOut, opts ...grpc.CallOption) (*StatusResponse, error)
	//CommunicationStatus request to determine if notification is prepared or released
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	//QueueRelease method for releasing queued massages and returns if notification status if released
	Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	//In method is for client request for particular notification respones from system
	In(ctx context.Context, in *MessageIn, opts ...grpc.CallOption) (*StatusResponse, error)
	//Search method is for client request for particular notification details from system
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (NotificationService_SearchClient, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Out(ctx context.Context, in *MessageOut, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/Out", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) In(ctx context.Context, in *MessageIn, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/In", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (NotificationService_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[0], "/notification.NotificationService/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_SearchClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type notificationServiceSearchClient struct {
	grpc.ClientStream
}

func (x *notificationServiceSearchClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	//Out method for queueing massages as requested
	Out(context.Context, *MessageOut) (*StatusResponse, error)
	//CommunicationStatus request to determine if notification is prepared or released
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	//QueueRelease method for releasing queued massages and returns if notification status if released
	Release(context.Context, *ReleaseRequest) (*StatusResponse, error)
	//In method is for client request for particular notification respones from system
	In(context.Context, *MessageIn) (*StatusResponse, error)
	//Search method is for client request for particular notification details from system
	Search(*SearchRequest, NotificationService_SearchServer) error
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) Out(context.Context, *MessageOut) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Out not implemented")
}
func (UnimplementedNotificationServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedNotificationServiceServer) Release(context.Context, *ReleaseRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (UnimplementedNotificationServiceServer) In(context.Context, *MessageIn) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method In not implemented")
}
func (UnimplementedNotificationServiceServer) Search(*SearchRequest, NotificationService_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_Out_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Out(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/Out",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Out(ctx, req.(*MessageOut))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Release(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_In_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).In(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/In",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).In(ctx, req.(*MessageIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).Search(m, &notificationServiceSearchServer{stream})
}

type NotificationService_SearchServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type notificationServiceSearchServer struct {
	grpc.ServerStream
}

func (x *notificationServiceSearchServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Out",
			Handler:    _NotificationService_Out_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _NotificationService_Status_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _NotificationService_Release_Handler,
		},
		{
			MethodName: "In",
			Handler:    _NotificationService_In_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _NotificationService_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notification.proto",
}
