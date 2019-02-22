package main

import (
<<<<<<< HEAD
	"autotest"
	"io/ioutil"
=======
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	autotest "github.com/scaumiao/autotest"
	"github.com/scaumiao/autotest/app/api"
	service "github.com/scaumiao/autotest/app/service"
	"github.com/scaumiao/autotest/app/store"
	"github.com/scaumiao/autotest/app/store/local"
	taskProto "github.com/scaumiao/autotest/proto/task"
	"google.golang.org/grpc"
>>>>>>> develop
)

type Server struct {
}

var grpcPort = flag.String("grpc_port", "50051", "grpc port,default 50051")
var gwPort = flag.String("gw_port", "8080", "grpc gateway port,default 8080")


func main() {
	flag.Parse()
	autotestServ := autotest.NewServer()
	localStore := local.NewLocalStore()
	taskStore := store.NewTaskStore()
	taskStore.SetStore(localStore)
	apiServ := api.NewApi()
	apiServ.SetTestServer(autotestServ)
	apiServ.SetTaskStore(taskStore)

	ctx := context.Background()
	mux, err := newGateway(ctx, ":"+*grpcPort)
	if err != nil {
		log.Fatal("gw failed to listen", err)
	}
	http.Handle("/", mux)
	fmt.Printf("grpc gateway start at port:%s\n", *gwPort)
	go http.ListenAndServe(":"+*gwPort, nil)
	// grpc
	_err := grpcServerStart(*grpcPort, apiServ)
	if _err != nil {
		log.Fatal("failed to listen: ", err)
	}

	// log.SetLevel("Debug")
	// log.Debug("{'request_id':123,'user_ip':10.18.105.2}", "zenmeban")
	// log.Info("{'request_id':123,'user_ip':10.18.105.2}", "zenmeban")
	// log.Warn("{'request_id':123,'user_ip':10.18.105.2}", "zenmeban")
}

func getTestFile(path string) string {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(bs)
}

func grpcServerStart(grpcPort string, apiServ *api.API) error {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatal("failed to listen: ", err)
		return err
	}
	s := grpc.NewServer()

	_taskService := &service.Service{
		Api: apiServ,
	}

	taskProto.RegisterTaskServiceServer(s, _taskService)
	fmt.Printf("grpc service start at port:%s", grpcPort)
	return s.Serve(lis)
}

func newGateway(ctx context.Context, Endpoint string, opts ...gwruntime.ServeMuxOption) (http.Handler, error) {
	mux := gwruntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := taskProto.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, Endpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
