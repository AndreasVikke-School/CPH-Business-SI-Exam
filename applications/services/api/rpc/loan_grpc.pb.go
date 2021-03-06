// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LoanServiceClient is the client API for LoanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoanServiceClient interface {
	GetLoan(ctx context.Context, in *wrapperspb.Int64Value, opts ...grpc.CallOption) (*Loan, error)
	GetAllLoans(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LoanList, error)
	GetAllLoansByUser(ctx context.Context, in *wrapperspb.Int64Value, opts ...grpc.CallOption) (*LoanList, error)
	CreateLoan(ctx context.Context, in *Loan, opts ...grpc.CallOption) (*Loan, error)
}

type loanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoanServiceClient(cc grpc.ClientConnInterface) LoanServiceClient {
	return &loanServiceClient{cc}
}

func (c *loanServiceClient) GetLoan(ctx context.Context, in *wrapperspb.Int64Value, opts ...grpc.CallOption) (*Loan, error) {
	out := new(Loan)
	err := c.cc.Invoke(ctx, "/rpc.LoanService/GetLoan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loanServiceClient) GetAllLoans(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LoanList, error) {
	out := new(LoanList)
	err := c.cc.Invoke(ctx, "/rpc.LoanService/GetAllLoans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loanServiceClient) GetAllLoansByUser(ctx context.Context, in *wrapperspb.Int64Value, opts ...grpc.CallOption) (*LoanList, error) {
	out := new(LoanList)
	err := c.cc.Invoke(ctx, "/rpc.LoanService/GetAllLoansByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loanServiceClient) CreateLoan(ctx context.Context, in *Loan, opts ...grpc.CallOption) (*Loan, error) {
	out := new(Loan)
	err := c.cc.Invoke(ctx, "/rpc.LoanService/CreateLoan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoanServiceServer is the server API for LoanService service.
// All implementations must embed UnimplementedLoanServiceServer
// for forward compatibility
type LoanServiceServer interface {
	GetLoan(context.Context, *wrapperspb.Int64Value) (*Loan, error)
	GetAllLoans(context.Context, *emptypb.Empty) (*LoanList, error)
	GetAllLoansByUser(context.Context, *wrapperspb.Int64Value) (*LoanList, error)
	CreateLoan(context.Context, *Loan) (*Loan, error)
	mustEmbedUnimplementedLoanServiceServer()
}

// UnimplementedLoanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoanServiceServer struct {
}

func (UnimplementedLoanServiceServer) GetLoan(context.Context, *wrapperspb.Int64Value) (*Loan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoan not implemented")
}
func (UnimplementedLoanServiceServer) GetAllLoans(context.Context, *emptypb.Empty) (*LoanList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLoans not implemented")
}
func (UnimplementedLoanServiceServer) GetAllLoansByUser(context.Context, *wrapperspb.Int64Value) (*LoanList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLoansByUser not implemented")
}
func (UnimplementedLoanServiceServer) CreateLoan(context.Context, *Loan) (*Loan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLoan not implemented")
}
func (UnimplementedLoanServiceServer) mustEmbedUnimplementedLoanServiceServer() {}

// UnsafeLoanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoanServiceServer will
// result in compilation errors.
type UnsafeLoanServiceServer interface {
	mustEmbedUnimplementedLoanServiceServer()
}

func RegisterLoanServiceServer(s grpc.ServiceRegistrar, srv LoanServiceServer) {
	s.RegisterService(&LoanService_ServiceDesc, srv)
}

func _LoanService_GetLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).GetLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LoanService/GetLoan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).GetLoan(ctx, req.(*wrapperspb.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoanService_GetAllLoans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).GetAllLoans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LoanService/GetAllLoans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).GetAllLoans(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoanService_GetAllLoansByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).GetAllLoansByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LoanService/GetAllLoansByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).GetAllLoansByUser(ctx, req.(*wrapperspb.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoanService_CreateLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Loan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).CreateLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LoanService/CreateLoan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).CreateLoan(ctx, req.(*Loan))
	}
	return interceptor(ctx, in, info, handler)
}

// LoanService_ServiceDesc is the grpc.ServiceDesc for LoanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.LoanService",
	HandlerType: (*LoanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLoan",
			Handler:    _LoanService_GetLoan_Handler,
		},
		{
			MethodName: "GetAllLoans",
			Handler:    _LoanService_GetAllLoans_Handler,
		},
		{
			MethodName: "GetAllLoansByUser",
			Handler:    _LoanService_GetAllLoansByUser_Handler,
		},
		{
			MethodName: "CreateLoan",
			Handler:    _LoanService_CreateLoan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loan.proto",
}
