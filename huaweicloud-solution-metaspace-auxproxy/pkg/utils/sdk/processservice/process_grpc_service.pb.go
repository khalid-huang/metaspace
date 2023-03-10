// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: process_grpc_service.proto

package processservice

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

// 定时检查进程的健康状态
type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_process_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_process_grpc_service_proto_rawDescGZIP(), []int{0}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HealthStatus bool `protobuf:"varint,1,opt,name=healthStatus,proto3" json:"healthStatus,omitempty"` // 健康状态，true表示健康，false表示不健康
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_process_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_process_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckResponse) GetHealthStatus() bool {
	if x != nil {
		return x.HealthStatus
	}
	return false
}

// 服务器会话属性详情
type SessionProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 属性名称（键）
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// 属性值（值）
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SessionProperty) Reset() {
	*x = SessionProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionProperty) ProtoMessage() {}

func (x *SessionProperty) ProtoReflect() protoreflect.Message {
	mi := &file_process_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionProperty.ProtoReflect.Descriptor instead.
func (*SessionProperty) Descriptor() ([]byte, []int) {
	return file_process_grpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *SessionProperty) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SessionProperty) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// 服务器会话
type ServerSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerSessionId   string             `protobuf:"bytes,1,opt,name=serverSessionId,proto3" json:"serverSessionId,omitempty"`     // 服务器会话ID
	FleetId           string             `protobuf:"bytes,2,opt,name=fleetId,proto3" json:"fleetId,omitempty"`                     // 队列ID
	Name              string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                           // 服务器会话名
	MaxClients        int32              `protobuf:"varint,4,opt,name=maxClients,proto3" json:"maxClients,omitempty"`              // 最大客户端会话数
	Joinable          bool               `protobuf:"varint,5,opt,name=joinable,proto3" json:"joinable,omitempty"`                  // 是否允许新的客户端会话加入
	SessionProperties []*SessionProperty `protobuf:"bytes,6,rep,name=sessionProperties,proto3" json:"sessionProperties,omitempty"` // 会话属性列表
	Port              int32              `protobuf:"varint,7,opt,name=port,proto3" json:"port,omitempty"`                          // 外部访问的端口，非SCASE服务使用
	IpAddress         string             `protobuf:"bytes,8,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`                 // 外部访问的IP地址，非SCASE服务使用
	SessionData       string             `protobuf:"bytes,9,opt,name=sessionData,proto3" json:"sessionData,omitempty"`             // 会话数据
}

func (x *ServerSession) Reset() {
	*x = ServerSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerSession) ProtoMessage() {}

func (x *ServerSession) ProtoReflect() protoreflect.Message {
	mi := &file_process_grpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerSession.ProtoReflect.Descriptor instead.
func (*ServerSession) Descriptor() ([]byte, []int) {
	return file_process_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *ServerSession) GetServerSessionId() string {
	if x != nil {
		return x.ServerSessionId
	}
	return ""
}

func (x *ServerSession) GetFleetId() string {
	if x != nil {
		return x.FleetId
	}
	return ""
}

func (x *ServerSession) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerSession) GetMaxClients() int32 {
	if x != nil {
		return x.MaxClients
	}
	return 0
}

func (x *ServerSession) GetJoinable() bool {
	if x != nil {
		return x.Joinable
	}
	return false
}

func (x *ServerSession) GetSessionProperties() []*SessionProperty {
	if x != nil {
		return x.SessionProperties
	}
	return nil
}

func (x *ServerSession) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServerSession) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *ServerSession) GetSessionData() string {
	if x != nil {
		return x.SessionData
	}
	return ""
}

// 启动服务器会话
type StartServerSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerSession *ServerSession `protobuf:"bytes,1,opt,name=serverSession,proto3" json:"serverSession,omitempty"` // 服务器会话信息
}

func (x *StartServerSessionRequest) Reset() {
	*x = StartServerSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_grpc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartServerSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartServerSessionRequest) ProtoMessage() {}

func (x *StartServerSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_process_grpc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartServerSessionRequest.ProtoReflect.Descriptor instead.
func (*StartServerSessionRequest) Descriptor() ([]byte, []int) {
	return file_process_grpc_service_proto_rawDescGZIP(), []int{4}
}

func (x *StartServerSessionRequest) GetServerSession() *ServerSession {
	if x != nil {
		return x.ServerSession
	}
	return nil
}

// 结束游戏进程
type ProcessTerminateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerminationTime int64 `protobuf:"varint,1,opt,name=terminationTime,proto3" json:"terminationTime,omitempty"` // 留给进程终止的时间，单位分钟
}

