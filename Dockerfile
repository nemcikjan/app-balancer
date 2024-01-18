FROM golang:1.20

WORKDIR /app
COPY . .
ARG TARGETOS
ARG TARGETARCH

RUN make build-app-balancer.$TARGETARCH

FROM alpine:3.19

COPY --from=0 /app/bin/app-balancer /bin/app-balancer

WORKDIR /bin
CMD ["app-balancer"]