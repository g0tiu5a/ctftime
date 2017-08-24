TEST=go test $(TEST_OPTS) .
TEST_OPTS=-v

TEST_PKGS=common event

.PHONY: test

