package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type SecurityEvent struct {
	User      string
	Action    string // (ör. "login", "logout", "failed_attempt")
	Timestamp string // Olay zaman damgası
	IPAddress string // Kullanıcının IP adresi
}

type SecurityBlock struct {
	Index        int
	PreviousHash string
	Timestamp    string
	Transactions  []SecurityEvent
	Hash         string
}

func NewSecurityBlock(index int, previousHash string, transactions []SecurityEvent) *SecurityBlock {
	timestamp := time.Now().Format(time.RFC3339)
	block := &SecurityBlock{
		Index:        index,
		PreviousHash: previousHash,
		Timestamp:    timestamp,
		Transactions:  transactions,
		Hash:         "",
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *SecurityBlock) calculateHash() string {
	data := fmt.Sprintf("%d%s%s%v", b.Index, b.PreviousHash, b.Timestamp, b.Transactions)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (b *SecurityBlock) PrintBlock() {
	fmt.Printf("Index: %d\n", b.Index)
	fmt.Printf("Previous Hash: %s\n", b.PreviousHash)
	fmt.Printf("Timestamp: %s\n", b.Timestamp)
	fmt.Println("Transactions:")
	for _, tx := range b.Transactions {
		fmt.Printf("  User: %s, Action: %s, Timestamp: %s, IP: %s\n", tx.User, tx.Action, tx.Timestamp, tx.IPAddress)
	}
	fmt.Printf("Hash: %s\n", b.Hash)
}

