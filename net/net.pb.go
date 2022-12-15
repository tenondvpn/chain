// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: net.proto

package net

import (
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

type BroadcastParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type              *uint32  `protobuf:"varint,1,opt,name=Type,proto3,oneof" json:"Type,omitempty"`
	NeighborCount     *uint32  `protobuf:"varint,2,opt,name=NeighborCount,proto3,oneof" json:"NeighborCount,omitempty"`
	StopTimes         *uint32  `protobuf:"varint,3,opt,name=StopTimes,proto3,oneof" json:"StopTimes,omitempty"`
	HopLimit          *uint32  `protobuf:"varint,4,opt,name=HopLimit,proto3,oneof" json:"HopLimit,omitempty"`
	LayerLeft         *uint64  `protobuf:"varint,5,opt,name=LayerLeft,proto3,oneof" json:"LayerLeft,omitempty"`
	LayerRight        *uint64  `protobuf:"varint,6,opt,name=LayerRight,proto3,oneof" json:"LayerRight,omitempty"`
	Overlap           *float32 `protobuf:"fixed32,7,opt,name=Overlap,proto3,oneof" json:"Overlap,omitempty"`
	HopToLayer        *uint32  `protobuf:"varint,8,opt,name=HopToLayer,proto3,oneof" json:"HopToLayer,omitempty"`
	Header            []byte   `protobuf:"bytes,9,opt,name=Header,proto3,oneof" json:"Header,omitempty"`
	Body              []byte   `protobuf:"bytes,10,opt,name=Body,proto3,oneof" json:"Body,omitempty"`
	NetCrossed        *bool    `protobuf:"varint,11,opt,name=NetCrossed,proto3,oneof" json:"NetCrossed,omitempty"`
	Bloomfilter       []uint64 `protobuf:"varint,12,rep,packed,name=Bloomfilter,proto3" json:"Bloomfilter,omitempty"`
	EvilRate          *float32 `protobuf:"fixed32,13,opt,name=EvilRate,proto3,oneof" json:"EvilRate,omitempty"`
	IgnBloomfilterHop *uint32  `protobuf:"varint,14,opt,name=IgnBloomfilterHop,proto3,oneof" json:"IgnBloomfilterHop,omitempty"`
}

func (x *BroadcastParam) Reset() {
	*x = BroadcastParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastParam) ProtoMessage() {}

func (x *BroadcastParam) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastParam.ProtoReflect.Descriptor instead.
func (*BroadcastParam) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{0}
}

func (x *BroadcastParam) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *BroadcastParam) GetNeighborCount() uint32 {
	if x != nil && x.NeighborCount != nil {
		return *x.NeighborCount
	}
	return 0
}

func (x *BroadcastParam) GetStopTimes() uint32 {
	if x != nil && x.StopTimes != nil {
		return *x.StopTimes
	}
	return 0
}

func (x *BroadcastParam) GetHopLimit() uint32 {
	if x != nil && x.HopLimit != nil {
		return *x.HopLimit
	}
	return 0
}

func (x *BroadcastParam) GetLayerLeft() uint64 {
	if x != nil && x.LayerLeft != nil {
		return *x.LayerLeft
	}
	return 0
}

func (x *BroadcastParam) GetLayerRight() uint64 {
	if x != nil && x.LayerRight != nil {
		return *x.LayerRight
	}
	return 0
}

func (x *BroadcastParam) GetOverlap() float32 {
	if x != nil && x.Overlap != nil {
		return *x.Overlap
	}
	return 0
}

func (x *BroadcastParam) GetHopToLayer() uint32 {
	if x != nil && x.HopToLayer != nil {
		return *x.HopToLayer
	}
	return 0
}

func (x *BroadcastParam) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BroadcastParam) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *BroadcastParam) GetNetCrossed() bool {
	if x != nil && x.NetCrossed != nil {
		return *x.NetCrossed
	}
	return false
}

func (x *BroadcastParam) GetBloomfilter() []uint64 {
	if x != nil {
		return x.Bloomfilter
	}
	return nil
}

func (x *BroadcastParam) GetEvilRate() float32 {
	if x != nil && x.EvilRate != nil {
		return *x.EvilRate
	}
	return 0
}

