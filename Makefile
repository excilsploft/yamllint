app := yamllint
source := $(app).go
test_source := $(app)_test.go
platforms := darwin linux windows
outdir := binaries
zipfiles := $(wildcard *.zip)

default: test build

.PHONY: build
build: $(platforms) $(source)

$(platforms):
	GOOS=$@ GOOARCH=amd64 go build -o $(app) $(source)
	zip '$@-amd64-$(app).zip' $(app)
	rm $(app)

.PHONY: test
test: $(test_source)
	go test -v 

.PHONY: install
install: $(source)
	@go install


.PHONY: clean
clean: $(zipfiles)
	rm  $(zipfiles)
