package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/natnael-alemayehu/expense-tracker/internal"
)

func main() {
	// Creates a subdomain add
	addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	deletecmd := flag.NewFlagSet("delete", flag.ExitOnError)
	summarycmd := flag.NewFlagSet("summary", flag.ExitOnError)

	// flags for add
	addDescription := addcmd.String("description", "", "Description of the expense")
	addAmount := addcmd.Float64("amount", 0, "Amount of the expense")

	// flags for delete
	deleteId := deletecmd.Int("ID", 0, "Deletes expense with the provided id")

	// flags for summary
	monthInt := summarycmd.Int("month", 0, "Months for summary report")

	if len(os.Args) < 2 {
		fmt.Print("Expected command (add, list, summary, delete) ")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addcmd.Parse(os.Args[2:])
		if *addDescription == "" {
			fmt.Println("Error: Description not provided")
			addcmd.PrintDefaults()
			os.Exit(1)
		}
		if *addAmount <= 0 {
			fmt.Println("Error: Ammount can not be negative or 0")
			addcmd.PrintDefaults()
			os.Exit(1)
		}
		err := internal.AddExpense(*addDescription, *addAmount)
		if err != nil {
			fmt.Print(err.Error())
		}

	case "list":
		if err := internal.BuildTale(); err != nil {
			fmt.Println("Building list table failed")
			os.Exit(1)
		}
	case "summary":
		summarycmd.Parse(os.Args[2:])
		if *monthInt == 0 {
			// Calcualte summary with 0 month will calculate the whole summary
			total, err := internal.CalculateSummary(0)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			output := fmt.Sprintf("Total expenses: %v \n", total)
			fmt.Print(output)
		} else if *monthInt <= 0 || *monthInt > 12 {
			fmt.Println("Month command is not correctly set: (1,12)")
			os.Exit(1)
		} else {
			total, err := internal.CalculateSummary(*monthInt)
			if err != nil {
				fmt.Printf("Error Calculating summary: %v \n", err)
				os.Exit(1)
			}
			fmt.Printf("Total expenses for %s: %v \n", time.Month(*monthInt), total)
		}
	case "delete":
		deletecmd.Parse(os.Args[2:])
		if *deleteId <= 0 {
			fmt.Println("Id doesn't exist")
			deletecmd.PrintDefaults()
			os.Exit(1)
		}
		fmt.Printf("Expense with ID: %d, Deleted Successfully\n", *deleteId)
	default:
		fmt.Println("Command does not exist does not exist")
		os.Exit(1)
	}

}
