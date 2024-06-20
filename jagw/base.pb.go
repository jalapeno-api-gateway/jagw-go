// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: core/base.proto

package jagw

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

type MultiTopologyIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OFlag *bool   `protobuf:"varint,1,req,name=o_flag,json=oFlag" json:"o_flag,omitempty"`
	AFlag *bool   `protobuf:"varint,2,req,name=a_flag,json=aFlag" json:"a_flag,omitempty"`
	Mtid  *uint32 `protobuf:"varint,3,req,name=mtid" json:"mtid,omitempty"` // protobuf does not support uint16
}

func (x *MultiTopologyIdentifier) Reset() {
	*x = MultiTopologyIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiTopologyIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiTopologyIdentifier) ProtoMessage() {}

func (x *MultiTopologyIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiTopologyIdentifier.ProtoReflect.Descriptor instead.
func (*MultiTopologyIdentifier) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{0}
}

func (x *MultiTopologyIdentifier) GetOFlag() bool {
	if x != nil && x.OFlag != nil {
		return *x.OFlag
	}
	return false
}

func (x *MultiTopologyIdentifier) GetAFlag() bool {
	if x != nil && x.AFlag != nil {
		return *x.AFlag
	}
	return false
}

func (x *MultiTopologyIdentifier) GetMtid() uint32 {
	if x != nil && x.Mtid != nil {
		return *x.Mtid
	}
	return 0
}

type Srv6CapabilitiesTlv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OFlag *bool `protobuf:"varint,1,req,name=o_flag,json=oFlag" json:"o_flag,omitempty"`
}

func (x *Srv6CapabilitiesTlv) Reset() {
	*x = Srv6CapabilitiesTlv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Srv6CapabilitiesTlv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Srv6CapabilitiesTlv) ProtoMessage() {}

func (x *Srv6CapabilitiesTlv) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Srv6CapabilitiesTlv.ProtoReflect.Descriptor instead.
func (*Srv6CapabilitiesTlv) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{1}
}

func (x *Srv6CapabilitiesTlv) GetOFlag() bool {
	if x != nil && x.OFlag != nil {
		return *x.OFlag
	}
	return false
}

type NodeMsd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsdType  *uint32 `protobuf:"varint,1,req,name=msd_type,json=msdType" json:"msd_type,omitempty"`    // protobuf does not support uint8
	MsdValue *uint32 `protobuf:"varint,2,req,name=msd_value,json=msdValue" json:"msd_value,omitempty"` // protobuf does not support uint8
}

func (x *NodeMsd) Reset() {
	*x = NodeMsd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMsd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMsd) ProtoMessage() {}

func (x *NodeMsd) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMsd.ProtoReflect.Descriptor instead.
func (*NodeMsd) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{2}
}

func (x *NodeMsd) GetMsdType() uint32 {
	if x != nil && x.MsdType != nil {
		return *x.MsdType
	}
	return 0
}

func (x *NodeMsd) GetMsdValue() uint32 {
	if x != nil && x.MsdValue != nil {
		return *x.MsdValue
	}
	return 0
}

type FlexAlgoDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlexAlgo        *uint32 `protobuf:"varint,1,req,name=flex_algo,json=flexAlgo" json:"flex_algo,omitempty"`                      // protobuf does not support uint8
	MetricType      *uint32 `protobuf:"varint,2,req,name=metric_type,json=metricType" json:"metric_type,omitempty"`                // protobuf does not support uint8
	CalculationType *uint32 `protobuf:"varint,3,req,name=calculation_type,json=calculationType" json:"calculation_type,omitempty"` // protobuf does not support uint8
	Priority        *uint32 `protobuf:"varint,4,req,name=priority" json:"priority,omitempty"`                                      // protobuf does not support uint8
}

func (x *FlexAlgoDefinition) Reset() {
	*x = FlexAlgoDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlexAlgoDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlexAlgoDefinition) ProtoMessage() {}

func (x *FlexAlgoDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlexAlgoDefinition.ProtoReflect.Descriptor instead.
func (*FlexAlgoDefinition) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{3}
}

func (x *FlexAlgoDefinition) GetFlexAlgo() uint32 {
	if x != nil && x.FlexAlgo != nil {
		return *x.FlexAlgo
	}
	return 0
}

func (x *FlexAlgoDefinition) GetMetricType() uint32 {
	if x != nil && x.MetricType != nil {
		return *x.MetricType
	}
	return 0
}

func (x *FlexAlgoDefinition) GetCalculationType() uint32 {
	if x != nil && x.CalculationType != nil {
		return *x.CalculationType
	}
	return 0
}

func (x *FlexAlgoDefinition) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

