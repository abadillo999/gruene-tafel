FROM golang:1.13 AS builder
ENV GOPATH=/code

COPY . /code/
WORKDIR /code
RUN echo $GOPATH
RUN go get "github.com/xeipuuv/gojsonschema" "github.com/gorilla/mux"
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o gruene-worker src/gruene-worker.go
RUN ls

FROM alpine:edge
RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app

COPY --from=builder /code/gruene-worker /app/

WORKDIR /app