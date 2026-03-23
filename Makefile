BINARY = tess
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DIR = dist

PLATFORMS = \
	darwin/amd64/macos-intel \
	darwin/arm64/macos-apple-silicon \
	linux/amd64/linux-x64 \
	linux/arm64/linux-arm64 \
	windows/amd64/windows-x64

LDFLAGS = -s -w -X main.version=$(VERSION)

.PHONY: build clean all release

# Build for current platform
build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/tess

# Build for all platforms
all: clean
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		os=$$(echo $$platform | cut -d/ -f1); \
		arch=$$(echo $$platform | cut -d/ -f2); \
		label=$$(echo $$platform | cut -d/ -f3); \
		dir=$(BUILD_DIR)/$(BINARY)-$${label}; \
		mkdir -p $$dir; \
		binary=$(BINARY); \
		if [ "$$os" = "windows" ]; then binary=$(BINARY).exe; fi; \
		echo "Building $$label..."; \
		GOOS=$$os GOARCH=$$arch go build -ldflags "$(LDFLAGS)" -o $$dir/$$binary ./cmd/tess || exit 1; \
	done
	@echo "Done. Binaries in $(BUILD_DIR)/"

# Build and create release archives
release: all
	@cd $(BUILD_DIR) && for dir in $(BINARY)-*; do \
		if [ -f "$$dir/$(BINARY).exe" ]; then \
			(cd $$dir && zip ../$$dir.zip $(BINARY).exe); \
		else \
			(cd $$dir && tar czf ../$$dir.tar.gz $(BINARY)); \
		fi \
	done
	@echo "Release archives in $(BUILD_DIR)/"

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(BINARY)
