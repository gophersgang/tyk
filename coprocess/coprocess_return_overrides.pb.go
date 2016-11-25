// Code generated by protoc-gen-go.
// source: coprocess_return_overrides.proto
// DO NOT EDIT!

package coprocess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ReturnOverrides struct {
	ResponseCode  int32  `protobuf:"varint,1,opt,name=response_code,json=responseCode" json:"response_code,omitempty"`
	ResponseError string `protobuf:"bytes,2,opt,name=response_error,json=responseError" json:"response_error,omitempty"`
}

func (m *ReturnOverrides) Reset()                    { *m = ReturnOverrides{} }
func (m *ReturnOverrides) String() string            { return proto.CompactTextString(m) }
func (*ReturnOverrides) ProtoMessage()               {}
func (*ReturnOverrides) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterType((*ReturnOverrides)(nil), "coprocess.ReturnOverrides")
}

func init() { proto.RegisterFile("coprocess_return_overrides.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0x28,
	0xca, 0x4f, 0x4e, 0x2d, 0x2e, 0x8e, 0x2f, 0x4a, 0x2d, 0x29, 0x2d, 0xca, 0x8b, 0xcf, 0x2f, 0x4b,
	0x2d, 0x2a, 0xca, 0x4c, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xab,
	0x50, 0x8a, 0xe5, 0xe2, 0x0f, 0x02, 0x2b, 0xf2, 0x87, 0xa9, 0x11, 0x52, 0xe6, 0xe2, 0x2d, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x8d, 0x4f, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0xe2, 0x81, 0x09, 0x3a, 0xe7, 0xa7, 0xa4, 0x0a, 0xa9, 0x72, 0xf1, 0xc1, 0x15, 0xa5,
	0x16, 0x15, 0xe5, 0x17, 0x49, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb5, 0xba, 0x82, 0x04,
	0x9d, 0x58, 0x3d, 0x98, 0x7f, 0x30, 0x32, 0x26, 0xb1, 0x81, 0xed, 0x35, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x10, 0x47, 0x70, 0x45, 0x9b, 0x00, 0x00, 0x00,
}
