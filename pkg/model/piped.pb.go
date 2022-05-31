// Copyright 2020 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/model/piped.proto

package model

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Piped_ConnectionStatus int32

const (
	Piped_UNKNOWN Piped_ConnectionStatus = 0
	Piped_ONLINE  Piped_ConnectionStatus = 1
	Piped_OFFLINE Piped_ConnectionStatus = 2
)

// Enum value maps for Piped_ConnectionStatus.
var (
	Piped_ConnectionStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "ONLINE",
		2: "OFFLINE",
	}
	Piped_ConnectionStatus_value = map[string]int32{
		"UNKNOWN": 0,
		"ONLINE":  1,
		"OFFLINE": 2,
	}
)

func (x Piped_ConnectionStatus) Enum() *Piped_ConnectionStatus {
	p := new(Piped_ConnectionStatus)
	*p = x
	return p
}

func (x Piped_ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Piped_ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_model_piped_proto_enumTypes[0].Descriptor()
}

func (Piped_ConnectionStatus) Type() protoreflect.EnumType {
	return &file_pkg_model_piped_proto_enumTypes[0]
}

func (x Piped_ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Piped_ConnectionStatus.Descriptor instead.
func (Piped_ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_model_piped_proto_rawDescGZIP(), []int{0, 0}
}

type Piped struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The generated unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the piped.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The additional description about the piped.
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// The hash value of the secret key generated for the piped.
	// This is used to authenticate while communicating with the control plane.
	//
	// Deprecated: Do not use.
	KeyHash string `protobuf:"bytes,4,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	// The ID of the project this piped belongs to.
	ProjectId string `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Currently running version.
	Version string `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	// Unix time when the piped is started up.
	StartedAt int64 `protobuf:"varint,8,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// List of configured cloud providers.
	CloudProviders []*Piped_CloudProvider `protobuf:"bytes,9,rep,name=cloud_providers,json=cloudProviders,proto3" json:"cloud_providers,omitempty"`
	// List of configured repositories.
	Repositories []*ApplicationGitRepository `protobuf:"bytes,10,rep,name=repositories,proto3" json:"repositories,omitempty"`
	// The latest connection status of piped.
	Status Piped_ConnectionStatus `protobuf:"varint,11,opt,name=status,proto3,enum=model.Piped_ConnectionStatus" json:"status,omitempty"`
	// The public key/service account for encrypting the secret data.
	SecretEncryption *Piped_SecretEncryption `protobuf:"bytes,21,opt,name=secret_encryption,json=secretEncryption,proto3" json:"secret_encryption,omitempty"`
	// The list keys can be used to authenticate.
	Keys []*PipedKey `protobuf:"bytes,20,rep,name=keys,proto3" json:"keys,omitempty"`
	// The desired version of Piped that should be run.
	// Launcher uses this value to compare with the currently running version
	// to determine whether Piped should be restarted with another version or not.
	DesiredVersion string `protobuf:"bytes,30,opt,name=desired_version,json=desiredVersion,proto3" json:"desired_version,omitempty"`
	// The flag to determine whether Piped should be restarted with current version or not.
	NeedRestart bool `protobuf:"varint,31,opt,name=need_restart,json=needRestart,proto3" json:"need_restart,omitempty"`
	// Whether the piped is disabled or not.
	Disabled bool `protobuf:"varint,13,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Unix time when the piped is created.
	CreatedAt int64 `protobuf:"varint,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Unix time of the last time when the piped is updated.
	UpdatedAt int64 `protobuf:"varint,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Piped) Reset() {
	*x = Piped{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_piped_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Piped) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Piped) ProtoMessage() {}

func (x *Piped) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_piped_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Piped.ProtoReflect.Descriptor instead.
func (*Piped) Descriptor() ([]byte, []int) {
	return file_pkg_model_piped_proto_rawDescGZIP(), []int{0}
}

func (x *Piped) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Piped) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Piped) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

// Deprecated: Do not use.
func (x *Piped) GetKeyHash() string {
	if x != nil {
		return x.KeyHash
	}
	return ""
}

func (x *Piped) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Piped) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Piped) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *Piped) GetCloudProviders() []*Piped_CloudProvider {
	if x != nil {
		return x.CloudProviders
	}
	return nil
}

func (x *Piped) GetRepositories() []*ApplicationGitRepository {
	if x != nil {
		return x.Repositories
	}
	return nil
}

func (x *Piped) GetStatus() Piped_ConnectionStatus {
	if x != nil {
		return x.Status
	}
	return Piped_UNKNOWN
}

func (x *Piped) GetSecretEncryption() *Piped_SecretEncryption {
	if x != nil {
		return x.SecretEncryption
	}
	return nil
}

func (x *Piped) GetKeys() []*PipedKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Piped) GetDesiredVersion() string {
	if x != nil {
		return x.DesiredVersion
	}
	return ""
}

func (x *Piped) GetNeedRestart() bool {
	if x != nil {
		return x.NeedRestart
	}
	return false
}

func (x *Piped) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Piped) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Piped) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type PipedKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash value of the key.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// The creator of the key.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Unix time when the key is created.
	CreatedAt int64 `protobuf:"varint,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PipedKey) Reset() {
	*x = PipedKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_piped_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipedKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipedKey) ProtoMessage() {}

