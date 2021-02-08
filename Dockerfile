FROM golang:1.15
COPY . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build /app/cmd/web
ENTRYPOINT ["./web"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./web"]
EXPOSE 4001