// Code generated by protoc-gen-go.
// source: thetest.proto
// DO NOT EDIT!

/*
Package gogoprototest is a generated protocol buffer package.

It is generated from these files:
	thetest.proto

It has these top-level messages:
	SimpleMessage
	NidOptNative
	NinOptNative
	NidRepNative
	NinRepNative
	NidRepPackedNative
	NinRepPackedNative
	NidOptStruct
	NinOptStruct
	NidRepStruct
	NinRepStruct
	NidEmbeddedStruct
	NinEmbeddedStruct
	NidNestedStruct
	NinNestedStruct
	NidOptUuidAsBytes
	NinOptUuidAsBytes
	NidRepUuidAsBytes
	NinRepUuidAsBytes
	NidOptUint128AsBytes
	NinOptUint128AsBytes
	NidRepUint128AsBytes
	NinRepUint128AsBytes
	NinOptNativeDefault
*/
package gogoprototest

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SimpleMessage struct {
	Id               []byte  `protobuf:"bytes,1,opt" json:"Id,omitempty"`
	SimpleName       *string `protobuf:"bytes,2,opt" json:"SimpleName,omitempty"`
	Time             *int64  `protobuf:"varint,3,opt" json:"Time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}

func (m *SimpleMessage) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SimpleMessage) GetSimpleName() string {
	if m != nil && m.SimpleName != nil {
		return *m.SimpleName
	}
	return ""
}

func (m *SimpleMessage) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type NidOptNative struct {
	Field1           *float64 `protobuf:"fixed64,1,opt" json:"Field1,omitempty"`
	Field2           *float32 `protobuf:"fixed32,2,opt" json:"Field2,omitempty"`
	Field3           *int32   `protobuf:"varint,3,opt" json:"Field3,omitempty"`
	Field4           *int64   `protobuf:"varint,4,opt" json:"Field4,omitempty"`
	Field5           *uint32  `protobuf:"varint,5,opt" json:"Field5,omitempty"`
	Field6           *uint64  `protobuf:"varint,6,opt" json:"Field6,omitempty"`
	Field7           *int32   `protobuf:"zigzag32,7,opt" json:"Field7,omitempty"`
	Field8           *int64   `protobuf:"zigzag64,8,opt" json:"Field8,omitempty"`
	Field9           *uint32  `protobuf:"fixed32,9,opt" json:"Field9,omitempty"`
	Field10          *int32   `protobuf:"fixed32,10,opt" json:"Field10,omitempty"`
	Field11          *uint64  `protobuf:"fixed64,11,opt" json:"Field11,omitempty"`
	Field12          *int64   `protobuf:"fixed64,12,opt" json:"Field12,omitempty"`
	Field13          *bool    `protobuf:"varint,13,opt" json:"Field13,omitempty"`
	Field14          *string  `protobuf:"bytes,14,opt" json:"Field14,omitempty"`
	Field15          []byte   `protobuf:"bytes,15,opt" json:"Field15,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NidOptNative) Reset()         { *m = NidOptNative{} }
func (m *NidOptNative) String() string { return proto.CompactTextString(m) }
func (*NidOptNative) ProtoMessage()    {}

func (m *NidOptNative) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NidOptNative) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func (m *NidOptNative) GetField3() int32 {
	if m != nil && m.Field3 != nil {
		return *m.Field3
	}
	return 0
}

func (m *NidOptNative) GetField4() int64 {
	if m != nil && m.Field4 != nil {
		return *m.Field4
	}
	return 0
}

func (m *NidOptNative) GetField5() uint32 {
	if m != nil && m.Field5 != nil {
		return *m.Field5
	}
	return 0
}

func (m *NidOptNative) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return 0
}

func (m *NidOptNative) GetField7() int32 {
	if m != nil && m.Field7 != nil {
		return *m.Field7
	}
	return 0
}

func (m *NidOptNative) GetField8() int64 {
	if m != nil && m.Field8 != nil {
		return *m.Field8
	}
	return 0
}

func (m *NidOptNative) GetField9() uint32 {
	if m != nil && m.Field9 != nil {
		return *m.Field9
	}
	return 0
}

func (m *NidOptNative) GetField10() int32 {
	if m != nil && m.Field10 != nil {
		return *m.Field10
	}
	return 0
}

