# Makefile

# Define the path for the public suffix list
PUBLIC_SUFFIX_LIST := ./data/public_suffix_list.dat
GENERATED_GO_FILE := ./parser/public_suffix_list.go

# The official URL for the public suffix list
URL := https://publicsuffix.org/list/public_suffix_list.dat

# Default target when running `make`.
# This will first run `clean` and then `generate`.
.PHONY: all
all: clean generate

# The 'generate' target creates the Go source file from the suffix list.
.PHONY: generate
generate: sync
	@echo "Generating parser/public_suffix_list.go..."
	@go run ./cmd/generate/main.go
	@echo "Generation complete."

# The 'sync' target ensures the public suffix list is downloaded.
# This is a phony target that depends on the actual file target.
.PHONY: sync
sync: $(PUBLIC_SUFFIX_LIST)

# Rule to download the public suffix list file.
# This rule runs only if the file does not exist.
$(PUBLIC_SUFFIX_LIST): ./data/
	@echo "Downloading public_suffix_list.dat..."
	@curl -sfL -o $(PUBLIC_SUFFIX_LIST) $(URL)
	@echo "Download complete."

# Rule to create the data directory if it doesn't exist.
./data/:
	@mkdir -p ./data/

# The 'clean' target removes the downloaded and generated files.
.PHONY: clean
clean:
	@echo "Cleaning up downloaded and generated files..."
	@rm -f $(PUBLIC_SUFFIX_LIST) $(GENERATED_GO_FILE)
