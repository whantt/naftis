// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/http_status.proto

package envoy_type

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type StatusCode int32

const (
	StatusCode_Empty                         StatusCode = 0
	StatusCode_Continue                      StatusCode = 100
	StatusCode_OK                            StatusCode = 200
	StatusCode_Created                       StatusCode = 201
	StatusCode_Accepted                      StatusCode = 202
	StatusCode_NonAuthoritativeInformation   StatusCode = 203
	StatusCode_NoContent                     StatusCode = 204
	StatusCode_ResetContent                  StatusCode = 205
	StatusCode_PartialContent                StatusCode = 206
	StatusCode_MultiStatus                   StatusCode = 207
	StatusCode_AlreadyReported               StatusCode = 208
	StatusCode_IMUsed                        StatusCode = 226
	StatusCode_MultipleChoices               StatusCode = 300
	StatusCode_MovedPermanently              StatusCode = 301
	StatusCode_Found                         StatusCode = 302
	StatusCode_SeeOther                      StatusCode = 303
	StatusCode_NotModified                   StatusCode = 304
	StatusCode_UseProxy                      StatusCode = 305
	StatusCode_TemporaryRedirect             StatusCode = 307
	StatusCode_PermanentRedirect             StatusCode = 308
	StatusCode_BadRequest                    StatusCode = 400
	StatusCode_Unauthorized                  StatusCode = 401
	StatusCode_PaymentRequired               StatusCode = 402
	StatusCode_Forbidden                     StatusCode = 403
	StatusCode_NotFound                      StatusCode = 404
	StatusCode_MethodNotAllowed              StatusCode = 405
	StatusCode_NotAcceptable                 StatusCode = 406
	StatusCode_ProxyAuthenticationRequired   StatusCode = 407
	StatusCode_RequestTimeout                StatusCode = 408
	StatusCode_Conflict                      StatusCode = 409
	StatusCode_Gone                          StatusCode = 410
	StatusCode_LengthRequired                StatusCode = 411
	StatusCode_PreconditionFailed            StatusCode = 412
	StatusCode_PayloadTooLarge               StatusCode = 413
	StatusCode_URITooLong                    StatusCode = 414
	StatusCode_UnsupportedMediaType          StatusCode = 415
	StatusCode_RangeNotSatisfiable           StatusCode = 416
	StatusCode_ExpectationFailed             StatusCode = 417
	StatusCode_MisdirectedRequest            StatusCode = 421
	StatusCode_UnprocessableEntity           StatusCode = 422
	StatusCode_Locked                        StatusCode = 423
	StatusCode_FailedDependency              StatusCode = 424
	StatusCode_UpgradeRequired               StatusCode = 426
	StatusCode_PreconditionRequired          StatusCode = 428
	StatusCode_TooManyRequests               StatusCode = 429
	StatusCode_RequestHeaderFieldsTooLarge   StatusCode = 431
	StatusCode_InternalServerError           StatusCode = 500
	StatusCode_NotImplemented                StatusCode = 501
	StatusCode_BadGateway                    StatusCode = 502
	StatusCode_ServiceUnavailable            StatusCode = 503
	StatusCode_GatewayTimeout                StatusCode = 504
	StatusCode_HTTPVersionNotSupported       StatusCode = 505
	StatusCode_VariantAlsoNegotiates         StatusCode = 506
	StatusCode_InsufficientStorage           StatusCode = 507
	StatusCode_LoopDetected                  StatusCode = 508
	StatusCode_NotExtended                   StatusCode = 510
	StatusCode_NetworkAuthenticationRequired StatusCode = 511
)

