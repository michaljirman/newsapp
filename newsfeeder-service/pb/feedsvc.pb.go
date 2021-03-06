// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feedsvc.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateFeedRequest struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Provider             string   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	RequestUri           string   `protobuf:"bytes,4,opt,name=request_uri,json=requestUri,proto3" json:"request_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFeedRequest) Reset()         { *m = CreateFeedRequest{} }
func (m *CreateFeedRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFeedRequest) ProtoMessage()    {}
func (*CreateFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{0}
}
func (m *CreateFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFeedRequest.Unmarshal(m, b)
}
func (m *CreateFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFeedRequest.Marshal(b, m, deterministic)
}
func (dst *CreateFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFeedRequest.Merge(dst, src)
}
func (m *CreateFeedRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFeedRequest.Size(m)
}
func (m *CreateFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFeedRequest proto.InternalMessageInfo

func (m *CreateFeedRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *CreateFeedRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *CreateFeedRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateFeedRequest) GetRequestUri() string {
	if m != nil {
		return m.RequestUri
	}
	return ""
}

type CreateFeedReply struct {
	FeedId               uint64   `protobuf:"varint,1,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	ResourceUri          string   `protobuf:"bytes,2,opt,name=resource_uri,json=resourceUri,proto3" json:"resource_uri,omitempty"`
	Err                  string   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFeedReply) Reset()         { *m = CreateFeedReply{} }
func (m *CreateFeedReply) String() string { return proto.CompactTextString(m) }
func (*CreateFeedReply) ProtoMessage()    {}
func (*CreateFeedReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{1}
}
func (m *CreateFeedReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFeedReply.Unmarshal(m, b)
}
func (m *CreateFeedReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFeedReply.Marshal(b, m, deterministic)
}
func (dst *CreateFeedReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFeedReply.Merge(dst, src)
}
func (m *CreateFeedReply) XXX_Size() int {
	return xxx_messageInfo_CreateFeedReply.Size(m)
}
func (m *CreateFeedReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFeedReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFeedReply proto.InternalMessageInfo

func (m *CreateFeedReply) GetFeedId() uint64 {
	if m != nil {
		return m.FeedId
	}
	return 0
}

func (m *CreateFeedReply) GetResourceUri() string {
	if m != nil {
		return m.ResourceUri
	}
	return ""
}

func (m *CreateFeedReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetArticlesRequest struct {
	FeedId               uint64   `protobuf:"varint,1,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArticlesRequest) Reset()         { *m = GetArticlesRequest{} }
func (m *GetArticlesRequest) String() string { return proto.CompactTextString(m) }
func (*GetArticlesRequest) ProtoMessage()    {}
func (*GetArticlesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{2}
}
func (m *GetArticlesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticlesRequest.Unmarshal(m, b)
}
func (m *GetArticlesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticlesRequest.Marshal(b, m, deterministic)
}
func (dst *GetArticlesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticlesRequest.Merge(dst, src)
}
func (m *GetArticlesRequest) XXX_Size() int {
	return xxx_messageInfo_GetArticlesRequest.Size(m)
}
func (m *GetArticlesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticlesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticlesRequest proto.InternalMessageInfo

func (m *GetArticlesRequest) GetFeedId() uint64 {
	if m != nil {
		return m.FeedId
	}
	return 0
}

type Article struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Link                 string               `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Published            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=published,proto3" json:"published,omitempty"`
	Guid                 string               `protobuf:"bytes,5,opt,name=guid,proto3" json:"guid,omitempty"`
	ThumbnailImageUrl    string               `protobuf:"bytes,6,opt,name=thumbnail_image_url,json=thumbnailImageUrl,proto3" json:"thumbnail_image_url,omitempty"`
	HtmlContent          string               `protobuf:"bytes,7,opt,name=html_content,json=htmlContent,proto3" json:"html_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{3}
}
func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (dst *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(dst, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Article) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Article) GetPublished() *timestamp.Timestamp {
	if m != nil {
		return m.Published
	}
	return nil
}

func (m *Article) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *Article) GetThumbnailImageUrl() string {
	if m != nil {
		return m.ThumbnailImageUrl
	}
	return ""
}

func (m *Article) GetHtmlContent() string {
	if m != nil {
		return m.HtmlContent
	}
	return ""
}

type GetArticlesReply struct {
	Articles             []*Article `protobuf:"bytes,1,rep,name=Articles,proto3" json:"Articles,omitempty"`
	Err                  string     `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetArticlesReply) Reset()         { *m = GetArticlesReply{} }
func (m *GetArticlesReply) String() string { return proto.CompactTextString(m) }
func (*GetArticlesReply) ProtoMessage()    {}
func (*GetArticlesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{4}
}
func (m *GetArticlesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticlesReply.Unmarshal(m, b)
}
func (m *GetArticlesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticlesReply.Marshal(b, m, deterministic)
}
func (dst *GetArticlesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticlesReply.Merge(dst, src)
}
func (m *GetArticlesReply) XXX_Size() int {
	return xxx_messageInfo_GetArticlesReply.Size(m)
}
func (m *GetArticlesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticlesReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticlesReply proto.InternalMessageInfo

func (m *GetArticlesReply) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

func (m *GetArticlesReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetArticleRequest struct {
	FeedId               uint64   `protobuf:"varint,1,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	ArticleGuid          string   `protobuf:"bytes,2,opt,name=article_guid,json=articleGuid,proto3" json:"article_guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArticleRequest) Reset()         { *m = GetArticleRequest{} }
func (m *GetArticleRequest) String() string { return proto.CompactTextString(m) }
func (*GetArticleRequest) ProtoMessage()    {}
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{5}
}
func (m *GetArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticleRequest.Unmarshal(m, b)
}
func (m *GetArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticleRequest.Marshal(b, m, deterministic)
}
func (dst *GetArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticleRequest.Merge(dst, src)
}
func (m *GetArticleRequest) XXX_Size() int {
	return xxx_messageInfo_GetArticleRequest.Size(m)
}
func (m *GetArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticleRequest proto.InternalMessageInfo

func (m *GetArticleRequest) GetFeedId() uint64 {
	if m != nil {
		return m.FeedId
	}
	return 0
}

func (m *GetArticleRequest) GetArticleGuid() string {
	if m != nil {
		return m.ArticleGuid
	}
	return ""
}

type GetArticleReply struct {
	Article              *Article `protobuf:"bytes,1,opt,name=Article,proto3" json:"Article,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArticleReply) Reset()         { *m = GetArticleReply{} }
func (m *GetArticleReply) String() string { return proto.CompactTextString(m) }
func (*GetArticleReply) ProtoMessage()    {}
func (*GetArticleReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{6}
}
func (m *GetArticleReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticleReply.Unmarshal(m, b)
}
func (m *GetArticleReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticleReply.Marshal(b, m, deterministic)
}
func (dst *GetArticleReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticleReply.Merge(dst, src)
}
func (m *GetArticleReply) XXX_Size() int {
	return xxx_messageInfo_GetArticleReply.Size(m)
}
func (m *GetArticleReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticleReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticleReply proto.InternalMessageInfo

func (m *GetArticleReply) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

func (m *GetArticleReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetFeedsRequest struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Provider             string   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedsRequest) Reset()         { *m = GetFeedsRequest{} }
func (m *GetFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedsRequest) ProtoMessage()    {}
func (*GetFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{7}
}
func (m *GetFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedsRequest.Unmarshal(m, b)
}
func (m *GetFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedsRequest.Marshal(b, m, deterministic)
}
func (dst *GetFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedsRequest.Merge(dst, src)
}
func (m *GetFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedsRequest.Size(m)
}
func (m *GetFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedsRequest proto.InternalMessageInfo

func (m *GetFeedsRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetFeedsRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

type Feed struct {
	FeedId               uint64               `protobuf:"varint,1,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Category             string               `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Provider             string               `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	Url                  string               `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Feed) Reset()         { *m = Feed{} }
func (m *Feed) String() string { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()    {}
func (*Feed) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{8}
}
func (m *Feed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feed.Unmarshal(m, b)
}
func (m *Feed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feed.Marshal(b, m, deterministic)
}
func (dst *Feed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed.Merge(dst, src)
}
func (m *Feed) XXX_Size() int {
	return xxx_messageInfo_Feed.Size(m)
}
func (m *Feed) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed.DiscardUnknown(m)
}

var xxx_messageInfo_Feed proto.InternalMessageInfo

func (m *Feed) GetFeedId() uint64 {
	if m != nil {
		return m.FeedId
	}
	return 0
}

func (m *Feed) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Feed) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Feed) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Feed) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Feed) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetFeedsReply struct {
	Feeds                []*Feed  `protobuf:"bytes,1,rep,name=Feeds,proto3" json:"Feeds,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedsReply) Reset()         { *m = GetFeedsReply{} }
func (m *GetFeedsReply) String() string { return proto.CompactTextString(m) }
func (*GetFeedsReply) ProtoMessage()    {}
func (*GetFeedsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_feedsvc_bb0db0a760f605f5, []int{9}
}
func (m *GetFeedsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedsReply.Unmarshal(m, b)
}
func (m *GetFeedsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedsReply.Marshal(b, m, deterministic)
}
func (dst *GetFeedsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedsReply.Merge(dst, src)
}
func (m *GetFeedsReply) XXX_Size() int {
	return xxx_messageInfo_GetFeedsReply.Size(m)
}
func (m *GetFeedsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedsReply proto.InternalMessageInfo

func (m *GetFeedsReply) GetFeeds() []*Feed {
	if m != nil {
		return m.Feeds
	}
	return nil
}

func (m *GetFeedsReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateFeedRequest)(nil), "pb.CreateFeedRequest")
	proto.RegisterType((*CreateFeedReply)(nil), "pb.CreateFeedReply")
	proto.RegisterType((*GetArticlesRequest)(nil), "pb.GetArticlesRequest")
	proto.RegisterType((*Article)(nil), "pb.Article")
	proto.RegisterType((*GetArticlesReply)(nil), "pb.GetArticlesReply")
	proto.RegisterType((*GetArticleRequest)(nil), "pb.GetArticleRequest")
	proto.RegisterType((*GetArticleReply)(nil), "pb.GetArticleReply")
	proto.RegisterType((*GetFeedsRequest)(nil), "pb.GetFeedsRequest")
	proto.RegisterType((*Feed)(nil), "pb.Feed")
	proto.RegisterType((*GetFeedsReply)(nil), "pb.GetFeedsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeederClient is the client API for Feeder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeederClient interface {
	CreateFeed(ctx context.Context, in *CreateFeedRequest, opts ...grpc.CallOption) (*CreateFeedReply, error)
	GetFeeds(ctx context.Context, in *GetFeedsRequest, opts ...grpc.CallOption) (*GetFeedsReply, error)
	GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*GetArticlesReply, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error)
}

type feederClient struct {
	cc *grpc.ClientConn
}

func NewFeederClient(cc *grpc.ClientConn) FeederClient {
	return &feederClient{cc}
}

func (c *feederClient) CreateFeed(ctx context.Context, in *CreateFeedRequest, opts ...grpc.CallOption) (*CreateFeedReply, error) {
	out := new(CreateFeedReply)
	err := c.cc.Invoke(ctx, "/pb.Feeder/CreateFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feederClient) GetFeeds(ctx context.Context, in *GetFeedsRequest, opts ...grpc.CallOption) (*GetFeedsReply, error) {
	out := new(GetFeedsReply)
	err := c.cc.Invoke(ctx, "/pb.Feeder/GetFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feederClient) GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*GetArticlesReply, error) {
	out := new(GetArticlesReply)
	err := c.cc.Invoke(ctx, "/pb.Feeder/GetArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feederClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error) {
	out := new(GetArticleReply)
	err := c.cc.Invoke(ctx, "/pb.Feeder/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeederServer is the server API for Feeder service.
type FeederServer interface {
	CreateFeed(context.Context, *CreateFeedRequest) (*CreateFeedReply, error)
	GetFeeds(context.Context, *GetFeedsRequest) (*GetFeedsReply, error)
	GetArticles(context.Context, *GetArticlesRequest) (*GetArticlesReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error)
}

func RegisterFeederServer(s *grpc.Server, srv FeederServer) {
	s.RegisterService(&_Feeder_serviceDesc, srv)
}

func _Feeder_CreateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeederServer).CreateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Feeder/CreateFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeederServer).CreateFeed(ctx, req.(*CreateFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feeder_GetFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeederServer).GetFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Feeder/GetFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeederServer).GetFeeds(ctx, req.(*GetFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feeder_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeederServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Feeder/GetArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeederServer).GetArticles(ctx, req.(*GetArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feeder_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeederServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Feeder/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeederServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Feeder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Feeder",
	HandlerType: (*FeederServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeed",
			Handler:    _Feeder_CreateFeed_Handler,
		},
		{
			MethodName: "GetFeeds",
			Handler:    _Feeder_GetFeeds_Handler,
		},
		{
			MethodName: "GetArticles",
			Handler:    _Feeder_GetArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _Feeder_GetArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedsvc.proto",
}

func init() { proto.RegisterFile("feedsvc.proto", fileDescriptor_feedsvc_bb0db0a760f605f5) }

var fileDescriptor_feedsvc_bb0db0a760f605f5 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xc6, 0xcd, 0xa3, 0xe9, 0x98, 0xaa, 0xcd, 0xb6, 0x80, 0xe5, 0x03, 0x0d, 0x96, 0x10, 0xbd,
	0xe0, 0x4a, 0x81, 0x03, 0x54, 0xe2, 0x10, 0x55, 0xa2, 0x0a, 0x12, 0x42, 0x8a, 0xc8, 0xd9, 0xf2,
	0x63, 0x9a, 0xac, 0xd8, 0xc4, 0x66, 0xbd, 0xae, 0x94, 0x0b, 0x7f, 0x95, 0x5f, 0x01, 0x67, 0xb4,
	0x2f, 0x3b, 0x0f, 0x4a, 0x0e, 0xdc, 0x76, 0xbe, 0x79, 0xec, 0x37, 0xdf, 0xcc, 0x2e, 0x1c, 0xdf,
	0x21, 0x66, 0xe5, 0x7d, 0x1a, 0x16, 0x3c, 0x17, 0x39, 0x39, 0x28, 0x12, 0xff, 0x62, 0x96, 0xe7,
	0x33, 0x86, 0x57, 0x0a, 0x49, 0xaa, 0xbb, 0x2b, 0x41, 0x17, 0x58, 0x8a, 0x78, 0x51, 0xe8, 0xa0,
	0xe0, 0x07, 0xf4, 0x6f, 0x38, 0xc6, 0x02, 0x3f, 0x22, 0x66, 0x13, 0xfc, 0x5e, 0x61, 0x29, 0x88,
	0x0f, 0xbd, 0x34, 0x16, 0x38, 0xcb, 0xf9, 0xca, 0x73, 0x06, 0xce, 0xe5, 0xd1, 0xa4, 0xb6, 0xa5,
	0xaf, 0xe0, 0xf9, 0x3d, 0xcd, 0x90, 0x7b, 0x07, 0xda, 0x67, 0x6d, 0x72, 0x0a, 0xad, 0x8a, 0x33,
	0xaf, 0xa5, 0x60, 0x79, 0x24, 0x17, 0xe0, 0x72, 0x5d, 0x34, 0xaa, 0x38, 0xf5, 0xda, 0xca, 0x03,
	0x06, 0x9a, 0x72, 0x1a, 0x44, 0x70, 0xb2, 0x7e, 0x7f, 0xc1, 0x56, 0xe4, 0x19, 0x1c, 0xca, 0x46,
	0x22, 0x9a, 0xa9, 0xcb, 0xdb, 0x93, 0xae, 0x34, 0xc7, 0x19, 0x79, 0x01, 0x8f, 0x39, 0x96, 0x79,
	0xc5, 0x53, 0x54, 0xd5, 0xf4, 0xf5, 0xae, 0xc5, 0xa6, 0x9c, 0x4a, 0x06, 0xc8, 0xb9, 0x65, 0x80,
	0x9c, 0x07, 0xaf, 0x81, 0xdc, 0xa2, 0x18, 0x71, 0x41, 0x53, 0x86, 0xa5, 0xed, 0xf0, 0xa1, 0x3b,
	0x82, 0xdf, 0x0e, 0x1c, 0x9a, 0x60, 0x72, 0x0e, 0x1d, 0x41, 0x05, 0x43, 0xa3, 0x81, 0x36, 0xc8,
	0x00, 0xdc, 0x0c, 0xcb, 0x94, 0xd3, 0x42, 0xd0, 0x7c, 0x69, 0x49, 0xac, 0x41, 0x84, 0x40, 0x9b,
	0xd1, 0xe5, 0x37, 0xc3, 0x42, 0x9d, 0xc9, 0x3b, 0x38, 0x2a, 0xaa, 0x84, 0xd1, 0x72, 0x8e, 0x99,
	0x92, 0xc1, 0x1d, 0xfa, 0xa1, 0x1e, 0x4e, 0x68, 0x87, 0x13, 0x7e, 0xb5, 0xc3, 0x99, 0x34, 0xc1,
	0xb2, 0xda, 0xac, 0xa2, 0x99, 0xd7, 0xd1, 0xd5, 0xe4, 0x99, 0x84, 0x70, 0x26, 0xe6, 0xd5, 0x22,
	0x59, 0xc6, 0x94, 0x45, 0x74, 0x11, 0xcf, 0xa4, 0x20, 0xcc, 0xeb, 0xaa, 0x90, 0x7e, 0xed, 0x1a,
	0x4b, 0xcf, 0x94, 0x33, 0xa9, 0xdc, 0x5c, 0x2c, 0x58, 0x94, 0xe6, 0x4b, 0x81, 0x4b, 0xe1, 0x1d,
	0x6a, 0xd2, 0x12, 0xbb, 0xd1, 0x50, 0xf0, 0x19, 0x4e, 0x37, 0x74, 0x92, 0x93, 0x78, 0x05, 0x3d,
	0x0b, 0x78, 0xce, 0xa0, 0x75, 0xe9, 0x0e, 0xdd, 0xb0, 0x48, 0x42, 0x83, 0x4d, 0x6a, 0xa7, 0x95,
	0xfd, 0xa0, 0x91, 0xfd, 0x0b, 0xf4, 0x9b, 0x72, 0xfb, 0x54, 0x97, 0xfc, 0x62, 0x1d, 0x1a, 0xa9,
	0x5e, 0x8d, 0xa8, 0x06, 0xbb, 0xad, 0x68, 0x16, 0x7c, 0x82, 0x93, 0xf5, 0x82, 0x92, 0xde, 0xcb,
	0x7a, 0x54, 0xaa, 0xdc, 0x16, 0xbb, 0x7a, 0x8c, 0xbb, 0xe4, 0xc6, 0xaa, 0x96, 0xdc, 0xb8, 0xf2,
	0x3f, 0x57, 0x3e, 0xf8, 0xe9, 0x40, 0x5b, 0x16, 0x7a, 0xb8, 0xb7, 0xf7, 0x00, 0xa9, 0xda, 0xf0,
	0x2c, 0x8a, 0x85, 0xca, 0xdf, 0x33, 0x7a, 0x13, 0x3d, 0x12, 0x32, 0xb5, 0x2a, 0x32, 0x9b, 0xda,
	0xda, 0x9f, 0x6a, 0xa2, 0x47, 0x9b, 0xfd, 0xb4, 0xff, 0xd1, 0x4f, 0xe7, 0xef, 0x4f, 0xb8, 0x5b,
	0x3f, 0xe1, 0x60, 0x04, 0xc7, 0x8d, 0x58, 0x52, 0xf6, 0xe7, 0xd0, 0x51, 0x96, 0x59, 0x89, 0x9e,
	0x14, 0x5d, 0xbd, 0x5e, 0x0d, 0xef, 0xea, 0x3d, 0xfc, 0xe5, 0x40, 0x57, 0xfa, 0x90, 0x93, 0x6b,
	0x80, 0xe6, 0xbd, 0x93, 0x27, 0x32, 0x77, 0xe7, 0xff, 0xf1, 0xcf, 0xb6, 0xe1, 0x82, 0xad, 0x82,
	0x47, 0xe4, 0x2d, 0xf4, 0x2c, 0x13, 0xa2, 0x42, 0xb6, 0x86, 0xe8, 0xf7, 0x37, 0x41, 0x9d, 0xf5,
	0x01, 0xdc, 0xb5, 0xc5, 0x26, 0x4f, 0x4d, 0xcc, 0xd6, 0x8f, 0xe0, 0x9f, 0xef, 0xe0, 0x3a, 0xfd,
	0x1a, 0xa0, 0x41, 0x35, 0xe1, 0x9d, 0xc5, 0xf6, 0xcf, 0xb6, 0x61, 0x95, 0x9b, 0x74, 0xd5, 0x8c,
	0xde, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x6b, 0x93, 0x7d, 0x99, 0x05, 0x00, 0x00,
}
