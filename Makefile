PACKAGE_NAME = intel8080
PACKAGE      = github.com/magarcia/$(PACKAGE_NAME)

GO      = go
GODOC   = godoc
GOFMT   = gofmt

Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m>\033[0m")


.PHONY: all
all: disassembler hexdump ; $(info $(M) building executables…) @ ## Build program binaries
#	$(GO) build -o bin/$(PACKAGE_NAME) main.go

.PHONY: disassembler
disassembler: cmd/disassembler.go ; $(info $(M) building disassembler) @ ## Build disassembler
	$Q $(GO) build -o ./bin/disassembler cmd/disassembler.go

.PHONY: hexdump
hexdump: cmd/hexdump.go ; $(info $(M) building hexdump) @ ## Build hexdump
	$Q $(GO) build -o ./bin/hexdump cmd/hexdump.go

.PHONY: fmt
fmt: ; $(info $(M) formatting files) @ ## Format .go files
	$Q $(GOFMT) -l -w ./**/*.go

.PHONY: clean
clean: ; $(info $(M) cleanning) @ ## Clean builds
	$Q rm -rf ./bin

.PHONY: test
test: ; $(info $(M) running tests) @ ## Run tests
	$Q echo "\tNot implemented yet"

.PHONY: lint
lint: ; $(info $(M) running lint) @ ## Run lint
	$Q echo "\tNot implemented yet"

.PHONY: vendor
vendor: ; $(info $(M) installing dependencies) @ ## Install dependencies
	$Q dep ensure

.PHONY: help
help: ; @ ## Show this help
	$Q grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