type Srv6EndXSidFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BFlag *bool `protobuf:"varint,1,req,name=b_flag,json=bFlag" json:"b_flag,omitempty"`
	SFlag *bool `protobuf:"varint,2,req,name=s_flag,json=sFlag" json:"s_flag,omitempty"`
	PFlag *bool `protobuf:"varint,3,req,name=p_flag,json=pFlag" json:"p_flag,omitempty"`
}

func (x *Srv6EndXSidFlags) Reset() {
	*x = Srv6EndXSidFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Srv6EndXSidFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Srv6EndXSidFlags) ProtoMessage() {}

func (x *Srv6EndXSidFlags) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Srv6EndXSidFlags.ProtoReflect.Descriptor instead.
func (*Srv6EndXSidFlags) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{4}
}

func (x *Srv6EndXSidFlags) GetBFlag() bool {
	if x != nil && x.BFlag != nil {
		return *x.BFlag
	}
	return false
}

func (x *Srv6EndXSidFlags) GetSFlag() bool {
	if x != nil && x.SFlag != nil {
		return *x.SFlag
	}
	return false
}

func (x *Srv6EndXSidFlags) GetPFlag() bool {
	if x != nil && x.PFlag != nil {
		return *x.PFlag
	}
	return false
}

type Srv6EndXSidTlv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             *uint32           `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`                                                 // protobuf does not support uint16
	Length           *uint32           `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`                                             // protobuf does not support uint16
	EndpointBehavior *uint32           `protobuf:"varint,3,req,name=endpoint_behavior,json=endpointBehavior" json:"endpoint_behavior,omitempty"` // protobuf does not support uint16
	Flags            *Srv6EndXSidFlags `protobuf:"bytes,4,opt,name=flags" json:"flags,omitempty"`
	Algorithm        *uint32           `protobuf:"varint,5,req,name=algorithm" json:"algorithm,omitempty"` // protobuf does not support uint8
	Weight           *uint32           `protobuf:"varint,6,req,name=weight" json:"weight,omitempty"`       // protobuf does not support uint8
	Sid              *string           `protobuf:"bytes,7,opt,name=sid" json:"sid,omitempty"`
}

func (x *Srv6EndXSidTlv) Reset() {
	*x = Srv6EndXSidTlv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Srv6EndXSidTlv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Srv6EndXSidTlv) ProtoMessage() {}

func (x *Srv6EndXSidTlv) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Srv6EndXSidTlv.ProtoReflect.Descriptor instead.
func (*Srv6EndXSidTlv) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{5}
}

func (x *Srv6EndXSidTlv) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *Srv6EndXSidTlv) GetLength() uint32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *Srv6EndXSidTlv) GetEndpointBehavior() uint32 {
	if x != nil && x.EndpointBehavior != nil {
		return *x.EndpointBehavior
	}
	return 0
}

func (x *Srv6EndXSidTlv) GetFlags() *Srv6EndXSidFlags {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Srv6EndXSidTlv) GetAlgorithm() uint32 {
	if x != nil && x.Algorithm != nil {
		return *x.Algorithm
	}
	return 0
}

func (x *Srv6EndXSidTlv) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *Srv6EndXSidTlv) GetSid() string {
	if x != nil && x.Sid != nil {
		return *x.Sid
	}
	return ""
}

type Srv6LocatorFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DFlag *bool `protobuf:"varint,1,req,name=d_flag,json=dFlag" json:"d_flag,omitempty"`
}

func (x *Srv6LocatorFlags) Reset() {
	*x = Srv6LocatorFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Srv6LocatorFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Srv6LocatorFlags) ProtoMessage() {}

func (x *Srv6LocatorFlags) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Srv6LocatorFlags.ProtoReflect.Descriptor instead.
func (*Srv6LocatorFlags) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{6}
}

func (x *Srv6LocatorFlags) GetDFlag() bool {
	if x != nil && x.DFlag != nil {
		return *x.DFlag
	}
	return false
}

type Srv6LocatorTlv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flags  *Srv6LocatorFlags `protobuf:"bytes,1,opt,name=flags" json:"flags,omitempty"`
	Algo   *uint32           `protobuf:"varint,2,req,name=algo" json:"algo,omitempty"` // protobuf does not support uint8
	Metric *uint32           `protobuf:"varint,3,req,name=metric" json:"metric,omitempty"`
}

func (x *Srv6LocatorTlv) Reset() {
	*x = Srv6LocatorTlv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Srv6LocatorTlv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Srv6LocatorTlv) ProtoMessage() {}

func (x *Srv6LocatorTlv) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Srv6LocatorTlv.ProtoReflect.Descriptor instead.
func (*Srv6LocatorTlv) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{7}
}

func (x *Srv6LocatorTlv) GetFlags() *Srv6LocatorFlags {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Srv6LocatorTlv) GetAlgo() uint32 {
	if x != nil && x.Algo != nil {
		return *x.Algo
	}
	return 0
}

