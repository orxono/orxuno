package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Vote struct {
	VoterID string
	Candidate string
	Timestamp time.Time
}

type VotingBlock struct {
	Index        int
	PreviousHash string
	Timestamp    time.Time
	Votes        []Vote
	Hash         string
}

func NewVotingBlock(index int, previousHash string, votes []Vote) *VotingBlock {
	block := &VotingBlock{
		Index:        index,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
		Votes:        votes,
	}
	block.Hash = block.CalculateHash()
	return block
}

func (b *VotingBlock) CalculateHash() string {
	record := fmt.Sprintf("%d%s%s%v", b.Index, b.PreviousHash, b.Timestamp, b.Votes)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (b *VotingBlock) PrintBlock() {
	fmt.Printf("Index: %d\n", b.Index)
	fmt.Printf("Previous Hash: %s\n", b.PreviousHash)
	fmt.Printf("Timestamp: %s\n", b.Timestamp)
	fmt.Println("Votes:")
	for _, vote := range b.Votes {
		fmt.Printf("  VoterID: %s, Candidate: %s, Timestamp: %s\n", vote.VoterID, vote.Candidate, vote.Timestamp)
	}
	fmt.Printf("Hash: %s\n", b.Hash)
}

