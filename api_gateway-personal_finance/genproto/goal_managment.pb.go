// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: goal_managment.proto

package genproto

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

type Responsee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Responsee) Reset() {
	*x = Responsee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Responsee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Responsee) ProtoMessage() {}

func (x *Responsee) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Responsee.ProtoReflect.Descriptor instead.
func (*Responsee) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{0}
}

func (x *Responsee) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateGoalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,4,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,5,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount,omitempty"`
	Deadline      string  `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status        string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateGoalRequest) Reset() {
	*x = CreateGoalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoalRequest) ProtoMessage() {}

func (x *CreateGoalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoalRequest.ProtoReflect.Descriptor instead.
func (*CreateGoalRequest) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGoalRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateGoalRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateGoalRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGoalRequest) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *CreateGoalRequest) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *CreateGoalRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *CreateGoalRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ListGoalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,4,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,5,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount,omitempty"`
	Deadline      string  `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status        string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ListGoalsRequest) Reset() {
	*x = ListGoalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoalsRequest) ProtoMessage() {}

func (x *ListGoalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoalsRequest.ProtoReflect.Descriptor instead.
func (*ListGoalsRequest) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{2}
}

func (x *ListGoalsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListGoalsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListGoalsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListGoalsRequest) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *ListGoalsRequest) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *ListGoalsRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *ListGoalsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetGoalByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoalId string `protobuf:"bytes,1,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
}

func (x *GetGoalByIdRequest) Reset() {
	*x = GetGoalByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoalByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoalByIdRequest) ProtoMessage() {}

func (x *GetGoalByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoalByIdRequest.ProtoReflect.Descriptor instead.
func (*GetGoalByIdRequest) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{3}
}

func (x *GetGoalByIdRequest) GetGoalId() string {
	if x != nil {
		return x.GoalId
	}
	return ""
}

type UpdateGoalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoalId        string  `protobuf:"bytes,1,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,4,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,5,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount,omitempty"`
	Deadline      string  `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status        string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateGoalRequest) Reset() {
	*x = UpdateGoalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoalRequest) ProtoMessage() {}

func (x *UpdateGoalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoalRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoalRequest) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGoalRequest) GetGoalId() string {
	if x != nil {
		return x.GoalId
	}
	return ""
}

func (x *UpdateGoalRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGoalRequest) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *UpdateGoalRequest) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *UpdateGoalRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *UpdateGoalRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteGoalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoalId string `protobuf:"bytes,1,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
}

func (x *DeleteGoalRequest) Reset() {
	*x = DeleteGoalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoalRequest) ProtoMessage() {}

func (x *DeleteGoalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoalRequest.ProtoReflect.Descriptor instead.
func (*DeleteGoalRequest) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGoalRequest) GetGoalId() string {
	if x != nil {
		return x.GoalId
	}
	return ""
}

type GoalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoalId        string  `protobuf:"bytes,1,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,4,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,5,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount,omitempty"`
	Deadline      string  `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status        string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GoalResponse) Reset() {
	*x = GoalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalResponse) ProtoMessage() {}

func (x *GoalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalResponse.ProtoReflect.Descriptor instead.
func (*GoalResponse) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{6}
}

func (x *GoalResponse) GetGoalId() string {
	if x != nil {
		return x.GoalId
	}
	return ""
}

func (x *GoalResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoalResponse) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *GoalResponse) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *GoalResponse) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *GoalResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ListGoalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goals []*GoalResponse `protobuf:"bytes,1,rep,name=goals,proto3" json:"goals,omitempty"`
}

func (x *ListGoalsResponse) Reset() {
	*x = ListGoalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoalsResponse) ProtoMessage() {}

func (x *ListGoalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoalsResponse.ProtoReflect.Descriptor instead.
func (*ListGoalsResponse) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{7}
}

func (x *ListGoalsResponse) GetGoals() []*GoalResponse {
	if x != nil {
		return x.Goals
	}
	return nil
}

type GoalDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GoalDeleteResponse) Reset() {
	*x = GoalDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_managment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalDeleteResponse) ProtoMessage() {}

func (x *GoalDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goal_managment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalDeleteResponse.ProtoReflect.Descriptor instead.
func (*GoalDeleteResponse) Descriptor() ([]byte, []int) {
	return file_goal_managment_proto_rawDescGZIP(), []int{8}
}

func (x *GoalDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_goal_managment_proto protoreflect.FileDescriptor

var file_goal_managment_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x25,
	0x0a, 0x09, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x61, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x6f, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x6f, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x6f, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xbb, 0x01, 0x0a, 0x0c, 0x47,
	0x6f, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67,
	0x6f, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x6f,
	0x61, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x6f, 0x61,
	0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xcd, 0x02, 0x0a, 0x0b, 0x47, 0x6f,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x61,
	0x6c, 0x73, 0x12, 0x18, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x6f,
	0x61, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x61, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x47, 0x6f, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f,
	0x61, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goal_managment_proto_rawDescOnce sync.Once
	file_goal_managment_proto_rawDescData = file_goal_managment_proto_rawDesc
)

func file_goal_managment_proto_rawDescGZIP() []byte {
	file_goal_managment_proto_rawDescOnce.Do(func() {
		file_goal_managment_proto_rawDescData = protoimpl.X.CompressGZIP(file_goal_managment_proto_rawDescData)
	})
	return file_goal_managment_proto_rawDescData
}

var file_goal_managment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_goal_managment_proto_goTypes = []interface{}{
	(*Responsee)(nil),          // 0: budget.Responsee
	(*CreateGoalRequest)(nil),  // 1: budget.CreateGoalRequest
	(*ListGoalsRequest)(nil),   // 2: budget.ListGoalsRequest
	(*GetGoalByIdRequest)(nil), // 3: budget.GetGoalByIdRequest
	(*UpdateGoalRequest)(nil),  // 4: budget.UpdateGoalRequest
	(*DeleteGoalRequest)(nil),  // 5: budget.DeleteGoalRequest
	(*GoalResponse)(nil),       // 6: budget.GoalResponse
	(*ListGoalsResponse)(nil),  // 7: budget.ListGoalsResponse
	(*GoalDeleteResponse)(nil), // 8: budget.GoalDeleteResponse
}
var file_goal_managment_proto_depIdxs = []int32{
	6, // 0: budget.ListGoalsResponse.goals:type_name -> budget.GoalResponse
	1, // 1: budget.GoalService.CreateGoal:input_type -> budget.CreateGoalRequest
	2, // 2: budget.GoalService.ListGoals:input_type -> budget.ListGoalsRequest
	3, // 3: budget.GoalService.GetGoalById:input_type -> budget.GetGoalByIdRequest
	4, // 4: budget.GoalService.UpdateGoal:input_type -> budget.UpdateGoalRequest
	5, // 5: budget.GoalService.DeleteGoal:input_type -> budget.DeleteGoalRequest
	0, // 6: budget.GoalService.CreateGoal:output_type -> budget.Responsee
	7, // 7: budget.GoalService.ListGoals:output_type -> budget.ListGoalsResponse
	6, // 8: budget.GoalService.GetGoalById:output_type -> budget.GoalResponse
	0, // 9: budget.GoalService.UpdateGoal:output_type -> budget.Responsee
	8, // 10: budget.GoalService.DeleteGoal:output_type -> budget.GoalDeleteResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goal_managment_proto_init() }
func file_goal_managment_proto_init() {
	if File_goal_managment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goal_managment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Responsee); i {
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
		file_goal_managment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoalRequest); i {
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
		file_goal_managment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoalsRequest); i {
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
		file_goal_managment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoalByIdRequest); i {
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
		file_goal_managment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoalRequest); i {
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
		file_goal_managment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoalRequest); i {
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
		file_goal_managment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalResponse); i {
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
		file_goal_managment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoalsResponse); i {
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
		file_goal_managment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalDeleteResponse); i {
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
			RawDescriptor: file_goal_managment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goal_managment_proto_goTypes,
		DependencyIndexes: file_goal_managment_proto_depIdxs,
		MessageInfos:      file_goal_managment_proto_msgTypes,
	}.Build()
	File_goal_managment_proto = out.File
	file_goal_managment_proto_rawDesc = nil
	file_goal_managment_proto_goTypes = nil
	file_goal_managment_proto_depIdxs = nil
}
