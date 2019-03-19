// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upload_service.proto

package upload

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// UploadMediaRequest is the request for Media Upload.
type UploadMediaRequest struct {
	// File is the file to upload.
	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// File key to store in S3.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The bucket to upload the file to.
	Bucket               string   `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadMediaRequest) Reset()         { *m = UploadMediaRequest{} }
func (m *UploadMediaRequest) String() string { return proto.CompactTextString(m) }
func (*UploadMediaRequest) ProtoMessage()    {}
func (*UploadMediaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{0}
}
func (m *UploadMediaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadMediaRequest.Unmarshal(m, b)
}
func (m *UploadMediaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadMediaRequest.Marshal(b, m, deterministic)
}
func (dst *UploadMediaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadMediaRequest.Merge(dst, src)
}
func (m *UploadMediaRequest) XXX_Size() int {
	return xxx_messageInfo_UploadMediaRequest.Size(m)
}
func (m *UploadMediaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadMediaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadMediaRequest proto.InternalMessageInfo

func (m *UploadMediaRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *UploadMediaRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UploadMediaRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

// UploadMediaResponse is the response for Media Upload.
type UploadMediaResponse struct {
	// The location that the file was uploaded to.
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadMediaResponse) Reset()         { *m = UploadMediaResponse{} }
func (m *UploadMediaResponse) String() string { return proto.CompactTextString(m) }
func (*UploadMediaResponse) ProtoMessage()    {}
func (*UploadMediaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{1}
}
func (m *UploadMediaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadMediaResponse.Unmarshal(m, b)
}
func (m *UploadMediaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadMediaResponse.Marshal(b, m, deterministic)
}
func (dst *UploadMediaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadMediaResponse.Merge(dst, src)
}
func (m *UploadMediaResponse) XXX_Size() int {
	return xxx_messageInfo_UploadMediaResponse.Size(m)
}
func (m *UploadMediaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadMediaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadMediaResponse proto.InternalMessageInfo

func (m *UploadMediaResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

// UploadMultipartRequest is the request for Multipart Upload.
type UploadMultipartRequest struct {
	// File to upload.
	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// File metadata.
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// File key to store in S3.
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// The bucket to upload the file to.
	Bucket               string   `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadMultipartRequest) Reset()         { *m = UploadMultipartRequest{} }
