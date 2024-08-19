# go-openapi3-sample

### Things todo list

1. Clone this repository: `git clone https://github.com/hendisantika/go-openapi3-sample.git`
2. Navigate to the folder: `cd go-openapi3-sample`
3. Run the application: `go run main.go`
4. Install Go Swagger CLI: `go install github.com/go-swagger/go-swagger/cmd/swagger@latest`
5. Ask go-swagger to generate our documentation by running the following command in a new terminal window:
   `swagger generate spec -o ./swagger.json`
6. Let’s now serve Swagger UI to actually see this documentation and try to interact with it by running the following
   command: `swagger serve -F=swagger swagger.json`