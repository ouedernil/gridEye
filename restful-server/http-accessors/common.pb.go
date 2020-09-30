// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	common.proto
	mon_msg.proto

It has these top-level messages:
	Command
	UMonitor
	IMonitor
	FMonitor
	EMonitor
	ThdsMonitor
	UpdateConfig
	MonMeasure
	RegMeasure
	GenericData
	MonitoringEvt
	Alarm
	MonMsg
	MonRemoteConfig
	MonRemoteConfigOneParam
*/
package http_accessors

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Command struct {
	Type             *uint32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Args             []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Command) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type UMonitor struct {
	Value            *float32 `protobuf:"fixed32,1,req,name=value" json:"value,omitempty"`
	Min              *float32 `protobuf:"fixed32,2,req,name=min" json:"min,omitempty"`
	Max              *float32 `protobuf:"fixed32,3,req,name=max" json:"max,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *UMonitor) Reset()                    { *m = UMonitor{} }
func (m *UMonitor) String() string            { return proto.CompactTextString(m) }
func (*UMonitor) ProtoMessage()               {}
func (*UMonitor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UMonitor) GetValue() float32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *UMonitor) GetMin() float32 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *UMonitor) GetMax() float32 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

type IMonitor struct {
	Value            *float32 `protobuf:"fixed32,1,req,name=value" json:"value,omitempty"`
	PMax             *float32 `protobuf:"fixed32,2,req,name=p_max,json=pMax" json:"p_max,omitempty"`
	NMax             *float32 `protobuf:"fixed32,3,req,name=n_max,json=nMax" json:"n_max,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IMonitor) Reset()                    { *m = IMonitor{} }
func (m *IMonitor) String() string            { return proto.CompactTextString(m) }
func (*IMonitor) ProtoMessage()               {}
func (*IMonitor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IMonitor) GetValue() float32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *IMonitor) GetPMax() float32 {
	if m != nil && m.PMax != nil {
		return *m.PMax
	}
	return 0
}

func (m *IMonitor) GetNMax() float32 {
	if m != nil && m.NMax != nil {
		return *m.NMax
	}
	return 0
}

