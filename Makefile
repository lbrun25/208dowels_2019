##
## EPITECH PROJECT, 2019
## Makefile
## File description:
## Makefile
##

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOENV = $(GOCMD) env
BINARY_NAME = 206neutrinos
BINARY_UNIX = $(BINARY_NAME)_unix

RM = rm -f

export GOPATH := $(shell pwd)

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)


build-linux:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

clean:
	$(GOCLEAN)

fclean: clean
	$(RM) $(BINARY_NAME)
	$(RM) $(BINARY_UNIX)

re: fclean all

.PHONY: all clean fclean re⏎
