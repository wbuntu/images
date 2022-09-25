.PHONY: image release
IMG ?= wbuntu/debian
VERSION ?= 11
PLATFORM ?= linux/amd64,linux/arm64

image: 
	@echo "Building image $(IMG):$(VERSION)"
	@docker build --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) .

release:
	@echo "Building and pushing all platform image $(IMG):$(VERSION)"
	@docker buildx build --platform $(PLATFORM) --build-arg VERSION=$(VERSION) -t $(IMG):$(VERSION) --push . 
	
