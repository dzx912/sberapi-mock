GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=sberapi-mock
VERSION?=0.0.0


build: 
	mkdir -p out/bin
	$(GOCMD) build -o out/bin/$(BINARY_NAME) .

clean:
	rm -fr ./bin
	rm -fr ./out


run-server:
	$(GOCMD) run *.go start