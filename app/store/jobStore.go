package store

import (
	taskProto "github.com/scaumiao/autotest/proto/task"
)

type JobStore struct {
	store Store
}

func NewJobStore() *JobStore {
	store := &JobStore{}
	return store
}

func (s *JobStore) SetStore(store Store) {
	s.store = store
}

func (s *JobStore) CreateJob(job *taskProto.Job) {
	s.store.Write("job", &job)
}

func (s *JobStore) GetJob(id string) *taskProto.Job {
	tmp := &taskProto.Job{}
	s.store.Read("job", id, &tmp)
	return tmp
}

func (s *JobStore) UpdateJob(job *taskProto.Job) *taskProto.Job {
	result := &taskProto.Job{}
	s.store.Update("job", job.Id, &job, &result)
	return result
}

func (s *JobStore) DeleteJob(id string) {
	s.store.Delete("job", id)
}

func (s *JobStore) GetJobList(jobList *taskProto.JobList) {
	var total []*taskProto.Job
	s.store.ReadAll("job", &total)
	jobList.Jobs = total
}
