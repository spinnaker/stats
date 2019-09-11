// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stats/log.proto

package stats

import (
	fmt "fmt"
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

// Sourced from https://github.com/spinnaker/orca/blob/master/orca-core/src/main/java/com/netflix/spinnaker/orca/ExecutionStatus.java
type Status int32

const (
	Status_UNKNOWN         Status = 0
	Status_NOT_STARTED     Status = 1
	Status_RUNNING         Status = 2
	Status_PAUSED          Status = 3
	Status_SUSPENDED       Status = 4
	Status_SUCCEEDED       Status = 5
	Status_FAILED_CONTINUE Status = 6
	Status_TERMINAL        Status = 7
	Status_CANCELED        Status = 8
	Status_REDIRECT        Status = 9
	Status_STOPPED         Status = 10
	Status_SKIPPED         Status = 11
	Status_BUFFERED        Status = 12
)

var Status_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "NOT_STARTED",
	2:  "RUNNING",
	3:  "PAUSED",
	4:  "SUSPENDED",
	5:  "SUCCEEDED",
	6:  "FAILED_CONTINUE",
	7:  "TERMINAL",
	8:  "CANCELED",
	9:  "REDIRECT",
	10: "STOPPED",
	11: "SKIPPED",
	12: "BUFFERED",
}

var Status_value = map[string]int32{
	"UNKNOWN":         0,
	"NOT_STARTED":     1,
	"RUNNING":         2,
	"PAUSED":          3,
	"SUSPENDED":       4,
	"SUCCEEDED":       5,
	"FAILED_CONTINUE": 6,
	"TERMINAL":        7,
	"CANCELED":        8,
	"REDIRECT":        9,
	"STOPPED":         10,
	"SKIPPED":         11,
	"BUFFERED":        12,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{0}
}

type Execution_Type int32

const (
	Execution_UNKNOWN                      Execution_Type = 0
	Execution_PIPELINE                     Execution_Type = 1
	Execution_ORCHESTRATION                Execution_Type = 2
	Execution_MANAGED_PIPELINE_TEMPLATE_V1 Execution_Type = 3
	Execution_MANAGED_PIPELINE_TEMPLATE_V2 Execution_Type = 4
)

var Execution_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "PIPELINE",
	2: "ORCHESTRATION",
	3: "MANAGED_PIPELINE_TEMPLATE_V1",
	4: "MANAGED_PIPELINE_TEMPLATE_V2",
}

var Execution_Type_value = map[string]int32{
	"UNKNOWN":                      0,
	"PIPELINE":                     1,
	"ORCHESTRATION":                2,
	"MANAGED_PIPELINE_TEMPLATE_V1": 3,
	"MANAGED_PIPELINE_TEMPLATE_V2": 4,
}

func (x Execution_Type) String() string {
	return proto.EnumName(Execution_Type_name, int32(x))
}

func (Execution_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{3, 0}
}

// Sourced from https://github.com/spinnaker/echo/tree/master/echo-model/src/main/java/com/netflix/spinnaker/echo/model/trigger
type Execution_Trigger_Type int32

const (
	Execution_Trigger_UNKNOWN     Execution_Trigger_Type = 0
	Execution_Trigger_ARTIFACTORY Execution_Trigger_Type = 1
	Execution_Trigger_BUILD       Execution_Trigger_Type = 2
	Execution_Trigger_DOCKER      Execution_Trigger_Type = 3
	Execution_Trigger_GIT         Execution_Trigger_Type = 4
	Execution_Trigger_MANUAL      Execution_Trigger_Type = 5
	Execution_Trigger_PUBSUB      Execution_Trigger_Type = 6
	Execution_Trigger_WEBHOOK     Execution_Trigger_Type = 7
)

var Execution_Trigger_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARTIFACTORY",
	2: "BUILD",
	3: "DOCKER",
	4: "GIT",
	5: "MANUAL",
	6: "PUBSUB",
	7: "WEBHOOK",
}

var Execution_Trigger_Type_value = map[string]int32{
	"UNKNOWN":     0,
	"ARTIFACTORY": 1,
	"BUILD":       2,
	"DOCKER":      3,
	"GIT":         4,
	"MANUAL":      5,
	"PUBSUB":      6,
	"WEBHOOK":     7,
}

