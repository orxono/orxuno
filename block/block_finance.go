package block

import (
    "crypto/sha256"
    "fmt"
    "time"
)

type FinanceTransaction struct {
    From   string
    To     string
    Amount float64
    Reason string
}

type FinanceBlock struct {
    Index        int
    PrevHash     string
    Timestamp    time.Time
    Transactions []FinanceTransaction
    Hash         string
}

func NewFinanceBlock(index int, prevHash string, transactions []FinanceTransaction) *FinanceBlock {
    block := &FinanceBlock{
        Index:        index,
        PrevHash:     prevHash,
        Timestamp:    time.Now(),
        Transactions: transactions,
    }
    block.Hash = CalculateFinanceHash(block)
    return block
}

func CalculateFinanceHash(block *FinanceBlock) string {
    var transactions string
    for _, tx := range block.Transactions {
        transactions += fmt.Sprintf("%s->%s:%.2f (%s),", tx.From, tx.To, tx.Amount, tx.Reason)
    }

    record := fmt.Sprintf("%d%s%s%s", block.Index, block.PrevHash, block.Timestamp.String(), transactions)
    h := sha256.New()
    h.Write([]byte(record))
    return fmt.Sprintf("%x", h.Sum(nil))
}

func (b *FinanceBlock) PrintBlock() {
    fmt.Printf("Index: %d\n", b.Index)
    fmt.Printf("Previous Hash: %s\n", b.PrevHash)
    fmt.Printf("Timestamp: %s\n", b.Timestamp)
    fmt.Println("Transactions:")
    for _, tx := range b.Transactions {
        fmt.Printf("  From: %s, To: %s, Amount: %.2f, Reason: %s\n", tx.From, tx.To, tx.Amount, tx.Reason)
    }
    fmt.Printf("Hash: %s\n", b.Hash)
}

