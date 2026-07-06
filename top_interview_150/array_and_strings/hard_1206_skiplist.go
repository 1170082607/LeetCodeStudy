package array_and_strings

import (
	"math/rand"
	"time"
)

const (
	skiplistMaxLevel    = 16
	skiplistProbability = 0.25
)

type skipNode struct {
	val  int
	next []*skipNode
}

// Skiplist implements randomized search, insert, and delete in logarithmic time.
type Skiplist struct {
	head  *skipNode
	level int
	rnd   *rand.Rand
}

// NewSkiplist constructs an empty skip list seeded with time-based randomness.
func NewSkiplist() *Skiplist {
	head := &skipNode{
		val:  -1,
		next: make([]*skipNode, skiplistMaxLevel),
	}
	return &Skiplist{
		head:  head,
		level: 1,
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func newSkipNode(val, level int) *skipNode {
	return &skipNode{
		val:  val,
		next: make([]*skipNode, level),
	}
}

func (s *Skiplist) randomLevel() int {
	lvl := 1
	for lvl < skiplistMaxLevel && s.rnd.Float64() < skiplistProbability {
		lvl++
	}
	return lvl
}

// Search reports whether the target value exists in the skip list.
func (s *Skiplist) Search(target int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < target {
			curr = curr.next[i]
		}
	}
	curr = curr.next[0]
	return curr != nil && curr.val == target
}

// Add inserts the given number into the skip list.
func (s *Skiplist) Add(num int) {
	update := make([]*skipNode, skiplistMaxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
		update[i] = curr
	}
	lvl := s.randomLevel()
	if lvl > s.level {
		for i := s.level; i < lvl; i++ {
			update[i] = s.head
		}
		s.level = lvl
	}
	node := newSkipNode(num, lvl)
	for i := 0; i < lvl; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}
}

// Erase removes the given number if present and reports success.
func (s *Skiplist) Erase(num int) bool {
	update := make([]*skipNode, skiplistMaxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
		update[i] = curr
	}
	target := curr.next[0]
	if target == nil || target.val != num {
		return false
	}
	for i := 0; i < s.level; i++ {
		if update[i].next[i] != target {
			continue
		}
		update[i].next[i] = target.next[i]
	}
	for s.level > 1 && s.head.next[s.level-1] == nil {
		s.level--
	}
	return true
}
