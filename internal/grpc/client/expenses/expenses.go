package client

import (
	"context"
	"log"

	"github.com/rubengomes8/golang-personal-finances-client/proto/expenses"
)

func CreateExpense(serviceClient expenses.ExpensesServiceClient, expense *expenses.ExpenseCreateRequest) int64 {

	log.Println("CreateExpense was invoked")

	res, err := serviceClient.CreateExpense(context.Background(), expense)
	if err != nil {
		log.Fatalf("client could not request for create expense: %v\n", err)
	}

	log.Printf("Requested create expense with ID: %d\n", res.Id)
	return res.Id
}

func CreateExpenses(serviceClient expenses.ExpensesServiceClient) {

	// TODO
	log.Println("CreateExpenses was invoked")

	expenses := expenses.ExpensesCreateRequest{}

	res, err := serviceClient.CreateExpenses(context.Background(), &expenses)
	if err != nil {
		log.Fatalf("client could not request for create expenses: %v\n", err)
	}

	log.Printf("Requested create expenses with IDs: %v\n", res.Ids)
}

func UpdateExpense(serviceClient expenses.ExpensesServiceClient, expense *expenses.ExpenseUpdateRequest) int64 {

	log.Println("UpdateExpense was invoked")

	res, err := serviceClient.UpdateExpense(context.Background(), expense)
	if err != nil {
		log.Fatalf("client could not request for update expense: %v\n", err)
	}

	log.Printf("Requested update expense with ID: %d\n", res.Id)
	return res.Id
}

func GetExpensesByCard(serviceClient expenses.ExpensesServiceClient, card *expenses.ExpensesGetRequestByCard) {

	log.Println("GetExpensesByCard was invoked")

	res, err := serviceClient.GetExpensesByCard(context.Background(), card)
	if err != nil {
		log.Fatalf("client could not request a get expense by card: %v\n", err)
	}

	log.Printf("Requested get expenses by card: %v\n", res.Expenses)
}

func GetExpensesByCategory(serviceClient expenses.ExpensesServiceClient) {

	log.Println("GetExpensesByCategory was invoked")

	expense := expenses.ExpensesGetRequestByCategory{
		Category: "House",
	}

	res, err := serviceClient.GetExpensesByCategory(context.Background(), &expense)
	if err != nil {
		log.Fatalf("client could not request a get expense by category: %v\n", err)
	}

	log.Printf("Requested get expenses by category: %v\n", res.Expenses)
}

func GetExpensesBySubCategory(serviceClient expenses.ExpensesServiceClient) {

	log.Println("GetExpensesBySubCategory was invoked")

	expense := expenses.ExpensesGetRequestBySubCategory{
		SubCategory: "Rent",
	}

	res, err := serviceClient.GetExpensesBySubCategory(context.Background(), &expense)
	if err != nil {
		log.Fatalf("client could not request a get expense by subcategory: %v\n", err)
	}

	log.Printf("Requested get expenses by subcategory: %v\n", res.Expenses)
}

func GetExpensesByDate(serviceClient expenses.ExpensesServiceClient) {

	log.Println("GetExpensesByDate was invoked")

	expense := expenses.ExpensesGetRequestByDate{
		MinDate: 1,
		MaxDate: 1000,
	}

	res, err := serviceClient.GetExpensesByDate(context.Background(), &expense)
	if err != nil {
		log.Fatalf("client could not request a get expense by date: %v\n", err)
	}

	log.Printf("Requested get expenses by date: %v\n", res.Expenses)
}
