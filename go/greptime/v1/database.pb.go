// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/database.proto

package v1

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

type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The `catalog` that is selected to be used in this request.
	Catalog string `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
	// The `schema` that is selected to be used in this request.
	Schema string `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	// The `authorization` header string, much like http's authorization header.
	Authorization string `protobuf:"bytes,3,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_greptime_v1_database_proto_rawDescGZIP(), []int{0}
}

func (x *RequestHeader) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

func (x *RequestHeader) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *RequestHeader) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

type GreptimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Types that are assignable to Request:
	//
	//	*GreptimeRequest_Insert
	//	*GreptimeRequest_Query
	//	*GreptimeRequest_Ddl
	Request isGreptimeRequest_Request `protobuf_oneof:"request"`
}

func (x *GreptimeRequest) Reset() {
	*x = GreptimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreptimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreptimeRequest) ProtoMessage() {}

func (x *GreptimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreptimeRequest.ProtoReflect.Descriptor instead.
func (*GreptimeRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_database_proto_rawDescGZIP(), []int{1}
}

func (x *GreptimeRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (m *GreptimeRequest) GetRequest() isGreptimeRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *GreptimeRequest) GetInsert() *InsertRequest {
	if x, ok := x.GetRequest().(*GreptimeRequest_Insert); ok {
		return x.Insert
	}
	return nil
}

func (x *GreptimeRequest) GetQuery() *QueryRequest {
	if x, ok := x.GetRequest().(*GreptimeRequest_Query); ok {
		return x.Query
	}
	return nil
}

func (x *GreptimeRequest) GetDdl() *DdlRequest {
	if x, ok := x.GetRequest().(*GreptimeRequest_Ddl); ok {
		return x.Ddl
	}
	return nil
}

type isGreptimeRequest_Request interface {
	isGreptimeRequest_Request()
}

type GreptimeRequest_Insert struct {
	Insert *InsertRequest `protobuf:"bytes,2,opt,name=insert,proto3,oneof"`
}

type GreptimeRequest_Query struct {
	Query *QueryRequest `protobuf:"bytes,3,opt,name=query,proto3,oneof"`
}

type GreptimeRequest_Ddl struct {
	Ddl *DdlRequest `protobuf:"bytes,4,opt,name=ddl,proto3,oneof"`
}

func (*GreptimeRequest_Insert) isGreptimeRequest_Request() {}

func (*GreptimeRequest_Query) isGreptimeRequest_Request() {}

func (*GreptimeRequest_Ddl) isGreptimeRequest_Request() {}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//
	//	*QueryRequest_Sql
	//	*QueryRequest_LogicalPlan
	Query isQueryRequest_Query `protobuf_oneof:"query"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_database_proto_rawDescGZIP(), []int{2}
}

func (m *QueryRequest) GetQuery() isQueryRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *QueryRequest) GetSql() string {
	if x, ok := x.GetQuery().(*QueryRequest_Sql); ok {
		return x.Sql
	}
	return ""
}

func (x *QueryRequest) GetLogicalPlan() []byte {
	if x, ok := x.GetQuery().(*QueryRequest_LogicalPlan); ok {
		return x.LogicalPlan
	}
	return nil
}

type isQueryRequest_Query interface {
	isQueryRequest_Query()
}

type QueryRequest_Sql struct {
	Sql string `protobuf:"bytes,1,opt,name=sql,proto3,oneof"`
}

type QueryRequest_LogicalPlan struct {
	LogicalPlan []byte `protobuf:"bytes,2,opt,name=logical_plan,json=logicalPlan,proto3,oneof"`
}

func (*QueryRequest_Sql) isQueryRequest_Query() {}

