package mgrd_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	. "gopkg.in/check.v1"

	"github.com/mvo5/qrest-skeleton/backend"
	"github.com/mvo5/qrest-skeleton/mgrd"
)

type apiSuite struct {
	d     *mgrd.Mgrd
	quotaMgr *backend.QuotaManager
}

var _ = Suite(&apiSuite{})

func (s *apiSuite) SetUpTest(c *C) {
	s.quotaMgr = backend.NewQuotaManager()
	s.d = mgrd.New("", s.quotaMgr)
}

func (s *apiSuite) doReq(method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	rr := httptest.NewRecorder()
	s.d.Server().Handler.ServeHTTP(rr, req)
	return rr
}

func (s *apiSuite) TestSlashRoot(c *C) {
	rr := s.doReq("GET", "/", nil)
	c.Check(rr.Code, Equals, 200)
	c.Check(rr.Body.String(), Equals, "rest API for quota service\n")
}

func (s *apiSuite) TestQuotaCreateHappy(c *C) {
	// ...
}

