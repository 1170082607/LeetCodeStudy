package array_and_strings

import "testing"

func TestRaftElection(t *testing.T) {
	nodes := newRaftCluster()
	queue := make([]Message, 0)
	leaderID := 0
	for round := 0; round < 20 && leaderID == 0; round++ {
		for _, node := range nodes {
			if out := node.Tick(); len(out) > 0 {
				queue = append(queue, out...)
			}
		}
		queue = driveMessages(queue, nodes)
		leaderID = singleLeader(nodes)
	}

	if leaderID == 0 {
		t.Fatalf("expected a leader to be elected")
	}

	for id, node := range nodes {
		if id == leaderID {
			if node.state != stateLeader {
				t.Fatalf("node %d should be leader, state=%v", id, node.state)
			}
		} else if node.state == stateLeader {
			t.Fatalf("multiple leaders detected: %d and %d", leaderID, id)
		}
	}
}

func TestRaftReplication(t *testing.T) {
	nodes := newRaftCluster()
	queue := make([]Message, 0)
	leaderID := 0
	for round := 0; round < 20 && leaderID == 0; round++ {
		for _, node := range nodes {
			if out := node.Tick(); len(out) > 0 {
				queue = append(queue, out...)
			}
		}
		queue = driveMessages(queue, nodes)
		leaderID = singleLeader(nodes)
	}
	if leaderID == 0 {
		t.Fatalf("leader not elected")
	}

	leader := nodes[leaderID]
	proposal := Message{From: leaderID, To: leaderID, Type: MsgPropose, Entries: []LogEntry{{Command: "set x=1"}}}
	queue = append(queue, leader.Step(proposal)...)
	queue = driveMessages(queue, nodes)

	for round := 0; round < 5; round++ {
		for _, node := range nodes {
			if out := node.Tick(); len(out) > 0 {
				queue = append(queue, out...)
			}
		}
		queue = driveMessages(queue, nodes)
	}

	if leader.commitIndex != 1 {
		t.Fatalf("expected leader commit index 1, got %d", leader.commitIndex)
	}

	for id, node := range nodes {
		if node.commitIndex != leader.commitIndex {
			t.Fatalf("node %d commit index %d != leader %d", id, node.commitIndex, leader.commitIndex)
		}
		if node.commitIndex == 0 {
			t.Fatalf("node %d did not commit the entry", id)
		}
		got := node.log[node.commitIndex].Command
		if got != "set x=1" {
			t.Fatalf("node %d committed command %q", id, got)
		}
	}
}

func newRaftCluster() map[int]*RaftNode {
	peers := []int{1, 2, 3}
	seeds := []int64{1, 2, 3}
	nodes := make(map[int]*RaftNode, len(peers))
	for i, id := range peers {
		nodes[id] = NewRaftNode(id, peers, seeds[i])
	}
	return nodes
}

func driveMessages(queue []Message, nodes map[int]*RaftNode) []Message {
	for len(queue) > 0 {
		msg := queue[0]
		queue = queue[1:]
		target, ok := nodes[msg.To]
		if !ok {
			continue
		}
		out := target.Step(msg)
		if len(out) > 0 {
			queue = append(queue, out...)
		}
	}
	return queue
}

func singleLeader(nodes map[int]*RaftNode) int {
	leader := 0
	for id, node := range nodes {
		if node.state == stateLeader {
			if leader != 0 {
				return 0
			}
			leader = id
		}
	}
	return leader
}
