.PHONY: image release
IMG ?= wbuntu/golang
VERSION ?= latest
PLATFORM ?= linux/amd64,linux/arm64
TARGETS ?= 1.23 1.24

image: 
	@echo "Building image $(IMG):$(VERSION)"
	@docker buildx build --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) --load .

release:
	@echo "Building and pushing all platform image $(IMG):$(VERSION)"
	@docker buildx build --platform $(PLATFORM) --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) --push .
	
all: 
	@for var in $(TARGETS);do $(MAKE) VERSION=$$var release;done