func (x *PipedKey) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_piped_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipedKey.ProtoReflect.Descriptor instead.
func (*PipedKey) Descriptor() ([]byte, []int) {
	return file_pkg_model_piped_proto_rawDescGZIP(), []int{1}
}

func (x *PipedKey) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *PipedKey) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *PipedKey) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type Piped_CloudProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Piped_CloudProvider) Reset() {
	*x = Piped_CloudProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_piped_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Piped_CloudProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Piped_CloudProvider) ProtoMessage() {}

func (x *Piped_CloudProvider) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_piped_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Piped_CloudProvider.ProtoReflect.Descriptor instead.
func (*Piped_CloudProvider) Descriptor() ([]byte, []int) {
	return file_pkg_model_piped_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Piped_CloudProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Piped_CloudProvider) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Piped_SecretEncryption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	PublicKey             string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	EncryptServiceAccount string `protobuf:"bytes,3,opt,name=encrypt_service_account,json=encryptServiceAccount,proto3" json:"encrypt_service_account,omitempty"`
}

func (x *Piped_SecretEncryption) Reset() {
	*x = Piped_SecretEncryption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_piped_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Piped_SecretEncryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Piped_SecretEncryption) ProtoMessage() {}

func (x *Piped_SecretEncryption) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_piped_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Piped_SecretEncryption.ProtoReflect.Descriptor instead.
func (*Piped_SecretEncryption) Descriptor() ([]byte, []int) {
	return file_pkg_model_piped_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Piped_SecretEncryption) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Piped_SecretEncryption) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Piped_SecretEncryption) GetEncryptServiceAccount() string {
	if x != nil {
		return x.EncryptServiceAccount
	}
	return ""
}

var File_pkg_model_piped_proto protoreflect.FileDescriptor

var file_pkg_model_piped_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf9, 0x07, 0x0a, 0x05, 0x50, 0x69, 0x70, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x64, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4a,
	0x0a, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65,
	0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x6e, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x26, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x49, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x1a, 0xa6, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x42, 0x24, 0x72, 0x22, 0x52, 0x08, 0x4b, 0x45,
	0x59, 0x5f, 0x50, 0x41, 0x49, 0x52, 0x52, 0x07, 0x47, 0x43, 0x50, 0x5f, 0x4b, 0x4d, 0x53, 0x52,
	0x07, 0x41, 0x57, 0x53, 0x5f, 0x4b, 0x4d, 0x53, 0x52, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46, 0x4c,
	0x49, 0x4e, 0x45, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x72, 0x0a, 0x08, 0x50,
	0x69, 0x70, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x2d, 0x63, 0x64, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x63, 0x64, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_model_piped_proto_rawDescOnce sync.Once
	file_pkg_model_piped_proto_rawDescData = file_pkg_model_piped_proto_rawDesc
)

func file_pkg_model_piped_proto_rawDescGZIP() []byte {
	file_pkg_model_piped_proto_rawDescOnce.Do(func() {
		file_pkg_model_piped_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_model_piped_proto_rawDescData)
	})
	return file_pkg_model_piped_proto_rawDescData
}

var file_pkg_model_piped_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_model_piped_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_model_piped_proto_goTypes = []interface{}{
	(Piped_ConnectionStatus)(0),      // 0: model.Piped.ConnectionStatus
	(*Piped)(nil),                    // 1: model.Piped
	(*PipedKey)(nil),                 // 2: model.PipedKey
	(*Piped_CloudProvider)(nil),      // 3: model.Piped.CloudProvider
	(*Piped_SecretEncryption)(nil),   // 4: model.Piped.SecretEncryption
	(*ApplicationGitRepository)(nil), // 5: model.ApplicationGitRepository
}
var file_pkg_model_piped_proto_depIdxs = []int32{
	3, // 0: model.Piped.cloud_providers:type_name -> model.Piped.CloudProvider
	5, // 1: model.Piped.repositories:type_name -> model.ApplicationGitRepository
	0, // 2: model.Piped.status:type_name -> model.Piped.ConnectionStatus
	4, // 3: model.Piped.secret_encryption:type_name -> model.Piped.SecretEncryption
	2, // 4: model.Piped.keys:type_name -> model.PipedKey
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_model_piped_proto_init() }
func file_pkg_model_piped_proto_init() {
	if File_pkg_model_piped_proto != nil {
		return
	}
	file_pkg_model_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_model_piped_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Piped); i {
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
		file_pkg_model_piped_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipedKey); i {
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
		file_pkg_model_piped_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Piped_CloudProvider); i {
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
		file_pkg_model_piped_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Piped_SecretEncryption); i {
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
			RawDescriptor: file_pkg_model_piped_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_model_piped_proto_goTypes,
		DependencyIndexes: file_pkg_model_piped_proto_depIdxs,
		EnumInfos:         file_pkg_model_piped_proto_enumTypes,
		MessageInfos:      file_pkg_model_piped_proto_msgTypes,
	}.Build()
	File_pkg_model_piped_proto = out.File
	file_pkg_model_piped_proto_rawDesc = nil
	file_pkg_model_piped_proto_goTypes = nil
	file_pkg_model_piped_proto_depIdxs = nil
}
