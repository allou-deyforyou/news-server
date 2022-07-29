// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: schema.proto

package custom

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type NewsType int32

const (
	NewsType_article NewsType = 0
	NewsType_media   NewsType = 1
)

// Enum value maps for NewsType.
var (
	NewsType_name = map[int32]string{
		0: "article",
		1: "media",
	}
	NewsType_value = map[string]int32{
		"article": 0,
		"media":   1,
	}
)

func (x NewsType) Enum() *NewsType {
	p := new(NewsType)
	*p = x
	return p
}

func (x NewsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NewsType) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_proto_enumTypes[0].Descriptor()
}

func (NewsType) Type() protoreflect.EnumType {
	return &file_schema_proto_enumTypes[0]
}

func (x NewsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NewsType.Descriptor instead.
func (NewsType) EnumDescriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

type MediaPost_Type int32

const (
	MediaPost_YOUTUBE MediaPost_Type = 0
	MediaPost_AUDIO   MediaPost_Type = 1
	MediaPost_VIDEO   MediaPost_Type = 2
)

// Enum value maps for MediaPost_Type.
var (
	MediaPost_Type_name = map[int32]string{
		0: "YOUTUBE",
		1: "AUDIO",
		2: "VIDEO",
	}
	MediaPost_Type_value = map[string]int32{
		"YOUTUBE": 0,
		"AUDIO":   1,
		"VIDEO":   2,
	}
)

func (x MediaPost_Type) Enum() *MediaPost_Type {
	p := new(MediaPost_Type)
	*p = x
	return p
}

func (x MediaPost_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaPost_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_proto_enumTypes[1].Descriptor()
}

func (MediaPost_Type) Type() protoreflect.EnumType {
	return &file_schema_proto_enumTypes[1]
}

func (x MediaPost_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaPost_Type.Descriptor instead.
func (MediaPost_Type) EnumDescriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{4, 0}
}

type ArticlePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Image       string               `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Date        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Link        string               `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Content     string               `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Source      string               `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	Logo        string               `protobuf:"bytes,7,opt,name=logo,proto3" json:"logo,omitempty"`
	Description string               `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ArticlePost) Reset() {
	*x = ArticlePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlePost) ProtoMessage() {}

func (x *ArticlePost) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlePost.ProtoReflect.Descriptor instead.
func (*ArticlePost) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

func (x *ArticlePost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticlePost) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ArticlePost) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ArticlePost) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *ArticlePost) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ArticlePost) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ArticlePost) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *ArticlePost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ArticlePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Country  string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Source   string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Query    string `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	Page     int64  `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ArticlePostRequest) Reset() {
	*x = ArticlePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlePostRequest) ProtoMessage() {}

func (x *ArticlePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlePostRequest.ProtoReflect.Descriptor instead.
func (*ArticlePostRequest) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{1}
}

func (x *ArticlePostRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *ArticlePostRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ArticlePostRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ArticlePostRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ArticlePostRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ArticlePostRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ArticlePostListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ArticlePost `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ArticlePostListResponse) Reset() {
	*x = ArticlePostListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlePostListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlePostListResponse) ProtoMessage() {}

func (x *ArticlePostListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlePostListResponse.ProtoReflect.Descriptor instead.
func (*ArticlePostListResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{2}
}

func (x *ArticlePostListResponse) GetData() []*ArticlePost {
	if x != nil {
		return x.Data
	}
	return nil
}

type ArticlePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ArticlePost `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ArticlePostResponse) Reset() {
	*x = ArticlePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlePostResponse) ProtoMessage() {}

func (x *ArticlePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlePostResponse.ProtoReflect.Descriptor instead.
func (*ArticlePostResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{3}
}

func (x *ArticlePostResponse) GetData() *ArticlePost {
	if x != nil {
		return x.Data
	}
	return nil
}

type MediaPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Image       string               `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Date        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Link        string               `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Content     string               `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Source      string               `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	Logo        string               `protobuf:"bytes,7,opt,name=logo,proto3" json:"logo,omitempty"`
	Description string               `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Live        bool                 `protobuf:"varint,9,opt,name=live,proto3" json:"live,omitempty"`
	Type        MediaPost_Type       `protobuf:"varint,10,opt,name=type,proto3,enum=MediaPost_Type" json:"type,omitempty"`
}

func (x *MediaPost) Reset() {
	*x = MediaPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaPost) ProtoMessage() {}

func (x *MediaPost) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaPost.ProtoReflect.Descriptor instead.
func (*MediaPost) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{4}
}

func (x *MediaPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MediaPost) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *MediaPost) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *MediaPost) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *MediaPost) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MediaPost) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *MediaPost) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *MediaPost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MediaPost) GetLive() bool {
	if x != nil {
		return x.Live
	}
	return false
}

func (x *MediaPost) GetType() MediaPost_Type {
	if x != nil {
		return x.Type
	}
	return MediaPost_YOUTUBE
}

type MediaPostListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*MediaPost `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MediaPostListResponse) Reset() {
	*x = MediaPostListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaPostListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaPostListResponse) ProtoMessage() {}

func (x *MediaPostListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaPostListResponse.ProtoReflect.Descriptor instead.
func (*MediaPostListResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{5}
}

func (x *MediaPostListResponse) GetData() []*MediaPost {
	if x != nil {
		return x.Data
	}
	return nil
}

type MediaPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *MediaPost `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MediaPostResponse) Reset() {
	*x = MediaPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaPostResponse) ProtoMessage() {}

func (x *MediaPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaPostResponse.ProtoReflect.Descriptor instead.
func (*MediaPostResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{6}
}

func (x *MediaPostResponse) GetData() *MediaPost {
	if x != nil {
		return x.Data
	}
	return nil
}

type NewsCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NewsCategory) Reset() {
	*x = NewsCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsCategory) ProtoMessage() {}

func (x *NewsCategory) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsCategory.ProtoReflect.Descriptor instead.
func (*NewsCategory) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{7}
}

func (x *NewsCategory) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *NewsCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NewsCategoryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*NewsCategory `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *NewsCategoryListResponse) Reset() {
	*x = NewsCategoryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsCategoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsCategoryListResponse) ProtoMessage() {}

func (x *NewsCategoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsCategoryListResponse.ProtoReflect.Descriptor instead.
func (*NewsCategoryListResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{8}
}

func (x *NewsCategoryListResponse) GetData() []*NewsCategory {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_schema_proto protoreflect.FileDescriptor

var file_schema_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe5, 0x01, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x3b, 0x0a, 0x17, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x37, 0x0a, 0x13, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc7, 0x02, 0x0a, 0x09, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x69,
	0x76, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x59, 0x4f, 0x55, 0x54, 0x55, 0x42, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f,
	0x10, 0x02, 0x22, 0x37, 0x0a, 0x15, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x11, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x38, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x18, 0x4e, 0x65,
	0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x22, 0x0a, 0x08, 0x4e, 0x65, 0x77,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x10, 0x01, 0x42, 0x09, 0x5a,
	0x07, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_proto_rawDescOnce sync.Once
	file_schema_proto_rawDescData = file_schema_proto_rawDesc
)

func file_schema_proto_rawDescGZIP() []byte {
	file_schema_proto_rawDescOnce.Do(func() {
		file_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_proto_rawDescData)
	})
	return file_schema_proto_rawDescData
}

var file_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_schema_proto_goTypes = []interface{}{
	(NewsType)(0),                    // 0: NewsType
	(MediaPost_Type)(0),              // 1: MediaPost.Type
	(*ArticlePost)(nil),              // 2: ArticlePost
	(*ArticlePostRequest)(nil),       // 3: ArticlePostRequest
	(*ArticlePostListResponse)(nil),  // 4: ArticlePostListResponse
	(*ArticlePostResponse)(nil),      // 5: ArticlePostResponse
	(*MediaPost)(nil),                // 6: MediaPost
	(*MediaPostListResponse)(nil),    // 7: MediaPostListResponse
	(*MediaPostResponse)(nil),        // 8: MediaPostResponse
	(*NewsCategory)(nil),             // 9: NewsCategory
	(*NewsCategoryListResponse)(nil), // 10: NewsCategoryListResponse
	(*timestamp.Timestamp)(nil),      // 11: google.protobuf.Timestamp
}
var file_schema_proto_depIdxs = []int32{
	11, // 0: ArticlePost.date:type_name -> google.protobuf.Timestamp
	2,  // 1: ArticlePostListResponse.data:type_name -> ArticlePost
	2,  // 2: ArticlePostResponse.data:type_name -> ArticlePost
	11, // 3: MediaPost.date:type_name -> google.protobuf.Timestamp
	1,  // 4: MediaPost.type:type_name -> MediaPost.Type
	6,  // 5: MediaPostListResponse.data:type_name -> MediaPost
	6,  // 6: MediaPostResponse.data:type_name -> MediaPost
	9,  // 7: NewsCategoryListResponse.data:type_name -> NewsCategory
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_schema_proto_init() }
func file_schema_proto_init() {
	if File_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlePost); i {
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
		file_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlePostRequest); i {
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
		file_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlePostListResponse); i {
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
		file_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlePostResponse); i {
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
		file_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaPost); i {
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
		file_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaPostListResponse); i {
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
		file_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaPostResponse); i {
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
		file_schema_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsCategory); i {
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
		file_schema_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsCategoryListResponse); i {
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
			RawDescriptor: file_schema_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_proto_goTypes,
		DependencyIndexes: file_schema_proto_depIdxs,
		EnumInfos:         file_schema_proto_enumTypes,
		MessageInfos:      file_schema_proto_msgTypes,
	}.Build()
	File_schema_proto = out.File
	file_schema_proto_rawDesc = nil
	file_schema_proto_goTypes = nil
	file_schema_proto_depIdxs = nil
}
