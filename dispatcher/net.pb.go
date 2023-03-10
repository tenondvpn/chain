// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: net.proto

package dispatcher

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

	NeighborCount     *uint32  `protobuf:"varint,1,opt,name=NeighborCount,def=7" json:"NeighborCount,omitempty"`
	StopTimes         *uint32  `protobuf:"varint,2,opt,name=StopTimes,def=2" json:"StopTimes,omitempty"`
	HopLimit          *uint32  `protobuf:"varint,3,opt,name=HopLimit,def=10" json:"HopLimit,omitempty"`
	LayerLeft         *uint64  `protobuf:"varint,4,opt,name=LayerLeft" json:"LayerLeft,omitempty"`
	LayerRight        *uint64  `protobuf:"varint,5,opt,name=LayerRight" json:"LayerRight,omitempty"`
	Overlap           *float32 `protobuf:"fixed32,6,opt,name=Overlap,def=0.1" json:"Overlap,omitempty"`
	HopToLayer        *uint32  `protobuf:"varint,7,opt,name=HopToLayer,def=2" json:"HopToLayer,omitempty"`
	Bloomfilter       []uint64 `protobuf:"varint,8,rep,name=Bloomfilter" json:"Bloomfilter,omitempty"`
	EvilRate          *float32 `protobuf:"fixed32,9,opt,name=EvilRate,def=0" json:"EvilRate,omitempty"`
	IgnBloomfilterHop *uint32  `protobuf:"varint,10,opt,name=IgnBloomfilterHop,def=1" json:"IgnBloomfilterHop,omitempty"`
}

// Default values for BroadcastParam fields.
const (
	Default_BroadcastParam_NeighborCount     = uint32(7)
	Default_BroadcastParam_StopTimes         = uint32(2)
	Default_BroadcastParam_HopLimit          = uint32(10)
	Default_BroadcastParam_Overlap           = float32(0.10000000149011612)
	Default_BroadcastParam_HopToLayer        = uint32(2)
	Default_BroadcastParam_EvilRate          = float32(0)
	Default_BroadcastParam_IgnBloomfilterHop = uint32(1)
)

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

func (x *BroadcastParam) GetNeighborCount() uint32 {
	if x != nil && x.NeighborCount != nil {
		return *x.NeighborCount
	}
	return Default_BroadcastParam_NeighborCount
}

func (x *BroadcastParam) GetStopTimes() uint32 {
	if x != nil && x.StopTimes != nil {
		return *x.StopTimes
	}
	return Default_BroadcastParam_StopTimes
}

func (x *BroadcastParam) GetHopLimit() uint32 {
	if x != nil && x.HopLimit != nil {
		return *x.HopLimit
	}
	return Default_BroadcastParam_HopLimit
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
	return Default_BroadcastParam_Overlap
}

func (x *BroadcastParam) GetHopToLayer() uint32 {
	if x != nil && x.HopToLayer != nil {
		return *x.HopToLayer
	}
	return Default_BroadcastParam_HopToLayer
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
	return Default_BroadcastParam_EvilRate
}

