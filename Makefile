BIN_OUT=bin
SERVE_NAME=glockserve
SERVE=cmd/serve/main.go
SOURCES=$(SERVE)

all: build

build: clean $(SOURCES)
	mkdir -p $(BIN_OUT)
	go build -o $(BIN_OUT)/$(SERVE_NAME) $(SERVE)

dev: $(SERVE)
	go run $(SERVE)

clean:
	if [ -d $(BIN_OUT) ]; then rm -r $(BIN_OUT);	fi
