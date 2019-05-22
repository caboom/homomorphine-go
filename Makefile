# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
CGO=$(GOCMD) tool cgo
SRCDIR=pkg

all: test build
build:
	$(GOBUILD) $(SRCDIR)/*.go
test: 
	$(GOTEST) -v ./$(SRCDIR)/...
clean: 
	$(GOCLEAN) $(SRCDIR)/*.go
install:
	$(GOINSTALL) $(SRCDIR)/*.go
