// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

/*
Package task is a generated protocol buffer package.

It is generated from these files:
	task.proto

It has these top-level messages:
	Task
	TaskList
*/
package task

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

// message CreateTaskRequest {
//     string name = 1;
// }
// message CreateTaskResponse {
//     string id = 1;
// }
// message UpdateTaskRequest {
//     string id = 1;
//     string name = 2;
// }
type Task struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TaskList struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TaskList) Reset()                    { *m = TaskList{} }
func (m *TaskList) String() string            { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()               {}
func (*TaskList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskList) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "autotest.protobuf.Task")
	proto.RegisterType((*TaskList)(nil), "autotest.protobuf.TaskList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskService service

type TaskServiceClient interface {
	CreateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	GetTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	GetTaskList(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*TaskList, error)
	DeleteTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	UpdateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/autotest.protobuf.TaskService/CreateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/autotest.protobuf.TaskService/GetTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTaskList(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*TaskList, error) {
	out := new(TaskList)
	err := grpc.Invoke(ctx, "/autotest.protobuf.TaskService/GetTaskList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/autotest.protobuf.TaskService/DeleteTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/autotest.protobuf.TaskService/UpdateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceServer interface {
	CreateTask(context.Context, *Task) (*Task, error)
	GetTask(context.Context, *Task) (*Task, error)
	GetTaskList(context.Context, *google_protobuf1.Empty) (*TaskList, error)
	DeleteTask(context.Context, *Task) (*Task, error)
	UpdateTask(context.Context, *Task) (*Task, error)
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.TaskService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.TaskService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.TaskService/GetTaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTaskList(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.TaskService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autotest.protobuf.TaskService/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "autotest.protobuf.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
		{
			MethodName: "GetTaskList",
			Handler:    _TaskService_GetTaskList_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TaskService_DeleteTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskService_UpdateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0x3b, 0x31,
	0x14, 0xc5, 0xe9, 0xc7, 0xff, 0x6f, 0x7b, 0x87, 0x2a, 0x8d, 0xa8, 0x65, 0xea, 0x42, 0x66, 0x55,
	0x8a, 0x26, 0x58, 0x57, 0xea, 0xce, 0x0f, 0xdc, 0xb8, 0x6a, 0x2d, 0x88, 0xbb, 0xb4, 0x8d, 0x35,
	0xd8, 0x99, 0x0c, 0xcd, 0x9d, 0x82, 0x88, 0x1b, 0x5f, 0xc1, 0xf7, 0xf1, 0x25, 0x7c, 0x05, 0x1f,
	0xc4, 0x49, 0x32, 0x83, 0xa8, 0xcc, 0x66, 0x36, 0x43, 0x92, 0x7b, 0xce, 0x2f, 0x27, 0x87, 0x01,
	0x40, 0xae, 0x1f, 0x69, 0xbc, 0x54, 0xa8, 0x48, 0x9b, 0x27, 0xa8, 0x50, 0x68, 0x74, 0xfb, 0x49,
	0x72, 0xef, 0xef, 0xce, 0x95, 0x9a, 0x2f, 0x04, 0xe3, 0xb1, 0x64, 0x3c, 0x8a, 0x14, 0x72, 0x94,
	0x2a, 0xd2, 0x4e, 0xe0, 0x77, 0xb3, 0x69, 0x2e, 0x67, 0x22, 0x8c, 0xf1, 0xc9, 0x0d, 0x83, 0x3e,
	0xd4, 0x6f, 0x52, 0x36, 0x59, 0x87, 0xaa, 0x9c, 0x75, 0x2a, 0x7b, 0x95, 0x5e, 0x73, 0x98, 0xae,
	0x08, 0x81, 0x7a, 0xc4, 0x43, 0xd1, 0xa9, 0xda, 0x13, 0xbb, 0x0e, 0x8e, 0xa1, 0x61, 0xb4, 0xd7,
	0x52, 0x23, 0x39, 0x80, 0x7f, 0x26, 0x93, 0x4e, 0x2d, 0xb5, 0x9e, 0x37, 0xd8, 0xa1, 0x7f, 0x52,
	0x51, 0xa3, 0x1d, 0x3a, 0xd5, 0xe0, 0xbd, 0x06, 0x9e, 0xd9, 0x8f, 0xc4, 0x72, 0x25, 0xa7, 0x82,
	0x8c, 0x00, 0xce, 0x97, 0x82, 0xa3, 0xb0, 0x97, 0x17, 0xb9, 0xfd, 0xa2, 0x41, 0xb0, 0xf9, 0xfa,
	0xf1, 0xf9, 0x56, 0x6d, 0x05, 0x0d, 0xb6, 0x3a, 0x64, 0xe6, 0x8a, 0x93, 0x4a, 0x9f, 0x0c, 0x61,
	0xed, 0x4a, 0x60, 0x49, 0xe2, 0x96, 0x25, 0x6e, 0x90, 0x56, 0x4e, 0x64, 0xcf, 0x72, 0xf6, 0x42,
	0xc6, 0xe0, 0x65, 0x4c, 0xfb, 0xec, 0x6d, 0xea, 0xca, 0xfc, 0x36, 0x5f, 0x9a, 0x32, 0xfd, 0x6e,
	0x01, 0xd6, 0x98, 0x82, 0xb6, 0x45, 0x7b, 0xa4, 0x99, 0xa3, 0x75, 0x8a, 0x85, 0x0b, 0xb1, 0x10,
	0xa5, 0xdf, 0x9f, 0xa5, 0xed, 0xff, 0x4a, 0x7b, 0x0b, 0x30, 0x8e, 0x67, 0xe5, 0x6b, 0xed, 0x58,
	0x2c, 0xf1, 0x7f, 0x62, 0xd3, 0x6e, 0xcf, 0xe8, 0xdd, 0xfe, 0x5c, 0xe2, 0x43, 0x32, 0xa1, 0x53,
	0x15, 0x32, 0x3d, 0xe5, 0x49, 0x28, 0xb9, 0x62, 0x39, 0xc7, 0xfd, 0x5c, 0xd6, 0x72, 0x6a, 0x3e,
	0x93, 0xff, 0xf6, 0xe0, 0xe8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x19, 0xe6, 0x96, 0xba, 0x02,
	0x00, 0x00,
}
