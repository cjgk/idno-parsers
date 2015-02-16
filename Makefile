.PHONY: test

test:
	@rm -f /tmp/idno-parsers-coverprofile.out
	@go test -v -coverprofile=/tmp/idno-parsers-coverprofile.out
	@go tool cover -func=/tmp/idno-parsers-coverprofile.out
	@rm -f /tmp/idno-parsers-coverprofile.out
