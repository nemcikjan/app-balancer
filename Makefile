COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z)
BUILDENVVAR=CGO_ENABLED=0

.PHONY: build-app-balancer.amd64
build-app-balancer.amd64:
	$(COMMONENVVAR) $(BUILDENVVAR) GOARCH=amd64 go build -ldflags '-w' -o bin/app-balancer main.go

.PHONY: build-app-balancer.arm64v8
build-app-balancer.arm64v8:
	$(COMMONENVVAR) $(BUILDENVVAR) GOARCH=arm64 go build -ldflags '-w' -o bin/app-balancer main.go