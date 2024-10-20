package block

import (
    "testing"
    "orxuno/block" // Block modülünü import edin
)

func TestNewWeatherBlock(t *testing.T) {
	reports := []WeatherReport{
		{City: "Istanbul", Temperature: 20.5, Humidity: 65.0, Timestamp: time.Now()},
		{City: "Ankara", Temperature: 15.0, Humidity: 50.0, Timestamp: time.Now()},
	}
	block := NewWeatherBlock(1, "0000000000000000", reports)

	if block.Index != 1 {
		t.Errorf("Expected index 1, got %d", block.Index)
	}
	if block.PreviousHash != "0000000000000000" {
		t.Errorf("Expected previous hash '0000000000000000', got %s", block.PreviousHash)
	}
	if len(block.Reports) != 2 {
		t.Errorf("Expected 2 reports, got %d", len(block.Reports))
	}
}

func TestCalculateWeatherHash(t *testing.T) {
	reports := []WeatherReport{
		{City: "Istanbul", Temperature: 20.5, Humidity: 65.0, Timestamp: time.Now()},
	}
	block := NewWeatherBlock(1, "0000000000000000", reports)
	expectedHash := block.CalculateHash()

	if block.Hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, block.Hash)
	}
}

func TestPrintWeatherBlock(t *testing.T) {
	reports := []WeatherReport{
		{City: "Istanbul", Temperature: 20.5, Humidity: 65.0, Timestamp: time.Now()},
		{City: "Ankara", Temperature: 15.0, Humidity: 50.0, Timestamp: time.Now()},
	}
	block := NewWeatherBlock(1, "0000000000000000", reports)

	block.PrintBlock() // Burada çıktıyı görsel olarak inceleyebilirsin.
}

