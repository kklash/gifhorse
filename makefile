ARCH := amd64
BINARY := gifhorse
PLATFORMS := darwin linux
VERSION ?= dev
DIST_DIR := dist
RELEASE_DIR := $(DIST_DIR)/$(VERSION)
SHASUMS := $(RELEASE_DIR)/SHA256SUMS.txt

os = $(word 1, $@)
output = $(BINARY)-$(os)_$(VERSION)_$(ARCH)
outdir = $(RELEASE_DIR)/$(output)

.PHONY: $(PLATFORMS) all

all: $(PLATFORMS)
	rm -f $(SHASUMS)
	find $(RELEASE_DIR) -name *.tar.gz -exec shasum -a 256 {} >>  $(SHASUMS) \;
	cat $(SHASUMS) | gpg --armor --batch --yes --output $(SHASUMS) --clear-sign

$(PLATFORMS):
	mkdir -p $(outdir)
	GOOS=$@ GOARCH=$(ARCH) go build -o $(outdir)/$(BINARY)
	tar -C $(RELEASE_DIR) -czf $(outdir).tar.gz $(output)
	rm -r $(outdir)

clean: dist
	rm -rf dist
