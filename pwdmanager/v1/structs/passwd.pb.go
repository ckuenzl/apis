// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: pwdmanager/v1/structs/passwd.proto

package structs

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Password struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pwd string `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
}

func (x *Password) Reset() {
	*x = Password{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwdmanager_v1_structs_passwd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Password) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Password) ProtoMessage() {}

func (x *Password) ProtoReflect() protoreflect.Message {
	mi := &file_pwdmanager_v1_structs_passwd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Password.ProtoReflect.Descriptor instead.
func (*Password) Descriptor() ([]byte, []int) {
	return file_pwdmanager_v1_structs_passwd_proto_rawDescGZIP(), []int{0}
}

func (x *Password) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Password) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type PasswordList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passwords []*Password `protobuf:"bytes,1,rep,name=passwords,proto3" json:"passwords,omitempty"`
}

func (x *PasswordList) Reset() {
	*x = PasswordList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwdmanager_v1_structs_passwd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordList) ProtoMessage() {}

func (x *PasswordList) ProtoReflect() protoreflect.Message {
	mi := &file_pwdmanager_v1_structs_passwd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordList.ProtoReflect.Descriptor instead.
func (*PasswordList) Descriptor() ([]byte, []int) {
	return file_pwdmanager_v1_structs_passwd_proto_rawDescGZIP(), []int{1}
}

func (x *PasswordList) GetPasswords() []*Password {
	if x != nil {
		return x.Passwords
	}
	return nil
}

var File_pwdmanager_v1_structs_passwd_proto protoreflect.FileDescriptor

var file_pwdmanager_v1_structs_passwd_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x77, 0x64, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x77, 0x64, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x22, 0x2c, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x70, 0x77, 0x64, 0x22, 0x4a, 0x0a, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x77, 0x64, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x09, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x70, 0x77, 0x64, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x3b, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pwdmanager_v1_structs_passwd_proto_rawDescOnce sync.Once
	file_pwdmanager_v1_structs_passwd_proto_rawDescData = file_pwdmanager_v1_structs_passwd_proto_rawDesc
)

func file_pwdmanager_v1_structs_passwd_proto_rawDescGZIP() []byte {
	file_pwdmanager_v1_structs_passwd_proto_rawDescOnce.Do(func() {
		file_pwdmanager_v1_structs_passwd_proto_rawDescData = protoimpl.X.CompressGZIP(file_pwdmanager_v1_structs_passwd_proto_rawDescData)
	})
	return file_pwdmanager_v1_structs_passwd_proto_rawDescData
}

var file_pwdmanager_v1_structs_passwd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pwdmanager_v1_structs_passwd_proto_goTypes = []interface{}{
	(*Password)(nil),     // 0: pwdmanager.structs.Password
	(*PasswordList)(nil), // 1: pwdmanager.structs.PasswordList
}
var file_pwdmanager_v1_structs_passwd_proto_depIdxs = []int32{
	0, // 0: pwdmanager.structs.PasswordList.passwords:type_name -> pwdmanager.structs.Password
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pwdmanager_v1_structs_passwd_proto_init() }
func file_pwdmanager_v1_structs_passwd_proto_init() {
	if File_pwdmanager_v1_structs_passwd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pwdmanager_v1_structs_passwd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Password); i {
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
		file_pwdmanager_v1_structs_passwd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordList); i {
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
			RawDescriptor: file_pwdmanager_v1_structs_passwd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pwdmanager_v1_structs_passwd_proto_goTypes,
		DependencyIndexes: file_pwdmanager_v1_structs_passwd_proto_depIdxs,
		MessageInfos:      file_pwdmanager_v1_structs_passwd_proto_msgTypes,
	}.Build()
	File_pwdmanager_v1_structs_passwd_proto = out.File
	file_pwdmanager_v1_structs_passwd_proto_rawDesc = nil
	file_pwdmanager_v1_structs_passwd_proto_goTypes = nil
	file_pwdmanager_v1_structs_passwd_proto_depIdxs = nil
}
