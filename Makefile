# Makefileの変数定義
PROTO_FILE := sample.proto
GO_OUT_DIR := .
GO_GRPC_OUT_DIR := .

# デフォルトのターゲット
.PHONY: all
all: generate

# コード生成ターゲット
.PHONY: generate
generate:
	cd proto && \
	protoc --go_out=$(GO_OUT_DIR) --go-grpc_out=require_unimplemented_servers=false:$(GO_GRPC_OUT_DIR) $(PROTO_FILE)

# クリーンターゲット
.PHONY: clean
clean:
	rm -f *.pb.go
