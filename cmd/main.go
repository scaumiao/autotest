package main

import (
	"io/ioutil"

	autotest "github.com/scaumiao/autotest"
)

func main() {

	autotest.NewServer()

}

func getTestFile(path string) string {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(bs)
}
