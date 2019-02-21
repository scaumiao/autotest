package api

import (
	server "github.com/scaumiao/autotest"
)

// server.go

type API struct {
	TestServer *server.Server
}

func NewApi() *API {
	api := &API{}
	return api
}

func (api *API) SetTestServer(s *server.Server) {
	api.TestServer = s
}
