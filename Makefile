
AUTH_DIR = cmd/auth
COUNT_DIR = cmd/count
HELLO_DIR = cmd/hello
QUERY_DIR = cmd/query

# Цель для запуска всех сервисов
.PHONY: all
all: run-auth run-count run-hello run-query

# Цели для запуска каждого сервиса в фоне
.PHONY: run-auth
run-auth:
	@echo "Running auth service..."
	go run $(AUTH_DIR)/main.go &

.PHONY: run-count
run-count:
	@echo "Running count service..."
	go run $(COUNT_DIR)/main.go &

.PHONY: run-hello
run-hello:
	@echo "Running hello service..."
	go run $(HELLO_DIR)/main.go &

.PHONY: run-query
run-query:
	@echo "Running query service..."
	go run $(QUERY_DIR)/main.go &