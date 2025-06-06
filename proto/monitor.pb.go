// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/monitor.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Heartbeat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NodoId        string                 `protobuf:"bytes,1,opt,name=nodoId,proto3" json:"nodoId,omitempty"`
	MarcaTiempo   int64                  `protobuf:"varint,2,opt,name=marcaTiempo,proto3" json:"marcaTiempo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	mi := &file_proto_monitor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_proto_monitor_proto_rawDescGZIP(), []int{0}
}

func (x *Heartbeat) GetNodoId() string {
	if x != nil {
		return x.NodoId
	}
	return ""
}

func (x *Heartbeat) GetMarcaTiempo() int64 {
	if x != nil {
		return x.MarcaTiempo
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mensaje       string                 `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_proto_monitor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

var File_proto_monitor_proto protoreflect.FileDescriptor

const file_proto_monitor_proto_rawDesc = "" +
	"\n" +
	"\x13proto/monitor.proto\x12\amonitor\"E\n" +
	"\tHeartbeat\x12\x16\n" +
	"\x06nodoId\x18\x01 \x01(\tR\x06nodoId\x12 \n" +
	"\vmarcaTiempo\x18\x02 \x01(\x03R\vmarcaTiempo\"\x1f\n" +
	"\x03Ack\x12\x18\n" +
	"\amensaje\x18\x01 \x01(\tR\amensaje2@\n" +
	"\aMonitor\x125\n" +
	"\x0fEnviarHeartbeat\x12\x12.monitor.Heartbeat\x1a\f.monitor.Ack(\x01B\tZ\a./protob\x06proto3"

var (
	file_proto_monitor_proto_rawDescOnce sync.Once
	file_proto_monitor_proto_rawDescData []byte
)

func file_proto_monitor_proto_rawDescGZIP() []byte {
	file_proto_monitor_proto_rawDescOnce.Do(func() {
		file_proto_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_monitor_proto_rawDesc), len(file_proto_monitor_proto_rawDesc)))
	})
	return file_proto_monitor_proto_rawDescData
}

var file_proto_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_monitor_proto_goTypes = []any{
	(*Heartbeat)(nil), // 0: monitor.Heartbeat
	(*Ack)(nil),       // 1: monitor.Ack
}
var file_proto_monitor_proto_depIdxs = []int32{
	0, // 0: monitor.Monitor.EnviarHeartbeat:input_type -> monitor.Heartbeat
	1, // 1: monitor.Monitor.EnviarHeartbeat:output_type -> monitor.Ack
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_monitor_proto_init() }
func file_proto_monitor_proto_init() {
	if File_proto_monitor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_monitor_proto_rawDesc), len(file_proto_monitor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_monitor_proto_goTypes,
		DependencyIndexes: file_proto_monitor_proto_depIdxs,
		MessageInfos:      file_proto_monitor_proto_msgTypes,
	}.Build()
	File_proto_monitor_proto = out.File
	file_proto_monitor_proto_goTypes = nil
	file_proto_monitor_proto_depIdxs = nil
}
