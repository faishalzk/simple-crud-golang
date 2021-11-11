# Simple CRUD Golang
This is a proof of concept of a Simple CRUD api in Golang

## Requirements
- Go 1.17.2
- Gorilla Mux

## How To Run
- Checkout this repository
- Install gorilla mux `go get -u github.com/gorilla/mux`
- Execute `go build` from `simple-crud-golang` directory with Terminal or Command Prompt
- Execute `go run main.go` from `simple-crud-golang` directory with Terminal or Command Prompt

## How To Run Test
- Checkout this repository
- Install gorilla mux `go get -u github.com/gorilla/mux`
- Execute `go build` from `simple-crud-golang` directory with Terminal or Command Prompt
- Execute `go test ./api/language -v LanguageService_test.go LanguageService.go` from `simple-crud-golang`
  directory with Terminal or Command Prompt
  
## Endpoints
- Hello page: Get `/`
- Check if a text is palindrome: Get `/palindrome/{text}`
- Get all programming languages from database: Get `/languages`
- Get programming language from database by id: Get `/language/{id}`
- Create new programming language: Post `/language`
- Update programming language by id: Patch `/language/{id}`
- Delete programming languages by id: Delete `/language/{id}`
