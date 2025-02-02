export LANG = en_US.UTF-8

PLATFORMS := linux/amd64 windows/amd64 darwin/arm64
CURRENT_PLARTFORM := $(shell go env GOOS)/$(shell go env GOARCH)

OUTPUT_DIR := build
APP_NAME := gomic

# build for current platform
.PHONY: build
build: $(CURRENT_PLARTFORM)

# build for all platforms
all: $(PLATFORMS)
	@echo "Build completed for all platforms."

# build for single platform
$(PLATFORMS):
	@mkdir -p $(OUTPUT_DIR)/$@
	@GOOS=$(word 1, $(subst /, ,$@)) GOARCH=$(word 2, $(subst /, ,$@)) \
		./scripts/build.sh $(word 1, $(subst /, ,$@)) $(word 2, $(subst /, ,$@)) \
		$(OUTPUT_DIR)/$@ $(APP_NAME)

# clean all build files
clean:
	@echo "Cleaning build files..."
	rm -rf $(OUTPUT_DIR)/*
	rm -rf assets/web/*
	@echo "All build files cleaned."

run:
	$(OUTPUT_DIR)/$(CURRENT_PLARTFORM)/$(APP_NAME)

dev-backend:
	go run .

dev-frontend:
	cd web
	npm run dev