.PHONY: run
run:build
	./apollo.exe

.PHONY: build
build: mod
	go build -o apollo.exe main.go

.PHONY: mod
mod:
	git add .
	go mod tidy

.PHONY: migrate
migrate: build
	./apollo.exe migrate

.PHONY: rollback
rollback: build
	./apollo.exe migrate rollbackLast