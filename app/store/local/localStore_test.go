package local

import (
	"fmt"
	"testing"

	taskProto "github.com/scaumiao/autotest/proto/task"
)

func TestLocalStore_Write(t *testing.T) {
	serv := NewLocalStore()
	task := &taskProto.Task{Name: "aaa", Id: "admin"}
	serv.Write("task", &task)
	task2 := &taskProto.Task{Name: "bbb"}
	serv.Write("task", &task2)
	// var out []*taskProto.Task
	// serv.ReadAll("task", &out)
	// if len(out) != 2 {
	// 	t.Error("error", out)
	// }
	var data *taskProto.Task
	serv.Read("task", "admin", &data)
	fmt.Println(data)
}

func TestLocalStore_Read(t *testing.T) {
	serv := NewLocalStore()
	task := taskProto.Task{Name: "aaa", Id: "admin"}
	serv.Write("task", &task)
	task2 := taskProto.Task{Name: "bbb"}
	serv.Write("task", &task2)
	// var out []*taskProto.Task
	// serv.ReadAll("task", &out)
	// if len(out) != 2 {
	// 	t.Error("error", out)
	// }
	var data taskProto.Task
	serv.Read("task", "admin", &data)
	fmt.Println(data)
}

func TestLocalStore_Update(t *testing.T) {
	serv := NewLocalStore()
	task := &taskProto.Task{Name: "aaa", Id: "admin"}
	serv.Write("task", &task)
	task2 := &taskProto.Task{Name: "bbb", Id: "admin"}
	// var out []*taskProto.Task
	// serv.ReadAll("task", &out)
	// if len(out) != 2 {
	// 	t.Error("error", out)
	// }
	var result *taskProto.Task
	serv.Update("task", "admin", &task2, &result)
	fmt.Println("update:", result)

	var total []*taskProto.Task
	serv.ReadAll("task", &total)

	fmt.Println("update:", total)
}

func TestLocalStore_Delete(t *testing.T) {
	serv := NewLocalStore()
	task := &taskProto.Task{Name: "aaa", Id: "admin"}
	task2 := taskProto.Task{Name: "aaa", Id: "admin"}
	serv.Write("task", &task)

	serv.Delete("task", "admin")
	var result []*taskProto.Task
	serv.Read("task", "admin", &result)
	fmt.Println(result)

	serv.Write("task", &task2)
	serv.Delete("task", "admin")
	serv.Read("task", "admin", &result)
	fmt.Println(result)

}
