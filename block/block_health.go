package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type HealthRecord struct {
	PatientID string
	Temperature float64
	BloodPressure string
	PulseRate int
	Timestamp time.Time
}

type HealthBlock struct {
	Index        int
	PreviousHash string
	Timestamp    time.Time
	Records      []HealthRecord
	Hash         string
}

func NewHealthBlock(index int, previousHash string, records []HealthRecord) *HealthBlock {
	block := &HealthBlock{
		Index:        index,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
		Records:      records,
	}
	block.Hash = block.CalculateHash()
	return block
}

func (b *HealthBlock) CalculateHash() string {
	record := fmt.Sprintf("%d%s%s%v", b.Index, b.PreviousHash, b.Timestamp, b.Records)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (b *HealthBlock) PrintBlock() {
	fmt.Printf("Index: %d\n", b.Index)
	fmt.Printf("Previous Hash: %s\n", b.PreviousHash)
	fmt.Printf("Timestamp: %s\n", b.Timestamp)
	fmt.Println("Health Records:")
	for _, record := range b.Records {
		fmt.Printf("  PatientID: %s, Temperature: %.2f, Blood Pressure: %s, Pulse Rate: %d, Timestamp: %s\n",
			record.PatientID, record.Temperature, record.BloodPressure, record.PulseRate, record.Timestamp)
	}
	fmt.Printf("Hash: %s\n", b.Hash)
}

