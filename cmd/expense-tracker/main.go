package main

import (
	"flag"
	"fmt"
	"os"
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
		fmt.Printf("description: %v, amount: %v \n", *addDescription, *addAmount)
	case "list":
		fmt.Println("Printing the list")
	case "summary":
		fmt.Println("Printing summary")
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
