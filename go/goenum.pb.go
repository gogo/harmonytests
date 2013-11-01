// Code generated by protoc-gen-go.
// source: goenum.proto
// DO NOT EDIT!

package gogoprototest

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type TheTestEnum int32

const (
	TheTestEnum_A TheTestEnum = 0
	TheTestEnum_B TheTestEnum = 1
	TheTestEnum_C TheTestEnum = 2
)

var TheTestEnum_name = map[int32]string{
	0: "A",
	1: "B",
	2: "C",
}
var TheTestEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
	"C": 2,
}

func (x TheTestEnum) Enum() *TheTestEnum {
	p := new(TheTestEnum)
	*p = x
	return p
}
func (x TheTestEnum) String() string {
	return proto.EnumName(TheTestEnum_name, int32(x))
}
func (x *TheTestEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TheTestEnum_value, data, "TheTestEnum")
	if err != nil {
		return err
	}
	*x = TheTestEnum(value)
	return nil
}

type AnotherTestEnum int32

const (
	AnotherTestEnum_D AnotherTestEnum = 10
	AnotherTestEnum_E AnotherTestEnum = 11
)

var AnotherTestEnum_name = map[int32]string{
	10: "D",
	11: "E",
}
var AnotherTestEnum_value = map[string]int32{
	"D": 10,
	"E": 11,
}

func (x AnotherTestEnum) Enum() *AnotherTestEnum {
	p := new(AnotherTestEnum)
	*p = x
	return p
}
func (x AnotherTestEnum) String() string {
	return proto.EnumName(AnotherTestEnum_name, int32(x))
}
func (x *AnotherTestEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AnotherTestEnum_value, data, "AnotherTestEnum")
	if err != nil {
		return err
	}
	*x = AnotherTestEnum(value)
	return nil
}

type NidOptEnum struct {
	Field1           *TheTestEnum `protobuf:"varint,1,opt,enum=gogoprototest.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NidOptEnum) Reset()         { *m = NidOptEnum{} }
func (m *NidOptEnum) String() string { return proto.CompactTextString(m) }
func (*NidOptEnum) ProtoMessage()    {}

func (m *NidOptEnum) GetField1() TheTestEnum {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return TheTestEnum_A
}

type NinOptEnum struct {
	Field1           *TheTestEnum `protobuf:"varint,1,opt,enum=gogoprototest.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NinOptEnum) Reset()         { *m = NinOptEnum{} }
func (m *NinOptEnum) String() string { return proto.CompactTextString(m) }
func (*NinOptEnum) ProtoMessage()    {}

func (m *NinOptEnum) GetField1() TheTestEnum {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return TheTestEnum_A
}

type NidRepEnum struct {
	Field1           []TheTestEnum `protobuf:"varint,1,rep,enum=gogoprototest.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NidRepEnum) Reset()         { *m = NidRepEnum{} }
func (m *NidRepEnum) String() string { return proto.CompactTextString(m) }
func (*NidRepEnum) ProtoMessage()    {}

func (m *NidRepEnum) GetField1() []TheTestEnum {
	if m != nil {
		return m.Field1
	}
	return nil
}

type NinRepEnum struct {
	Field1           []TheTestEnum `protobuf:"varint,1,rep,enum=gogoprototest.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NinRepEnum) Reset()         { *m = NinRepEnum{} }
func (m *NinRepEnum) String() string { return proto.CompactTextString(m) }
func (*NinRepEnum) ProtoMessage()    {}

func (m *NinRepEnum) GetField1() []TheTestEnum {
	if m != nil {
		return m.Field1
	}
	return nil
}

type NinOptEnumDefault struct {
	Field1           *TheTestEnum `protobuf:"varint,1,opt,enum=gogoprototest.TheTestEnum,def=2" json:"Field1,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NinOptEnumDefault) Reset()         { *m = NinOptEnumDefault{} }
func (m *NinOptEnumDefault) String() string { return proto.CompactTextString(m) }
func (*NinOptEnumDefault) ProtoMessage()    {}

const Default_NinOptEnumDefault_Field1 TheTestEnum = TheTestEnum_C

func (m *NinOptEnumDefault) GetField1() TheTestEnum {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return Default_NinOptEnumDefault_Field1
}

type AnotherNinOptEnum struct {
	Field1           *AnotherTestEnum `protobuf:"varint,1,opt,enum=gogoprototest.AnotherTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *AnotherNinOptEnum) Reset()         { *m = AnotherNinOptEnum{} }
func (m *AnotherNinOptEnum) String() string { return proto.CompactTextString(m) }
func (*AnotherNinOptEnum) ProtoMessage()    {}

func (m *AnotherNinOptEnum) GetField1() AnotherTestEnum {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return AnotherTestEnum_D
}

type AnotherNinOptEnumDefault struct {
	Field1           *AnotherTestEnum `protobuf:"varint,1,opt,enum=gogoprototest.AnotherTestEnum,def=11" json:"Field1,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *AnotherNinOptEnumDefault) Reset()         { *m = AnotherNinOptEnumDefault{} }
func (m *AnotherNinOptEnumDefault) String() string { return proto.CompactTextString(m) }
func (*AnotherNinOptEnumDefault) ProtoMessage()    {}

const Default_AnotherNinOptEnumDefault_Field1 AnotherTestEnum = AnotherTestEnum_E

func (m *AnotherNinOptEnumDefault) GetField1() AnotherTestEnum {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return Default_AnotherNinOptEnumDefault_Field1
}

func init() {
	proto.RegisterEnum("gogoprototest.TheTestEnum", TheTestEnum_name, TheTestEnum_value)
	proto.RegisterEnum("gogoprototest.AnotherTestEnum", AnotherTestEnum_name, AnotherTestEnum_value)
}
