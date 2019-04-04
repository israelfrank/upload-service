
#build stage
FROM golang:alpine AS builder
ENV GO111MODULE=on
WORKDIR /go/src/app
RUN apk add --no-cache git make protobuf
RUN go get -u github.com/golang/protobuf/protoc-gen-go
COPY . .
RUN make build

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app/upload-service /upload-service
ENTRYPOINT ./upload-service
LABEL Name=upload-service Version=0.0.1
EXPOSE 8080
