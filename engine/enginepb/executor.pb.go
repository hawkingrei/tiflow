// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: engine/proto/executor.proto

package enginepb

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

type PreDispatchTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskTypeId  int64        `protobuf:"varint,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	TaskConfig  []byte       `protobuf:"bytes,2,opt,name=task_config,json=taskConfig,proto3" json:"task_config,omitempty"`
	MasterId    string       `protobuf:"bytes,3,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	WorkerId    string       `protobuf:"bytes,4,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	ProjectInfo *ProjectInfo `protobuf:"bytes,5,opt,name=project_info,json=projectInfo,proto3" json:"project_info,omitempty"`
	// request_id should be a UUID unique for each RPC call.
	RequestId   string `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	WorkerEpoch int64  `protobuf:"varint,7,opt,name=worker_epoch,json=workerEpoch,proto3" json:"worker_epoch,omitempty"`
}

func (x *PreDispatchTaskRequest) Reset() {
	*x = PreDispatchTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_executor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreDispatchTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreDispatchTaskRequest) ProtoMessage() {}

func (x *PreDispatchTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_executor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreDispatchTaskRequest.ProtoReflect.Descriptor instead.
func (*PreDispatchTaskRequest) Descriptor() ([]byte, []int) {
	return file_engine_proto_executor_proto_rawDescGZIP(), []int{0}
}

func (x *PreDispatchTaskRequest) GetTaskTypeId() int64 {
	if x != nil {
		return x.TaskTypeId
	}
	return 0
}

func (x *PreDispatchTaskRequest) GetTaskConfig() []byte {
	if x != nil {
		return x.TaskConfig
	}
	return nil
}

func (x *PreDispatchTaskRequest) GetMasterId() string {
	if x != nil {
		return x.MasterId
	}
	return ""
}

func (x *PreDispatchTaskRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *PreDispatchTaskRequest) GetProjectInfo() *ProjectInfo {
	if x != nil {
		return x.ProjectInfo
	}
	return nil
}

func (x *PreDispatchTaskRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *PreDispatchTaskRequest) GetWorkerEpoch() int64 {
	if x != nil {
		return x.WorkerEpoch
	}
	return 0
}

type PreDispatchTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PreDispatchTaskResponse) Reset() {
	*x = PreDispatchTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_executor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreDispatchTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreDispatchTaskResponse) ProtoMessage() {}

func (x *PreDispatchTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_executor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreDispatchTaskResponse.ProtoReflect.Descriptor instead.
func (*PreDispatchTaskResponse) Descriptor() ([]byte, []int) {
	return file_engine_proto_executor_proto_rawDescGZIP(), []int{1}
}

type ConfirmDispatchTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId  string `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *ConfirmDispatchTaskRequest) Reset() {
	*x = ConfirmDispatchTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_executor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmDispatchTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmDispatchTaskRequest) ProtoMessage() {}

