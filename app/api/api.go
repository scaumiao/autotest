package api

import (
	server "github.com/scaumiao/autotest"
	logger "github.com/scaumiao/autotest/app/log"
	"github.com/scaumiao/autotest/app/store"
)

// server.go

type API struct {
	TestServer *server.Server
	TaskStore  *store.TaskStore
	JobStore   *store.JobStore
	LogServer  *logger.LogServer
	Config     GrpcConfig
	// channel
}

type GrpcConfig struct {
	GrpcHost string
	GrpcPort string
}

func NewApi(grpcHost string, grpcPort string) *API {
	api := &API{
		Config: GrpcConfig{
			GrpcHost: grpcHost,
			GrpcPort: grpcPort,
		},
	}
	return api
}

func (api *API) SetTestServer(s *server.Server) {
	api.TestServer = s
}
func (api *API) SetTaskStore(s *store.TaskStore) {
	api.TaskStore = s
}
func (api *API) SetJobStore(s *store.JobStore) {
	api.JobStore = s
}
func (api *API) SetLogServer(s *logger.LogServer) {
	api.LogServer = s
}
