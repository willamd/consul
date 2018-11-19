// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/rds.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

import bytes "bytes"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RouteConfiguration struct {
	// The name of the route configuration. For example, it might match
	// :ref:`route_config_name
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.Rds.route_config_name>` in
	// :ref:`envoy_api_msg_config.filter.network.http_connection_manager.v2.Rds`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An array of virtual hosts that make up the route table.
	VirtualHosts []route.VirtualHost `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts" json:"virtual_hosts"`
	// Optionally specifies a list of HTTP headers that the connection manager
	// will consider to be internal only. If they are found on external requests they will be cleaned
	// prior to filter invocation. See :ref:`config_http_conn_man_headers_x-envoy-internal` for more
	// information.
	InternalOnlyHeaders []string `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders" json:"internal_only_headers,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response that
	// the connection manager encodes. Headers specified at this level are applied
	// after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	ResponseHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// that the connection manager encodes.
	ResponseHeadersToRemove []string `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove" json:"response_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request
	// routed by the HTTP connection manager. Headers specified at this level are
	// applied after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	RequestHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd" json:"request_headers_to_add,omitempty"`
	// An optional boolean that specifies whether the clusters that the route
	// table refers to will be validated by the cluster manager. If set to true
	// and a route refers to a non-existent cluster, the route table will not
	// load. If set to false and a route refers to a non-existent cluster, the
	// route table will load and the router filter will return a 404 if the route
	// is selected at runtime. This setting defaults to true if the route table
	// is statically defined via the :ref:`route_config
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.route_config>`
	// option. This setting default to false if the route table is loaded dynamically via the
	// :ref:`rds
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.rds>`
	// option. Users may which to override the default behavior in certain cases (for example when
	// using CDS with a static route table).
	ValidateClusters     *types.BoolValue `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_rds_49b8428ed0bf9596, []int{0}
}
func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(dst, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []route.VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetValidateClusters() *types.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.api.v2.RouteConfiguration")
}
func (this *RouteConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteConfiguration)
	if !ok {
		that2, ok := that.(RouteConfiguration)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.VirtualHosts) != len(that1.VirtualHosts) {
		return false
	}
	for i := range this.VirtualHosts {
		if !this.VirtualHosts[i].Equal(&that1.VirtualHosts[i]) {
			return false
		}
	}
	if len(this.InternalOnlyHeaders) != len(that1.InternalOnlyHeaders) {
		return false
	}
	for i := range this.InternalOnlyHeaders {
		if this.InternalOnlyHeaders[i] != that1.InternalOnlyHeaders[i] {
			return false
		}
	}
	if len(this.ResponseHeadersToAdd) != len(that1.ResponseHeadersToAdd) {
		return false
	}
	for i := range this.ResponseHeadersToAdd {
		if !this.ResponseHeadersToAdd[i].Equal(that1.ResponseHeadersToAdd[i]) {
			return false
		}
	}
	if len(this.ResponseHeadersToRemove) != len(that1.ResponseHeadersToRemove) {
		return false
	}
	for i := range this.ResponseHeadersToRemove {
		if this.ResponseHeadersToRemove[i] != that1.ResponseHeadersToRemove[i] {
			return false
		}
	}
	if len(this.RequestHeadersToAdd) != len(that1.RequestHeadersToAdd) {
		return false
	}
	for i := range this.RequestHeadersToAdd {
		if !this.RequestHeadersToAdd[i].Equal(that1.RequestHeadersToAdd[i]) {
			return false
		}
	}
	if !this.ValidateClusters.Equal(that1.ValidateClusters) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RouteDiscoveryService service

type RouteDiscoveryServiceClient interface {
	StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error)
	IncrementalRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_IncrementalRoutesClient, error)
	FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type routeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteDiscoveryServiceClient(cc *grpc.ClientConn) RouteDiscoveryServiceClient {
	return &routeDiscoveryServiceClient{cc}
}

