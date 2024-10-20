package block

import (
    "fmt"
)

// SmartContract yapısı
type SmartContract struct {
    ID          string // Sözleşme kimliği
    Active      bool   // Sözleşmenin aktif durumu
    Transactions []Transaction // Sözleşmedeki işlemler
}

// NewSmartContract yeni bir akıllı sözleşme oluşturur
func NewSmartContract(id string) *SmartContract {
    return &SmartContract{
        ID:          id,
        Active:      true,
        Transactions: []Transaction{},
    }
}

// AddTransaction bir işlem ekler
func (sc *SmartContract) AddTransaction(tx Transaction) {
    if !sc.Active {
        fmt.Println("Sözleşme aktif değil, işlem eklenemedi.")
        return
    }
    sc.Transactions = append(sc.Transactions, tx)
    fmt.Printf("İşlem eklendi: %s -> %s, Miktar: %d\n", tx.From, tx.To, tx.Amount)
}

// Disable devre dışı bırakır
func (sc *SmartContract) Disable() {
    sc.Active = false
    fmt.Printf("Sözleşme devre dışı bırakıldı: %s\n", sc.ID)
}

// Enable etkinleştirir
func (sc *SmartContract) Enable() {
    sc.Active = true
    fmt.Printf("Sözleşme aktif hale getirildi: %s\n", sc.ID)
}

