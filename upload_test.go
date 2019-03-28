package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"reflect"
	"testing"
	"time"
	pb "upload-service/proto"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// Declaring global variables.
var s3Endpoint string
var newSession = session.Must(session.NewSession())
var s3Client *s3.S3
var lis *bufconn.Listener

func init() {
	// Fetch env vars
	s3AccessKey := os.Getenv("S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Endpoint = os.Getenv("S3_ENDPOINT")
	s3Token := ""

	// Configure to use S3 Server
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3AccessKey, s3SecretKey, s3Token),
		Endpoint:         aws.String(s3Endpoint),
		Region:           aws.String("eu-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	// Init real client.
	newSession = session.New(s3Config)
	s3Client = s3.New(newSession)

	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(10 << 20))
	server := &UploadHandler{UploadService: UploadService{s3Client: s3Client}}
	pb.RegisterUploadServer(grpcServer, server)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestUploadService_UploadFile(t *testing.T) {
	metadata := make(map[string]*string)
	metadata["test"] = aws.String("testt")
	type fields struct {
		s3Client *s3.S3
	}
	type args struct {
		file     io.Reader
		key      *string
		bucket   *string
		metadata map[string]*string
		ctx      context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		{
			name:   "upload text file",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   aws.String("testbucket"),
				file:     bytes.NewReader([]byte("Hello, World!")),
				metadata: nil,
				ctx:      context.Background(),
			},
			wantErr: false,
			want:    aws.String(fmt.Sprintf("%s/testbucket/testfile.txt", s3Endpoint)),
		},
		{
			name:   "upload text file in a folder",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfolder/testfile.txt"),
				bucket:   aws.String("testbucket"),
				file:     bytes.NewReader([]byte("Hello, World!")),
				metadata: metadata,
				ctx:      context.Background(),
			},
			wantErr: false,
			want:    aws.String(fmt.Sprintf("%s/testbucket/testfolder/testfile.txt", s3Endpoint)),
		},
		{
			name:   "upload text file with empty key",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String(""),
				bucket:   aws.String("testbucket"),
				file:     bytes.NewReader([]byte("Hello, World!")),
				metadata: nil,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload text file with empty bucket",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   aws.String(""),
				file:     bytes.NewReader([]byte("Hello, World!")),
				metadata: nil,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload text file with nil key",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      nil,
				bucket:   aws.String("testbucket"),
				file:     bytes.NewReader([]byte("Hello, World!")),
				metadata: nil,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload text file with nil bucket",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   nil,
				file:     bytes.NewReader([]byte("Hello, World!")),
				metadata: nil,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload nil file",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   aws.String("testbucket"),
				file:     nil,
				metadata: nil,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UploadService{
				s3Client: tt.fields.s3Client,
			}

			got, err := s.UploadFile(tt.args.ctx, tt.args.file, tt.args.key, tt.args.bucket, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadService.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && *got != *tt.want {
				t.Errorf("UploadService.UploadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUploadHandler_UploadMedia(t *testing.T) {
	hugefile := make([]byte, 5<<20)
	rand.Read(hugefile)

	uploadservice := UploadService{
		s3Client: s3Client,
	}
	type fields struct {
		UploadService UploadService
	}
	type args struct {
		ctx     context.Context
		request *pb.UploadMediaRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UploadMediaResponse
		wantErr bool
	}{
		{
			name:   "UploadMedia - text file",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMediaRequest{
					Key:    "testfile.txt",
					Bucket: "testbucket",
					File:   []byte("Hello, World!"),
				},
			},
			wantErr: false,
			want: &pb.UploadMediaResponse{
				Output: fmt.Sprintf("%s/testbucket/testfile.txt", s3Endpoint),
			},
		},
		{
			name:   "UploadMedia - text file - without key",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMediaRequest{
					Key:    "",
					Bucket: "testbucket",
					File:   []byte("Hello, World!"),
				},
			},
			wantErr: true,
		},
		{
			name:   "UploadMedia - text file - without bucket",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMediaRequest{
					Key:    "testfile.txt",
					Bucket: "",
					File:   []byte("Hello, World!"),
				},
			},
			wantErr: true,
		},
		{
			name:   "UploadMedia - text file - with nil file",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMediaRequest{
					Key:    "testfile.txt",
					Bucket: "testbucket",
					File:   nil,
				},
			},
			wantErr: false,
			want: &pb.UploadMediaResponse{
				Output: fmt.Sprintf("%s/testbucket/testfile.txt", s3Endpoint),
			},
		},
		{
			name:   "UploadMedia - text file - huge file",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMediaRequest{
					Key:    "testfile.txt",
					Bucket: "testbucket",
					File:   hugefile,
				},
			},
			wantErr: false,
			want: &pb.UploadMediaResponse{
				Output: fmt.Sprintf("%s/testbucket/testfile.txt", s3Endpoint),
			},
		},
	}

	// Create connection to server
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewUploadClient(conn)

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.UploadMedia(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadHandler.UploadMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadHandler.UploadMedia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUploadService_UploadInit(t *testing.T) {
	metadata := make(map[string]*string)
	metadata["test"] = aws.String("testt")
	type fields struct {
		s3Client *s3.S3
	}
	type args struct {
		key      *string
		bucket   *string
		metadata map[string]*string
		ctx      context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *s3.CreateMultipartUploadOutput
		wantErr bool
	}{
		{
			name:   "init upload",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   aws.String("testbucket"),
				metadata: metadata,
				ctx:      context.Background(),
			},
			wantErr: false,
			want: &s3.CreateMultipartUploadOutput{
				Bucket: aws.String("testbucket"),
				Key:    aws.String("testfile.txt"),
			},
		},
		{
			name:   "init upload in folder",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfolder/testfile.txt"),
				bucket:   aws.String("testbucket"),
				metadata: metadata,
				ctx:      context.Background(),
			},
			wantErr: false,
			want: &s3.CreateMultipartUploadOutput{
				Bucket: aws.String("testbucket"),
				Key:    aws.String("testfolder/testfile.txt"),
			},
		},
		{
			name:   "init upload with missing key",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String(""),
				bucket:   aws.String("testbucket"),
				metadata: metadata,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "init upload with nil key",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      nil,
				bucket:   aws.String("testbucket"),
				metadata: metadata,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "init upload with missing bucket",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   aws.String(""),
				metadata: metadata,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "init upload with nil bucket",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   nil,
				metadata: metadata,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "init upload with empty metadata",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   aws.String("testbucket"),
				metadata: aws.StringMap(make(map[string]string)),
				ctx:      context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "init upload with nil metadata",
			fields: fields{s3Client: s3Client},
			args: args{
				key:      aws.String("testfile.txt"),
				bucket:   aws.String("testbucket"),
				metadata: nil,
				ctx:      context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UploadService{
				s3Client: tt.fields.s3Client,
			}
			got, err := s.UploadInit(tt.args.ctx, tt.args.key, tt.args.bucket, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadService.UploadInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && got.UploadId == nil {
				t.Errorf("UploadService.UploadInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUploadHandler_UploadInit(t *testing.T) {
	metadata := make(map[string]string)
	metadata["test"] = "testt"
	uploadservice := UploadService{
		s3Client: s3Client,
	}
	type fields struct {
		UploadService UploadService
	}
	type args struct {
		ctx     context.Context
		request *pb.UploadInitRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UploadInitResponse
		wantErr bool
	}{
		{
			name:   "UploadInit",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadInitRequest{
					Key:      "testfile.txt",
					Bucket:   "testbucket",
					Metadata: metadata,
				},
			},
			wantErr: false,
			want: &pb.UploadInitResponse{
				Key:    "testfile.txt",
				Bucket: "testbucket",
			},
		},
		{
			name:   "UploadInit folder",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadInitRequest{
					Key:      "testfolder/testfile.txt",
					Bucket:   "testbucket",
					Metadata: metadata,
				},
			},
			wantErr: false,
			want: &pb.UploadInitResponse{
				Key:    "testfolder/testfile.txt",
				Bucket: "testbucket",
			},
		},
		{
			name:   "UploadInit with empty key",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadInitRequest{
					Key:      "",
					Bucket:   "testbucket",
					Metadata: metadata,
				},
			},
			wantErr: true,
		},
		{
			name:   "UploadInit with empty bucket",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadInitRequest{
					Key:      "testfile.txt",
					Bucket:   "",
					Metadata: metadata,
				},
			},
			wantErr: true,
		},
		{
			name:   "UploadInit with empty metadata",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadInitRequest{
					Key:      "testfile.txt",
					Bucket:   "testbucket",
					Metadata: make(map[string]string),
				},
			},
			wantErr: true,
		},
		{
			name:   "UploadInit with nil metadata",
			fields: fields{UploadService: uploadservice},
			args: args{
				ctx: context.Background(),
				request: &pb.UploadInitRequest{
					Key:      "testfile.txt",
					Bucket:   "testbucket",
					Metadata: nil,
				},
			},
			wantErr: true,
		},
	}

	// Create connection to server
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewUploadClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.UploadInit(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadHandler.UploadInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && &got.UploadId == nil {
				t.Errorf("UploadHandler.UploadInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUploadService_UploadPart(t *testing.T) {
	metadata := make(map[string]*string)
	metadata["test"] = aws.String("meta")
	file := make([]byte, 50<<20)
	rand.Read(file)
	fileReader := bytes.NewReader(file)

	type fields struct {
		s3Client *s3.S3
	}
	type args struct {
		initKey    *string
		initBucket *string
		key        *string
		bucket     *string
		partNumber *int64
		body       io.ReadSeeker
		ctx        context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "upload part",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile.txt"),
				bucket:     aws.String("testbucket"),
				partNumber: aws.Int64(1),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: false,
		},
		{
			name:   "upload part in folder",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("testfolder/partfile.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("testfolder/partfile.txt"),
				bucket:     aws.String("testbucket"),
				partNumber: aws.Int64(1),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: false,
		},
		{
			name:   "upload part with empty key",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile1.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String(""),
				bucket:     aws.String("testbucket"),
				partNumber: aws.Int64(1),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part with nil key",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile2.txt"),
				initBucket: aws.String("testbucket"),
				key:        nil,
				bucket:     aws.String("testbucket"),
				partNumber: aws.Int64(1),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part with key mismatch",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile3.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile.txt"),
				bucket:     aws.String("testbucket"),
				partNumber: aws.Int64(1),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part with empty bucket",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile4.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile4.txt"),
				bucket:     aws.String(""),
				partNumber: aws.Int64(1),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part with nil bucket",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile5.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile5.txt"),
				bucket:     nil,
				partNumber: aws.Int64(1),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part with bucket mismatch",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile6.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile6.txt"),
				bucket:     aws.String("testbucket1"),
				partNumber: aws.Int64(1),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part with nil body",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile6.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile6.txt"),
				bucket:     aws.String("testbucket"),
				partNumber: aws.Int64(1),
				body:       nil,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part without part number",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile8.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile8.txt"),
				bucket:     aws.String("testbucket"),
				partNumber: nil,
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part with part number lower than 1",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile7.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile7.txt"),
				bucket:     aws.String("testbucket"),
				partNumber: aws.Int64(0),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "upload part with part number greater than 10000",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("partfile7.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("partfile7.txt"),
				bucket:     aws.String("testbucket"),
				partNumber: aws.Int64(10001),
				body:       fileReader,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UploadService{
				s3Client: tt.fields.s3Client,
			}

			initOutput, err := s.UploadInit(tt.args.ctx, tt.args.initKey, tt.args.initBucket, metadata)
			if err != nil {
				t.Errorf("UploadService.UploadInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := s.UploadPart(tt.args.ctx, initOutput.UploadId, tt.args.key, tt.args.bucket, tt.args.partNumber, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadService.UploadPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil || got.ETag == nil || *got.ETag == "") != tt.wantErr {
				t.Errorf("UploadService.UploadPart() = %v", got)
			}
		})
	}

	t.Run("UploadPart - nil UploadID", func(t *testing.T) {
		s := UploadService{
			s3Client: s3Client,
		}

		ctx := context.Background()
		got, err := s.UploadPart(ctx, nil, aws.String("testfile10.txt"), aws.String("testbucket"), aws.Int64(1), fileReader)
		if err == nil {
			t.Errorf("UploadService.UploadPart() error = %v, wantErr %v", err, true)
			return
		}
		if (got == nil || got.ETag == nil || *got.ETag == "") && err == nil {
			t.Errorf("UploadService.UploadPart() = %v", got)
		}
	})
	t.Run("UploadPart - empty UploadID", func(t *testing.T) {
		s := UploadService{
			s3Client: s3Client,
		}

		ctx := context.Background()
		got, err := s.UploadPart(ctx, aws.String(""), aws.String("testfile10.txt"), aws.String("testbucket"), aws.Int64(1), fileReader)
		if err == nil {
			t.Errorf("UploadService.UploadPart() error = %v, wantErr %v", err, true)
			return
		}
		if (got == nil || got.ETag == nil || *got.ETag == "") && err == nil {
			t.Errorf("UploadService.UploadPart() = %v", got)
		}
	})
}

