PROJECT?=github.com/toschneck/go-cds2018
BUILD_PATH?=cmd/cdays
RELEASE?=0.1.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOOS?=linux
GOARCH?=amd64
APP?=cds2018
PORT?=8000
DIAGNOSE_PORT=8088

REGISTRY?=docker.io/toschneck
NAMESPACE?=tobi
C_NAME=${NAMESPACE}-${APP}
C_IMAGE=${REGISTRY}/${C_NAME}

build: clean test
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
		go build \
	  	-ldflags "-s -w -X ${PROJECT}/internal/version.Release=${RELEASE} \
	  	-X ${PROJECT}/internal/version.Commit=${COMMIT} \
	  	-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
	  	-o ./bin/${GOOS}-${GOARCH}/${APP} ${PROJECT}/${BUILD_PATH}

image: build
	docker build -t ${C_IMAGE}:${RELEASE} .

run: image
	docker run --name ${APP} -p ${PORT}:${PORT} -p ${DIAGNOSE_PORT}:${DIAGNOSE_PORT} --rm \
		-e "PORT=${PORT}" -e "DIAGNOSE_PORT=${DIAGNOSE_PORT}" \
		${C_IMAGE}:${RELEASE}

test:
	go test -race ./...

clean:
	rm -rf ./bin
