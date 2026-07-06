package array_and_strings

import (
	"math/rand"
	"testing"
)

func TestSkiplistBasicOperations(t *testing.T) {
	sl := NewSkiplist()
	sl.rnd = rand.New(rand.NewSource(1))

	if sl.Search(3) {
		t.Fatalf("expected empty skiplist to return false on search")
	}

	sl.Add(1)
	sl.Add(2)
	sl.Add(3)

	if !sl.Search(1) || !sl.Search(2) || !sl.Search(3) {
		t.Fatalf("expected to find inserted elements")
	}

	if sl.Search(4) {
		t.Fatalf("did not expect to find absent element")
	}

	if !sl.Erase(2) {
		t.Fatalf("expected erase to succeed")
	}

	if sl.Search(2) {
		t.Fatalf("expected element to be removed")
	}

	if sl.Erase(2) {
		t.Fatalf("expected erase to fail when element missing")
	}
}

func TestSkiplistAllowsDuplicates(t *testing.T) {
	sl := NewSkiplist()
	sl.rnd = rand.New(rand.NewSource(2))

	sl.Add(5)
	sl.Add(5)

	if !sl.Search(5) {
		t.Fatalf("expected to find value after inserting duplicates")
	}

	if !sl.Erase(5) {
		t.Fatalf("expected first erase to succeed")
	}

	if !sl.Search(5) {
		t.Fatalf("expected value to remain after removing one duplicate")
	}

	if !sl.Erase(5) {
		t.Fatalf("expected second erase to remove remaining duplicate")
	}

	if sl.Search(5) {
		t.Fatalf("expected value to be fully removed")
	}
}
