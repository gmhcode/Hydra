// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: Ship.proto

package hydraproto

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

// option go_package = "github.com/Hydra/hydracomms/hydramessages/protobuffr";
type Ship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shipname    string             `protobuf:"bytes,1,opt,name=shipname,proto3" json:"shipname,omitempty"`
	CaptainName string             `protobuf:"bytes,2,opt,name=CaptainName,proto3" json:"CaptainName,omitempty"`
	Crew        []*Ship_CrewMember `protobuf:"bytes,3,rep,name=Crew,proto3" json:"Crew,omitempty"`
}

func (x *Ship) Reset() {
	*x = Ship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ship_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ship) ProtoMessage() {}

func (x *Ship) ProtoReflect() protoreflect.Message {
	mi := &file_Ship_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ship.ProtoReflect.Descriptor instead.
func (*Ship) Descriptor() ([]byte, []int) {
	return file_Ship_proto_rawDescGZIP(), []int{0}
}

func (x *Ship) GetShipname() string {
	if x != nil {
		return x.Shipname
	}
	return ""
}

func (x *Ship) GetCaptainName() string {
	if x != nil {
		return x.CaptainName
	}
	return ""
}

func (x *Ship) GetCrew() []*Ship_CrewMember {
	if x != nil {
		return x.Crew
	}
	return nil
}

type Ship_CrewMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SecClearance int32  `protobuf:"varint,3,opt,name=secClearance,proto3" json:"secClearance,omitempty"`
	Position     string `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Ship_CrewMember) Reset() {
	*x = Ship_CrewMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ship_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ship_CrewMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ship_CrewMember) ProtoMessage() {}

func (x *Ship_CrewMember) ProtoReflect() protoreflect.Message {
	mi := &file_Ship_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ship_CrewMember.ProtoReflect.Descriptor instead.
func (*Ship_CrewMember) Descriptor() ([]byte, []int) {
	return file_Ship_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Ship_CrewMember) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ship_CrewMember) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ship_CrewMember) GetSecClearance() int32 {
	if x != nil {
		return x.SecClearance
	}
	return 0
}

func (x *Ship_CrewMember) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

var File_Ship_proto protoreflect.FileDescriptor

var file_Ship_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x53, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x79,
	0x64, 0x72, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x04, 0x53, 0x68, 0x69,
	0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x04, 0x43, 0x72, 0x65, 0x77, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x68, 0x79, 0x64, 0x72, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x2e,
	0x43, 0x72, 0x65, 0x77, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x04, 0x43, 0x72, 0x65, 0x77,
	0x1a, 0x70, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x77, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Ship_proto_rawDescOnce sync.Once
	file_Ship_proto_rawDescData = file_Ship_proto_rawDesc
)

func file_Ship_proto_rawDescGZIP() []byte {
	file_Ship_proto_rawDescOnce.Do(func() {
		file_Ship_proto_rawDescData = protoimpl.X.CompressGZIP(file_Ship_proto_rawDescData)
	})
	return file_Ship_proto_rawDescData
}

var file_Ship_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Ship_proto_goTypes = []interface{}{
	(*Ship)(nil),            // 0: hydraproto.Ship
	(*Ship_CrewMember)(nil), // 1: hydraproto.Ship.CrewMember
}
var file_Ship_proto_depIdxs = []int32{
	1, // 0: hydraproto.Ship.Crew:type_name -> hydraproto.Ship.CrewMember
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Ship_proto_init() }
func file_Ship_proto_init() {
	if File_Ship_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Ship_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ship); i {
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
		file_Ship_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ship_CrewMember); i {
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
			RawDescriptor: file_Ship_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Ship_proto_goTypes,
		DependencyIndexes: file_Ship_proto_depIdxs,
		MessageInfos:      file_Ship_proto_msgTypes,
	}.Build()
	File_Ship_proto = out.File
	file_Ship_proto_rawDesc = nil
	file_Ship_proto_goTypes = nil
	file_Ship_proto_depIdxs = nil
}
