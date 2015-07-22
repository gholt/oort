// Code generated by protoc-gen-go.
// source: ort_api.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ort_api.proto

It has these top-level messages:
	FileRequest
	Attr
	File
	WriteResponse
	DirRequest
	DirEntries
	DirEnt
	FileEnt
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

// FileRequest is the file path we want to operate on
type FileRequest struct {
	Fpath string `protobuf:"bytes,1,opt,name=fpath" json:"fpath,omitempty"`
	Inode uint64 `protobuf:"varint,2,opt,name=inode" json:"inode,omitempty"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto1.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}

// Attr. fields are optional by default in proto3, so
// clients don't have to send all fields when performing an
// attr update for example. These might not all be needed
// but i got tired of constantly forgetting fields.
type Attr struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Inode  uint64 `protobuf:"varint,2,opt,name=inode" json:"inode,omitempty"`
	Atime  int64  `protobuf:"varint,3,opt,name=atime" json:"atime,omitempty"`
	Mtime  int64  `protobuf:"varint,4,opt,name=mtime" json:"mtime,omitempty"`
	Ctime  int64  `protobuf:"varint,5,opt,name=ctime" json:"ctime,omitempty"`
	Crtime int64  `protobuf:"varint,6,opt,name=crtime" json:"crtime,omitempty"`
	Mode   uint32 `protobuf:"varint,7,opt,name=mode" json:"mode,omitempty"`
	Valid  int32  `protobuf:"varint,9,opt,name=valid" json:"valid,omitempty"`
	Parent string `protobuf:"bytes,10,opt,name=parent" json:"parent,omitempty"`
	Size   uint64 `protobuf:"varint,11,opt,name=size" json:"size,omitempty"`
}

func (m *Attr) Reset()         { *m = Attr{} }
func (m *Attr) String() string { return proto1.CompactTextString(m) }
func (*Attr) ProtoMessage()    {}

// File contains the files name and its contents
type File struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Inode   uint64 `protobuf:"varint,2,opt,name=inode" json:"inode,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto1.CompactTextString(m) }
func (*File) ProtoMessage()    {}

// WriteRepsonse place holder. Maybe use an enum so
// we can map to fuse errors ?
type WriteResponse struct {
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto1.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}

// DirRequest is the dir we want to operate on
type DirRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Inode uint64 `protobuf:"varint,2,opt,name=inode" json:"inode,omitempty"`
}

func (m *DirRequest) Reset()         { *m = DirRequest{} }
func (m *DirRequest) String() string { return proto1.CompactTextString(m) }
func (*DirRequest) ProtoMessage()    {}

// DirEntries just contains a list of directory entries
type DirEntries struct {
	DirEntries  []*DirEnt  `protobuf:"bytes,1,rep" json:"DirEntries,omitempty"`
	FileEntries []*FileEnt `protobuf:"bytes,2,rep" json:"FileEntries,omitempty"`
}

func (m *DirEntries) Reset()         { *m = DirEntries{} }
func (m *DirEntries) String() string { return proto1.CompactTextString(m) }
func (*DirEntries) ProtoMessage()    {}

func (m *DirEntries) GetDirEntries() []*DirEnt {
	if m != nil {
		return m.DirEntries
	}
	return nil
}

func (m *DirEntries) GetFileEntries() []*FileEnt {
	if m != nil {
		return m.FileEntries
	}
	return nil
}

// DirEnt is a directory entry
type DirEnt struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Attr *Attr  `protobuf:"bytes,2,opt,name=attr" json:"attr,omitempty"`
}

func (m *DirEnt) Reset()         { *m = DirEnt{} }
func (m *DirEnt) String() string { return proto1.CompactTextString(m) }
func (*DirEnt) ProtoMessage()    {}

func (m *DirEnt) GetAttr() *Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

type FileEnt struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Attr *Attr  `protobuf:"bytes,2,opt,name=attr" json:"attr,omitempty"`
}

func (m *FileEnt) Reset()         { *m = FileEnt{} }
func (m *FileEnt) String() string { return proto1.CompactTextString(m) }
func (*FileEnt) ProtoMessage()    {}

func (m *FileEnt) GetAttr() *Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

// Client API for FileApi service

type FileApiClient interface {
	GetAttr(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*Attr, error)
	SetAttr(ctx context.Context, in *Attr, opts ...grpc.CallOption) (*Attr, error)
	Read(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*File, error)
	Write(ctx context.Context, in *File, opts ...grpc.CallOption) (*WriteResponse, error)
}

type fileApiClient struct {
	cc *grpc.ClientConn
}

func NewFileApiClient(cc *grpc.ClientConn) FileApiClient {
	return &fileApiClient{cc}
}

