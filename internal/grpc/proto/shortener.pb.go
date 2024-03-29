// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: internal/grpc/proto/shortener.proto

package proto

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

type URLInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ShortUrl string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *URLInfo) Reset() {
	*x = URLInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLInfo) ProtoMessage() {}

func (x *URLInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLInfo.ProtoReflect.Descriptor instead.
func (*URLInfo) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *URLInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *URLInfo) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type CreateShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateShortURLRequest) Reset() {
	*x = CreateShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLRequest) ProtoMessage() {}

func (x *CreateShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLRequest.ProtoReflect.Descriptor instead.
func (*CreateShortURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateShortURLResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateShortURLResponce) Reset() {
	*x = CreateShortURLResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLResponce) ProtoMessage() {}

func (x *CreateShortURLResponce) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLResponce.ProtoReflect.Descriptor instead.
func (*CreateShortURLResponce) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *CreateShortURLResponce) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *CreateShortURLResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CreateBatchShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url []string `protobuf:"bytes,1,rep,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateBatchShortURLRequest) Reset() {
	*x = CreateBatchShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchShortURLRequest) ProtoMessage() {}

func (x *CreateBatchShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchShortURLRequest.ProtoReflect.Descriptor instead.
func (*CreateBatchShortURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBatchShortURLRequest) GetUrl() []string {
	if x != nil {
		return x.Url
	}
	return nil
}

type CreateBatchShortURLResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl []string `protobuf:"bytes,1,rep,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	Error    string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateBatchShortURLResponce) Reset() {
	*x = CreateBatchShortURLResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchShortURLResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchShortURLResponce) ProtoMessage() {}

func (x *CreateBatchShortURLResponce) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchShortURLResponce.ProtoReflect.Descriptor instead.
func (*CreateBatchShortURLResponce) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBatchShortURLResponce) GetShortUrl() []string {
	if x != nil {
		return x.ShortUrl
	}
	return nil
}

func (x *CreateBatchShortURLResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetFullURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetFullURLRequest) Reset() {
	*x = GetFullURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullURLRequest) ProtoMessage() {}

func (x *GetFullURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullURLRequest.ProtoReflect.Descriptor instead.
func (*GetFullURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *GetFullURLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetFullURLResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetFullURLResponce) Reset() {
	*x = GetFullURLResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullURLResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullURLResponce) ProtoMessage() {}

func (x *GetFullURLResponce) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullURLResponce.ProtoReflect.Descriptor instead.
func (*GetFullURLResponce) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *GetFullURLResponce) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetFullURLResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetUserURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserURLRequest) Reset() {
	*x = GetUserURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserURLRequest) ProtoMessage() {}

func (x *GetUserURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserURLRequest.ProtoReflect.Descriptor instead.
func (*GetUserURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserURLRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserURLResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlInfo []*URLInfo `protobuf:"bytes,1,rep,name=url_info,json=urlInfo,proto3" json:"url_info,omitempty"`
}

func (x *GetUserURLResponce) Reset() {
	*x = GetUserURLResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserURLResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserURLResponce) ProtoMessage() {}

func (x *GetUserURLResponce) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserURLResponce.ProtoReflect.Descriptor instead.
func (*GetUserURLResponce) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserURLResponce) GetUrlInfo() []*URLInfo {
	if x != nil {
		return x.UrlInfo
	}
	return nil
}

type DeleteUserURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShortUrl []string `protobuf:"bytes,2,rep,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *DeleteUserURLRequest) Reset() {
	*x = DeleteUserURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserURLRequest) ProtoMessage() {}

func (x *DeleteUserURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserURLRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserURLRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteUserURLRequest) GetShortUrl() []string {
	if x != nil {
		return x.ShortUrl
	}
	return nil
}

type DeleteUserURLResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteUserURLResponce) Reset() {
	*x = DeleteUserURLResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserURLResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserURLResponce) ProtoMessage() {}

func (x *DeleteUserURLResponce) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserURLResponce.ProtoReflect.Descriptor instead.
func (*DeleteUserURLResponce) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteUserURLResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetShortenerStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetShortenerStatsRequest) Reset() {
	*x = GetShortenerStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShortenerStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShortenerStatsRequest) ProtoMessage() {}

func (x *GetShortenerStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShortenerStatsRequest.ProtoReflect.Descriptor instead.
func (*GetShortenerStatsRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{11}
}

func (x *GetShortenerStatsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetShortenerStatsResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlCount  int32 `protobuf:"varint,1,opt,name=url_count,json=urlCount,proto3" json:"url_count,omitempty"`
	UserCount int32 `protobuf:"varint,2,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
}

func (x *GetShortenerStatsResponce) Reset() {
	*x = GetShortenerStatsResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_shortener_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShortenerStatsResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShortenerStatsResponce) ProtoMessage() {}

func (x *GetShortenerStatsResponce) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_shortener_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShortenerStatsResponce.ProtoReflect.Descriptor instead.
func (*GetShortenerStatsResponce) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_shortener_proto_rawDescGZIP(), []int{12}
}

