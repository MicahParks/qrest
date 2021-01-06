package backend

import (
	"fmt"
	"sort"
	"sync"
)

// QuotaManager manages the quota groups for snaps.
type QuotaManager struct {
	mu sync.Mutex

	groups map[string]*QuotaGroup
}

func NewQuotaManager() *QuotaManager {
	return &QuotaManager{
		groups: make(map[string]*QuotaGroup),
	}
}

func (q *QuotaManager) GetGroup(name string) *QuotaGroup {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.groups[name]
}

func (q *QuotaManager) Groups() []*QuotaGroup {
	q.mu.Lock()
	defer q.mu.Unlock()

	l := make([]*QuotaGroup, 0, len(q.groups))
	for _, v := range q.groups {
		l = append(l, v)
	}
	sort.Sort(byGroupName(l))
	return l
}

func (q *QuotaManager) AddGroup(name string, maxMemory uint64) *QuotaGroup {
	q.mu.Lock()
	defer q.mu.Unlock()

	// XXX: validate name
	// XXX2: validate that group name does not exist already?
	q.groups[name] = newQuotaGroup(q, name, maxMemory)
	return q.groups[name]
}

func (q *QuotaManager) RemoveGroup(name string) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if _, ok := q.groups[name]; !ok {
		return fmt.Errorf("cannot remove %q: does not exist", name)
	}
	delete(q.groups, name)
	return nil
}

type byGroupName []*QuotaGroup

func (a byGroupName) Len() int           { return len(a) }
func (a byGroupName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byGroupName) Less(i, j int) bool { return a[i].Name() < a[j].Name() }

// QuotaGroup represents a single quota group
type QuotaGroup struct {
	mu sync.Mutex

	quota *QuotaManager

	name      string
	maxMemory uint64

	snaps  map[string]bool
	groups map[string]bool
}

func newQuotaGroup(q *QuotaManager, name string, maxMemory uint64) *QuotaGroup {
	// XXX: validate name
	return &QuotaGroup{
		quota:     q,
		name:      name,
		maxMemory: maxMemory,
		snaps:     make(map[string]bool),
		groups:    make(map[string]bool),
	}
}

func (q *QuotaGroup) Name() string {
	return q.name
}

func (q *QuotaGroup) MaxMemory() uint64 {
	return q.maxMemory
}

func (q *QuotaGroup) SetMaxMemory(new uint64) error {
	// XXX: This is cheating, this ensures we never overflow the allocated
	// boxes. Instead this should re-check against all users of the
	// quota group.
	if new > q.maxMemory {
		return fmt.Errorf("cannot increase max-memory yet")
	}
	q.maxMemory = new
	return nil
}

func (q *QuotaGroup) AddSnap(snap string) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.snaps[snap] = true
}

func (q *QuotaGroup) RemoveSnap(snap string) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if _, ok := q.snaps[snap]; !ok {
		return fmt.Errorf("cannot remove snap %q: does not exist", snap)
	}
	delete(q.snaps, snap)
	return nil
}

func (q *QuotaGroup) Snaps() []string {
	q.mu.Lock()
	defer q.mu.Unlock()

	snaps := make([]string, 0, len(q.snaps))
	for k := range q.snaps {
		snaps = append(snaps, k)
	}
	sort.Strings(snaps)
	return snaps
}

func (q *QuotaGroup) AddGroup(grpName string) error {
	// XXX: validate name

	if _, ok := q.quota.groups[grpName]; !ok {
		return fmt.Errorf("cannot add %q: does not exist", grpName)
	}
	// ensure the group has enough "space"
	maxMem := q.quota.GetGroup(grpName).maxMemory
	for grpName := range q.groups {
		grp := q.quota.GetGroup(grpName)
		maxMem += grp.maxMemory
	}
	if maxMem > q.maxMemory {
		return fmt.Errorf("cannot add %q: does not fit", grpName)
	}
	q.groups[grpName] = true
	return nil
}

func (q *QuotaGroup) RemoveGroup(grpName string) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if _, ok := q.groups[grpName]; !ok {
		return fmt.Errorf("cannot remove group %q: does not exist", grpName)
	}
	delete(q.groups, grpName)
	return nil
}

func (q *QuotaGroup) Groups() []string {
	q.mu.Lock()
	defer q.mu.Unlock()

	groups := make([]string, 0, len(q.groups))
	for k := range q.groups {
		groups = append(groups, k)
	}
	sort.Strings(groups)
	return groups
}
