.PHONY: test
test:
	go test -v -cover -coverpkg=./core/... ./... -count=1

.PHONY: version-check
version-check:
	./scripts/version-up.sh --patch

.PHONY: version-apply
version-apply:
	./scripts/version-up.sh --patch --apply