func TestUploadService_UploadComplete(t *testing.T) {
	metadata := make(map[string]*string)
	metadata["test"] = aws.String("meta")
	file := make([]byte, 50<<20)
	rand.Read(file)
	fileReader := bytes.NewReader(file)

	type fields struct {
		s3Client *s3.S3
	}
	type args struct {
		initKey    *string
		initBucket *string
		key        *string
		bucket     *string
		ctx        context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Upload Complete",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("file.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("file.txt"),
				bucket:     aws.String("testbucket"),
				ctx:        context.Background(),
			},
			wantErr: false,
		},
		{
			name:   "Upload Complete to folder",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("testfolder/file.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("testfolder/file.txt"),
				bucket:     aws.String("testbucket"),
				ctx:        context.Background(),
			},
			wantErr: false,
		},
		{
			name:   "Upload Complete with empty key",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("file.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String(""),
				bucket:     aws.String("testbucket"),
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "Upload Complete with nil key",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("file.txt"),
				initBucket: aws.String("testbucket"),
				key:        nil,
				bucket:     aws.String("testbucket"),
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "Upload Complete with key mismatch",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("file.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("file1.txt"),
				bucket:     aws.String("testbucket"),
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "Upload Complete with empty bucket",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("file.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("file.txt"),
				bucket:     aws.String(""),
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "Upload Complete with nil bucket",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("file.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("file.txt"),
				bucket:     nil,
				ctx:        context.Background(),
			},
			wantErr: true,
		},
		{
			name:   "Upload Complete with bucket mismatch",
			fields: fields{s3Client: s3Client},
			args: args{
				initKey:    aws.String("file.txt"),
				initBucket: aws.String("testbucket"),
				key:        aws.String("file1.txt"),
				bucket:     aws.String("testbucket1"),
				ctx:        context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UploadService{
				s3Client: tt.fields.s3Client,
			}

			initOutput, err := s.UploadInit(tt.args.ctx, tt.args.initKey, tt.args.initBucket, metadata)
			if err != nil {
				t.Errorf("UploadService.UploadInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			_, err = s.UploadPart(
				tt.args.ctx,
				initOutput.UploadId,
				tt.args.initKey,
				tt.args.initBucket,
				aws.Int64(1),
				fileReader)

			if err != nil {
				t.Errorf("UploadService.UploadPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := s.UploadComplete(tt.args.ctx, initOutput.UploadId, tt.args.key, tt.args.bucket)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadService.UploadComplete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantErr {
				t.Errorf("UploadService.UploadComplete() = %v", got)
			}
		})
	}
	t.Run("UploadComplete - empty uploadID ", func(t *testing.T) {
		s := UploadService{
			s3Client: s3Client,
		}

		ctx := context.Background()
		got, err := s.UploadComplete(ctx, aws.String(""), aws.String("tests.txt"), aws.String("testbucket"))
		if err == nil {
			t.Errorf("UploadService.UploadComplete() error = %v, wantErr %v", err, true)
			return
		}
		if got != nil && err != nil {
			t.Errorf("UploadService.UploadComplete() = %v", got)
			return
		}
	})

	t.Run("UploadComplete - nil uploadID ", func(t *testing.T) {
		s := UploadService{
			s3Client: s3Client,
		}

		ctx := context.Background()
		got, err := s.UploadComplete(ctx, nil, aws.String("tests.txt"), aws.String("testbucket"))
		if err == nil {
			t.Errorf("UploadService.UploadComplete() error = %v, wantErr %v", err, true)
			return
		}
		if got != nil && err != nil {
			t.Errorf("UploadService.UploadComplete() = %v", got)
			return
		}
	})
}

func TestUploadHandler_UploadMultipart(t *testing.T) {
	// Init global values to use in tests.
	file := make([]byte, 50<<20)
	rand.Read(file)
	metadata := make(map[string]string)
	metadata["test"] = "testt"

	type args struct {
		ctx     context.Context
		request *pb.UploadMultipartRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.UploadMultipartResponse
		wantErr bool
	}{
		{
			name: "Upload Multipart",
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMultipartRequest{
					Key:      "testfile.txt",
					Bucket:   "testbucket",
					File:     file,
					Metadata: metadata,
				},
			},
			wantErr: false,
			want: &pb.UploadMultipartResponse{
				Location: fmt.Sprintf("%s/testbucket/testfile.txt", s3Endpoint),
			},
		},
		{
			name: "Upload Multipart to folder",
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMultipartRequest{
					Key:      "testfolder/testfile.txt",
					Bucket:   "testbucket",
					File:     file,
					Metadata: metadata,
				},
			},
			wantErr: false,
			want: &pb.UploadMultipartResponse{
				Location: fmt.Sprintf("%s/testbucket/testfolder/testfile.txt", s3Endpoint),
			},
		},
		{
			name: "Upload Multipart with empty key",
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMultipartRequest{
					Key:      "",
					Bucket:   "testbucket",
					File:     file,
					Metadata: metadata,
				},
			},
			wantErr: true,
		},
		{
			name: "Upload Multipart with empty bucket",
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMultipartRequest{
					Key:      "testfile.txt",
					Bucket:   "",
					File:     file,
					Metadata: metadata,
				},
			},
			wantErr: true,
		},
		{
			name: "Upload Multipart with nil file",
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMultipartRequest{
					Key:      "testfile.txt",
					Bucket:   "testbucket",
					File:     nil,
					Metadata: metadata,
				},
			},
			wantErr: false,
			want: &pb.UploadMultipartResponse{
				Location: fmt.Sprintf("%s/testbucket/testfile.txt", s3Endpoint),
			},
		},
		{
			name: "Upload Multipart with empty file",
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMultipartRequest{
					Key:      "testfile.txt",
					Bucket:   "testbucket",
					File:     make([]byte, 0),
					Metadata: metadata,
				},
			},
			wantErr: false,
			want: &pb.UploadMultipartResponse{
				Location: fmt.Sprintf("%s/testbucket/testfile.txt", s3Endpoint),
			},
		},
		{
			name: "Upload Multipart with nil metadata",
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMultipartRequest{
					Key:      "testfile.txt",
					Bucket:   "testbucket",
					File:     file,
					Metadata: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "Upload Multipart with empty metadata",
			args: args{
				ctx: context.Background(),
				request: &pb.UploadMultipartRequest{
					Key:      "testfile.txt",
					Bucket:   "testbucket",
					File:     file,
					Metadata: make(map[string]string),
				},
			},
			wantErr: true,
		},
	}

	// Create connection to server
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewUploadClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.UploadMultipart(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadHandler.UploadMultipart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadHandler.UploadMultipart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUploadService_UploadAbort(t *testing.T) {
	metadata := make(map[string]*string)
	metadata["test"] = aws.String("testt")
	file := make([]byte, 50<<20)
	rand.Read(file)
	fileReader := bytes.NewReader(file)
	type fields struct {
		s3Client *s3.S3
	}
	type args struct {
		ctx    aws.Context
		key    *string
		bucket *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "init upload",
			fields: fields{s3Client: s3Client},
			args: args{
				key:    aws.String("testfile.txt"),
				bucket: aws.String("testbucket"),
				ctx:    context.Background(),
			},
			wantErr: false,
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UploadService{
				s3Client: tt.fields.s3Client,
			}

			initOutput, err := s.UploadInit(tt.args.ctx, tt.args.key, tt.args.bucket, metadata)
			if err != nil {
				t.Errorf("UploadService.UploadInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			_, err = s.UploadPart(tt.args.ctx, initOutput.UploadId, tt.args.key, tt.args.bucket, aws.Int64(1), fileReader)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadService.UploadPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := s.UploadAbort(tt.args.ctx, initOutput.UploadId, tt.args.key, tt.args.bucket)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadService.UploadAbort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UploadService.UploadAbort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO:
func TestUploadHandler_UploadAbort(t *testing.T) {
	type fields struct {
		UploadService UploadService
	}
	type args struct {
		ctx     context.Context
		request *pb.UploadAbortRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UploadAbortResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := UploadHandler{
				UploadService: tt.fields.UploadService,
			}
			got, err := h.UploadAbort(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadHandler.UploadAbort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadHandler.UploadAbort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO:
func TestUploadHandler_UploadPart(t *testing.T) {
	type fields struct {
		UploadService UploadService
	}
	type args struct {
		stream pb.Upload_UploadPartServer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := UploadHandler{
				UploadService: tt.fields.UploadService,
			}
			if err := h.UploadPart(tt.args.stream); (err != nil) != tt.wantErr {
				t.Errorf("UploadHandler.UploadPart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO:
func TestUploadHandler_UploadComplete(t *testing.T) {
	type fields struct {
		UploadService UploadService
	}
	type args struct {
		ctx     context.Context
		request *pb.UploadCompleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UploadCompleteResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := UploadHandler{
				UploadService: tt.fields.UploadService,
			}
			got, err := h.UploadComplete(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadHandler.UploadComplete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadHandler.UploadComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}
