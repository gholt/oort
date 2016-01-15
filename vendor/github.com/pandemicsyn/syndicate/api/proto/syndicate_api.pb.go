// Code generated by protoc-gen-go.
// source: syndicate_api.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	syndicate_api.proto

It has these top-level messages:
	RingMsg
	StoreResult
	StatusRequest
	StatusMsg
	EmptyMsg
	RingStatus
	Node
	ModifyMsg
	RingConf
	Conf
	RegisterRequest
	NodeConfig
	Ring
	SearchResult
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RingMsg struct {
	Version  int64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Ring     []byte `protobuf:"bytes,2,opt,name=ring,proto3" json:"ring,omitempty"`
	Builder  []byte `protobuf:"bytes,3,opt,name=builder,proto3" json:"builder,omitempty"`
	Deadline int64  `protobuf:"varint,4,opt,name=deadline" json:"deadline,omitempty"`
	Rollback int64  `protobuf:"varint,5,opt,name=rollback" json:"rollback,omitempty"`
}

func (m *RingMsg) Reset()         { *m = RingMsg{} }
func (m *RingMsg) String() string { return proto1.CompactTextString(m) }
func (*RingMsg) ProtoMessage()    {}

type StoreResult struct {
	Version int64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Ring    bool   `protobuf:"varint,2,opt,name=ring" json:"ring,omitempty"`
	Builder bool   `protobuf:"varint,3,opt,name=builder" json:"builder,omitempty"`
	ErrMsg  string `protobuf:"bytes,4,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
}

func (m *StoreResult) Reset()         { *m = StoreResult{} }
func (m *StoreResult) String() string { return proto1.CompactTextString(m) }
func (*StoreResult) ProtoMessage()    {}

type StatusRequest struct {
	Ring    bool `protobuf:"varint,1,opt,name=ring" json:"ring,omitempty"`
	Builder bool `protobuf:"varint,2,opt,name=builder" json:"builder,omitempty"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto1.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}

type StatusMsg struct {
	Version      int64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Ringstats    string `protobuf:"bytes,2,opt,name=ringstats" json:"ringstats,omitempty"`
	Builderstats string `protobuf:"bytes,3,opt,name=builderstats" json:"builderstats,omitempty"`
	Master       string `protobuf:"bytes,4,opt,name=master" json:"master,omitempty"`
}

func (m *StatusMsg) Reset()         { *m = StatusMsg{} }
func (m *StatusMsg) String() string { return proto1.CompactTextString(m) }
func (*StatusMsg) ProtoMessage()    {}

type EmptyMsg struct {
}

func (m *EmptyMsg) Reset()         { *m = EmptyMsg{} }
func (m *EmptyMsg) String() string { return proto1.CompactTextString(m) }
func (*EmptyMsg) ProtoMessage()    {}

type RingStatus struct {
	Status  bool  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Version int64 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *RingStatus) Reset()         { *m = RingStatus{} }
func (m *RingStatus) String() string { return proto1.CompactTextString(m) }
func (*RingStatus) ProtoMessage()    {}

