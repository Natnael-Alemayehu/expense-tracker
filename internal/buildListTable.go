package internal

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/natnael-alemayehu/expense-tracker/internal/data"
)

var (
	id          = "#"
	date        = "Date"
	description = "Description"
	amount      = "Amount"
	rowHeader   = table.Row{id, date, description, amount}
)

func BuildTale() error {
	expenses, err := ReadExpenses(FILENAME)
	if err != nil {
		return nil
	}

	GoPretty(expenses)
	return nil
}

func GoPretty(expenses []data.Expense) error {

	// Fetching Total
	total, err := CalculateSummary()
	if err != nil {
		return err
	}

	tw := table.NewWriter()
	tw.AppendHeader(rowHeader)
	rowFooter := table.Row{"", "", "Total", total}
	tw.AppendFooter(rowFooter)

	for _, val := range expenses {
		tw.AppendRow(table.Row{
			val.Id, val.Date, val.Description, val.Amount,
		})
	}
	fmt.Println(tw.Render())
	return nil

}
