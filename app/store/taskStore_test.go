package store

import (
	"fmt"
	"testing"

	"github.com/scaumiao/autotest/app/store/local"
	taskProto "github.com/scaumiao/autotest/proto/task"
)

func TestTaskStore_CreateTask(t *testing.T) {
	taskStore := NewTaskStore()
	localStore := local.NewLocalStore()
	taskStore.SetStore(localStore)
	task := &taskProto.Task{
		Name: "task1",
	}
	taskStore.CreateTask(task)
	var to []*taskProto.Task
	taskStore.GetTaskList(&to)
	fmt.Print(to)
}
