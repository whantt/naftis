// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/squash/v2/squash.proto

package envoy_config_filter_http_squash_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Squash struct {
	Cluster              string             `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	AttachmentTemplate   *_struct.Struct    `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate,proto3" json:"attachment_template,omitempty"`
	RequestTimeout       *duration.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	AttachmentTimeout    *duration.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,proto3" json:"attachment_timeout,omitempty"`
	AttachmentPollPeriod *duration.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,proto3" json:"attachment_poll_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Squash) Reset()         { *m = Squash{} }
func (m *Squash) String() string { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()    {}
func (*Squash) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fc8434388b1e13, []int{0}
}

func (m *Squash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Squash.Unmarshal(m, b)
}
func (m *Squash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Squash.Marshal(b, m, deterministic)
}
func (m *Squash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Squash.Merge(m, src)
}
func (m *Squash) XXX_Size() int {
	return xxx_messageInfo_Squash.Size(m)
}
func (m *Squash) XXX_DiscardUnknown() {
	xxx_messageInfo_Squash.DiscardUnknown(m)
}

var xxx_messageInfo_Squash proto.InternalMessageInfo

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *_struct.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *duration.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *duration.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *duration.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.config.filter.http.squash.v2.Squash")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/squash/v2/squash.proto", fileDescriptor_63fc8434388b1e13)
}

var fileDescriptor_63fc8434388b1e13 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0xae, 0x12, 0x31,
	0x14, 0xc6, 0x33, 0xe3, 0xbd, 0xdc, 0xd8, 0x9b, 0xf8, 0xa7, 0x1a, 0x19, 0x89, 0x12, 0x64, 0x23,
	0x6e, 0x5a, 0x03, 0x6f, 0x30, 0x71, 0xc1, 0xce, 0x09, 0xb0, 0x27, 0x65, 0xa6, 0x0c, 0x4d, 0x3a,
	0xd3, 0xa1, 0x3d, 0x9d, 0xc0, 0xce, 0x37, 0x70, 0xeb, 0xb3, 0xf8, 0x04, 0x6e, 0x4d, 0x7c, 0x12,
	0x97, 0x2e, 0x8c, 0xa1, 0x2d, 0x01, 0x2f, 0x0b, 0x76, 0xcd, 0x7c, 0xf3, 0xfb, 0xe5, 0x3b, 0xed,
	0x41, 0x94, 0xd7, 0xad, 0xda, 0xd3, 0x5c, 0xd5, 0x6b, 0x51, 0xd2, 0xb5, 0x90, 0xc0, 0x35, 0xdd,
	0x00, 0x34, 0xd4, 0x6c, 0x2d, 0x33, 0x1b, 0xda, 0x8e, 0xc3, 0x89, 0x34, 0x5a, 0x81, 0xc2, 0x43,
	0x07, 0x10, 0x0f, 0x10, 0x0f, 0x90, 0x03, 0x40, 0xc2, 0x6f, 0xed, 0xb8, 0xd7, 0x2f, 0x95, 0x2a,
	0x25, 0xa7, 0x8e, 0x58, 0xd9, 0x35, 0x2d, 0xac, 0x66, 0x20, 0x54, 0xed, 0x1d, 0xbd, 0x37, 0x0f,
	0x73, 0x03, 0xda, 0xe6, 0x10, 0xd2, 0xbe, 0x2d, 0x1a, 0x46, 0x59, 0x5d, 0x2b, 0x70, 0x90, 0xa1,
	0x95, 0x28, 0x35, 0x03, 0x1e, 0xf2, 0xb7, 0x17, 0xb9, 0x01, 0x06, 0xd6, 0x84, 0xb8, 0xdb, 0x32,
	0x29, 0x0a, 0x06, 0x9c, 0x1e, 0x0f, 0x3e, 0x18, 0xfe, 0x8a, 0x51, 0x67, 0xee, 0x3a, 0xe2, 0x77,
	0xe8, 0x2e, 0x97, 0xd6, 0x00, 0xd7, 0x49, 0x34, 0x88, 0x46, 0x8f, 0xd3, 0xbb, 0x3f, 0xe9, 0x8d,
	0x8e, 0x07, 0xd1, 0xec, 0xf8, 0x1d, 0x4f, 0xd1, 0x0b, 0x06, 0xc0, 0xf2, 0x4d, 0xc5, 0x6b, 0x58,
	0x02, 0xaf, 0x1a, 0xc9, 0x80, 0x27, 0xf1, 0x20, 0x1a, 0xdd, 0x8f, 0xbb, 0xc4, 0x4f, 0x40, 0x8e,
	0x13, 0x90, 0xb9, 0x9b, 0x60, 0x86, 0x4f, 0xcc, 0x22, 0x20, 0x38, 0x45, 0x4f, 0x35, 0xdf, 0x5a,
	0x6e, 0x60, 0x09, 0xa2, 0xe2, 0xca, 0x42, 0xf2, 0xc8, 0x59, 0x5e, 0x5f, 0x58, 0x3e, 0x85, 0x7b,
	0x9a, 0x3d, 0x09, 0xc4, 0xc2, 0x03, 0x78, 0x8a, 0xf0, 0x79, 0x9b, 0xa0, 0xb9, 0xb9, 0xa6, 0x79,
	0x7e, 0x56, 0x27, 0x98, 0x3e, 0xa3, 0x57, 0x67, 0xa6, 0x46, 0x49, 0xb9, 0x6c, 0xb8, 0x16, 0xaa,
	0x48, 0x6e, 0xaf, 0xd9, 0x5e, 0x9e, 0xc0, 0x4c, 0x49, 0x99, 0x39, 0x2c, 0xdd, 0xfd, 0xfe, 0xf6,
	0xf7, 0xeb, 0xed, 0x07, 0xfc, 0xde, 0x2f, 0x06, 0xdf, 0x01, 0xaf, 0xcd, 0xe1, 0x59, 0xc2, 0x72,
	0x98, 0xff, 0xb7, 0x63, 0xf2, 0xfd, 0xcb, 0x8f, 0x9f, 0x9d, 0xf8, 0x59, 0x84, 0x3e, 0x0a, 0x45,
	0x1c, 0xd3, 0x68, 0xb5, 0xdb, 0x93, 0xeb, 0x7b, 0x95, 0xde, 0xfb, 0xe7, 0xcb, 0x0e, 0xcd, 0xb2,
	0x68, 0xd5, 0x71, 0x15, 0x27, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x08, 0x88, 0x5b, 0x95, 0xc4,
	0x02, 0x00, 0x00,
}
