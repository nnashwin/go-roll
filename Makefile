BIN_NAME="roll"

all: test build move

.PHONY: build
build:
	go build -o $(BIN_NAME)

.PHONY: test
test:
	go test

# Moves to the GOPATH/bin if GOPATH is set
.PHONY: move
move:
	if [ "$(GOPATH)" = "" ] ; \
	then \
		@echo "Your GOPATH is not set.  You must set it to continue"; \
	else \
		mv $(BIN_NAME) $(GOPATH)/bin/; \
	fi;
