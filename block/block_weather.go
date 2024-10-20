package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type WeatherReport struct {
	City      string
	Temperature float64
	Humidity    float64
	Timestamp   time.Time
}

type WeatherBlock struct {
	Index        int
	PreviousHash string
	Timestamp    time.Time
	Reports      []WeatherReport
	Hash         string
}

func NewWeatherBlock(index int, previousHash string, reports []WeatherReport) *WeatherBlock {
	block := &WeatherBlock{
		Index:        index,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
		Reports:      reports,
	}
	block.Hash = block.CalculateHash()
	return block
}

func (b *WeatherBlock) CalculateHash() string {
	record := fmt.Sprintf("%d%s%s%v", b.Index, b.PreviousHash, b.Timestamp, b.Reports)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (b *WeatherBlock) PrintBlock() {
	fmt.Printf("Index: %d\n", b.Index)
	fmt.Printf("Previous Hash: %s\n", b.PreviousHash)
	fmt.Printf("Timestamp: %s\n", b.Timestamp)
	fmt.Println("Weather Reports:")
	for _, report := range b.Reports {
		fmt.Printf("  City: %s, Temperature: %.2f, Humidity: %.2f, Timestamp: %s\n",
			report.City, report.Temperature, report.Humidity, report.Timestamp)
	}
	fmt.Printf("Hash: %s\n", b.Hash)
}

