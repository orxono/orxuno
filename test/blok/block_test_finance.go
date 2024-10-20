package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)

func TestNewFinanceBlock(t *testing.T) {
    transactions := []FinanceTransaction{
        {"Alice", "Bob", 100.00, "Payment for services"},
        {"Charlie", "David", 250.50, "Loan repayment"},
    }

    block := NewFinanceBlock(1, "0000000000000000", transactions)

    if block.Index != 1 {
        t.Errorf("Expected block index to be 1, got %d", block.Index)
    }
    if block.PrevHash != "0000000000000000" {
        t.Errorf("Expected previous hash to be 0000000000000000, got %s", block.PrevHash)
    }
    if len(block.Transactions) != 2 {
        t.Errorf("Expected 2 transactions, got %d", len(block.Transactions))
    }
}

func TestPrintFinanceBlock(t *testing.T) {
    transactions := []FinanceTransaction{
        {"Alice", "Bob", 100.00, "Payment for services"},
        {"Charlie", "David", 250.50, "Loan repayment"},
    }

    block := NewFinanceBlock(1, "0000000000000000", transactions)

    block.PrintBlock()
}

func TestCalculateFinanceHash(t *testing.T) {
    transactions := []FinanceTransaction{
        {"Alice", "Bob", 100.00, "Payment for services"},
    }

    block := NewFinanceBlock(1, "0000000000000000", transactions)

    // Beklenen hash değerini burada hesaplayın
    expectedHash := CalculateFinanceHash(block)

    // Hesaplanan hash ile beklenen hash değerini karşılaştır
    if block.Hash != expectedHash {
        t.Errorf("Expected hash %s, got %s", expectedHash, block.Hash)
    }
}

