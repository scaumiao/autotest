package autotest

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/scaumiao/autotest/proto/stages"
	// "gitlab.iyorhe.com/wfgz/reverseproxy/proto/autotest/report"
)

//Report autotest 报告
type Report struct {
}

type Stage stages.Stage

func AddAction(s *stages.Stage, a stages.Action) {

	s.Actions = append(s.Actions, &a)
}

type Stages []*stages.Stage

// func (s *Stages) AddStage(sd Stage) {
// 	s = append(s, sd)
// }

type Action struct {
	*stages.Action
}

func NewStage(id, description string, line int32) *stages.Stage {

	return &stages.Stage{Id: id,
		Description: description,
		Line:        line,
		StartAt:     ptypes.TimestampNow(),
	}
}

func NewStages() Stages {
	return []*stages.Stage{}
}

func NewAction() stages.Action {
	return stages.Action{}
}
