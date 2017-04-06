# CMD
BIN_OUT=bin

## serve
SERVE_NAME=serve
SERVE_PATH=cmd/serve/main.go
SERVE_BIN_PATH=$(BIN_OUT)/$(SERVE_NAME)


all: build


dev: $(SERVE_PATH)
	go run $(SERVE_PATH)


clean:
	if [ -d $(BIN_OUT) ]; then rm -r $(BIN_OUT);	fi


build-serve:
	mkdir -p $(BIN_OUT)
	if [ -f $(SERVE_BIN_PATH) ]; then rm $(SERVE_BIN_PATH);	fi
	CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -s' -o $(SERVE_BIN_PATH) $(SERVE_PATH)


build: clean build-serve


run: build
	./bin/$(SERVE_NAME)
