# Directory containing the proto files
PROTO_DIR := api
# Go output directory (usually in your Go workspace or project)
OUT_DIR := .

# The protobuf compiler (protoc)
PROTOC := protoc

# The Go plugin for protoc
GO_PLUGIN := protoc-gen-go
# The Go gRPC plugin for protoc
GRPC_PLUGIN := protoc-gen-go-grpc

# Add necessary flags for protoc
IMPORT_PATHS := $(shell find $(PROTO_DIR) -type d -exec echo -I {} \;)
GO_FLAGS := --go_out=$(OUT_DIR) --go_opt=paths=source_relative $(IMPORT_PATH)
GRPC_FLAGS := --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative $(IMPORT_PATH)

# Find all .proto files in the PROTO_DIR
PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")

# Default target to generate Go files and gRPC files
protos: $(PROTO_FILES:.proto=.pb.go) $(PROTO_FILES:.proto=.pb.go.grpc.go)

# Rule to generate Go files for each .proto file
%.pb.go: %.proto
	@$(PROTOC) $(GO_FLAGS) $<

# Rule to generate gRPC code for each .proto file
%.pb.go.grpc.go: %.proto
	@$(PROTOC) $(GRPC_FLAGS) $<

