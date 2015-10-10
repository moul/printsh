NAME = printsh

#TESTCMD ?= for i in {0..6}; do date; sleep .2; date >&2; sleep .2; done
TESTCMD ?= test-streams


$(NAME): $(shell find . -name "*.go")
	go build -o $@ ./cmd/$(NAME)


.PHONY: test
test: test1


.PHONY: test1
test1: $(NAME)
	$(TESTCMD) | ./$(NAME)


.PHONY: test2
test2: $(NAME)
	$(NAME) sh -c "$(TESTCMD)"
