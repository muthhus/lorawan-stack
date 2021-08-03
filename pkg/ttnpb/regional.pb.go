// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/regional.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ConcentratorConfig struct {
	Channels             []*ConcentratorConfig_Channel           `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
	LoraStandardChannel  *ConcentratorConfig_LoRaStandardChannel `protobuf:"bytes,2,opt,name=lora_standard_channel,json=loraStandardChannel,proto3" json:"lora_standard_channel,omitempty"`
	FskChannel           *ConcentratorConfig_FSKChannel          `protobuf:"bytes,3,opt,name=fsk_channel,json=fskChannel,proto3" json:"fsk_channel,omitempty"`
	Lbt                  *ConcentratorConfig_LBTConfiguration    `protobuf:"bytes,4,opt,name=lbt,proto3" json:"lbt,omitempty"`
	PingSlot             *ConcentratorConfig_Channel             `protobuf:"bytes,5,opt,name=ping_slot,json=pingSlot,proto3" json:"ping_slot,omitempty"`
	Radios               []*GatewayRadio                         `protobuf:"bytes,6,rep,name=radios,proto3" json:"radios,omitempty"`
	ClockSource          uint32                                  `protobuf:"varint,7,opt,name=clock_source,json=clockSource,proto3" json:"clock_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ConcentratorConfig) Reset()      { *m = ConcentratorConfig{} }
func (*ConcentratorConfig) ProtoMessage() {}
func (*ConcentratorConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe48e1cbcf8ee88, []int{0}
}
func (m *ConcentratorConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcentratorConfig.Unmarshal(m, b)
}
func (m *ConcentratorConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcentratorConfig.Marshal(b, m, deterministic)
}
func (m *ConcentratorConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcentratorConfig.Merge(m, src)
}
func (m *ConcentratorConfig) XXX_Size() int {
	return xxx_messageInfo_ConcentratorConfig.Size(m)
}
func (m *ConcentratorConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcentratorConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConcentratorConfig proto.InternalMessageInfo

func (m *ConcentratorConfig) GetChannels() []*ConcentratorConfig_Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *ConcentratorConfig) GetLoraStandardChannel() *ConcentratorConfig_LoRaStandardChannel {
	if m != nil {
		return m.LoraStandardChannel
	}
	return nil
}

func (m *ConcentratorConfig) GetFskChannel() *ConcentratorConfig_FSKChannel {
	if m != nil {
		return m.FskChannel
	}
	return nil
}

func (m *ConcentratorConfig) GetLbt() *ConcentratorConfig_LBTConfiguration {
	if m != nil {
		return m.Lbt
	}
	return nil
}

func (m *ConcentratorConfig) GetPingSlot() *ConcentratorConfig_Channel {
	if m != nil {
		return m.PingSlot
	}
	return nil
}

func (m *ConcentratorConfig) GetRadios() []*GatewayRadio {
	if m != nil {
		return m.Radios
	}
	return nil
}

func (m *ConcentratorConfig) GetClockSource() uint32 {
	if m != nil {
		return m.ClockSource
	}
	return 0
}

