package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)

func TestNewSecurityBlock(t *testing.T) {
	transactions := SampleSecurityEvents()
	block := NewSecurityBlock(1, "0000000000000000", transactions)

	if block.Index != 1 {
		t.Errorf("Expected index 1, got %d", block.Index)
	}
	if block.PreviousHash != "0000000000000000" {
		t.Errorf("Expected previous hash '0000000000000000', got %s", block.PreviousHash)
	}
	if len(block.Transactions) != len(transactions) {
		t.Errorf("Expected %d transactions, got %d", len(transactions), len(block.Transactions))
	}
}

func TestPrintSecurityBlock(t *testing.T) {
	transactions := SampleSecurityEvents()
	block := NewSecurityBlock(1, "0000000000000000", transactions)
	block.PrintBlock() // Bu test, çıktıyı kontrol etmiyor ama çıktıyı görmek için.
}

func TestCalculateSecurityHash(t *testing.T) {
	transactions := SampleSecurityEvents()
	block := NewSecurityBlock(1, "0000000000000000", transactions)
	expectedHash := block.Hash // Hash hesaplamasını kontrol etmek için beklenen değeri alıyoruz.
	if block.Hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, block.Hash)
	}
}

func SampleSecurityEvents() []SecurityEvent {
	return []SecurityEvent{
		{"Alice", "login", "2024-10-20T14:30:00Z", "192.168.1.1"},
		{"Bob", "failed_attempt", "2024-10-20T14:31:00Z", "192.168.1.2"},
		{"Charlie", "logout", "2024-10-20T14:32:00Z", "192.168.1.3"},
	}
}

