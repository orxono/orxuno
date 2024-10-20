package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)

// TestNewStakeBlock yeni bir StakeBlock oluşturmayı test eder
func TestNewStakeBlock(t *testing.T) {
    transactions := []StakeTransaction{
        {From: "Alice", Amount: 100.0, Status: "completed"},
    }
    block := NewStakeBlock(1, "0000000000000000", transactions)

    if block.Index != 1 {
        t.Errorf("Beklenen blok indeksi 1, ancak %d alındı.", block.Index)
    }
    if block.PrevHash != "0000000000000000" {
        t.Errorf("Beklenen önceki hash '0000000000000000', ancak %s alındı.", block.PrevHash)
    }
    if len(block.Transactions) != 1 {
        t.Errorf("Beklenen işlem sayısı 1, ancak %d alındı.", len(block.Transactions))
    }
}

// TestStakeTransactionValidity geçerli bir stake işlemini test eder
func TestStakeTransactionValidity(t *testing.T) {
    tx := StakeTransaction{From: "Bob", Amount: 50.0, Status: "pending"}

    if tx.Amount <= 0 {
        t.Error("Miktar sıfırdan büyük olmalıdır.")
    }
}

// TestStakeBlockPrint stake bloğunun yazdırılmasını test eder
func TestStakeBlockPrint(t *testing.T) {
    transactions := []StakeTransaction{
        {From: "Alice", Amount: 100.0, Status: "completed"},
        {From: "Bob", Amount: 50.0, Status: "pending"},
    }
    block := NewStakeBlock(1, "0000000000000000", transactions)

    // Bu test, Print fonksiyonunun çıktısını kontrol etmek için kullanılabilir.
    // Konsol çıktısını kontrol etmek genellikle zordur, bu nedenle
    // işlevin başarılı bir şekilde çağrıldığını kontrol ediyoruz.
    block.Print() // Konsola yazdırır, bu bir hata olmamalıdır.
}

// TestSampleStakeTransactions örnek stake işlemlerinin varlığını test eder
func TestSampleStakeTransactions(t *testing.T) {
    if len(sampleStakeTransactions) == 0 {
        t.Error("Örnek stake işlemleri boş olmamalıdır.")
    }
}

