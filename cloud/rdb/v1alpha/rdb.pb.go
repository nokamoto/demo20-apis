// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cloud/rdb/v1alpha/rdb.proto

package v1alpha

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/nokamoto/demo20-apis/cloud/api"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string   `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Cluster   *Cluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_cloud_rdb_v1alpha_rdb_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClusterRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateClusterRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type DeleteClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteClusterRequest) Reset() {
	*x = DeleteClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterRequest) ProtoMessage() {}

func (x *DeleteClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return file_cloud_rdb_v1alpha_rdb_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetClusterRequest) Reset() {
	*x = GetClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRequest) ProtoMessage() {}

func (x *GetClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRequest.ProtoReflect.Descriptor instead.
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return file_cloud_rdb_v1alpha_rdb_proto_rawDescGZIP(), []int{2}
}

func (x *GetClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListClusterRequest) Reset() {
	*x = ListClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClusterRequest) ProtoMessage() {}

func (x *ListClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClusterRequest.ProtoReflect.Descriptor instead.
func (*ListClusterRequest) Descriptor() ([]byte, []int) {
	return file_cloud_rdb_v1alpha_rdb_proto_rawDescGZIP(), []int{3}
}

func (x *ListClusterRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListClusterRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextPageToken string     `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	Clusters      []*Cluster `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *ListClusterResponse) Reset() {
	*x = ListClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClusterResponse) ProtoMessage() {}

func (x *ListClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClusterResponse.ProtoReflect.Descriptor instead.
func (*ListClusterResponse) Descriptor() ([]byte, []int) {
	return file_cloud_rdb_v1alpha_rdb_proto_rawDescGZIP(), []int{4}
}

func (x *ListClusterResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListClusterResponse) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type UpdateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster    *Cluster              `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateClusterRequest) Reset() {
	*x = UpdateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClusterRequest) ProtoMessage() {}

func (x *UpdateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_rdb_v1alpha_rdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClusterRequest.ProtoReflect.Descriptor instead.
func (*UpdateClusterRequest) Descriptor() ([]byte, []int) {
	return file_cloud_rdb_v1alpha_rdb_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateClusterRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *UpdateClusterRequest) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_cloud_rdb_v1alpha_rdb_proto protoreflect.FileDescriptor

var file_cloud_rdb_v1alpha_rdb_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x72, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x64, 0x62,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34,
	0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x75, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0xac,
	0x04, 0x0a, 0x03, 0x52, 0x64, 0x62, 0x12, 0x6d, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x17, 0x82, 0xb5,
	0x18, 0x13, 0x0a, 0x11, 0x72, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x69, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x82, 0xb5, 0x18, 0x13, 0x0a, 0x11, 0x72,
	0x64, 0x62, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x66, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x16, 0x82, 0xb5, 0x18, 0x12, 0x0a, 0x10, 0x72, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x12, 0x74, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xb5, 0x18, 0x12, 0x0a, 0x10, 0x72, 0x64,
	0x62, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x12, 0x6d,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x17, 0x82, 0xb5, 0x18, 0x13, 0x0a, 0x11, 0x72, 0x64, 0x62, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6b, 0x61,
	0x6d, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x32, 0x30, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_rdb_v1alpha_rdb_proto_rawDescOnce sync.Once
	file_cloud_rdb_v1alpha_rdb_proto_rawDescData = file_cloud_rdb_v1alpha_rdb_proto_rawDesc
)

func file_cloud_rdb_v1alpha_rdb_proto_rawDescGZIP() []byte {
	file_cloud_rdb_v1alpha_rdb_proto_rawDescOnce.Do(func() {
		file_cloud_rdb_v1alpha_rdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_rdb_v1alpha_rdb_proto_rawDescData)
	})
	return file_cloud_rdb_v1alpha_rdb_proto_rawDescData
}

var file_cloud_rdb_v1alpha_rdb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloud_rdb_v1alpha_rdb_proto_goTypes = []interface{}{
	(*CreateClusterRequest)(nil), // 0: cloud.rdb.v1alpha.CreateClusterRequest
	(*DeleteClusterRequest)(nil), // 1: cloud.rdb.v1alpha.DeleteClusterRequest
	(*GetClusterRequest)(nil),    // 2: cloud.rdb.v1alpha.GetClusterRequest
	(*ListClusterRequest)(nil),   // 3: cloud.rdb.v1alpha.ListClusterRequest
	(*ListClusterResponse)(nil),  // 4: cloud.rdb.v1alpha.ListClusterResponse
	(*UpdateClusterRequest)(nil), // 5: cloud.rdb.v1alpha.UpdateClusterRequest
	(*Cluster)(nil),              // 6: cloud.rdb.v1alpha.Cluster
	(*field_mask.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*empty.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_cloud_rdb_v1alpha_rdb_proto_depIdxs = []int32{
	6, // 0: cloud.rdb.v1alpha.CreateClusterRequest.cluster:type_name -> cloud.rdb.v1alpha.Cluster
	6, // 1: cloud.rdb.v1alpha.ListClusterResponse.clusters:type_name -> cloud.rdb.v1alpha.Cluster
	6, // 2: cloud.rdb.v1alpha.UpdateClusterRequest.cluster:type_name -> cloud.rdb.v1alpha.Cluster
	7, // 3: cloud.rdb.v1alpha.UpdateClusterRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 4: cloud.rdb.v1alpha.Rdb.CreateCluster:input_type -> cloud.rdb.v1alpha.CreateClusterRequest
	1, // 5: cloud.rdb.v1alpha.Rdb.DeleteCluster:input_type -> cloud.rdb.v1alpha.DeleteClusterRequest
	2, // 6: cloud.rdb.v1alpha.Rdb.GetCluster:input_type -> cloud.rdb.v1alpha.GetClusterRequest
	3, // 7: cloud.rdb.v1alpha.Rdb.ListCluster:input_type -> cloud.rdb.v1alpha.ListClusterRequest
	5, // 8: cloud.rdb.v1alpha.Rdb.UpdateCluster:input_type -> cloud.rdb.v1alpha.UpdateClusterRequest
	6, // 9: cloud.rdb.v1alpha.Rdb.CreateCluster:output_type -> cloud.rdb.v1alpha.Cluster
	8, // 10: cloud.rdb.v1alpha.Rdb.DeleteCluster:output_type -> google.protobuf.Empty
	6, // 11: cloud.rdb.v1alpha.Rdb.GetCluster:output_type -> cloud.rdb.v1alpha.Cluster
	4, // 12: cloud.rdb.v1alpha.Rdb.ListCluster:output_type -> cloud.rdb.v1alpha.ListClusterResponse
	6, // 13: cloud.rdb.v1alpha.Rdb.UpdateCluster:output_type -> cloud.rdb.v1alpha.Cluster
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cloud_rdb_v1alpha_rdb_proto_init() }
func file_cloud_rdb_v1alpha_rdb_proto_init() {
	if File_cloud_rdb_v1alpha_rdb_proto != nil {
		return
	}
	file_cloud_rdb_v1alpha_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_rdb_v1alpha_rdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterRequest); i {
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
		file_cloud_rdb_v1alpha_rdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterRequest); i {
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
		file_cloud_rdb_v1alpha_rdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterRequest); i {
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
		file_cloud_rdb_v1alpha_rdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClusterRequest); i {
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
		file_cloud_rdb_v1alpha_rdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClusterResponse); i {
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
		file_cloud_rdb_v1alpha_rdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClusterRequest); i {
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
			RawDescriptor: file_cloud_rdb_v1alpha_rdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_rdb_v1alpha_rdb_proto_goTypes,
		DependencyIndexes: file_cloud_rdb_v1alpha_rdb_proto_depIdxs,
		MessageInfos:      file_cloud_rdb_v1alpha_rdb_proto_msgTypes,
	}.Build()
	File_cloud_rdb_v1alpha_rdb_proto = out.File
	file_cloud_rdb_v1alpha_rdb_proto_rawDesc = nil
	file_cloud_rdb_v1alpha_rdb_proto_goTypes = nil
	file_cloud_rdb_v1alpha_rdb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RdbClient is the client API for Rdb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RdbClient interface {
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	ListCluster(ctx context.Context, in *ListClusterRequest, opts ...grpc.CallOption) (*ListClusterResponse, error)
	UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
}

type rdbClient struct {
	cc grpc.ClientConnInterface
}

func NewRdbClient(cc grpc.ClientConnInterface) RdbClient {
	return &rdbClient{cc}
}

func (c *rdbClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/cloud.rdb.v1alpha.Rdb/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cloud.rdb.v1alpha.Rdb/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/cloud.rdb.v1alpha.Rdb/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbClient) ListCluster(ctx context.Context, in *ListClusterRequest, opts ...grpc.CallOption) (*ListClusterResponse, error) {
	out := new(ListClusterResponse)
	err := c.cc.Invoke(ctx, "/cloud.rdb.v1alpha.Rdb/ListCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbClient) UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/cloud.rdb.v1alpha.Rdb/UpdateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RdbServer is the server API for Rdb service.
type RdbServer interface {
	CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*empty.Empty, error)
	GetCluster(context.Context, *GetClusterRequest) (*Cluster, error)
	ListCluster(context.Context, *ListClusterRequest) (*ListClusterResponse, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error)
}

// UnimplementedRdbServer can be embedded to have forward compatible implementations.
type UnimplementedRdbServer struct {
}

func (*UnimplementedRdbServer) CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (*UnimplementedRdbServer) DeleteCluster(context.Context, *DeleteClusterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (*UnimplementedRdbServer) GetCluster(context.Context, *GetClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (*UnimplementedRdbServer) ListCluster(context.Context, *ListClusterRequest) (*ListClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCluster not implemented")
}
func (*UnimplementedRdbServer) UpdateCluster(context.Context, *UpdateClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCluster not implemented")
}

func RegisterRdbServer(s *grpc.Server, srv RdbServer) {
	s.RegisterService(&_Rdb_serviceDesc, srv)
}

func _Rdb_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.rdb.v1alpha.Rdb/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rdb_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.rdb.v1alpha.Rdb/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rdb_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.rdb.v1alpha.Rdb/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rdb_ListCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbServer).ListCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.rdb.v1alpha.Rdb/ListCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbServer).ListCluster(ctx, req.(*ListClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rdb_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.rdb.v1alpha.Rdb/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbServer).UpdateCluster(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rdb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.rdb.v1alpha.Rdb",
	HandlerType: (*RdbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _Rdb_CreateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _Rdb_DeleteCluster_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _Rdb_GetCluster_Handler,
		},
		{
			MethodName: "ListCluster",
			Handler:    _Rdb_ListCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _Rdb_UpdateCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/rdb/v1alpha/rdb.proto",
}