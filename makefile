build: # Build an executable on bin folder
	@go build -o ./starwars-api ./cmd/main.go

install: # install dependencies
	@go mod tidy

start: # start api process
	@go run cmd/main.go

test: # run all unit tests	
	@go test ./... -timeout 5s -cover -coverprofile=cover.out	