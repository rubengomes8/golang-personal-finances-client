package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	client "github.com/rubengomes8/golang-personal-finances-client/internal/grpc/client/expenses"
	"github.com/rubengomes8/golang-personal-finances-client/internal/pb/expenses"
)

const ADDR = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// cardsclient.CreateCard(cCards)
	// time.Sleep(500 * time.Millisecond)

	// cCards := cards.NewCardServiceClient(conn)

	cExpenses := expenses.NewExpensesServiceClient(conn)

	// --------------------------
	fmt.Println()
	restaurantExpense := createTextExpense(50.0, "Laser", "Restaurants", "Food allowance", time.Now().UTC().Unix())
	_, err = client.CreateExpense(cExpenses, &restaurantExpense)
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------
	fmt.Println()
	houseExpenseDate := time.Now().UTC().Add(72 * time.Hour).Unix()
	houseExpense := createTextExpense(200.0, "House", "Rent", "CGD", houseExpenseDate)
	houseExpenseId, err := client.CreateExpense(cExpenses, &houseExpense)
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------
	fmt.Println()
	houseExpenseUpdate := updateTextExpense(houseExpenseId, 250.0, "House", "Rent", "CGD", houseExpenseDate)
	_, err = client.UpdateExpense(cExpenses, &houseExpenseUpdate)
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------
	fmt.Println()
	card := expenses.ExpensesGetRequestByCard{
		Card: "CGD",
	}
	expensesGetByCard, err := client.GetExpensesByCard(cExpenses, &card) // SHOULD RETURN CGD
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(expensesGetByCard.Expenses)

	// --------------------------
	fmt.Println()
	expCategory := expenses.ExpensesGetRequestByCategory{
		Category: "House",
	}
	expensesGetByCategory, err := client.GetExpensesByCategory(cExpenses, &expCategory) // SHOULD RETURN RENT
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(expensesGetByCategory.Expenses)

	// --------------------------
	fmt.Println()
	expSubCategory := expenses.ExpensesGetRequestBySubCategory{
		SubCategory: "Restaurants",
	}
	expensesGetBySubCategory, err := client.GetExpensesBySubCategory(cExpenses, &expSubCategory) // SHOULD RETURN RESTAURANTS
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(expensesGetBySubCategory.Expenses)

	// --------------------------
	fmt.Println()
	minDateUnix := time.Now().UTC().Add(70 * time.Hour).Unix()
	maxDateUnix := time.Now().UTC().Add(74 * time.Hour).Unix()
	fmt.Printf("minDateUnix: %v\n", minDateUnix)
	fmt.Printf("maxDateUnix: %v\n", maxDateUnix)
	expDate := expenses.ExpensesGetRequestByDate{
		MinDate: time.Now().UTC().Add(70 * time.Hour).Unix(),
		MaxDate: time.Now().UTC().Add(74 * time.Hour).Unix(),
	}
	expensesGetByDate, err := client.GetExpensesByDate(cExpenses, &expDate) // SHOULD RETURN RESTAURANTS
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(expensesGetByDate.Expenses)

}

func createTextExpense(value float64, category, subCategory, card string, expenseDate int64) expenses.ExpenseCreateRequest {
	return expenses.ExpenseCreateRequest{
		Value:       value,
		Date:        expenseDate,
		Category:    category,
		SubCategory: subCategory,
		Card:        card,
		Description: "TEST",
	}
}

func updateTextExpense(id int64, value float64, category, subCategory, card string, expenseDate int64) expenses.ExpenseUpdateRequest {
	return expenses.ExpenseUpdateRequest{
		Id:          id,
		Value:       value,
		Date:        expenseDate,
		Category:    category,
		SubCategory: subCategory,
		Card:        card,
		Description: "TEST",
	}
}
