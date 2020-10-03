.PHONY: build
build:
	git pull
	cd cmd/ && CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo && cd ../ && ./cmd/cmd