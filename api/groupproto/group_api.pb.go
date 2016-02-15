// Code generated by protoc-gen-go.
// source: group_api.proto
// DO NOT EDIT!

/*
Package groupproto is a generated protocol buffer package.

It is generated from these files:
	group_api.proto

It has these top-level messages:
	EmptyMsg
	WriteRequest
	LookupRequest
	ReadRequest
	DeleteRequest
	LookupGroupRequest
	ReadGroupRequest
	WriteResponse
	LookupResponse
	LookupGroupResponse
	LookupGroupItem
	ReadGroupResponse
	ReadGroupItem
	ReadResponse
	DeleteResponse
*/
package groupproto

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

type EmptyMsg struct {
}

func (m *EmptyMsg) Reset()         { *m = EmptyMsg{} }
func (m *EmptyMsg) String() string { return proto.CompactTextString(m) }
func (*EmptyMsg) ProtoMessage()    {}

type WriteRequest struct {
	KeyA      uint64 `protobuf:"varint,1,opt,name=keyA" json:"keyA,omitempty"`
	KeyB      uint64 `protobuf:"varint,2,opt,name=keyB" json:"keyB,omitempty"`
	ChildKeyA uint64 `protobuf:"varint,3,opt,name=childKeyA" json:"childKeyA,omitempty"`
	ChildKeyB uint64 `protobuf:"varint,4,opt,name=childKeyB" json:"childKeyB,omitempty"`
	Value     []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Tsm       int64  `protobuf:"varint,6,opt,name=tsm" json:"tsm,omitempty"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}

type LookupRequest struct {
	KeyA      uint64 `protobuf:"varint,1,opt,name=keyA" json:"keyA,omitempty"`
	KeyB      uint64 `protobuf:"varint,2,opt,name=keyB" json:"keyB,omitempty"`
	ChildKeyA uint64 `protobuf:"varint,3,opt,name=childKeyA" json:"childKeyA,omitempty"`
	ChildKeyB uint64 `protobuf:"varint,4,opt,name=childKeyB" json:"childKeyB,omitempty"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}

type ReadRequest struct {
	KeyA      uint64 `protobuf:"varint,1,opt,name=keyA" json:"keyA,omitempty"`
	KeyB      uint64 `protobuf:"varint,2,opt,name=keyB" json:"keyB,omitempty"`
	ChildKeyA uint64 `protobuf:"varint,3,opt,name=childKeyA" json:"childKeyA,omitempty"`
	ChildKeyB uint64 `protobuf:"varint,4,opt,name=childKeyB" json:"childKeyB,omitempty"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}

type DeleteRequest struct {
	KeyA      uint64 `protobuf:"varint,1,opt,name=keyA" json:"keyA,omitempty"`
	KeyB      uint64 `protobuf:"varint,2,opt,name=keyB" json:"keyB,omitempty"`
	ChildKeyA uint64 `protobuf:"varint,3,opt,name=childKeyA" json:"childKeyA,omitempty"`
	ChildKeyB uint64 `protobuf:"varint,4,opt,name=childKeyB" json:"childKeyB,omitempty"`
	Tsm       int64  `protobuf:"varint,5,opt,name=tsm" json:"tsm,omitempty"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}

type LookupGroupRequest struct {
	A uint64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B uint64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *LookupGroupRequest) Reset()         { *m = LookupGroupRequest{} }
func (m *LookupGroupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupGroupRequest) ProtoMessage()    {}

type ReadGroupRequest struct {
	A uint64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B uint64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *ReadGroupRequest) Reset()         { *m = ReadGroupRequest{} }
func (m *ReadGroupRequest) String() string { return proto.CompactTextString(m) }
func (*ReadGroupRequest) ProtoMessage()    {}

type WriteResponse struct {
	Tsm int64  `protobuf:"varint,1,opt,name=tsm" json:"tsm,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}

type LookupResponse struct {
	Tsm    int64  `protobuf:"varint,1,opt,name=tsm" json:"tsm,omitempty"`
	Length uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Err    string `protobuf:"bytes,3,opt,name=err" json:"err,omitempty"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}