func (m *NidOptNative) GetField11() uint64 {
	if m != nil && m.Field11 != nil {
		return *m.Field11
	}
	return 0
}

func (m *NidOptNative) GetField12() int64 {
	if m != nil && m.Field12 != nil {
		return *m.Field12
	}
	return 0
}

func (m *NidOptNative) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return false
}

func (m *NidOptNative) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *NidOptNative) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NinOptNative struct {
	Field1           *float64 `protobuf:"fixed64,1,opt" json:"Field1,omitempty"`
	Field2           *float32 `protobuf:"fixed32,2,opt" json:"Field2,omitempty"`
	Field3           *int32   `protobuf:"varint,3,opt" json:"Field3,omitempty"`
	Field4           *int64   `protobuf:"varint,4,opt" json:"Field4,omitempty"`
	Field5           *uint32  `protobuf:"varint,5,opt" json:"Field5,omitempty"`
	Field6           *uint64  `protobuf:"varint,6,opt" json:"Field6,omitempty"`
	Field7           *int32   `protobuf:"zigzag32,7,opt" json:"Field7,omitempty"`
	Field8           *int64   `protobuf:"zigzag64,8,opt" json:"Field8,omitempty"`
	Field9           *uint32  `protobuf:"fixed32,9,opt" json:"Field9,omitempty"`
	Field10          *int32   `protobuf:"fixed32,10,opt" json:"Field10,omitempty"`
	Field11          *uint64  `protobuf:"fixed64,11,opt" json:"Field11,omitempty"`
	Field12          *int64   `protobuf:"fixed64,12,opt" json:"Field12,omitempty"`
	Field13          *bool    `protobuf:"varint,13,opt" json:"Field13,omitempty"`
	Field14          *string  `protobuf:"bytes,14,opt" json:"Field14,omitempty"`
	Field15          []byte   `protobuf:"bytes,15,opt" json:"Field15,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NinOptNative) Reset()         { *m = NinOptNative{} }
func (m *NinOptNative) String() string { return proto.CompactTextString(m) }
func (*NinOptNative) ProtoMessage()    {}

func (m *NinOptNative) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NinOptNative) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func (m *NinOptNative) GetField3() int32 {
	if m != nil && m.Field3 != nil {
		return *m.Field3
	}
	return 0
}

func (m *NinOptNative) GetField4() int64 {
	if m != nil && m.Field4 != nil {
		return *m.Field4
	}
	return 0
}

func (m *NinOptNative) GetField5() uint32 {
	if m != nil && m.Field5 != nil {
		return *m.Field5
	}
	return 0
}

func (m *NinOptNative) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return 0
}

func (m *NinOptNative) GetField7() int32 {
	if m != nil && m.Field7 != nil {
		return *m.Field7
	}
	return 0
}

func (m *NinOptNative) GetField8() int64 {
	if m != nil && m.Field8 != nil {
		return *m.Field8
	}
	return 0
}

func (m *NinOptNative) GetField9() uint32 {
	if m != nil && m.Field9 != nil {
		return *m.Field9
	}
	return 0
}

func (m *NinOptNative) GetField10() int32 {
	if m != nil && m.Field10 != nil {
		return *m.Field10
	}
	return 0
}

func (m *NinOptNative) GetField11() uint64 {
	if m != nil && m.Field11 != nil {
		return *m.Field11
	}
	return 0
}

func (m *NinOptNative) GetField12() int64 {
	if m != nil && m.Field12 != nil {
		return *m.Field12
	}
	return 0
}

func (m *NinOptNative) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return false
}

func (m *NinOptNative) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *NinOptNative) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NidRepNative struct {
	Field1           []float64 `protobuf:"fixed64,1,rep" json:"Field1,omitempty"`
	Field2           []float32 `protobuf:"fixed32,2,rep" json:"Field2,omitempty"`
	Field3           []int32   `protobuf:"varint,3,rep" json:"Field3,omitempty"`
	Field4           []int64   `protobuf:"varint,4,rep" json:"Field4,omitempty"`
	Field5           []uint32  `protobuf:"varint,5,rep" json:"Field5,omitempty"`
	Field6           []uint64  `protobuf:"varint,6,rep" json:"Field6,omitempty"`
	Field7           []int32   `protobuf:"zigzag32,7,rep" json:"Field7,omitempty"`
	Field8           []int64   `protobuf:"zigzag64,8,rep" json:"Field8,omitempty"`
	Field9           []uint32  `protobuf:"fixed32,9,rep" json:"Field9,omitempty"`
	Field10          []int32   `protobuf:"fixed32,10,rep" json:"Field10,omitempty"`
	Field11          []uint64  `protobuf:"fixed64,11,rep" json:"Field11,omitempty"`
	Field12          []int64   `protobuf:"fixed64,12,rep" json:"Field12,omitempty"`
	Field13          []bool    `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	Field14          []string  `protobuf:"bytes,14,rep" json:"Field14,omitempty"`
	Field15          [][]byte  `protobuf:"bytes,15,rep" json:"Field15,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NidRepNative) Reset()         { *m = NidRepNative{} }
func (m *NidRepNative) String() string { return proto.CompactTextString(m) }
func (*NidRepNative) ProtoMessage()    {}

func (m *NidRepNative) GetField1() []float64 {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NidRepNative) GetField2() []float32 {
	if m != nil {
		return m.Field2
	}
	return nil
}

func (m *NidRepNative) GetField3() []int32 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NidRepNative) GetField4() []int64 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NidRepNative) GetField5() []uint32 {
	if m != nil {
		return m.Field5
	}
	return nil
}

func (m *NidRepNative) GetField6() []uint64 {
	if m != nil {
		return m.Field6
	}
	return nil
}

func (m *NidRepNative) GetField7() []int32 {
	if m != nil {
		return m.Field7
	}
	return nil
}

func (m *NidRepNative) GetField8() []int64 {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NidRepNative) GetField9() []uint32 {
	if m != nil {
		return m.Field9
	}
	return nil
}

func (m *NidRepNative) GetField10() []int32 {
	if m != nil {
		return m.Field10
	}
	return nil
}

func (m *NidRepNative) GetField11() []uint64 {
	if m != nil {
		return m.Field11
	}
	return nil
}

func (m *NidRepNative) GetField12() []int64 {
	if m != nil {
		return m.Field12
	}
	return nil
}

func (m *NidRepNative) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

func (m *NidRepNative) GetField14() []string {
	if m != nil {
		return m.Field14
	}
	return nil
}

func (m *NidRepNative) GetField15() [][]byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NinRepNative struct {
	Field1           []float64 `protobuf:"fixed64,1,rep" json:"Field1,omitempty"`
	Field2           []float32 `protobuf:"fixed32,2,rep" json:"Field2,omitempty"`
	Field3           []int32   `protobuf:"varint,3,rep" json:"Field3,omitempty"`
	Field4           []int64   `protobuf:"varint,4,rep" json:"Field4,omitempty"`
	Field5           []uint32  `protobuf:"varint,5,rep" json:"Field5,omitempty"`
	Field6           []uint64  `protobuf:"varint,6,rep" json:"Field6,omitempty"`
	Field7           []int32   `protobuf:"zigzag32,7,rep" json:"Field7,omitempty"`
	Field8           []int64   `protobuf:"zigzag64,8,rep" json:"Field8,omitempty"`
	Field9           []uint32  `protobuf:"fixed32,9,rep" json:"Field9,omitempty"`
	Field10          []int32   `protobuf:"fixed32,10,rep" json:"Field10,omitempty"`
	Field11          []uint64  `protobuf:"fixed64,11,rep" json:"Field11,omitempty"`
	Field12          []int64   `protobuf:"fixed64,12,rep" json:"Field12,omitempty"`
	Field13          []bool    `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	Field14          []string  `protobuf:"bytes,14,rep" json:"Field14,omitempty"`
	Field15          [][]byte  `protobuf:"bytes,15,rep" json:"Field15,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NinRepNative) Reset()         { *m = NinRepNative{} }
func (m *NinRepNative) String() string { return proto.CompactTextString(m) }
func (*NinRepNative) ProtoMessage()    {}

func (m *NinRepNative) GetField1() []float64 {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinRepNative) GetField2() []float32 {
	if m != nil {
		return m.Field2
	}
	return nil
}

func (m *NinRepNative) GetField3() []int32 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinRepNative) GetField4() []int64 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinRepNative) GetField5() []uint32 {
	if m != nil {
		return m.Field5
	}
	return nil
}

func (m *NinRepNative) GetField6() []uint64 {
	if m != nil {
		return m.Field6
	}
	return nil
}

func (m *NinRepNative) GetField7() []int32 {
	if m != nil {
		return m.Field7
	}
	return nil
}

func (m *NinRepNative) GetField8() []int64 {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NinRepNative) GetField9() []uint32 {
	if m != nil {
		return m.Field9
	}
	return nil
}

func (m *NinRepNative) GetField10() []int32 {
	if m != nil {
		return m.Field10
	}
	return nil
}

func (m *NinRepNative) GetField11() []uint64 {
	if m != nil {
		return m.Field11
	}
	return nil
}

func (m *NinRepNative) GetField12() []int64 {
	if m != nil {
		return m.Field12
	}
	return nil
}

func (m *NinRepNative) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

func (m *NinRepNative) GetField14() []string {
	if m != nil {
		return m.Field14
	}
	return nil
}

func (m *NinRepNative) GetField15() [][]byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NidRepPackedNative struct {
	Field3           []int32 `protobuf:"varint,3,rep,packed" json:"Field3,omitempty"`
	Field4           []int64 `protobuf:"varint,4,rep,packed" json:"Field4,omitempty"`
	Field13          []bool  `protobuf:"varint,13,rep,packed" json:"Field13,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NidRepPackedNative) Reset()         { *m = NidRepPackedNative{} }
func (m *NidRepPackedNative) String() string { return proto.CompactTextString(m) }
func (*NidRepPackedNative) ProtoMessage()    {}

func (m *NidRepPackedNative) GetField3() []int32 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NidRepPackedNative) GetField4() []int64 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NidRepPackedNative) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

