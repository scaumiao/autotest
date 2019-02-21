package taskService

import (
	"github.com/scaumiao/autotest/app/api"
	taskProto "github.com/scaumiao/autotest/proto/task"
	"github.com/scaumiao/autotest/utils"
	"golang.org/x/net/context"
)

type Service struct {
	Api *api.API
}

// 创建脚本
func (s *Service) CreateTask(ctx context.Context, task *taskProto.CreateTaskRequest) (*taskProto.CreateTaskResponse, error) {
	// data := []*usr.UserInfo{}
	id := utils.NewID()
	data := &taskProto.CreateTaskResponse{
		Id: id,
	}
	// 自动生成ID data.Id = xxx
	// 调用store的write方法保存
	// 返回保存值
	return data, nil
}
