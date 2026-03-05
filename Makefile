build: # сборка утилиты
	go build -o bin/gendiff ./cmd/gendiff

lint: # проверка кода линтером golangci-lint
	golangci-lint run
	
test: # запуск тестов
	go test -v ./...
