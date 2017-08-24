TEST=go test $(TEST_OPTS) .
TEST_OPTS=-v

TEST_PKGS=common event

.PHONY: test

test:
	@for pkg in $(TEST_PKGS); do\
		(\
			cd $$pkg && \
				$(TEST)\
		);\
	done