type ConcentratorConfig_Channel struct {
	// Frequency (Hz).
	Frequency            uint64   `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Radio                uint32   `protobuf:"varint,2,opt,name=radio,proto3" json:"radio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcentratorConfig_Channel) Reset()      { *m = ConcentratorConfig_Channel{} }
func (*ConcentratorConfig_Channel) ProtoMessage() {}
func (*ConcentratorConfig_Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe48e1cbcf8ee88, []int{0, 0}
}
func (m *ConcentratorConfig_Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcentratorConfig_Channel.Unmarshal(m, b)
}
func (m *ConcentratorConfig_Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcentratorConfig_Channel.Marshal(b, m, deterministic)
}
func (m *ConcentratorConfig_Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcentratorConfig_Channel.Merge(m, src)
}
func (m *ConcentratorConfig_Channel) XXX_Size() int {
	return xxx_messageInfo_ConcentratorConfig_Channel.Size(m)
}
func (m *ConcentratorConfig_Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcentratorConfig_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_ConcentratorConfig_Channel proto.InternalMessageInfo

func (m *ConcentratorConfig_Channel) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *ConcentratorConfig_Channel) GetRadio() uint32 {
	if m != nil {
		return m.Radio
	}
	return 0
}

type ConcentratorConfig_LoRaStandardChannel struct {
	// Frequency (Hz).
	Frequency uint64 `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Radio     uint32 `protobuf:"varint,2,opt,name=radio,proto3" json:"radio,omitempty"`
	// Bandwidth (Hz).
	Bandwidth            uint32   `protobuf:"varint,3,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	SpreadingFactor      uint32   `protobuf:"varint,4,opt,name=spreading_factor,json=spreadingFactor,proto3" json:"spreading_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcentratorConfig_LoRaStandardChannel) Reset() {
	*m = ConcentratorConfig_LoRaStandardChannel{}
}
func (*ConcentratorConfig_LoRaStandardChannel) ProtoMessage() {}
func (*ConcentratorConfig_LoRaStandardChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe48e1cbcf8ee88, []int{0, 1}
}
func (m *ConcentratorConfig_LoRaStandardChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcentratorConfig_LoRaStandardChannel.Unmarshal(m, b)
}
func (m *ConcentratorConfig_LoRaStandardChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcentratorConfig_LoRaStandardChannel.Marshal(b, m, deterministic)
}
func (m *ConcentratorConfig_LoRaStandardChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcentratorConfig_LoRaStandardChannel.Merge(m, src)
}
func (m *ConcentratorConfig_LoRaStandardChannel) XXX_Size() int {
	return xxx_messageInfo_ConcentratorConfig_LoRaStandardChannel.Size(m)
}
func (m *ConcentratorConfig_LoRaStandardChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcentratorConfig_LoRaStandardChannel.DiscardUnknown(m)
}

var xxx_messageInfo_ConcentratorConfig_LoRaStandardChannel proto.InternalMessageInfo

func (m *ConcentratorConfig_LoRaStandardChannel) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *ConcentratorConfig_LoRaStandardChannel) GetRadio() uint32 {
	if m != nil {
		return m.Radio
	}
	return 0
}

func (m *ConcentratorConfig_LoRaStandardChannel) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *ConcentratorConfig_LoRaStandardChannel) GetSpreadingFactor() uint32 {
	if m != nil {
		return m.SpreadingFactor
	}
	return 0
}

type ConcentratorConfig_FSKChannel struct {
	// Frequency (Hz).
	Frequency            uint64   `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Radio                uint32   `protobuf:"varint,2,opt,name=radio,proto3" json:"radio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcentratorConfig_FSKChannel) Reset()      { *m = ConcentratorConfig_FSKChannel{} }
func (*ConcentratorConfig_FSKChannel) ProtoMessage() {}
func (*ConcentratorConfig_FSKChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe48e1cbcf8ee88, []int{0, 2}
}
func (m *ConcentratorConfig_FSKChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcentratorConfig_FSKChannel.Unmarshal(m, b)
}
func (m *ConcentratorConfig_FSKChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcentratorConfig_FSKChannel.Marshal(b, m, deterministic)
}
func (m *ConcentratorConfig_FSKChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcentratorConfig_FSKChannel.Merge(m, src)
}
func (m *ConcentratorConfig_FSKChannel) XXX_Size() int {
	return xxx_messageInfo_ConcentratorConfig_FSKChannel.Size(m)
}
func (m *ConcentratorConfig_FSKChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcentratorConfig_FSKChannel.DiscardUnknown(m)
}

var xxx_messageInfo_ConcentratorConfig_FSKChannel proto.InternalMessageInfo

func (m *ConcentratorConfig_FSKChannel) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *ConcentratorConfig_FSKChannel) GetRadio() uint32 {
	if m != nil {
		return m.Radio
	}
	return 0
}

