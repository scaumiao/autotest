package store

import (
	jobProto "github.com/scaumiao/autotest/proto/job"
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

func (s *JobStore) CreateJob(job *jobProto.Job) {
	s.store.Write("job", &job)
}

func (s *JobStore) GetJob(id string) *jobProto.Job {
	tmp := &jobProto.Job{}
	s.store.Read("job", id, &tmp)
	return tmp
}

func (s *JobStore) UpdateJob(job *jobProto.Job) *jobProto.Job {
	result := &jobProto.Job{}
	s.store.Update("job", job.Id, &job, &result)
	return result
}

func (s *JobStore) DeleteJob(id string) {
	s.store.Delete("job", id)
}

func (s *JobStore) GetJobList(jobList *jobProto.JobList) {
	var total []*jobProto.Job
	s.store.ReadAll("job", &total)
	jobList.Jobs = total
}
