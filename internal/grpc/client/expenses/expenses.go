package client

import (
	"context"
	"fmt"
	"log"

	"github.com/rubengomes8/golang-personal-finances-client/internal/pb/expenses"
)

func CreateExpense(serviceClient expenses.ExpensesServiceClient, expense *expenses.ExpenseCreateRequest) (int64, error) {

	log.Println("CreateExpense was invoked")

	res, err := serviceClient.CreateExpense(context.Background(), expense)
	if err != nil {
		return 0, fmt.Errorf("client could not request for create expense: %v", err)
	}

	log.Printf("Requested create expense with ID: %d\n", res.Id)
	return res.Id, nil
}

func CreateExpenses(serviceClient expenses.ExpensesServiceClient) ([]int64, error) {

	// TODO
	log.Println("CreateExpenses was invoked")

	expenses := expenses.ExpensesCreateRequest{}

	res, err := serviceClient.CreateExpenses(context.Background(), &expenses)
	if err != nil {
		return []int64{}, fmt.Errorf("client could not request for create expenses: %v", err)
	}

	var ids []int64
	for _, resId := range res.Ids {
		ids = append(ids, resId.Id)
	}

	return ids, nil
}

func UpdateExpense(serviceClient expenses.ExpensesServiceClient, expense *expenses.ExpenseUpdateRequest) (int64, error) {

	log.Println("UpdateExpense was invoked")

	res, err := serviceClient.UpdateExpense(context.Background(), expense)
	if err != nil {
		return 0, fmt.Errorf("client could not request for update expense: %v", err)
	}

	return res.Id, nil
}

func GetExpensesByCard(serviceClient expenses.ExpensesServiceClient, cardReq *expenses.ExpensesGetRequestByCard) (*expenses.ExpensesGetResponse, error) {

	log.Println("GetExpensesByCard was invoked")

	res, err := serviceClient.GetExpensesByCard(context.Background(), cardReq)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("client could not request a get expense by card: %v", err)
	}

	return res, nil
}

func GetExpensesByCategory(serviceClient expenses.ExpensesServiceClient, categoryReq *expenses.ExpensesGetRequestByCategory) (*expenses.ExpensesGetResponse, error) {

	log.Println("GetExpensesByCategory was invoked")

	res, err := serviceClient.GetExpensesByCategory(context.Background(), categoryReq)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("client could not request a get expense by category: %v", err)
	}

	return res, nil
}

func GetExpensesBySubCategory(serviceClient expenses.ExpensesServiceClient, subCategoryReq *expenses.ExpensesGetRequestBySubCategory) (*expenses.ExpensesGetResponse, error) {

	log.Println("GetExpensesBySubCategory was invoked")

	res, err := serviceClient.GetExpensesBySubCategory(context.Background(), subCategoryReq)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("client could not request a get expense by subcategory: %v", err)
	}

	return res, nil
}

func GetExpensesByDate(serviceClient expenses.ExpensesServiceClient, datesReq *expenses.ExpensesGetRequestByDate) (*expenses.ExpensesGetResponse, error) {

	log.Println("GetExpensesByDate was invoked")

	res, err := serviceClient.GetExpensesByDate(context.Background(), datesReq)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("client could not request a get expense by date: %v", err)
	}

	return res, nil
}
