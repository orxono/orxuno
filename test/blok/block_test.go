package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)

func TestBlockValidatorWithValidSet(t *testing.T) {
    bv := NewBlockValidator()

    // Geçerli validatörler ekle
    bv.AddValidator("Alice")
    bv.AddValidator("Bob")

    // Geçerli bloklar oluştur
    block1 := NewBlock(0, "", []Transaction{
        {From: "Alice", To: "Bob", Amount: 10, Type: "transfer", Signature: "Alice's Signature"},
    }, 0) // Doğru sırada parametreleri geçin

    block2 := NewBlock(1, block1.Hash, []Transaction{
        {From: "Bob", To: "Alice", Amount: 5, Type: "transfer", Signature: "Bob's Signature"},
    }, 1) // Doğru sırada parametreleri geçin

    // Zincir oluştur
    chain := []*Block{block1, block2}

    // Zinciri doğrula
    if !bv.ValidateBlockchain(chain) {
        t.Error("Zincir geçerli olmalıydı, ancak geçersiz olarak işaretlendi.")
    }
}

func TestBlockValidatorWithInvalidSet(t *testing.T) {
    bv := NewBlockValidator()

    // Geçerli validatörler ekle
    bv.AddValidator("Alice")
    bv.AddValidator("Bob")

    // Geçersiz işlem içeren blok oluştur
    block1 := NewBlock(0, "", []Transaction{
        {From: "Alice", To: "Bob", Amount: 10, Type: "transfer", Signature: "Alice's Signature"},
    }, 0) // Doğru sırada parametreleri geçin

    block2 := NewBlock(1, block1.Hash, []Transaction{
        {From: "Charlie", To: "Alice", Amount: 5, Type: "transfer", Signature: "Charlie's Signature"}, // Geçersiz işlem
    }, 1) // Doğru sırada parametreleri geçin

    // Zincir oluştur
    chain := []*Block{block1, block2}

    // Zinciri doğrula
    if bv.ValidateBlockchain(chain) {
        t.Error("Zincir geçersiz olmalıydı, ancak geçerli olarak işaretlendi.")
    }
}