func (x *BroadcastParam) GetIgnBloomfilterHop() uint32 {
	if x != nil && x.IgnBloomfilterHop != nil {
		return *x.IgnBloomfilterHop
	}
	return 0
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcDhtKey     []byte          `protobuf:"bytes,1,opt,name=SrcDhtKey,proto3,oneof" json:"SrcDhtKey,omitempty"`
	DesDhtKey     []byte          `protobuf:"bytes,2,opt,name=DesDhtKey,proto3,oneof" json:"DesDhtKey,omitempty"`
	SrcNodeId     []byte          `protobuf:"bytes,3,opt,name=SrcNodeId,proto3,oneof" json:"SrcNodeId,omitempty"`
	DesNodeId     []byte          `protobuf:"bytes,4,opt,name=DesNodeId,proto3,oneof" json:"DesNodeId,omitempty"`
	HopCount      *uint32         `protobuf:"varint,5,opt,name=HopCount,proto3,oneof" json:"HopCount,omitempty"`
	Data          []byte          `protobuf:"bytes,6,opt,name=Data,proto3,oneof" json:"Data,omitempty"`
	Priority      *uint32         `protobuf:"varint,7,opt,name=Priority,proto3,oneof" json:"Priority,omitempty"`
	Debug         []byte          `protobuf:"bytes,8,opt,name=Debug,proto3,oneof" json:"Debug,omitempty"`
	FromIp        []byte          `protobuf:"bytes,9,opt,name=FromIp,proto3,oneof" json:"FromIp,omitempty"`       // will clear before transport send, not use
	FromPort      *uint32         `protobuf:"varint,10,opt,name=FromPort,proto3,oneof" json:"FromPort,omitempty"` // will clear before transport send, not use
	ToIp          []byte          `protobuf:"bytes,11,opt,name=ToIp,proto3,oneof" json:"ToIp,omitempty"`          // will clear before transport send, not use
	ToPort        *uint32         `protobuf:"varint,12,opt,name=ToPort,proto3,oneof" json:"ToPort,omitempty"`     // will clear before transport send, not use
	Id            *uint32         `protobuf:"varint,13,opt,name=Id,proto3,oneof" json:"Id,omitempty"`
	Hash          *uint64         `protobuf:"varint,14,opt,name=Hash,proto3,oneof" json:"Hash,omitempty"`
	Type          *uint32         `protobuf:"varint,15,opt,name=Type,proto3,oneof" json:"Type,omitempty"`
	Client        *bool           `protobuf:"varint,16,opt,name=Client,proto3,oneof" json:"Client,omitempty"`
	ClientRelayed *bool           `protobuf:"varint,17,opt,name=ClientRelayed,proto3,oneof" json:"ClientRelayed,omitempty"`
	ClientProxy   *bool           `protobuf:"varint,18,opt,name=ClientProxy,proto3,oneof" json:"ClientProxy,omitempty"` // for local node to handle client message and direct send back
	ClientDhtKey  []byte          `protobuf:"bytes,19,opt,name=ClientDhtKey,proto3,oneof" json:"ClientDhtKey,omitempty"`
	ClientHandled *bool           `protobuf:"varint,20,opt,name=ClientHandled,proto3,oneof" json:"ClientHandled,omitempty"`
	Universal     *bool           `protobuf:"varint,21,opt,name=Universal,proto3,oneof" json:"Universal,omitempty"`
	Broadcast     *BroadcastParam `protobuf:"bytes,22,opt,name=Broadcast,proto3,oneof" json:"Broadcast,omitempty"`
	Handled       *bool           `protobuf:"varint,23,opt,name=Handled,proto3,oneof" json:"Handled,omitempty"`
	DesDhtKeyHash *uint64         `protobuf:"varint,24,opt,name=DesDhtKeyHash,proto3,oneof" json:"DesDhtKeyHash,omitempty"`
	Pubkey        []byte          `protobuf:"bytes,25,opt,name=Pubkey,proto3,oneof" json:"Pubkey,omitempty"`
	Sign          []byte          `protobuf:"bytes,26,opt,name=Sign,proto3,oneof" json:"Sign,omitempty"`
	Local         *bool           `protobuf:"varint,27,opt,name=Local,proto3,oneof" json:"Local,omitempty"`
	DesNetwork    *int32          `protobuf:"varint,28,opt,name=DesNetwork,proto3,oneof" json:"DesNetwork,omitempty"`
	DesCountry    []byte          `protobuf:"bytes,29,opt,name=DesCountry,proto3,oneof" json:"DesCountry,omitempty"`
	TransportType *int32          `protobuf:"varint,30,opt,name=TransportType,proto3,oneof" json:"TransportType,omitempty"`
	Version       *int32          `protobuf:"varint,31,opt,name=Version,proto3,oneof" json:"Version,omitempty"`
	Timestamps    []uint64        `protobuf:"varint,32,rep,packed,name=Timestamps,proto3" json:"Timestamps,omitempty"`
	ThreadIdx     *uint32         `protobuf:"varint,33,opt,name=ThreadIdx,proto3,oneof" json:"ThreadIdx,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetSrcDhtKey() []byte {
	if x != nil {
		return x.SrcDhtKey
	}
	return nil
}

func (x *Header) GetDesDhtKey() []byte {
	if x != nil {
		return x.DesDhtKey
	}
	return nil
}

func (x *Header) GetSrcNodeId() []byte {
	if x != nil {
		return x.SrcNodeId
	}
	return nil
}

func (x *Header) GetDesNodeId() []byte {
	if x != nil {
		return x.DesNodeId
	}
	return nil
}

func (x *Header) GetHopCount() uint32 {
	if x != nil && x.HopCount != nil {
		return *x.HopCount
	}
	return 0
}

func (x *Header) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Header) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

func (x *Header) GetDebug() []byte {
	if x != nil {
		return x.Debug
	}
	return nil
}

func (x *Header) GetFromIp() []byte {
	if x != nil {
		return x.FromIp
	}
	return nil
}

func (x *Header) GetFromPort() uint32 {
	if x != nil && x.FromPort != nil {
		return *x.FromPort
	}
	return 0
}

func (x *Header) GetToIp() []byte {
	if x != nil {
		return x.ToIp
	}
	return nil
}

func (x *Header) GetToPort() uint32 {
	if x != nil && x.ToPort != nil {
		return *x.ToPort
	}
	return 0
}

func (x *Header) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Header) GetHash() uint64 {
	if x != nil && x.Hash != nil {
		return *x.Hash
	}
	return 0
}

func (x *Header) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *Header) GetClient() bool {
	if x != nil && x.Client != nil {
		return *x.Client
	}
	return false
}

func (x *Header) GetClientRelayed() bool {
	if x != nil && x.ClientRelayed != nil {
		return *x.ClientRelayed
	}
	return false
}

func (x *Header) GetClientProxy() bool {
	if x != nil && x.ClientProxy != nil {
		return *x.ClientProxy
	}
	return false
}

func (x *Header) GetClientDhtKey() []byte {
	if x != nil {
		return x.ClientDhtKey
	}
	return nil
}

func (x *Header) GetClientHandled() bool {
	if x != nil && x.ClientHandled != nil {
		return *x.ClientHandled
	}
	return false
}

func (x *Header) GetUniversal() bool {
	if x != nil && x.Universal != nil {
		return *x.Universal
	}
	return false
}

func (x *Header) GetBroadcast() *BroadcastParam {
	if x != nil {
		return x.Broadcast
	}
	return nil
}

func (x *Header) GetHandled() bool {
	if x != nil && x.Handled != nil {
		return *x.Handled
	}
	return false
}

func (x *Header) GetDesDhtKeyHash() uint64 {
	if x != nil && x.DesDhtKeyHash != nil {
		return *x.DesDhtKeyHash
	}
	return 0
}

func (x *Header) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *Header) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

func (x *Header) GetLocal() bool {
	if x != nil && x.Local != nil {
		return *x.Local
	}
	return false
}

func (x *Header) GetDesNetwork() int32 {
	if x != nil && x.DesNetwork != nil {
		return *x.DesNetwork
	}
	return 0
}

func (x *Header) GetDesCountry() []byte {
	if x != nil {
		return x.DesCountry
	}
	return nil
}

func (x *Header) GetTransportType() int32 {
	if x != nil && x.TransportType != nil {
		return *x.TransportType
	}
	return 0
}

func (x *Header) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *Header) GetTimestamps() []uint64 {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

func (x *Header) GetThreadIdx() uint32 {
	if x != nil && x.ThreadIdx != nil {
		return *x.ThreadIdx
	}
	return 0
}

var File_net_proto protoreflect.FileDescriptor

var file_net_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x74,
	0x22, 0xa9, 0x05, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d,
	0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x0d, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x53, 0x74,
	0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x48, 0x6f,
	0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x08,
	0x48, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x48, 0x04,
	0x52, 0x09, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x66, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23,
	0x0a, 0x0a, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x05, 0x52, 0x0a, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x67, 0x68, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x06, 0x52, 0x07, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x70, 0x88,
	0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x48, 0x6f, 0x70, 0x54, 0x6f, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x07, 0x52, 0x0a, 0x48, 0x6f, 0x70, 0x54, 0x6f, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x08, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0a, 0x4e, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x0a, 0x52, 0x0a, 0x4e, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x08, 0x45, 0x76, 0x69, 0x6c, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x48, 0x0b, 0x52, 0x08, 0x45, 0x76, 0x69, 0x6c, 0x52, 0x61,
	0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x49, 0x67, 0x6e, 0x42, 0x6c, 0x6f, 0x6f,
	0x6d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x0c, 0x52, 0x11, 0x49, 0x67, 0x6e, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x48, 0x6f, 0x70, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x48, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x66, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x67, 0x68, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x70, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x48, 0x6f, 0x70, 0x54,
	0x6f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x42, 0x6f, 0x64, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x4e,
	0x65, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x65, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x45, 0x76,
	0x69, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x49, 0x67, 0x6e, 0x42, 0x6c,
	0x6f, 0x6f, 0x6d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x70, 0x22, 0xf1, 0x0b, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x09, 0x53, 0x72, 0x63, 0x44, 0x68,
	0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x53, 0x72,
	0x63, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x44, 0x65,
	0x73, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52,
	0x09, 0x44, 0x65, 0x73, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a,
	0x09, 0x53, 0x72, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x02, 0x52, 0x09, 0x53, 0x72, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x09, 0x44, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x03, 0x52, 0x09, 0x44, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x48, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x08, 0x48, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x05, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x06, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x07, 0x52,
	0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x46, 0x72, 0x6f,
	0x6d, 0x49, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x08, 0x52, 0x06, 0x46, 0x72, 0x6f,
	0x6d, 0x49, 0x70, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x09, 0x52, 0x08, 0x46, 0x72, 0x6f, 0x6d,
	0x50, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x6f, 0x49, 0x70, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x0a, 0x52, 0x04, 0x54, 0x6f, 0x49, 0x70, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x54, 0x6f, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x0b, 0x52, 0x06, 0x54, 0x6f, 0x50, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x0c, 0x52, 0x02, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x0d, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x0e, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x0f, 0x52, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x29, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x48, 0x10, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x11, 0x52, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x68, 0x74,
	0x4b, 0x65, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x12, 0x52, 0x0c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x13, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x55, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x48, 0x14, 0x52, 0x09, 0x55, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x09, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x48, 0x15, 0x52, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x16, 0x52, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x29, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x18, 0x20, 0x01, 0x28, 0x04, 0x48, 0x17, 0x52, 0x0d, 0x44, 0x65, 0x73, 0x44,
	0x68, 0x74, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x18, 0x52, 0x06,
	0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x53, 0x69, 0x67,
	0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x19, 0x52, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x1b, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x1a, 0x52, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0a, 0x44, 0x65, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x1c, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x1b, 0x52, 0x0a, 0x44, 0x65, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x88,
	0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x1c, 0x52, 0x0a, 0x44, 0x65, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x1d,
	0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x1e, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x18,
	0x20, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x73, 0x12, 0x21, 0x0a, 0x09, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x78, 0x18, 0x21,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x1f, 0x52, 0x09, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64,
	0x78, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x53, 0x72, 0x63, 0x44, 0x68, 0x74, 0x4b,
	0x65, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x44, 0x65, 0x73, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x53, 0x72, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x44, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x48, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x46, 0x72,
	0x6f, 0x6d, 0x49, 0x70, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72,
	0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x6f, 0x49, 0x70, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x54,
	0x6f, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x48, 0x61, 0x73, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x44, 0x65, 0x73, 0x44, 0x68,
	0x74, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x50, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x53, 0x69, 0x67, 0x6e, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x44, 0x65, 0x73, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x44, 0x65, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x78,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6e, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_net_proto_rawDescOnce sync.Once
	file_net_proto_rawDescData = file_net_proto_rawDesc
)

func file_net_proto_rawDescGZIP() []byte {
	file_net_proto_rawDescOnce.Do(func() {
		file_net_proto_rawDescData = protoimpl.X.CompressGZIP(file_net_proto_rawDescData)
	})
	return file_net_proto_rawDescData
}

var file_net_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_net_proto_goTypes = []interface{}{
	(*BroadcastParam)(nil), // 0: net.BroadcastParam
	(*Header)(nil),         // 1: net.Header
}
var file_net_proto_depIdxs = []int32{
	0, // 0: net.Header.Broadcast:type_name -> net.BroadcastParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_net_proto_init() }
func file_net_proto_init() {
	if File_net_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_net_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastParam); i {
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
		file_net_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
	file_net_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_net_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_net_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_net_proto_goTypes,
		DependencyIndexes: file_net_proto_depIdxs,
		MessageInfos:      file_net_proto_msgTypes,
	}.Build()
	File_net_proto = out.File
	file_net_proto_rawDesc = nil
	file_net_proto_goTypes = nil
	file_net_proto_depIdxs = nil
}
