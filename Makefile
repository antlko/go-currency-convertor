build:
	go build -o go-currency-convertor cmd/service/main.go

start: build
	./go-currency-convertor