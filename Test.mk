include .env

ENV ?= dev

run:
	@if [ "$(ENV)" = "dev" ]; then \
		echo "dev mode"; \
	else \
		echo "prod mode"; \
	fi