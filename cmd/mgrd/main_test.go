package main_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	. "gopkg.in/check.v1"

	main "github.com/mvo5/qrest-skeleton/cmd/mgrd"
)

func Test(t *testing.T) { TestingT(t) }

type mainSuite struct{}

var _ = Suite(&mainSuite{})

func (s *mainSuite) TestIntegration(c *C) {
	// start the restd
	errCh := make(chan error, 1)
	go func() {
		errCh <- main.Run()
	}()

	// wait until port is ready
	for i := 0; i < 100; i++ {
		_, err := http.Get("http://localhost:8080/")
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	// check that it serves API requests
	resp, err := http.Get("http://localhost:8080/")
	c.Assert(err, IsNil)
	c.Check(resp.StatusCode, Equals, 200)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	c.Assert(err, IsNil)
	c.Check(string(body), Equals, "rest API for quota service\n")

	// starting twice will yield an error
	err = main.Run()
	c.Assert(err, ErrorMatches, ".* bind: address already in use")

	// shuting down works
	p, err := os.FindProcess(os.Getpid())
	c.Assert(err, IsNil)
	err = p.Signal(os.Interrupt)
	c.Assert(err, IsNil)

	// check that shutdown worked without errors
	err = <-errCh
	c.Assert(err, IsNil)
}