func (x *GetShortenerStatsResponce) GetUrlCount() int32 {
	if x != nil {
		return x.UrlCount
	}
	return 0
}

func (x *GetShortenerStatsResponce) GetUserCount() int32 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

var File_internal_grpc_proto_shortener_proto protoreflect.FileDescriptor

var file_internal_grpc_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x07,
	0x55, 0x52, 0x4c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x29, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x4b, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2e,
	0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x50,
	0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x30, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x52, 0x4c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x4c, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2d, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x33, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x57, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x72, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x75, 0x72, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xe2, 0x03, 0x0a, 0x09, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c,
	0x55, 0x52, 0x4c, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x64,
	0x79, 0x61, 0x63, 0x68, 0x6b, 0x6f, 0x76, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_grpc_proto_shortener_proto_rawDescOnce sync.Once
	file_internal_grpc_proto_shortener_proto_rawDescData = file_internal_grpc_proto_shortener_proto_rawDesc
)

func file_internal_grpc_proto_shortener_proto_rawDescGZIP() []byte {
	file_internal_grpc_proto_shortener_proto_rawDescOnce.Do(func() {
		file_internal_grpc_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_proto_shortener_proto_rawDescData)
	})
	return file_internal_grpc_proto_shortener_proto_rawDescData
}

var file_internal_grpc_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_internal_grpc_proto_shortener_proto_goTypes = []interface{}{
	(*URLInfo)(nil),                     // 0: proto.URLInfo
	(*CreateShortURLRequest)(nil),       // 1: proto.CreateShortURLRequest
	(*CreateShortURLResponce)(nil),      // 2: proto.CreateShortURLResponce
	(*CreateBatchShortURLRequest)(nil),  // 3: proto.CreateBatchShortURLRequest
	(*CreateBatchShortURLResponce)(nil), // 4: proto.CreateBatchShortURLResponce
	(*GetFullURLRequest)(nil),           // 5: proto.GetFullURLRequest
	(*GetFullURLResponce)(nil),          // 6: proto.GetFullURLResponce
	(*GetUserURLRequest)(nil),           // 7: proto.GetUserURLRequest
	(*GetUserURLResponce)(nil),          // 8: proto.GetUserURLResponce
	(*DeleteUserURLRequest)(nil),        // 9: proto.DeleteUserURLRequest
	(*DeleteUserURLResponce)(nil),       // 10: proto.DeleteUserURLResponce
	(*GetShortenerStatsRequest)(nil),    // 11: proto.GetShortenerStatsRequest
	(*GetShortenerStatsResponce)(nil),   // 12: proto.GetShortenerStatsResponce
}
var file_internal_grpc_proto_shortener_proto_depIdxs = []int32{
	0,  // 0: proto.GetUserURLResponce.url_info:type_name -> proto.URLInfo
	1,  // 1: proto.Shortener.CreateShortURL:input_type -> proto.CreateShortURLRequest
	3,  // 2: proto.Shortener.CreateBatchShortURL:input_type -> proto.CreateBatchShortURLRequest
	5,  // 3: proto.Shortener.GetFullURL:input_type -> proto.GetFullURLRequest
	7,  // 4: proto.Shortener.GetUserURL:input_type -> proto.GetUserURLRequest
	9,  // 5: proto.Shortener.DeleteUserURL:input_type -> proto.DeleteUserURLRequest
	11, // 6: proto.Shortener.GetShortenerStats:input_type -> proto.GetShortenerStatsRequest
	2,  // 7: proto.Shortener.CreateShortURL:output_type -> proto.CreateShortURLResponce
	4,  // 8: proto.Shortener.CreateBatchShortURL:output_type -> proto.CreateBatchShortURLResponce
	6,  // 9: proto.Shortener.GetFullURL:output_type -> proto.GetFullURLResponce
	8,  // 10: proto.Shortener.GetUserURL:output_type -> proto.GetUserURLResponce
	10, // 11: proto.Shortener.DeleteUserURL:output_type -> proto.DeleteUserURLResponce
	12, // 12: proto.Shortener.GetShortenerStats:output_type -> proto.GetShortenerStatsResponce
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_internal_grpc_proto_shortener_proto_init() }
func file_internal_grpc_proto_shortener_proto_init() {
	if File_internal_grpc_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpc_proto_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLInfo); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLRequest); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLResponce); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchShortURLRequest); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchShortURLResponce); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullURLRequest); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullURLResponce); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserURLRequest); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserURLResponce); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserURLRequest); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserURLResponce); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShortenerStatsRequest); i {
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
		file_internal_grpc_proto_shortener_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShortenerStatsResponce); i {
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
			RawDescriptor: file_internal_grpc_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_proto_shortener_proto_goTypes,
		DependencyIndexes: file_internal_grpc_proto_shortener_proto_depIdxs,
		MessageInfos:      file_internal_grpc_proto_shortener_proto_msgTypes,
	}.Build()
	File_internal_grpc_proto_shortener_proto = out.File
	file_internal_grpc_proto_shortener_proto_rawDesc = nil
	file_internal_grpc_proto_shortener_proto_goTypes = nil
	file_internal_grpc_proto_shortener_proto_depIdxs = nil
}
