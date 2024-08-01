// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: todos/todo-service.proto

package todo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Reminder    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_todos_todo_service_proto_rawDescGZIP(), []int{0}
}

func (x *ToDo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ToDo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToDo) GetReminder() *timestamppb.Timestamp {
	if x != nil {
		return x.Reminder
	}
	return nil
}

type CreateToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToDo *ToDo `protobuf:"bytes,1,opt,name=to_do,json=toDo,proto3" json:"to_do,omitempty"`
}

func (x *CreateToDoRequest) Reset() {
	*x = CreateToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateToDoRequest) ProtoMessage() {}

func (x *CreateToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateToDoRequest.ProtoReflect.Descriptor instead.
func (*CreateToDoRequest) Descriptor() ([]byte, []int) {
	return file_todos_todo_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateToDoRequest) GetToDo() *ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

type CreateToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateToDoResponse) Reset() {
	*x = CreateToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateToDoResponse) ProtoMessage() {}

func (x *CreateToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateToDoResponse.ProtoReflect.Descriptor instead.
func (*CreateToDoResponse) Descriptor() ([]byte, []int) {
	return file_todos_todo_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateToDoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReadToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadToDoRequest) Reset() {
	*x = ReadToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadToDoRequest) ProtoMessage() {}

func (x *ReadToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadToDoRequest.ProtoReflect.Descriptor instead.
func (*ReadToDoRequest) Descriptor() ([]byte, []int) {
	return file_todos_todo_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReadToDoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReadToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToDo *ToDo `protobuf:"bytes,1,opt,name=to_do,json=toDo,proto3" json:"to_do,omitempty"`
}

func (x *ReadToDoResponse) Reset() {
	*x = ReadToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadToDoResponse) ProtoMessage() {}

func (x *ReadToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadToDoResponse.ProtoReflect.Descriptor instead.
func (*ReadToDoResponse) Descriptor() ([]byte, []int) {
	return file_todos_todo_service_proto_rawDescGZIP(), []int{4}
}

func (x *ReadToDoResponse) GetToDo() *ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

type ReadAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadAllRequest) Reset() {
	*x = ReadAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllRequest) ProtoMessage() {}

func (x *ReadAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllRequest.ProtoReflect.Descriptor instead.
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return file_todos_todo_service_proto_rawDescGZIP(), []int{5}
}

type ReadAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToDo []*ToDo `protobuf:"bytes,1,rep,name=to_do,json=toDo,proto3" json:"to_do,omitempty"`
}

func (x *ReadAllResponse) Reset() {
	*x = ReadAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllResponse) ProtoMessage() {}

func (x *ReadAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllResponse.ProtoReflect.Descriptor instead.
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return file_todos_todo_service_proto_rawDescGZIP(), []int{6}
}

func (x *ReadAllResponse) GetToDo() []*ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

var File_todos_todo_service_proto protoreflect.FileDescriptor

var file_todos_todo_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x86, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x05, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x22, 0x24, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x54, 0x6f, 0x44,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x5f,
	0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f,
	0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x0f, 0x52, 0x65,
	0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x05, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x32, 0xb3, 0x01, 0x0a,
	0x0b, 0x54, 0x6f, 0x44, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x6f,
	0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x72, 0x69, 0x65, 0x66, 0x72, 0x6f, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d,
	0x74, 0x6f, 0x2d, 0x64, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todos_todo_service_proto_rawDescOnce sync.Once
	file_todos_todo_service_proto_rawDescData = file_todos_todo_service_proto_rawDesc
)

func file_todos_todo_service_proto_rawDescGZIP() []byte {
	file_todos_todo_service_proto_rawDescOnce.Do(func() {
		file_todos_todo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_todos_todo_service_proto_rawDescData)
	})
	return file_todos_todo_service_proto_rawDescData
}

var file_todos_todo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_todos_todo_service_proto_goTypes = []interface{}{
	(*ToDo)(nil),                  // 0: pb.ToDo
	(*CreateToDoRequest)(nil),     // 1: pb.CreateToDoRequest
	(*CreateToDoResponse)(nil),    // 2: pb.CreateToDoResponse
	(*ReadToDoRequest)(nil),       // 3: pb.ReadToDoRequest
	(*ReadToDoResponse)(nil),      // 4: pb.ReadToDoResponse
	(*ReadAllRequest)(nil),        // 5: pb.ReadAllRequest
	(*ReadAllResponse)(nil),       // 6: pb.ReadAllResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_todos_todo_service_proto_depIdxs = []int32{
	7, // 0: pb.ToDo.reminder:type_name -> google.protobuf.Timestamp
	0, // 1: pb.CreateToDoRequest.to_do:type_name -> pb.ToDo
	0, // 2: pb.ReadToDoResponse.to_do:type_name -> pb.ToDo
	0, // 3: pb.ReadAllResponse.to_do:type_name -> pb.ToDo
	1, // 4: pb.ToDoService.Create:input_type -> pb.CreateToDoRequest
	3, // 5: pb.ToDoService.Read:input_type -> pb.ReadToDoRequest
	5, // 6: pb.ToDoService.ReadAll:input_type -> pb.ReadAllRequest
	2, // 7: pb.ToDoService.Create:output_type -> pb.CreateToDoResponse
	4, // 8: pb.ToDoService.Read:output_type -> pb.ReadToDoResponse
	6, // 9: pb.ToDoService.ReadAll:output_type -> pb.ReadAllResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_todos_todo_service_proto_init() }
func file_todos_todo_service_proto_init() {
	if File_todos_todo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todos_todo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
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
		file_todos_todo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateToDoRequest); i {
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
		file_todos_todo_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateToDoResponse); i {
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
		file_todos_todo_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadToDoRequest); i {
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
		file_todos_todo_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadToDoResponse); i {
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
		file_todos_todo_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllRequest); i {
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
		file_todos_todo_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllResponse); i {
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
			RawDescriptor: file_todos_todo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todos_todo_service_proto_goTypes,
		DependencyIndexes: file_todos_todo_service_proto_depIdxs,
		MessageInfos:      file_todos_todo_service_proto_msgTypes,
	}.Build()
	File_todos_todo_service_proto = out.File
	file_todos_todo_service_proto_rawDesc = nil
	file_todos_todo_service_proto_goTypes = nil
	file_todos_todo_service_proto_depIdxs = nil
}
