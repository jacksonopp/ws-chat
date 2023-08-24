GOVERSION := $(shell go version | cut -d ' ' -f 3 | cut -d '.' -f 2)

.PHONY: build ui
.DEFAULT_GOAL := build
PROTON_COMMIT := "d6e14c68d5a661d2a517613de560342c1cf2cea6"

ui:
	@echo " > generating ui build"
	@cd ui && $(MAKE) 

build: ui
	CGO_ENABLED=0 go build -o bin/ws-chat .

serve-prod: ui build
	./bin/ws-chat
.PHONY: serve-prod

dev:
	make -j serve-be serve-fe
.PHONY: dev

serve-be:
	CGO_ENABLED=0 go run main.go
.PHONY: serve-be

serve-fe:
	@cd ui && npm run dev
.PHONY: serve-fe