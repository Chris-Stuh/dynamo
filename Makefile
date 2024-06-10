GO := go
SRC_DIR := ./src

# Variable to hold directories with source code for easy reference
PACKAGES := hashing network node

# Command to build each package, outputting binaries to a 'bin' directory under each package folder
BINARIES := $(addprefix $(SRC_DIR)/,$(addsuffix /bin,$(PACKAGES)))

# Default rule for building all packages
all: $(BINARIES)

# Rule to build each package
$(BINARIES):
	@echo "Building $@"
	@mkdir -p $@
	@$(GO) build -o $@/$(notdir $@) $(patsubst %/bin,%,$@)

# Rule to run tests
test:
	@$(GO) test ./tests

# Rule to clean up built binaries and other artifacts
clean:
	@echo "Cleaning up..."
	@find . -name 'bin' -type d -exec rm -r {} +
	@find . -type f -name '*.test' -delete
	@find . -type f -name '*.out' -delete

.PHONY: all test clean
