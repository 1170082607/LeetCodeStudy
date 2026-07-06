package array_and_strings

import (
	"math/rand"
	"sort"
)

type raftState int

const (
	stateFollower raftState = iota
	stateCandidate
	stateLeader
)

const (
	minElectionTimeout = 4
	heartbeatTicks     = 2
)

// LogEntry represents a single replicated command and the term it was created in.
type LogEntry struct {
	Term    int
	Command string
}

// MessageType models the subset of Raft RPCs required for a teaching demo.
type MessageType int

const (
	MsgHup MessageType = iota
	MsgRequestVote
	MsgRequestVoteResp
	MsgAppend
	MsgAppendResp
	MsgPropose
)

// Message models communication between nodes in the in-memory Raft simulation.
type Message struct {
	From        int
	To          int
	Type        MessageType
	Term        int
	LogIndex    int
	LogTerm     int
	Entries     []LogEntry
	CommitIndex int
	VoteGranted bool
	Reject      bool
}

// RaftNode is a simplified Raft replica suitable for deterministic unit tests.
type RaftNode struct {
	id          int
	peers       []int
	state       raftState
	currentTerm int
	votedFor    int
	log         []LogEntry
	commitIndex int
	lastApplied int

	electionTimeout  int
	electionElapsed  int
	heartbeatTimeout int
	heartbeatElapsed int

	votes map[int]bool
	match map[int]int
	next  map[int]int

	rnd *rand.Rand
}

// NewRaftNode constructs a Raft replica with deterministic timing for tests.
func NewRaftNode(id int, peers []int, seed int64) *RaftNode {
	r := &RaftNode{
		id:               id,
		peers:            append([]int{}, peers...),
		state:            stateFollower,
		votedFor:         -1,
		log:              []LogEntry{{Term: 0}}, // log index starts at 1
		heartbeatTimeout: heartbeatTicks,
		votes:            make(map[int]bool),
		match:            make(map[int]int),
		next:             make(map[int]int),
		rnd:              rand.New(rand.NewSource(seed)),
	}
	r.resetElectionTimeout()
	return r
}

// Tick advances the replica's timers and triggers elections or heartbeats when needed.
func (r *RaftNode) Tick() []Message {
	switch r.state {
	case stateLeader:
		r.heartbeatElapsed++
		if r.heartbeatElapsed >= r.heartbeatTimeout {
			r.heartbeatElapsed = 0
			return r.broadcastHeartbeat()
		}
	default:
		r.electionElapsed++
		if r.electionElapsed >= r.electionTimeout {
			r.electionElapsed = 0
			return r.campaign()
		}
	}
	return nil
}

// Step processes inbound messages and returns outbound messages to forward.
func (r *RaftNode) Step(msg Message) []Message {
	if msg.Term > r.currentTerm {
		r.becomeFollower(msg.Term)
	}

	switch msg.Type {
	case MsgHup:
		if r.state != stateLeader {
			return r.campaign()
		}
	case MsgRequestVote:
		return r.handleRequestVote(msg)
	case MsgRequestVoteResp:
		if r.state == stateCandidate {
			return r.handleVoteResponse(msg)
		}
	case MsgAppend:
		return r.handleAppendEntries(msg)
	case MsgAppendResp:
		if r.state == stateLeader {
			r.handleAppendResponse(msg)
		}
	case MsgPropose:
		if r.state == stateLeader {
			return r.handleProposal(msg)
		}
	}
	return nil
}

func (r *RaftNode) campaign() []Message {
	r.currentTerm++
	r.becomeCandidate()
	msgs := make([]Message, 0, len(r.peers))
	for _, id := range r.peers {
		if id == r.id {
			continue
		}
		msgs = append(msgs, Message{
			From:     r.id,
			To:       id,
			Type:     MsgRequestVote,
			Term:     r.currentTerm,
			LogIndex: r.lastIndex(),
			LogTerm:  r.lastTerm(),
		})
	}
	return msgs
}

func (r *RaftNode) handleRequestVote(msg Message) []Message {
	voteGranted := false
	if msg.Term < r.currentTerm {
		voteGranted = false
	} else {
		upToDate := msg.LogTerm > r.lastTerm() || (msg.LogTerm == r.lastTerm() && msg.LogIndex >= r.lastIndex())
		if (r.votedFor == -1 || r.votedFor == msg.From) && upToDate {
			voteGranted = true
			r.votedFor = msg.From
			r.resetElectionTimeout()
		}
	}
	return []Message{{
		From:        r.id,
		To:          msg.From,
		Type:        MsgRequestVoteResp,
		Term:        r.currentTerm,
		VoteGranted: voteGranted,
	}}
}

func (r *RaftNode) handleVoteResponse(msg Message) []Message {
	r.votes[msg.From] = msg.VoteGranted
	if msg.VoteGranted && r.hasMajorityVotes() {
		r.becomeLeader()
		return r.broadcastHeartbeat()
	}
	denied := 0
	for _, granted := range r.votes {
		if !granted {
			denied++
		}
	}
	if denied >= r.quorum() {
		r.becomeFollower(r.currentTerm)
	}
	return nil
}

func (r *RaftNode) handleAppendEntries(msg Message) []Message {
	if msg.Term < r.currentTerm {
		return []Message{{
			From:     r.id,
			To:       msg.From,
			Type:     MsgAppendResp,
			Term:     r.currentTerm,
			LogIndex: r.lastIndex(),
			Reject:   true,
		}}
	}

	r.becomeFollower(msg.Term)

	if msg.LogIndex > r.lastIndex() {
		return []Message{{
			From:     r.id,
			To:       msg.From,
			Type:     MsgAppendResp,
			Term:     r.currentTerm,
			LogIndex: r.lastIndex(),
			Reject:   true,
		}}
	}

	if msg.LogIndex >= 0 && msg.LogIndex < len(r.log) && r.log[msg.LogIndex].Term != msg.LogTerm {
		return []Message{{
			From:     r.id,
			To:       msg.From,
			Type:     MsgAppendResp,
			Term:     r.currentTerm,
			LogIndex: r.lastIndex(),
			Reject:   true,
		}}
	}

	for i, entry := range msg.Entries {
		index := msg.LogIndex + 1 + i
		if index <= r.lastIndex() {
			if r.log[index].Term != entry.Term {
				r.log = r.log[:index]
				r.log = append(r.log, entry)
			}
			continue
		}
		r.log = append(r.log, entry)
	}

	if msg.CommitIndex > r.commitIndex {
		r.commitIndex = minInt(msg.CommitIndex, r.lastIndex())
	}

	return []Message{{
		From:     r.id,
		To:       msg.From,
		Type:     MsgAppendResp,
		Term:     r.currentTerm,
		LogIndex: r.lastIndex(),
		Reject:   false,
	}}
}

func (r *RaftNode) handleAppendResponse(msg Message) {
	if msg.Reject {
		r.next[msg.From] = maxInt(1, r.next[msg.From]-1)
		return
	}
	r.match[msg.From] = msg.LogIndex
	r.next[msg.From] = msg.LogIndex + 1

	matches := make([]int, 0, len(r.match)+1)
	matches = append(matches, r.lastIndex())
	for _, m := range r.match {
		matches = append(matches, m)
	}
	sort.Ints(matches)
	idx := matches[len(matches)/2]
	if idx > r.commitIndex && r.log[idx].Term == r.currentTerm {
		r.commitIndex = idx
	}
}

func (r *RaftNode) handleProposal(msg Message) []Message {
	if len(msg.Entries) == 0 {
		return nil
	}
	r.log = append(r.log, LogEntry{Term: r.currentTerm, Command: msg.Entries[0].Command})
	r.match[r.id] = r.lastIndex()
	r.next[r.id] = r.lastIndex() + 1
	return r.broadcastAppendEntries(r.commitIndex)
}

func (r *RaftNode) broadcastHeartbeat() []Message {
	return r.broadcastAppendEntries(r.commitIndex)
}

func (r *RaftNode) broadcastAppendEntries(commit int) []Message {
	msgs := make([]Message, 0, len(r.peers))
	for _, id := range r.peers {
		if id == r.id {
			continue
		}
		nextIndex := r.next[id]
		if nextIndex == 0 {
			nextIndex = r.lastIndex() + 1
		}
		prevIndex := nextIndex - 1
		if prevIndex < 0 {
			prevIndex = 0
		}
		prevTerm := r.log[prevIndex].Term
		entries := make([]LogEntry, 0)
		if nextIndex <= r.lastIndex() {
			entries = append(entries, r.log[nextIndex:]...)
		}
		msgs = append(msgs, Message{
			From:        r.id,
			To:          id,
			Type:        MsgAppend,
			Term:        r.currentTerm,
			LogIndex:    prevIndex,
			LogTerm:     prevTerm,
			Entries:     append([]LogEntry(nil), entries...),
			CommitIndex: commit,
		})
	}
	r.heartbeatElapsed = 0
	return msgs
}

func (r *RaftNode) becomeFollower(term int) {
	r.state = stateFollower
	r.currentTerm = term
	r.votedFor = -1
	r.votes = make(map[int]bool)
	r.resetElectionTimeout()
}

func (r *RaftNode) becomeCandidate() {
	r.state = stateCandidate
	r.votedFor = r.id
	r.votes = map[int]bool{r.id: true}
	r.resetElectionTimeout()
}

func (r *RaftNode) becomeLeader() {
	r.state = stateLeader
	r.votes = make(map[int]bool)
	r.match = make(map[int]int)
	r.next = make(map[int]int)
	for _, id := range r.peers {
		if id == r.id {
			continue
		}
		r.next[id] = r.lastIndex() + 1
		r.match[id] = 0
	}
	r.match[r.id] = r.lastIndex()
	r.next[r.id] = r.lastIndex() + 1
	r.heartbeatElapsed = 0
	r.electionElapsed = 0
}

func (r *RaftNode) resetElectionTimeout() {
	r.electionTimeout = minElectionTimeout + r.rnd.Intn(minElectionTimeout)
	r.electionElapsed = 0
}

func (r *RaftNode) hasMajorityVotes() bool {
	granted := 0
	for _, ok := range r.votes {
		if ok {
			granted++
		}
	}
	return granted >= r.quorum()
}

func (r *RaftNode) quorum() int {
	return len(r.peers)/2 + 1
}

func (r *RaftNode) lastIndex() int {
	return len(r.log) - 1
}

func (r *RaftNode) lastTerm() int {
	if len(r.log) == 0 {
		return 0
	}
	return r.log[len(r.log)-1].Term
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
