# The binaries to build (just the basenames).
BINS := apg

VERSION ?= $(shell grep VersionString apg.go | head -1 | awk '{print $$5}' | sed 's/"//g')
ALL_PLATFORMS := darwin/amd64 darwin/arm64                      \
    linux/amd64 linux/arm linux/arm64 linux/ppc64le linux/s390x \
    freebsd/386 freebsd/amd64 freebsd/arm                       \
    openbsd/386 openbsd/amd64 openbsd/arm                       \
    netbsd/386 netbsd/amd64 netbsd/arm                          \
    windows/386 windows/amd64
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
TAG := $(VERSION)__$(OS)_$(ARCH)

BIN_EXTENSION :=
ifeq ($(OS), windows)
  BIN_EXTENSION := .exe
endif

# If you want to build all binaries, see the 'all-build' rule.
all: # @HELP builds binaries for one platform ($OS/$ARCH)
all: build

# For the following OS/ARCH expansions, we transform OS/ARCH into OS_ARCH
# because make pattern rules don't match with embedded '/' characters.

build-%:
	@$(MAKE) build                        \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

all-build: # @HELP builds binaries for all platforms
all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS)))
OUTBINS = $(foreach bin,$(BINS),bin/v$(VERSION)/$(OS)_$(ARCH)/$(bin)$(BIN_EXTENSION))
build: $(OUTBINS)
BUILD_DIRS := bin/v$(VERSION)/$(OS)_$(ARCH)     \
              bin/zip-files/v$(VERSION)         \
              .go/bin/v$(VERSION)/$(OS)_$(ARCH) \
              .go/cache

# Each outbin target is just a facade for the respective stampfile target.
# This `eval` establishes the dependencies for each.
$(foreach outbin,$(OUTBINS),$(eval  \
    $(outbin): .go/$(outbin).stamp  \
))
# This is the target definition for all outbins.
$(OUTBINS):
	@true

# Each stampfile target can reference an $(OUTBIN) variable.
$(foreach outbin,$(OUTBINS),$(eval $(strip   \
    .go/$(outbin).stamp: OUTBIN = $(outbin)  \
)))
# This is the target definition for all stampfiles.
# This will build the binary under ./.go and update the real binary iff needed.
STAMPS = $(foreach outbin,$(OUTBINS),.go/$(outbin).stamp)
.PHONY: $(STAMPS)
$(STAMPS): go-build
	@echo "binary: $(OUTBIN)"
	@if ! cmp -s .go/$(OUTBIN) $(OUTBIN); then  \
	    mv .go/$(OUTBIN) $(OUTBIN);             \
	    zip -9 ./bin/zip-files/v$(VERSION)/apg_v$(VERSION)_$(OS)_$(ARCH).zip $(OUTBIN) \
	    date >$@;                               \
	fi

# This runs the actual `go build` which updates all binaries.
go-build: $(BUILD_DIRS)
	@echo
	@echo "building for $(OS)/$(ARCH)"
	@go build -o .go/bin/v$(VERSION)/$(OS)_$(ARCH)/apg$(BIN_EXTENSION) apg.go

version: # @HELP outputs the version string
version:
	@echo $(VERSION)

	    "

$(BUILD_DIRS):
	@mkdir -p $@

clean: # @HELP removes built binaries and temporary files
clean: bin-clean

bin-clean:
	rm -rf .go bin

help: # @HELP prints this message
help:
	@echo "VARIABLES:"
	@echo "  BINS = $(BINS)"
	@echo "  OS = $(OS)"
	@echo "  ARCH = $(ARCH)"
	@echo "  REGISTRY = $(REGISTRY)"
	@echo
	@echo "TARGETS:"
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)    \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '
