// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common

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

// * Represents an empty message
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_450217a1e26bcceb, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// * A type which contains attestation data for specific platform.
type AttestationData struct {
	// * Type of attestation to perform.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// * The attestation data.
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestationData) Reset()         { *m = AttestationData{} }
func (m *AttestationData) String() string { return proto.CompactTextString(m) }
func (*AttestationData) ProtoMessage()    {}
func (*AttestationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_450217a1e26bcceb, []int{1}
}
func (m *AttestationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationData.Unmarshal(m, b)
}
func (m *AttestationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationData.Marshal(b, m, deterministic)
}
func (dst *AttestationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationData.Merge(dst, src)
}
func (m *AttestationData) XXX_Size() int {
	return xxx_messageInfo_AttestationData.Size(m)
}
func (m *AttestationData) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationData.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationData proto.InternalMessageInfo

func (m *AttestationData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AttestationData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// * A type which describes the conditions under which a registration
// entry is matched.
type Selector struct {
	// * A selector type represents the type of attestation used in attesting
	// the entity (Eg: AWS, K8).
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// * The value to be attested.
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_450217a1e26bcceb, []int{2}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (dst *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(dst, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Selector) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// * Represents a type with a list of Selector.
type Selectors struct {
	// * A list of Selector.
	Entries              []*Selector `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Selectors) Reset()         { *m = Selectors{} }
func (m *Selectors) String() string { return proto.CompactTextString(m) }
func (*Selectors) ProtoMessage()    {}
func (*Selectors) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_450217a1e26bcceb, []int{3}
}
func (m *Selectors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selectors.Unmarshal(m, b)
}
func (m *Selectors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selectors.Marshal(b, m, deterministic)
}
func (dst *Selectors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selectors.Merge(dst, src)
}
func (m *Selectors) XXX_Size() int {
	return xxx_messageInfo_Selectors.Size(m)
}
func (m *Selectors) XXX_DiscardUnknown() {
	xxx_messageInfo_Selectors.DiscardUnknown(m)
}

var xxx_messageInfo_Selectors proto.InternalMessageInfo

func (m *Selectors) GetEntries() []*Selector {
	if m != nil {
		return m.Entries
	}
	return nil
}

// * This is a curated record that the Server uses to set up and
// manage the various registered nodes and workloads that are controlled by it.
type RegistrationEntry struct {
	// * A list of selectors.
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors" json:"selectors,omitempty"`
	// * The SPIFFE ID of an entity that is authorized to attest the validity
	// of a selector
	ParentId string `protobuf:"bytes,2,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	// * The SPIFFE ID is a structured string used to identify a resource or
	// caller. It is defined as a URI comprising a “trust domain” and an
	// associated path.
	SpiffeId string `protobuf:"bytes,3,opt,name=spiffe_id,json=spiffeId" json:"spiffe_id,omitempty"`
	// * Time to live.
	Ttl int32 `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
	// * A list of federated bundle spiffe ids.
	FbSpiffeIds []string `protobuf:"bytes,5,rep,name=fb_spiffe_ids,json=fbSpiffeIds" json:"fb_spiffe_ids,omitempty"`
	// * Entry ID
	EntryId              string   `protobuf:"bytes,6,opt,name=entry_id,json=entryId" json:"entry_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationEntry) Reset()         { *m = RegistrationEntry{} }
func (m *RegistrationEntry) String() string { return proto.CompactTextString(m) }
func (*RegistrationEntry) ProtoMessage()    {}
func (*RegistrationEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_450217a1e26bcceb, []int{4}
}
func (m *RegistrationEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationEntry.Unmarshal(m, b)
}
func (m *RegistrationEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationEntry.Marshal(b, m, deterministic)
}
func (dst *RegistrationEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationEntry.Merge(dst, src)
}
func (m *RegistrationEntry) XXX_Size() int {
	return xxx_messageInfo_RegistrationEntry.Size(m)
}
func (m *RegistrationEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationEntry proto.InternalMessageInfo

func (m *RegistrationEntry) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *RegistrationEntry) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *RegistrationEntry) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *RegistrationEntry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *RegistrationEntry) GetFbSpiffeIds() []string {
	if m != nil {
		return m.FbSpiffeIds
	}
	return nil
}

func (m *RegistrationEntry) GetEntryId() string {
	if m != nil {
		return m.EntryId
	}
	return ""
}

// * A list of registration entries.
type RegistrationEntries struct {
	// * A list of RegistrationEntry.
	Entries              []*RegistrationEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegistrationEntries) Reset()         { *m = RegistrationEntries{} }
func (m *RegistrationEntries) String() string { return proto.CompactTextString(m) }
func (*RegistrationEntries) ProtoMessage()    {}
func (*RegistrationEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_450217a1e26bcceb, []int{5}
}
func (m *RegistrationEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationEntries.Unmarshal(m, b)
}
func (m *RegistrationEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationEntries.Marshal(b, m, deterministic)
}
func (dst *RegistrationEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationEntries.Merge(dst, src)
}
func (m *RegistrationEntries) XXX_Size() int {
	return xxx_messageInfo_RegistrationEntries.Size(m)
}
func (m *RegistrationEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationEntries.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationEntries proto.InternalMessageInfo

func (m *RegistrationEntries) GetEntries() []*RegistrationEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "spire.common.Empty")
	proto.RegisterType((*AttestationData)(nil), "spire.common.AttestationData")
	proto.RegisterType((*Selector)(nil), "spire.common.Selector")
	proto.RegisterType((*Selectors)(nil), "spire.common.Selectors")
	proto.RegisterType((*RegistrationEntry)(nil), "spire.common.RegistrationEntry")
	proto.RegisterType((*RegistrationEntries)(nil), "spire.common.RegistrationEntries")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_450217a1e26bcceb) }

var fileDescriptor_common_450217a1e26bcceb = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x25, 0xbf, 0x34, 0x6d, 0x32, 0xed, 0x0f, 0x75, 0x15, 0x89, 0x78, 0x30, 0xec, 0x29, 0xa7,
	0x20, 0xda, 0x4b, 0x0f, 0x1e, 0x14, 0x7b, 0xe8, 0x4d, 0xb6, 0x37, 0x2f, 0x65, 0xdb, 0x4c, 0x64,
	0xa1, 0xf9, 0xc3, 0xee, 0x28, 0xe4, 0x7b, 0xfa, 0x81, 0x24, 0xbb, 0x4d, 0xd5, 0x2a, 0x78, 0x9b,
	0x7d, 0xf3, 0xde, 0x63, 0xde, 0x63, 0x61, 0xb2, 0xa9, 0xcb, 0xb2, 0xae, 0xb2, 0x46, 0xd7, 0x54,
	0xb3, 0x89, 0x69, 0x94, 0xc6, 0xcc, 0x61, 0x7c, 0x04, 0xc1, 0xbc, 0x6c, 0xa8, 0xe5, 0x33, 0x38,
	0xba, 0x27, 0x42, 0x43, 0x92, 0x54, 0x5d, 0x3d, 0x4a, 0x92, 0x8c, 0xc1, 0x80, 0xda, 0x06, 0x63,
	0x2f, 0xf1, 0xd2, 0x48, 0xd8, 0xb9, 0xc3, 0x72, 0x49, 0x32, 0xfe, 0x97, 0x78, 0xe9, 0x44, 0xd8,
	0x99, 0x4f, 0x21, 0x5c, 0xe2, 0x16, 0x37, 0x54, 0xeb, 0x5f, 0x35, 0x67, 0x10, 0xbc, 0xc9, 0xed,
	0x2b, 0x5a, 0x51, 0x24, 0xdc, 0x83, 0xdf, 0x41, 0xd4, 0xab, 0x0c, 0xbb, 0x86, 0x11, 0x56, 0xa4,
	0x15, 0x9a, 0xd8, 0x4b, 0xfc, 0x74, 0x7c, 0x73, 0x9e, 0x7d, 0x3d, 0x33, 0xeb, 0x99, 0xa2, 0xa7,
	0xf1, 0x77, 0x0f, 0x4e, 0x04, 0xbe, 0x28, 0x43, 0xda, 0x5e, 0x3c, 0xaf, 0x48, 0xb7, 0x6c, 0x0a,
	0x91, 0xe9, 0x4d, 0xff, 0x70, 0xfa, 0x24, 0xb2, 0x4b, 0x88, 0x1a, 0xa9, 0xb1, 0xa2, 0x95, 0xca,
	0x77, 0x47, 0x86, 0x0e, 0x58, 0xe4, 0xdd, 0xd2, 0x34, 0xaa, 0x28, 0xb0, 0x5b, 0xfa, 0x6e, 0xe9,
	0x80, 0x45, 0xce, 0x8e, 0xc1, 0x27, 0xda, 0xc6, 0x83, 0xc4, 0x4b, 0x03, 0xd1, 0x8d, 0x8c, 0xc3,
	0xff, 0x62, 0xbd, 0xda, 0x2b, 0x4c, 0x1c, 0x24, 0x7e, 0x1a, 0x89, 0x71, 0xb1, 0x5e, 0xee, 0x44,
	0x86, 0x5d, 0x40, 0xd8, 0xc5, 0x68, 0x3b, 0xc7, 0xa1, 0x75, 0xb4, 0xb1, 0xda, 0x45, 0xce, 0x9f,
	0xe0, 0xf4, 0x30, 0x95, 0x42, 0xc3, 0x66, 0x87, 0xfd, 0x5c, 0x7d, 0x4f, 0xf5, 0xa3, 0x89, 0x7d,
	0x51, 0x0f, 0xe1, 0xf3, 0xd0, 0x91, 0xd6, 0x43, 0xfb, 0x01, 0x6e, 0x3f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xb9, 0xa4, 0x13, 0xc7, 0x10, 0x02, 0x00, 0x00,
}