func (x *ProcessTerminateRequest) Reset() {
	*x = ProcessTerminateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_grpc_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessTerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessTerminateRequest) ProtoMessage() {}

func (x *ProcessTerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_process_grpc_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessTerminateRequest.ProtoReflect.Descriptor instead.
func (*ProcessTerminateRequest) Descriptor() ([]byte, []int) {
	return file_process_grpc_service_proto_rawDescGZIP(), []int{5}
}

func (x *ProcessTerminateRequest) GetTerminationTime() int64 {
	if x != nil {
		return x.TerminationTime
	}
	return 0
}

// 返回结果
type ProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProcessResponse) Reset() {
	*x = ProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_grpc_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessResponse) ProtoMessage() {}

func (x *ProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_process_grpc_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessResponse.ProtoReflect.Descriptor instead.
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return file_process_grpc_service_proto_rawDescGZIP(), []int{6}
}

var File_process_grpc_service_proto protoreflect.FileDescriptor

var file_process_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x39, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x39, 0x0a,
	0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc6, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x4d,
	0x0a, 0x11, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x11, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x60, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43,
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x0f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbb, 0x02, 0x0a, 0x15,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x70, 0x63, 0x53, 0x64, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x4f, 0x6e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x64, 0x0a, 0x14, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x12, 0x4f, 0x6e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_process_grpc_service_proto_rawDescOnce sync.Once
	file_process_grpc_service_proto_rawDescData = file_process_grpc_service_proto_rawDesc
)

func file_process_grpc_service_proto_rawDescGZIP() []byte {
	file_process_grpc_service_proto_rawDescOnce.Do(func() {
		file_process_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_process_grpc_service_proto_rawDescData)
	})
	return file_process_grpc_service_proto_rawDescData
}

var file_process_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_process_grpc_service_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),        // 0: processService.HealthCheckRequest
	(*HealthCheckResponse)(nil),       // 1: processService.HealthCheckResponse
	(*SessionProperty)(nil),           // 2: processService.SessionProperty
	(*ServerSession)(nil),             // 3: processService.ServerSession
	(*StartServerSessionRequest)(nil), // 4: processService.StartServerSessionRequest
	(*ProcessTerminateRequest)(nil),   // 5: processService.ProcessTerminateRequest
	(*ProcessResponse)(nil),           // 6: processService.ProcessResponse
}
var file_process_grpc_service_proto_depIdxs = []int32{
	2, // 0: processService.ServerSession.sessionProperties:type_name -> processService.SessionProperty
	3, // 1: processService.StartServerSessionRequest.serverSession:type_name -> processService.ServerSession
	0, // 2: processService.ProcessGrpcSdkService.OnHealthCheck:input_type -> processService.HealthCheckRequest
	4, // 3: processService.ProcessGrpcSdkService.OnStartServerSession:input_type -> processService.StartServerSessionRequest
	5, // 4: processService.ProcessGrpcSdkService.OnProcessTerminate:input_type -> processService.ProcessTerminateRequest
	1, // 5: processService.ProcessGrpcSdkService.OnHealthCheck:output_type -> processService.HealthCheckResponse
	6, // 6: processService.ProcessGrpcSdkService.OnStartServerSession:output_type -> processService.ProcessResponse
	6, // 7: processService.ProcessGrpcSdkService.OnProcessTerminate:output_type -> processService.ProcessResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_process_grpc_service_proto_init() }
func file_process_grpc_service_proto_init() {
	if File_process_grpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_process_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_process_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_process_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionProperty); i {
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
		file_process_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerSession); i {
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
		file_process_grpc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartServerSessionRequest); i {
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
		file_process_grpc_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessTerminateRequest); i {
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
		file_process_grpc_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessResponse); i {
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
			RawDescriptor: file_process_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_process_grpc_service_proto_goTypes,
		DependencyIndexes: file_process_grpc_service_proto_depIdxs,
		MessageInfos:      file_process_grpc_service_proto_msgTypes,
	}.Build()
	File_process_grpc_service_proto = out.File
	file_process_grpc_service_proto_rawDesc = nil
	file_process_grpc_service_proto_goTypes = nil
	file_process_grpc_service_proto_depIdxs = nil
}
