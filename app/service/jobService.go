package service

import (
	"github.com/golang/protobuf/ptypes"
	empty_proto "github.com/golang/protobuf/ptypes/empty"
	jobProto "github.com/scaumiao/autotest/proto/job"
	"github.com/scaumiao/autotest/utils"
	"golang.org/x/net/context"
)

func (s *Service) GetJobList(ctx context.Context, empty *empty_proto.Empty) (*jobProto.JobList, error) {
	var result = &jobProto.JobList{}
	s.Api.JobStore.GetJobList(result)
	s.Api.LogServer.Info(result, "GetJobList success")
	return result, nil
}
func (s *Service) CreateJob(ctx context.Context, job *jobProto.Job) (*jobProto.Job, error) {
	id := utils.NewID()
	_job := &jobProto.Job{
		Id:      id,
		TaskId:  job.Id,
		Script:  job.Script, //task.Script
		StartAt: ptypes.TimestampNow(),
	}
	s.Api.JobStore.CreateJob(_job)
	s.Api.LogServer.Info(_job, "CreateJob")
	return _job, nil
}

func (s *Service) UpdateJob(ctx context.Context, job *jobProto.Job) (*jobProto.Job, error) {
	result := s.Api.JobStore.UpdateJob(job)
	if result == nil {
		return &jobProto.Job{}, nil
	}
	return result, nil
}
