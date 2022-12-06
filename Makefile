JAGW_PROTO_DIR ?= $(JAGW_DIR)/proto
JAGW_GO_PB_DIR = ./jagw

.PHONY: help

help:
	@echo "Env Variables:"
	@echo "	JAGW_DIR: $(JAGW_DIR)"
	@echo "	JAGW_PROTO_DIR: $(JAGW_PROTO_DIR)"	
	@echo ""
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  	help: Show this help"
	@echo "  	proto-gen: Generate proto files"

proto-gen:
	mkdir -p $(JAGW_GO_PB_DIR)
	protoc --proto_path=$(JAGW_PROTO_DIR) \
	--go_out=$(JAGW_GO_PB_DIR) \
	--go_opt=module=github.com/jalapeno-api-gateway/jagw-go \
	--go-grpc_out=$(JAGW_GO_PB_DIR) \
	--go-grpc_opt=module=github.com/jalapeno-api-gateway/jagw-go $(JAGW_PROTO_DIR)/**/*.proto