package internal

import (
	"testing"
)

func TestBuildTable(t *testing.T) {
	// Setup
	cleanup := setupTestFile(t)
	defer cleanup()

	// Add some test expenses
	AddExpense("Test Expense 1", 100.00)
	AddExpense("Test Expense 2", 200.00)

	err := BuildTale()
	if err != nil {
		t.Errorf("BuildTale() error = %v", err)
	}
}

func TestGoPretty(t *testing.T) {
	// Setup
	cleanup := setupTestFile(t)
	defer cleanup()

	expenses, err := ReadExpenses(FILENAME)
	if err != nil {
		t.Fatalf("Failed to read expenses: %v", err)
	}

	err = GoPretty(expenses)
	if err != nil {
		t.Errorf("GoPretty() error = %v", err)
	}
}
