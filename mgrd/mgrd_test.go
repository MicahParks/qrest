package mgrd_test

import (
	"fmt"
	"net/http"
	"testing"

	. "gopkg.in/check.v1"

	"github.com/mvo5/qrest-skeleton/mgrd"
)

func Test(t *testing.T) { TestingT(t) }

type mgrdSuite struct {
	addr string
	d    *mgrd.Mgrd
}

var _ = Suite(&mgrdSuite{})

func (s *mgrdSuite) SetUpTest(c *C) {
	s.addr = "localhost:18080"
	s.d = mgrd.New(s.addr, nil)
	c.Assert(s.d, NotNil)
}

func (s *mgrdSuite) TestStartStop(c *C) {
	err := s.d.Start()
	c.Assert(err, IsNil)

	resp, err := http.Get(fmt.Sprintf("http://%s/", s.addr))
	c.Assert(err, IsNil)
	c.Check(resp.StatusCode, Equals, 200)

	s.d.Stop()
	err = s.d.Wait()
	c.Assert(err, IsNil)
}

func (s *mgrdSuite) TestStopErr(c *C) {
	err := s.d.Start()
	c.Assert(err, IsNil)

	mgrd.InjectErr(s.d, fmt.Errorf("boom-1"))
	mgrd.InjectErr(s.d, fmt.Errorf("boom-2"))

	s.d.Stop()
	err = s.d.Wait()
	c.Assert(err, ErrorMatches, `mgrd errored: \[boom-1 boom-2\]`)
}
