package block

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "time"
)

// Block yapısı
type Block struct {
    Index        int
    Timestamp    string
    Transactions []Transaction
    PrevHash     string
    Hash         string
    Nonce        int
}

// Transaction yapısı
type Transaction struct {
    From      string
    To        string
    Amount    int
    Type      string
    Signature string
}

// Blockchain yapısı
type Blockchain struct {
    Blocks []*Block
}

// CalculateHash belirli bir bloğun hash değerini hesaplar
func CalculateHash(block *Block) string {
    record := fmt.Sprintf("%d%s%d", block.Index, block.Timestamp, block.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    return hex.EncodeToString(h.Sum(nil))
}

// NewBlock yeni bir blok oluşturur
func NewBlock(index int, prevHash string, transactions []Transaction, nonce int) *Block {
    block := &Block{
        Index:        index,
        Timestamp:    time.Now().String(),
        PrevHash:     prevHash,
        Nonce:        nonce,
        Transactions: transactions,
    }
    block.Hash = CalculateHash(block) // Bloğun hash'ini hesapla
    return block
}

