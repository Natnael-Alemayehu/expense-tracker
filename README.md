# Expense Tracker

A command-line expense tracking application written in Go that helps you manage and monitor your expenses efficiently.

## Features

- Add new expenses with descriptions and amounts
- List all recorded expenses
- Generate expense summaries (monthly or total)
- Delete expenses by ID
- Persistent storage using text files

## Installation

1. Make sure you have Go installed on your system
2. Clone the repository:
```bash
git clone https://github.com/natnael-alemayehu/expense-tracker.git
cd expense-tracker
```
3. Build the application:
```bash
go build ./cmd/expense-tracker
```

## Usage

The application supports the following commands:

### Add an Expense
```bash
./expense-tracker add -description "Grocery shopping" -amount 50.50
```

### List All Expenses
```bash
./expense-tracker list
```

### Get Expense Summary
- For total summary:
```bash
./expense-tracker summary
```
- For monthly summary (replace N with month number 1-12):
```bash
./expense-tracker summary -month N
```

### Delete an Expense
```bash
./expense-tracker delete -id N
```

## Project Structure

- `cmd/expense-tracker/`: Contains the main application entry point
- `internal/`: Internal package code including expense calculations and data handling
- `expenseCalculate_test.go`: Tests for expense calculations
- `buildListTable_test.go`: Tests for table building functionality
- `expenses.txt`: Storage file for expense records

## License

This project is licensed under the terms included in the LICENSE file.

## Contributing

Feel free to submit issues and pull requests to help improve this project.

## Acknowledgments

This project is based on the [roadmap.sh](https://roadmap.sh/projects/expense-tracker) project ideas.