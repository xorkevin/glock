# METADATA
VERSION=v0.1.0
PORT=8080
BASEDIR=public


# CMD
BIN_OUT=bin

## serve
SERVE_NAME=serve
SERVE_PATH=cmd/serve/main.go
SERVE_BIN_PATH=$(BIN_OUT)/$(SERVE_NAME)

## fsserve
## serve
FSSERVE_NAME=fsserve
FSSERVE_PATH=cmd/fsserve/main.go
FSSERVE_BIN_PATH=$(BIN_OUT)/$(FSSERVE_NAME)


# DOCKER
SERVE_IMAGE_NAME=glockserver
SERVE_CONTAINER_NAME=sglock


# DEV_POSTGRES
POSTGRES_VOLUME=glockpgvol
POSTGRES_CONTAINER=glockpg
POSTGRES_PASS=admin


.RECIPEPREFIX +=


all: build


dev:
  VERSION=$(VERSION) MODE=DEBUG go run $(SERVE_PATH)


dev-fsserve:
  BASEDIR=$(BASEDIR) go run $(FSSERVE_PATH)


clean:
  if [ -d $(BIN_OUT) ]; then rm -r $(BIN_OUT); fi


build-serve:
  mkdir -p $(BIN_OUT)
  if [ -f $(SERVE_BIN_PATH) ]; then rm $(SERVE_BIN_PATH);	fi
  CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -s' -o $(SERVE_BIN_PATH) $(SERVE_PATH)


build-fsserve:
  mkdir -p $(BIN_OUT)
  if [ -f $(FSSERVE_BIN_PATH) ]; then rm $(FSSERVE_BIN_PATH);	fi
  CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -s' -o $(FSSERVE_BIN_PATH) $(FSSERVE_PATH)


build: clean build-serve build-fsserve


## docker
docker-build: build
  docker build -t $(SERVE_IMAGE_NAME):$(VERSION) .
  docker build -t $(SERVE_IMAGE_NAME) .


docker-run:
  docker run -it --rm --name $(SERVE_CONTAINER_NAME) -e VERSION=$(VERSION) -e MODE=INFO -p $(PORT):$(PORT) $(SERVE_IMAGE_NAME)


docker: docker-build docker-run


## postgres
pg-setup:
  docker volume create --name $(POSTGRES_VOLUME)


pg-run:
  docker run -it --rm --name $(POSTGRES_CONTAINER) -p 5432:5432 -v $(POSTGRES_VOLUME):/var/lib/postgresql/data -e POSTGRES_PASSWORD=$(POSTGRES_PASS) postgres:alpine
