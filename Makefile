NAME ?= isaac
VERSION ?= 0.0.0
LOG_LEVEL ?= debug

.PHONY: build run version test coverage

all: clean build run

build:
	@echo "Building app $(NAME) ..."
	@go build -ldflags="-X 'main.BuildVersion=$(VERSION)' -X 'main.BuildName=$(NAME)'" -o $(NAME) .

run: $(NAME)
	@echo "Starting app $(NAME)"
	@./$(NAME)

dev:
	@echo "Starting app $(NAME) in dev mode"
	@go run -ldflags="-X 'main.BuildVersion=$(VERSION)' -X 'main.BuildName=$(NAME)'" main.go

version:
	@echo $(VERSION)

clean:
	rm -f $(NAME)

test:
	@echo "Running unit tests .."
	@go install github.com/jstemmer/go-junit-report
	@go test ./... -v -cover -coverprofile=c.out 2>&1 | go-junit-report > report.xml

coverage:
	@echo "Coverage .."
	@go tool cover -html=c.out -o coverage.html
	@go install github.com/t-yuki/gocover-cobertura
	@gocover-cobertura < c.out > coverage.xml
