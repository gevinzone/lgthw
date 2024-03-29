// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: blog.proto

package gen

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

type BlogNature int32

const (
	BlogNature_Original   BlogNature = 0
	BlogNature_Reprinting BlogNature = 1
)

// Enum value maps for BlogNature.
var (
	BlogNature_name = map[int32]string{
		0: "Original",
		1: "Reprinting",
	}
	BlogNature_value = map[string]int32{
		"Original":   0,
		"Reprinting": 1,
	}
)

func (x BlogNature) Enum() *BlogNature {
	p := new(BlogNature)
	*p = x
	return p
}

func (x BlogNature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlogNature) Descriptor() protoreflect.EnumDescriptor {
	return file_blog_proto_enumTypes[0].Descriptor()
}

func (BlogNature) Type() protoreflect.EnumType {
	return &file_blog_proto_enumTypes[0]
}

func (x BlogNature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlogNature.Descriptor instead.
func (BlogNature) EnumDescriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{0}
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Author  string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Slug    *string  `protobuf:"bytes,4,opt,name=slug,proto3,oneof" json:"slug,omitempty"`
	Tags    []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *BlogRequest) Reset() {
	*x = BlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogRequest) ProtoMessage() {}

func (x *BlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogRequest.ProtoReflect.Descriptor instead.
func (*BlogRequest) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{1}
}

func (x *BlogRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *BlogRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *BlogRequest) GetSlug() string {
	if x != nil && x.Slug != nil {
		return *x.Slug
	}
	return ""
}

func (x *BlogRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content     string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	ContentHtml *string                `protobuf:"bytes,4,opt,name=content_html,json=contentHtml,proto3,oneof" json:"content_html,omitempty"`
	Author      string                 `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Slug        *string                `protobuf:"bytes,8,opt,name=slug,proto3,oneof" json:"slug,omitempty"`
	Tags        []string               `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{2}
}

func (x *Blog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Blog) GetContentHtml() string {
	if x != nil && x.ContentHtml != nil {
		return *x.ContentHtml
	}
	return ""
}

func (x *Blog) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Blog) GetSlug() string {
	if x != nil && x.Slug != nil {
		return *x.Slug
	}
	return ""
}

func (x *Blog) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Blog) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Blog) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type BlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Blog     *Blog     `protobuf:"bytes,2,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *BlogResponse) Reset() {
	*x = BlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogResponse) ProtoMessage() {}

func (x *BlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogResponse.ProtoReflect.Descriptor instead.
func (*BlogResponse) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{4}
}

func (x *BlogResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *BlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

var File_blog_proto protoreflect.FileDescriptor

var file_blog_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6c,
	0x6f, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x22, 0xc5, 0x02, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x74, 0x6d, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x48, 0x74, 0x6d, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x17, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3a, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x22,
	0x46, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5a, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62,
	0x6c, 0x6f, 0x67, 0x2a, 0x2a, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x67, 0x4e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x32,
	0xc2, 0x01, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x67, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x24, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x08, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f,
	0x67, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42,
	0x6c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x08, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_proto_rawDescOnce sync.Once
	file_blog_proto_rawDescData = file_blog_proto_rawDesc
)

func file_blog_proto_rawDescGZIP() []byte {
	file_blog_proto_rawDescOnce.Do(func() {
		file_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_proto_rawDescData)
	})
	return file_blog_proto_rawDescData
}

var file_blog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_blog_proto_goTypes = []interface{}{
	(BlogNature)(0),               // 0: blog.BlogNature
	(*Id)(nil),                    // 1: blog.Id
	(*BlogRequest)(nil),           // 2: blog.BlogRequest
	(*Blog)(nil),                  // 3: blog.Blog
	(*Response)(nil),              // 4: blog.Response
	(*BlogResponse)(nil),          // 5: blog.BlogResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_blog_proto_depIdxs = []int32{
	6, // 0: blog.Blog.createTime:type_name -> google.protobuf.Timestamp
	6, // 1: blog.Blog.updateTime:type_name -> google.protobuf.Timestamp
	4, // 2: blog.BlogResponse.response:type_name -> blog.Response
	3, // 3: blog.BlogResponse.blog:type_name -> blog.Blog
	1, // 4: blog.BlogAdmin.GetArticle:input_type -> blog.Id
	2, // 5: blog.BlogAdmin.CreateArticle:input_type -> blog.BlogRequest
	2, // 6: blog.BlogAdmin.UpdateArticle:input_type -> blog.BlogRequest
	1, // 7: blog.BlogAdmin.DeleteArticle:input_type -> blog.Id
	3, // 8: blog.BlogAdmin.GetArticle:output_type -> blog.Blog
	3, // 9: blog.BlogAdmin.CreateArticle:output_type -> blog.Blog
	3, // 10: blog.BlogAdmin.UpdateArticle:output_type -> blog.Blog
	4, // 11: blog.BlogAdmin.DeleteArticle:output_type -> blog.Response
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_blog_proto_init() }
func file_blog_proto_init() {
	if File_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogRequest); i {
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
		file_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogResponse); i {
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
	file_blog_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_blog_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_proto_goTypes,
		DependencyIndexes: file_blog_proto_depIdxs,
		EnumInfos:         file_blog_proto_enumTypes,
		MessageInfos:      file_blog_proto_msgTypes,
	}.Build()
	File_blog_proto = out.File
	file_blog_proto_rawDesc = nil
	file_blog_proto_goTypes = nil
	file_blog_proto_depIdxs = nil
}
