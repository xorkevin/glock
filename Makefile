# CMD
BIN_OUT=bin

## serve
SERVE_NAME=serve
SERVE_PATH=cmd/serve/main.go
SERVE_BIN_PATH=$(BIN_OUT)/$(SERVE_NAME)

PORT=8080

SERVE_IMAGE_NAME=glockserver
SERVE_CONTAINER_NAME=sglock
VERSION=v1


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


run:
	./$(SERVE_BIN_PATH)


## docker
build-docker: build
	docker build -t $(SERVE_IMAGE_NAME):$(VERSION) .
	docker build -t $(SERVE_IMAGE_NAME) .


run-docker:
	docker run -it --rm --name $(SERVE_CONTAINER_NAME) -p $(PORT):$(PORT) $(SERVE_IMAGE_NAME)


docker: build-docker run-docker
