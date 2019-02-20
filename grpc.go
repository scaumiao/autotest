package autotest

import (
	"log"
	"net"

	service "github.com/scaumiao/autotest/app/service"

	taskProto "github.com/scaumiao/autotest/proto/task"

	"google.golang.org/grpc"
)

func (serv *Server) GrpcServerStart(grpcPort string) error {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatal("failed to listen: ", err)
		return err
	}
	s := grpc.NewServer()

	_taskService := &service.Service{}

	taskProto.RegisterTaskServiceServer(s, _taskService)
	return s.Serve(lis)
}
