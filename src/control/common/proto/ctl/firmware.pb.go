//
// (C) Copyright 2020-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: ctl/firmware.proto

package ctl

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

type FirmwareUpdateReq_DeviceType int32

const (
	FirmwareUpdateReq_SCM  FirmwareUpdateReq_DeviceType = 0
	FirmwareUpdateReq_NVMe FirmwareUpdateReq_DeviceType = 1
)

// Enum value maps for FirmwareUpdateReq_DeviceType.
var (
	FirmwareUpdateReq_DeviceType_name = map[int32]string{
		0: "SCM",
		1: "NVMe",
	}
	FirmwareUpdateReq_DeviceType_value = map[string]int32{
		"SCM":  0,
		"NVMe": 1,
	}
)

func (x FirmwareUpdateReq_DeviceType) Enum() *FirmwareUpdateReq_DeviceType {
	p := new(FirmwareUpdateReq_DeviceType)
	*p = x
	return p
}

func (x FirmwareUpdateReq_DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FirmwareUpdateReq_DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_ctl_firmware_proto_enumTypes[0].Descriptor()
}

func (FirmwareUpdateReq_DeviceType) Type() protoreflect.EnumType {
	return &file_ctl_firmware_proto_enumTypes[0]
}

func (x FirmwareUpdateReq_DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FirmwareUpdateReq_DeviceType.Descriptor instead.
func (FirmwareUpdateReq_DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{4, 0}
}

type FirmwareQueryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryScm    bool     `protobuf:"varint,1,opt,name=queryScm,proto3" json:"queryScm,omitempty"`      // Should we query SCM devices?
	QueryNvme   bool     `protobuf:"varint,2,opt,name=queryNvme,proto3" json:"queryNvme,omitempty"`    // Should we query NVMe devices?
	DeviceIDs   []string `protobuf:"bytes,3,rep,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`     // Filter by specific devices
	ModelID     string   `protobuf:"bytes,4,opt,name=modelID,proto3" json:"modelID,omitempty"`         // Filter by model ID
	FirmwareRev string   `protobuf:"bytes,5,opt,name=firmwareRev,proto3" json:"firmwareRev,omitempty"` // Filter by current firmware revision
}

func (x *FirmwareQueryReq) Reset() {
	*x = FirmwareQueryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_firmware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareQueryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareQueryReq) ProtoMessage() {}

func (x *FirmwareQueryReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_firmware_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareQueryReq.ProtoReflect.Descriptor instead.
func (*FirmwareQueryReq) Descriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{0}
}

func (x *FirmwareQueryReq) GetQueryScm() bool {
	if x != nil {
		return x.QueryScm
	}
	return false
}

func (x *FirmwareQueryReq) GetQueryNvme() bool {
	if x != nil {
		return x.QueryNvme
	}
	return false
}

func (x *FirmwareQueryReq) GetDeviceIDs() []string {
	if x != nil {
		return x.DeviceIDs
	}
	return nil
}

func (x *FirmwareQueryReq) GetModelID() string {
	if x != nil {
		return x.ModelID
	}
	return ""
}

func (x *FirmwareQueryReq) GetFirmwareRev() string {
	if x != nil {
		return x.FirmwareRev
	}
	return ""
}

type ScmFirmwareQueryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module            *ScmModule `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`                        // The module of this firmware
	ActiveVersion     string     `protobuf:"bytes,2,opt,name=activeVersion,proto3" json:"activeVersion,omitempty"`          // Active FW version
	StagedVersion     string     `protobuf:"bytes,3,opt,name=stagedVersion,proto3" json:"stagedVersion,omitempty"`          // Staged FW version
	ImageMaxSizeBytes uint32     `protobuf:"varint,4,opt,name=imageMaxSizeBytes,proto3" json:"imageMaxSizeBytes,omitempty"` // Maximum size of FW image accepted
	UpdateStatus      uint32     `protobuf:"varint,5,opt,name=updateStatus,proto3" json:"updateStatus,omitempty"`           // Status of FW update
	Error             string     `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`                          // Error string, if any
}

func (x *ScmFirmwareQueryResp) Reset() {
	*x = ScmFirmwareQueryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_firmware_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScmFirmwareQueryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScmFirmwareQueryResp) ProtoMessage() {}

func (x *ScmFirmwareQueryResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_firmware_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScmFirmwareQueryResp.ProtoReflect.Descriptor instead.
func (*ScmFirmwareQueryResp) Descriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{1}
}

func (x *ScmFirmwareQueryResp) GetModule() *ScmModule {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *ScmFirmwareQueryResp) GetActiveVersion() string {
	if x != nil {
		return x.ActiveVersion
	}
	return ""
}

func (x *ScmFirmwareQueryResp) GetStagedVersion() string {
	if x != nil {
		return x.StagedVersion
	}
	return ""
}

func (x *ScmFirmwareQueryResp) GetImageMaxSizeBytes() uint32 {
	if x != nil {
		return x.ImageMaxSizeBytes
	}
	return 0
}

func (x *ScmFirmwareQueryResp) GetUpdateStatus() uint32 {
	if x != nil {
		return x.UpdateStatus
	}
	return 0
}

func (x *ScmFirmwareQueryResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type NvmeFirmwareQueryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *NvmeController `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"` // Controller information includes FW rev
}

func (x *NvmeFirmwareQueryResp) Reset() {
	*x = NvmeFirmwareQueryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_firmware_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NvmeFirmwareQueryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NvmeFirmwareQueryResp) ProtoMessage() {}

func (x *NvmeFirmwareQueryResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_firmware_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NvmeFirmwareQueryResp.ProtoReflect.Descriptor instead.
func (*NvmeFirmwareQueryResp) Descriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{2}
}

func (x *NvmeFirmwareQueryResp) GetDevice() *NvmeController {
	if x != nil {
		return x.Device
	}
	return nil
}

type FirmwareQueryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScmResults  []*ScmFirmwareQueryResp  `protobuf:"bytes,1,rep,name=scmResults,proto3" json:"scmResults,omitempty"`
	NvmeResults []*NvmeFirmwareQueryResp `protobuf:"bytes,2,rep,name=nvmeResults,proto3" json:"nvmeResults,omitempty"`
}

func (x *FirmwareQueryResp) Reset() {
	*x = FirmwareQueryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_firmware_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareQueryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareQueryResp) ProtoMessage() {}

func (x *FirmwareQueryResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_firmware_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareQueryResp.ProtoReflect.Descriptor instead.
func (*FirmwareQueryResp) Descriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{3}
}

func (x *FirmwareQueryResp) GetScmResults() []*ScmFirmwareQueryResp {
	if x != nil {
		return x.ScmResults
	}
	return nil
}

func (x *FirmwareQueryResp) GetNvmeResults() []*NvmeFirmwareQueryResp {
	if x != nil {
		return x.NvmeResults
	}
	return nil
}

type FirmwareUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirmwarePath string                       `protobuf:"bytes,1,opt,name=firmwarePath,proto3" json:"firmwarePath,omitempty"`                        // Path to firmware file
	Type         FirmwareUpdateReq_DeviceType `protobuf:"varint,2,opt,name=type,proto3,enum=ctl.FirmwareUpdateReq_DeviceType" json:"type,omitempty"` // Type of device this firmware applies to
	DeviceIDs    []string                     `protobuf:"bytes,3,rep,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`                              // Devices this update applies to
	ModelID      string                       `protobuf:"bytes,4,opt,name=modelID,proto3" json:"modelID,omitempty"`                                  // Model ID this update applies to
	FirmwareRev  string                       `protobuf:"bytes,5,opt,name=firmwareRev,proto3" json:"firmwareRev,omitempty"`                          // Starting FW rev this update applies to
}

func (x *FirmwareUpdateReq) Reset() {
	*x = FirmwareUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_firmware_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareUpdateReq) ProtoMessage() {}

func (x *FirmwareUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_firmware_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareUpdateReq.ProtoReflect.Descriptor instead.
func (*FirmwareUpdateReq) Descriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{4}
}

func (x *FirmwareUpdateReq) GetFirmwarePath() string {
	if x != nil {
		return x.FirmwarePath
	}
	return ""
}

func (x *FirmwareUpdateReq) GetType() FirmwareUpdateReq_DeviceType {
	if x != nil {
		return x.Type
	}
	return FirmwareUpdateReq_SCM
}

func (x *FirmwareUpdateReq) GetDeviceIDs() []string {
	if x != nil {
		return x.DeviceIDs
	}
	return nil
}

func (x *FirmwareUpdateReq) GetModelID() string {
	if x != nil {
		return x.ModelID
	}
	return ""
}

func (x *FirmwareUpdateReq) GetFirmwareRev() string {
	if x != nil {
		return x.FirmwareRev
	}
	return ""
}

type ScmFirmwareUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module *ScmModule `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"` // SCM device
	Error  string     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`   // empty if successful
}

func (x *ScmFirmwareUpdateResp) Reset() {
	*x = ScmFirmwareUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_firmware_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScmFirmwareUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScmFirmwareUpdateResp) ProtoMessage() {}

func (x *ScmFirmwareUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_firmware_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScmFirmwareUpdateResp.ProtoReflect.Descriptor instead.
func (*ScmFirmwareUpdateResp) Descriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{5}
}

func (x *ScmFirmwareUpdateResp) GetModule() *ScmModule {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *ScmFirmwareUpdateResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type NvmeFirmwareUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PciAddr string `protobuf:"bytes,1,opt,name=pciAddr,proto3" json:"pciAddr,omitempty"` // PCI address of the NVMe device
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`     // empty if successful
}

func (x *NvmeFirmwareUpdateResp) Reset() {
	*x = NvmeFirmwareUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_firmware_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NvmeFirmwareUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NvmeFirmwareUpdateResp) ProtoMessage() {}

func (x *NvmeFirmwareUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_firmware_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NvmeFirmwareUpdateResp.ProtoReflect.Descriptor instead.
func (*NvmeFirmwareUpdateResp) Descriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{6}
}

func (x *NvmeFirmwareUpdateResp) GetPciAddr() string {
	if x != nil {
		return x.PciAddr
	}
	return ""
}

func (x *NvmeFirmwareUpdateResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FirmwareUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScmResults  []*ScmFirmwareUpdateResp  `protobuf:"bytes,1,rep,name=scmResults,proto3" json:"scmResults,omitempty"`   // results for SCM update
	NvmeResults []*NvmeFirmwareUpdateResp `protobuf:"bytes,2,rep,name=nvmeResults,proto3" json:"nvmeResults,omitempty"` // results for NVMe update
}

func (x *FirmwareUpdateResp) Reset() {
	*x = FirmwareUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_firmware_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareUpdateResp) ProtoMessage() {}

func (x *FirmwareUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_firmware_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareUpdateResp.ProtoReflect.Descriptor instead.
func (*FirmwareUpdateResp) Descriptor() ([]byte, []int) {
	return file_ctl_firmware_proto_rawDescGZIP(), []int{7}
}

func (x *FirmwareUpdateResp) GetScmResults() []*ScmFirmwareUpdateResp {
	if x != nil {
		return x.ScmResults
	}
	return nil
}

func (x *FirmwareUpdateResp) GetNvmeResults() []*NvmeFirmwareUpdateResp {
	if x != nil {
		return x.NvmeResults
	}
	return nil
}

var File_ctl_firmware_proto protoreflect.FileDescriptor

var file_ctl_firmware_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x74, 0x6c, 0x2f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x74, 0x6c, 0x1a, 0x15, 0x63, 0x74, 0x6c, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x63, 0x74, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x76,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x72,
	0x6d, 0x77, 0x61, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x4e, 0x76, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x4e, 0x76, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x76, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65,
	0x76, 0x22, 0xf2, 0x01, 0x0a, 0x14, 0x53, 0x63, 0x6d, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x74, 0x6c,
	0x2e, 0x53, 0x63, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x74, 0x61, 0x67, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c,
	0x0a, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x4d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x44, 0x0a, 0x15, 0x4e, 0x76, 0x6d, 0x65, 0x46, 0x69,
	0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2b, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x4e, 0x76, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x8c, 0x01, 0x0a,
	0x11, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x63, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x63, 0x6d,
	0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x0a, 0x73, 0x63, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x3c, 0x0a,
	0x0b, 0x6e, 0x76, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x4e, 0x76, 0x6d, 0x65, 0x46, 0x69, 0x72,
	0x6d, 0x77, 0x61, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0b,
	0x6e, 0x76, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xe9, 0x01, 0x0a, 0x11,
	0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65,
	0x52, 0x65, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x6d, 0x77,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x76, 0x22, 0x1f, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x43, 0x4d, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x56, 0x4d, 0x65, 0x10, 0x01, 0x22, 0x55, 0x0a, 0x15, 0x53, 0x63, 0x6d, 0x46, 0x69,
	0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x63, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x48,
	0x0a, 0x16, 0x4e, 0x76, 0x6d, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x63, 0x69, 0x41,
	0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x63, 0x69, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x8f, 0x01, 0x0a, 0x12, 0x46, 0x69, 0x72,
	0x6d, 0x77, 0x61, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3a, 0x0a, 0x0a, 0x73, 0x63, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x63, 0x6d, 0x46, 0x69, 0x72,
	0x6d, 0x77, 0x61, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x0a, 0x73, 0x63, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x6e,
	0x76, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x4e, 0x76, 0x6d, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77,
	0x61, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0b, 0x6e,
	0x76, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6f, 0x73, 0x2d, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x64, 0x61, 0x6f, 0x73, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x74, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ctl_firmware_proto_rawDescOnce sync.Once
	file_ctl_firmware_proto_rawDescData = file_ctl_firmware_proto_rawDesc
)

