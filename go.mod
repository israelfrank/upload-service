module github.com/meateam/upload-service

go 1.12

require (
	github.com/aws/aws-sdk-go v1.19.22
	github.com/golang/protobuf v1.2.0
	go.elastic.co/apm/module/apmgrpc v1.3.0
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.20.1
)

replace go.elastic.co/apm/module/apmgrpc => github.com/omrishtam/apm-agent-go/module/apmgrpc v1.3.1-0.20190514172539-1b2e35db8668
