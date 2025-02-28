package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	loadExpenses()

	for {

		fmt.Println("Welcome to Track Expenser")
		fmt.Println("1. Add expense")
		fmt.Println("2. View expenses")
		fmt.Println("3. Total all expenses by category")
		fmt.Println("4. Export to CSV")
		fmt.Println("5. Exit")
		fmt.Println("Choose an option-> ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var category, desciption string
			var amount float64
			fmt.Println("Enter amount: ", amount)
			fmt.Scan(&amount)
			fmt.Println("Enter category: ", category)
			fmt.Scan(&category)
			fmt.Println("Enter description: ", desciption)
			fmt.Scan(&desciption)

			addExpense(amount, category, desciption)

		case 2:
			viewExpenses()

		case 3:
			totalByCategory()

		case 4:
			exportToCSV()

		case 5:
			fmt.Println("Exiting TrackEX.... GOOD BYEEE")
			return

		default:
			fmt.Println("X Invalid option. Try again")
		}
	}
}

//Expense struct to store each expense and operate on them.

type Expense struct {
	Date       string  `json: "date"`
	Amount     float64 `json: "amount"`
	Category   string  `json: "category"`
	Desciption string  `json: "description"`
}

var expenses []Expense

const fileName = "expenses.json"

// Load expenses from JSON file
func loadExpenses() {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("No data found. Starting afresh")
		return
	}
	json.Unmarshal(file, &expenses)
}

func addExpense(amount float64, category string, description string) {
	newExpense := Expense{
		Date:       time.Now().Format("2006-01-02"),
		Amount:     amount,
		Category:   category,
		Desciption: description,
	}

	expenses = append(expenses, newExpense)
	saveExpenses()
	fmt.Println("Expense added successfully")
}

func saveExpenses() {
	data, _ := json.MarshalIndent(expenses, "", " ")
	os.WriteFile(fileName, data, 0644)
}

func viewExpenses() {
	fmt.Println("All Expenses: ")
	for _, exp := range expenses {
		fmt.Printf("[%s]  %s - Rs.%.2f (%s)\n", exp.Date, exp.Category, exp.Amount, exp.Desciption)
	}
}

func totalByCategory() {
	categoryTotals := make(map[string]float64)
	for _, exp := range expenses {
		categoryTotals[exp.Category] += exp.Amount
	}

	for cap, total := range categoryTotals {
		fmt.Printf("%s: Rs.%2f", cap, total)
	}
}

func exportToCSV() {
	file, err := os.Create("expenses.csv")
	if err != nil {
		fmt.Println("Error creating file, please retry")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	//write headers
	writer.Write([]string{"Date", "Category", "Amount", "Description"})

	//write expenses
	for _, exp := range expenses {
		writer.Write([]string{exp.Date, exp.Category, strconv.FormatFloat(exp.Amount, 'f', 2, 64), exp.Desciption})
	}

	fmt.Println("Files updated as CSV successfully expenses.csv")

}