func (m *UploadMultipartRequest) String() string { return proto.CompactTextString(m) }
func (*UploadMultipartRequest) ProtoMessage()    {}
func (*UploadMultipartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{2}
}
func (m *UploadMultipartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadMultipartRequest.Unmarshal(m, b)
}
func (m *UploadMultipartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadMultipartRequest.Marshal(b, m, deterministic)
}
func (dst *UploadMultipartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadMultipartRequest.Merge(dst, src)
}
func (m *UploadMultipartRequest) XXX_Size() int {
	return xxx_messageInfo_UploadMultipartRequest.Size(m)
}
func (m *UploadMultipartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadMultipartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadMultipartRequest proto.InternalMessageInfo

func (m *UploadMultipartRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *UploadMultipartRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *UploadMultipartRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UploadMultipartRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

// UploadMultipartResponse is the response for Multipart Upload.
type UploadMultipartResponse struct {
	// The location that the file was uploaded to.
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadMultipartResponse) Reset()         { *m = UploadMultipartResponse{} }
func (m *UploadMultipartResponse) String() string { return proto.CompactTextString(m) }
func (*UploadMultipartResponse) ProtoMessage()    {}
func (*UploadMultipartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{3}
}
func (m *UploadMultipartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadMultipartResponse.Unmarshal(m, b)
}
func (m *UploadMultipartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadMultipartResponse.Marshal(b, m, deterministic)
}
func (dst *UploadMultipartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadMultipartResponse.Merge(dst, src)
}
func (m *UploadMultipartResponse) XXX_Size() int {
	return xxx_messageInfo_UploadMultipartResponse.Size(m)
}
func (m *UploadMultipartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadMultipartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadMultipartResponse proto.InternalMessageInfo

func (m *UploadMultipartResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

// UploadResumableInitResponse is the response for Initiating Resumable Upload.
type UploadResumableInitResponse struct {
	// Upload ID generated for resumable upload of a file.
	UploadId             string   `protobuf:"bytes,1,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResumableInitResponse) Reset()         { *m = UploadResumableInitResponse{} }
func (m *UploadResumableInitResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResumableInitResponse) ProtoMessage()    {}
func (*UploadResumableInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{4}
}
func (m *UploadResumableInitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResumableInitResponse.Unmarshal(m, b)
}
func (m *UploadResumableInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResumableInitResponse.Marshal(b, m, deterministic)
}
func (dst *UploadResumableInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResumableInitResponse.Merge(dst, src)
}
func (m *UploadResumableInitResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResumableInitResponse.Size(m)
}
func (m *UploadResumableInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResumableInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResumableInitResponse proto.InternalMessageInfo

func (m *UploadResumableInitResponse) GetUploadId() string {
	if m != nil {
		return m.UploadId
	}
	return ""
}

// UploadResumableRequest is the request for Resumable Part Upload.
type UploadResumableRequest struct {
	// Types that are valid to be assigned to Value:
	//	*UploadResumableRequest_UploadResumableInit_
	//	*UploadResumableRequest_Part_
	Value                isUploadResumableRequest_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UploadResumableRequest) Reset()         { *m = UploadResumableRequest{} }
func (m *UploadResumableRequest) String() string { return proto.CompactTextString(m) }
func (*UploadResumableRequest) ProtoMessage()    {}
func (*UploadResumableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{5}
}
func (m *UploadResumableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResumableRequest.Unmarshal(m, b)
}
func (m *UploadResumableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResumableRequest.Marshal(b, m, deterministic)
}
func (dst *UploadResumableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResumableRequest.Merge(dst, src)
}
func (m *UploadResumableRequest) XXX_Size() int {
	return xxx_messageInfo_UploadResumableRequest.Size(m)
}
func (m *UploadResumableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResumableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResumableRequest proto.InternalMessageInfo

type isUploadResumableRequest_Value interface {
	isUploadResumableRequest_Value()
}

type UploadResumableRequest_UploadResumableInit_ struct {
	UploadResumableInit *UploadResumableRequest_UploadResumableInit `protobuf:"bytes,1,opt,name=uploadResumableInit,proto3,oneof"`
}
type UploadResumableRequest_Part_ struct {
	Part *UploadResumableRequest_Part `protobuf:"bytes,2,opt,name=part,proto3,oneof"`
}

func (*UploadResumableRequest_UploadResumableInit_) isUploadResumableRequest_Value() {}
func (*UploadResumableRequest_Part_) isUploadResumableRequest_Value()                {}

func (m *UploadResumableRequest) GetValue() isUploadResumableRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UploadResumableRequest) GetUploadResumableInit() *UploadResumableRequest_UploadResumableInit {
	if x, ok := m.GetValue().(*UploadResumableRequest_UploadResumableInit_); ok {
		return x.UploadResumableInit
	}
	return nil
}

func (m *UploadResumableRequest) GetPart() *UploadResumableRequest_Part {
	if x, ok := m.GetValue().(*UploadResumableRequest_Part_); ok {
		return x.Part
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UploadResumableRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UploadResumableRequest_OneofMarshaler, _UploadResumableRequest_OneofUnmarshaler, _UploadResumableRequest_OneofSizer, []interface{}{
		(*UploadResumableRequest_UploadResumableInit_)(nil),
		(*UploadResumableRequest_Part_)(nil),
	}
}

func _UploadResumableRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UploadResumableRequest)
	// value
	switch x := m.Value.(type) {
	case *UploadResumableRequest_UploadResumableInit_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UploadResumableInit); err != nil {
			return err
		}
	case *UploadResumableRequest_Part_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Part); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UploadResumableRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _UploadResumableRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UploadResumableRequest)
	switch tag {
	case 1: // value.uploadResumableInit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UploadResumableRequest_UploadResumableInit)
		err := b.DecodeMessage(msg)
		m.Value = &UploadResumableRequest_UploadResumableInit_{msg}
		return true, err
	case 2: // value.part
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UploadResumableRequest_Part)
		err := b.DecodeMessage(msg)
		m.Value = &UploadResumableRequest_Part_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UploadResumableRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UploadResumableRequest)
	// value
	switch x := m.Value.(type) {
	case *UploadResumableRequest_UploadResumableInit_:
		s := proto.Size(x.UploadResumableInit)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UploadResumableRequest_Part_:
		s := proto.Size(x.Part)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// UploadResumableInit is the data for Initiating Resumable Upload.
type UploadResumableRequest_UploadResumableInit struct {
	// File key to store in S3.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The bucket to upload the file to.
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// File metadata.
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UploadResumableRequest_UploadResumableInit) Reset() {
	*m = UploadResumableRequest_UploadResumableInit{}
}
func (m *UploadResumableRequest_UploadResumableInit) String() string {
	return proto.CompactTextString(m)
}
func (*UploadResumableRequest_UploadResumableInit) ProtoMessage() {}
func (*UploadResumableRequest_UploadResumableInit) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{5, 0}
}
func (m *UploadResumableRequest_UploadResumableInit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResumableRequest_UploadResumableInit.Unmarshal(m, b)
}
func (m *UploadResumableRequest_UploadResumableInit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResumableRequest_UploadResumableInit.Marshal(b, m, deterministic)
}
func (dst *UploadResumableRequest_UploadResumableInit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResumableRequest_UploadResumableInit.Merge(dst, src)
}
func (m *UploadResumableRequest_UploadResumableInit) XXX_Size() int {
	return xxx_messageInfo_UploadResumableRequest_UploadResumableInit.Size(m)
}
func (m *UploadResumableRequest_UploadResumableInit) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResumableRequest_UploadResumableInit.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResumableRequest_UploadResumableInit proto.InternalMessageInfo

func (m *UploadResumableRequest_UploadResumableInit) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UploadResumableRequest_UploadResumableInit) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *UploadResumableRequest_UploadResumableInit) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Part is the current part to upload.
type UploadResumableRequest_Part struct {
	// File part chunk.
	Part []byte `protobuf:"bytes,1,opt,name=part,proto3" json:"part,omitempty"`
	// Part number.
	PartNumber           int64    `protobuf:"varint,2,opt,name=partNumber,proto3" json:"partNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResumableRequest_Part) Reset()         { *m = UploadResumableRequest_Part{} }
func (m *UploadResumableRequest_Part) String() string { return proto.CompactTextString(m) }
func (*UploadResumableRequest_Part) ProtoMessage()    {}
func (*UploadResumableRequest_Part) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{5, 1}
}
func (m *UploadResumableRequest_Part) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResumableRequest_Part.Unmarshal(m, b)
}
func (m *UploadResumableRequest_Part) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResumableRequest_Part.Marshal(b, m, deterministic)
}
func (dst *UploadResumableRequest_Part) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResumableRequest_Part.Merge(dst, src)
}
func (m *UploadResumableRequest_Part) XXX_Size() int {
	return xxx_messageInfo_UploadResumableRequest_Part.Size(m)
}
func (m *UploadResumableRequest_Part) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResumableRequest_Part.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResumableRequest_Part proto.InternalMessageInfo

func (m *UploadResumableRequest_Part) GetPart() []byte {
	if m != nil {
		return m.Part
	}
	return nil
}

func (m *UploadResumableRequest_Part) GetPartNumber() int64 {
	if m != nil {
		return m.PartNumber
	}
	return 0
}

// UploadResumableResponse is the response for Resumable Part Upload.
type UploadResumableResponse struct {
	// The location that the file was uploaded to.
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResumableResponse) Reset()         { *m = UploadResumableResponse{} }
func (m *UploadResumableResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResumableResponse) ProtoMessage()    {}
func (*UploadResumableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_service_3b4eef93eea338f0, []int{6}
}
func (m *UploadResumableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResumableResponse.Unmarshal(m, b)
}
func (m *UploadResumableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResumableResponse.Marshal(b, m, deterministic)
}
func (dst *UploadResumableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResumableResponse.Merge(dst, src)
}
func (m *UploadResumableResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResumableResponse.Size(m)
}
func (m *UploadResumableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResumableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResumableResponse proto.InternalMessageInfo

func (m *UploadResumableResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func init() {
	proto.RegisterType((*UploadMediaRequest)(nil), "upload.UploadMediaRequest")
	proto.RegisterType((*UploadMediaResponse)(nil), "upload.UploadMediaResponse")
	proto.RegisterType((*UploadMultipartRequest)(nil), "upload.UploadMultipartRequest")
	proto.RegisterMapType((map[string]string)(nil), "upload.UploadMultipartRequest.MetadataEntry")
	proto.RegisterType((*UploadMultipartResponse)(nil), "upload.UploadMultipartResponse")
	proto.RegisterType((*UploadResumableInitResponse)(nil), "upload.UploadResumableInitResponse")
	proto.RegisterType((*UploadResumableRequest)(nil), "upload.UploadResumableRequest")
	proto.RegisterType((*UploadResumableRequest_UploadResumableInit)(nil), "upload.UploadResumableRequest.UploadResumableInit")
	proto.RegisterMapType((map[string]string)(nil), "upload.UploadResumableRequest.UploadResumableInit.MetadataEntry")
	proto.RegisterType((*UploadResumableRequest_Part)(nil), "upload.UploadResumableRequest.Part")
	proto.RegisterType((*UploadResumableResponse)(nil), "upload.UploadResumableResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UploadClient is the client API for Upload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UploadClient interface {
	// The function Uploads the given file.
	//
	// Returns the Location of the file as output.
	//
	// In case of an error the error is returned.
	UploadMedia(ctx context.Context, in *UploadMediaRequest, opts ...grpc.CallOption) (*UploadMediaResponse, error)
	UploadMultipart(ctx context.Context, in *UploadMultipartRequest, opts ...grpc.CallOption) (*UploadMultipartResponse, error)
	UploadResumable(ctx context.Context, opts ...grpc.CallOption) (Upload_UploadResumableClient, error)
}

type uploadClient struct {
	cc *grpc.ClientConn
}

func NewUploadClient(cc *grpc.ClientConn) UploadClient {
	return &uploadClient{cc}
}

func (c *uploadClient) UploadMedia(ctx context.Context, in *UploadMediaRequest, opts ...grpc.CallOption) (*UploadMediaResponse, error) {
	out := new(UploadMediaResponse)
	err := c.cc.Invoke(ctx, "/upload.Upload/UploadMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadClient) UploadMultipart(ctx context.Context, in *UploadMultipartRequest, opts ...grpc.CallOption) (*UploadMultipartResponse, error) {
	out := new(UploadMultipartResponse)
	err := c.cc.Invoke(ctx, "/upload.Upload/UploadMultipart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadClient) UploadResumable(ctx context.Context, opts ...grpc.CallOption) (Upload_UploadResumableClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Upload_serviceDesc.Streams[0], "/upload.Upload/UploadResumable", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadUploadResumableClient{stream}
	return x, nil
}

type Upload_UploadResumableClient interface {
	Send(*UploadResumableRequest) error
	CloseAndRecv() (*UploadResumableResponse, error)
	grpc.ClientStream
}

type uploadUploadResumableClient struct {
	grpc.ClientStream
}

func (x *uploadUploadResumableClient) Send(m *UploadResumableRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadUploadResumableClient) CloseAndRecv() (*UploadResumableResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResumableResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServer is the server API for Upload service.
type UploadServer interface {
	// The function Uploads the given file.
	//
	// Returns the Location of the file as output.
	//
	// In case of an error the error is returned.
	UploadMedia(context.Context, *UploadMediaRequest) (*UploadMediaResponse, error)
	UploadMultipart(context.Context, *UploadMultipartRequest) (*UploadMultipartResponse, error)
	UploadResumable(Upload_UploadResumableServer) error
}

func RegisterUploadServer(s *grpc.Server, srv UploadServer) {
	s.RegisterService(&_Upload_serviceDesc, srv)
}

func _Upload_UploadMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).UploadMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/upload.Upload/UploadMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).UploadMedia(ctx, req.(*UploadMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upload_UploadMultipart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadMultipartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).UploadMultipart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/upload.Upload/UploadMultipart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).UploadMultipart(ctx, req.(*UploadMultipartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upload_UploadResumable_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServer).UploadResumable(&uploadUploadResumableServer{stream})
}

type Upload_UploadResumableServer interface {
	SendAndClose(*UploadResumableResponse) error
	Recv() (*UploadResumableRequest, error)
	grpc.ServerStream
}

type uploadUploadResumableServer struct {
	grpc.ServerStream
}

func (x *uploadUploadResumableServer) SendAndClose(m *UploadResumableResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadUploadResumableServer) Recv() (*UploadResumableRequest, error) {
	m := new(UploadResumableRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Upload_serviceDesc = grpc.ServiceDesc{
	ServiceName: "upload.Upload",
	HandlerType: (*UploadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadMedia",
			Handler:    _Upload_UploadMedia_Handler,
		},
		{
			MethodName: "UploadMultipart",
			Handler:    _Upload_UploadMultipart_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadResumable",
			Handler:       _Upload_UploadResumable_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "upload_service.proto",
}

func init() {
	proto.RegisterFile("upload_service.proto", fileDescriptor_upload_service_3b4eef93eea338f0)
}

var fileDescriptor_upload_service_3b4eef93eea338f0 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x6d, 0x3e, 0x8c, 0xcf, 0x1b, 0x45, 0x99, 0xf7, 0xa8, 0x25, 0x85, 0x2a, 0x71, 0xd3, 0x85,
	0x06, 0x8c, 0x1b, 0x5b, 0x37, 0x22, 0x08, 0xed, 0xa2, 0x22, 0x41, 0x5d, 0x09, 0x32, 0x69, 0xa6,
	0x10, 0x9a, 0x36, 0x31, 0x99, 0x29, 0x74, 0xe7, 0xef, 0x14, 0xfa, 0x5f, 0x4c, 0x66, 0xf2, 0x9d,
	0xb4, 0x45, 0xde, 0xaa, 0x33, 0x93, 0x7b, 0x4e, 0xef, 0x39, 0x73, 0xee, 0xc0, 0x1d, 0x8b, 0x82,
	0x10, 0x7b, 0xbf, 0x12, 0x12, 0x1f, 0xfc, 0x35, 0xb1, 0xa2, 0x38, 0xa4, 0x21, 0xd2, 0xc4, 0xa9,
	0xe9, 0x00, 0xfa, 0xce, 0x57, 0x2b, 0xe2, 0xf9, 0xd8, 0x21, 0xbf, 0x19, 0x49, 0x28, 0x42, 0xa0,
	0x6e, 0xfc, 0x80, 0x8c, 0xa4, 0x97, 0xd2, 0xf4, 0xb1, 0xc3, 0xd7, 0xe8, 0x19, 0x28, 0x5b, 0x72,
	0x1c, 0xc9, 0xe9, 0xd1, 0x23, 0x27, 0x5b, 0xa2, 0x21, 0x68, 0x2e, 0x5b, 0x6f, 0x09, 0x1d, 0x29,
	0xfc, 0x30, 0xdf, 0x99, 0x6f, 0xe0, 0xb6, 0xc1, 0x99, 0x44, 0xe1, 0x3e, 0x21, 0x59, 0x79, 0xc8,
	0x68, 0xc4, 0x28, 0xa7, 0x4d, 0xcb, 0xc5, 0xce, 0x3c, 0x49, 0x30, 0xcc, 0xeb, 0x59, 0x40, 0xfd,
	0x08, 0xc7, 0xf4, 0x52, 0x1f, 0x0b, 0xb8, 0xd9, 0x11, 0x8a, 0x3d, 0x4c, 0x71, 0xda, 0x8c, 0x32,
	0xd5, 0xed, 0xd7, 0x96, 0x10, 0x63, 0xf5, 0xb3, 0x58, 0xab, 0xbc, 0xfc, 0xf3, 0x9e, 0xc6, 0x47,
	0xa7, 0x44, 0x17, 0x8a, 0x94, 0x3e, 0x45, 0x6a, 0x5d, 0x91, 0xf1, 0x01, 0x9e, 0x34, 0x48, 0x0a,
	0xa8, 0x54, 0x41, 0xef, 0xe0, 0xc1, 0x01, 0x07, 0x8c, 0xe4, 0x06, 0x89, 0xcd, 0x5c, 0x7e, 0x2f,
	0x99, 0x6f, 0xe1, 0x79, 0xa7, 0xb1, 0x2b, 0x96, 0xcc, 0x60, 0x2c, 0x20, 0x69, 0x25, 0xdb, 0x61,
	0x37, 0x20, 0xcb, 0xbd, 0x5f, 0xc1, 0x0c, 0xb8, 0x11, 0x8a, 0x97, 0x5e, 0x0e, 0x2c, 0xf7, 0xe6,
	0x49, 0x29, 0xdc, 0x2c, 0xb1, 0x85, 0x9b, 0x1b, 0xb8, 0x65, 0x5d, 0x56, 0xce, 0xa0, 0xdb, 0x76,
	0xd3, 0xc4, 0x36, 0xd8, 0xea, 0xe9, 0x67, 0x31, 0x70, 0xfa, 0x08, 0xd1, 0x0c, 0xd4, 0x4c, 0x25,
	0x77, 0x42, 0xb7, 0x5f, 0x5d, 0x21, 0xfe, 0x9a, 0x96, 0xa6, 0x4c, 0x1c, 0x62, 0xfc, 0x95, 0x8a,
	0xec, 0x34, 0x29, 0xbb, 0x7e, 0x57, 0x57, 0x25, 0xd7, 0xaf, 0x0a, 0xfd, 0xac, 0xc5, 0x43, 0xe1,
	0xf1, 0xf8, 0xf8, 0xff, 0xca, 0xce, 0x45, 0xe6, 0x5e, 0x41, 0x30, 0xe6, 0xa0, 0x66, 0x62, 0xb3,
	0x54, 0x73, 0x7f, 0xf2, 0x54, 0x67, 0x6b, 0x34, 0x01, 0xc8, 0x7e, 0xbf, 0xb0, 0x9d, 0x4b, 0x62,
	0x0e, 0x55, 0x9c, 0xda, 0xc9, 0xa7, 0x87, 0x39, 0x6b, 0x95, 0xa6, 0x9a, 0x8e, 0xcb, 0x69, 0xb2,
	0xff, 0xc8, 0xa0, 0x09, 0x4c, 0x3a, 0x3c, 0x7a, 0x6d, 0x34, 0x91, 0xd1, 0x9a, 0x9c, 0xda, 0x1b,
	0x60, 0x8c, 0x7b, 0xbf, 0x89, 0xbf, 0x32, 0x07, 0xe8, 0x1b, 0x3c, 0x6d, 0xa5, 0x1a, 0x4d, 0x2e,
	0xcf, 0xa1, 0xf1, 0xe2, 0xec, 0xf7, 0x92, 0xf5, 0x47, 0xc1, 0x5a, 0xaa, 0x6b, 0xb3, 0xb6, 0xaf,
	0xaf, 0xcd, 0xda, 0xb1, 0xc5, 0x1c, 0x4c, 0x25, 0x57, 0xe3, 0xaf, 0xde, 0xbb, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0xfb, 0x40, 0x6d, 0x0d, 0x05, 0x00, 0x00,
}
