package block

import (
    "fmt"
    "time"
)

// SosyalTransaction temsil eden yapı
type SosyalTransaction struct {
    From     string
    To       string
    Content  string
    Amount   float64
    Timestamp time.Time
}

// SosyalBlock sosyal işlemleri içeren blok yapısı
type SosyalBlock struct {
    Index        int
    PreviousHash string
    Transactions []SosyalTransaction
    Timestamp    time.Time
    Hash         string
}

// Yeni sosyal bloğu oluştur
func NewSosyalBlock(index int, previousHash string, transactions []SosyalTransaction) *SosyalBlock {
    return &SosyalBlock{
        Index:        index,
        PreviousHash: previousHash,
        Transactions: transactions,
        Timestamp:    time.Now(),
        Hash:         "", // Hash hesaplanacak
    }
}

// Sosyal veri setini oluştur
func CreateSampleSosyalTransactions() []SosyalTransaction {
    return []SosyalTransaction{
        {
            From:     "Alice",
            To:       "Bob",
            Content:  "Hello, Bob! How are you?",
            Amount:   0.0,
            Timestamp: time.Now(),
        },
        {
            From:     "Bob",
            To:       "Charlie",
            Content:  "Hi Alice! I'm good, thanks!",
            Amount:   0.0,
            Timestamp: time.Now(),
        },
        {
            From:     "Charlie",
            To:       "David",
            Content:  "Just sharing a funny meme!",
            Amount:   0.0,
            Timestamp: time.Now(),
        },
    }
}

// PrintSosyalBlock bloğun bilgilerini yazdır
func (b *SosyalBlock) PrintSosyalBlock() {
    fmt.Printf("Index: %d\n", b.Index)
    fmt.Printf("Previous Hash: %s\n", b.PreviousHash)
    fmt.Printf("Timestamp: %s\n", b.Timestamp)
    fmt.Println("Transactions:")
    for _, tx := range b.Transactions {
        fmt.Printf("  From: %s, To: %s, Content: %s, Amount: %.2f\n", tx.From, tx.To, tx.Content, tx.Amount)
    }
    fmt.Printf("Hash: %s\n", b.Hash)
}