func (x *Srv6LocatorTlv) GetMetric() uint32 {
	if x != nil && x.Metric != nil {
		return *x.Metric
	}
	return 0
}

type Srv6EndpointBehavior struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointBehavior *uint32 `protobuf:"varint,1,req,name=endpoint_behavior,json=endpointBehavior" json:"endpoint_behavior,omitempty"` // protobuf does not support uint16
	Flag             *uint32 `protobuf:"varint,2,req,name=flag" json:"flag,omitempty"`                                                 // protobuf does not support uint8
	Algorithm        *uint32 `protobuf:"varint,3,req,name=algorithm" json:"algorithm,omitempty"`                                       // protobuf does not support uint8
}

func (x *Srv6EndpointBehavior) Reset() {
	*x = Srv6EndpointBehavior{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Srv6EndpointBehavior) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Srv6EndpointBehavior) ProtoMessage() {}

func (x *Srv6EndpointBehavior) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Srv6EndpointBehavior.ProtoReflect.Descriptor instead.
func (*Srv6EndpointBehavior) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{8}
}

func (x *Srv6EndpointBehavior) GetEndpointBehavior() uint32 {
	if x != nil && x.EndpointBehavior != nil {
		return *x.EndpointBehavior
	}
	return 0
}

func (x *Srv6EndpointBehavior) GetFlag() uint32 {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return 0
}

func (x *Srv6EndpointBehavior) GetAlgorithm() uint32 {
	if x != nil && x.Algorithm != nil {
		return *x.Algorithm
	}
	return 0
}

type Srv6SidStructure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type               *uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`                                                         // protobuf does not support uint16
	Length             *uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`                                                     // protobuf does not support uint16
	LocatorBlockLength *uint32 `protobuf:"varint,3,req,name=locator_block_length,json=locatorBlockLength" json:"locator_block_length,omitempty"` // protobuf does not support uint8
	LocatorNodeLength  *uint32 `protobuf:"varint,4,req,name=locator_node_length,json=locatorNodeLength" json:"locator_node_length,omitempty"`    // protobuf does not support uint8
	FunctionLength     *uint32 `protobuf:"varint,5,req,name=function_length,json=functionLength" json:"function_length,omitempty"`               // protobuf does not support uint8
	ArgumentLength     *uint32 `protobuf:"varint,6,req,name=argument_length,json=argumentLength" json:"argument_length,omitempty"`               // protobuf does not support uint8
}

func (x *Srv6SidStructure) Reset() {
	*x = Srv6SidStructure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_base_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Srv6SidStructure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Srv6SidStructure) ProtoMessage() {}

func (x *Srv6SidStructure) ProtoReflect() protoreflect.Message {
	mi := &file_core_base_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Srv6SidStructure.ProtoReflect.Descriptor instead.
func (*Srv6SidStructure) Descriptor() ([]byte, []int) {
	return file_core_base_proto_rawDescGZIP(), []int{9}
}

func (x *Srv6SidStructure) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *Srv6SidStructure) GetLength() uint32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *Srv6SidStructure) GetLocatorBlockLength() uint32 {
	if x != nil && x.LocatorBlockLength != nil {
		return *x.LocatorBlockLength
	}
	return 0
}

func (x *Srv6SidStructure) GetLocatorNodeLength() uint32 {
	if x != nil && x.LocatorNodeLength != nil {
		return *x.LocatorNodeLength
	}
	return 0
}

func (x *Srv6SidStructure) GetFunctionLength() uint32 {
	if x != nil && x.FunctionLength != nil {
		return *x.FunctionLength
	}
	return 0
}

func (x *Srv6SidStructure) GetArgumentLength() uint32 {
	if x != nil && x.ArgumentLength != nil {
		return *x.ArgumentLength
	}
	return 0
}

var File_core_base_proto protoreflect.FileDescriptor

