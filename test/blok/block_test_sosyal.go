package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)

// TestNewSosyalBlock yeni bir sosyal bloğun doğru şekilde oluşturulup oluşturulmadığını kontrol eder.
func TestNewSosyalBlock(t *testing.T) {
    transactions := CreateSampleSosyalTransactions()
    block := NewSosyalBlock(1, "0000000000000000", transactions)

    if block.Index != 1 {
        t.Errorf("Beklenen Index: 1, Ancak bulundu: %d", block.Index)
    }

    if block.PreviousHash != "0000000000000000" {
        t.Errorf("Beklenen Önceki Hash: 0000000000000000, Ancak bulundu: %s", block.PreviousHash)
    }

    if len(block.Transactions) != len(transactions) {
        t.Errorf("Beklenen İşlem Sayısı: %d, Ancak bulundu: %d", len(transactions), len(block.Transactions))
    }
}

// TestPrintSosyalBlock bloğun bilgilerini yazdırma fonksiyonunun düzgün çalışıp çalışmadığını kontrol eder.
func TestPrintSosyalBlock(t *testing.T) {
    transactions := CreateSampleSosyalTransactions()
    block := NewSosyalBlock(1, "0000000000000000", transactions)

    // PrintSosyalBlock() fonksiyonu konsola yazdırdığı için burada kontrol edemeyiz.
    // Ancak, bu testin geçebilmesi için PrintSosyalBlock() fonksiyonu hatasız çalışmalıdır.
    block.PrintSosyalBlock()
}

// TestSampleSosyalTransactions sosyal işlemlerin doğru bir şekilde oluşturulup oluşturulmadığını kontrol eder.
func TestSampleSosyalTransactions(t *testing.T) {
    transactions := CreateSampleSosyalTransactions()

    if len(transactions) != 3 {
        t.Errorf("Beklenen işlem sayısı: 3, Ancak bulundu: %d", len(transactions))
    }

    if transactions[0].From != "Alice" {
        t.Errorf("Beklenen From: Alice, Ancak bulundu: %s", transactions[0].From)
    }

    if transactions[1].Content != "Hi Alice! I'm good, thanks!" {
        t.Errorf("Beklenen İçerik: Hi Alice! I'm good, thanks!, Ancak bulundu: %s", transactions[1].Content)
    }
}