func (x Execution_Trigger_Type) String() string {
	return proto.EnumName(Execution_Trigger_Type_name, int32(x))
}

func (Execution_Trigger_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{3, 0, 0}
}

type Event struct {
	SpinnakerInstance    *SpinnakerInstance `protobuf:"bytes,1,opt,name=spinnaker_instance,json=spinnakerInstance,proto3" json:"spinnaker_instance,omitempty"`
	Application          *Application       `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
	Execution            *Execution         `protobuf:"bytes,3,opt,name=execution,proto3" json:"execution,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetSpinnakerInstance() *SpinnakerInstance {
	if m != nil {
		return m.SpinnakerInstance
	}
	return nil
}

func (m *Event) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *Event) GetExecution() *Execution {
	if m != nil {
		return m.Execution
	}
	return nil
}

type SpinnakerInstance struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpinnakerInstance) Reset()         { *m = SpinnakerInstance{} }
func (m *SpinnakerInstance) String() string { return proto.CompactTextString(m) }
func (*SpinnakerInstance) ProtoMessage()    {}
func (*SpinnakerInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{1}
}

func (m *SpinnakerInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpinnakerInstance.Unmarshal(m, b)
}
func (m *SpinnakerInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpinnakerInstance.Marshal(b, m, deterministic)
}
func (m *SpinnakerInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpinnakerInstance.Merge(m, src)
}
func (m *SpinnakerInstance) XXX_Size() int {
	return xxx_messageInfo_SpinnakerInstance.Size(m)
}
func (m *SpinnakerInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_SpinnakerInstance.DiscardUnknown(m)
}

var xxx_messageInfo_SpinnakerInstance proto.InternalMessageInfo

func (m *SpinnakerInstance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SpinnakerInstance) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Application struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{2}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Execution struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Execution_Type     `protobuf:"varint,2,opt,name=type,proto3,enum=io.spinnaker.protos.stats.Execution_Type" json:"type,omitempty"`
	Trigger              *Execution_Trigger `protobuf:"bytes,3,opt,name=trigger,proto3" json:"trigger,omitempty"`
	Stages               []*Stage           `protobuf:"bytes,4,rep,name=stages,proto3" json:"stages,omitempty"`
	Status               Status             `protobuf:"varint,5,opt,name=status,proto3,enum=io.spinnaker.protos.stats.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Execution) Reset()         { *m = Execution{} }
func (m *Execution) String() string { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()    {}
func (*Execution) Descriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{3}
}

func (m *Execution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Execution.Unmarshal(m, b)
}
func (m *Execution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Execution.Marshal(b, m, deterministic)
}
func (m *Execution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execution.Merge(m, src)
}
func (m *Execution) XXX_Size() int {
	return xxx_messageInfo_Execution.Size(m)
}
func (m *Execution) XXX_DiscardUnknown() {
	xxx_messageInfo_Execution.DiscardUnknown(m)
}

var xxx_messageInfo_Execution proto.InternalMessageInfo

func (m *Execution) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Execution) GetType() Execution_Type {
	if m != nil {
		return m.Type
	}
	return Execution_UNKNOWN
}

func (m *Execution) GetTrigger() *Execution_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *Execution) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

func (m *Execution) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNKNOWN
}

type Execution_Trigger struct {
	Type                 Execution_Trigger_Type `protobuf:"varint,1,opt,name=type,proto3,enum=io.spinnaker.protos.stats.Execution_Trigger_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Execution_Trigger) Reset()         { *m = Execution_Trigger{} }
func (m *Execution_Trigger) String() string { return proto.CompactTextString(m) }
func (*Execution_Trigger) ProtoMessage()    {}
func (*Execution_Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{3, 0}
}

