.DEFAULT_GOAL: build
.PHONY: build
build:
	@docker build . -t "deciphernow/parrot-service:latest"