func file_ctl_firmware_proto_rawDescGZIP() []byte {
	file_ctl_firmware_proto_rawDescOnce.Do(func() {
		file_ctl_firmware_proto_rawDescData = protoimpl.X.CompressGZIP(file_ctl_firmware_proto_rawDescData)
	})
	return file_ctl_firmware_proto_rawDescData
}

var file_ctl_firmware_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ctl_firmware_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ctl_firmware_proto_goTypes = []interface{}{
	(FirmwareUpdateReq_DeviceType)(0), // 0: ctl.FirmwareUpdateReq.DeviceType
	(*FirmwareQueryReq)(nil),          // 1: ctl.FirmwareQueryReq
	(*ScmFirmwareQueryResp)(nil),      // 2: ctl.ScmFirmwareQueryResp
	(*NvmeFirmwareQueryResp)(nil),     // 3: ctl.NvmeFirmwareQueryResp
	(*FirmwareQueryResp)(nil),         // 4: ctl.FirmwareQueryResp
	(*FirmwareUpdateReq)(nil),         // 5: ctl.FirmwareUpdateReq
	(*ScmFirmwareUpdateResp)(nil),     // 6: ctl.ScmFirmwareUpdateResp
	(*NvmeFirmwareUpdateResp)(nil),    // 7: ctl.NvmeFirmwareUpdateResp
	(*FirmwareUpdateResp)(nil),        // 8: ctl.FirmwareUpdateResp
	(*ScmModule)(nil),                 // 9: ctl.ScmModule
	(*NvmeController)(nil),            // 10: ctl.NvmeController
}
var file_ctl_firmware_proto_depIdxs = []int32{
	9,  // 0: ctl.ScmFirmwareQueryResp.module:type_name -> ctl.ScmModule
	10, // 1: ctl.NvmeFirmwareQueryResp.device:type_name -> ctl.NvmeController
	2,  // 2: ctl.FirmwareQueryResp.scmResults:type_name -> ctl.ScmFirmwareQueryResp
	3,  // 3: ctl.FirmwareQueryResp.nvmeResults:type_name -> ctl.NvmeFirmwareQueryResp
	0,  // 4: ctl.FirmwareUpdateReq.type:type_name -> ctl.FirmwareUpdateReq.DeviceType
	9,  // 5: ctl.ScmFirmwareUpdateResp.module:type_name -> ctl.ScmModule
	6,  // 6: ctl.FirmwareUpdateResp.scmResults:type_name -> ctl.ScmFirmwareUpdateResp
	7,  // 7: ctl.FirmwareUpdateResp.nvmeResults:type_name -> ctl.NvmeFirmwareUpdateResp
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_ctl_firmware_proto_init() }
func file_ctl_firmware_proto_init() {
	if File_ctl_firmware_proto != nil {
		return
	}
	file_ctl_storage_scm_proto_init()
	file_ctl_storage_nvme_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ctl_firmware_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareQueryReq); i {
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
		file_ctl_firmware_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScmFirmwareQueryResp); i {
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
		file_ctl_firmware_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NvmeFirmwareQueryResp); i {
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
		file_ctl_firmware_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareQueryResp); i {
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
		file_ctl_firmware_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareUpdateReq); i {
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
		file_ctl_firmware_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScmFirmwareUpdateResp); i {
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
		file_ctl_firmware_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NvmeFirmwareUpdateResp); i {
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
		file_ctl_firmware_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareUpdateResp); i {
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
			RawDescriptor: file_ctl_firmware_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ctl_firmware_proto_goTypes,
		DependencyIndexes: file_ctl_firmware_proto_depIdxs,
		EnumInfos:         file_ctl_firmware_proto_enumTypes,
		MessageInfos:      file_ctl_firmware_proto_msgTypes,
	}.Build()
	File_ctl_firmware_proto = out.File
	file_ctl_firmware_proto_rawDesc = nil
	file_ctl_firmware_proto_goTypes = nil
	file_ctl_firmware_proto_depIdxs = nil
}
