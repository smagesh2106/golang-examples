# Compiler settings
GO := go
CFLAGS := -buildmode=plugin
SRC_DIR := src/plugins
PLUGINS_DIR := plugins


# Build the main program and plugins
all: clean plugins main

# Build the main program (main.go)
main: main.go
	@echo "Building main program..."
	$(GO) build -o main main.go

# Build the plugins and place them into the plugins directory
plugins: $(SRC_DIR)/plugin1.go $(SRC_DIR)/plugin2.go
	@echo "Building plugins..."

	# Ensure the plugins directory exists
	mkdir -p $(PLUGINS_DIR)

	# Build plugin1 and place it into the plugins directory
	@echo "Building plugin1.so..."
	$(GO) build $(CFLAGS) -o $(PLUGINS_DIR)/plugin1.so $(SRC_DIR)/plugin1.go
	# Check if plugin1.so was created
	@if [ ! -f $(PLUGINS_DIR)/plugin1.so ]; then \
		echo "Error: plugin1.so not generated!"; \
		exit 1; \
	fi

	@echo "Finished Building plugin1.so..."

	# Build plugin2 and place it into the plugins directory
	@echo "Building plugin2.so..."
	$(GO) build $(CFLAGS) -o $(PLUGINS_DIR)/plugin2.so $(SRC_DIR)/plugin2.go
	# Check if plugin2.so was created
	@if [ ! -f $(PLUGINS_DIR)/plugin2.so ]; then \
		echo "Error: plugin2.so not generated!"; \
		exit 1; \
	fi

	@echo "Finished Building plugin2.so..."

# Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -f main $(PLUGINS_DIR)/*.so
	rm -rf $(PLUGINS_DIR)
	@echo "Clean up done."
