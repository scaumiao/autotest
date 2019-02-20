// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stages.proto

/*
Package stages is a generated protocol buffer package.

It is generated from these files:
	stages.proto

It has these top-level messages:
	CreateStagesRequest
	CreateStagesResponse
	SearchStagesRequest
	SearchStagesResponse
	UpdateStagesRequest
	UpdateStagesResponse
	RemoveStagesRequest
	RemoveStagesResponse
	Stages
	Stage
	Action
*/
package stages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type State int32

const (
	State_ok   State = 0
	State_fail State = 1
)

var State_name = map[int32]string{
	0: "ok",
	1: "fail",
}
var State_value = map[string]int32{
	"ok":   0,
	"fail": 1,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateStagesRequest struct {
	Stages string `protobuf:"bytes,1,opt,name=stages" json:"stages,omitempty"`
}

func (m *CreateStagesRequest) Reset()                    { *m = CreateStagesRequest{} }
func (m *CreateStagesRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateStagesRequest) ProtoMessage()               {}
func (*CreateStagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateStagesRequest) GetStages() string {
	if m != nil {
		return m.Stages
	}
	return ""
}

type CreateStagesResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateStagesResponse) Reset()                    { *m = CreateStagesResponse{} }
func (m *CreateStagesResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateStagesResponse) ProtoMessage()               {}
func (*CreateStagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateStagesResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchStagesRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchStagesRequest) Reset()                    { *m = SearchStagesRequest{} }
func (m *SearchStagesRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchStagesRequest) ProtoMessage()               {}
func (*SearchStagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SearchStagesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SearchStagesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchStagesRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SearchStagesResponse struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Total  int32   `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	Stages *Stages `protobuf:"bytes,3,opt,name=stages" json:"stages,omitempty"`
}

func (m *SearchStagesResponse) Reset()                    { *m = SearchStagesResponse{} }
func (m *SearchStagesResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchStagesResponse) ProtoMessage()               {}
func (*SearchStagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SearchStagesResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SearchStagesResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *SearchStagesResponse) GetStages() *Stages {
	if m != nil {
		return m.Stages
	}
	return nil
}

type UpdateStagesRequest struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Raw string `protobuf:"bytes,2,opt,name=raw" json:"raw,omitempty"`
}

func (m *UpdateStagesRequest) Reset()                    { *m = UpdateStagesRequest{} }
func (m *UpdateStagesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateStagesRequest) ProtoMessage()               {}
func (*UpdateStagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateStagesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateStagesRequest) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

type UpdateStagesResponse struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *UpdateStagesResponse) Reset()                    { *m = UpdateStagesResponse{} }
func (m *UpdateStagesResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateStagesResponse) ProtoMessage()               {}
func (*UpdateStagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateStagesResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateStagesResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RemoveStagesRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RemoveStagesRequest) Reset()                    { *m = RemoveStagesRequest{} }
func (m *RemoveStagesRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveStagesRequest) ProtoMessage()               {}
func (*RemoveStagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RemoveStagesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveStagesResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *RemoveStagesResponse) Reset()                    { *m = RemoveStagesResponse{} }
func (m *RemoveStagesResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveStagesResponse) ProtoMessage()               {}
func (*RemoveStagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RemoveStagesResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Stages struct {
	Stages []*Stage `protobuf:"bytes,1,rep,name=stages" json:"stages,omitempty"`
}

func (m *Stages) Reset()                    { *m = Stages{} }
func (m *Stages) String() string            { return proto.CompactTextString(m) }
func (*Stages) ProtoMessage()               {}
func (*Stages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Stages) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

type Stage struct {
	Id          string                      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string                      `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Actions     []*Action                   `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
	State       State                       `protobuf:"varint,4,opt,name=state,enum=autotest.protobuf.State" json:"state,omitempty"`
	Line        int32                       `protobuf:"varint,5,opt,name=line" json:"line,omitempty"`
	Image       string                      `protobuf:"bytes,6,opt,name=image" json:"image,omitempty"`
	StartAt     *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=start_at,json=startAt" json:"start_at,omitempty"`
	EndAt       *google_protobuf1.Timestamp `protobuf:"bytes,8,opt,name=end_at,json=endAt" json:"end_at,omitempty"`
	TimeElapsed float64                     `protobuf:"fixed64,9,opt,name=time_elapsed,json=timeElapsed" json:"time_elapsed,omitempty"`
}

func (m *Stage) Reset()                    { *m = Stage{} }
func (m *Stage) String() string            { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()               {}
func (*Stage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Stage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Stage) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Stage) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Stage) GetState() State {
	if m != nil {
		return m.State
	}
	return State_ok
}

func (m *Stage) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *Stage) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Stage) GetStartAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartAt
	}
	return nil
}

func (m *Stage) GetEndAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EndAt
	}
	return nil
}

func (m *Stage) GetTimeElapsed() float64 {
	if m != nil {
		return m.TimeElapsed
	}
	return 0
}

type Action struct {
	Msg         string                      `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Command     string                      `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	State       State                       `protobuf:"varint,3,opt,name=state,enum=autotest.protobuf.State" json:"state,omitempty"`
	Line        int32                       `protobuf:"varint,4,opt,name=line" json:"line,omitempty"`
	StartAt     *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=start_at,json=startAt" json:"start_at,omitempty"`
	EndAt       *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=end_at,json=endAt" json:"end_at,omitempty"`
	TimeElapsed float64                     `protobuf:"fixed64,7,opt,name=time_elapsed,json=timeElapsed" json:"time_elapsed,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Action) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Action) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Action) GetState() State {
	if m != nil {
		return m.State
	}
	return State_ok
}

func (m *Action) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *Action) GetStartAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartAt
	}
	return nil
}

func (m *Action) GetEndAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EndAt
	}
	return nil
}

func (m *Action) GetTimeElapsed() float64 {
	if m != nil {
		return m.TimeElapsed
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateStagesRequest)(nil), "autotest.protobuf.CreateStagesRequest")
	proto.RegisterType((*CreateStagesResponse)(nil), "autotest.protobuf.CreateStagesResponse")
	proto.RegisterType((*SearchStagesRequest)(nil), "autotest.protobuf.SearchStagesRequest")
	proto.RegisterType((*SearchStagesResponse)(nil), "autotest.protobuf.SearchStagesResponse")
	proto.RegisterType((*UpdateStagesRequest)(nil), "autotest.protobuf.UpdateStagesRequest")
	proto.RegisterType((*UpdateStagesResponse)(nil), "autotest.protobuf.UpdateStagesResponse")
	proto.RegisterType((*RemoveStagesRequest)(nil), "autotest.protobuf.RemoveStagesRequest")
	proto.RegisterType((*RemoveStagesResponse)(nil), "autotest.protobuf.RemoveStagesResponse")
	proto.RegisterType((*Stages)(nil), "autotest.protobuf.Stages")
	proto.RegisterType((*Stage)(nil), "autotest.protobuf.Stage")
	proto.RegisterType((*Action)(nil), "autotest.protobuf.Action")
	proto.RegisterEnum("autotest.protobuf.State", State_name, State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StagesSvic service

type StagesSvicClient interface {
	CreateStages(ctx context.Context, in *CreateStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error)
	GetStages(ctx context.Context, in *SearchStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error)
	ListStages(ctx context.Context, in *SearchStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error)
	UpdateStages(ctx context.Context, in *UpdateStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error)
	RemoveStages(ctx context.Context, in *CreateStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error)
}

type stagesSvicClient struct {
	cc *grpc.ClientConn
}

func NewStagesSvicClient(cc *grpc.ClientConn) StagesSvicClient {
	return &stagesSvicClient{cc}
}

func (c *stagesSvicClient) CreateStages(ctx context.Context, in *CreateStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error) {
	out := new(CreateStagesResponse)
	err := grpc.Invoke(ctx, "/autotest.protobuf.StagesSvic/CreateStages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagesSvicClient) GetStages(ctx context.Context, in *SearchStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error) {
	out := new(CreateStagesResponse)
	err := grpc.Invoke(ctx, "/autotest.protobuf.StagesSvic/GetStages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagesSvicClient) ListStages(ctx context.Context, in *SearchStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error) {
	out := new(CreateStagesResponse)
	err := grpc.Invoke(ctx, "/autotest.protobuf.StagesSvic/ListStages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagesSvicClient) UpdateStages(ctx context.Context, in *UpdateStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error) {
	out := new(CreateStagesResponse)
	err := grpc.Invoke(ctx, "/autotest.protobuf.StagesSvic/UpdateStages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagesSvicClient) RemoveStages(ctx context.Context, in *CreateStagesRequest, opts ...grpc.CallOption) (*CreateStagesResponse, error) {
	out := new(CreateStagesResponse)
	err := grpc.Invoke(ctx, "/autotest.protobuf.StagesSvic/RemoveStages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StagesSvic service

type StagesSvicServer interface {
	CreateStages(context.Context, *CreateStagesRequest) (*CreateStagesResponse, error)
	GetStages(context.Context, *SearchStagesRequest) (*CreateStagesResponse, error)
	ListStages(context.Context, *SearchStagesRequest) (*CreateStagesResponse, error)
	UpdateStages(context.Context, *UpdateStagesRequest) (*CreateStagesResponse, error)
	RemoveStages(context.Context, *CreateStagesRequest) (*CreateStagesResponse, error)
}

func RegisterStagesSvicServer(s *grpc.Server, srv StagesSvicServer) {
	s.RegisterService(&_StagesSvic_serviceDesc, srv)
}

func _StagesSvic_CreateStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagesSvicServer).CreateStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.StagesSvic/CreateStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagesSvicServer).CreateStages(ctx, req.(*CreateStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagesSvic_GetStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagesSvicServer).GetStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.StagesSvic/GetStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagesSvicServer).GetStages(ctx, req.(*SearchStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagesSvic_ListStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagesSvicServer).ListStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.StagesSvic/ListStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagesSvicServer).ListStages(ctx, req.(*SearchStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagesSvic_UpdateStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagesSvicServer).UpdateStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.StagesSvic/UpdateStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagesSvicServer).UpdateStages(ctx, req.(*UpdateStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagesSvic_RemoveStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagesSvicServer).RemoveStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.StagesSvic/RemoveStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagesSvicServer).RemoveStages(ctx, req.(*CreateStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StagesSvic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "autotest.protobuf.StagesSvic",
	HandlerType: (*StagesSvicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStages",
			Handler:    _StagesSvic_CreateStages_Handler,
		},
		{
			MethodName: "GetStages",
			Handler:    _StagesSvic_GetStages_Handler,
		},
		{
			MethodName: "ListStages",
			Handler:    _StagesSvic_ListStages_Handler,
		},
		{
			MethodName: "UpdateStages",
			Handler:    _StagesSvic_UpdateStages_Handler,
		},
		{
			MethodName: "RemoveStages",
			Handler:    _StagesSvic_RemoveStages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stages.proto",
}

func init() { proto.RegisterFile("stages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0xec, 0x34, 0xd3, 0xa8, 0xb4, 0x1b, 0xab, 0xb8, 0x51, 0xa5, 0x16, 0x4b, 0x94,
	0xaa, 0x12, 0x71, 0x7f, 0x84, 0x40, 0xe5, 0x54, 0x10, 0xe2, 0xc2, 0xc9, 0x81, 0x0b, 0x97, 0x6a,
	0x1b, 0x6f, 0xd2, 0x15, 0xb1, 0xd7, 0xd8, 0x9b, 0x20, 0x51, 0x21, 0x21, 0x24, 0x9e, 0x80, 0x47,
	0x43, 0xbc, 0x01, 0xaf, 0x81, 0xc4, 0xfe, 0xd8, 0x8d, 0x93, 0x38, 0x6d, 0xa9, 0x04, 0xa7, 0xec,
	0x8c, 0xbf, 0x99, 0x6f, 0xe7, 0x9b, 0x9d, 0x09, 0x34, 0x53, 0x8e, 0x07, 0x24, 0xed, 0xc4, 0x09,
	0xe3, 0x0c, 0xad, 0xe1, 0x11, 0x67, 0x9c, 0xa4, 0x5c, 0xdb, 0x67, 0xa3, 0x7e, 0x7b, 0x73, 0xc0,
	0xd8, 0x60, 0x48, 0x3c, 0x1c, 0x53, 0x0f, 0x47, 0x11, 0xe3, 0x98, 0x53, 0x16, 0x65, 0x01, 0xed,
	0xad, 0xec, 0x6b, 0x0e, 0xf7, 0x38, 0x0d, 0x45, 0x38, 0x0e, 0x63, 0x0d, 0x70, 0x1f, 0x41, 0xeb,
	0x45, 0x42, 0x30, 0x27, 0x5d, 0xc5, 0xe3, 0x93, 0x0f, 0x23, 0xf1, 0x1d, 0xad, 0x83, 0xa5, 0x89,
	0x1d, 0x63, 0xdb, 0xd8, 0x6d, 0xf8, 0x99, 0xe5, 0xee, 0x80, 0x3d, 0x0d, 0x4f, 0x63, 0x41, 0x46,
	0xd0, 0x0a, 0x54, 0x68, 0x90, 0x61, 0xc5, 0xc9, 0xed, 0x42, 0xab, 0x4b, 0x70, 0xd2, 0x3b, 0x9f,
	0x4e, 0x3b, 0x03, 0x43, 0x36, 0x98, 0x43, 0x1a, 0x52, 0xee, 0x54, 0x85, 0xcb, 0xf4, 0xb5, 0x21,
	0xc9, 0x59, 0xbf, 0x9f, 0x12, 0xee, 0xd4, 0x94, 0x3b, 0xb3, 0x5c, 0x06, 0xf6, 0x74, 0xd2, 0x72,
	0x72, 0x99, 0x55, 0xa8, 0x84, 0x87, 0x4e, 0x45, 0x67, 0x55, 0x06, 0x3a, 0xb8, 0x2c, 0x49, 0x92,
	0x2d, 0x1f, 0x6e, 0x74, 0xe6, 0xc4, 0xec, 0x64, 0x89, 0xf3, 0x6a, 0x9f, 0x40, 0xeb, 0x6d, 0x1c,
	0xcc, 0x89, 0x33, 0xcb, 0xb7, 0x0a, 0xd5, 0x04, 0x7f, 0x54, 0x6c, 0x0d, 0x5f, 0x1e, 0xdd, 0xa7,
	0x60, 0x4f, 0x07, 0x2e, 0xb8, 0xa9, 0x88, 0x0c, 0xd3, 0x41, 0x1e, 0x29, 0x8e, 0xee, 0x03, 0x68,
	0xf9, 0x24, 0x64, 0xe3, 0xab, 0x29, 0xdd, 0x5d, 0xb0, 0xa7, 0x61, 0x19, 0x41, 0x96, 0xd0, 0x98,
	0x24, 0x3c, 0x06, 0x4b, 0x63, 0xd0, 0x7e, 0xa1, 0xa7, 0x55, 0x21, 0x80, 0xb3, 0x48, 0x80, 0xcb,
	0xfa, 0x7f, 0x56, 0xc0, 0x54, 0x9e, 0xb9, 0x8b, 0x6f, 0xc3, 0x72, 0x40, 0xd2, 0x5e, 0x42, 0x63,
	0xf9, 0xda, 0xb2, 0x02, 0x8a, 0x2e, 0x74, 0x04, 0x75, 0xdc, 0x53, 0x4f, 0x51, 0xe8, 0x5d, 0x5d,
	0xa0, 0xf7, 0x89, 0x42, 0xf8, 0x39, 0x12, 0x75, 0xc0, 0x14, 0xd4, 0x9c, 0xa8, 0xc6, 0xaf, 0x2c,
	0xba, 0x21, 0x27, 0xbe, 0x86, 0x21, 0x04, 0xb5, 0x21, 0x8d, 0x88, 0x63, 0xaa, 0x46, 0xab, 0xb3,
	0xec, 0x3e, 0x0d, 0xc5, 0x9d, 0x1d, 0x4b, 0x5d, 0x4a, 0x1b, 0xe8, 0x31, 0x2c, 0x89, 0x90, 0x84,
	0x9f, 0x62, 0xee, 0xd4, 0x55, 0xff, 0xdb, 0x1d, 0x3d, 0x1b, 0x93, 0xd4, 0x6f, 0xf2, 0xd9, 0xf0,
	0xeb, 0x0a, 0x7b, 0xc2, 0xe5, 0xa3, 0x21, 0x51, 0x20, 0x83, 0x96, 0xae, 0x0d, 0x32, 0x05, 0x52,
	0x84, 0xdc, 0x87, 0xa6, 0x1c, 0xb2, 0x53, 0x32, 0xc4, 0x71, 0x4a, 0x02, 0xa7, 0x21, 0x02, 0x0d,
	0x7f, 0x59, 0xfa, 0x5e, 0x6a, 0x97, 0xfb, 0xad, 0x02, 0x96, 0x2e, 0x7d, 0xbe, 0x61, 0xc8, 0x81,
	0x7a, 0x8f, 0x85, 0x21, 0x8e, 0x82, 0x4c, 0xd6, 0xdc, 0x9c, 0xa8, 0x53, 0xfd, 0x3b, 0x75, 0x6a,
	0x05, 0x75, 0x8a, 0x3a, 0x98, 0xb7, 0xd1, 0xc1, 0xba, 0xad, 0x0e, 0xf5, 0x39, 0x1d, 0xf6, 0x36,
	0xd4, 0xf3, 0x12, 0x37, 0xb5, 0xa0, 0xc2, 0xde, 0xaf, 0xde, 0x41, 0x4b, 0x50, 0xeb, 0x63, 0x3a,
	0x5c, 0x35, 0x0e, 0x7f, 0xd7, 0x00, 0xf4, 0xbb, 0xed, 0x8e, 0x69, 0x0f, 0x5d, 0x40, 0xb3, 0xb8,
	0x77, 0xd0, 0x4e, 0x49, 0xed, 0x25, 0x7b, 0xac, 0xfd, 0xf0, 0x5a, 0x9c, 0x1e, 0x1c, 0xb7, 0xfd,
	0xf5, 0xc7, 0xaf, 0xef, 0x15, 0xdb, 0xbd, 0xeb, 0x8d, 0x0f, 0x3c, 0x89, 0xf7, 0xf4, 0x0c, 0x1c,
	0x1b, 0x7b, 0xe8, 0x13, 0x34, 0x5e, 0x11, 0x7e, 0x05, 0x73, 0xc9, 0xaa, 0xbb, 0x39, 0xf3, 0xa6,
	0x62, 0x5e, 0x47, 0xf6, 0x0c, 0xb3, 0x77, 0x41, 0x83, 0xcf, 0x68, 0x0c, 0xf0, 0x9a, 0xa6, 0xff,
	0x8c, 0xfc, 0x9e, 0x22, 0x5f, 0x43, 0xb3, 0x65, 0xa3, 0x2f, 0x06, 0x34, 0x8b, 0x2b, 0xac, 0x94,
	0xba, 0x64, 0x39, 0xde, 0x9c, 0x7a, 0x4b, 0x51, 0x6f, 0xb4, 0x4b, 0xeb, 0x96, 0xb2, 0x8b, 0x9e,
	0x17, 0x77, 0xdc, 0x7f, 0xed, 0xf9, 0xf3, 0xc3, 0x77, 0xfb, 0x03, 0xca, 0xcf, 0x47, 0x67, 0x1d,
	0x31, 0x7d, 0x5e, 0xda, 0xc3, 0xa3, 0x90, 0x62, 0xe6, 0xe5, 0x99, 0xf5, 0x1f, 0x6a, 0x86, 0x7f,
	0xa6, 0x7f, 0xce, 0x2c, 0xe5, 0x3c, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x72, 0xad, 0xfe, 0xea,
	0xb4, 0x07, 0x00, 0x00,
}