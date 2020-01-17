source = *.go
output := yamllint
test_source := *_test.go

default: test build

build: $(source)
	go build -o $(output) $(source)

.PHONY: test
test: $(test_source)
	go test

.PHONY: install
install: $(soruce)
	go install 

clean: $(output)
	@rm $(output)

