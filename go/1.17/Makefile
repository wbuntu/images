.PHONY: image release
GOVERSION ?= 1.17.11
IMG ?= wbuntu/go-1.17
VERSION ?= v1.0
PLATFORM ?= linux/amd64,linux/arm64

image: 
	@echo "Building image $(IMG):$(VERSION)"
	@docker buildx build --build-arg GOVERSION=$(GOVERSION) --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) .

release:
	@echo "Building and pushing all platform image $(IMG):$(VERSION)"
	@docker buildx build --platform $(PLATFORM) --build-arg GOVERSION=$(GOVERSION) --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) . --push
	
