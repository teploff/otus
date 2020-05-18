FROM golang:1.14 as builder
LABEL mainater="Alexander Teplov teploff.aa@gmail.com"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOSE=linux GO111MODULE=on go build -mod=vendor -a -installsuffix nocgo -o scheduler /app/cmd/scheduler/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/ ./

CMD ["./scheduler", "--config=./init/config_stage.yaml"]