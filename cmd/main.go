package main

import (
	"flag"
	"io/ioutil"

	autotest "github.com/scaumiao/autotest"
)

var grpcPort = flag.String("grpc_port", "50051", "grpc port,default 50051")
var gwPort = flag.String("gw_port", "8080", "grpc gateway port,default 8080")

func main() {
	flag.Parse()
	autotest.NewServer(*grpcPort, *gwPort)

}

func getTestFile(path string) string {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(bs)
}