type FMonitor struct {
	Value            *float32 `protobuf:"fixed32,1,req,name=value" json:"value,omitempty"`
	Min              *float32 `protobuf:"fixed32,2,req,name=min" json:"min,omitempty"`
	Max              *float32 `protobuf:"fixed32,3,req,name=max" json:"max,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FMonitor) Reset()                    { *m = FMonitor{} }
func (m *FMonitor) String() string            { return proto.CompactTextString(m) }
func (*FMonitor) ProtoMessage()               {}
func (*FMonitor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FMonitor) GetValue() float32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *FMonitor) GetMin() float32 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *FMonitor) GetMax() float32 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

type EMonitor struct {
	PValue           *float32 `protobuf:"fixed32,1,req,name=p_value,json=pValue" json:"p_value,omitempty"`
	NValue           *float32 `protobuf:"fixed32,2,req,name=n_value,json=nValue" json:"n_value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EMonitor) Reset()                    { *m = EMonitor{} }
func (m *EMonitor) String() string            { return proto.CompactTextString(m) }
func (*EMonitor) ProtoMessage()               {}
func (*EMonitor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EMonitor) GetPValue() float32 {
	if m != nil && m.PValue != nil {
		return *m.PValue
	}
	return 0
}

func (m *EMonitor) GetNValue() float32 {
	if m != nil && m.NValue != nil {
		return *m.NValue
	}
	return 0
}

type ThdsMonitor struct {
	Value            *float32 `protobuf:"fixed32,1,req,name=value" json:"value,omitempty"`
	Min              *float32 `protobuf:"fixed32,2,req,name=min" json:"min,omitempty"`
	Max              *float32 `protobuf:"fixed32,3,req,name=max" json:"max,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ThdsMonitor) Reset()                    { *m = ThdsMonitor{} }
func (m *ThdsMonitor) String() string            { return proto.CompactTextString(m) }
func (*ThdsMonitor) ProtoMessage()               {}
func (*ThdsMonitor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ThdsMonitor) GetValue() float32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *ThdsMonitor) GetMin() float32 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *ThdsMonitor) GetMax() float32 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

type UpdateConfig struct {
	Id               *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateConfig) Reset()                    { *m = UpdateConfig{} }
func (m *UpdateConfig) String() string            { return proto.CompactTextString(m) }
func (*UpdateConfig) ProtoMessage()               {}
func (*UpdateConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateConfig) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *UpdateConfig) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type MonMeasure struct {
	SrcId            *uint32        `protobuf:"varint,1,req,name=src_id,json=srcId" json:"src_id,omitempty"`
	Timestamp        *uint64        `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	U                []*UMonitor    `protobuf:"bytes,3,rep,name=u" json:"u,omitempty"`
	I                []*IMonitor    `protobuf:"bytes,4,rep,name=i" json:"i,omitempty"`
	Freq             []*FMonitor    `protobuf:"bytes,5,rep,name=freq" json:"freq,omitempty"`
	RealEng          []*EMonitor    `protobuf:"bytes,6,rep,name=real_eng,json=realEng" json:"real_eng,omitempty"`
	ReactEng         []*EMonitor    `protobuf:"bytes,7,rep,name=react_eng,json=reactEng" json:"react_eng,omitempty"`
	Thds             []*ThdsMonitor `protobuf:"bytes,8,rep,name=thds" json:"thds,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *MonMeasure) Reset()                    { *m = MonMeasure{} }
func (m *MonMeasure) String() string            { return proto.CompactTextString(m) }
func (*MonMeasure) ProtoMessage()               {}
func (*MonMeasure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MonMeasure) GetSrcId() uint32 {
	if m != nil && m.SrcId != nil {
		return *m.SrcId
	}
	return 0
}

func (m *MonMeasure) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *MonMeasure) GetU() []*UMonitor {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *MonMeasure) GetI() []*IMonitor {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *MonMeasure) GetFreq() []*FMonitor {
	if m != nil {
		return m.Freq
	}
	return nil
}

func (m *MonMeasure) GetRealEng() []*EMonitor {
	if m != nil {
		return m.RealEng
	}
	return nil
}

func (m *MonMeasure) GetReactEng() []*EMonitor {
	if m != nil {
		return m.ReactEng
	}
	return nil
}

func (m *MonMeasure) GetThds() []*ThdsMonitor {
	if m != nil {
		return m.Thds
	}
	return nil
}

type RegMeasure struct {
	Timestamp        *uint64     `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	U                []*UMonitor `protobuf:"bytes,2,rep,name=u" json:"u,omitempty"`
	I                []*IMonitor `protobuf:"bytes,3,rep,name=i" json:"i,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RegMeasure) Reset()                    { *m = RegMeasure{} }
func (m *RegMeasure) String() string            { return proto.CompactTextString(m) }
func (*RegMeasure) ProtoMessage()               {}
func (*RegMeasure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RegMeasure) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *RegMeasure) GetU() []*UMonitor {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *RegMeasure) GetI() []*IMonitor {
	if m != nil {
		return m.I
	}
	return nil
}

type GenericData struct {
	DataType         *uint32  `protobuf:"varint,1,req,name=data_type,json=dataType" json:"data_type,omitempty"`
	PhasePosition    *uint32  `protobuf:"varint,2,req,name=phase_position,json=phasePosition" json:"phase_position,omitempty"`
	Value            *float32 `protobuf:"fixed32,3,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GenericData) Reset()                    { *m = GenericData{} }
func (m *GenericData) String() string            { return proto.CompactTextString(m) }
func (*GenericData) ProtoMessage()               {}
func (*GenericData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GenericData) GetDataType() uint32 {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return 0
}

func (m *GenericData) GetPhasePosition() uint32 {
	if m != nil && m.PhasePosition != nil {
		return *m.PhasePosition
	}
	return 0
}

func (m *GenericData) GetValue() float32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Command)(nil), "command")
	proto.RegisterType((*UMonitor)(nil), "u_monitor")
	proto.RegisterType((*IMonitor)(nil), "i_monitor")
	proto.RegisterType((*FMonitor)(nil), "f_monitor")
	proto.RegisterType((*EMonitor)(nil), "e_monitor")
	proto.RegisterType((*ThdsMonitor)(nil), "thds_monitor")
	proto.RegisterType((*UpdateConfig)(nil), "update_config")
	proto.RegisterType((*MonMeasure)(nil), "mon_measure")
	proto.RegisterType((*RegMeasure)(nil), "reg_measure")
	proto.RegisterType((*GenericData)(nil), "generic_data")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x5d, 0x8b, 0xd4, 0x30,
	0x14, 0x86, 0x99, 0xb6, 0x33, 0x9d, 0x9e, 0x99, 0x2e, 0x12, 0x95, 0x0d, 0x28, 0x32, 0x16, 0x16,
	0xe7, 0x6a, 0x40, 0xc1, 0x4b, 0x2f, 0x17, 0x14, 0x59, 0x90, 0x20, 0xde, 0x86, 0xd0, 0x66, 0x3a,
	0x81, 0xcd, 0x87, 0x49, 0x2a, 0xe3, 0x4f, 0xf0, 0x5f, 0x4b, 0xd2, 0xd0, 0x8e, 0xb8, 0xe2, 0xcd,
	0x5e, 0x35, 0x79, 0xdf, 0x73, 0x1e, 0xce, 0x47, 0x03, 0xdb, 0x56, 0x4b, 0xa9, 0xd5, 0xc1, 0x58,
	0xed, 0x75, 0xf3, 0x16, 0xca, 0x70, 0x67, 0xaa, 0x43, 0x08, 0x0a, 0xff, 0xd3, 0x70, 0xbc, 0xd8,
	0x65, 0xfb, 0x9a, 0xc4, 0x73, 0xd0, 0x98, 0xed, 0x1d, 0xce, 0x76, 0xf9, 0xbe, 0x22, 0xf1, 0xdc,
	0xdc, 0x42, 0x35, 0x50, 0xa9, 0x95, 0xf0, 0xda, 0xa2, 0x67, 0xb0, 0xfc, 0xc1, 0xee, 0x87, 0x31,
	0x2b, 0x23, 0xe3, 0x05, 0x3d, 0x81, 0x5c, 0x0a, 0x85, 0xb3, 0xa8, 0x85, 0x63, 0x54, 0xd8, 0x19,
	0xe7, 0x49, 0x61, 0xe7, 0xe6, 0x33, 0x54, 0xe2, 0x3f, 0x98, 0xa7, 0xb0, 0x34, 0x34, 0xa4, 0x8d,
	0xa0, 0xc2, 0xdc, 0xb1, 0x73, 0x10, 0x15, 0x9d, 0x59, 0x85, 0xba, 0x63, 0xe7, 0x50, 0xd3, 0xf1,
	0x11, 0x6a, 0xfa, 0x00, 0x15, 0x9f, 0x30, 0xd7, 0x50, 0x1a, 0x7a, 0x09, 0x5a, 0x99, 0x6f, 0x91,
	0x74, 0x0d, 0xa5, 0x4a, 0xc6, 0x48, 0x5b, 0xa9, 0x68, 0x34, 0x1f, 0x61, 0xeb, 0x4f, 0x9d, 0x7b,
	0x84, 0x42, 0xde, 0x43, 0x3d, 0x98, 0x8e, 0x79, 0x4e, 0x5b, 0xad, 0x8e, 0xa2, 0x47, 0x57, 0x90,
	0x89, 0x2e, 0xad, 0x26, 0x13, 0xdd, 0x8c, 0x0e, 0x98, 0x2a, 0xa1, 0x9b, 0x5f, 0x19, 0x6c, 0xa4,
	0x56, 0x54, 0x72, 0xe6, 0x06, 0xcb, 0xd1, 0x73, 0x58, 0x39, 0xdb, 0xd2, 0x29, 0x73, 0xe9, 0x6c,
	0xfb, 0xa9, 0x43, 0x2f, 0xa1, 0xf2, 0x42, 0x72, 0xe7, 0x99, 0x34, 0x11, 0x50, 0x90, 0x59, 0x40,
	0x18, 0x16, 0x03, 0xce, 0x77, 0xf9, 0x7e, 0xf3, 0x0e, 0x0e, 0xd3, 0xa6, 0xc9, 0x62, 0x08, 0x8e,
	0xc0, 0x45, 0x72, 0xc4, 0xec, 0x08, 0xf4, 0x0a, 0x8a, 0xa3, 0xe5, 0xdf, 0xf1, 0x32, 0x99, 0xd3,
	0x32, 0x48, 0xd4, 0xd1, 0x0d, 0xac, 0x2d, 0x67, 0xf7, 0x94, 0xab, 0x1e, 0xaf, 0x52, 0xcc, 0x34,
	0x69, 0x52, 0x06, 0xef, 0x56, 0xf5, 0xe8, 0x0d, 0x54, 0x96, 0xb3, 0xd6, 0xc7, 0xb8, 0xf2, 0xaf,
	0xb8, 0x75, 0x34, 0x43, 0xe0, 0x6b, 0x28, 0xc2, 0xa4, 0xf1, 0x3a, 0xc6, 0xd4, 0x87, 0xcb, 0xb1,
	0x93, 0x68, 0x35, 0x14, 0x36, 0x96, 0xf7, 0xd3, 0x28, 0xfe, 0xe8, 0x79, 0xf1, 0x60, 0xcf, 0xd9,
	0x3f, 0x7b, 0xce, 0x1f, 0xe8, 0xb9, 0x39, 0xc1, 0xb6, 0xe7, 0x8a, 0x5b, 0xd1, 0xd2, 0x8e, 0x79,
	0x86, 0x5e, 0x40, 0x15, 0xbe, 0xf4, 0xe2, 0x11, 0xad, 0x83, 0xf0, 0x35, 0x3c, 0xa4, 0x1b, 0xb8,
	0x32, 0x27, 0xe6, 0x38, 0x35, 0xda, 0x09, 0x2f, 0xf4, 0xb8, 0xff, 0x9a, 0xd4, 0x51, 0xfd, 0x92,
	0xc4, 0x79, 0xad, 0xf9, 0xc5, 0x1f, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xfe, 0x74, 0xca,
	0xb3, 0x03, 0x00, 0x00,
}
