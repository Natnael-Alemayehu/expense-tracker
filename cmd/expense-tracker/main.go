package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/natnael-alemayehu/expense-tracker/internal"
)

func main() {
	// Creates a subdomain add
	addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	deletecmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// flags for add
	addDescription := addcmd.String("description", "", "Description of the expense")
	addAmount := addcmd.Float64("amount", 0, "Amount of the expense")

	// flags for delete
	deleteId := deletecmd.Int("ID", 0, "Deletes expense with the provided id")

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
		total, err := internal.CalculateSummary()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		output := fmt.Sprintf("Total expenses: $", total)
		fmt.Print(output)
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
