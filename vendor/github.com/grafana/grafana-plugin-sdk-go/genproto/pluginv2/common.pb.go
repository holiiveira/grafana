// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package pluginv2

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

type TimeRange struct {
	FromRaw              string   `protobuf:"bytes,1,opt,name=fromRaw,proto3" json:"fromRaw,omitempty"`
	ToRaw                string   `protobuf:"bytes,2,opt,name=toRaw,proto3" json:"toRaw,omitempty"`
	FromEpochMs          int64    `protobuf:"varint,3,opt,name=fromEpochMs,proto3" json:"fromEpochMs,omitempty"`
	ToEpochMs            int64    `protobuf:"varint,4,opt,name=toEpochMs,proto3" json:"toEpochMs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetFromRaw() string {
	if m != nil {
		return m.FromRaw
	}
	return ""
}

func (m *TimeRange) GetToRaw() string {
	if m != nil {
		return m.ToRaw
	}
	return ""
}

func (m *TimeRange) GetFromEpochMs() int64 {
	if m != nil {
		return m.FromEpochMs
	}
	return 0
}

func (m *TimeRange) GetToEpochMs() int64 {
	if m != nil {
		return m.ToEpochMs
	}
	return 0
}

type DatasourceInfo struct {
	Id                      int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrgId                   int64             `protobuf:"varint,2,opt,name=orgId,proto3" json:"orgId,omitempty"`
	Name                    string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                    string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Url                     string            `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	JsonData                string            `protobuf:"bytes,6,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
	DecryptedSecureJsonData map[string]string `protobuf:"bytes,7,rep,name=decryptedSecureJsonData,proto3" json:"decryptedSecureJsonData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *DatasourceInfo) Reset()         { *m = DatasourceInfo{} }
func (m *DatasourceInfo) String() string { return proto.CompactTextString(m) }
func (*DatasourceInfo) ProtoMessage()    {}
func (*DatasourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *DatasourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasourceInfo.Unmarshal(m, b)
}
func (m *DatasourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasourceInfo.Marshal(b, m, deterministic)
}
func (m *DatasourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasourceInfo.Merge(m, src)
}
func (m *DatasourceInfo) XXX_Size() int {
	return xxx_messageInfo_DatasourceInfo.Size(m)
}
func (m *DatasourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DatasourceInfo proto.InternalMessageInfo

func (m *DatasourceInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DatasourceInfo) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *DatasourceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatasourceInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DatasourceInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DatasourceInfo) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *DatasourceInfo) GetDecryptedSecureJsonData() map[string]string {
	if m != nil {
		return m.DecryptedSecureJsonData
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeRange)(nil), "pluginv2.TimeRange")
	proto.RegisterType((*DatasourceInfo)(nil), "pluginv2.DatasourceInfo")
	proto.RegisterMapType((map[string]string)(nil), "pluginv2.DatasourceInfo.DecryptedSecureJsonDataEntry")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0xb6, 0xff, 0x32, 0x95, 0x22, 0x8b, 0x60, 0x28, 0x3d, 0x84, 0x9e, 0x72, 0xca,
	0xa1, 0x22, 0x88, 0xe7, 0xf6, 0xd0, 0x82, 0x97, 0xd5, 0x17, 0x58, 0x93, 0x6d, 0x8c, 0x26, 0x3b,
	0x61, 0xb3, 0xa9, 0x04, 0x9f, 0xd0, 0xb7, 0x92, 0x9d, 0x1a, 0xff, 0x80, 0xf6, 0xf6, 0x7d, 0xbf,
	0x9d, 0xd9, 0x6f, 0x98, 0x81, 0xb3, 0x14, 0xab, 0x0a, 0x75, 0x52, 0x1b, 0xb4, 0xc8, 0x27, 0x75,
	0xd9, 0xe6, 0x85, 0x3e, 0xac, 0x96, 0x6f, 0x10, 0x3c, 0x14, 0x95, 0x12, 0x52, 0xe7, 0x8a, 0x87,
	0x30, 0xde, 0x1b, 0xac, 0x84, 0x7c, 0x0d, 0xbd, 0xc8, 0x8b, 0x03, 0xd1, 0x5b, 0x7e, 0x01, 0x43,
	0x8b, 0x8e, 0xfb, 0xc4, 0x8f, 0x86, 0x47, 0x30, 0x75, 0x05, 0x9b, 0x1a, 0xd3, 0xa7, 0xbb, 0x26,
	0x64, 0x91, 0x17, 0x33, 0xf1, 0x13, 0xf1, 0x05, 0x04, 0x16, 0xfb, 0xf7, 0x01, 0xbd, 0x7f, 0x83,
	0xe5, 0xbb, 0x0f, 0xb3, 0xb5, 0xb4, 0xb2, 0xc1, 0xd6, 0xa4, 0x6a, 0xab, 0xf7, 0xc8, 0x67, 0xe0,
	0x17, 0x19, 0xa5, 0x33, 0xe1, 0x17, 0x99, 0x0b, 0x46, 0x93, 0x6f, 0x33, 0x0a, 0x66, 0xe2, 0x68,
	0x38, 0x87, 0x81, 0x96, 0x95, 0xa2, 0xc4, 0x40, 0x90, 0x76, 0xcc, 0x76, 0xb5, 0xa2, 0x94, 0x40,
	0x90, 0xe6, 0xe7, 0xc0, 0x5a, 0x53, 0x86, 0x43, 0x42, 0x4e, 0xf2, 0x39, 0x4c, 0x9e, 0x1b, 0xd4,
	0x2e, 0x35, 0x1c, 0x11, 0xfe, 0xf2, 0x1c, 0xe1, 0x32, 0x53, 0xa9, 0xe9, 0x6a, 0xab, 0xb2, 0x7b,
	0x95, 0xb6, 0x46, 0xed, 0xfa, 0xd2, 0x71, 0xc4, 0xe2, 0xe9, 0xea, 0x3a, 0xe9, 0xf7, 0x96, 0xfc,
	0x1e, 0x3b, 0x59, 0xff, 0xdd, 0xb7, 0xd1, 0xd6, 0x74, 0xe2, 0xbf, 0x5f, 0xe7, 0x3b, 0x58, 0x9c,
	0x6a, 0x74, 0xe3, 0xbf, 0xa8, 0xee, 0xf3, 0x16, 0x4e, 0xba, 0x75, 0x1c, 0x64, 0xd9, 0xaa, 0xfe,
	0x0e, 0x64, 0x6e, 0xfd, 0x1b, 0xef, 0x71, 0x44, 0x97, 0xbd, 0xfa, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x82, 0xb7, 0xa6, 0x7c, 0xe9, 0x01, 0x00, 0x00,
}
