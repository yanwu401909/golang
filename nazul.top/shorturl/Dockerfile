FROM golang:1.23 AS short-url-builder 

WORKDIR /usr/src/app
COPY . .
RUN go get -v ./...
RUN go build -v -o /usr/local/bin/app ./...


FROM debian:stretch-slim
COPY --from=short-url-builder /usr/local/bin/app /usr/local/bin 
EXPOSE 9999
ENTRYPOINT ["/usr/local/bin/app"]