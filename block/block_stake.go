package block

import (
    "fmt"
)

// StakeTransaction bir stake işlemini temsil eder
type StakeTransaction struct {
    From    string  // Stake'i yapan kişinin adı
    Amount  float64 // Stake edilen miktar
    Status  string  // İşlem durumu (örneğin: "pending", "completed")
}

// StakeBlock bir stake bloğunu temsil eder
type StakeBlock struct {
    Index        int                 // Blokun indeksi
    Timestamp    string              // Zaman damgası
    Transactions []StakeTransaction   // Bu bloktaki stake işlemleri
    PrevHash     string              // Önceki blokun hash'i
    Hash         string              // Bu bloğun hash'i
}

// NewStakeBlock yeni bir StakeBlock döndürür
func NewStakeBlock(index int, prevHash string, transactions []StakeTransaction) *StakeBlock {
    return &StakeBlock{
        Index:        index,
        Timestamp:    fmt.Sprint("2024-10-20T13:14:35+03:00"), // Buraya gerçek zaman damgası eklenebilir
        Transactions: transactions,
        PrevHash:     prevHash,
        Hash:         "", // Hash daha sonra hesaplanacak
    }
}

// Örnek veri seti
var sampleStakeTransactions = []StakeTransaction{
    {From: "Alice", Amount: 100.0, Status: "completed"},
    {From: "Bob", Amount: 50.0, Status: "pending"},
    {From: "Charlie", Amount: 75.0, Status: "completed"},
}

// Örnek StakeBlock
var sampleStakeBlock = NewStakeBlock(1, "0000000000000000", sampleStakeTransactions)

// Fonksiyonlar
// StakeBlock'u yazdır
func (sb *StakeBlock) Print() {
    fmt.Printf("Index: %d\nTimestamp: %s\nPrevious Hash: %s\nHash: %s\n", sb.Index, sb.Timestamp, sb.PrevHash, sb.Hash)
    fmt.Println("Transactions:")
    for _, tx := range sb.Transactions {
        fmt.Printf("  From: %s, Amount: %.2f, Status: %s\n", tx.From, tx.Amount, tx.Status)
    }
}

