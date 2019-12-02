FROM golang:1.13 AS builder
ENV GOPATH=/code

COPY . /code/
WORKDIR /code
RUN echo $GOPATH
RUN go get "github.com/xeipuuv/gojsonschema" "github.com/gorilla/mux"
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o gruene-worker src/gruene-worker.go

FROM fiji/fiji:fiji-oracle-jdk6

COPY --from=builder /code/gruene-worker /app/

WORKDIR /app