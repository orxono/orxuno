package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)

func TestNewVotingBlock(t *testing.T) {
	votes := []Vote{
		{VoterID: "user1", Candidate: "Alice", Timestamp: time.Now()},
		{VoterID: "user2", Candidate: "Bob", Timestamp: time.Now()},
	}
	block := NewVotingBlock(1, "0000000000000000", votes)

	if block.Index != 1 {
		t.Errorf("Expected index 1, got %d", block.Index)
	}
	if block.PreviousHash != "0000000000000000" {
		t.Errorf("Expected previous hash '0000000000000000', got %s", block.PreviousHash)
	}
	if len(block.Votes) != 2 {
		t.Errorf("Expected 2 votes, got %d", len(block.Votes))
	}
}

func TestCalculateVotingHash(t *testing.T) {
	votes := []Vote{
		{VoterID: "user1", Candidate: "Alice", Timestamp: time.Now()},
	}
	block := NewVotingBlock(1, "0000000000000000", votes)
	expectedHash := block.CalculateHash()

	if block.Hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, block.Hash)
	}
}

func TestPrintVotingBlock(t *testing.T) {
	votes := []Vote{
		{VoterID: "user1", Candidate: "Alice", Timestamp: time.Now()},
		{VoterID: "user2", Candidate: "Bob", Timestamp: time.Now()},
	}
	block := NewVotingBlock(1, "0000000000000000", votes)

	block.PrintBlock() // Burada çıktıyı görsel olarak inceleyebilirsin.
}

