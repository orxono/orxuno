package block

import (
    "fmt"
)

// BlockValidator yapısı
type BlockValidator struct {
    ValidatorSet map[string]bool
}

// NewBlockValidator yeni bir BlockValidator döndürür
func NewBlockValidator() *BlockValidator {
    return &BlockValidator{
        ValidatorSet: make(map[string]bool),
    }
}

// AddValidator yeni bir validatör ekler
func (bv *BlockValidator) AddValidator(validator string) {
    bv.ValidatorSet[validator] = true
}

// ValidateTransaction bir işlemin geçerliliğini kontrol eder
func (bv *BlockValidator) ValidateTransaction(transaction Transaction) bool {
    // İşlem validatörü: Geçerli bir validatör tarafından imzalanmış mı?
    return bv.ValidatorSet[transaction.From]
}

// ValidateBlock belirli bir bloğun geçerliliğini kontrol eder
func (bv *BlockValidator) ValidateBlock(currentBlock, previousBlock *Block) bool {
    if currentBlock.PrevHash != previousBlock.Hash {
        return false
    }
    if currentBlock.Hash != CalculateHash(currentBlock) {
        return false
    }

    // Tüm işlemleri doğrula
    for _, tx := range currentBlock.Transactions {
        if !bv.ValidateTransaction(tx) {
            return false // Geçersiz işlem
        }
    }

    return true
}

// ValidateBlockchain zincirin geçerliliğini kontrol eder
func (bv *BlockValidator) ValidateBlockchain(chain []*Block) bool {
    if len(chain) == 0 {
        return false
    }

    for i := 1; i < len(chain); i++ {
        if !bv.ValidateBlock(chain[i], chain[i-1]) {
            fmt.Println("Geçersiz blok bulundu:", chain[i].Hash)
            return false
        }
    }
    return true
}

