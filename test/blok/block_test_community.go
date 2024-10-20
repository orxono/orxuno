package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)

// TestNewCommunityBlock yeni bir CommunityBlock oluşturma testidir.
func TestNewCommunityBlock(t *testing.T) {
    transactions := []CommunityTransaction{
        {From: "Alice", To: "Bob", Content: "Hello Bob!", Amount: 0.0},
        {From: "Bob", To: "Charlie", Content: "Hi Charlie!", Amount: 0.0},
    }

    block := NewCommunityBlock(1, "0000000000000000", transactions)

    if block.Index != 1 {
        t.Errorf("Beklenen index: 1, mevcut index: %d", block.Index)
    }
    if block.PrevHash != "0000000000000000" {
        t.Errorf("Beklenen önceki hash: 0000000000000000, mevcut önceki hash: %s", block.PrevHash)
    }
    if len(block.Transactions) != 2 {
        t.Errorf("Beklenen işlem sayısı: 2, mevcut işlem sayısı: %d", len(block.Transactions))
    }
}

// TestPrintCommunityBlock bloğun içeriğini yazdırma testidir.
func TestPrintCommunityBlock(t *testing.T) {
    transactions := []CommunityTransaction{
        {From: "Alice", To: "Bob", Content: "Hello Bob!", Amount: 0.0},
        {From: "Bob", To: "Charlie", Content: "Hi Charlie!", Amount: 0.0},
    }

    block := NewCommunityBlock(1, "0000000000000000", transactions)

    // PrintBlock metodunu çağırıyoruz; burada çıktı kontrolü yapmayacağız
    block.PrintBlock() // Konsolda çıkacak, test ederken gözle kontrol edebilirsiniz
}

// TestSampleCommunityTransactions örnek topluluk işlemleri testidir.
func TestSampleCommunityTransactions(t *testing.T) {
    transactions := []CommunityTransaction{
        {From: "Alice", To: "Bob", Content: "Hello Bob!", Amount: 0.0},
        {From: "Bob", To: "Charlie", Content: "Hi Charlie!", Amount: 0.0},
        {From: "Charlie", To: "David", Content: "Just sharing a funny meme!", Amount: 0.0},
    }

    if len(transactions) != 3 {
        t.Errorf("Beklenen işlem sayısı: 3, mevcut işlem sayısı: %d", len(transactions))
    }
}