type LookupGroupResponse struct {
	Items []*LookupGroupItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *LookupGroupResponse) Reset()         { *m = LookupGroupResponse{} }
func (m *LookupGroupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupGroupResponse) ProtoMessage()    {}

func (m *LookupGroupResponse) GetItems() []*LookupGroupItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type LookupGroupItem struct {
	ChildKeyA      uint64 `protobuf:"varint,1,opt,name=childKeyA" json:"childKeyA,omitempty"`
	ChildKeyB      uint64 `protobuf:"varint,2,opt,name=childKeyB" json:"childKeyB,omitempty"`
	TimestampMicro int64  `protobuf:"varint,3,opt,name=timestampMicro" json:"timestampMicro,omitempty"`
	Length         uint32 `protobuf:"varint,4,opt,name=length" json:"length,omitempty"`
}

func (m *LookupGroupItem) Reset()         { *m = LookupGroupItem{} }
func (m *LookupGroupItem) String() string { return proto.CompactTextString(m) }
func (*LookupGroupItem) ProtoMessage()    {}

type ReadGroupResponse struct {
	Items []*ReadGroupItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ReadGroupResponse) Reset()         { *m = ReadGroupResponse{} }
func (m *ReadGroupResponse) String() string { return proto.CompactTextString(m) }
func (*ReadGroupResponse) ProtoMessage()    {}

func (m *ReadGroupResponse) GetItems() []*ReadGroupItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ReadGroupItem struct {
	ChildKeyA      uint64 `protobuf:"varint,1,opt,name=childKeyA" json:"childKeyA,omitempty"`
	ChildKeyB      uint64 `protobuf:"varint,2,opt,name=childKeyB" json:"childKeyB,omitempty"`
	TimestampMicro int64  `protobuf:"varint,3,opt,name=timestampMicro" json:"timestampMicro,omitempty"`
	Value          []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ReadGroupItem) Reset()         { *m = ReadGroupItem{} }
func (m *ReadGroupItem) String() string { return proto.CompactTextString(m) }
func (*ReadGroupItem) ProtoMessage()    {}

type ReadResponse struct {
	Tsm   int64  `protobuf:"varint,1,opt,name=tsm" json:"tsm,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Err   string `protobuf:"bytes,3,opt,name=err" json:"err,omitempty"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}

