package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	client "github.com/rubengomes8/golang-personal-finances-client/internal/grpc/client/expenses"
	"github.com/rubengomes8/golang-personal-finances-client/proto/expenses"
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

	restaurantExpense := createTextExpense(50.0, "Laser", "Restaurants", "Food allowance")
	_ = client.CreateExpense(cExpenses, &restaurantExpense)

	houseExpense := createTextExpense(200.0, "House", "Rent", "CGD")
	houseExpenseId := client.CreateExpense(cExpenses, &houseExpense)

	houseExpenseUpdate := updateTextExpense(houseExpenseId, 250.0, "House", "Rent", "CGD")
	_ = client.UpdateExpense(cExpenses, &houseExpenseUpdate)

	card := expenses.ExpensesGetRequestByCard{
		Card: "CGD",
	}
	client.GetExpensesByCard(cExpenses, &card)
	time.Sleep(500 * time.Millisecond)

	// client.GetExpensesByCategory(c)
	// time.Sleep(500 * time.Millisecond)

	// client.GetExpensesBySubCategory(c)
	// time.Sleep(500 * time.Millisecond)

	// client.GetExpensesByDate(c)
	// time.Sleep(500 * time.Millisecond)
}

func createTextExpense(value float64, category, subCategory, card string) expenses.ExpenseCreateRequest {
	return expenses.ExpenseCreateRequest{
		Value:       value,
		Date:        time.Now().UTC().Unix(),
		Category:    category,
		SubCategory: subCategory,
		Card:        card,
		Description: "TEST",
	}
}

func updateTextExpense(id int64, value float64, category, subCategory, card string) expenses.ExpenseUpdateRequest {
	return expenses.ExpenseUpdateRequest{
		Id:          id,
		Value:       value,
		Date:        time.Now().UTC().Unix(),
		Category:    category,
		SubCategory: subCategory,
		Card:        card,
		Description: "TEST",
	}
}
