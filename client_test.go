package iex

import (
	"net/http"
	"testing"
	"time"
)

func setupTestClient() *Client {
	return NewClient(&http.Client{
		Timeout: 5 * time.Second,
	})
}

func testTOPS(t *testing.T, symbols []string) {
	c := setupTestClient()
	result, err := c.GetTOPS(symbols)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != len(symbols) {
		t.Fatalf("Received %v results, expected %v", len(result), len(symbols))
	}

	// TODO(palpant): Test parsing with a mock http client.
}

func TestTOPS_OneSymbol(t *testing.T) {
	testTOPS(t, []string{"SPY"})
}

func TestTOPS_TwoSymbols(t *testing.T) {
	testTOPS(t, []string{"SPY", "AAPL"})
}

func TestTOPS_AllSymbols(t *testing.T) {
	c := setupTestClient()
	result, err := c.GetTOPS(nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatalf("Received %v results", len(result))
	}
}

func TestLast(t *testing.T) {
	c := setupTestClient()
	symbols := []string{"SPY", "AAPL"}
	result, err := c.GetLast(symbols)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != len(symbols) {
		t.Fatalf("Received %v results, expected %v", len(result), len(symbols))
	}
}

func TestHIST_OneDate(t *testing.T) {
	c := setupTestClient()
	testDate := time.Date(2017, time.June, 6, 0, 0, 0, 0, time.UTC)
	result, err := c.GetHIST(testDate)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatalf("Received zero results")
	}
}

func TestHIST_AllDates(t *testing.T) {
	c := setupTestClient()
	result, err := c.GetAllAvailableHIST()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatalf("Received zero results")
	}
}

func TestDEEP(t *testing.T) {
	c := setupTestClient()
	result, err := c.GetDEEP("SPY")
	if err != nil {
		t.Fatal(err)
	}

	if result.Symbol != "SPY" {
		t.Fatalf("Expected symbol = %v, got %v", "SPY", result.Symbol)
	}
}

func TestBook(t *testing.T) {
	c := setupTestClient()
	symbols := []string{"SPY"}
	result, err := c.GetBook(symbols)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != len(symbols) {
		t.Log(result)
		t.Fatalf("Received %v results, expected %v", len(result), len(symbols))
	}
}

func TestSymbols(t *testing.T) {
	c := setupTestClient()
	symbols, err := c.GetSymbols()
	if err != nil {
		t.Fatal(err)
	}

	if len(symbols) == 0 {
		t.Fatal("Received zero symbols")
	}

	symbol := symbols[0]
	if symbol.Symbol == "" || symbol.Name == "" || symbol.Date == "" {
		t.Fatal("Failed to decode symbol correctly")
	}
}

func TestMarkets(t *testing.T) {
	c := setupTestClient()
	markets, err := c.GetMarkets()
	if err != nil {
		t.Fatal(err)
	}

	if len(markets) == 0 {
		t.Fatal("Received zero markets")
	}
}

func TestGetHistoricalDaily(t *testing.T) {
	c := setupTestClient()
	stats, err := c.GetHistoricalDaily(&HistoricalDailyRequest{Last: 5})
	if err != nil {
		t.Fatal(err)
	}

	if len(stats) != 5 {
		t.Fatalf("Received %d historical daily stats, expected %d", len(stats), 5)
	}
}
