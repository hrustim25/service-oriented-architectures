// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.15.8
// source: stat_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetEventsCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetEventsCountRequest) Reset() {
	*x = GetEventsCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsCountRequest) ProtoMessage() {}

func (x *GetEventsCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsCountRequest.ProtoReflect.Descriptor instead.
func (*GetEventsCountRequest) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventsCountRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type GetEventsCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ViewCount uint32 `protobuf:"varint,1,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	LikeCount uint32 `protobuf:"varint,2,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
}

func (x *GetEventsCountResponse) Reset() {
	*x = GetEventsCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsCountResponse) ProtoMessage() {}

func (x *GetEventsCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsCountResponse.ProtoReflect.Descriptor instead.
func (*GetEventsCountResponse) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventsCountResponse) GetViewCount() uint32 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *GetEventsCountResponse) GetLikeCount() uint32 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

type GetTopTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType uint32 `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
}

func (x *GetTopTasksRequest) Reset() {
	*x = GetTopTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopTasksRequest) ProtoMessage() {}

func (x *GetTopTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopTasksRequest.ProtoReflect.Descriptor instead.
func (*GetTopTasksRequest) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTopTasksRequest) GetEventType() uint32 {
	if x != nil {
		return x.EventType
	}
	return 0
}

type GetTopTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*TopTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetTopTasksResponse) Reset() {
	*x = GetTopTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopTasksResponse) ProtoMessage() {}

func (x *GetTopTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopTasksResponse.ProtoReflect.Descriptor instead.
func (*GetTopTasksResponse) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTopTasksResponse) GetTasks() []*TopTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type GetTopAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*TopAuthor `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *GetTopAuthorsResponse) Reset() {
	*x = GetTopAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopAuthorsResponse) ProtoMessage() {}

func (x *GetTopAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopAuthorsResponse.ProtoReflect.Descriptor instead.
func (*GetTopAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTopAuthorsResponse) GetAuthors() []*TopAuthor {
	if x != nil {
		return x.Authors
	}
	return nil
}

type TopTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId       uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskAuthorId uint64 `protobuf:"varint,2,opt,name=task_author_id,json=taskAuthorId,proto3" json:"task_author_id,omitempty"`
	ViewCount    uint32 `protobuf:"varint,3,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	LikeCount    uint32 `protobuf:"varint,4,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
}

func (x *TopTask) Reset() {
	*x = TopTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopTask) ProtoMessage() {}

func (x *TopTask) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopTask.ProtoReflect.Descriptor instead.
func (*TopTask) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{5}
}

func (x *TopTask) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TopTask) GetTaskAuthorId() uint64 {
	if x != nil {
		return x.TaskAuthorId
	}
	return 0
}

func (x *TopTask) GetViewCount() uint32 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *TopTask) GetLikeCount() uint32 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

type TopAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskAuthorId uint64 `protobuf:"varint,1,opt,name=task_author_id,json=taskAuthorId,proto3" json:"task_author_id,omitempty"`
	LikeCount    uint32 `protobuf:"varint,2,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
}

func (x *TopAuthor) Reset() {
	*x = TopAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopAuthor) ProtoMessage() {}

func (x *TopAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopAuthor.ProtoReflect.Descriptor instead.
func (*TopAuthor) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{6}
}

func (x *TopAuthor) GetTaskAuthorId() uint64 {
	if x != nil {
		return x.TaskAuthorId
	}
	return 0
}

func (x *TopAuthor) GetLikeCount() uint32 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

var File_stat_service_proto protoreflect.FileDescriptor

var file_stat_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x35, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x54, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x3d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x70, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x54, 0x6f, 0x70, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x54, 0x6f, 0x70, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x50, 0x0a, 0x09, 0x54, 0x6f, 0x70, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0e,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0xcb, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x70, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_stat_service_proto_rawDescOnce sync.Once
	file_stat_service_proto_rawDescData = file_stat_service_proto_rawDesc
)

func file_stat_service_proto_rawDescGZIP() []byte {
	file_stat_service_proto_rawDescOnce.Do(func() {
		file_stat_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stat_service_proto_rawDescData)
	})
	return file_stat_service_proto_rawDescData
}

var file_stat_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_stat_service_proto_goTypes = []interface{}{
	(*GetEventsCountRequest)(nil),  // 0: GetEventsCountRequest
	(*GetEventsCountResponse)(nil), // 1: GetEventsCountResponse
	(*GetTopTasksRequest)(nil),     // 2: GetTopTasksRequest
	(*GetTopTasksResponse)(nil),    // 3: GetTopTasksResponse
	(*GetTopAuthorsResponse)(nil),  // 4: GetTopAuthorsResponse
	(*TopTask)(nil),                // 5: TopTask
	(*TopAuthor)(nil),              // 6: TopAuthor
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_stat_service_proto_depIdxs = []int32{
	5, // 0: GetTopTasksResponse.tasks:type_name -> TopTask
	6, // 1: GetTopAuthorsResponse.authors:type_name -> TopAuthor
	0, // 2: StatService.GetEventsCount:input_type -> GetEventsCountRequest
	2, // 3: StatService.GetTopTasks:input_type -> GetTopTasksRequest
	7, // 4: StatService.GetTopAuthors:input_type -> google.protobuf.Empty
	1, // 5: StatService.GetEventsCount:output_type -> GetEventsCountResponse
	3, // 6: StatService.GetTopTasks:output_type -> GetTopTasksResponse
	4, // 7: StatService.GetTopAuthors:output_type -> GetTopAuthorsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stat_service_proto_init() }
func file_stat_service_proto_init() {
	if File_stat_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stat_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsCountRequest); i {
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
		file_stat_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsCountResponse); i {
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
		file_stat_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopTasksRequest); i {
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
		file_stat_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopTasksResponse); i {
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
		file_stat_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopAuthorsResponse); i {
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
		file_stat_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopTask); i {
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
		file_stat_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopAuthor); i {
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
			RawDescriptor: file_stat_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stat_service_proto_goTypes,
		DependencyIndexes: file_stat_service_proto_depIdxs,
		MessageInfos:      file_stat_service_proto_msgTypes,
	}.Build()
	File_stat_service_proto = out.File
	file_stat_service_proto_rawDesc = nil
	file_stat_service_proto_goTypes = nil
	file_stat_service_proto_depIdxs = nil
}