func (x *BroadcastParam) GetIgnBloomfilterHop() uint32 {
	if x != nil && x.IgnBloomfilterHop != nil {
		return *x.IgnBloomfilterHop
	}
	return Default_BroadcastParam_IgnBloomfilterHop
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcDhtKey     *string         `protobuf:"bytes,1,opt,name=SrcDhtKey" json:"SrcDhtKey,omitempty"`
	DesDhtKey     *string         `protobuf:"bytes,2,opt,name=DesDhtKey" json:"DesDhtKey,omitempty"`
	SrcNodeId     *string         `protobuf:"bytes,3,opt,name=SrcNodeId" json:"SrcNodeId,omitempty"`
	DesNodeId     *string         `protobuf:"bytes,4,opt,name=DesNodeId" json:"DesNodeId,omitempty"`
	HopCount      *uint32         `protobuf:"varint,5,opt,name=HopCount" json:"HopCount,omitempty"`
	Data          *string         `protobuf:"bytes,6,opt,name=Data" json:"Data,omitempty"`
	Priority      *uint32         `protobuf:"varint,7,opt,name=Priority" json:"Priority,omitempty"`
	Debug         *string         `protobuf:"bytes,8,opt,name=Debug" json:"Debug,omitempty"`
	FromIp        *string         `protobuf:"bytes,9,opt,name=FromIp" json:"FromIp,omitempty"`       // will clear before transport send, not use
	FromPort      *uint32         `protobuf:"varint,10,opt,name=FromPort" json:"FromPort,omitempty"` // will clear before transport send, not use
	ToIp          *string         `protobuf:"bytes,11,opt,name=ToIp" json:"ToIp,omitempty"`          // will clear before transport send, not use
	ToPort        *uint32         `protobuf:"varint,12,opt,name=ToPort" json:"ToPort,omitempty"`     // will clear before transport send, not use
	Id            *uint32         `protobuf:"varint,13,opt,name=Id" json:"Id,omitempty"`
	Hash          *uint64         `protobuf:"varint,14,opt,name=Hash" json:"Hash,omitempty"`
	Type          *uint32         `protobuf:"varint,15,opt,name=Type" json:"Type,omitempty"`
	Client        *bool           `protobuf:"varint,16,opt,name=Client" json:"Client,omitempty"`
	ClientRelayed *bool           `protobuf:"varint,17,opt,name=ClientRelayed" json:"ClientRelayed,omitempty"`
	ClientProxy   *bool           `protobuf:"varint,18,opt,name=ClientProxy" json:"ClientProxy,omitempty"` // for local node to handle client message and direct send back
	ClientDhtKey  *string         `protobuf:"bytes,19,opt,name=ClientDhtKey" json:"ClientDhtKey,omitempty"`
	ClientHandled *bool           `protobuf:"varint,20,opt,name=ClientHandled" json:"ClientHandled,omitempty"`
	Universal     *bool           `protobuf:"varint,21,opt,name=Universal" json:"Universal,omitempty"`
	Broadcast     *BroadcastParam `protobuf:"bytes,22,opt,name=Broadcast" json:"Broadcast,omitempty"`
	Handled       *bool           `protobuf:"varint,23,opt,name=Handled" json:"Handled,omitempty"`
	DesDhtKeyHash *uint64         `protobuf:"varint,24,opt,name=DesDhtKeyHash" json:"DesDhtKeyHash,omitempty"`
	Pubkey        *string         `protobuf:"bytes,25,opt,name=Pubkey" json:"Pubkey,omitempty"`
	Sign          *string         `protobuf:"bytes,26,opt,name=Sign" json:"Sign,omitempty"`
	Local         *bool           `protobuf:"varint,27,opt,name=Local" json:"Local,omitempty"`
	DesNetwork    *int32          `protobuf:"varint,28,opt,name=DesNetwork" json:"DesNetwork,omitempty"`
	DesCountry    *string         `protobuf:"bytes,29,opt,name=DesCountry" json:"DesCountry,omitempty"`
	TransportType *int32          `protobuf:"varint,30,opt,name=TransportType" json:"TransportType,omitempty"`
	Version       *int32          `protobuf:"varint,31,opt,name=Version" json:"Version,omitempty"`
	Timestamps    []uint64        `protobuf:"varint,32,rep,name=Timestamps" json:"Timestamps,omitempty"` // for test delay
	ThreadIdx     *uint32         `protobuf:"varint,33,opt,name=ThreadIdx" json:"ThreadIdx,omitempty"`
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

func (x *Header) GetSrcDhtKey() string {
	if x != nil && x.SrcDhtKey != nil {
		return *x.SrcDhtKey
	}
	return ""
}

func (x *Header) GetDesDhtKey() string {
	if x != nil && x.DesDhtKey != nil {
		return *x.DesDhtKey
	}
	return ""
}

func (x *Header) GetSrcNodeId() string {
	if x != nil && x.SrcNodeId != nil {
		return *x.SrcNodeId
	}
	return ""
}

func (x *Header) GetDesNodeId() string {
	if x != nil && x.DesNodeId != nil {
		return *x.DesNodeId
	}
	return ""
}

func (x *Header) GetHopCount() uint32 {
	if x != nil && x.HopCount != nil {
		return *x.HopCount
	}
	return 0
}

func (x *Header) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

func (x *Header) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

func (x *Header) GetDebug() string {
	if x != nil && x.Debug != nil {
		return *x.Debug
	}
	return ""
}

func (x *Header) GetFromIp() string {
	if x != nil && x.FromIp != nil {
		return *x.FromIp
	}
	return ""
}

func (x *Header) GetFromPort() uint32 {
	if x != nil && x.FromPort != nil {
		return *x.FromPort
	}
	return 0
}

func (x *Header) GetToIp() string {
	if x != nil && x.ToIp != nil {
		return *x.ToIp
	}
	return ""
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

func (x *Header) GetClientDhtKey() string {
	if x != nil && x.ClientDhtKey != nil {
		return *x.ClientDhtKey
	}
	return ""
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

func (x *Header) GetPubkey() string {
	if x != nil && x.Pubkey != nil {
		return *x.Pubkey
	}
	return ""
}

func (x *Header) GetSign() string {
	if x != nil && x.Sign != nil {
		return *x.Sign
	}
	return ""
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

func (x *Header) GetDesCountry() string {
	if x != nil && x.DesCountry != nil {
		return *x.DesCountry
	}
	return ""
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
	0x0a, 0x09, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0xec, 0x02, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x27, 0x0a, 0x0d, 0x4e, 0x65,
	0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x3a, 0x01, 0x37, 0x52, 0x0d, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x32, 0x52, 0x09, 0x53, 0x74, 0x6f, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x08, 0x48, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x02, 0x31, 0x30, 0x52, 0x08, 0x48, 0x6f, 0x70, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x66,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65,
	0x66, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x3a, 0x03, 0x30, 0x2e, 0x31, 0x52, 0x07, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x70, 0x12, 0x21, 0x0a, 0x0a, 0x48, 0x6f, 0x70, 0x54, 0x6f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x32, 0x52, 0x0a, 0x48, 0x6f, 0x70, 0x54, 0x6f, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x42, 0x6c, 0x6f, 0x6f, 0x6d,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x08, 0x45, 0x76, 0x69, 0x6c, 0x52, 0x61,
	0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x45, 0x76, 0x69,
	0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x11, 0x49, 0x67, 0x6e, 0x42, 0x6c, 0x6f, 0x6f,
	0x6d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x3a, 0x01, 0x31, 0x52, 0x11, 0x49, 0x67, 0x6e, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x48, 0x6f, 0x70, 0x22, 0xbc, 0x07, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x72, 0x63, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x72, 0x63, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x44, 0x65, 0x73, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x44, 0x65, 0x73, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x72, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x53, 0x72, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x44,
	0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x44, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x70,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x48, 0x6f, 0x70,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x46,
	0x72, 0x6f, 0x6d, 0x49, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x72, 0x6f,
	0x6d, 0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x6f, 0x49, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x6f, 0x49, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x6f, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x54, 0x6f, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x68, 0x74,
	0x4b, 0x65, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x09, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x44, 0x68, 0x74, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x44, 0x65, 0x73, 0x44, 0x68, 0x74, 0x4b, 0x65,
	0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x69, 0x67,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x44, 0x65, 0x73,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x49, 0x64, 0x78, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x49, 0x64, 0x78, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
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
	(*BroadcastParam)(nil), // 0: dispatcher.BroadcastParam
	(*Header)(nil),         // 1: dispatcher.Header
}
var file_net_proto_depIdxs = []int32{
	0, // 0: dispatcher.Header.Broadcast:type_name -> dispatcher.BroadcastParam
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