type DeleteResponse struct {
	Tsm int64  `protobuf:"varint,1,opt,name=tsm" json:"tsm,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GroupStore service

type GroupStoreClient interface {
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	StreamWrite(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamWriteClient, error)
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	StreamLookup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamLookupClient, error)
	LookupGroup(ctx context.Context, in *LookupGroupRequest, opts ...grpc.CallOption) (*LookupGroupResponse, error)
	StreamLookupGroup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamLookupGroupClient, error)
	ReadGroup(ctx context.Context, in *ReadGroupRequest, opts ...grpc.CallOption) (*ReadGroupResponse, error)
	StreamReadGroup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamReadGroupClient, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	StreamRead(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamReadClient, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	StreamDelete(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamDeleteClient, error)
}

type groupStoreClient struct {
	cc *grpc.ClientConn
}

func NewGroupStoreClient(cc *grpc.ClientConn) GroupStoreClient {
	return &groupStoreClient{cc}
}

func (c *groupStoreClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) StreamWrite(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamWriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[0], c.cc, "/groupproto.GroupStore/StreamWrite", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreStreamWriteClient{stream}
	return x, nil
}

type GroupStore_StreamWriteClient interface {
	Send(*WriteRequest) error
	Recv() (*WriteResponse, error)
	grpc.ClientStream
}

type groupStoreStreamWriteClient struct {
	grpc.ClientStream
}

func (x *groupStoreStreamWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreStreamWriteClient) Recv() (*WriteResponse, error) {
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) StreamLookup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamLookupClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[1], c.cc, "/groupproto.GroupStore/StreamLookup", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreStreamLookupClient{stream}
	return x, nil
}

type GroupStore_StreamLookupClient interface {
	Send(*LookupRequest) error
	Recv() (*LookupResponse, error)
	grpc.ClientStream
}

type groupStoreStreamLookupClient struct {
	grpc.ClientStream
}

func (x *groupStoreStreamLookupClient) Send(m *LookupRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreStreamLookupClient) Recv() (*LookupResponse, error) {
	m := new(LookupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) LookupGroup(ctx context.Context, in *LookupGroupRequest, opts ...grpc.CallOption) (*LookupGroupResponse, error) {
	out := new(LookupGroupResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/LookupGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) StreamLookupGroup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamLookupGroupClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[2], c.cc, "/groupproto.GroupStore/StreamLookupGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreStreamLookupGroupClient{stream}
	return x, nil
}

type GroupStore_StreamLookupGroupClient interface {
	Send(*LookupGroupRequest) error
	Recv() (*LookupGroupResponse, error)
	grpc.ClientStream
}

type groupStoreStreamLookupGroupClient struct {
	grpc.ClientStream
}

func (x *groupStoreStreamLookupGroupClient) Send(m *LookupGroupRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreStreamLookupGroupClient) Recv() (*LookupGroupResponse, error) {
	m := new(LookupGroupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) ReadGroup(ctx context.Context, in *ReadGroupRequest, opts ...grpc.CallOption) (*ReadGroupResponse, error) {
	out := new(ReadGroupResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/ReadGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) StreamReadGroup(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamReadGroupClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[3], c.cc, "/groupproto.GroupStore/StreamReadGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreStreamReadGroupClient{stream}
	return x, nil
}

type GroupStore_StreamReadGroupClient interface {
	Send(*ReadGroupRequest) error
	Recv() (*ReadGroupResponse, error)
	grpc.ClientStream
}

type groupStoreStreamReadGroupClient struct {
	grpc.ClientStream
}

func (x *groupStoreStreamReadGroupClient) Send(m *ReadGroupRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreStreamReadGroupClient) Recv() (*ReadGroupResponse, error) {
	m := new(ReadGroupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) StreamRead(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[4], c.cc, "/groupproto.GroupStore/StreamRead", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreStreamReadClient{stream}
	return x, nil
}

type GroupStore_StreamReadClient interface {
	Send(*ReadRequest) error
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type groupStoreStreamReadClient struct {
	grpc.ClientStream
}

func (x *groupStoreStreamReadClient) Send(m *ReadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreStreamReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groupStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) StreamDelete(ctx context.Context, opts ...grpc.CallOption) (GroupStore_StreamDeleteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroupStore_serviceDesc.Streams[5], c.cc, "/groupproto.GroupStore/StreamDelete", opts...)
	if err != nil {
		return nil, err
	}
	x := &groupStoreStreamDeleteClient{stream}
	return x, nil
}

type GroupStore_StreamDeleteClient interface {
	Send(*DeleteRequest) error
	Recv() (*DeleteResponse, error)
	grpc.ClientStream
}

type groupStoreStreamDeleteClient struct {
	grpc.ClientStream
}

func (x *groupStoreStreamDeleteClient) Send(m *DeleteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *groupStoreStreamDeleteClient) Recv() (*DeleteResponse, error) {
	m := new(DeleteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GroupStore service

type GroupStoreServer interface {
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	StreamWrite(GroupStore_StreamWriteServer) error
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	StreamLookup(GroupStore_StreamLookupServer) error
	LookupGroup(context.Context, *LookupGroupRequest) (*LookupGroupResponse, error)
	StreamLookupGroup(GroupStore_StreamLookupGroupServer) error
	ReadGroup(context.Context, *ReadGroupRequest) (*ReadGroupResponse, error)
	StreamReadGroup(GroupStore_StreamReadGroupServer) error
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	StreamRead(GroupStore_StreamReadServer) error
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	StreamDelete(GroupStore_StreamDeleteServer) error
}

func RegisterGroupStoreServer(s *grpc.Server, srv GroupStoreServer) {
	s.RegisterService(&_GroupStore_serviceDesc, srv)
}

func _GroupStore_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).Write(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_StreamWrite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).StreamWrite(&groupStoreStreamWriteServer{stream})
}

type GroupStore_StreamWriteServer interface {
	Send(*WriteResponse) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type groupStoreStreamWriteServer struct {
	grpc.ServerStream
}

func (x *groupStoreStreamWriteServer) Send(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreStreamWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).Lookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_StreamLookup_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).StreamLookup(&groupStoreStreamLookupServer{stream})
}

type GroupStore_StreamLookupServer interface {
	Send(*LookupResponse) error
	Recv() (*LookupRequest, error)
	grpc.ServerStream
}

type groupStoreStreamLookupServer struct {
	grpc.ServerStream
}

func (x *groupStoreStreamLookupServer) Send(m *LookupResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreStreamLookupServer) Recv() (*LookupRequest, error) {
	m := new(LookupRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_LookupGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LookupGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).LookupGroup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_StreamLookupGroup_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).StreamLookupGroup(&groupStoreStreamLookupGroupServer{stream})
}

type GroupStore_StreamLookupGroupServer interface {
	Send(*LookupGroupResponse) error
	Recv() (*LookupGroupRequest, error)
	grpc.ServerStream
}

type groupStoreStreamLookupGroupServer struct {
	grpc.ServerStream
}

func (x *groupStoreStreamLookupGroupServer) Send(m *LookupGroupResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreStreamLookupGroupServer) Recv() (*LookupGroupRequest, error) {
	m := new(LookupGroupRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_ReadGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ReadGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).ReadGroup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_StreamReadGroup_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).StreamReadGroup(&groupStoreStreamReadGroupServer{stream})
}

type GroupStore_StreamReadGroupServer interface {
	Send(*ReadGroupResponse) error
	Recv() (*ReadGroupRequest, error)
	grpc.ServerStream
}

type groupStoreStreamReadGroupServer struct {
	grpc.ServerStream
}

func (x *groupStoreStreamReadGroupServer) Send(m *ReadGroupResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreStreamReadGroupServer) Recv() (*ReadGroupRequest, error) {
	m := new(ReadGroupRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).Read(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_StreamRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).StreamRead(&groupStoreStreamReadServer{stream})
}

type GroupStore_StreamReadServer interface {
	Send(*ReadResponse) error
	Recv() (*ReadRequest, error)
	grpc.ServerStream
}

type groupStoreStreamReadServer struct {
	grpc.ServerStream
}

func (x *groupStoreStreamReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreStreamReadServer) Recv() (*ReadRequest, error) {
	m := new(ReadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GroupStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_StreamDelete_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GroupStoreServer).StreamDelete(&groupStoreStreamDeleteServer{stream})
}

type GroupStore_StreamDeleteServer interface {
	Send(*DeleteResponse) error
	Recv() (*DeleteRequest, error)
	grpc.ServerStream
}

type groupStoreStreamDeleteServer struct {
	grpc.ServerStream
}

func (x *groupStoreStreamDeleteServer) Send(m *DeleteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *groupStoreStreamDeleteServer) Recv() (*DeleteRequest, error) {
	m := new(DeleteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GroupStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "groupproto.GroupStore",
	HandlerType: (*GroupStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _GroupStore_Write_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _GroupStore_Lookup_Handler,
		},
		{
			MethodName: "LookupGroup",
			Handler:    _GroupStore_LookupGroup_Handler,
		},
		{
			MethodName: "ReadGroup",
			Handler:    _GroupStore_ReadGroup_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _GroupStore_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GroupStore_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamWrite",
			Handler:       _GroupStore_StreamWrite_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamLookup",
			Handler:       _GroupStore_StreamLookup_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamLookupGroup",
			Handler:       _GroupStore_StreamLookupGroup_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamReadGroup",
			Handler:       _GroupStore_StreamReadGroup_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamRead",
			Handler:       _GroupStore_StreamRead_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamDelete",
			Handler:       _GroupStore_StreamDelete_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
