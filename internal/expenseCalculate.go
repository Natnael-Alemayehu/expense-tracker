package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/natnael-alemayehu/expense-tracker/internal/data"
)

var (
	dateForamt = "2006-01-02"
)

func ReadFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0644)
	if err == os.ErrNotExist {
		file, err := os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("error creating file: %q", err)
		}
		return file, nil
	} else if err != nil {
		return nil, fmt.Errorf("error opening file: %q", err)
	}
	return file, nil
}

func AddExpense(description string, amount float64) error {

	expenses, err := ReadExpenses(FILENAME)
	if err != nil {
		return fmt.Errorf("failed to read expenses: %v", err)
	}
	nextID := 1
	if len(expenses) > 0 {
		lastExpense := expenses[len(expenses)-1]
		nextID = lastExpense.Id + 1
	}

	file, err := ReadFile(FILENAME)
	if err != nil {
		return err
	}
	defer file.Close()

	date := time.Now().Format(dateForamt)

	_, err = file.WriteString(fmt.Sprintf("id: %d,date: %v,description: %v,amount: %.2f\n", nextID, date, description, amount))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func ReadExpenses(filename string) ([]data.Expense, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var expenses []data.Expense
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		// Extract id
		idPart := strings.TrimPrefix(parts[0], "id: ")
		idPart = strings.TrimSpace(idPart)
		id, err := strconv.Atoi(idPart)
		if err != nil {
			return nil, fmt.Errorf("error parsing ID: %v", err)
		}

		// Extract time
		timePart := strings.TrimPrefix(parts[1], "date: ")
		timePart = strings.TrimSpace(timePart)
		date, err := time.Parse(dateForamt, timePart)
		if err != nil {
			return nil, fmt.Errorf("error parsing date: %v", err)
		}
		formattedDate := date.Format(dateForamt)

		// Extract description
		descPart := strings.TrimPrefix(parts[2], "description: ")
		descPart = strings.TrimSpace(descPart)

		// Extract amount
		amountPart := strings.TrimPrefix(parts[3], "amount: ")
		amountPart = strings.TrimSpace(amountPart)
		amount, err := strconv.ParseFloat(amountPart, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse amount: %v", err)
		}

		expenses = append(expenses, data.Expense{
			Id:          id,
			Date:        formattedDate,
			Description: descPart,
			Amount:      amount,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %v", err)
	}

	return expenses, nil
}

// Calcualte summary with 0 month will calculate the whole summary
func CalculateSummary(month int) (total int, err error) {
	expenses, err := ReadExpenses(FILENAME)
	if err != nil {
		return 0, err
	}

	month_ := time.Month(month)

	total = 0
	if month == 0 {
		for _, val := range expenses {
			total += int(val.Amount)
		}
		return total, nil
	} else if month >= 1 && month <= 12 {
		for _, val := range expenses {
			d, err := time.Parse(dateForamt, val.Date)
			if err != nil {
				return 0, fmt.Errorf("error Parsing Date: %v", err)
			}
			if d.Month() == month_ {
				total += int(val.Amount)
			}
		}
		return total, nil
	} else {
		return 0, fmt.Errorf("invalid month")
	}
}

func Delete(idArg int) error {
	expenses, err := ReadExpenses(FILENAME)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	idFound := true
	for _, v := range expenses {
		if v.Id == idArg {
			expenses = append(expenses[:idArg-1], expenses[idArg:]...)
			idFound = false
			break
		}
	}
	if idFound {
		return fmt.Errorf("id not found")
	}

	if err = WriteExpense(expenses); err != nil {
		return err
	}

	return nil
}

func WriteExpense(expenses []data.Expense) error {
	file, err := ReadFile(FILENAME)
	if err != nil {
		return err
	}
	defer file.Close()

	os.Truncate(file.Name(), 0)
	file.Seek(0, 0)

	for _, v := range expenses {
		_, err = file.WriteString(fmt.Sprintf("id: %d,date: %v,description: %v,amount: %.2f\n", v.Id, v.Date, v.Description, v.Amount))
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}

	}
	return nil
}
