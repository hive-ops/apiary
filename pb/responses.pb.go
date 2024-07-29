// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: responses.proto

package pb

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

type GetEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries  []*Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	NotFound []string `protobuf:"bytes,2,rep,name=not_found,json=notFound,proto3" json:"not_found,omitempty"`
}

func (x *GetEntriesResponse) Reset() {
	*x = GetEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntriesResponse) ProtoMessage() {}

func (x *GetEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntriesResponse.ProtoReflect.Descriptor instead.
func (*GetEntriesResponse) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{0}
}

func (x *GetEntriesResponse) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *GetEntriesResponse) GetNotFound() []string {
	if x != nil {
		return x.NotFound
	}
	return nil
}

type SetEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful []string `protobuf:"bytes,1,rep,name=successful,proto3" json:"successful,omitempty"`
	Failed     []string `protobuf:"bytes,2,rep,name=failed,proto3" json:"failed,omitempty"`
}

func (x *SetEntriesResponse) Reset() {
	*x = SetEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEntriesResponse) ProtoMessage() {}

func (x *SetEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEntriesResponse.ProtoReflect.Descriptor instead.
func (*SetEntriesResponse) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{1}
}

func (x *SetEntriesResponse) GetSuccessful() []string {
	if x != nil {
		return x.Successful
	}
	return nil
}

func (x *SetEntriesResponse) GetFailed() []string {
	if x != nil {
		return x.Failed
	}
	return nil
}

type DeleteEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful []string `protobuf:"bytes,1,rep,name=successful,proto3" json:"successful,omitempty"`
	NotFound   []string `protobuf:"bytes,2,rep,name=not_found,json=notFound,proto3" json:"not_found,omitempty"`
	Failed     []string `protobuf:"bytes,3,rep,name=failed,proto3" json:"failed,omitempty"`
}

func (x *DeleteEntriesResponse) Reset() {
	*x = DeleteEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntriesResponse) ProtoMessage() {}

func (x *DeleteEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntriesResponse.ProtoReflect.Descriptor instead.
func (*DeleteEntriesResponse) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteEntriesResponse) GetSuccessful() []string {
	if x != nil {
		return x.Successful
	}
	return nil
}

func (x *DeleteEntriesResponse) GetNotFound() []string {
	if x != nil {
		return x.NotFound
	}
	return nil
}

func (x *DeleteEntriesResponse) GetFailed() []string {
	if x != nil {
		return x.Failed
	}
	return nil
}

type ClearEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
}

func (x *ClearEntriesResponse) Reset() {
	*x = ClearEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearEntriesResponse) ProtoMessage() {}

func (x *ClearEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearEntriesResponse.ProtoReflect.Descriptor instead.
func (*ClearEntriesResponse) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{3}
}

func (x *ClearEntriesResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

var File_responses_proto protoreflect.FileDescriptor

var file_responses_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x53, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x22, 0x4c, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x22, 0x6c, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x22, 0x36, 0x0a, 0x14, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x42, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x76, 0x65, 0x2d, 0x6f, 0x70,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_responses_proto_rawDescOnce sync.Once
	file_responses_proto_rawDescData = file_responses_proto_rawDesc
)

func file_responses_proto_rawDescGZIP() []byte {
	file_responses_proto_rawDescOnce.Do(func() {
		file_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_responses_proto_rawDescData)
	})
	return file_responses_proto_rawDescData
}

var file_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_responses_proto_goTypes = []any{
	(*GetEntriesResponse)(nil),    // 0: GetEntriesResponse
	(*SetEntriesResponse)(nil),    // 1: SetEntriesResponse
	(*DeleteEntriesResponse)(nil), // 2: DeleteEntriesResponse
	(*ClearEntriesResponse)(nil),  // 3: ClearEntriesResponse
	(*Entry)(nil),                 // 4: Entry
}
var file_responses_proto_depIdxs = []int32{
	4, // 0: GetEntriesResponse.entries:type_name -> Entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_responses_proto_init() }
func file_responses_proto_init() {
	if File_responses_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_responses_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetEntriesResponse); i {
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
		file_responses_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetEntriesResponse); i {
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
		file_responses_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteEntriesResponse); i {
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
		file_responses_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ClearEntriesResponse); i {
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
			RawDescriptor: file_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_responses_proto_goTypes,
		DependencyIndexes: file_responses_proto_depIdxs,
		MessageInfos:      file_responses_proto_msgTypes,
	}.Build()
	File_responses_proto = out.File
	file_responses_proto_rawDesc = nil
	file_responses_proto_goTypes = nil
	file_responses_proto_depIdxs = nil
}
