.PHONY: image release all
IMG ?= wbuntu/alpine
VERSION ?= latest
PLATFORM ?= linux/amd64,linux/arm64
TARGETS ?= 3.18 3.19

image: 
	@echo "Building image $(IMG):$(VERSION)"
	@docker buildx build --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) --load .

release:
	@echo "Building and pushing all platform image $(IMG):$(VERSION)"
	@docker buildx build --platform $(PLATFORM) --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) --push . 

all: 
	@for var in $(TARGETS);do $(MAKE) VERSION=$$var release;done