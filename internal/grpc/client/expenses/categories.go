package client

import (
	"context"
	"log"

	"github.com/rubengomes8/golang-personal-finances-client/internal/pb/expenses/categories"
)

func CreateCategory(serviceClient categories.ExpenseCategoryServiceClient) {

	log.Println("CreateCategory was invoked")

	category := categories.ExpenseCategoryCreateRequest{
		Name: "Test",
	}

	res, err := serviceClient.CreateExpenseCategory(context.Background(), &category)
	if err != nil {
		log.Fatalf("client could not request for create category: %v\n", err)
	}

	log.Printf("Expense category created with ID: %d\n", res.Id)
}