func (c *routeDiscoveryServiceClient) StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.RouteDiscoveryService/StreamRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceStreamRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_StreamRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceStreamRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceStreamRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) IncrementalRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_IncrementalRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.RouteDiscoveryService/IncrementalRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceIncrementalRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_IncrementalRoutesClient interface {
	Send(*IncrementalDiscoveryRequest) error
	Recv() (*IncrementalDiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceIncrementalRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceIncrementalRoutesClient) Send(m *IncrementalDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceIncrementalRoutesClient) Recv() (*IncrementalDiscoveryResponse, error) {
	m := new(IncrementalDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.RouteDiscoveryService/FetchRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RouteDiscoveryService service

type RouteDiscoveryServiceServer interface {
	StreamRoutes(RouteDiscoveryService_StreamRoutesServer) error
	IncrementalRoutes(RouteDiscoveryService_IncrementalRoutesServer) error
	FetchRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterRouteDiscoveryServiceServer(s *grpc.Server, srv RouteDiscoveryServiceServer) {
	s.RegisterService(&_RouteDiscoveryService_serviceDesc, srv)
}

func _RouteDiscoveryService_StreamRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).StreamRoutes(&routeDiscoveryServiceStreamRoutesServer{stream})
}

type RouteDiscoveryService_StreamRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceStreamRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceStreamRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_IncrementalRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).IncrementalRoutes(&routeDiscoveryServiceIncrementalRoutesServer{stream})
}

