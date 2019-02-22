package main

import (
	"autotest"
	"io/ioutil"
)

func main() {

	autotest.NewServer()

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
