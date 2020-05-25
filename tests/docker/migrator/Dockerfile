FROM golang:1.14 as builder
WORKDIR /app
COPY . .
RUN cd /app/calendar/migrations && CGO_ENABLED=0 GOOSE=linux GO111MODULE=on go build -mod=vendor -a -installsuffix nocgo -o migrator *.go && mv migrator ../../


FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/ ./