type RouteDiscoveryService_IncrementalRoutesServer interface {
	Send(*IncrementalDiscoveryResponse) error
	Recv() (*IncrementalDiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceIncrementalRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceIncrementalRoutesServer) Send(m *IncrementalDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceIncrementalRoutesServer) Recv() (*IncrementalDiscoveryRequest, error) {
	m := new(IncrementalDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_FetchRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.RouteDiscoveryService/FetchRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.RouteDiscoveryService",
	HandlerType: (*RouteDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRoutes",
			Handler:    _RouteDiscoveryService_FetchRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRoutes",
			Handler:       _RouteDiscoveryService_StreamRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "IncrementalRoutes",
			Handler:       _RouteDiscoveryService_IncrementalRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/rds.proto",
}

func (m *RouteConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RouteConfiguration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRds(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.VirtualHosts) > 0 {
		for _, msg := range m.VirtualHosts {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRds(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.InternalOnlyHeaders) > 0 {
		for _, s := range m.InternalOnlyHeaders {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.ResponseHeadersToAdd) > 0 {
		for _, msg := range m.ResponseHeadersToAdd {
			dAtA[i] = 0x22
			i++
			i = encodeVarintRds(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ResponseHeadersToRemove) > 0 {
		for _, s := range m.ResponseHeadersToRemove {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.RequestHeadersToAdd) > 0 {
		for _, msg := range m.RequestHeadersToAdd {
			dAtA[i] = 0x32
			i++
			i = encodeVarintRds(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.ValidateClusters != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintRds(dAtA, i, uint64(m.ValidateClusters.Size()))
		n1, err := m.ValidateClusters.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRds(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RouteConfiguration) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRds(uint64(l))
	}
	if len(m.VirtualHosts) > 0 {
		for _, e := range m.VirtualHosts {
			l = e.Size()
			n += 1 + l + sovRds(uint64(l))
		}
	}
	if len(m.InternalOnlyHeaders) > 0 {
		for _, s := range m.InternalOnlyHeaders {
			l = len(s)
			n += 1 + l + sovRds(uint64(l))
		}
	}
	if len(m.ResponseHeadersToAdd) > 0 {
		for _, e := range m.ResponseHeadersToAdd {
			l = e.Size()
			n += 1 + l + sovRds(uint64(l))
		}
	}
	if len(m.ResponseHeadersToRemove) > 0 {
		for _, s := range m.ResponseHeadersToRemove {
			l = len(s)
			n += 1 + l + sovRds(uint64(l))
		}
	}
	if len(m.RequestHeadersToAdd) > 0 {
		for _, e := range m.RequestHeadersToAdd {
			l = e.Size()
			n += 1 + l + sovRds(uint64(l))
		}
	}
	if m.ValidateClusters != nil {
		l = m.ValidateClusters.Size()
		n += 1 + l + sovRds(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRds(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRds(x uint64) (n int) {
	return sovRds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RouteConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRds
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RouteConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RouteConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRds
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualHosts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualHosts = append(m.VirtualHosts, route.VirtualHost{})
			if err := m.VirtualHosts[len(m.VirtualHosts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalOnlyHeaders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRds
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InternalOnlyHeaders = append(m.InternalOnlyHeaders, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseHeadersToAdd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseHeadersToAdd = append(m.ResponseHeadersToAdd, &core.HeaderValueOption{})
			if err := m.ResponseHeadersToAdd[len(m.ResponseHeadersToAdd)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseHeadersToRemove", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRds
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseHeadersToRemove = append(m.ResponseHeadersToRemove, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestHeadersToAdd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestHeadersToAdd = append(m.RequestHeadersToAdd, &core.HeaderValueOption{})
			if err := m.RequestHeadersToAdd[len(m.RequestHeadersToAdd)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidateClusters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidateClusters == nil {
				m.ValidateClusters = &types.BoolValue{}
			}
			if err := m.ValidateClusters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRds
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRds
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRds(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRds   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/api/v2/rds.proto", fileDescriptor_rds_49b8428ed0bf9596) }

var fileDescriptor_rds_49b8428ed0bf9596 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xef, 0x25, 0xa1, 0xa8, 0x97, 0x20, 0xd1, 0x4b, 0xda, 0x86, 0x08, 0x39, 0x51, 0xc4, 0x10,
	0x3a, 0xd8, 0xc8, 0x6c, 0x65, 0x22, 0x45, 0x50, 0x58, 0x2a, 0xb9, 0x50, 0x09, 0x31, 0x58, 0x17,
	0xfb, 0x35, 0xb1, 0xe4, 0xdc, 0x33, 0x77, 0x67, 0xa3, 0xac, 0x4c, 0xec, 0x7c, 0x09, 0x3e, 0x4a,
	0x47, 0x10, 0x3b, 0x42, 0x11, 0xe2, 0x73, 0x20, 0x9f, 0x6d, 0xa8, 0x5b, 0x90, 0x40, 0x62, 0x89,
	0x2e, 0xef, 0xf7, 0xe7, 0xfd, 0xee, 0xbd, 0x33, 0xdd, 0x05, 0x91, 0xe1, 0xca, 0xe1, 0x49, 0xe4,
	0x64, 0xae, 0x23, 0x43, 0x65, 0x27, 0x12, 0x35, 0xb2, 0x8e, 0xa9, 0xdb, 0x3c, 0x89, 0xec, 0xcc,
	0x1d, 0xdc, 0xae, 0xb1, 0x02, 0x94, 0xe0, 0xcc, 0xb8, 0x82, 0x82, 0x7b, 0x09, 0x0d, 0x23, 0x15,
	0x60, 0x06, 0x72, 0x55, 0xa2, 0x56, 0xbd, 0x03, 0xa6, 0x1a, 0x8a, 0xdf, 0x4a, 0x3d, 0x47, 0x9c,
	0xc7, 0x60, 0x08, 0x5c, 0x08, 0xd4, 0x5c, 0x47, 0x28, 0x54, 0xa5, 0x2e, 0x51, 0xf3, 0x6f, 0x96,
	0x9e, 0x39, 0x6f, 0x24, 0x4f, 0x12, 0x90, 0x15, 0xde, 0x9b, 0xe3, 0x1c, 0xcd, 0xd1, 0xc9, 0x4f,
	0x45, 0x75, 0xfc, 0xbd, 0x49, 0x99, 0x97, 0xf7, 0x38, 0x44, 0x71, 0x16, 0xcd, 0x53, 0x69, 0x3c,
	0x19, 0xa3, 0x2d, 0xc1, 0x97, 0xd0, 0x27, 0x23, 0x32, 0xd9, 0xf2, 0xcc, 0x99, 0x3d, 0xa3, 0x37,
	0xb2, 0x48, 0xea, 0x94, 0xc7, 0xfe, 0x02, 0x95, 0x56, 0xfd, 0xc6, 0xa8, 0x39, 0x69, 0xbb, 0x43,
	0xfb, 0xe2, 0x00, 0xec, 0x22, 0xf0, 0x69, 0x41, 0x3c, 0x42, 0xa5, 0xa7, 0xad, 0xf3, 0x2f, 0xc3,
	0x0d, 0xaf, 0x93, 0xfd, 0x2a, 0x29, 0xe6, 0xd2, 0x9d, 0x48, 0x68, 0x90, 0x82, 0xc7, 0x3e, 0x8a,
	0x78, 0xe5, 0x2f, 0x80, 0x87, 0x20, 0x55, 0xbf, 0x39, 0x6a, 0x4e, 0xb6, 0xbc, 0x6e, 0x05, 0x1e,
	0x8b, 0x78, 0x75, 0x54, 0x40, 0xec, 0x15, 0xdd, 0x93, 0xa0, 0x12, 0x14, 0x0a, 0x2a, 0xba, 0xaf,
	0xd1, 0xe7, 0x61, 0xd8, 0x6f, 0x99, 0x24, 0x77, 0xea, 0x49, 0xf2, 0xe1, 0xdb, 0x85, 0xf8, 0x94,
	0xc7, 0x29, 0x1c, 0x27, 0xf9, 0xd5, 0xbc, 0x5e, 0x65, 0x52, 0xfa, 0x3e, 0xc7, 0x87, 0x61, 0xc8,
	0x1e, 0xd0, 0xc1, 0xef, 0xcc, 0x25, 0x2c, 0x31, 0x83, 0xfe, 0x35, 0x93, 0x6a, 0xef, 0x8a, 0xd2,
	0x33, 0x30, 0x7b, 0x49, 0x77, 0x25, 0xbc, 0x4e, 0x41, 0xe9, 0xcb, 0xc1, 0x36, 0xff, 0x21, 0x58,
	0xb7, 0xf4, 0xa8, 0xe5, 0x7a, 0x42, 0xb7, 0x33, 0x1e, 0x47, 0x21, 0xd7, 0xe0, 0x07, 0x71, 0xaa,
	0x74, 0x3e, 0xa4, 0xeb, 0x23, 0x32, 0x69, 0xbb, 0x03, 0xbb, 0xd8, 0xb8, 0x5d, 0x6d, 0xdc, 0x9e,
	0x22, 0xc6, 0xc6, 0xd1, 0xbb, 0x59, 0x89, 0x0e, 0x4b, 0x8d, 0xfb, 0xa9, 0x41, 0x77, 0xcc, 0xa2,
	0x1f, 0x55, 0xaf, 0xee, 0x04, 0x64, 0x16, 0x05, 0xc0, 0x5e, 0xd0, 0xce, 0x89, 0x96, 0xc0, 0x97,
	0x06, 0x56, 0xcc, 0xaa, 0xa7, 0xfd, 0xc9, 0xf7, 0x8a, 0x78, 0x83, 0xe1, 0x1f, 0xf1, 0x62, 0x3a,
	0xe3, 0x8d, 0x09, 0xb9, 0x47, 0x58, 0x42, 0xb7, 0x9f, 0x8a, 0x40, 0xc2, 0x12, 0x84, 0xe6, 0x71,
	0xe9, 0x7d, 0xb7, 0xae, 0xbd, 0x40, 0xb8, 0xd2, 0x66, 0xff, 0x6f, 0xa8, 0xb5, 0x8e, 0x48, 0xdb,
	0x8f, 0x41, 0x07, 0x8b, 0xff, 0x75, 0x8f, 0xe1, 0xdb, 0xcf, 0xdf, 0xde, 0x37, 0x6e, 0x8d, 0x7b,
	0xb5, 0x8f, 0xf5, 0xc0, 0xbc, 0x70, 0x75, 0x40, 0xf6, 0xa7, 0xdd, 0x0f, 0x6b, 0x8b, 0x9c, 0xaf,
	0x2d, 0xf2, 0x71, 0x6d, 0x91, 0xaf, 0x6b, 0x8b, 0xbc, 0x23, 0x64, 0xb6, 0x69, 0xd6, 0x71, 0xff,
	0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xf0, 0xa4, 0xcb, 0x30, 0x04, 0x00, 0x00,
}