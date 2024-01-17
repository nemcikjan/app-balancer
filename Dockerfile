ARG ARCH
FROM golang:1.20

WORKDIR /app
COPY . .
ARG ARCH

RUN make build-app-balancer.$ARCH

FROM $ARCH/alpine:3.19

COPY --from=0 /app/bin/app-balancer /bin/app-balancer

WORKDIR /bin
CMD ["app-balancer"]