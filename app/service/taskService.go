package service

import (
	"errors"
	"io/ioutil"

	"github.com/golang/protobuf/ptypes"
	empty_proto "github.com/golang/protobuf/ptypes/empty"
	grpcClient "github.com/scaumiao/autotest/app/client"
	jobProto "github.com/scaumiao/autotest/proto/job"
	taskProto "github.com/scaumiao/autotest/proto/task"

	"github.com/scaumiao/autotest/utils"
	"golang.org/x/net/context"
)

// 创建任务
func (s *Service) CreateTask(ctx context.Context, task *taskProto.Task) (*taskProto.Task, error) {
	// data := []*usr.UserInfo{}
	id := utils.NewID()
	data := &taskProto.Task{
		Id:     id,
		Name:   task.Name,
		Script: task.Script,
	}
	s.Api.TaskStore.CreateTask(data)

	// 自动生成ID data.Id = xxx
	// 调用TaskStore的CreateTask方法保存
	// 返回保存值
	return data, nil
}

// 获取任务列表
func (s *Service) GetTaskList(ctx context.Context, empty *empty_proto.Empty) (*taskProto.TaskList, error) {
	var result = &taskProto.TaskList{}
	s.Api.TaskStore.GetTaskList(result)
	return result, nil
}

// 删除
func (s *Service) DeleteTask(ctx context.Context, task *taskProto.Task) (*taskProto.Task, error) {

	data := &taskProto.Task{
		Id: task.Id,
	}

	s.Api.TaskStore.DeleteTask(task.Id)

	return data, nil
}

func (s *Service) UpdateTask(ctx context.Context, task *taskProto.Task) (*taskProto.Task, error) {
	result := s.Api.TaskStore.UpdateTask(task)
	if result == nil {
		return &taskProto.Task{}, nil
	}
	return result, nil
}

func (s *Service) GetTask(ctx context.Context, task *taskProto.Task) (*taskProto.Task, error) {

	data := s.Api.TaskStore.GetTask(task.Id)
	// if data == nil {
	// 	return &taskProto.Task{}, nil
	// }
	// fmt.Println(data)
	return data, nil
}

func (s *Service) RunTask(ctx context.Context, task *taskProto.Task) (*jobProto.Job, error) {

	// get task
	data := s.Api.TaskStore.GetTask(task.Id)
	if data.Id == "" {
		err := errors.New("taskId not exit")
		return nil, err
	}
	if data.Script == "" {
		err := errors.New("script can not be empty")
		return nil, err
	}
	serv := s.Api.TestServer
	serv.Start()
	defer serv.Stop()

	// connect grpc
	cli := grpcClient.NewJobClient(s.Api.Config.GrpcHost + ":" + s.Api.Config.GrpcPort)
	err := cli.Start()
	if err != nil {
		return nil, err
	}
	c := cli.GetJobClient()

	// create job
	job, err := c.CreateJob(ctx, &jobProto.Job{
		TaskId:  task.Id,
		Script:  data.Script, //task.Script
		StartAt: ptypes.TimestampNow(),
	})
	if err != nil {
		return nil, err
	}

	// run test
	rs := serv.RunTest(data.Script)
	// update job
	job.EndAt = ptypes.TimestampNow()
	job.Stages = rs
	j, err := c.UpdateJob(ctx, job)
	if err != nil {
		return nil, err
	}
	// stop client
	defer cli.Stop()

	return j, nil
}

func getTestFile(path string) string {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(bs)
}
