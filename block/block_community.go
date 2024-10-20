package block

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "time"
)

// CommunityTransaction topluluk işlemleri için yapı
type CommunityTransaction struct {
    From    string  // İşlemi gerçekleştiren kullanıcı
    To      string  // İşlemi alan kullanıcı
    Content string  // İçerik
    Amount  float64 // Miktar
}

// CommunityBlock topluluk blokları için yapı
type CommunityBlock struct {
    Index        int                    // Blok numarası
    Timestamp    time.Time              // Zaman damgası
    PrevHash     string                 // Önceki blok hash'i
    Hash         string                 // Bu blok hash'i
    Transactions []CommunityTransaction   // Bu bloktaki işlemler
}

// NewCommunityBlock yeni bir CommunityBlock döndürür
func NewCommunityBlock(index int, prevHash string, transactions []CommunityTransaction) *CommunityBlock {
    block := &CommunityBlock{
        Index:        index,
        Timestamp:    time.Now(),
        PrevHash:     prevHash,
        Transactions: transactions,
    }
    block.Hash = CalculateCommunityHash(block) // CalculateHash'ı buradan çağırın
    return block
}

// CalculateCommunityHash belirli bir CommunityBlock'un hash'ini hesaplar
func CalculateCommunityHash(block *CommunityBlock) string {
    record := fmt.Sprintf("%d%s%s", block.Index, block.Timestamp.String(), block.PrevHash)
    h := sha256.New()
    h.Write([]byte(record))
    return hex.EncodeToString(h.Sum(nil))
}

// PrintBlock bloğun içeriğini yazdırır
func (b *CommunityBlock) PrintBlock() {
    fmt.Printf("Index: %d\n", b.Index)
    fmt.Printf("Previous Hash: %s\n", b.PrevHash)
    fmt.Printf("Timestamp: %s\n", b.Timestamp)
    fmt.Println("Transactions:")
    for _, tx := range b.Transactions {
        fmt.Printf("  From: %s, To: %s, Content: %s, Amount: %.2f\n", tx.From, tx.To, tx.Content, tx.Amount)
    }
    fmt.Printf("Hash: %s\n", b.Hash)
}