var StatusCode_name = map[int32]string{
	0:   "Empty",
	100: "Continue",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "NonAuthoritativeInformation",
	204: "NoContent",
	205: "ResetContent",
	206: "PartialContent",
	207: "MultiStatus",
	208: "AlreadyReported",
	226: "IMUsed",
	300: "MultipleChoices",
	301: "MovedPermanently",
	302: "Found",
	303: "SeeOther",
	304: "NotModified",
	305: "UseProxy",
	307: "TemporaryRedirect",
	308: "PermanentRedirect",
	400: "BadRequest",
	401: "Unauthorized",
	402: "PaymentRequired",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	406: "NotAcceptable",
	407: "ProxyAuthenticationRequired",
	408: "RequestTimeout",
	409: "Conflict",
	410: "Gone",
	411: "LengthRequired",
	412: "PreconditionFailed",
	413: "PayloadTooLarge",
	414: "URITooLong",
	415: "UnsupportedMediaType",
	416: "RangeNotSatisfiable",
	417: "ExpectationFailed",
	421: "MisdirectedRequest",
	422: "UnprocessableEntity",
	423: "Locked",
	424: "FailedDependency",
	426: "UpgradeRequired",
	428: "PreconditionRequired",
	429: "TooManyRequests",
	431: "RequestHeaderFieldsTooLarge",
	500: "InternalServerError",
	501: "NotImplemented",
	502: "BadGateway",
	503: "ServiceUnavailable",
	504: "GatewayTimeout",
	505: "HTTPVersionNotSupported",
	506: "VariantAlsoNegotiates",
	507: "InsufficientStorage",
	508: "LoopDetected",
	510: "NotExtended",
	511: "NetworkAuthenticationRequired",
}

var StatusCode_value = map[string]int32{
	"Empty":                         0,
	"Continue":                      100,
	"OK":                            200,
	"Created":                       201,
	"Accepted":                      202,
	"NonAuthoritativeInformation":   203,
	"NoContent":                     204,
	"ResetContent":                  205,
	"PartialContent":                206,
	"MultiStatus":                   207,
	"AlreadyReported":               208,
	"IMUsed":                        226,
	"MultipleChoices":               300,
	"MovedPermanently":              301,
	"Found":                         302,
	"SeeOther":                      303,
	"NotModified":                   304,
	"UseProxy":                      305,
	"TemporaryRedirect":             307,
	"PermanentRedirect":             308,
	"BadRequest":                    400,
	"Unauthorized":                  401,
	"PaymentRequired":               402,
	"Forbidden":                     403,
	"NotFound":                      404,
	"MethodNotAllowed":              405,
	"NotAcceptable":                 406,
	"ProxyAuthenticationRequired":   407,
	"RequestTimeout":                408,
	"Conflict":                      409,
	"Gone":                          410,
	"LengthRequired":                411,
	"PreconditionFailed":            412,
	"PayloadTooLarge":               413,
	"URITooLong":                    414,
	"UnsupportedMediaType":          415,
	"RangeNotSatisfiable":           416,
	"ExpectationFailed":             417,
	"MisdirectedRequest":            421,
	"UnprocessableEntity":           422,
	"Locked":                        423,
	"FailedDependency":              424,
	"UpgradeRequired":               426,
	"PreconditionRequired":          428,
	"TooManyRequests":               429,
	"RequestHeaderFieldsTooLarge":   431,
	"InternalServerError":           500,
	"NotImplemented":                501,
	"BadGateway":                    502,
	"ServiceUnavailable":            503,
	"GatewayTimeout":                504,
	"HTTPVersionNotSupported":       505,
	"VariantAlsoNegotiates":         506,
	"InsufficientStorage":           507,
	"LoopDetected":                  508,
	"NotExtended":                   510,
	"NetworkAuthenticationRequired": 511,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7544d7adacd3389b, []int{0}
}

type HttpStatus struct {
	Code                 StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.type.StatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HttpStatus) Reset()         { *m = HttpStatus{} }
func (m *HttpStatus) String() string { return proto.CompactTextString(m) }
func (*HttpStatus) ProtoMessage()    {}
func (*HttpStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7544d7adacd3389b, []int{0}
}

