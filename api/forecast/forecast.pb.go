// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/forecast/forecast.proto

package forecast

import (
	fmt "fmt"
	data "github.com/danilvpetrov/weathersource/data"
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

// Forecast is the container for the latest weather forecast data.
type Forecast struct {
	Current              *data.Data   `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Minutely             []*data.Data `protobuf:"bytes,2,rep,name=minutely,proto3" json:"minutely,omitempty"`
	Hourly               []*data.Data `protobuf:"bytes,3,rep,name=hourly,proto3" json:"hourly,omitempty"`
	Daily                []*data.Data `protobuf:"bytes,4,rep,name=daily,proto3" json:"daily,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Forecast) Reset()         { *m = Forecast{} }
func (m *Forecast) String() string { return proto.CompactTextString(m) }
func (*Forecast) ProtoMessage()    {}
func (*Forecast) Descriptor() ([]byte, []int) {
	return fileDescriptor_03d31b6fcc6cd6d4, []int{0}
}

func (m *Forecast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Forecast.Unmarshal(m, b)
}
func (m *Forecast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Forecast.Marshal(b, m, deterministic)
}
func (m *Forecast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forecast.Merge(m, src)
}
func (m *Forecast) XXX_Size() int {
	return xxx_messageInfo_Forecast.Size(m)
}
func (m *Forecast) XXX_DiscardUnknown() {
	xxx_messageInfo_Forecast.DiscardUnknown(m)
}

var xxx_messageInfo_Forecast proto.InternalMessageInfo

func (m *Forecast) GetCurrent() *data.Data {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *Forecast) GetMinutely() []*data.Data {
	if m != nil {
		return m.Minutely
	}
	return nil
}

func (m *Forecast) GetHourly() []*data.Data {
	if m != nil {
		return m.Hourly
	}
	return nil
}

func (m *Forecast) GetDaily() []*data.Data {
	if m != nil {
		return m.Daily
	}
	return nil
}

func init() {
	proto.RegisterType((*Forecast)(nil), "weathersource.api.forecast.Forecast")
}

func init() {
	proto.RegisterFile("api/forecast/forecast.proto", fileDescriptor_03d31b6fcc6cd6d4)
}

var fileDescriptor_03d31b6fcc6cd6d4 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0xbf, 0x4a, 0xc5, 0x30,
	0x14, 0xc7, 0x71, 0xea, 0xd5, 0x6b, 0x89, 0x83, 0x90, 0x29, 0xd4, 0xa5, 0x38, 0x39, 0x25, 0x52,
	0xfb, 0x04, 0x22, 0x3e, 0x80, 0xa3, 0xdb, 0x69, 0x1a, 0x6d, 0x20, 0x6d, 0xc2, 0xe9, 0x49, 0xa5,
	0x2f, 0xea, 0xf3, 0x88, 0xfd, 0x87, 0x5d, 0xec, 0x12, 0x02, 0xf9, 0x7c, 0x33, 0xfc, 0xd8, 0x1d,
	0x04, 0xab, 0x3e, 0x3c, 0x1a, 0x0d, 0x3d, 0x6d, 0x17, 0x19, 0xd0, 0x93, 0xe7, 0xd9, 0x97, 0x01,
	0x6a, 0x0c, 0xf6, 0x3e, 0xa2, 0x36, 0x12, 0x82, 0x95, 0xab, 0xc8, 0x6e, 0x6b, 0x20, 0x50, 0xbf,
	0xc7, 0x8c, 0xef, 0xbf, 0x13, 0x96, 0xbe, 0x2e, 0xaf, 0xbc, 0x60, 0xd7, 0x3a, 0x22, 0x9a, 0x8e,
	0x44, 0x92, 0x27, 0x0f, 0x37, 0x85, 0x90, 0xfb, 0xbf, 0xa6, 0xf0, 0x05, 0x08, 0xde, 0x56, 0xc8,
	0x4b, 0x96, 0xb6, 0xb6, 0x8b, 0x64, 0xdc, 0x28, 0x2e, 0xf2, 0xd3, 0xbf, 0xd1, 0x26, 0xf9, 0x23,
	0x3b, 0x37, 0x3e, 0xa2, 0x1b, 0xc5, 0xe9, 0xa0, 0x59, 0x1c, 0x97, 0xec, 0xaa, 0x06, 0xeb, 0x46,
	0x71, 0x79, 0x10, 0xcc, 0xec, 0xb9, 0x7c, 0x2f, 0x3e, 0x2d, 0x35, 0xb1, 0x92, 0xda, 0xb7, 0xaa,
	0x86, 0xce, 0xba, 0x21, 0x18, 0x42, 0x3f, 0xa8, 0x5d, 0xa9, 0xfe, 0x4e, 0x59, 0x9d, 0xa7, 0x55,
	0x9e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x32, 0x4c, 0xdf, 0x61, 0x01, 0x00, 0x00,
}
