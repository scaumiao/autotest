package store

import (
	taskProto "github.com/scaumiao/autotest/proto/task"
)

type TaskStore struct {
	store Store
}

func NewTaskStore() *TaskStore {
	store := &TaskStore{}
	return store
}

func (s *TaskStore) SetStore(store Store) {
	s.store = store
}

func (s *TaskStore) CreateTask(task *taskProto.Task) {
	s.store.Write("task", &task)
}

func (s *TaskStore) GetTaskList(taskList *taskProto.TaskList) {
	var total []*taskProto.Task
	s.store.ReadAll("task", &total)
	taskList.Tasks = total
}

func (s *TaskStore) GetTask(id string) *taskProto.Task {
	tmp := &taskProto.Task{}
	s.store.Read("task", id, &tmp)
	return tmp
}

func (s *TaskStore) UpdateTask(task *taskProto.Task) *taskProto.Task {
	result := &taskProto.Task{}
	s.store.Update("task", task.Id, &task, &result)
	return result
}

func (s *TaskStore) DeleteTask(id string) {
	s.store.Delete("task", id)
}
