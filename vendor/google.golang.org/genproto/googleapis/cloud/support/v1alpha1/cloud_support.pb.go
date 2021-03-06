// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/support/v1alpha1/cloud_support.proto

/*
Package support is a generated protocol buffer package.

It is generated from these files:
	google/cloud/support/v1alpha1/cloud_support.proto

It has these top-level messages:
	GetSupportAccountRequest
	ListSupportAccountsRequest
	ListSupportAccountsResponse
	GetCaseRequest
	ListCasesRequest
	ListCasesResponse
	ListCommentsRequest
	ListCommentsResponse
	CreateCaseRequest
	UpdateCaseRequest
	CreateCommentRequest
	GetIssueTaxonomyRequest
*/
package support

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_cloud_support_common "google.golang.org/genproto/googleapis/cloud/support/common"
import _ "github.com/golang/protobuf/ptypes/empty"
import google_protobuf3 "google.golang.org/genproto/protobuf/field_mask"

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

// The request message for `GetSupportAccount`.
type GetSupportAccountRequest struct {
	// The resource name of the support accounts. For example:
	// `supportAccounts/accountA`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetSupportAccountRequest) Reset()                    { *m = GetSupportAccountRequest{} }
func (m *GetSupportAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSupportAccountRequest) ProtoMessage()               {}
func (*GetSupportAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetSupportAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for `ListSupportAccount`.
type ListSupportAccountsRequest struct {
	// The filter applied to search results. It only supports filtering a support
	// account list by a cloud_resource. For example, to filter results by support
	// accounts associated with an Organization, its value should be:
	// "cloud_resource:organizations/<organization_id>"
	Filter string `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	// Maximum number of accounts fetched with each request.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// A token identifying the page of results to return. If unspecified, the
	// first page is retrieved.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListSupportAccountsRequest) Reset()                    { *m = ListSupportAccountsRequest{} }
func (m *ListSupportAccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSupportAccountsRequest) ProtoMessage()               {}
func (*ListSupportAccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListSupportAccountsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListSupportAccountsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListSupportAccountsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for `ListSupportAccount`.
type ListSupportAccountsResponse struct {
	// A list of support accounts.
	Accounts []*google_cloud_support_common.SupportAccount `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
	// A token to retrieve the next page of results. This should be passed on in
	// `page_token` field of `ListSupportAccountRequest` for next request. If
	// unspecified, there are no more results to retrieve.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListSupportAccountsResponse) Reset()                    { *m = ListSupportAccountsResponse{} }
func (m *ListSupportAccountsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSupportAccountsResponse) ProtoMessage()               {}
func (*ListSupportAccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListSupportAccountsResponse) GetAccounts() []*google_cloud_support_common.SupportAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ListSupportAccountsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for `GetCase` method.
type GetCaseRequest struct {
	// Name of case resource requested.
	// For example: "supportAccounts/accountA/cases/123"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetCaseRequest) Reset()                    { *m = GetCaseRequest{} }
func (m *GetCaseRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCaseRequest) ProtoMessage()               {}
func (*GetCaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetCaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for `ListCase` method.
type ListCasesRequest struct {
	// Name of the account resource for which cases are requested. For example:
	// "supportAccounts/accountA"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The filter applied to the search results. Currently it only accepts "OPEN"
	// or "CLOSED" strings, filtering out cases that are open or resolved.
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// Maximum number of cases fetched with each request.
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// A token identifying the page of results to return. If unspecified, the
	// first page is retrieved.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListCasesRequest) Reset()                    { *m = ListCasesRequest{} }
func (m *ListCasesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListCasesRequest) ProtoMessage()               {}
func (*ListCasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListCasesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListCasesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListCasesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCasesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for `ListCase` method.
type ListCasesResponse struct {
	// A list of cases.
	Cases []*google_cloud_support_common.Case `protobuf:"bytes,1,rep,name=cases" json:"cases,omitempty"`
	// A token to retrieve the next page of results. This should be passed on in
	// `page_token` field of `ListCaseRequest` for next request. If unspecified,
	// there are no more results to retrieve.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListCasesResponse) Reset()                    { *m = ListCasesResponse{} }
func (m *ListCasesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCasesResponse) ProtoMessage()               {}
func (*ListCasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListCasesResponse) GetCases() []*google_cloud_support_common.Case {
	if m != nil {
		return m.Cases
	}
	return nil
}

func (m *ListCasesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for `ListComments` method.
type ListCommentsRequest struct {
	// The resource name of case for which comments should be listed.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ListCommentsRequest) Reset()                    { *m = ListCommentsRequest{} }
func (m *ListCommentsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListCommentsRequest) ProtoMessage()               {}
func (*ListCommentsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListCommentsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message for `ListComments` method.
type ListCommentsResponse struct {
	// A list of comments.
	Comments []*google_cloud_support_common.Comment `protobuf:"bytes,1,rep,name=comments" json:"comments,omitempty"`
}

func (m *ListCommentsResponse) Reset()                    { *m = ListCommentsResponse{} }
func (m *ListCommentsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCommentsResponse) ProtoMessage()               {}
func (*ListCommentsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListCommentsResponse) GetComments() []*google_cloud_support_common.Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

// The request message for `CreateCase` method.
type CreateCaseRequest struct {
	// The resource name for `SupportAccount` under which this case is created.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The case resource to create.
	Case *google_cloud_support_common.Case `protobuf:"bytes,2,opt,name=case" json:"case,omitempty"`
}

func (m *CreateCaseRequest) Reset()                    { *m = CreateCaseRequest{} }
func (m *CreateCaseRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCaseRequest) ProtoMessage()               {}
func (*CreateCaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateCaseRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateCaseRequest) GetCase() *google_cloud_support_common.Case {
	if m != nil {
		return m.Case
	}
	return nil
}

// The request message for `UpdateCase` method.
type UpdateCaseRequest struct {
	// The case resource to update.
	Case *google_cloud_support_common.Case `protobuf:"bytes,1,opt,name=case" json:"case,omitempty"`
	// A field that represents attributes of a Case object that should be updated
	// as part of this request.
	UpdateMask *google_protobuf3.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateCaseRequest) Reset()                    { *m = UpdateCaseRequest{} }
func (m *UpdateCaseRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateCaseRequest) ProtoMessage()               {}
func (*UpdateCaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateCaseRequest) GetCase() *google_cloud_support_common.Case {
	if m != nil {
		return m.Case
	}
	return nil
}

func (m *UpdateCaseRequest) GetUpdateMask() *google_protobuf3.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The request message for `CreateComment` method.
type CreateCommentRequest struct {
	// The resource name of case to which this comment should be added.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The `Comment` to be added to this case.
	Comment *google_cloud_support_common.Comment `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
}

func (m *CreateCommentRequest) Reset()                    { *m = CreateCommentRequest{} }
func (m *CreateCommentRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCommentRequest) ProtoMessage()               {}
func (*CreateCommentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreateCommentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCommentRequest) GetComment() *google_cloud_support_common.Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

// The request message for `GetIssueTaxonomy` method.
type GetIssueTaxonomyRequest struct {
}

func (m *GetIssueTaxonomyRequest) Reset()                    { *m = GetIssueTaxonomyRequest{} }
func (m *GetIssueTaxonomyRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIssueTaxonomyRequest) ProtoMessage()               {}
func (*GetIssueTaxonomyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*GetSupportAccountRequest)(nil), "google.cloud.support.v1alpha1.GetSupportAccountRequest")
	proto.RegisterType((*ListSupportAccountsRequest)(nil), "google.cloud.support.v1alpha1.ListSupportAccountsRequest")
	proto.RegisterType((*ListSupportAccountsResponse)(nil), "google.cloud.support.v1alpha1.ListSupportAccountsResponse")
	proto.RegisterType((*GetCaseRequest)(nil), "google.cloud.support.v1alpha1.GetCaseRequest")
	proto.RegisterType((*ListCasesRequest)(nil), "google.cloud.support.v1alpha1.ListCasesRequest")
	proto.RegisterType((*ListCasesResponse)(nil), "google.cloud.support.v1alpha1.ListCasesResponse")
	proto.RegisterType((*ListCommentsRequest)(nil), "google.cloud.support.v1alpha1.ListCommentsRequest")
	proto.RegisterType((*ListCommentsResponse)(nil), "google.cloud.support.v1alpha1.ListCommentsResponse")
	proto.RegisterType((*CreateCaseRequest)(nil), "google.cloud.support.v1alpha1.CreateCaseRequest")
	proto.RegisterType((*UpdateCaseRequest)(nil), "google.cloud.support.v1alpha1.UpdateCaseRequest")
	proto.RegisterType((*CreateCommentRequest)(nil), "google.cloud.support.v1alpha1.CreateCommentRequest")
	proto.RegisterType((*GetIssueTaxonomyRequest)(nil), "google.cloud.support.v1alpha1.GetIssueTaxonomyRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CloudSupport service

type CloudSupportClient interface {
	// Retrieves the support account details given an account identifier.
	// The authenticated user calling this method must be the account owner.
	GetSupportAccount(ctx context.Context, in *GetSupportAccountRequest, opts ...grpc.CallOption) (*google_cloud_support_common.SupportAccount, error)
	// Retrieves the list of accounts the current authenticated user has access
	// to.
	ListSupportAccounts(ctx context.Context, in *ListSupportAccountsRequest, opts ...grpc.CallOption) (*ListSupportAccountsResponse, error)
	// Retrieves the details for a support case. The current authenticated user
	// calling this method must have permissions to view this case.
	GetCase(ctx context.Context, in *GetCaseRequest, opts ...grpc.CallOption) (*google_cloud_support_common.Case, error)
	// Retrieves the list of support cases associated with an account. The current
	// authenticated user must have the permission to list and view these cases.
	ListCases(ctx context.Context, in *ListCasesRequest, opts ...grpc.CallOption) (*ListCasesResponse, error)
	// Lists all comments from a case.
	ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error)
	// Creates a case and associates it with a
	// [SupportAccount][google.cloud.support.v1alpha2.SupportAcccount]. The
	// authenticated user attempting this action must have permissions to create a
	// `Case` under that [SupportAccount].
	CreateCase(ctx context.Context, in *CreateCaseRequest, opts ...grpc.CallOption) (*google_cloud_support_common.Case, error)
	// Updates a support case. Only a small set of details (priority, subject and
	// cc_address) can be update after a case is created.
	UpdateCase(ctx context.Context, in *UpdateCaseRequest, opts ...grpc.CallOption) (*google_cloud_support_common.Case, error)
	// Adds a new comment to a case.
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*google_cloud_support_common.Comment, error)
	// Retrieves the taxonomy of product categories and components to be used
	// while creating a support case.
	GetIssueTaxonomy(ctx context.Context, in *GetIssueTaxonomyRequest, opts ...grpc.CallOption) (*google_cloud_support_common.IssueTaxonomy, error)
}

type cloudSupportClient struct {
	cc *grpc.ClientConn
}

func NewCloudSupportClient(cc *grpc.ClientConn) CloudSupportClient {
	return &cloudSupportClient{cc}
}

func (c *cloudSupportClient) GetSupportAccount(ctx context.Context, in *GetSupportAccountRequest, opts ...grpc.CallOption) (*google_cloud_support_common.SupportAccount, error) {
	out := new(google_cloud_support_common.SupportAccount)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/GetSupportAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSupportClient) ListSupportAccounts(ctx context.Context, in *ListSupportAccountsRequest, opts ...grpc.CallOption) (*ListSupportAccountsResponse, error) {
	out := new(ListSupportAccountsResponse)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/ListSupportAccounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSupportClient) GetCase(ctx context.Context, in *GetCaseRequest, opts ...grpc.CallOption) (*google_cloud_support_common.Case, error) {
	out := new(google_cloud_support_common.Case)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/GetCase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSupportClient) ListCases(ctx context.Context, in *ListCasesRequest, opts ...grpc.CallOption) (*ListCasesResponse, error) {
	out := new(ListCasesResponse)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/ListCases", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSupportClient) ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error) {
	out := new(ListCommentsResponse)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/ListComments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSupportClient) CreateCase(ctx context.Context, in *CreateCaseRequest, opts ...grpc.CallOption) (*google_cloud_support_common.Case, error) {
	out := new(google_cloud_support_common.Case)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/CreateCase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSupportClient) UpdateCase(ctx context.Context, in *UpdateCaseRequest, opts ...grpc.CallOption) (*google_cloud_support_common.Case, error) {
	out := new(google_cloud_support_common.Case)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/UpdateCase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSupportClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*google_cloud_support_common.Comment, error) {
	out := new(google_cloud_support_common.Comment)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/CreateComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudSupportClient) GetIssueTaxonomy(ctx context.Context, in *GetIssueTaxonomyRequest, opts ...grpc.CallOption) (*google_cloud_support_common.IssueTaxonomy, error) {
	out := new(google_cloud_support_common.IssueTaxonomy)
	err := grpc.Invoke(ctx, "/google.cloud.support.v1alpha1.CloudSupport/GetIssueTaxonomy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudSupport service

type CloudSupportServer interface {
	// Retrieves the support account details given an account identifier.
	// The authenticated user calling this method must be the account owner.
	GetSupportAccount(context.Context, *GetSupportAccountRequest) (*google_cloud_support_common.SupportAccount, error)
	// Retrieves the list of accounts the current authenticated user has access
	// to.
	ListSupportAccounts(context.Context, *ListSupportAccountsRequest) (*ListSupportAccountsResponse, error)
	// Retrieves the details for a support case. The current authenticated user
	// calling this method must have permissions to view this case.
	GetCase(context.Context, *GetCaseRequest) (*google_cloud_support_common.Case, error)
	// Retrieves the list of support cases associated with an account. The current
	// authenticated user must have the permission to list and view these cases.
	ListCases(context.Context, *ListCasesRequest) (*ListCasesResponse, error)
	// Lists all comments from a case.
	ListComments(context.Context, *ListCommentsRequest) (*ListCommentsResponse, error)
	// Creates a case and associates it with a
	// [SupportAccount][google.cloud.support.v1alpha2.SupportAcccount]. The
	// authenticated user attempting this action must have permissions to create a
	// `Case` under that [SupportAccount].
	CreateCase(context.Context, *CreateCaseRequest) (*google_cloud_support_common.Case, error)
	// Updates a support case. Only a small set of details (priority, subject and
	// cc_address) can be update after a case is created.
	UpdateCase(context.Context, *UpdateCaseRequest) (*google_cloud_support_common.Case, error)
	// Adds a new comment to a case.
	CreateComment(context.Context, *CreateCommentRequest) (*google_cloud_support_common.Comment, error)
	// Retrieves the taxonomy of product categories and components to be used
	// while creating a support case.
	GetIssueTaxonomy(context.Context, *GetIssueTaxonomyRequest) (*google_cloud_support_common.IssueTaxonomy, error)
}

func RegisterCloudSupportServer(s *grpc.Server, srv CloudSupportServer) {
	s.RegisterService(&_CloudSupport_serviceDesc, srv)
}

func _CloudSupport_GetSupportAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupportAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).GetSupportAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/GetSupportAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).GetSupportAccount(ctx, req.(*GetSupportAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSupport_ListSupportAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSupportAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).ListSupportAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/ListSupportAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).ListSupportAccounts(ctx, req.(*ListSupportAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSupport_GetCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).GetCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/GetCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).GetCase(ctx, req.(*GetCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSupport_ListCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).ListCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/ListCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).ListCases(ctx, req.(*ListCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSupport_ListComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).ListComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/ListComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).ListComments(ctx, req.(*ListCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSupport_CreateCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).CreateCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/CreateCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).CreateCase(ctx, req.(*CreateCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSupport_UpdateCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).UpdateCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/UpdateCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).UpdateCase(ctx, req.(*UpdateCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSupport_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudSupport_GetIssueTaxonomy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIssueTaxonomyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudSupportServer).GetIssueTaxonomy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v1alpha1.CloudSupport/GetIssueTaxonomy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudSupportServer).GetIssueTaxonomy(ctx, req.(*GetIssueTaxonomyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.support.v1alpha1.CloudSupport",
	HandlerType: (*CloudSupportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSupportAccount",
			Handler:    _CloudSupport_GetSupportAccount_Handler,
		},
		{
			MethodName: "ListSupportAccounts",
			Handler:    _CloudSupport_ListSupportAccounts_Handler,
		},
		{
			MethodName: "GetCase",
			Handler:    _CloudSupport_GetCase_Handler,
		},
		{
			MethodName: "ListCases",
			Handler:    _CloudSupport_ListCases_Handler,
		},
		{
			MethodName: "ListComments",
			Handler:    _CloudSupport_ListComments_Handler,
		},
		{
			MethodName: "CreateCase",
			Handler:    _CloudSupport_CreateCase_Handler,
		},
		{
			MethodName: "UpdateCase",
			Handler:    _CloudSupport_UpdateCase_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _CloudSupport_CreateComment_Handler,
		},
		{
			MethodName: "GetIssueTaxonomy",
			Handler:    _CloudSupport_GetIssueTaxonomy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/support/v1alpha1/cloud_support.proto",
}

func init() { proto.RegisterFile("google/cloud/support/v1alpha1/cloud_support.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x41, 0x4f, 0x33, 0x45,
	0x18, 0xce, 0xb4, 0xc8, 0x07, 0x2f, 0xdf, 0xa7, 0x5f, 0x47, 0x82, 0x65, 0x0b, 0x49, 0x3b, 0x21,
	0xa6, 0x56, 0xdd, 0x85, 0x36, 0x88, 0x96, 0x40, 0x14, 0x88, 0x8d, 0x89, 0x26, 0xa4, 0x60, 0x62,
	0xbc, 0x34, 0x43, 0x19, 0xd6, 0x95, 0xee, 0xce, 0xda, 0x99, 0x1a, 0x40, 0xbd, 0x78, 0xf1, 0xa6,
	0x07, 0x6f, 0x7a, 0xe1, 0xe2, 0x59, 0x0f, 0xfe, 0x13, 0xff, 0x82, 0xfe, 0x0f, 0xb3, 0xb3, 0xb3,
	0xed, 0x76, 0x69, 0x77, 0x17, 0x6e, 0xdd, 0x77, 0xde, 0xe7, 0x7d, 0x9f, 0x79, 0xe6, 0x9d, 0x67,
	0x0a, 0x3b, 0x36, 0xe7, 0xf6, 0x80, 0x59, 0xfd, 0x01, 0x1f, 0x5d, 0x5a, 0x62, 0xe4, 0xfb, 0x7c,
	0x28, 0xad, 0x6f, 0x77, 0xe8, 0xc0, 0xff, 0x8a, 0xee, 0x84, 0xe1, 0x9e, 0x0e, 0x9b, 0xfe, 0x90,
	0x4b, 0x8e, 0x37, 0x43, 0x88, 0xa9, 0xd6, 0xcc, 0x68, 0x2d, 0x82, 0x18, 0x1b, 0xba, 0x22, 0xf5,
	0x1d, 0x8b, 0x7a, 0x1e, 0x97, 0x54, 0x3a, 0xdc, 0x13, 0x21, 0xd8, 0xa8, 0xcd, 0xec, 0xd7, 0xe7,
	0xae, 0xcb, 0x3d, 0x9d, 0x52, 0xd1, 0x29, 0xea, 0xeb, 0x62, 0x74, 0x65, 0x31, 0xd7, 0x97, 0xb7,
	0x7a, 0xb1, 0x9a, 0x5c, 0xbc, 0x72, 0xd8, 0xe0, 0xb2, 0xe7, 0x52, 0x71, 0x1d, 0x66, 0x10, 0x13,
	0xca, 0x1d, 0x26, 0xcf, 0xc2, 0xca, 0x1f, 0xf5, 0xfb, 0x7c, 0xe4, 0xc9, 0x2e, 0xfb, 0x66, 0xc4,
	0x84, 0xc4, 0x18, 0x16, 0x3c, 0xea, 0xb2, 0x32, 0xaa, 0xa2, 0xfa, 0x72, 0x57, 0xfd, 0x26, 0x3e,
	0x18, 0x9f, 0x3a, 0x22, 0x01, 0x10, 0x11, 0x62, 0x0d, 0x16, 0xaf, 0x9c, 0x81, 0x64, 0x43, 0x8d,
	0xd1, 0x5f, 0xb8, 0x02, 0xcb, 0x3e, 0xb5, 0x59, 0x4f, 0x38, 0x77, 0xac, 0x5c, 0xa8, 0xa2, 0x7a,
	0xb1, 0xbb, 0x14, 0x04, 0xce, 0x9c, 0x3b, 0x86, 0x37, 0x01, 0xd4, 0xa2, 0xe4, 0xd7, 0xcc, 0x2b,
	0x17, 0x15, 0x50, 0xa5, 0x9f, 0x07, 0x01, 0xf2, 0x33, 0x82, 0xca, 0xcc, 0x96, 0xc2, 0xe7, 0x9e,
	0x60, 0xb8, 0x03, 0x4b, 0x54, 0xc7, 0xca, 0xa8, 0x5a, 0xac, 0xaf, 0x34, 0xdf, 0x36, 0x67, 0x6a,
	0xae, 0x65, 0x4b, 0xec, 0x75, 0x0c, 0xc6, 0x6f, 0xc2, 0x6b, 0x1e, 0xbb, 0x91, 0xbd, 0x18, 0x99,
	0x82, 0x22, 0xf3, 0x22, 0x08, 0x9f, 0x8e, 0x09, 0x6d, 0xc1, 0xab, 0x1d, 0x26, 0x8f, 0xa9, 0x60,
	0x69, 0x42, 0xdd, 0xc1, 0xcb, 0x80, 0x75, 0x90, 0x26, 0x52, 0xf2, 0x62, 0x92, 0x15, 0xe6, 0x4b,
	0x56, 0x4c, 0x95, 0x6c, 0x21, 0x29, 0x99, 0x84, 0x52, 0xac, 0xb7, 0xd6, 0x69, 0x0f, 0x5e, 0xe9,
	0x07, 0x01, 0x2d, 0x52, 0x2d, 0x55, 0x24, 0xb5, 0xbb, 0x30, 0x3f, 0xb7, 0x2e, 0x6f, 0xc1, 0xeb,
	0xaa, 0x2b, 0x77, 0x5d, 0x16, 0x9b, 0x89, 0x59, 0xe2, 0x7c, 0x01, 0xab, 0xd3, 0xa9, 0x9a, 0xe3,
	0x87, 0xb0, 0xd4, 0xd7, 0x31, 0x4d, 0x73, 0x2b, 0x9d, 0x66, 0x98, 0xdc, 0x1d, 0xa3, 0xc8, 0x05,
	0x94, 0x8e, 0x87, 0x8c, 0x4a, 0x16, 0x3f, 0x9f, 0x35, 0x58, 0xf4, 0xe9, 0x90, 0x79, 0x32, 0x1a,
	0xcb, 0xf0, 0x0b, 0xef, 0xc2, 0x42, 0xb0, 0x45, 0xb5, 0x9d, 0x5c, 0x8a, 0xa8, 0x74, 0xf2, 0x13,
	0x82, 0xd2, 0xe7, 0xfe, 0x65, 0xa2, 0x49, 0x54, 0x0c, 0x3d, 0xaa, 0x18, 0xde, 0x87, 0x95, 0x91,
	0xaa, 0xa5, 0x6e, 0xa5, 0xa6, 0x62, 0x44, 0xe8, 0xe8, 0xe2, 0x9a, 0x1f, 0x07, 0x17, 0xf7, 0x33,
	0x2a, 0xae, 0xbb, 0x10, 0xa6, 0x07, 0xbf, 0xc9, 0xd7, 0xb0, 0xaa, 0x77, 0xab, 0x85, 0x48, 0x19,
	0xb4, 0x43, 0x78, 0xa6, 0x55, 0xd2, 0x4d, 0xf2, 0x49, 0x1b, 0x81, 0xc8, 0x3a, 0xbc, 0xd1, 0x61,
	0xf2, 0x13, 0x21, 0x46, 0xec, 0x9c, 0xde, 0x70, 0x8f, 0xbb, 0xb7, 0xba, 0x5d, 0xf3, 0xbf, 0x15,
	0x78, 0x7e, 0x1c, 0x14, 0xd1, 0x77, 0x0b, 0xff, 0x81, 0xa0, 0xf4, 0xc0, 0x56, 0xf0, 0x9e, 0x99,
	0xea, 0x85, 0xe6, 0x3c, 0x23, 0x32, 0x1e, 0x73, 0xa1, 0x49, 0xe3, 0xc7, 0x7f, 0xfe, 0xfd, 0xb5,
	0xb0, 0x85, 0xc9, 0xc4, 0x97, 0xbf, 0x0b, 0x04, 0x38, 0x10, 0xd3, 0x06, 0x62, 0x35, 0x7e, 0xc0,
	0x7f, 0xa2, 0x70, 0x66, 0x13, 0xde, 0x82, 0x3f, 0xc8, 0x60, 0x3a, 0xdf, 0x02, 0x8d, 0xf6, 0x53,
	0xa0, 0xe1, 0xf8, 0x93, 0x9a, 0xa2, 0x5e, 0xc1, 0xeb, 0x13, 0xea, 0x09, 0xd2, 0xf8, 0x17, 0x04,
	0xcf, 0xb4, 0xfb, 0xe0, 0x77, 0xb3, 0xf5, 0x8c, 0x0d, 0xa8, 0x91, 0x3d, 0x92, 0xa4, 0xa9, 0x08,
	0xbc, 0x83, 0x1b, 0x99, 0xda, 0x59, 0xca, 0x1b, 0x02, 0x0d, 0xef, 0x11, 0x2c, 0x8f, 0xdd, 0x06,
	0x5b, 0x39, 0xb6, 0x1f, 0xf7, 0x44, 0x63, 0x3b, 0x3f, 0x40, 0xab, 0xb4, 0xad, 0x48, 0x36, 0x70,
	0x3d, 0xfb, 0x80, 0x43, 0x96, 0xf8, 0x6f, 0x04, 0xcf, 0xe3, 0x7e, 0x83, 0x9b, 0x79, 0x9a, 0x4e,
	0xfb, 0x98, 0xd1, 0x7a, 0x14, 0x46, 0x73, 0xdd, 0x57, 0x5c, 0x77, 0x71, 0x2b, 0xbf, 0xa0, 0x56,
	0xe4, 0x65, 0xf8, 0x77, 0x04, 0x30, 0x31, 0x33, 0x9c, 0xa5, 0xd4, 0x03, 0xdf, 0xcb, 0x73, 0xe2,
	0xef, 0x2b, 0x82, 0x4d, 0x12, 0x3f, 0xf1, 0xd0, 0x1d, 0xe7, 0xca, 0xd9, 0x0e, 0x8d, 0xeb, 0x1e,
	0x01, 0x4c, 0x5c, 0x30, 0x93, 0xdd, 0x03, 0xc3, 0xcc, 0xc3, 0xee, 0x40, 0xb1, 0xdb, 0x6b, 0x5a,
	0x31, 0x76, 0x41, 0x73, 0x33, 0x43, 0x43, 0x4d, 0xf1, 0x2f, 0x04, 0x2f, 0xa6, 0xfc, 0x11, 0xb7,
	0xf2, 0x69, 0x38, 0xe5, 0xa6, 0x46, 0x2e, 0xa3, 0x24, 0x27, 0x8a, 0xeb, 0x21, 0x79, 0xca, 0x51,
	0xb7, 0x23, 0x97, 0xc5, 0xbf, 0x21, 0x78, 0x99, 0xb4, 0x59, 0xfc, 0x5e, 0xf6, 0x45, 0x9f, 0xe5,
	0xcb, 0x46, 0x23, 0x95, 0xf8, 0x14, 0x84, 0x10, 0x45, 0x7f, 0x03, 0x1b, 0x63, 0xfa, 0x6d, 0x3b,
	0x51, 0xf6, 0xe8, 0x7b, 0xa8, 0xf5, 0xb9, 0x9b, 0x4e, 0xe6, 0xa8, 0x14, 0x7f, 0x09, 0x4e, 0x83,
	0xf7, 0xeb, 0xcb, 0x13, 0x8d, 0xb0, 0xf9, 0x80, 0x7a, 0xb6, 0xc9, 0x87, 0xb6, 0x65, 0x33, 0x4f,
	0xbd, 0x6d, 0x56, 0xb8, 0x44, 0x7d, 0x47, 0xcc, 0xf9, 0x57, 0xbd, 0xaf, 0x03, 0x17, 0x8b, 0x0a,
	0xd0, 0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xad, 0xe1, 0xf2, 0x57, 0x85, 0x0b, 0x00, 0x00,
}
