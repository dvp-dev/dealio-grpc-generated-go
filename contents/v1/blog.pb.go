// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: contents/v1/blog.proto

package contents_grpc

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

type BlogGetListRequest_SortBy int32

const (
	BlogGetListRequest_OLDEST_UNSPECIFIED BlogGetListRequest_SortBy = 0
	BlogGetListRequest_NEWEST             BlogGetListRequest_SortBy = 1
	BlogGetListRequest_TITLE_A_TO_Z       BlogGetListRequest_SortBy = 2
	BlogGetListRequest_TITLE_Z_TO_A       BlogGetListRequest_SortBy = 3
)

// Enum value maps for BlogGetListRequest_SortBy.
var (
	BlogGetListRequest_SortBy_name = map[int32]string{
		0: "OLDEST_UNSPECIFIED",
		1: "NEWEST",
		2: "TITLE_A_TO_Z",
		3: "TITLE_Z_TO_A",
	}
	BlogGetListRequest_SortBy_value = map[string]int32{
		"OLDEST_UNSPECIFIED": 0,
		"NEWEST":             1,
		"TITLE_A_TO_Z":       2,
		"TITLE_Z_TO_A":       3,
	}
)

func (x BlogGetListRequest_SortBy) Enum() *BlogGetListRequest_SortBy {
	p := new(BlogGetListRequest_SortBy)
	*p = x
	return p
}

func (x BlogGetListRequest_SortBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlogGetListRequest_SortBy) Descriptor() protoreflect.EnumDescriptor {
	return file_contents_v1_blog_proto_enumTypes[0].Descriptor()
}

func (BlogGetListRequest_SortBy) Type() protoreflect.EnumType {
	return &file_contents_v1_blog_proto_enumTypes[0]
}

