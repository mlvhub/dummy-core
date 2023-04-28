.PHONY: test
test:
	go test -v -cover -coverpkg=./core/... ./... -count=1
