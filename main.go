package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var valueTotal float64

type ConvertNumber struct {
	Number float64 `json:number`
}

type Expense struct {
	Description string
	Amount      float64
}

func readExpensesFromCSV(filename string) ([]Expense, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var expenses []Expense
	for _, line := range lines {
		amount, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, Expense{
			Description: line[0],
			Amount:      amount,
		})
	}

	return expenses, nil
}

func calculateTotal(expenses []Expense) float64 {
	var total float64
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total
}

func main() {
	expenses, err := readExpensesFromCSV("expenses.csv")
	if err != nil {
		log.Fatal(err)
	}
	total := calculateTotal(expenses)
	valueTotal = total

	http.HandleFunc("/", ListValue)

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}

func ListValue(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Server Runing")
	value := ConvertNumber{valueTotal}
	totalValue, _ := json.Marshal(value)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write(totalValue)
	if err != nil {
		return
	}
}