type NinRepPackedNative struct {
	Field3           []int32 `protobuf:"varint,3,rep" json:"Field3,omitempty"`
	Field4           []int64 `protobuf:"varint,4,rep" json:"Field4,omitempty"`
	Field13          []bool  `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NinRepPackedNative) Reset()         { *m = NinRepPackedNative{} }
func (m *NinRepPackedNative) String() string { return proto.CompactTextString(m) }
func (*NinRepPackedNative) ProtoMessage()    {}

func (m *NinRepPackedNative) GetField3() []int32 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinRepPackedNative) GetField4() []int64 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinRepPackedNative) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

type NidOptStruct struct {
	Field1           *float64      `protobuf:"fixed64,1,opt" json:"Field1,omitempty"`
	Field2           *float32      `protobuf:"fixed32,2,opt" json:"Field2,omitempty"`
	Field3           *NidOptNative `protobuf:"bytes,3,opt" json:"Field3,omitempty"`
	Field4           *NinOptNative `protobuf:"bytes,4,opt" json:"Field4,omitempty"`
	Field6           *uint64       `protobuf:"varint,6,opt" json:"Field6,omitempty"`
	Field7           *int32        `protobuf:"zigzag32,7,opt" json:"Field7,omitempty"`
	Field8           *NidOptNative `protobuf:"bytes,8,opt" json:"Field8,omitempty"`
	Field13          *bool         `protobuf:"varint,13,opt" json:"Field13,omitempty"`
	Field14          *string       `protobuf:"bytes,14,opt" json:"Field14,omitempty"`
	Field15          []byte        `protobuf:"bytes,15,opt" json:"Field15,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NidOptStruct) Reset()         { *m = NidOptStruct{} }
func (m *NidOptStruct) String() string { return proto.CompactTextString(m) }
func (*NidOptStruct) ProtoMessage()    {}

func (m *NidOptStruct) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NidOptStruct) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func (m *NidOptStruct) GetField3() *NidOptNative {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NidOptStruct) GetField4() *NinOptNative {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NidOptStruct) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return 0
}

