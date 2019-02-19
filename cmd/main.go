package main

import (
	"io/ioutil"

	autotest "github.com/scaumiao/autotest"
)

func main() {

	s := autotest.NewServer()
	defer s.Stop()
	raw := getTestFile("../test/wf.tl")
	s.RunTest(raw)

}

func getTestFile(path string) string {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(bs)
}