func (x *ConfirmDispatchTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_executor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmDispatchTaskRequest.ProtoReflect.Descriptor instead.
func (*ConfirmDispatchTaskRequest) Descriptor() ([]byte, []int) {
	return file_engine_proto_executor_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmDispatchTaskRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *ConfirmDispatchTaskRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type ConfirmDispatchTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfirmDispatchTaskResponse) Reset() {
	*x = ConfirmDispatchTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_executor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmDispatchTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmDispatchTaskResponse) ProtoMessage() {}

func (x *ConfirmDispatchTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_executor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmDispatchTaskResponse.ProtoReflect.Descriptor instead.
func (*ConfirmDispatchTaskResponse) Descriptor() ([]byte, []int) {
	return file_engine_proto_executor_proto_rawDescGZIP(), []int{3}
}

type RemoveLocalResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	CreatorId  string `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *RemoveLocalResourceRequest) Reset() {
	*x = RemoveLocalResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_executor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveLocalResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveLocalResourceRequest) ProtoMessage() {}

func (x *RemoveLocalResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_executor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveLocalResourceRequest.ProtoReflect.Descriptor instead.
func (*RemoveLocalResourceRequest) Descriptor() ([]byte, []int) {
	return file_engine_proto_executor_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveLocalResourceRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RemoveLocalResourceRequest) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

type RemoveLocalResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveLocalResourceResponse) Reset() {
	*x = RemoveLocalResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_executor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveLocalResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveLocalResourceResponse) ProtoMessage() {}

func (x *RemoveLocalResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_executor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveLocalResourceResponse.ProtoReflect.Descriptor instead.
func (*RemoveLocalResourceResponse) Descriptor() ([]byte, []int) {
	return file_engine_proto_executor_proto_rawDescGZIP(), []int{5}
}

var File_engine_proto_executor_proto protoreflect.FileDescriptor

var file_engine_proto_executor_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x1a, 0x1b, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x19, 0x0a, 0x17, 0x50, 0x72, 0x65, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x1d, 0x0a,
	0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x0a, 0x1a,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd1, 0x01, 0x0a, 0x0f, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a,
	0x0f, 0x50, 0x72, 0x65, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x20, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x72,
	0x65, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x70, 0x0a,
	0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f,
	0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x24, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x70, 0x2f, 0x74, 0x69, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_engine_proto_executor_proto_rawDescOnce sync.Once
	file_engine_proto_executor_proto_rawDescData = file_engine_proto_executor_proto_rawDesc
)

func file_engine_proto_executor_proto_rawDescGZIP() []byte {
	file_engine_proto_executor_proto_rawDescOnce.Do(func() {
		file_engine_proto_executor_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_proto_executor_proto_rawDescData)
	})
	return file_engine_proto_executor_proto_rawDescData
}

var file_engine_proto_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_engine_proto_executor_proto_goTypes = []interface{}{
	(*PreDispatchTaskRequest)(nil),      // 0: enginepb.PreDispatchTaskRequest
	(*PreDispatchTaskResponse)(nil),     // 1: enginepb.PreDispatchTaskResponse
	(*ConfirmDispatchTaskRequest)(nil),  // 2: enginepb.ConfirmDispatchTaskRequest
	(*ConfirmDispatchTaskResponse)(nil), // 3: enginepb.ConfirmDispatchTaskResponse
	(*RemoveLocalResourceRequest)(nil),  // 4: enginepb.RemoveLocalResourceRequest
	(*RemoveLocalResourceResponse)(nil), // 5: enginepb.RemoveLocalResourceResponse
	(*ProjectInfo)(nil),                 // 6: enginepb.ProjectInfo
}
var file_engine_proto_executor_proto_depIdxs = []int32{
	6, // 0: enginepb.PreDispatchTaskRequest.project_info:type_name -> enginepb.ProjectInfo
	0, // 1: enginepb.ExecutorService.PreDispatchTask:input_type -> enginepb.PreDispatchTaskRequest
	2, // 2: enginepb.ExecutorService.ConfirmDispatchTask:input_type -> enginepb.ConfirmDispatchTaskRequest
	4, // 3: enginepb.BrokerService.RemoveResource:input_type -> enginepb.RemoveLocalResourceRequest
	1, // 4: enginepb.ExecutorService.PreDispatchTask:output_type -> enginepb.PreDispatchTaskResponse
	3, // 5: enginepb.ExecutorService.ConfirmDispatchTask:output_type -> enginepb.ConfirmDispatchTaskResponse
	5, // 6: enginepb.BrokerService.RemoveResource:output_type -> enginepb.RemoveLocalResourceResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_engine_proto_executor_proto_init() }
func file_engine_proto_executor_proto_init() {
	if File_engine_proto_executor_proto != nil {
		return
	}
	file_engine_proto_projects_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_engine_proto_executor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreDispatchTaskRequest); i {
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
		file_engine_proto_executor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreDispatchTaskResponse); i {
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
		file_engine_proto_executor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmDispatchTaskRequest); i {
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
		file_engine_proto_executor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmDispatchTaskResponse); i {
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
		file_engine_proto_executor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveLocalResourceRequest); i {
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
		file_engine_proto_executor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveLocalResourceResponse); i {
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
			RawDescriptor: file_engine_proto_executor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_engine_proto_executor_proto_goTypes,
		DependencyIndexes: file_engine_proto_executor_proto_depIdxs,
		MessageInfos:      file_engine_proto_executor_proto_msgTypes,
	}.Build()
	File_engine_proto_executor_proto = out.File
	file_engine_proto_executor_proto_rawDesc = nil
	file_engine_proto_executor_proto_goTypes = nil
	file_engine_proto_executor_proto_depIdxs = nil
}