func (c *fileApiClient) GetAttr(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*Attr, error) {
	out := new(Attr)
	err := grpc.Invoke(ctx, "/proto.FileApi/GetAttr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileApiClient) SetAttr(ctx context.Context, in *Attr, opts ...grpc.CallOption) (*Attr, error) {
	out := new(Attr)
	err := grpc.Invoke(ctx, "/proto.FileApi/SetAttr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileApiClient) Read(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := grpc.Invoke(ctx, "/proto.FileApi/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileApiClient) Write(ctx context.Context, in *File, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/proto.FileApi/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileApi service

type FileApiServer interface {
	GetAttr(context.Context, *FileRequest) (*Attr, error)
	SetAttr(context.Context, *Attr) (*Attr, error)
	Read(context.Context, *FileRequest) (*File, error)
	Write(context.Context, *File) (*WriteResponse, error)
}

func RegisterFileApiServer(s *grpc.Server, srv FileApiServer) {
	s.RegisterService(&_FileApi_serviceDesc, srv)
}

func _FileApi_GetAttr_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(FileRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(FileApiServer).GetAttr(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileApi_SetAttr_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Attr)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(FileApiServer).SetAttr(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileApi_Read_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(FileRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(FileApiServer).Read(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FileApi_Write_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(File)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(FileApiServer).Write(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _FileApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FileApi",
	HandlerType: (*FileApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAttr",
			Handler:    _FileApi_GetAttr_Handler,
		},
		{
			MethodName: "SetAttr",
			Handler:    _FileApi_SetAttr_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _FileApi_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _FileApi_Write_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for DirApi service

type DirApiClient interface {
	GetAttr(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*Attr, error)
	MkDir(ctx context.Context, in *DirEnt, opts ...grpc.CallOption) (*DirEnt, error)
	Create(ctx context.Context, in *FileEnt, opts ...grpc.CallOption) (*FileEnt, error)
	Remove(ctx context.Context, in *DirEnt, opts ...grpc.CallOption) (*WriteResponse, error)
	Lookup(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*DirEnt, error)
	ReadDirAll(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*DirEntries, error)
}

type dirApiClient struct {
	cc *grpc.ClientConn
}

func NewDirApiClient(cc *grpc.ClientConn) DirApiClient {
	return &dirApiClient{cc}
}

func (c *dirApiClient) GetAttr(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*Attr, error) {
	out := new(Attr)
	err := grpc.Invoke(ctx, "/proto.DirApi/GetAttr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dirApiClient) MkDir(ctx context.Context, in *DirEnt, opts ...grpc.CallOption) (*DirEnt, error) {
	out := new(DirEnt)
	err := grpc.Invoke(ctx, "/proto.DirApi/MkDir", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dirApiClient) Create(ctx context.Context, in *FileEnt, opts ...grpc.CallOption) (*FileEnt, error) {
	out := new(FileEnt)
	err := grpc.Invoke(ctx, "/proto.DirApi/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dirApiClient) Remove(ctx context.Context, in *DirEnt, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/proto.DirApi/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dirApiClient) Lookup(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*DirEnt, error) {
	out := new(DirEnt)
	err := grpc.Invoke(ctx, "/proto.DirApi/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dirApiClient) ReadDirAll(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*DirEntries, error) {
	out := new(DirEntries)
	err := grpc.Invoke(ctx, "/proto.DirApi/ReadDirAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DirApi service

type DirApiServer interface {
	GetAttr(context.Context, *DirRequest) (*Attr, error)
	MkDir(context.Context, *DirEnt) (*DirEnt, error)
	Create(context.Context, *FileEnt) (*FileEnt, error)
	Remove(context.Context, *DirEnt) (*WriteResponse, error)
	Lookup(context.Context, *DirRequest) (*DirEnt, error)
	ReadDirAll(context.Context, *DirRequest) (*DirEntries, error)
}

func RegisterDirApiServer(s *grpc.Server, srv DirApiServer) {
	s.RegisterService(&_DirApi_serviceDesc, srv)
}

func _DirApi_GetAttr_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DirRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(DirApiServer).GetAttr(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _DirApi_MkDir_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DirEnt)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(DirApiServer).MkDir(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _DirApi_Create_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(FileEnt)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(DirApiServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _DirApi_Remove_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DirEnt)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(DirApiServer).Remove(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _DirApi_Lookup_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DirRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(DirApiServer).Lookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _DirApi_ReadDirAll_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DirRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(DirApiServer).ReadDirAll(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _DirApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DirApi",
	HandlerType: (*DirApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAttr",
			Handler:    _DirApi_GetAttr_Handler,
		},
		{
			MethodName: "MkDir",
			Handler:    _DirApi_MkDir_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DirApi_Create_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _DirApi_Remove_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _DirApi_Lookup_Handler,
		},
		{
			MethodName: "ReadDirAll",
			Handler:    _DirApi_ReadDirAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
