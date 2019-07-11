// Code generated by protoc-gen-go.
// source: ThesisProtocol.proto
// DO NOT EDIT!

/*
Package ThesisProtocol is a generated protocol buffer package.

It is generated from these files:
	ThesisProtocol.proto

It has these top-level messages:
	ReceivedMessage
*/
package ThesisProtocol

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

type ReceivedMessage struct {
	DeviceX          *int32                    `protobuf:"varint,1,req,name=deviceX" json:"deviceX,omitempty"`
	DeviceY          *int32                    `protobuf:"varint,2,req,name=deviceY" json:"deviceY,omitempty"`
	DeviceVoltage    *float32                  `protobuf:"fixed32,3,opt,name=deviceVoltage" json:"deviceVoltage,omitempty"`
	LastPost         *bool                     `protobuf:"varint,4,opt,name=lastPost" json:"lastPost,omitempty"`
	SubmitTime       *int32                    `protobuf:"varint,5,opt,name=submitTime" json:"submitTime,omitempty"`
	WifiTime         *int32                    `protobuf:"varint,6,opt,name=wifiTime" json:"wifiTime,omitempty"`
	Devices          []*ReceivedMessage_Device `protobuf:"bytes,7,rep,name=devices" json:"devices,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *ReceivedMessage) Reset()                    { *m = ReceivedMessage{} }
func (m *ReceivedMessage) String() string            { return proto.CompactTextString(m) }
func (*ReceivedMessage) ProtoMessage()               {}
func (*ReceivedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReceivedMessage) GetDeviceX() int32 {
	if m != nil && m.DeviceX != nil {
		return *m.DeviceX
	}
	return 0
}

func (m *ReceivedMessage) GetDeviceY() int32 {
	if m != nil && m.DeviceY != nil {
		return *m.DeviceY
	}
	return 0
}

func (m *ReceivedMessage) GetDeviceVoltage() float32 {
	if m != nil && m.DeviceVoltage != nil {
		return *m.DeviceVoltage
	}
	return 0
}

func (m *ReceivedMessage) GetLastPost() bool {
	if m != nil && m.LastPost != nil {
		return *m.LastPost
	}
	return false
}

func (m *ReceivedMessage) GetSubmitTime() int32 {
	if m != nil && m.SubmitTime != nil {
		return *m.SubmitTime
	}
	return 0
}

func (m *ReceivedMessage) GetWifiTime() int32 {
	if m != nil && m.WifiTime != nil {
		return *m.WifiTime
	}
	return 0
}

func (m *ReceivedMessage) GetDevices() []*ReceivedMessage_Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type ReceivedMessage_Device struct {
	Mac              *string                      `protobuf:"bytes,1,req,name=mac" json:"mac,omitempty"`
	Datapoints       []*ReceivedMessage_DataPoint `protobuf:"bytes,2,rep,name=datapoints" json:"datapoints,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *ReceivedMessage_Device) Reset()                    { *m = ReceivedMessage_Device{} }
func (m *ReceivedMessage_Device) String() string            { return proto.CompactTextString(m) }
func (*ReceivedMessage_Device) ProtoMessage()               {}
func (*ReceivedMessage_Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ReceivedMessage_Device) GetMac() string {
	if m != nil && m.Mac != nil {
		return *m.Mac
	}
	return ""
}

func (m *ReceivedMessage_Device) GetDatapoints() []*ReceivedMessage_DataPoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

type ReceivedMessage_DataPoint struct {
	Rssi             *int32 `protobuf:"zigzag32,1,req,name=rssi" json:"rssi,omitempty"`
	SecondsSince     *int32 `protobuf:"varint,2,req,name=secondsSince" json:"secondsSince,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReceivedMessage_DataPoint) Reset()                    { *m = ReceivedMessage_DataPoint{} }
func (m *ReceivedMessage_DataPoint) String() string            { return proto.CompactTextString(m) }
func (*ReceivedMessage_DataPoint) ProtoMessage()               {}
func (*ReceivedMessage_DataPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *ReceivedMessage_DataPoint) GetRssi() int32 {
	if m != nil && m.Rssi != nil {
		return *m.Rssi
	}
	return 0
}

func (m *ReceivedMessage_DataPoint) GetSecondsSince() int32 {
	if m != nil && m.SecondsSince != nil {
		return *m.SecondsSince
	}
	return 0
}

func init() {
	proto.RegisterType((*ReceivedMessage)(nil), "ReceivedMessage")
	proto.RegisterType((*ReceivedMessage_Device)(nil), "ReceivedMessage.Device")
	proto.RegisterType((*ReceivedMessage_DataPoint)(nil), "ReceivedMessage.DataPoint")
}

func init() { proto.RegisterFile("ThesisProtocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x99, 0x76, 0x3a, 0x3f, 0x77, 0x46, 0x46, 0x2f, 0x23, 0x86, 0xae, 0x8a, 0xab, 0xae,
	0x2a, 0xf8, 0x0c, 0xba, 0x14, 0x8a, 0x0e, 0xa2, 0xcb, 0x98, 0x5e, 0xf5, 0x42, 0xdb, 0x0c, 0x73,
	0x63, 0x7d, 0x21, 0x1f, 0xd4, 0x24, 0x45, 0x05, 0x67, 0x77, 0xf2, 0xe5, 0xe4, 0xe3, 0x04, 0xb6,
	0xbb, 0x77, 0x12, 0x96, 0xfa, 0x60, 0x9d, 0x35, 0xb6, 0xad, 0xf6, 0x21, 0x5c, 0x7e, 0x25, 0xb0,
	0xb9, 0x27, 0x43, 0x3c, 0x50, 0x73, 0x47, 0x22, 0xfa, 0x8d, 0x70, 0x03, 0xf3, 0x86, 0x06, 0x36,
	0xf4, 0xa4, 0x26, 0x45, 0x52, 0x66, 0x7f, 0xe0, 0x59, 0x25, 0x11, 0x9c, 0xc3, 0xc9, 0x08, 0x1e,
	0x6d, 0xeb, 0xfc, 0x13, 0x95, 0x16, 0x93, 0x32, 0xc1, 0x53, 0x58, 0xb4, 0x5a, 0x5c, 0x6d, 0xc5,
	0xa9, 0xa9, 0x27, 0x0b, 0x44, 0x00, 0xf9, 0x78, 0xe9, 0xd8, 0xed, 0xb8, 0x23, 0x95, 0x79, 0x96,
	0x85, 0xd6, 0x27, 0xbf, 0x72, 0x24, 0xb3, 0x48, 0xca, 0x1f, 0xbf, 0xa8, 0x79, 0x91, 0x96, 0xab,
	0xeb, 0x8b, 0xea, 0xdf, 0xa6, 0xea, 0x26, 0xde, 0xe7, 0xb7, 0x30, 0x1b, 0x13, 0xae, 0x20, 0xed,
	0xb4, 0x89, 0x03, 0x97, 0x58, 0x01, 0x34, 0xda, 0xe9, 0xbd, 0xe5, 0xde, 0x89, 0xdf, 0x18, 0x1c,
	0xf9, 0xb1, 0xc3, 0x57, 0xea, 0x50, 0xc9, 0xaf, 0x60, 0xf9, 0x7b, 0xc0, 0x35, 0x4c, 0x0f, 0x22,
	0x1c, 0x55, 0x67, 0xb8, 0x85, 0xb5, 0x90, 0xb1, 0x7d, 0x23, 0x0f, 0xdc, 0x1b, 0x1a, 0x3f, 0xfc,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x56, 0x73, 0x0f, 0x2e, 0x3d, 0x01, 0x00, 0x00,
}
