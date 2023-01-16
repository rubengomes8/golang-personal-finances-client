GO_MODULE=github.com/rubengomes8/golang-personal-finances-client
BIN_DIR=bin
EXPENSES_CODE_DIR=./internal/pb/expenses/
EXPENSES_CATEGORIES_CODE_DIR=./internal/pb/expenses/categories/
EXPENSES_SUBCATEGORIES_CODE_DIR=./internal/pb/expenses/subcategories/
CARDS_CODE_DIR=./internal/pb/cards/

# CARDS #
cards:
	protoc --proto_path=./proto --go_out=${CARDS_CODE_DIR} --go_opt=module=${GO_MODULE} --go-grpc_out=${CARDS_CODE_DIR} --go-grpc_opt=module=${GO_MODULE} cards.proto

# EXPENSES CATEGORIES #
expense_categories:
	protoc --proto_path=./proto --go_out=${EXPENSES_CATEGORIES_CODE_DIR} --go_opt=module=${GO_MODULE} --go-grpc_out=${EXPENSES_CATEGORIES_CODE_DIR} --go-grpc_opt=module=${GO_MODULE} expense_categories.proto


# EXPENSES SUB CATEGORIES #
expense_subcategories:
	protoc --proto_path=./proto --go_out=${EXPENSES_SUBCATEGORIES_CODE_DIR} --go_opt=module=${GO_MODULE} --go-grpc_out=${EXPENSES_SUBCATEGORIES_CODE_DIR} --go-grpc_opt=module=${GO_MODULE} expense_subcategories.proto

# EXPENSES #
expenses:
	protoc --proto_path=./proto --go_out=${EXPENSES_CODE_DIR} --go_opt=module=${GO_MODULE} --go-grpc_out=${EXPENSES_CODE_DIR} --go-grpc_opt=module=${GO_MODULE} expenses.proto

all: cards expense_categories expense_subcategories expenses
# BUILD GO #
build-expenses:
	go build -o ${BIN_DIR}/expenses/client ./cmd/grpc-client/main.go

build: build-expenses