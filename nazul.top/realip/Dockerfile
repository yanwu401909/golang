FROM golang:1.23 AS realip-builder 

WORKDIR /usr/src/app

# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

COPY . .
RUN go get -v ./...
RUN go build -v -o /usr/local/bin/app ./...


FROM alpine
WORKDIR /go/bin
COPY --from=realip-builder /usr/local/bin/app .
# RUN apt-get update
# RUN apt-get install ca-certificates
ENTRYPOINT ["./app"]