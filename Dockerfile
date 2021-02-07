FROM golang:1.15
ADD . /app
WORKDIR /app
RUN go build -o snippetbox cmd/web/*
ENTRYPOINT /app/snippetbox
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=0 /app/snippetbox .
CMD ["./snippetbox"]
