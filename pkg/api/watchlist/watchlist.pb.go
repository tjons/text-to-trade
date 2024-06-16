// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: watchlist/watchlist.proto

package watchlist

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Watchlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbols []string `protobuf:"bytes,3,rep,name=symbols,proto3" json:"symbols,omitempty"`
	UserId  uint32   `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Watchlist) Reset() {
	*x = Watchlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watchlist_watchlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Watchlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Watchlist) ProtoMessage() {}

func (x *Watchlist) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_watchlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Watchlist.ProtoReflect.Descriptor instead.
func (*Watchlist) Descriptor() ([]byte, []int) {
	return file_watchlist_watchlist_proto_rawDescGZIP(), []int{0}
}

func (x *Watchlist) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Watchlist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Watchlist) GetSymbols() []string {
	if x != nil {
		return x.Symbols
	}
	return nil
}

func (x *Watchlist) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type WatchlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *WatchlistRequest) Reset() {
	*x = WatchlistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watchlist_watchlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchlistRequest) ProtoMessage() {}

func (x *WatchlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_watchlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchlistRequest.ProtoReflect.Descriptor instead.
func (*WatchlistRequest) Descriptor() ([]byte, []int) {
	return file_watchlist_watchlist_proto_rawDescGZIP(), []int{1}
}

func (x *WatchlistRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WatchlistRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type WatchlistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Watchlist *Watchlist `protobuf:"bytes,1,opt,name=watchlist,proto3" json:"watchlist,omitempty"`
}

func (x *WatchlistResponse) Reset() {
	*x = WatchlistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watchlist_watchlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchlistResponse) ProtoMessage() {}

func (x *WatchlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_watchlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchlistResponse.ProtoReflect.Descriptor instead.
func (*WatchlistResponse) Descriptor() ([]byte, []int) {
	return file_watchlist_watchlist_proto_rawDescGZIP(), []int{2}
}

func (x *WatchlistResponse) GetWatchlist() *Watchlist {
	if x != nil {
		return x.Watchlist
	}
	return nil
}

type WatchlistListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Watchlists []*Watchlist `protobuf:"bytes,1,rep,name=watchlists,proto3" json:"watchlists,omitempty"`
}

func (x *WatchlistListResponse) Reset() {
	*x = WatchlistListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watchlist_watchlist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchlistListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchlistListResponse) ProtoMessage() {}

func (x *WatchlistListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_watchlist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchlistListResponse.ProtoReflect.Descriptor instead.
func (*WatchlistListResponse) Descriptor() ([]byte, []int) {
	return file_watchlist_watchlist_proto_rawDescGZIP(), []int{3}
}

func (x *WatchlistListResponse) GetWatchlists() []*Watchlist {
	if x != nil {
		return x.Watchlists
	}
	return nil
}

var File_watchlist_watchlist_proto protoreflect.FileDescriptor

var file_watchlist_watchlist_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x10, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x4d,
	0x0a, 0x15, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x0a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x32, 0xb3, 0x03,
	0x0a, 0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x78, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x62, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x1a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1b,
	0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x1a, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x1a, 0x1a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x42, 0x3f, 0x92, 0x41, 0x0a, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a,
	0x01, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x6a, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x74, 0x6f, 0x2d, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_watchlist_watchlist_proto_rawDescOnce sync.Once
	file_watchlist_watchlist_proto_rawDescData = file_watchlist_watchlist_proto_rawDesc
)

func file_watchlist_watchlist_proto_rawDescGZIP() []byte {
	file_watchlist_watchlist_proto_rawDescOnce.Do(func() {
		file_watchlist_watchlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_watchlist_watchlist_proto_rawDescData)
	})
	return file_watchlist_watchlist_proto_rawDescData
}

var file_watchlist_watchlist_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_watchlist_watchlist_proto_goTypes = []interface{}{
	(*Watchlist)(nil),             // 0: watchlist.Watchlist
	(*WatchlistRequest)(nil),      // 1: watchlist.WatchlistRequest
	(*WatchlistResponse)(nil),     // 2: watchlist.WatchlistResponse
	(*WatchlistListResponse)(nil), // 3: watchlist.WatchlistListResponse
}
var file_watchlist_watchlist_proto_depIdxs = []int32{
	0, // 0: watchlist.WatchlistResponse.watchlist:type_name -> watchlist.Watchlist
	0, // 1: watchlist.WatchlistListResponse.watchlists:type_name -> watchlist.Watchlist
	1, // 2: watchlist.WatchlistService.GetWatchlist:input_type -> watchlist.WatchlistRequest
	1, // 3: watchlist.WatchlistService.ListWatchlists:input_type -> watchlist.WatchlistRequest
	0, // 4: watchlist.WatchlistService.CreateWatchlist:input_type -> watchlist.Watchlist
	0, // 5: watchlist.WatchlistService.UpdateWatchlist:input_type -> watchlist.Watchlist
	2, // 6: watchlist.WatchlistService.GetWatchlist:output_type -> watchlist.WatchlistResponse
	3, // 7: watchlist.WatchlistService.ListWatchlists:output_type -> watchlist.WatchlistListResponse
	2, // 8: watchlist.WatchlistService.CreateWatchlist:output_type -> watchlist.WatchlistResponse
	2, // 9: watchlist.WatchlistService.UpdateWatchlist:output_type -> watchlist.WatchlistResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_watchlist_watchlist_proto_init() }
func file_watchlist_watchlist_proto_init() {
	if File_watchlist_watchlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_watchlist_watchlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Watchlist); i {
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
		file_watchlist_watchlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchlistRequest); i {
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
		file_watchlist_watchlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchlistResponse); i {
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
		file_watchlist_watchlist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchlistListResponse); i {
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
			RawDescriptor: file_watchlist_watchlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_watchlist_watchlist_proto_goTypes,
		DependencyIndexes: file_watchlist_watchlist_proto_depIdxs,
		MessageInfos:      file_watchlist_watchlist_proto_msgTypes,
	}.Build()
	File_watchlist_watchlist_proto = out.File
	file_watchlist_watchlist_proto_rawDesc = nil
	file_watchlist_watchlist_proto_goTypes = nil
	file_watchlist_watchlist_proto_depIdxs = nil
}
