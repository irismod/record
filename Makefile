PACKAGES_UNITTEST=$(shell go list ./... | grep -v '/simulation')

test-unit:
	go test -mod=readonly -tags='ledger test_ledger_mock' ${PACKAGES_UNITTEST}

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs goimports -w -local github.com/irismod/record