func (m *HttpStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpStatus.Unmarshal(m, b)
}
func (m *HttpStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpStatus.Marshal(b, m, deterministic)
}
func (m *HttpStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStatus.Merge(m, src)
}
func (m *HttpStatus) XXX_Size() int {
	return xxx_messageInfo_HttpStatus.Size(m)
}
func (m *HttpStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStatus proto.InternalMessageInfo

func (m *HttpStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Empty
}

func init() {
	proto.RegisterEnum("envoy.type.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*HttpStatus)(nil), "envoy.type.HttpStatus")
}

func init() { proto.RegisterFile("envoy/type/http_status.proto", fileDescriptor_7544d7adacd3389b) }

var fileDescriptor_7544d7adacd3389b = []byte{
	// 936 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x49, 0x6f, 0x1c, 0x45,
	0x14, 0x4e, 0x4f, 0x3b, 0x76, 0x5c, 0x71, 0xec, 0x97, 0xca, 0x62, 0x63, 0x12, 0xc9, 0xca, 0x09,
	0x71, 0xb0, 0x25, 0x10, 0xdc, 0xbd, 0xc6, 0x16, 0x9e, 0xc9, 0x68, 0x3c, 0x93, 0x2b, 0x2a, 0x77,
	0xbd, 0x99, 0x29, 0xa5, 0xa7, 0x5e, 0xa7, 0xfa, 0xf5, 0xd8, 0xcd, 0x09, 0xf1, 0x0b, 0xd8, 0xf7,
	0xf5, 0xc0, 0x22, 0x94, 0x10, 0x10, 0x70, 0xe1, 0x17, 0x84, 0x1d, 0xfe, 0x02, 0xbf, 0x81, 0x35,
	0x20, 0x40, 0x55, 0xb3, 0xd8, 0x97, 0xdc, 0x66, 0x5e, 0xbd, 0xe5, 0x5b, 0x5e, 0x3f, 0x71, 0x09,
	0x6d, 0x9f, 0xca, 0x15, 0x2e, 0x33, 0x5c, 0xe9, 0x32, 0x67, 0x4f, 0xe6, 0xac, 0xb8, 0xc8, 0x97,
	0x33, 0x47, 0x4c, 0x52, 0x84, 0xd7, 0x65, 0xff, 0xba, 0x78, 0xb9, 0xd0, 0x99, 0x5a, 0x51, 0xd6,
	0x12, 0x2b, 0x36, 0x64, 0xf3, 0x95, 0xe3, 0xa9, 0x8b, 0xf3, 0x7d, 0x95, 0x1a, 0xad, 0x18, 0x57,
	0x46, 0x3f, 0x06, 0x0f, 0x57, 0x36, 0x84, 0xd8, 0x66, 0xce, 0xf6, 0x42, 0xb2, 0x7c, 0x5c, 0x4c,
	0x24, 0xa4, 0x71, 0x21, 0x5a, 0x8a, 0x1e, 0x9a, 0x7d, 0xe4, 0xe2, 0xf2, 0xd1, 0x80, 0xe5, 0x41,
	0xc6, 0x3a, 0x69, 0x5c, 0x13, 0xf7, 0xd6, 0xa6, 0x9e, 0x89, 0x26, 0x20, 0x5a, 0x3a, 0xd1, 0x08,
	0xf9, 0x0f, 0x7f, 0x39, 0x2d, 0xc4, 0x51, 0x82, 0x9c, 0x16, 0x27, 0x37, 0x7b, 0x19, 0x97, 0x70,
	0x42, 0xce, 0x88, 0x53, 0xeb, 0x64, 0xd9, 0xd8, 0x02, 0x41, 0xcb, 0x29, 0x51, 0xb9, 0xf6, 0x04,
	0xdc, 0x8d, 0xe4, 0x8c, 0x98, 0x5a, 0x77, 0xa8, 0x18, 0x35, 0x7c, 0x1d, 0xc9, 0x33, 0xe2, 0xd4,
	0x6a, 0x92, 0x60, 0xe6, 0xff, 0x7e, 0x13, 0xc9, 0x25, 0xf1, 0x60, 0x8d, 0xec, 0x6a, 0xc1, 0x5d,
	0x72, 0xc6, 0xd3, 0xe9, 0xe3, 0x8e, 0x6d, 0x93, 0xeb, 0x05, 0x66, 0xf0, 0x6d, 0x24, 0x67, 0xc5,
	0x74, 0x8d, 0x7c, 0x5f, 0xb4, 0x0c, 0xdf, 0x45, 0xf2, 0xac, 0x98, 0x69, 0x60, 0x8e, 0x3c, 0x0a,
	0x7d, 0x1f, 0xc9, 0x73, 0x62, 0xb6, 0xae, 0x1c, 0x1b, 0x95, 0x8e, 0x82, 0x3f, 0x44, 0x12, 0xc4,
	0xe9, 0x6a, 0x91, 0xb2, 0x19, 0x60, 0x85, 0x1f, 0x23, 0x79, 0x5e, 0xcc, 0xad, 0xa6, 0x0e, 0x95,
	0x2e, 0x1b, 0x98, 0x91, 0xf3, 0x08, 0x7e, 0x8a, 0xe4, 0x69, 0x31, 0xb9, 0x53, 0x6d, 0xe5, 0xa8,
	0xe1, 0x97, 0x90, 0x12, 0x8a, 0xb2, 0x14, 0xd7, 0xbb, 0x64, 0x12, 0xcc, 0xe1, 0x56, 0x45, 0x5e,
	0x10, 0x50, 0xa5, 0x3e, 0xea, 0x3a, 0xba, 0x9e, 0xb2, 0x68, 0x39, 0x2d, 0xe1, 0x76, 0x45, 0x0a,
	0x71, 0x72, 0x8b, 0x0a, 0xab, 0xe1, 0x93, 0x8a, 0xa7, 0xb5, 0x87, 0x78, 0x8d, 0xbb, 0xe8, 0xe0,
	0x4e, 0xc5, 0x0f, 0xaf, 0x11, 0x57, 0x49, 0x9b, 0xb6, 0x41, 0x0d, 0x9f, 0x86, 0x84, 0x56, 0x8e,
	0x75, 0x47, 0x87, 0x25, 0x7c, 0x56, 0x91, 0x17, 0xc5, 0xd9, 0x26, 0xf6, 0x32, 0x72, 0xca, 0x95,
	0x0d, 0xd4, 0xc6, 0x61, 0xc2, 0xf0, 0x79, 0x88, 0x8f, 0xa7, 0x8c, 0xe3, 0x5f, 0x54, 0xe4, 0x9c,
	0x10, 0x6b, 0x4a, 0x37, 0xf0, 0x66, 0x81, 0x39, 0xc3, 0xb3, 0xb1, 0x97, 0xa1, 0x65, 0xd5, 0x40,
	0xb7, 0xa7, 0x50, 0xc3, 0x73, 0xb1, 0x07, 0x5f, 0x57, 0x65, 0x2f, 0x54, 0xde, 0x2c, 0x8c, 0x43,
	0x0d, 0xcf, 0xc7, 0x5e, 0xbf, 0x2d, 0x72, 0xfb, 0x46, 0x6b, 0xb4, 0xf0, 0x42, 0xec, 0x81, 0xd4,
	0x88, 0x07, 0xc0, 0x5f, 0x8c, 0x03, 0x37, 0xe4, 0x2e, 0xe9, 0x1a, 0xf1, 0x6a, 0x9a, 0xd2, 0x01,
	0x6a, 0x78, 0x29, 0x96, 0x52, 0x9c, 0xf1, 0x81, 0xe0, 0x94, 0xda, 0x4f, 0x11, 0x5e, 0x8e, 0xbd,
	0x57, 0x01, 0xbf, 0x77, 0x0b, 0x2d, 0x9b, 0x24, 0x78, 0x34, 0x9e, 0xf5, 0x4a, 0xec, 0x8d, 0x18,
	0x42, 0x6c, 0x9a, 0x1e, 0x52, 0xc1, 0xf0, 0x6a, 0x18, 0xb8, 0x4e, 0xb6, 0x9d, 0x9a, 0x84, 0xe1,
	0xb5, 0x58, 0x4e, 0x8b, 0x89, 0xab, 0x64, 0x11, 0x5e, 0x0f, 0xe9, 0xbb, 0x68, 0x3b, 0xdc, 0x1d,
	0xf7, 0x78, 0x23, 0x96, 0xf3, 0x42, 0xd6, 0x1d, 0x26, 0x64, 0xb5, 0xf1, 0xed, 0xb7, 0x94, 0x49,
	0x51, 0xc3, 0x9b, 0x23, 0x7a, 0x29, 0x29, 0xdd, 0x24, 0xda, 0x55, 0xae, 0x83, 0xf0, 0x56, 0xec,
	0x85, 0x69, 0x35, 0x76, 0x7c, 0x84, 0x6c, 0x07, 0xde, 0x8e, 0xe5, 0x03, 0xe2, 0x7c, 0xcb, 0xe6,
	0x45, 0x36, 0x70, 0xb8, 0x8a, 0xda, 0xa8, 0x66, 0x99, 0x21, 0xbc, 0x13, 0xcb, 0x05, 0x71, 0xae,
	0xa1, 0x6c, 0x07, 0x6b, 0xc4, 0x7b, 0x8a, 0x4d, 0xde, 0x36, 0x81, 0xda, 0xbb, 0xb1, 0x97, 0x7d,
	0xf3, 0x30, 0xc3, 0x64, 0xf0, 0x41, 0x0d, 0x67, 0xbe, 0x17, 0xc0, 0x54, 0x4d, 0x3e, 0xb0, 0x01,
	0xc7, 0xf2, 0xbf, 0x1f, 0x5a, 0xb5, 0x6c, 0xe6, 0x28, 0xc1, 0x3c, 0xf7, 0x4d, 0x36, 0x2d, 0x1b,
	0x2e, 0xe1, 0x83, 0xd8, 0xef, 0xd3, 0x2e, 0x25, 0x37, 0x50, 0xc3, 0x87, 0x41, 0xdd, 0x41, 0xb3,
	0x0d, 0xcc, 0xd0, 0x6a, 0xb4, 0x49, 0x09, 0x1f, 0x05, 0x2a, 0xad, 0xac, 0xe3, 0x94, 0xc6, 0x31,
	0xf3, 0x8f, 0x03, 0xf2, 0xe3, 0xcc, 0xc7, 0x4f, 0xb7, 0x42, 0x41, 0x93, 0xa8, 0xaa, 0x6c, 0x39,
	0xc4, 0x90, 0xc3, 0xed, 0x60, 0xc8, 0xf0, 0xef, 0x36, 0x2a, 0x8d, 0x6e, 0xcb, 0x60, 0xaa, 0xf3,
	0xb1, 0x3a, 0x77, 0x02, 0xcc, 0x1d, 0xcb, 0xe8, 0xac, 0x4a, 0xf7, 0xd0, 0xf5, 0xd1, 0x6d, 0x3a,
	0x47, 0x0e, 0x7e, 0x0d, 0xda, 0xd7, 0x88, 0x77, 0x7a, 0x59, 0x8a, 0x7e, 0x63, 0x50, 0xc3, 0x6f,
	0xf1, 0x70, 0xcb, 0xae, 0x2a, 0xc6, 0x03, 0x55, 0xc2, 0xef, 0x81, 0xbf, 0xaf, 0x33, 0x09, 0xb6,
	0xac, 0xea, 0x2b, 0x93, 0x06, 0xc1, 0xfe, 0x08, 0xe5, 0xc3, 0xb4, 0x91, 0xd3, 0x7f, 0xc6, 0xf2,
	0x92, 0x98, 0xdf, 0x6e, 0x36, 0xeb, 0xd7, 0xd1, 0xe5, 0x86, 0xac, 0x57, 0x79, 0x64, 0x03, 0xfc,
	0x15, 0xcb, 0x45, 0x71, 0xe1, 0xba, 0x72, 0x46, 0x59, 0x5e, 0x4d, 0x73, 0xaa, 0x61, 0x87, 0xd8,
	0x28, 0xc6, 0x1c, 0xee, 0x0d, 0x71, 0xe6, 0x45, 0xbb, 0x6d, 0x12, 0x83, 0x96, 0xf7, 0x98, 0x9c,
	0xea, 0x20, 0xfc, 0x1d, 0xf6, 0x7c, 0x97, 0x28, 0xdb, 0x40, 0x0e, 0x16, 0xc0, 0x3f, 0xf1, 0xf0,
	0xe3, 0xda, 0x3c, 0x64, 0xaf, 0xa8, 0x86, 0x7f, 0x63, 0x79, 0x45, 0x5c, 0xae, 0x21, 0x1f, 0x90,
	0xbb, 0x71, 0x9f, 0xdd, 0xfc, 0x2f, 0x5e, 0x7b, 0xec, 0xab, 0xa7, 0xef, 0xfe, 0x3c, 0x59, 0x81,
	0x48, 0x2c, 0x18, 0x1a, 0x5c, 0xbb, 0xcc, 0x6f, 0xf3, 0xb1, 0xc3, 0xb7, 0x36, 0x77, 0x74, 0x1f,
	0xeb, 0xfe, 0x64, 0xd6, 0xa3, 0xfd, 0xc9, 0x70, 0x3b, 0x1f, 0xfd, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0x93, 0x0c, 0x30, 0xee, 0x9f, 0x05, 0x00, 0x00,
}