type Node struct {
	Id        uint64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Active    bool     `protobuf:"varint,2,opt,name=active" json:"active,omitempty"`
	Capacity  uint32   `protobuf:"varint,3,opt,name=capacity" json:"capacity,omitempty"`
	Tiers     []string `protobuf:"bytes,4,rep,name=tiers" json:"tiers,omitempty"`
	Addresses []string `protobuf:"bytes,5,rep,name=addresses" json:"addresses,omitempty"`
	Meta      string   `protobuf:"bytes,6,opt,name=meta" json:"meta,omitempty"`
	Conf      []byte   `protobuf:"bytes,7,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto1.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

type ModifyMsg struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Id    uint64 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
}

func (m *ModifyMsg) Reset()         { *m = ModifyMsg{} }
func (m *ModifyMsg) String() string { return proto1.CompactTextString(m) }
func (*ModifyMsg) ProtoMessage()    {}

type RingConf struct {
	Status *RingStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Conf   *Conf       `protobuf:"bytes,2,opt,name=conf" json:"conf,omitempty"`
}

func (m *RingConf) Reset()         { *m = RingConf{} }
func (m *RingConf) String() string { return proto1.CompactTextString(m) }
func (*RingConf) ProtoMessage()    {}

func (m *RingConf) GetStatus() *RingStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *RingConf) GetConf() *Conf {
	if m != nil {
		return m.Conf
	}
	return nil
}

type Conf struct {
	Conf            []byte `protobuf:"bytes,1,opt,name=conf,proto3" json:"conf,omitempty"`
	RestartRequired bool   `protobuf:"varint,2,opt,name=restartRequired" json:"restartRequired,omitempty"`
}

func (m *Conf) Reset()         { *m = Conf{} }
func (m *Conf) String() string { return proto1.CompactTextString(m) }
func (*Conf) ProtoMessage()    {}

type RegisterRequest struct {
	Hostname   string   `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Addrs      []string `protobuf:"bytes,2,rep,name=addrs" json:"addrs,omitempty"`
	Tiers      []string `protobuf:"bytes,3,rep,name=tiers" json:"tiers,omitempty"`
	Disks      int32    `protobuf:"varint,4,opt,name=disks" json:"disks,omitempty"`
	Cores      int32    `protobuf:"varint,5,opt,name=cores" json:"cores,omitempty"`
	Hardwareid string   `protobuf:"bytes,6,opt,name=hardwareid" json:"hardwareid,omitempty"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto1.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}

type NodeConfig struct {
	Localid uint64 `protobuf:"varint,1,opt,name=localid" json:"localid,omitempty"`
	Ring    []byte `protobuf:"bytes,2,opt,name=ring,proto3" json:"ring,omitempty"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto1.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}

type Ring struct {
	Version uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Ring    []byte `protobuf:"bytes,2,opt,name=ring,proto3" json:"ring,omitempty"`
}

func (m *Ring) Reset()         { *m = Ring{} }
func (m *Ring) String() string { return proto1.CompactTextString(m) }
func (*Ring) ProtoMessage()    {}

type SearchResult struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *SearchResult) Reset()         { *m = SearchResult{} }
func (m *SearchResult) String() string { return proto1.CompactTextString(m) }
func (*SearchResult) ProtoMessage()    {}

