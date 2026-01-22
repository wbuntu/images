.PHONY: image release
IMG ?= wbuntu/n8n
VERSION ?= 1.101.1
PLATFORM ?= linux/amd64,linux/arm64

image: 
	@echo "Building image $(IMG):$(VERSION)"
	@docker buildx build --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) --load .

release:
	@echo "Building and pushing all platform image $(IMG):$(VERSION)"
	@docker buildx build --platform $(PLATFORM) --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) --push . 
	@docker buildx build --platform $(PLATFORM) --build-arg VERSION=$(VERSION) -t $(IMG):latest --push . 