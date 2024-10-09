FROM golang:1.23 AS monitor-builder 

WORKDIR /usr/src/app

# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

COPY . .
RUN go get -v .
RUN go build -v -o /usr/local/bin/app ./...


FROM debian:sid-slim
COPY --from=monitor-builder /usr/local/bin/app /usr/local/bin 
ENTRYPOINT ["/usr/local/bin/app"]