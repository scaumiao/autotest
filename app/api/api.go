package api

import (
	server "github.com/scaumiao/autotest"
	"github.com/scaumiao/autotest/app/store"
)

// server.go

type API struct {
	TestServer *server.Server
	TaskStore  *store.TaskStore
}

func NewApi() *API {
	api := &API{}
	return api
}

func (api *API) SetTestServer(s *server.Server) {
	api.TestServer = s
}
func (api *API) SetTaskStore(s *store.TaskStore) {
	api.TaskStore = s
}
