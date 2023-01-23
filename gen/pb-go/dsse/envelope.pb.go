// https://raw.githubusercontent.com/secure-systems-lab/dsse/9c813476bd36de70a5738c72e784f123ecea16af/envelope.proto

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: envelope.proto

package dsse

import (
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

// An authenticated message of arbitrary type.
type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message to be signed. (In JSON, this is encoded as base64.)
	// REQUIRED.
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// String unambiguously identifying how to interpret payload.
	// REQUIRED.
	PayloadType string `protobuf:"bytes,2,opt,name=payloadType,proto3" json:"payloadType,omitempty"`
	// Signature over:
	//     PAE(type, body)
	// Where PAE is defined as:
	// PAE(type, body) = "DSSEv1" + SP + LEN(type) + SP + type + SP + LEN(body) + SP + body
	// +               = concatenation
	// SP              = ASCII space [0x20]
	// "DSSEv1"        = ASCII [0x44, 0x53, 0x53, 0x45, 0x76, 0x31]
	// LEN(s)          = ASCII decimal encoding of the byte length of s, with no leading zeros
	// REQUIRED (length >= 1).
	Signatures []*Signature `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_envelope_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Envelope) GetPayloadType() string {
	if x != nil {
		return x.PayloadType
	}
	return ""
}

func (x *Envelope) GetSignatures() []*Signature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signature itself. (In JSON, this is encoded as base64.)
	// REQUIRED.
	Sig []byte `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	// *Unauthenticated* hint identifying which public key was used.
	// OPTIONAL.
	Keyid string `protobuf:"bytes,2,opt,name=keyid,proto3" json:"keyid,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_envelope_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *Signature) GetKeyid() string {
	if x != nil {
		return x.Keyid
	}
	return ""
}

var File_envelope_proto protoreflect.FileDescriptor

var file_envelope_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x69, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a,
	0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x6b, 0x65, 0x79, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b,
	0x65, 0x79, 0x69, 0x64, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x73, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envelope_proto_rawDescOnce sync.Once
	file_envelope_proto_rawDescData = file_envelope_proto_rawDesc
)

func file_envelope_proto_rawDescGZIP() []byte {
	file_envelope_proto_rawDescOnce.Do(func() {
		file_envelope_proto_rawDescData = protoimpl.X.CompressGZIP(file_envelope_proto_rawDescData)
	})
	return file_envelope_proto_rawDescData
}

var file_envelope_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envelope_proto_goTypes = []interface{}{
	(*Envelope)(nil),  // 0: io.intoto.Envelope
	(*Signature)(nil), // 1: io.intoto.Signature
}
var file_envelope_proto_depIdxs = []int32{
	1, // 0: io.intoto.Envelope.signatures:type_name -> io.intoto.Signature
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envelope_proto_init() }
func file_envelope_proto_init() {
	if File_envelope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envelope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_envelope_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
			RawDescriptor: file_envelope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envelope_proto_goTypes,
		DependencyIndexes: file_envelope_proto_depIdxs,
		MessageInfos:      file_envelope_proto_msgTypes,
	}.Build()
	File_envelope_proto = out.File
	file_envelope_proto_rawDesc = nil
	file_envelope_proto_goTypes = nil
	file_envelope_proto_depIdxs = nil
}
