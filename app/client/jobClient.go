package client

import (
	"context"
	"log"

	jobProto "github.com/scaumiao/autotest/proto/job"

	"google.golang.org/grpc"
)

type JobClient struct {
	conn    *grpc.ClientConn
	ctx     context.Context
	address string
}

func NewJobClient(addr string) *JobClient {
	cli := &JobClient{
		address: addr,
	}
	return cli
}

func (cli *JobClient) Start() error {
	conn, err := grpc.Dial(cli.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	cli.conn = conn
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return nil
}

func (cli *JobClient) GetJobClient() jobProto.JobServiceClient {
	c := jobProto.NewJobServiceClient(cli.conn)
	return c
}

func (cli *JobClient) Stop() {
	cli.conn.Close()
}