func (m *NidOptStruct) GetField7() int32 {
	if m != nil && m.Field7 != nil {
		return *m.Field7
	}
	return 0
}

func (m *NidOptStruct) GetField8() *NidOptNative {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NidOptStruct) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return false
}

func (m *NidOptStruct) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *NidOptStruct) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NinOptStruct struct {
	Field1           *float64      `protobuf:"fixed64,1,opt" json:"Field1,omitempty"`
	Field2           *float32      `protobuf:"fixed32,2,opt" json:"Field2,omitempty"`
	Field3           *NidOptNative `protobuf:"bytes,3,opt" json:"Field3,omitempty"`
	Field4           *NinOptNative `protobuf:"bytes,4,opt" json:"Field4,omitempty"`
	Field6           *uint64       `protobuf:"varint,6,opt" json:"Field6,omitempty"`
	Field7           *int32        `protobuf:"zigzag32,7,opt" json:"Field7,omitempty"`
	Field8           *NidOptNative `protobuf:"bytes,8,opt" json:"Field8,omitempty"`
	Field13          *bool         `protobuf:"varint,13,opt" json:"Field13,omitempty"`
	Field14          *string       `protobuf:"bytes,14,opt" json:"Field14,omitempty"`
	Field15          []byte        `protobuf:"bytes,15,opt" json:"Field15,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NinOptStruct) Reset()         { *m = NinOptStruct{} }
func (m *NinOptStruct) String() string { return proto.CompactTextString(m) }
func (*NinOptStruct) ProtoMessage()    {}

func (m *NinOptStruct) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NinOptStruct) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func (m *NinOptStruct) GetField3() *NidOptNative {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinOptStruct) GetField4() *NinOptNative {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinOptStruct) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return 0
}

func (m *NinOptStruct) GetField7() int32 {
	if m != nil && m.Field7 != nil {
		return *m.Field7
	}
	return 0
}

func (m *NinOptStruct) GetField8() *NidOptNative {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NinOptStruct) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return false
}

func (m *NinOptStruct) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *NinOptStruct) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NidRepStruct struct {
	Field1           []float64       `protobuf:"fixed64,1,rep" json:"Field1,omitempty"`
	Field2           []float32       `protobuf:"fixed32,2,rep" json:"Field2,omitempty"`
	Field3           []*NidOptNative `protobuf:"bytes,3,rep" json:"Field3,omitempty"`
	Field4           []*NinOptNative `protobuf:"bytes,4,rep" json:"Field4,omitempty"`
	Field6           []uint64        `protobuf:"varint,6,rep" json:"Field6,omitempty"`
	Field7           []int32         `protobuf:"zigzag32,7,rep" json:"Field7,omitempty"`
	Field8           []*NidOptNative `protobuf:"bytes,8,rep" json:"Field8,omitempty"`
	Field13          []bool          `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	Field14          []string        `protobuf:"bytes,14,rep" json:"Field14,omitempty"`
	Field15          [][]byte        `protobuf:"bytes,15,rep" json:"Field15,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NidRepStruct) Reset()         { *m = NidRepStruct{} }
func (m *NidRepStruct) String() string { return proto.CompactTextString(m) }
func (*NidRepStruct) ProtoMessage()    {}

func (m *NidRepStruct) GetField1() []float64 {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NidRepStruct) GetField2() []float32 {
	if m != nil {
		return m.Field2
	}
	return nil
}

func (m *NidRepStruct) GetField3() []*NidOptNative {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NidRepStruct) GetField4() []*NinOptNative {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NidRepStruct) GetField6() []uint64 {
	if m != nil {
		return m.Field6
	}
	return nil
}

func (m *NidRepStruct) GetField7() []int32 {
	if m != nil {
		return m.Field7
	}
	return nil
}

func (m *NidRepStruct) GetField8() []*NidOptNative {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NidRepStruct) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

func (m *NidRepStruct) GetField14() []string {
	if m != nil {
		return m.Field14
	}
	return nil
}

func (m *NidRepStruct) GetField15() [][]byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NinRepStruct struct {
	Field1           []float64       `protobuf:"fixed64,1,rep" json:"Field1,omitempty"`
	Field2           []float32       `protobuf:"fixed32,2,rep" json:"Field2,omitempty"`
	Field3           []*NidOptNative `protobuf:"bytes,3,rep" json:"Field3,omitempty"`
	Field4           []*NinOptNative `protobuf:"bytes,4,rep" json:"Field4,omitempty"`
	Field6           []uint64        `protobuf:"varint,6,rep" json:"Field6,omitempty"`
	Field7           []int32         `protobuf:"zigzag32,7,rep" json:"Field7,omitempty"`
	Field8           []*NidOptNative `protobuf:"bytes,8,rep" json:"Field8,omitempty"`
	Field13          []bool          `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	Field14          []string        `protobuf:"bytes,14,rep" json:"Field14,omitempty"`
	Field15          [][]byte        `protobuf:"bytes,15,rep" json:"Field15,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NinRepStruct) Reset()         { *m = NinRepStruct{} }
func (m *NinRepStruct) String() string { return proto.CompactTextString(m) }
func (*NinRepStruct) ProtoMessage()    {}

func (m *NinRepStruct) GetField1() []float64 {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinRepStruct) GetField2() []float32 {
	if m != nil {
		return m.Field2
	}
	return nil
}

func (m *NinRepStruct) GetField3() []*NidOptNative {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinRepStruct) GetField4() []*NinOptNative {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinRepStruct) GetField6() []uint64 {
	if m != nil {
		return m.Field6
	}
	return nil
}

func (m *NinRepStruct) GetField7() []int32 {
	if m != nil {
		return m.Field7
	}
	return nil
}

func (m *NinRepStruct) GetField8() []*NidOptNative {
	if m != nil {
		return m.Field8
	}
	return nil
}

func (m *NinRepStruct) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

func (m *NinRepStruct) GetField14() []string {
	if m != nil {
		return m.Field14
	}
	return nil
}

func (m *NinRepStruct) GetField15() [][]byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NidEmbeddedStruct struct {
	Field1           *NidOptNative `protobuf:"bytes,1,opt" json:"Field1,omitempty"`
	Field200         *NidOptNative `protobuf:"bytes,200,opt" json:"Field200,omitempty"`
	Field210         *bool         `protobuf:"varint,210,opt" json:"Field210,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NidEmbeddedStruct) Reset()         { *m = NidEmbeddedStruct{} }
func (m *NidEmbeddedStruct) String() string { return proto.CompactTextString(m) }
func (*NidEmbeddedStruct) ProtoMessage()    {}

func (m *NidEmbeddedStruct) GetField1() *NidOptNative {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NidEmbeddedStruct) GetField200() *NidOptNative {
	if m != nil {
		return m.Field200
	}
	return nil
}

func (m *NidEmbeddedStruct) GetField210() bool {
	if m != nil && m.Field210 != nil {
		return *m.Field210
	}
	return false
}

type NinEmbeddedStruct struct {
	Field1           *NidOptNative `protobuf:"bytes,1,opt" json:"Field1,omitempty"`
	Field200         *NidOptNative `protobuf:"bytes,200,opt" json:"Field200,omitempty"`
	Field210         *bool         `protobuf:"varint,210,opt" json:"Field210,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NinEmbeddedStruct) Reset()         { *m = NinEmbeddedStruct{} }
func (m *NinEmbeddedStruct) String() string { return proto.CompactTextString(m) }
func (*NinEmbeddedStruct) ProtoMessage()    {}

func (m *NinEmbeddedStruct) GetField1() *NidOptNative {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinEmbeddedStruct) GetField200() *NidOptNative {
	if m != nil {
		return m.Field200
	}
	return nil
}

func (m *NinEmbeddedStruct) GetField210() bool {
	if m != nil && m.Field210 != nil {
		return *m.Field210
	}
	return false
}

type NidNestedStruct struct {
	Field1           *NidOptStruct   `protobuf:"bytes,1,opt" json:"Field1,omitempty"`
	Field2           []*NidRepStruct `protobuf:"bytes,2,rep" json:"Field2,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NidNestedStruct) Reset()         { *m = NidNestedStruct{} }
func (m *NidNestedStruct) String() string { return proto.CompactTextString(m) }
func (*NidNestedStruct) ProtoMessage()    {}

func (m *NidNestedStruct) GetField1() *NidOptStruct {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NidNestedStruct) GetField2() []*NidRepStruct {
	if m != nil {
		return m.Field2
	}
	return nil
}

type NinNestedStruct struct {
	Field1           *NidOptStruct   `protobuf:"bytes,1,opt" json:"Field1,omitempty"`
	Field2           []*NidRepStruct `protobuf:"bytes,2,rep" json:"Field2,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NinNestedStruct) Reset()         { *m = NinNestedStruct{} }
func (m *NinNestedStruct) String() string { return proto.CompactTextString(m) }
func (*NinNestedStruct) ProtoMessage()    {}

func (m *NinNestedStruct) GetField1() *NidOptStruct {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinNestedStruct) GetField2() []*NidRepStruct {
	if m != nil {
		return m.Field2
	}
	return nil
}

type NidOptUuidAsBytes struct {
	Id               []byte `protobuf:"bytes,1,opt" json:"Id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NidOptUuidAsBytes) Reset()         { *m = NidOptUuidAsBytes{} }
func (m *NidOptUuidAsBytes) String() string { return proto.CompactTextString(m) }
func (*NidOptUuidAsBytes) ProtoMessage()    {}

func (m *NidOptUuidAsBytes) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type NinOptUuidAsBytes struct {
	Id               []byte `protobuf:"bytes,1,opt" json:"Id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NinOptUuidAsBytes) Reset()         { *m = NinOptUuidAsBytes{} }
func (m *NinOptUuidAsBytes) String() string { return proto.CompactTextString(m) }
func (*NinOptUuidAsBytes) ProtoMessage()    {}

func (m *NinOptUuidAsBytes) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type NidRepUuidAsBytes struct {
	Id               [][]byte `protobuf:"bytes,1,rep" json:"Id,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NidRepUuidAsBytes) Reset()         { *m = NidRepUuidAsBytes{} }
func (m *NidRepUuidAsBytes) String() string { return proto.CompactTextString(m) }
func (*NidRepUuidAsBytes) ProtoMessage()    {}

func (m *NidRepUuidAsBytes) GetId() [][]byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type NinRepUuidAsBytes struct {
	Id               [][]byte `protobuf:"bytes,1,rep" json:"Id,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NinRepUuidAsBytes) Reset()         { *m = NinRepUuidAsBytes{} }
func (m *NinRepUuidAsBytes) String() string { return proto.CompactTextString(m) }
func (*NinRepUuidAsBytes) ProtoMessage()    {}

func (m *NinRepUuidAsBytes) GetId() [][]byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type NidOptUint128AsBytes struct {
	Value            []byte `protobuf:"bytes,1,opt" json:"Value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NidOptUint128AsBytes) Reset()         { *m = NidOptUint128AsBytes{} }
func (m *NidOptUint128AsBytes) String() string { return proto.CompactTextString(m) }
func (*NidOptUint128AsBytes) ProtoMessage()    {}

func (m *NidOptUint128AsBytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NinOptUint128AsBytes struct {
	Value            []byte `protobuf:"bytes,1,opt" json:"Value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NinOptUint128AsBytes) Reset()         { *m = NinOptUint128AsBytes{} }
func (m *NinOptUint128AsBytes) String() string { return proto.CompactTextString(m) }
func (*NinOptUint128AsBytes) ProtoMessage()    {}

func (m *NinOptUint128AsBytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NidRepUint128AsBytes struct {
	Value            [][]byte `protobuf:"bytes,1,rep" json:"Value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NidRepUint128AsBytes) Reset()         { *m = NidRepUint128AsBytes{} }
func (m *NidRepUint128AsBytes) String() string { return proto.CompactTextString(m) }
func (*NidRepUint128AsBytes) ProtoMessage()    {}

func (m *NidRepUint128AsBytes) GetValue() [][]byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NinRepUint128AsBytes struct {
	Value            [][]byte `protobuf:"bytes,1,rep" json:"Value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NinRepUint128AsBytes) Reset()         { *m = NinRepUint128AsBytes{} }
func (m *NinRepUint128AsBytes) String() string { return proto.CompactTextString(m) }
func (*NinRepUint128AsBytes) ProtoMessage()    {}

func (m *NinRepUint128AsBytes) GetValue() [][]byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NinOptNativeDefault struct {
	Field1           *float64 `protobuf:"fixed64,1,opt,def=1234.1234" json:"Field1,omitempty"`
	Field2           *float32 `protobuf:"fixed32,2,opt,def=1234.1234" json:"Field2,omitempty"`
	Field3           *int32   `protobuf:"varint,3,opt,def=1234" json:"Field3,omitempty"`
	Field4           *int64   `protobuf:"varint,4,opt,def=1234" json:"Field4,omitempty"`
	Field5           *uint32  `protobuf:"varint,5,opt,def=1234" json:"Field5,omitempty"`
	Field6           *uint64  `protobuf:"varint,6,opt,def=1234" json:"Field6,omitempty"`
	Field7           *int32   `protobuf:"zigzag32,7,opt,def=1234" json:"Field7,omitempty"`
	Field8           *int64   `protobuf:"zigzag64,8,opt,def=1234" json:"Field8,omitempty"`
	Field9           *uint32  `protobuf:"fixed32,9,opt,def=1234" json:"Field9,omitempty"`
	Field10          *int32   `protobuf:"fixed32,10,opt,def=1234" json:"Field10,omitempty"`
	Field11          *uint64  `protobuf:"fixed64,11,opt,def=1234" json:"Field11,omitempty"`
	Field12          *int64   `protobuf:"fixed64,12,opt,def=1234" json:"Field12,omitempty"`
	Field13          *bool    `protobuf:"varint,13,opt,def=1" json:"Field13,omitempty"`
	Field14          *string  `protobuf:"bytes,14,opt,def=1234" json:"Field14,omitempty"`
	Field15          []byte   `protobuf:"bytes,15,opt" json:"Field15,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NinOptNativeDefault) Reset()         { *m = NinOptNativeDefault{} }
func (m *NinOptNativeDefault) String() string { return proto.CompactTextString(m) }
func (*NinOptNativeDefault) ProtoMessage()    {}

const Default_NinOptNativeDefault_Field1 float64 = 1234.1234
const Default_NinOptNativeDefault_Field2 float32 = 1234.1234
const Default_NinOptNativeDefault_Field3 int32 = 1234
const Default_NinOptNativeDefault_Field4 int64 = 1234
const Default_NinOptNativeDefault_Field5 uint32 = 1234
const Default_NinOptNativeDefault_Field6 uint64 = 1234
const Default_NinOptNativeDefault_Field7 int32 = 1234
const Default_NinOptNativeDefault_Field8 int64 = 1234
const Default_NinOptNativeDefault_Field9 uint32 = 1234
const Default_NinOptNativeDefault_Field10 int32 = 1234
const Default_NinOptNativeDefault_Field11 uint64 = 1234
const Default_NinOptNativeDefault_Field12 int64 = 1234
const Default_NinOptNativeDefault_Field13 bool = true
const Default_NinOptNativeDefault_Field14 string = "1234"

func (m *NinOptNativeDefault) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return Default_NinOptNativeDefault_Field1
}

func (m *NinOptNativeDefault) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return Default_NinOptNativeDefault_Field2
}

func (m *NinOptNativeDefault) GetField3() int32 {
	if m != nil && m.Field3 != nil {
		return *m.Field3
	}
	return Default_NinOptNativeDefault_Field3
}

func (m *NinOptNativeDefault) GetField4() int64 {
	if m != nil && m.Field4 != nil {
		return *m.Field4
	}
	return Default_NinOptNativeDefault_Field4
}

func (m *NinOptNativeDefault) GetField5() uint32 {
	if m != nil && m.Field5 != nil {
		return *m.Field5
	}
	return Default_NinOptNativeDefault_Field5
}

func (m *NinOptNativeDefault) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return Default_NinOptNativeDefault_Field6
}

func (m *NinOptNativeDefault) GetField7() int32 {
	if m != nil && m.Field7 != nil {
		return *m.Field7
	}
	return Default_NinOptNativeDefault_Field7
}

func (m *NinOptNativeDefault) GetField8() int64 {
	if m != nil && m.Field8 != nil {
		return *m.Field8
	}
	return Default_NinOptNativeDefault_Field8
}

func (m *NinOptNativeDefault) GetField9() uint32 {
	if m != nil && m.Field9 != nil {
		return *m.Field9
	}
	return Default_NinOptNativeDefault_Field9
}

func (m *NinOptNativeDefault) GetField10() int32 {
	if m != nil && m.Field10 != nil {
		return *m.Field10
	}
	return Default_NinOptNativeDefault_Field10
}

func (m *NinOptNativeDefault) GetField11() uint64 {
	if m != nil && m.Field11 != nil {
		return *m.Field11
	}
	return Default_NinOptNativeDefault_Field11
}

func (m *NinOptNativeDefault) GetField12() int64 {
	if m != nil && m.Field12 != nil {
		return *m.Field12
	}
	return Default_NinOptNativeDefault_Field12
}

func (m *NinOptNativeDefault) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return Default_NinOptNativeDefault_Field13
}

func (m *NinOptNativeDefault) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return Default_NinOptNativeDefault_Field14
}

func (m *NinOptNativeDefault) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

func init() {
}
