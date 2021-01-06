package backend_test

import (
	"fmt"
	"sync"
	"testing"

	. "gopkg.in/check.v1"

	"github.com/mvo5/qrest-skeleton/backend"
)

func Test(t *testing.T) { TestingT(t) }

type backendSuite struct {
	quotaMgr *backend.QuotaManager
}

var _ = Suite(&backendSuite{})

func (s *backendSuite) SetUpTest(c *C) {
	s.quotaMgr = backend.NewQuotaManager()
	c.Assert(s.quotaMgr.Groups(), HasLen, 0)
}

func (s *backendSuite) TestQuotaGroupSimple(c *C) {
	qg := s.quotaMgr.AddGroup("foo", 16*1024)
	c.Check(qg.Name(), Equals, "foo")
	c.Check(qg.MaxMemory(), Equals, uint64(16*1024))
	c.Check(qg.Snaps(), HasLen, 0)
}

func (s *backendSuite) TestQuotaGroupAddRemoveSnaps(c *C) {
	qg := s.quotaMgr.AddGroup("foo", 16*1024)

	qg.AddSnap("snap1")
	c.Check(qg.Snaps(), DeepEquals, []string{"snap1"})
	qg.AddSnap("snap2")
	c.Check(qg.Snaps(), DeepEquals, []string{"snap1", "snap2"})

	err := qg.RemoveSnap("snap1")
	c.Assert(err, IsNil)
	c.Check(qg.Snaps(), DeepEquals, []string{"snap2"})
	err = qg.RemoveSnap("snap1")
	c.Check(err, ErrorMatches, `cannot remove snap "snap1": does not exist`)

	err = qg.RemoveSnap("snap2")
	c.Assert(err, IsNil)
	c.Check(qg.Snaps(), HasLen, 0)
}

func (s *backendSuite) TestQuotaGroupAddGroupUnhappy(c *C) {
	qg := s.quotaMgr.AddGroup("grp1", 16*1024)
	err := qg.AddGroup("other-group")
	c.Assert(err, ErrorMatches, `cannot add "other-group": does not exist`)
}

func (s *backendSuite) TestQuotaGroupAddRemoveGroups(c *C) {
	s.quotaMgr.AddGroup("group1", 16*1024)
	s.quotaMgr.AddGroup("group2", 16*1024)

	qg := s.quotaMgr.AddGroup("qg", 32*1024)
	err := qg.AddGroup("group1")
	c.Assert(err, IsNil)
	c.Check(qg.Groups(), DeepEquals, []string{"group1"})
	err = qg.AddGroup("group2")
	c.Assert(err, IsNil)
	c.Check(qg.Groups(), DeepEquals, []string{"group1", "group2"})

	err = qg.RemoveGroup("group1")
	c.Assert(err, IsNil)
	c.Check(qg.Groups(), DeepEquals, []string{"group2"})
	err = qg.RemoveGroup("group1")
	c.Check(err, ErrorMatches, `cannot remove group "group1": does not exist`)

	err = qg.RemoveGroup("group2")
	c.Assert(err, IsNil)
	c.Check(qg.Groups(), HasLen, 0)
}

func (s *backendSuite) TestQuotaGroupAddChecksGroupConstraints(c *C) {
	grp32 := s.quotaMgr.AddGroup("grp32", 32*1024)
	grp16 := s.quotaMgr.AddGroup("grp16", 16*1024)

	err := grp32.AddGroup("grp16")
	c.Assert(err, IsNil)
	err = grp16.AddGroup("grp32")
	c.Assert(err, ErrorMatches, `cannot add "grp32": does not fit`)
}

func (s *backendSuite) TestAddRemoveGroup(c *C) {
	qg := s.quotaMgr.AddGroup("foo", 16*1024)
	c.Check(s.quotaMgr.Groups(), HasLen, 1)
	c.Check(qg, DeepEquals, s.quotaMgr.GetGroup("foo"))

	err := s.quotaMgr.RemoveGroup("foo")
	c.Assert(err, IsNil)
	c.Check(s.quotaMgr.Groups(), HasLen, 0)
}

func (s *backendSuite) TestAddRemoveGroupUnhappy(c *C) {
	err := s.quotaMgr.RemoveGroup("some-group")
	c.Assert(err, ErrorMatches, `cannot remove "some-group": does not exist`)
}

func (s *backendSuite) TestAddGroupAddRemoveThreadSafe(c *C) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			s.quotaMgr.AddGroup(fmt.Sprintf("foo-%v", i), 16*1024)
			s.quotaMgr.RemoveGroup(fmt.Sprintf("foo-%v", i))
			wg.Done()
		}(i)
	}
	wg.Wait()

	c.Check(s.quotaMgr.Groups(), HasLen, 0)
}

func (s *backendSuite) TestGroupAddRemoveThreadSafe(c *C) {
	var wg sync.WaitGroup
	grp := s.quotaMgr.AddGroup("grp", 1024)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			grp.AddSnap(fmt.Sprintf("snap-%v", i))
			grp.RemoveSnap(fmt.Sprintf("snap-%v", i))
			wg.Done()
		}(i)
	}
	wg.Wait()

	c.Check(grp.Snaps(), HasLen, 0)
}
