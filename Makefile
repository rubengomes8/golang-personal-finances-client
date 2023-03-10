GO_MODULE=github.com/rubengomes8/golang-personal-finances-client
BIN_DIR=bin

# CARDS #
cards:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} cards.proto

# EXPENSES CATEGORIES #
expense_categories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expense_categories.proto

# EXPENSES SUB CATEGORIES #
expense_subcategories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expense_subcategories.proto

# EXPENSES #
expenses:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expenses.proto

all: cards expense_categories expense_subcategories expenses
# BUILD GO #
build-expenses:
	go build -o ${BIN_DIR}/expenses/client ./cmd/grpc-client/main.go

build: build-expenses