func (m *Execution_Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Execution_Trigger.Unmarshal(m, b)
}
func (m *Execution_Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Execution_Trigger.Marshal(b, m, deterministic)
}
func (m *Execution_Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execution_Trigger.Merge(m, src)
}
func (m *Execution_Trigger) XXX_Size() int {
	return xxx_messageInfo_Execution_Trigger.Size(m)
}
func (m *Execution_Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Execution_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Execution_Trigger proto.InternalMessageInfo

func (m *Execution_Trigger) GetType() Execution_Trigger_Type {
	if m != nil {
		return m.Type
	}
	return Execution_Trigger_UNKNOWN
}

type Stage struct {
	Type                 string         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Status               Status         `protobuf:"varint,2,opt,name=status,proto3,enum=io.spinnaker.protos.stats.Status" json:"status,omitempty"`
	CloudProvider        *CloudProvider `protobuf:"bytes,3,opt,name=cloud_provider,json=cloudProvider,proto3" json:"cloud_provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Stage) Reset()         { *m = Stage{} }
func (m *Stage) String() string { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()    {}
func (*Stage) Descriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{4}
}

func (m *Stage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stage.Unmarshal(m, b)
}
func (m *Stage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stage.Marshal(b, m, deterministic)
}
func (m *Stage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stage.Merge(m, src)
}
func (m *Stage) XXX_Size() int {
	return xxx_messageInfo_Stage.Size(m)
}
func (m *Stage) XXX_DiscardUnknown() {
	xxx_messageInfo_Stage.DiscardUnknown(m)
}

var xxx_messageInfo_Stage proto.InternalMessageInfo

func (m *Stage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Stage) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNKNOWN
}

func (m *Stage) GetCloudProvider() *CloudProvider {
	if m != nil {
		return m.CloudProvider
	}
	return nil
}

type CloudProvider struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Variant              string   `protobuf:"bytes,2,opt,name=variant,proto3" json:"variant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudProvider) Reset()         { *m = CloudProvider{} }
func (m *CloudProvider) String() string { return proto.CompactTextString(m) }
func (*CloudProvider) ProtoMessage()    {}
func (*CloudProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_be815d6cfeb20831, []int{5}
}

func (m *CloudProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudProvider.Unmarshal(m, b)
}
func (m *CloudProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudProvider.Marshal(b, m, deterministic)
}
func (m *CloudProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudProvider.Merge(m, src)
}
func (m *CloudProvider) XXX_Size() int {
	return xxx_messageInfo_CloudProvider.Size(m)
}
func (m *CloudProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudProvider.DiscardUnknown(m)
}

var xxx_messageInfo_CloudProvider proto.InternalMessageInfo

func (m *CloudProvider) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CloudProvider) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func init() {
	proto.RegisterEnum("io.spinnaker.protos.stats.Status", Status_name, Status_value)
	proto.RegisterEnum("io.spinnaker.protos.stats.Execution_Type", Execution_Type_name, Execution_Type_value)
	proto.RegisterEnum("io.spinnaker.protos.stats.Execution_Trigger_Type", Execution_Trigger_Type_name, Execution_Trigger_Type_value)
	proto.RegisterType((*Event)(nil), "io.spinnaker.protos.stats.Event")
	proto.RegisterType((*SpinnakerInstance)(nil), "io.spinnaker.protos.stats.SpinnakerInstance")
	proto.RegisterType((*Application)(nil), "io.spinnaker.protos.stats.Application")
	proto.RegisterType((*Execution)(nil), "io.spinnaker.protos.stats.Execution")
	proto.RegisterType((*Execution_Trigger)(nil), "io.spinnaker.protos.stats.Execution.Trigger")
	proto.RegisterType((*Stage)(nil), "io.spinnaker.protos.stats.Stage")
	proto.RegisterType((*CloudProvider)(nil), "io.spinnaker.protos.stats.CloudProvider")
}

func init() { proto.RegisterFile("stats/log.proto", fileDescriptor_be815d6cfeb20831) }

var fileDescriptor_be815d6cfeb20831 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xf9, 0xac, 0x27, 0x4d, 0xbb, 0x5d, 0x2e, 0x01, 0x81, 0x94, 0x5a, 0x08, 0x15, 0x54,
	0x39, 0x6a, 0xb8, 0xd0, 0x43, 0x0f, 0x8e, 0xbd, 0x49, 0xad, 0x24, 0x6b, 0x6b, 0xbd, 0xa6, 0x02,
	0x0e, 0x91, 0x9b, 0x98, 0x60, 0x91, 0xda, 0x51, 0xec, 0x54, 0xed, 0x2f, 0xe2, 0x47, 0x70, 0xe5,
	0xce, 0x7f, 0xe1, 0x17, 0x20, 0x6f, 0x9c, 0x34, 0xb4, 0x25, 0xb4, 0x37, 0xbf, 0xf1, 0x7b, 0x6f,
	0x67, 0xe6, 0xed, 0xc2, 0x6e, 0x9c, 0x78, 0x49, 0xdc, 0x98, 0x44, 0x63, 0x75, 0x3a, 0x8b, 0x92,
	0x08, 0x3f, 0x0b, 0x22, 0x35, 0x9e, 0x06, 0x61, 0xe8, 0x7d, 0xf3, 0x67, 0x8b, 0x5a, 0xac, 0x0a,
	0x92, 0xf2, 0x5b, 0x82, 0x22, 0xb9, 0xf4, 0xc3, 0x04, 0x7f, 0x06, 0xbc, 0xe2, 0x0c, 0x82, 0x30,
	0x4e, 0xbc, 0x70, 0xe8, 0xd7, 0xa4, 0xba, 0x74, 0x50, 0x69, 0x1e, 0xaa, 0xff, 0x74, 0x50, 0x9d,
	0x65, 0xd9, 0xcc, 0x34, 0x6c, 0x2f, 0xbe, 0x5d, 0xc2, 0xa7, 0x50, 0xf1, 0xa6, 0xd3, 0x49, 0x30,
	0xf4, 0x92, 0x20, 0x0a, 0x6b, 0x39, 0xe1, 0xfa, 0x7a, 0x83, 0xab, 0x76, 0xc3, 0x66, 0xeb, 0x52,
	0xdc, 0x02, 0xd9, 0xbf, 0xf2, 0x87, 0x73, 0xe1, 0x93, 0x17, 0x3e, 0xaf, 0x36, 0xf8, 0x90, 0x25,
	0x97, 0xdd, 0xc8, 0x94, 0x13, 0xd8, 0xbb, 0xd3, 0x35, 0xde, 0x81, 0x5c, 0x30, 0x12, 0xf3, 0xca,
	0x2c, 0x17, 0x8c, 0x70, 0x0d, 0xca, 0x97, 0xfe, 0x2c, 0x5e, 0xb6, 0x2b, 0xb3, 0x25, 0x54, 0x5e,
	0x42, 0x65, 0xad, 0xbd, 0xdb, 0x42, 0xe5, 0x57, 0x01, 0xe4, 0xd5, 0xb1, 0x77, 0x6c, 0x4f, 0xa0,
	0x90, 0x5c, 0x4f, 0x7d, 0xe1, 0xb9, 0xd3, 0x7c, 0xf3, 0x90, 0xd6, 0x55, 0x7e, 0x3d, 0xf5, 0x99,
	0x90, 0xe1, 0x36, 0x94, 0x93, 0x59, 0x30, 0x1e, 0xfb, 0xb3, 0x6c, 0xf8, 0xc3, 0x87, 0x39, 0x2c,
	0x34, 0x6c, 0x29, 0xc6, 0xef, 0xa1, 0x14, 0x27, 0xde, 0xd8, 0x8f, 0x6b, 0x85, 0x7a, 0xfe, 0xa0,
	0xd2, 0xac, 0x6f, 0x4a, 0x38, 0x25, 0xb2, 0x8c, 0x8f, 0x8f, 0x85, 0x32, 0x99, 0xc7, 0xb5, 0xa2,
	0x18, 0x61, 0x7f, 0xb3, 0x32, 0x99, 0xc7, 0x2c, 0x13, 0x3c, 0xff, 0x21, 0x41, 0x39, 0xeb, 0x04,
	0x93, 0x6c, 0x0f, 0x92, 0x30, 0x39, 0x7a, 0xcc, 0x14, 0x6b, 0xfb, 0x50, 0x02, 0x28, 0xa4, 0x08,
	0x57, 0xa0, 0xec, 0xd2, 0x2e, 0xb5, 0xce, 0x28, 0x7a, 0x82, 0x77, 0xa1, 0xa2, 0x31, 0x6e, 0xb6,
	0x35, 0x9d, 0x5b, 0xec, 0x23, 0x92, 0xb0, 0x0c, 0xc5, 0x96, 0x6b, 0xf6, 0x0c, 0x94, 0xc3, 0x00,
	0x25, 0xc3, 0xd2, 0xbb, 0x84, 0xa1, 0x3c, 0x2e, 0x43, 0xbe, 0x63, 0x72, 0x54, 0x48, 0x8b, 0x7d,
	0x8d, 0xba, 0x5a, 0x0f, 0x15, 0xd3, 0x6f, 0xdb, 0x6d, 0x39, 0x6e, 0x0b, 0x95, 0x52, 0xd7, 0x33,
	0xd2, 0x3a, 0xb5, 0xac, 0x2e, 0x2a, 0x2b, 0x57, 0xf7, 0x1d, 0xb5, 0x0d, 0x5b, 0xb6, 0x69, 0x93,
	0x9e, 0x49, 0x09, 0x92, 0xf0, 0x1e, 0x54, 0x2d, 0xa6, 0x9f, 0x12, 0x87, 0x33, 0x8d, 0x9b, 0x16,
	0x45, 0x39, 0x5c, 0x87, 0x17, 0x7d, 0x8d, 0x6a, 0x1d, 0x62, 0x0c, 0x96, 0xc4, 0x01, 0x27, 0x7d,
	0xbb, 0xa7, 0x71, 0x32, 0xf8, 0x70, 0x84, 0xf2, 0xff, 0x61, 0x34, 0x51, 0x41, 0xf9, 0x2e, 0x41,
	0x51, 0x84, 0x80, 0xf1, 0xda, 0xd6, 0xe4, 0xec, 0x4a, 0xdc, 0x04, 0x92, 0x7b, 0x64, 0x20, 0xd8,
	0x82, 0x9d, 0xe1, 0x24, 0x9a, 0x8f, 0x06, 0xd3, 0x59, 0x74, 0x19, 0x8c, 0x56, 0x97, 0xea, 0x60,
	0x83, 0x85, 0x9e, 0x0a, 0xec, 0x8c, 0xcf, 0xaa, 0xc3, 0x75, 0xa8, 0x1c, 0x43, 0xf5, 0xaf, 0xff,
	0xf7, 0xbe, 0x2a, 0x6f, 0x16, 0x78, 0x61, 0xb2, 0x7a, 0x55, 0x0b, 0xf8, 0xf6, 0xa7, 0x04, 0xa5,
	0x45, 0x7b, 0x77, 0xc2, 0xa4, 0x16, 0x1f, 0x38, 0x5c, 0x63, 0x9c, 0x18, 0x48, 0x4a, 0xff, 0x32,
	0x97, 0x52, 0x93, 0x76, 0x16, 0x71, 0xda, 0x9a, 0xeb, 0x10, 0x03, 0xe5, 0x71, 0x15, 0x64, 0xc7,
	0x75, 0x6c, 0x42, 0x0d, 0x62, 0xa0, 0xc2, 0x02, 0xea, 0x3a, 0x21, 0x29, 0x2c, 0xe2, 0xa7, 0xb0,
	0xdb, 0xd6, 0xcc, 0x1e, 0x31, 0x06, 0xba, 0x45, 0xb9, 0x49, 0x5d, 0x82, 0x4a, 0x69, 0x7c, 0x9c,
	0xb0, 0xbe, 0x49, 0xb5, 0x1e, 0x2a, 0xa7, 0x48, 0xd7, 0xa8, 0x4e, 0x7a, 0xc4, 0x40, 0x5b, 0x29,
	0x62, 0xc4, 0x30, 0x19, 0xd1, 0x39, 0x92, 0xd3, 0x53, 0x1d, 0x6e, 0xd9, 0x36, 0x31, 0x10, 0x08,
	0xd0, 0x35, 0x05, 0xa8, 0xa4, 0xbc, 0x96, 0xdb, 0x6e, 0x13, 0x46, 0x0c, 0xb4, 0xdd, 0xea, 0xc0,
	0xfe, 0x30, 0xba, 0x50, 0x43, 0x3f, 0xf9, 0x32, 0x09, 0xae, 0x6e, 0x2f, 0x71, 0xb1, 0x43, 0x5b,
	0xfa, 0x54, 0x1f, 0x07, 0xc9, 0xd7, 0xf9, 0xb9, 0x3a, 0x8c, 0x2e, 0x1a, 0x2b, 0x4e, 0x43, 0x70,
	0x1a, 0x82, 0x73, 0x5e, 0x12, 0xe0, 0xdd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x65, 0x9d,
	0x71, 0xce, 0x05, 0x00, 0x00,
}
