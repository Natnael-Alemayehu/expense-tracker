package internal

import (
	"os"
	"testing"
	"time"

	"github.com/natnael-alemayehu/expense-tracker/internal/data"
)

const testFileName = "test_expenses.txt"

func TestAddExpense(t *testing.T) {
	// Setup
	cleanup := setupTestFile(t)
	defer cleanup()

	tests := []struct {
		name        string
		description string
		amount      float64
		wantErr     bool
	}{
		{
			name:        "Valid expense",
			description: "Test expense",
			amount:      100.50,
			wantErr:     false,
		},
		{
			name:        "Zero amount",
			description: "Zero expense",
			amount:      0,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddExpense(tt.description, tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddExpense() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadExpenses(t *testing.T) {
	// Setup
	cleanup := setupTestFile(t)
	defer cleanup()

	// Add test data
	testExpenses := []data.Expense{
		{Id: 1, Description: "Test 1", Amount: 100.00, Date: time.Now().Format("2006-01-02")},
		{Id: 2, Description: "Test 2", Amount: 200.00, Date: time.Now().Format("2006-01-02")},
	}

	for _, exp := range testExpenses {
		AddExpense(exp.Description, exp.Amount)
	}

	// Test reading
	expenses, err := ReadExpenses(testFileName)
	if err != nil {
		t.Fatalf("ReadExpenses() error = %v", err)
	}

	if len(expenses) != len(testExpenses) {
		t.Errorf("ReadExpenses() got %d expenses, want %d", len(expenses), len(testExpenses))
	}
}

func TestCalculateSummary(t *testing.T) {
	// Setup
	cleanup := setupTestFile(t)
	defer cleanup()

	// Add test expenses
	AddExpense("January Expense", 100.00)
	AddExpense("February Expense", 200.00)

	tests := []struct {
		name      string
		month     int
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "Total summary",
			month:     0,
			wantTotal: 300,
			wantErr:   false,
		},
		{
			name:      "Invalid month",
			month:     13,
			wantTotal: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total, err := CalculateSummary(tt.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateSummary() error = %v, wantErr %v", err, tt.wantErr)
			}
			if total != tt.wantTotal && !tt.wantErr {
				t.Errorf("CalculateSummary() = %v, want %v", total, tt.wantTotal)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	// Setup
	cleanup := setupTestFile(t)
	defer cleanup()

	// Add test expense
	AddExpense("Test Expense", 100.00)

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "Delete existing expense",
			id:      1,
			wantErr: false,
		},
		{
			name:    "Delete non-existing expense",
			id:      999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Helper function to setup test environment
func setupTestFile(t *testing.T) func() {
	originalFilename := FILENAME
	FILENAME = testFileName

	return func() {
		FILENAME = originalFilename
		os.Remove(testFileName)
	}
}