type ConcentratorConfig_LBTConfiguration struct {
	// Received signal strength (dBm).
	RssiTarget float32 `protobuf:"fixed32,1,opt,name=rssi_target,json=rssiTarget,proto3" json:"rssi_target,omitempty"`
	// Received signal strength offset (dBm).
	RssiOffset           float32       `protobuf:"fixed32,2,opt,name=rssi_offset,json=rssiOffset,proto3" json:"rssi_offset,omitempty"`
	ScanTime             time.Duration `protobuf:"bytes,3,opt,name=scan_time,json=scanTime,proto3,stdduration" json:"scan_time"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConcentratorConfig_LBTConfiguration) Reset()      { *m = ConcentratorConfig_LBTConfiguration{} }
func (*ConcentratorConfig_LBTConfiguration) ProtoMessage() {}
func (*ConcentratorConfig_LBTConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe48e1cbcf8ee88, []int{0, 3}
}
func (m *ConcentratorConfig_LBTConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcentratorConfig_LBTConfiguration.Unmarshal(m, b)
}
func (m *ConcentratorConfig_LBTConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcentratorConfig_LBTConfiguration.Marshal(b, m, deterministic)
}
func (m *ConcentratorConfig_LBTConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcentratorConfig_LBTConfiguration.Merge(m, src)
}
func (m *ConcentratorConfig_LBTConfiguration) XXX_Size() int {
	return xxx_messageInfo_ConcentratorConfig_LBTConfiguration.Size(m)
}
func (m *ConcentratorConfig_LBTConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcentratorConfig_LBTConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ConcentratorConfig_LBTConfiguration proto.InternalMessageInfo

func (m *ConcentratorConfig_LBTConfiguration) GetRssiTarget() float32 {
	if m != nil {
		return m.RssiTarget
	}
	return 0
}

func (m *ConcentratorConfig_LBTConfiguration) GetRssiOffset() float32 {
	if m != nil {
		return m.RssiOffset
	}
	return 0
}

func (m *ConcentratorConfig_LBTConfiguration) GetScanTime() time.Duration {
	if m != nil {
		return m.ScanTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ConcentratorConfig)(nil), "ttn.lorawan.v3.ConcentratorConfig")
	golang_proto.RegisterType((*ConcentratorConfig)(nil), "ttn.lorawan.v3.ConcentratorConfig")
	proto.RegisterType((*ConcentratorConfig_Channel)(nil), "ttn.lorawan.v3.ConcentratorConfig.Channel")
	golang_proto.RegisterType((*ConcentratorConfig_Channel)(nil), "ttn.lorawan.v3.ConcentratorConfig.Channel")
	proto.RegisterType((*ConcentratorConfig_LoRaStandardChannel)(nil), "ttn.lorawan.v3.ConcentratorConfig.LoRaStandardChannel")
	golang_proto.RegisterType((*ConcentratorConfig_LoRaStandardChannel)(nil), "ttn.lorawan.v3.ConcentratorConfig.LoRaStandardChannel")
	proto.RegisterType((*ConcentratorConfig_FSKChannel)(nil), "ttn.lorawan.v3.ConcentratorConfig.FSKChannel")
	golang_proto.RegisterType((*ConcentratorConfig_FSKChannel)(nil), "ttn.lorawan.v3.ConcentratorConfig.FSKChannel")
	proto.RegisterType((*ConcentratorConfig_LBTConfiguration)(nil), "ttn.lorawan.v3.ConcentratorConfig.LBTConfiguration")
	golang_proto.RegisterType((*ConcentratorConfig_LBTConfiguration)(nil), "ttn.lorawan.v3.ConcentratorConfig.LBTConfiguration")
}

func init() { proto.RegisterFile("lorawan-stack/api/regional.proto", fileDescriptor_cfe48e1cbcf8ee88) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/regional.proto", fileDescriptor_cfe48e1cbcf8ee88)
}

var fileDescriptor_cfe48e1cbcf8ee88 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x3d, 0x6f, 0x13, 0x3f,
	0x1c, 0x8e, 0xfb, 0x5e, 0xe7, 0x9f, 0x3f, 0x95, 0x0b, 0xd2, 0x11, 0x55, 0x97, 0xc0, 0x14, 0x90,
	0xea, 0x93, 0x1a, 0xc4, 0x86, 0x54, 0xb5, 0xd0, 0x0e, 0xbc, 0x49, 0x4e, 0x27, 0x96, 0xc8, 0xb9,
	0xf3, 0x39, 0x26, 0x57, 0x3b, 0xd8, 0xbf, 0x34, 0xea, 0xc6, 0x47, 0x40, 0x0c, 0x88, 0x8f, 0xc0,
	0xc8, 0x47, 0x60, 0xe4, 0x23, 0x30, 0x81, 0x48, 0x17, 0x46, 0x3e, 0x02, 0x3a, 0xdf, 0x25, 0x51,
	0x5b, 0x86, 0xd2, 0xcd, 0x7e, 0xfc, 0xbc, 0xd8, 0x7e, 0xce, 0x87, 0x9b, 0x99, 0xb1, 0x7c, 0xcc,
	0xf5, 0xb6, 0x03, 0x1e, 0x0f, 0x22, 0x3e, 0x54, 0x91, 0x15, 0x52, 0x19, 0xcd, 0x33, 0x3a, 0xb4,
	0x06, 0x0c, 0xf9, 0x1f, 0x40, 0xd3, 0x92, 0x45, 0x4f, 0xda, 0xf5, 0x6d, 0xa9, 0xa0, 0x3f, 0xea,
	0xd1, 0xd8, 0x1c, 0x47, 0xd2, 0x48, 0x13, 0x79, 0x5a, 0x6f, 0x94, 0xfa, 0x99, 0x9f, 0xf8, 0x51,
	0x21, 0xaf, 0x87, 0xd2, 0x18, 0x99, 0x89, 0x39, 0x2b, 0x19, 0x59, 0x0e, 0xca, 0xe8, 0x72, 0xbd,
	0x71, 0x79, 0x03, 0x92, 0x83, 0x18, 0xf3, 0xd3, 0x82, 0x70, 0xf7, 0xf3, 0x2a, 0x26, 0xfb, 0x46,
	0xc7, 0x42, 0x83, 0xe5, 0x60, 0xec, 0xbe, 0xd1, 0xa9, 0x92, 0xe4, 0x00, 0xaf, 0xc5, 0x7d, 0xae,
	0xb5, 0xc8, 0x5c, 0x80, 0x9a, 0x8b, 0xad, 0xea, 0xce, 0x7d, 0x7a, 0x7e, 0xa7, 0xf4, 0xb2, 0x8a,
	0xee, 0x17, 0x12, 0x36, 0xd3, 0x92, 0xd7, 0xf8, 0x56, 0x2e, 0xe9, 0x3a, 0xe0, 0x3a, 0xe1, 0x36,
	0xe9, 0x96, 0x2b, 0xc1, 0x42, 0x13, 0xb5, 0xaa, 0x3b, 0x0f, 0xaf, 0x60, 0xfa, 0xcc, 0x30, 0xde,
	0x29, 0xe5, 0xd3, 0x80, 0xcd, 0x5c, 0x72, 0x01, 0x24, 0x2f, 0x70, 0x35, 0x75, 0x83, 0x59, 0xc2,
	0xa2, 0x4f, 0xd8, 0xbe, 0x42, 0xc2, 0x41, 0xe7, 0xe9, 0xd4, 0x18, 0xa7, 0x6e, 0x30, 0xf5, 0x7b,
	0x82, 0x17, 0xb3, 0x1e, 0x04, 0x4b, 0xde, 0xa7, 0x7d, 0x95, 0x9d, 0xee, 0x1d, 0x15, 0xa3, 0xb2,
	0x03, 0x96, 0xeb, 0xc9, 0x21, 0x5e, 0x1f, 0x2a, 0x2d, 0xbb, 0x2e, 0x33, 0x10, 0x2c, 0x7b, 0xb3,
	0x7f, 0xba, 0xcb, 0x5c, 0xdc, 0xc9, 0x0c, 0x90, 0x07, 0x78, 0xc5, 0xf2, 0x44, 0x19, 0x17, 0xac,
	0xf8, 0x46, 0xb6, 0x2e, 0xba, 0x1c, 0x16, 0xcd, 0xb2, 0x9c, 0xc4, 0x4a, 0x2e, 0xb9, 0x83, 0xff,
	0x8b, 0x33, 0x13, 0x0f, 0xba, 0xce, 0x8c, 0x6c, 0x2c, 0x82, 0xd5, 0x26, 0x6a, 0xd5, 0x58, 0xd5,
	0x63, 0x1d, 0x0f, 0xd5, 0x1f, 0xe1, 0xd5, 0xe9, 0x99, 0xb7, 0xf0, 0x7a, 0x6a, 0xc5, 0x9b, 0x91,
	0xd0, 0xf1, 0x69, 0x80, 0x9a, 0xa8, 0xb5, 0xc4, 0xe6, 0x00, 0xb9, 0x89, 0x97, 0xbd, 0xab, 0x6f,
	0xaf, 0xc6, 0x8a, 0x49, 0xfd, 0x3d, 0xc2, 0x9b, 0x7f, 0x29, 0xe9, 0x3a, 0x5e, 0xb9, 0xa6, 0xc7,
	0x75, 0x32, 0x56, 0x09, 0xf4, 0x7d, 0x83, 0x35, 0x36, 0x07, 0xc8, 0x3d, 0xbc, 0xe1, 0x86, 0x56,
	0xf0, 0x24, 0xbf, 0xcf, 0x94, 0xc7, 0x60, 0xac, 0xaf, 0xa7, 0xc6, 0x6e, 0xcc, 0xf0, 0x03, 0x0f,
	0xd7, 0x77, 0x31, 0x9e, 0xd7, 0x7a, 0xad, 0x63, 0x7d, 0x40, 0x78, 0xe3, 0x62, 0xa3, 0xa4, 0x81,
	0xab, 0xd6, 0x39, 0xd5, 0x05, 0x6e, 0xa5, 0x00, 0x6f, 0xb5, 0xc0, 0x70, 0x0e, 0x1d, 0x79, 0x64,
	0x46, 0x30, 0x69, 0xea, 0x04, 0x78, 0xc7, 0x92, 0xf0, 0xd2, 0x23, 0x64, 0x17, 0xaf, 0xbb, 0x98,
	0xeb, 0x2e, 0xa8, 0x63, 0x51, 0x7e, 0xa3, 0xb7, 0x69, 0xf1, 0x8a, 0xe9, 0xf4, 0x15, 0xd3, 0xc7,
	0x65, 0xde, 0xde, 0xda, 0xd7, 0xef, 0x8d, 0xca, 0xc7, 0x1f, 0x0d, 0xc4, 0xd6, 0x72, 0xd5, 0x91,
	0x3a, 0x16, 0x7b, 0xcf, 0xbf, 0xfd, 0x0c, 0x2b, 0x6f, 0x27, 0x21, 0xfa, 0x34, 0x09, 0xd1, 0xaf,
	0x49, 0x58, 0xf9, 0x3d, 0x09, 0xd1, 0xbb, 0xb3, 0xb0, 0xf2, 0xe5, 0x2c, 0x44, 0xaf, 0x22, 0x69,
	0x28, 0xf4, 0x05, 0xf4, 0x95, 0x96, 0x8e, 0x6a, 0x01, 0x63, 0x63, 0x07, 0xd1, 0xf9, 0xdf, 0xc0,
	0x49, 0x3b, 0x1a, 0x0e, 0x64, 0x04, 0xa0, 0x87, 0xbd, 0xde, 0x8a, 0x4f, 0x6d, 0xff, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xc2, 0xa7, 0x0a, 0xc3, 0xac, 0x04, 0x00, 0x00,
}

func (this *ConcentratorConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConcentratorConfig)
	if !ok {
		that2, ok := that.(ConcentratorConfig)
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
	if len(this.Channels) != len(that1.Channels) {
		return false
	}
	for i := range this.Channels {
		if !this.Channels[i].Equal(that1.Channels[i]) {
			return false
		}
	}
	if !this.LoraStandardChannel.Equal(that1.LoraStandardChannel) {
		return false
	}
	if !this.FskChannel.Equal(that1.FskChannel) {
		return false
	}
	if !this.Lbt.Equal(that1.Lbt) {
		return false
	}
	if !this.PingSlot.Equal(that1.PingSlot) {
		return false
	}
	if len(this.Radios) != len(that1.Radios) {
		return false
	}
	for i := range this.Radios {
		if !this.Radios[i].Equal(that1.Radios[i]) {
			return false
		}
	}
	if this.ClockSource != that1.ClockSource {
		return false
	}
	return true
}
func (this *ConcentratorConfig_Channel) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConcentratorConfig_Channel)
	if !ok {
		that2, ok := that.(ConcentratorConfig_Channel)
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
	if this.Frequency != that1.Frequency {
		return false
	}
	if this.Radio != that1.Radio {
		return false
	}
	return true
}
func (this *ConcentratorConfig_LoRaStandardChannel) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConcentratorConfig_LoRaStandardChannel)
	if !ok {
		that2, ok := that.(ConcentratorConfig_LoRaStandardChannel)
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
	if this.Frequency != that1.Frequency {
		return false
	}
	if this.Radio != that1.Radio {
		return false
	}
	if this.Bandwidth != that1.Bandwidth {
		return false
	}
	if this.SpreadingFactor != that1.SpreadingFactor {
		return false
	}
	return true
}
func (this *ConcentratorConfig_FSKChannel) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConcentratorConfig_FSKChannel)
	if !ok {
		that2, ok := that.(ConcentratorConfig_FSKChannel)
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
	if this.Frequency != that1.Frequency {
		return false
	}
	if this.Radio != that1.Radio {
		return false
	}
	return true
}
func (this *ConcentratorConfig_LBTConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConcentratorConfig_LBTConfiguration)
	if !ok {
		that2, ok := that.(ConcentratorConfig_LBTConfiguration)
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
	if this.RssiTarget != that1.RssiTarget {
		return false
	}
	if this.RssiOffset != that1.RssiOffset {
		return false
	}
	if this.ScanTime != that1.ScanTime {
		return false
	}
	return true
}
func (this *ConcentratorConfig) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForChannels := "[]*ConcentratorConfig_Channel{"
	for _, f := range this.Channels {
		repeatedStringForChannels += strings.Replace(fmt.Sprintf("%v", f), "ConcentratorConfig_Channel", "ConcentratorConfig_Channel", 1) + ","
	}
	repeatedStringForChannels += "}"
	repeatedStringForRadios := "[]*GatewayRadio{"
	for _, f := range this.Radios {
		repeatedStringForRadios += strings.Replace(fmt.Sprintf("%v", f), "GatewayRadio", "GatewayRadio", 1) + ","
	}
	repeatedStringForRadios += "}"
	s := strings.Join([]string{`&ConcentratorConfig{`,
		`Channels:` + repeatedStringForChannels + `,`,
		`LoraStandardChannel:` + strings.Replace(fmt.Sprintf("%v", this.LoraStandardChannel), "ConcentratorConfig_LoRaStandardChannel", "ConcentratorConfig_LoRaStandardChannel", 1) + `,`,
		`FskChannel:` + strings.Replace(fmt.Sprintf("%v", this.FskChannel), "ConcentratorConfig_FSKChannel", "ConcentratorConfig_FSKChannel", 1) + `,`,
		`Lbt:` + strings.Replace(fmt.Sprintf("%v", this.Lbt), "ConcentratorConfig_LBTConfiguration", "ConcentratorConfig_LBTConfiguration", 1) + `,`,
		`PingSlot:` + strings.Replace(fmt.Sprintf("%v", this.PingSlot), "ConcentratorConfig_Channel", "ConcentratorConfig_Channel", 1) + `,`,
		`Radios:` + repeatedStringForRadios + `,`,
		`ClockSource:` + fmt.Sprintf("%v", this.ClockSource) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConcentratorConfig_Channel) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConcentratorConfig_Channel{`,
		`Frequency:` + fmt.Sprintf("%v", this.Frequency) + `,`,
		`Radio:` + fmt.Sprintf("%v", this.Radio) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConcentratorConfig_LoRaStandardChannel) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConcentratorConfig_LoRaStandardChannel{`,
		`Frequency:` + fmt.Sprintf("%v", this.Frequency) + `,`,
		`Radio:` + fmt.Sprintf("%v", this.Radio) + `,`,
		`Bandwidth:` + fmt.Sprintf("%v", this.Bandwidth) + `,`,
		`SpreadingFactor:` + fmt.Sprintf("%v", this.SpreadingFactor) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConcentratorConfig_FSKChannel) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConcentratorConfig_FSKChannel{`,
		`Frequency:` + fmt.Sprintf("%v", this.Frequency) + `,`,
		`Radio:` + fmt.Sprintf("%v", this.Radio) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConcentratorConfig_LBTConfiguration) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConcentratorConfig_LBTConfiguration{`,
		`RssiTarget:` + fmt.Sprintf("%v", this.RssiTarget) + `,`,
		`RssiOffset:` + fmt.Sprintf("%v", this.RssiOffset) + `,`,
		`ScanTime:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ScanTime), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRegional(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