func (x BlogGetListRequest_SortBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlogGetListRequest_SortBy.Descriptor instead.
func (BlogGetListRequest_SortBy) EnumDescriptor() ([]byte, []int) {
	return file_contents_v1_blog_proto_rawDescGZIP(), []int{1, 0}
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                 string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Post                  string                 `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	Slug                  string                 `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	Keyword               []string               `protobuf:"bytes,5,rep,name=keyword,proto3" json:"keyword,omitempty"`
	MetaDescription       string                 `protobuf:"bytes,6,opt,name=meta_description,json=metaDescription,proto3" json:"meta_description,omitempty"`
	UtmLink               string                 `protobuf:"bytes,7,opt,name=utm_link,json=utmLink,proto3" json:"utm_link,omitempty"`
	ImageCover            string                 `protobuf:"bytes,8,opt,name=image_cover,json=imageCover,proto3" json:"image_cover,omitempty"`
	Prods                 int32                  `protobuf:"varint,9,opt,name=prods,proto3" json:"prods,omitempty"`
	RelatedDealsId        []int32                `protobuf:"varint,10,rep,packed,name=related_deals_id,json=relatedDealsId,proto3" json:"related_deals_id,omitempty"`
	RelatedProductDealsId []int32                `protobuf:"varint,11,rep,packed,name=related_product_deals_id,json=relatedProductDealsId,proto3" json:"related_product_deals_id,omitempty"`
	TotalClick            int32                  `protobuf:"varint,12,opt,name=total_click,json=totalClick,proto3" json:"total_click,omitempty"`
	Datetime              *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=datetime,proto3" json:"datetime,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_v1_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_contents_v1_blog_proto_msgTypes[0]
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
	return file_contents_v1_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetId() int32 {
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

func (x *Blog) GetPost() string {
	if x != nil {
		return x.Post
	}
	return ""
}

func (x *Blog) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Blog) GetKeyword() []string {
	if x != nil {
		return x.Keyword
	}
	return nil
}

func (x *Blog) GetMetaDescription() string {
	if x != nil {
		return x.MetaDescription
	}
	return ""
}

func (x *Blog) GetUtmLink() string {
	if x != nil {
		return x.UtmLink
	}
	return ""
}

func (x *Blog) GetImageCover() string {
	if x != nil {
		return x.ImageCover
	}
	return ""
}

func (x *Blog) GetProds() int32 {
	if x != nil {
		return x.Prods
	}
	return 0
}

func (x *Blog) GetRelatedDealsId() []int32 {
	if x != nil {
		return x.RelatedDealsId
	}
	return nil
}

func (x *Blog) GetRelatedProductDealsId() []int32 {
	if x != nil {
		return x.RelatedProductDealsId
	}
	return nil
}

func (x *Blog) GetTotalClick() int32 {
	if x != nil {
		return x.TotalClick
	}
	return 0
}

func (x *Blog) GetDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

type BlogGetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlyActive     bool                      `protobuf:"varint,1,opt,name=only_active,json=onlyActive,proto3" json:"only_active,omitempty"`
	Page           int32                     `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	ContentPerPage int32                     `protobuf:"varint,3,opt,name=content_per_page,json=contentPerPage,proto3" json:"content_per_page,omitempty"`
	SortBy         BlogGetListRequest_SortBy `protobuf:"varint,4,opt,name=sort_by,json=sortBy,proto3,enum=contents.v1.BlogGetListRequest_SortBy" json:"sort_by,omitempty"`
}

func (x *BlogGetListRequest) Reset() {
	*x = BlogGetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_v1_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogGetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogGetListRequest) ProtoMessage() {}

func (x *BlogGetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contents_v1_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogGetListRequest.ProtoReflect.Descriptor instead.
func (*BlogGetListRequest) Descriptor() ([]byte, []int) {
	return file_contents_v1_blog_proto_rawDescGZIP(), []int{1}
}

func (x *BlogGetListRequest) GetOnlyActive() bool {
	if x != nil {
		return x.OnlyActive
	}
	return false
}

func (x *BlogGetListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *BlogGetListRequest) GetContentPerPage() int32 {
	if x != nil {
		return x.ContentPerPage
	}
	return 0
}

func (x *BlogGetListRequest) GetSortBy() BlogGetListRequest_SortBy {
	if x != nil {
		return x.SortBy
	}
	return BlogGetListRequest_OLDEST_UNSPECIFIED
}

type BlogGetListResponseStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Blog `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlogGetListResponseStream) Reset() {
	*x = BlogGetListResponseStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_v1_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogGetListResponseStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogGetListResponseStream) ProtoMessage() {}

func (x *BlogGetListResponseStream) ProtoReflect() protoreflect.Message {
	mi := &file_contents_v1_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogGetListResponseStream.ProtoReflect.Descriptor instead.
func (*BlogGetListResponseStream) Descriptor() ([]byte, []int) {
	return file_contents_v1_blog_proto_rawDescGZIP(), []int{2}
}

func (x *BlogGetListResponseStream) GetData() *Blog {
	if x != nil {
		return x.Data
	}
	return nil
}

type BlogGetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *BlogGetOneRequest) Reset() {
	*x = BlogGetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_v1_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogGetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogGetOneRequest) ProtoMessage() {}

func (x *BlogGetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contents_v1_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogGetOneRequest.ProtoReflect.Descriptor instead.
func (*BlogGetOneRequest) Descriptor() ([]byte, []int) {
	return file_contents_v1_blog_proto_rawDescGZIP(), []int{3}
}

func (x *BlogGetOneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlogGetOneRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type BlogGetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Blog `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlogGetOneResponse) Reset() {
	*x = BlogGetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_v1_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogGetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogGetOneResponse) ProtoMessage() {}

func (x *BlogGetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contents_v1_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogGetOneResponse.ProtoReflect.Descriptor instead.
func (*BlogGetOneResponse) Descriptor() ([]byte, []int) {
	return file_contents_v1_blog_proto_rawDescGZIP(), []int{4}
}

func (x *BlogGetOneResponse) GetData() *Blog {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_contents_v1_blog_proto protoreflect.FileDescriptor

var file_contents_v1_blog_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x03, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x74, 0x6d, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x74, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x6f, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x72, 0x6f, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x64, 0x65, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x73, 0x49, 0x64, 0x12, 0x37,
	0x0a, 0x18, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x15, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x65, 0x61, 0x6c, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x86, 0x02, 0x0a, 0x12, 0x42, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x6e, 0x6c, 0x79, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6f, 0x6e,
	0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x52,
	0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x22, 0x50, 0x0a, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x4c, 0x44, 0x45, 0x53, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x45, 0x57,
	0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x41,
	0x5f, 0x54, 0x4f, 0x5f, 0x5a, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x54, 0x4c, 0x45,
	0x5f, 0x5a, 0x5f, 0x54, 0x4f, 0x5f, 0x41, 0x10, 0x03, 0x22, 0x42, 0x0a, 0x19, 0x42, 0x6c, 0x6f,
	0x67, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a,
	0x11, 0x42, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x3b, 0x0a, 0x12, 0x42, 0x6c, 0x6f, 0x67, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x76, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x65, 0x61, 0x6c, 0x69, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d,
	0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contents_v1_blog_proto_rawDescOnce sync.Once
	file_contents_v1_blog_proto_rawDescData = file_contents_v1_blog_proto_rawDesc
)

func file_contents_v1_blog_proto_rawDescGZIP() []byte {
	file_contents_v1_blog_proto_rawDescOnce.Do(func() {
		file_contents_v1_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_contents_v1_blog_proto_rawDescData)
	})
	return file_contents_v1_blog_proto_rawDescData
}

var file_contents_v1_blog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contents_v1_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contents_v1_blog_proto_goTypes = []interface{}{
	(BlogGetListRequest_SortBy)(0),    // 0: contents.v1.BlogGetListRequest.SortBy
	(*Blog)(nil),                      // 1: contents.v1.Blog
	(*BlogGetListRequest)(nil),        // 2: contents.v1.BlogGetListRequest
	(*BlogGetListResponseStream)(nil), // 3: contents.v1.BlogGetListResponseStream
	(*BlogGetOneRequest)(nil),         // 4: contents.v1.BlogGetOneRequest
	(*BlogGetOneResponse)(nil),        // 5: contents.v1.BlogGetOneResponse
	(*timestamppb.Timestamp)(nil),     // 6: google.protobuf.Timestamp
}
var file_contents_v1_blog_proto_depIdxs = []int32{
	6, // 0: contents.v1.Blog.datetime:type_name -> google.protobuf.Timestamp
	0, // 1: contents.v1.BlogGetListRequest.sort_by:type_name -> contents.v1.BlogGetListRequest.SortBy
	1, // 2: contents.v1.BlogGetListResponseStream.data:type_name -> contents.v1.Blog
	1, // 3: contents.v1.BlogGetOneResponse.data:type_name -> contents.v1.Blog
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_contents_v1_blog_proto_init() }
func file_contents_v1_blog_proto_init() {
	if File_contents_v1_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contents_v1_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_contents_v1_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogGetListRequest); i {
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
		file_contents_v1_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogGetListResponseStream); i {
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
		file_contents_v1_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogGetOneRequest); i {
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
		file_contents_v1_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogGetOneResponse); i {
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
			RawDescriptor: file_contents_v1_blog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contents_v1_blog_proto_goTypes,
		DependencyIndexes: file_contents_v1_blog_proto_depIdxs,
		EnumInfos:         file_contents_v1_blog_proto_enumTypes,
		MessageInfos:      file_contents_v1_blog_proto_msgTypes,
	}.Build()
	File_contents_v1_blog_proto = out.File
	file_contents_v1_blog_proto_rawDesc = nil
	file_contents_v1_blog_proto_goTypes = nil
	file_contents_v1_blog_proto_depIdxs = nil
}
