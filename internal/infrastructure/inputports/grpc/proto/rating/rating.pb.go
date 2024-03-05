// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: rating.proto

package rating

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

type UserRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId int32 `protobuf:"varint,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Vote   int32 `protobuf:"varint,4,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *UserRatingRequest) Reset() {
	*x = UserRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRatingRequest) ProtoMessage() {}

func (x *UserRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRatingRequest.ProtoReflect.Descriptor instead.
func (*UserRatingRequest) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{0}
}

func (x *UserRatingRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRatingRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRatingRequest) GetPostId() int32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *UserRatingRequest) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

type RatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RatingResponse) Reset() {
	*x = RatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingResponse) ProtoMessage() {}

func (x *RatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingResponse.ProtoReflect.Descriptor instead.
func (*RatingResponse) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{1}
}

func (x *RatingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rating_proto protoreflect.FileDescriptor

var file_rating_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x69, 0x0a,
	0x11, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x5c, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74,
	0x65, 0x12, 0x20, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rating_proto_rawDescOnce sync.Once
	file_rating_proto_rawDescData = file_rating_proto_rawDesc
)

func file_rating_proto_rawDescGZIP() []byte {
	file_rating_proto_rawDescOnce.Do(func() {
		file_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_rating_proto_rawDescData)
	})
	return file_rating_proto_rawDescData
}

var file_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rating_proto_goTypes = []interface{}{
	(*UserRatingRequest)(nil), // 0: ratingService.UserRatingRequest
	(*RatingResponse)(nil),    // 1: ratingService.RatingResponse
}
var file_rating_proto_depIdxs = []int32{
	0, // 0: ratingService.RatingService.UserVote:input_type -> ratingService.UserRatingRequest
	1, // 1: ratingService.RatingService.UserVote:output_type -> ratingService.RatingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rating_proto_init() }
func file_rating_proto_init() {
	if File_rating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRatingRequest); i {
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
		file_rating_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingResponse); i {
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
			RawDescriptor: file_rating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rating_proto_goTypes,
		DependencyIndexes: file_rating_proto_depIdxs,
		MessageInfos:      file_rating_proto_msgTypes,
	}.Build()
	File_rating_proto = out.File
	file_rating_proto_rawDesc = nil
	file_rating_proto_goTypes = nil
	file_rating_proto_depIdxs = nil
}
