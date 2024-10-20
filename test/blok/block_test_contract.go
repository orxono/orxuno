package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)


func TestSmartContract(t *testing.T) {
    sc := NewSmartContract("contract-001")
    
    tx := Transaction{From: "Alice", To: "Bob", Amount: 10, Type: "transfer", Signature: "Alice's Signature"}
    sc.AddTransaction(tx)
    
    if len(sc.Transactions) != 1 {
        t.Errorf("Beklenen işlem sayısı 1, ancak %d alındı.", len(sc.Transactions))
    }

    sc.Disable()
    sc.AddTransaction(Transaction{From: "Bob", To: "Charlie", Amount: 5, Type: "transfer", Signature: "Bob's Signature"})
}

