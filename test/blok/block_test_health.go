package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)


func TestNewHealthBlock(t *testing.T) {
	records := []HealthRecord{
		{PatientID: "patient-001", Temperature: 37.5, BloodPressure: "120/80", PulseRate: 75, Timestamp: time.Now()},
		{PatientID: "patient-002", Temperature: 38.0, BloodPressure: "130/85", PulseRate: 80, Timestamp: time.Now()},
	}
	block := NewHealthBlock(1, "0000000000000000", records)

	if block.Index != 1 {
		t.Errorf("Expected index 1, got %d", block.Index)
	}
	if block.PreviousHash != "0000000000000000" {
		t.Errorf("Expected previous hash '0000000000000000', got %s", block.PreviousHash)
	}
	if len(block.Records) != 2 {
		t.Errorf("Expected 2 records, got %d", len(block.Records))
	}
}

func TestCalculateHealthHash(t *testing.T) {
	records := []HealthRecord{
		{PatientID: "patient-001", Temperature: 37.5, BloodPressure: "120/80", PulseRate: 75, Timestamp: time.Now()},
	}
	block := NewHealthBlock(1, "0000000000000000", records)
	expectedHash := block.CalculateHash()

	if block.Hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, block.Hash)
	}
}

func TestPrintHealthBlock(t *testing.T) {
	records := []HealthRecord{
		{PatientID: "patient-001", Temperature: 37.5, BloodPressure: "120/80", PulseRate: 75, Timestamp: time.Now()},
		{PatientID: "patient-002", Temperature: 38.0, BloodPressure: "130/85", PulseRate: 80, Timestamp: time.Now()},
	}
	block := NewHealthBlock(1, "0000000000000000", records)

	block.PrintBlock() // Burada çıktıyı görsel olarak inceleyebilirsin.
}

