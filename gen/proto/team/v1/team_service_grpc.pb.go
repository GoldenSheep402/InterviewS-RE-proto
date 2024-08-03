// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: team/v1/team_service.proto

package teamV1

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
	TeamService_CreateTeam_FullMethodName      = "/team.v1.TeamService/CreateTeam"
	TeamService_UpdateTeam_FullMethodName      = "/team.v1.TeamService/UpdateTeam"
	TeamService_GetTeamInfo_FullMethodName     = "/team.v1.TeamService/GetTeamInfo"
	TeamService_GetTeamMembers_FullMethodName  = "/team.v1.TeamService/GetTeamMembers"
	TeamService_CreateGroup_FullMethodName     = "/team.v1.TeamService/CreateGroup"
	TeamService_GetGroupInfo_FullMethodName    = "/team.v1.TeamService/GetGroupInfo"
	TeamService_GetGroupMembers_FullMethodName = "/team.v1.TeamService/GetGroupMembers"
)

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	// CreateTeam only administrator can create team
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error)
	// UpdateTeam only administrator or its manager can update team
	UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error)
	// GetTeamInfo get team info if not manager only des
	GetTeamInfo(ctx context.Context, in *GetTeamInfoRequest, opts ...grpc.CallOption) (*GetTeamInfoResponse, error)
	// GetTeamMembers only manager
	GetTeamMembers(ctx context.Context, in *GetTeamMembersRequest, opts ...grpc.CallOption) (*GetTeamMembersResponse, error)
	// CreateGroup only team manager
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	// GetGroupInfo all user
	GetGroupInfo(ctx context.Context, in *GetGroupInfoRequest, opts ...grpc.CallOption) (*GetGroupInfoResponse, error)
	// GetGroupMembers only team manager
	GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTeamResponse)
	err := c.cc.Invoke(ctx, TeamService_CreateTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTeamResponse)
	err := c.cc.Invoke(ctx, TeamService_UpdateTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamInfo(ctx context.Context, in *GetTeamInfoRequest, opts ...grpc.CallOption) (*GetTeamInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamInfoResponse)
	err := c.cc.Invoke(ctx, TeamService_GetTeamInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamMembers(ctx context.Context, in *GetTeamMembersRequest, opts ...grpc.CallOption) (*GetTeamMembersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamMembersResponse)
	err := c.cc.Invoke(ctx, TeamService_GetTeamMembers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, TeamService_CreateGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetGroupInfo(ctx context.Context, in *GetGroupInfoRequest, opts ...grpc.CallOption) (*GetGroupInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupInfoResponse)
	err := c.cc.Invoke(ctx, TeamService_GetGroupInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupMembersResponse)
	err := c.cc.Invoke(ctx, TeamService_GetGroupMembers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
// All implementations must embed UnimplementedTeamServiceServer
// for forward compatibility.
type TeamServiceServer interface {
	// CreateTeam only administrator can create team
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error)
	// UpdateTeam only administrator or its manager can update team
	UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamResponse, error)
	// GetTeamInfo get team info if not manager only des
	GetTeamInfo(context.Context, *GetTeamInfoRequest) (*GetTeamInfoResponse, error)
	// GetTeamMembers only manager
	GetTeamMembers(context.Context, *GetTeamMembersRequest) (*GetTeamMembersResponse, error)
	// CreateGroup only team manager
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	// GetGroupInfo all user
	GetGroupInfo(context.Context, *GetGroupInfoRequest) (*GetGroupInfoResponse, error)
	// GetGroupMembers only team manager
	GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error)
	mustEmbedUnimplementedTeamServiceServer()
}

// UnimplementedTeamServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTeamServiceServer struct{}

func (UnimplementedTeamServiceServer) CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (UnimplementedTeamServiceServer) UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamInfo(context.Context, *GetTeamInfoRequest) (*GetTeamInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamInfo not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamMembers(context.Context, *GetTeamMembersRequest) (*GetTeamMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamMembers not implemented")
}
func (UnimplementedTeamServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedTeamServiceServer) GetGroupInfo(context.Context, *GetGroupInfoRequest) (*GetGroupInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupInfo not implemented")
}
func (UnimplementedTeamServiceServer) GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMembers not implemented")
}
func (UnimplementedTeamServiceServer) mustEmbedUnimplementedTeamServiceServer() {}
func (UnimplementedTeamServiceServer) testEmbeddedByValue()                     {}

// UnsafeTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServiceServer will
// result in compilation errors.
type UnsafeTeamServiceServer interface {
	mustEmbedUnimplementedTeamServiceServer()
}

func RegisterTeamServiceServer(s grpc.ServiceRegistrar, srv TeamServiceServer) {
	// If the following call pancis, it indicates UnimplementedTeamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TeamService_ServiceDesc, srv)
}

func _TeamService_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_CreateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).CreateTeam(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_UpdateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).UpdateTeam(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamInfo(ctx, req.(*GetTeamInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamMembers(ctx, req.(*GetTeamMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetGroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetGroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetGroupInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetGroupInfo(ctx, req.(*GetGroupInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetGroupMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetGroupMembers(ctx, req.(*GetGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamService_ServiceDesc is the grpc.ServiceDesc for TeamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "team.v1.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _TeamService_CreateTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _TeamService_UpdateTeam_Handler,
		},
		{
			MethodName: "GetTeamInfo",
			Handler:    _TeamService_GetTeamInfo_Handler,
		},
		{
			MethodName: "GetTeamMembers",
			Handler:    _TeamService_GetTeamMembers_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _TeamService_CreateGroup_Handler,
		},
		{
			MethodName: "GetGroupInfo",
			Handler:    _TeamService_GetGroupInfo_Handler,
		},
		{
			MethodName: "GetGroupMembers",
			Handler:    _TeamService_GetGroupMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team/v1/team_service.proto",
}