var file_core_base_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6a, 0x61, 0x67, 0x77, 0x22, 0x5b, 0x0a, 0x17, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x08, 0x52, 0x05, 0x6f, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x5f, 0x66,
	0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x02, 0x28, 0x08, 0x52, 0x05, 0x61, 0x46, 0x6c, 0x61, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04,
	0x6d, 0x74, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x13, 0x53, 0x72, 0x76, 0x36, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x54, 0x6c, 0x76, 0x12, 0x15, 0x0a, 0x06, 0x6f,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x05, 0x6f, 0x46, 0x6c,
	0x61, 0x67, 0x22, 0x41, 0x0a, 0x07, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x73, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x73, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x07, 0x6d, 0x73, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x73, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x73, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x12, 0x46, 0x6c, 0x65, 0x78, 0x41, 0x6c,
	0x67, 0x6f, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x6c, 0x65, 0x78, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x08, 0x66, 0x6c, 0x65, 0x78, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x22, 0x57, 0x0a, 0x10, 0x53, 0x72, 0x76, 0x36, 0x45, 0x6e, 0x64, 0x58, 0x53, 0x69, 0x64,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x05, 0x62, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x06,
	0x73, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x02, 0x28, 0x08, 0x52, 0x05, 0x73, 0x46,
	0x6c, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x08, 0x52, 0x05, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x22, 0xdf, 0x01, 0x0a, 0x0e, 0x53,
	0x72, 0x76, 0x36, 0x45, 0x6e, 0x64, 0x58, 0x53, 0x69, 0x64, 0x54, 0x6c, 0x76, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x10, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x53, 0x72, 0x76,
	0x36, 0x45, 0x6e, 0x64, 0x58, 0x53, 0x69, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x05, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x10,
	0x53, 0x72, 0x76, 0x36, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x15, 0x0a, 0x06, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08,
	0x52, 0x05, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x6a, 0x0a, 0x0e, 0x53, 0x72, 0x76, 0x36, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x6c, 0x76, 0x12, 0x2c, 0x0a, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e,
	0x53, 0x72, 0x76, 0x36, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x73,
	0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x22, 0x75, 0x0a, 0x14, 0x53, 0x72, 0x76, 0x36, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x10, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x22, 0xf2, 0x01, 0x0a, 0x10, 0x53,
	0x72, 0x76, 0x36, 0x53, 0x69, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x14, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x2e, 0x0a,
	0x13, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x27, 0x0a,
	0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x0e, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61,
	0x6c, 0x61, 0x70, 0x65, 0x6e, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x6a, 0x61, 0x67, 0x77, 0x2d, 0x67, 0x6f, 0x3b, 0x6a, 0x61, 0x67, 0x77,
}

var (
	file_core_base_proto_rawDescOnce sync.Once
	file_core_base_proto_rawDescData = file_core_base_proto_rawDesc
)

func file_core_base_proto_rawDescGZIP() []byte {
	file_core_base_proto_rawDescOnce.Do(func() {
		file_core_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_base_proto_rawDescData)
	})
	return file_core_base_proto_rawDescData
}

var file_core_base_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_core_base_proto_goTypes = []any{
	(*MultiTopologyIdentifier)(nil), // 0: jagw.MultiTopologyIdentifier
	(*Srv6CapabilitiesTlv)(nil),     // 1: jagw.Srv6CapabilitiesTlv
	(*NodeMsd)(nil),                 // 2: jagw.NodeMsd
	(*FlexAlgoDefinition)(nil),      // 3: jagw.FlexAlgoDefinition
	(*Srv6EndXSidFlags)(nil),        // 4: jagw.Srv6EndXSidFlags
	(*Srv6EndXSidTlv)(nil),          // 5: jagw.Srv6EndXSidTlv
	(*Srv6LocatorFlags)(nil),        // 6: jagw.Srv6LocatorFlags
	(*Srv6LocatorTlv)(nil),          // 7: jagw.Srv6LocatorTlv
	(*Srv6EndpointBehavior)(nil),    // 8: jagw.Srv6EndpointBehavior
	(*Srv6SidStructure)(nil),        // 9: jagw.Srv6SidStructure
}
var file_core_base_proto_depIdxs = []int32{
	4, // 0: jagw.Srv6EndXSidTlv.flags:type_name -> jagw.Srv6EndXSidFlags
	6, // 1: jagw.Srv6LocatorTlv.flags:type_name -> jagw.Srv6LocatorFlags
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_base_proto_init() }
func file_core_base_proto_init() {
	if File_core_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_base_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MultiTopologyIdentifier); i {
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
		file_core_base_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Srv6CapabilitiesTlv); i {
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
		file_core_base_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NodeMsd); i {
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
		file_core_base_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FlexAlgoDefinition); i {
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
		file_core_base_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Srv6EndXSidFlags); i {
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
		file_core_base_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Srv6EndXSidTlv); i {
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
		file_core_base_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Srv6LocatorFlags); i {
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
		file_core_base_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Srv6LocatorTlv); i {
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
		file_core_base_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Srv6EndpointBehavior); i {
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
		file_core_base_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Srv6SidStructure); i {
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
			RawDescriptor: file_core_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_base_proto_goTypes,
		DependencyIndexes: file_core_base_proto_depIdxs,
		MessageInfos:      file_core_base_proto_msgTypes,
	}.Build()
	File_core_base_proto = out.File
	file_core_base_proto_rawDesc = nil
	file_core_base_proto_goTypes = nil
	file_core_base_proto_depIdxs = nil
}