func (*QueryRequest_LogicalPlan) isQueryRequest_Query() {}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// Data is represented here.
	Columns []*Column `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	// The row_count of all columns, which include null and non-null values.
	//
	// Note: the row_count of all columns in a InsertRequest must be same.
	RowCount uint32 `protobuf:"varint,4,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	// The region number of current insert request.
	RegionNumber uint32 `protobuf:"varint,5,opt,name=region_number,json=regionNumber,proto3" json:"region_number,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRequest.ProtoReflect.Descriptor instead.
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_database_proto_rawDescGZIP(), []int{3}
}

func (x *InsertRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *InsertRequest) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *InsertRequest) GetRowCount() uint32 {
	if x != nil {
		return x.RowCount
	}
	return 0
}

func (x *InsertRequest) GetRegionNumber() uint32 {
	if x != nil {
		return x.RegionNumber
	}
	return 0
}

type AffectedRows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AffectedRows) Reset() {
	*x = AffectedRows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AffectedRows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AffectedRows) ProtoMessage() {}

func (x *AffectedRows) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AffectedRows.ProtoReflect.Descriptor instead.
func (*AffectedRows) Descriptor() ([]byte, []int) {
	return file_greptime_v1_database_proto_rawDescGZIP(), []int{4}
}

func (x *AffectedRows) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type FlightMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AffectedRows *AffectedRows `protobuf:"bytes,1,opt,name=affected_rows,json=affectedRows,proto3" json:"affected_rows,omitempty"`
}

func (x *FlightMetadata) Reset() {
	*x = FlightMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightMetadata) ProtoMessage() {}

func (x *FlightMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightMetadata.ProtoReflect.Descriptor instead.
func (*FlightMetadata) Descriptor() ([]byte, []int) {
	return file_greptime_v1_database_proto_rawDescGZIP(), []int{5}
}

func (x *FlightMetadata) GetAffectedRows() *AffectedRows {
	if x != nil {
		return x.AffectedRows
	}
	return nil
}

var File_greptime_v1_database_proto protoreflect.FileDescriptor

var file_greptime_v1_database_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x67, 0x72, 0x65, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xe6, 0x01, 0x0a, 0x0f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x69,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x03, 0x64, 0x64, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x64, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x03, 0x64, 0x64,
	0x6c, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x0c,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x03,
	0x73, 0x71, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x73, 0x71, 0x6c,
	0x12, 0x23, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61,
	0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x9f,
	0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x24, 0x0a, 0x0c, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x77, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x50, 0x0a, 0x0e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x0d, 0x61, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x0c, 0x61, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x42, 0x51, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_database_proto_rawDescOnce sync.Once
	file_greptime_v1_database_proto_rawDescData = file_greptime_v1_database_proto_rawDesc
)

func file_greptime_v1_database_proto_rawDescGZIP() []byte {
	file_greptime_v1_database_proto_rawDescOnce.Do(func() {
		file_greptime_v1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_database_proto_rawDescData)
	})
	return file_greptime_v1_database_proto_rawDescData
}

var file_greptime_v1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_greptime_v1_database_proto_goTypes = []interface{}{
	(*RequestHeader)(nil),   // 0: greptime.v1.RequestHeader
	(*GreptimeRequest)(nil), // 1: greptime.v1.GreptimeRequest
	(*QueryRequest)(nil),    // 2: greptime.v1.QueryRequest
	(*InsertRequest)(nil),   // 3: greptime.v1.InsertRequest
	(*AffectedRows)(nil),    // 4: greptime.v1.AffectedRows
	(*FlightMetadata)(nil),  // 5: greptime.v1.FlightMetadata
	(*DdlRequest)(nil),      // 6: greptime.v1.DdlRequest
	(*Column)(nil),          // 7: greptime.v1.Column
}
var file_greptime_v1_database_proto_depIdxs = []int32{
	0, // 0: greptime.v1.GreptimeRequest.header:type_name -> greptime.v1.RequestHeader
	3, // 1: greptime.v1.GreptimeRequest.insert:type_name -> greptime.v1.InsertRequest
	2, // 2: greptime.v1.GreptimeRequest.query:type_name -> greptime.v1.QueryRequest
	6, // 3: greptime.v1.GreptimeRequest.ddl:type_name -> greptime.v1.DdlRequest
	7, // 4: greptime.v1.InsertRequest.columns:type_name -> greptime.v1.Column
	4, // 5: greptime.v1.FlightMetadata.affected_rows:type_name -> greptime.v1.AffectedRows
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_greptime_v1_database_proto_init() }
func file_greptime_v1_database_proto_init() {
	if File_greptime_v1_database_proto != nil {
		return
	}
	file_greptime_v1_ddl_proto_init()
	file_greptime_v1_column_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
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
		file_greptime_v1_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreptimeRequest); i {
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
		file_greptime_v1_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_greptime_v1_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRequest); i {
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
		file_greptime_v1_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AffectedRows); i {
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
		file_greptime_v1_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightMetadata); i {
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
	file_greptime_v1_database_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GreptimeRequest_Insert)(nil),
		(*GreptimeRequest_Query)(nil),
		(*GreptimeRequest_Ddl)(nil),
	}
	file_greptime_v1_database_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*QueryRequest_Sql)(nil),
		(*QueryRequest_LogicalPlan)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greptime_v1_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_greptime_v1_database_proto_goTypes,
		DependencyIndexes: file_greptime_v1_database_proto_depIdxs,
		MessageInfos:      file_greptime_v1_database_proto_msgTypes,
	}.Build()
	File_greptime_v1_database_proto = out.File
	file_greptime_v1_database_proto_rawDesc = nil
	file_greptime_v1_database_proto_goTypes = nil
	file_greptime_v1_database_proto_depIdxs = nil
}
