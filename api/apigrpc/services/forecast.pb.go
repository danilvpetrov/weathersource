// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: api/apigrpc/services/forecast.proto

package services

import (
	context "context"
	forecast "github.com/danilvpetrov/weathersource/api/forecast"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// ForecastRequest is the request to query the latest weather forecast data.
type ForecastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForecastRequest) Reset() {
	*x = ForecastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apigrpc_services_forecast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecastRequest) ProtoMessage() {}

func (x *ForecastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_apigrpc_services_forecast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecastRequest.ProtoReflect.Descriptor instead.
func (*ForecastRequest) Descriptor() ([]byte, []int) {
	return file_api_apigrpc_services_forecast_proto_rawDescGZIP(), []int{0}
}

// ForecastResponse is the response with the latest weather forecast data.
type ForecastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forecast *forecast.Forecast `protobuf:"bytes,1,opt,name=forecast,proto3" json:"forecast,omitempty"`
}

func (x *ForecastResponse) Reset() {
	*x = ForecastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apigrpc_services_forecast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecastResponse) ProtoMessage() {}

func (x *ForecastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_apigrpc_services_forecast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecastResponse.ProtoReflect.Descriptor instead.
func (*ForecastResponse) Descriptor() ([]byte, []int) {
	return file_api_apigrpc_services_forecast_proto_rawDescGZIP(), []int{1}
}

func (x *ForecastResponse) GetForecast() *forecast.Forecast {
	if x != nil {
		return x.Forecast
	}
	return nil
}

var File_api_apigrpc_services_forecast_proto protoreflect.FileDescriptor

var file_api_apigrpc_services_forecast_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x1a,
	0x1b, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x66, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x54, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x08, 0x66, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x32, 0x88, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x12, 0x7c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x39, 0x2e, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x66, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x73, 0x74, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x2e,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x61, 0x6e, 0x69, 0x6c, 0x76, 0x70, 0x65, 0x74, 0x72, 0x6f, 0x76, 0x2f, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_apigrpc_services_forecast_proto_rawDescOnce sync.Once
	file_api_apigrpc_services_forecast_proto_rawDescData = file_api_apigrpc_services_forecast_proto_rawDesc
)

func file_api_apigrpc_services_forecast_proto_rawDescGZIP() []byte {
	file_api_apigrpc_services_forecast_proto_rawDescOnce.Do(func() {
		file_api_apigrpc_services_forecast_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_apigrpc_services_forecast_proto_rawDescData)
	})
	return file_api_apigrpc_services_forecast_proto_rawDescData
}

var file_api_apigrpc_services_forecast_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_apigrpc_services_forecast_proto_goTypes = []interface{}{
	(*ForecastRequest)(nil),   // 0: weathersource.api.grpc.services.forecast.ForecastRequest
	(*ForecastResponse)(nil),  // 1: weathersource.api.grpc.services.forecast.ForecastResponse
	(*forecast.Forecast)(nil), // 2: weathersource.api.forecast.Forecast
}
var file_api_apigrpc_services_forecast_proto_depIdxs = []int32{
	2, // 0: weathersource.api.grpc.services.forecast.ForecastResponse.forecast:type_name -> weathersource.api.forecast.Forecast
	0, // 1: weathersource.api.grpc.services.forecast.Forecast.Get:input_type -> weathersource.api.grpc.services.forecast.ForecastRequest
	1, // 2: weathersource.api.grpc.services.forecast.Forecast.Get:output_type -> weathersource.api.grpc.services.forecast.ForecastResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_apigrpc_services_forecast_proto_init() }
func file_api_apigrpc_services_forecast_proto_init() {
	if File_api_apigrpc_services_forecast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_apigrpc_services_forecast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_apigrpc_services_forecast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_apigrpc_services_forecast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_apigrpc_services_forecast_proto_goTypes,
		DependencyIndexes: file_api_apigrpc_services_forecast_proto_depIdxs,
		MessageInfos:      file_api_apigrpc_services_forecast_proto_msgTypes,
	}.Build()
	File_api_apigrpc_services_forecast_proto = out.File
	file_api_apigrpc_services_forecast_proto_rawDesc = nil
	file_api_apigrpc_services_forecast_proto_goTypes = nil
	file_api_apigrpc_services_forecast_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ForecastClient is the client API for Forecast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForecastClient interface {
	// Get retrieves the latest weather forecast from the forecast service API.
	Get(ctx context.Context, in *ForecastRequest, opts ...grpc.CallOption) (*ForecastResponse, error)
}

type forecastClient struct {
	cc grpc.ClientConnInterface
}

func NewForecastClient(cc grpc.ClientConnInterface) ForecastClient {
	return &forecastClient{cc}
}

func (c *forecastClient) Get(ctx context.Context, in *ForecastRequest, opts ...grpc.CallOption) (*ForecastResponse, error) {
	out := new(ForecastResponse)
	err := c.cc.Invoke(ctx, "/weathersource.api.grpc.services.forecast.Forecast/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForecastServer is the server API for Forecast service.
type ForecastServer interface {
	// Get retrieves the latest weather forecast from the forecast service API.
	Get(context.Context, *ForecastRequest) (*ForecastResponse, error)
}

// UnimplementedForecastServer can be embedded to have forward compatible implementations.
type UnimplementedForecastServer struct {
}

func (*UnimplementedForecastServer) Get(context.Context, *ForecastRequest) (*ForecastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterForecastServer(s *grpc.Server, srv ForecastServer) {
	s.RegisterService(&_Forecast_serviceDesc, srv)
}

func _Forecast_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForecastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForecastServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weathersource.api.grpc.services.forecast.Forecast/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForecastServer).Get(ctx, req.(*ForecastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Forecast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weathersource.api.grpc.services.forecast.Forecast",
	HandlerType: (*ForecastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Forecast_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/apigrpc/services/forecast.proto",
}