func (m *SearchResult) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto1.RegisterType((*RingMsg)(nil), "proto.RingMsg")
	proto1.RegisterType((*StoreResult)(nil), "proto.StoreResult")
	proto1.RegisterType((*StatusRequest)(nil), "proto.StatusRequest")
	proto1.RegisterType((*StatusMsg)(nil), "proto.StatusMsg")
	proto1.RegisterType((*EmptyMsg)(nil), "proto.EmptyMsg")
	proto1.RegisterType((*RingStatus)(nil), "proto.RingStatus")
	proto1.RegisterType((*Node)(nil), "proto.Node")
	proto1.RegisterType((*ModifyMsg)(nil), "proto.ModifyMsg")
	proto1.RegisterType((*RingConf)(nil), "proto.RingConf")
	proto1.RegisterType((*Conf)(nil), "proto.Conf")
	proto1.RegisterType((*RegisterRequest)(nil), "proto.RegisterRequest")
	proto1.RegisterType((*NodeConfig)(nil), "proto.NodeConfig")
	proto1.RegisterType((*Ring)(nil), "proto.Ring")
	proto1.RegisterType((*SearchResult)(nil), "proto.SearchResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for RingDist service

type RingDistClient interface {
	Store(ctx context.Context, in *RingMsg, opts ...grpc.CallOption) (*StoreResult, error)
	Revert(ctx context.Context, in *RingMsg, opts ...grpc.CallOption) (*StoreResult, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusMsg, error)
	Setup(ctx context.Context, in *RingMsg, opts ...grpc.CallOption) (*StoreResult, error)
}

type ringDistClient struct {
	cc *grpc.ClientConn
}

func NewRingDistClient(cc *grpc.ClientConn) RingDistClient {
	return &ringDistClient{cc}
}

func (c *ringDistClient) Store(ctx context.Context, in *RingMsg, opts ...grpc.CallOption) (*StoreResult, error) {
	out := new(StoreResult)
	err := grpc.Invoke(ctx, "/proto.RingDist/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ringDistClient) Revert(ctx context.Context, in *RingMsg, opts ...grpc.CallOption) (*StoreResult, error) {
	out := new(StoreResult)
	err := grpc.Invoke(ctx, "/proto.RingDist/Revert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ringDistClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusMsg, error) {
	out := new(StatusMsg)
	err := grpc.Invoke(ctx, "/proto.RingDist/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ringDistClient) Setup(ctx context.Context, in *RingMsg, opts ...grpc.CallOption) (*StoreResult, error) {
	out := new(StoreResult)
	err := grpc.Invoke(ctx, "/proto.RingDist/Setup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RingDist service

type RingDistServer interface {
	Store(context.Context, *RingMsg) (*StoreResult, error)
	Revert(context.Context, *RingMsg) (*StoreResult, error)
	Status(context.Context, *StatusRequest) (*StatusMsg, error)
	Setup(context.Context, *RingMsg) (*StoreResult, error)
}

func RegisterRingDistServer(s *grpc.Server, srv RingDistServer) {
	s.RegisterService(&_RingDist_serviceDesc, srv)
}

func _RingDist_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RingDistServer).Store(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RingDist_Revert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RingDistServer).Revert(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RingDist_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RingDistServer).Status(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RingDist_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RingDistServer).Setup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _RingDist_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RingDist",
	HandlerType: (*RingDistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _RingDist_Store_Handler,
		},
		{
			MethodName: "Revert",
			Handler:    _RingDist_Revert_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _RingDist_Status_Handler,
		},
		{
			MethodName: "Setup",
			Handler:    _RingDist_Setup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Syndicate service

type SyndicateClient interface {
	AddNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*RingStatus, error)
	RemoveNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*RingStatus, error)
	ModNode(ctx context.Context, in *ModifyMsg, opts ...grpc.CallOption) (*RingStatus, error)
	SetConf(ctx context.Context, in *Conf, opts ...grpc.CallOption) (*RingStatus, error)
	SetActive(ctx context.Context, in *Node, opts ...grpc.CallOption) (*RingStatus, error)
	GetVersion(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*RingStatus, error)
	GetGlobalConfig(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*RingConf, error)
	GetNodeConfig(ctx context.Context, in *Node, opts ...grpc.CallOption) (*RingConf, error)
	SearchNodes(ctx context.Context, in *Node, opts ...grpc.CallOption) (*SearchResult, error)
	GetRing(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*Ring, error)
	RegisterNode(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*NodeConfig, error)
}

type syndicateClient struct {
	cc *grpc.ClientConn
}

func NewSyndicateClient(cc *grpc.ClientConn) SyndicateClient {
	return &syndicateClient{cc}
}

func (c *syndicateClient) AddNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*RingStatus, error) {
	out := new(RingStatus)
	err := grpc.Invoke(ctx, "/proto.Syndicate/AddNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) RemoveNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*RingStatus, error) {
	out := new(RingStatus)
	err := grpc.Invoke(ctx, "/proto.Syndicate/RemoveNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) ModNode(ctx context.Context, in *ModifyMsg, opts ...grpc.CallOption) (*RingStatus, error) {
	out := new(RingStatus)
	err := grpc.Invoke(ctx, "/proto.Syndicate/ModNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) SetConf(ctx context.Context, in *Conf, opts ...grpc.CallOption) (*RingStatus, error) {
	out := new(RingStatus)
	err := grpc.Invoke(ctx, "/proto.Syndicate/SetConf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) SetActive(ctx context.Context, in *Node, opts ...grpc.CallOption) (*RingStatus, error) {
	out := new(RingStatus)
	err := grpc.Invoke(ctx, "/proto.Syndicate/SetActive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) GetVersion(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*RingStatus, error) {
	out := new(RingStatus)
	err := grpc.Invoke(ctx, "/proto.Syndicate/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) GetGlobalConfig(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*RingConf, error) {
	out := new(RingConf)
	err := grpc.Invoke(ctx, "/proto.Syndicate/GetGlobalConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) GetNodeConfig(ctx context.Context, in *Node, opts ...grpc.CallOption) (*RingConf, error) {
	out := new(RingConf)
	err := grpc.Invoke(ctx, "/proto.Syndicate/GetNodeConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) SearchNodes(ctx context.Context, in *Node, opts ...grpc.CallOption) (*SearchResult, error) {
	out := new(SearchResult)
	err := grpc.Invoke(ctx, "/proto.Syndicate/SearchNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) GetRing(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*Ring, error) {
	out := new(Ring)
	err := grpc.Invoke(ctx, "/proto.Syndicate/GetRing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syndicateClient) RegisterNode(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*NodeConfig, error) {
	out := new(NodeConfig)
	err := grpc.Invoke(ctx, "/proto.Syndicate/RegisterNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Syndicate service

type SyndicateServer interface {
	AddNode(context.Context, *Node) (*RingStatus, error)
	RemoveNode(context.Context, *Node) (*RingStatus, error)
	ModNode(context.Context, *ModifyMsg) (*RingStatus, error)
	SetConf(context.Context, *Conf) (*RingStatus, error)
	SetActive(context.Context, *Node) (*RingStatus, error)
	GetVersion(context.Context, *EmptyMsg) (*RingStatus, error)
	GetGlobalConfig(context.Context, *EmptyMsg) (*RingConf, error)
	GetNodeConfig(context.Context, *Node) (*RingConf, error)
	SearchNodes(context.Context, *Node) (*SearchResult, error)
	GetRing(context.Context, *EmptyMsg) (*Ring, error)
	RegisterNode(context.Context, *RegisterRequest) (*NodeConfig, error)
}

func RegisterSyndicateServer(s *grpc.Server, srv SyndicateServer) {
	s.RegisterService(&_Syndicate_serviceDesc, srv)
}

func _Syndicate_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).AddNode(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).RemoveNode(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_ModNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ModifyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).ModNode(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_SetConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Conf)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).SetConf(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_SetActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).SetActive(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).GetVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_GetGlobalConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).GetGlobalConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_GetNodeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).GetNodeConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_SearchNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).SearchNodes(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_GetRing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).GetRing(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Syndicate_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SyndicateServer).RegisterNode(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Syndicate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Syndicate",
	HandlerType: (*SyndicateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNode",
			Handler:    _Syndicate_AddNode_Handler,
		},
		{
			MethodName: "RemoveNode",
			Handler:    _Syndicate_RemoveNode_Handler,
		},
		{
			MethodName: "ModNode",
			Handler:    _Syndicate_ModNode_Handler,
		},
		{
			MethodName: "SetConf",
			Handler:    _Syndicate_SetConf_Handler,
		},
		{
			MethodName: "SetActive",
			Handler:    _Syndicate_SetActive_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Syndicate_GetVersion_Handler,
		},
		{
			MethodName: "GetGlobalConfig",
			Handler:    _Syndicate_GetGlobalConfig_Handler,
		},
		{
			MethodName: "GetNodeConfig",
			Handler:    _Syndicate_GetNodeConfig_Handler,
		},
		{
			MethodName: "SearchNodes",
			Handler:    _Syndicate_SearchNodes_Handler,
		},
		{
			MethodName: "GetRing",
			Handler:    _Syndicate_GetRing_Handler,
		},
		{
			MethodName: "RegisterNode",
			Handler:    _Syndicate_RegisterNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}