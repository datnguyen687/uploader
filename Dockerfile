#### Build stage
FROM amd64/golang:1.17-alpine as builder
WORKDIR /app
RUN apk update && apk add g++ make git

ARG ENV=dev
ENV ENV=${ENV} \
    CGO_ENABLED=1
COPY . .

RUN go mod download
RUN echo "✅ Build for Linux"
RUN make init; make swagger; make mockery; make test; make build

#### Runtime
FROM alpine:3.12

WORKDIR /app
COPY --from=builder /app/bin/ecommerce .
COPY --from=builder /app/docs .

ENTRYPOINT [ "/app/ecommerce" ]
