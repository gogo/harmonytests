// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf/gogoproto
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package gogoprototest

import (
	gogoproto "code.google.com/p/gogoprotobuf/proto"
	gotest "code.google.com/p/gogoprototest/go"
	gogotest "code.google.com/p/gogoprototest/gogo"
	goproto "code.google.com/p/goprotobuf/proto"
	"testing"
)

func TestDefaultProtoVimAnotherNidGoOptEnum(t *testing.T) {
	data, err := goproto.Marshal(govimAnotherNidOptEnum)
	if err != nil {
		panic(err)
	}
	msg := &gotest.AnotherNidOptEnum{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if govimAnotherNidOptEnum.GetField1() != msg.GetField1() {
		t.Fatalf("expect %v got %v", govimAnotherNidOptEnum.GetField1(), msg.GetField1())
	}
}

func TestDefaultProtoDefAnotherNidGoOptEnum(t *testing.T) {
	data, err := goproto.Marshal(godefAnotherNidOptEnum)
	if err != nil {
		panic(err)
	}
	msg := &gotest.AnotherNidOptEnum{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if godefAnotherNidOptEnum.GetField1() != msg.GetField1() {
		t.Fatalf("expect %v got %v", godefAnotherNidOptEnum.GetField1(), msg.GetField1())
	}
}

func TestDefaultProtoVimAnotherNidGoGoOptEnum(t *testing.T) {
	data, err := gogoproto.Marshal(gogovimAnotherNidOptEnum)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.AnotherNidOptEnum{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogovimAnotherNidOptEnum.Field1 != msg.Field1 {
		t.Fatalf("expect %v got %v", gogovimAnotherNidOptEnum.Field1, msg.Field1)
	}
}

func TestDefaultProtoDefAnotherNidGoGoOptEnum(t *testing.T) {
	data, err := gogoproto.Marshal(gogodefAnotherNidOptEnum)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.AnotherNidOptEnum{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogodefAnotherNidOptEnum.Field1 != msg.Field1 {
		t.Fatalf("expect %v got %v", gogodefAnotherNidOptEnum.Field1, msg.Field1)
	}
}

func TestDefaultProtoVimAnotherNidGoToGoGoOptEnum(t *testing.T) {
	data, err := goproto.Marshal(govimAnotherNidOptEnum)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.AnotherNidOptEnum{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(govimAnotherNidOptEnum.GetField1()) != int32(msg.Field1) {
		t.Fatalf("expect %v got %v", govimAnotherNidOptEnum.GetField1(), msg.Field1)
	}
	t.Logf("expect %v got %v", govimAnotherNidOptEnum.GetField1(), msg.Field1)
}

func TestDefaultProtoDefAnotherNidGoToGoGoOptEnum(t *testing.T) {
	data, err := goproto.Marshal(godefAnotherNidOptEnum)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.AnotherNidOptEnum{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(godefAnotherNidOptEnum.GetField1()) != int32(msg.Field1) {
		t.Fatalf("expect %v got %v", godefAnotherNidOptEnum.GetField1(), msg.Field1)
	}
	t.Logf("expect %v got %v", godefAnotherNidOptEnum.GetField1(), msg.Field1)
}

func TestDefaultProtoVimAnotherNidGoGoToGoOptEnum(t *testing.T) {
	data, err := gogoproto.Marshal(gogovimAnotherNidOptEnum)
	if err != nil {
		panic(err)
	}
	msg := &gotest.AnotherNidOptEnum{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(gogovimAnotherNidOptEnum.Field1) != int32(msg.GetField1()) {
		t.Fatalf("expect %v got %v", gogovimAnotherNidOptEnum.Field1, msg.GetField1())
	}
	t.Logf("expect %v got %v", gogovimAnotherNidOptEnum.Field1, msg.GetField1())
}

func TestDefaultProtoDefAnotherNidGoGoToGoOptEnum(t *testing.T) {
	data, err := gogoproto.Marshal(gogodefAnotherNidOptEnum)
	if err != nil {
		panic(err)
	}
	msg := &gotest.AnotherNidOptEnum{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(gogodefAnotherNidOptEnum.Field1) != int32(msg.GetField1()) {
		t.Fatalf("expect %v got %v", gogodefAnotherNidOptEnum.Field1, msg.GetField1())
	}
	t.Logf("expect %v got %v", gogodefAnotherNidOptEnum.Field1, msg.GetField1())
}

func TestDefaultProtoVimAnotherNidGoOptEnumDefault(t *testing.T) {
	data, err := goproto.Marshal(govimAnotherNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.AnotherNidOptEnumDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if govimAnotherNidOptEnumDefault.GetField1() != msg.GetField1() {
		t.Fatalf("expect %v got %v", govimAnotherNidOptEnumDefault.GetField1(), msg.GetField1())
	}
}

func TestDefaultProtoDefAnotherNidGoOptEnumDefault(t *testing.T) {
	data, err := goproto.Marshal(godefAnotherNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.AnotherNidOptEnumDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if godefAnotherNidOptEnumDefault.GetField1() != msg.GetField1() {
		t.Fatalf("expect %v got %v", godefAnotherNidOptEnumDefault.GetField1(), msg.GetField1())
	}
}

func TestDefaultProtoVimAnotherNidGoGoOptEnumDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogovimAnotherNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.AnotherNidOptEnumDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogovimAnotherNidOptEnumDefault.Field1 != msg.Field1 {
		t.Fatalf("expect %v got %v", gogovimAnotherNidOptEnumDefault.Field1, msg.Field1)
	}
}

func TestDefaultProtoDefAnotherNidGoGoOptEnumDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogodefAnotherNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.AnotherNidOptEnumDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogodefAnotherNidOptEnumDefault.Field1 != msg.Field1 {
		t.Fatalf("expect %v got %v", gogodefAnotherNidOptEnumDefault.Field1, msg.Field1)
	}
}

func TestDefaultProtoVimAnotherNidGoToGoGoOptEnumDefault(t *testing.T) {
	data, err := goproto.Marshal(govimAnotherNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.AnotherNidOptEnumDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(govimAnotherNidOptEnumDefault.GetField1()) != int32(msg.Field1) {
		t.Fatalf("expect %v got %v", govimAnotherNidOptEnumDefault.GetField1(), msg.Field1)
	}
	t.Logf("expect %v got %v", govimAnotherNidOptEnumDefault.GetField1(), msg.Field1)
}

func TestDefaultProtoDefAnotherNidGoToGoGoOptEnumDefault(t *testing.T) {
	data, err := goproto.Marshal(godefAnotherNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.AnotherNidOptEnumDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(godefAnotherNidOptEnumDefault.GetField1()) != int32(msg.Field1) {
		t.Fatalf("expect %v got %v", godefAnotherNidOptEnumDefault.GetField1(), msg.Field1)
	}
	t.Logf("expect %v got %v", godefAnotherNidOptEnumDefault.GetField1(), msg.Field1)
}

func TestDefaultProtoVimAnotherNidGoGoToGoOptEnumDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogovimAnotherNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.AnotherNidOptEnumDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(gogovimAnotherNidOptEnumDefault.Field1) != int32(msg.GetField1()) {
		t.Fatalf("expect %v got %v", gogovimAnotherNidOptEnumDefault.Field1, msg.GetField1())
	}
	t.Logf("expect %v got %v", gogovimAnotherNidOptEnumDefault.Field1, msg.GetField1())
}

func TestDefaultProtoDefAnotherNidGoGoToGoOptEnumDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogodefAnotherNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.AnotherNidOptEnumDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(gogodefAnotherNidOptEnumDefault.Field1) != int32(msg.GetField1()) {
		t.Fatalf("expect %v got %v", gogodefAnotherNidOptEnumDefault.Field1, msg.GetField1())
	}
	t.Logf("expect %v got %v", gogodefAnotherNidOptEnumDefault.Field1, msg.GetField1())
}

func TestDefaultProtoVimNidGoOptEnumDefault(t *testing.T) {
	data, err := goproto.Marshal(govimNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.NidOptEnumDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if govimNidOptEnumDefault.GetField1() != msg.GetField1() {
		t.Fatalf("expect %v got %v", govimNidOptEnumDefault.GetField1(), msg.GetField1())
	}
}

func TestDefaultProtoDefNidGoOptEnumDefault(t *testing.T) {
	data, err := goproto.Marshal(godefNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.NidOptEnumDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if godefNidOptEnumDefault.GetField1() != msg.GetField1() {
		t.Fatalf("expect %v got %v", godefNidOptEnumDefault.GetField1(), msg.GetField1())
	}
}

func TestDefaultProtoVimNidGoGoOptEnumDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogovimNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.NidOptEnumDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogovimNidOptEnumDefault.Field1 != msg.Field1 {
		t.Fatalf("expect %v got %v", gogovimNidOptEnumDefault.Field1, msg.Field1)
	}
}

func TestDefaultProtoDefNidGoGoOptEnumDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogodefNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.NidOptEnumDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogodefNidOptEnumDefault.Field1 != msg.Field1 {
		t.Fatalf("expect %v got %v", gogodefNidOptEnumDefault.Field1, msg.Field1)
	}
}

func TestDefaultProtoVimNidGoToGoGoOptEnumDefault(t *testing.T) {
	data, err := goproto.Marshal(govimNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.NidOptEnumDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(govimNidOptEnumDefault.GetField1()) != int32(msg.Field1) {
		t.Fatalf("expect %v got %v", govimNidOptEnumDefault.GetField1(), msg.Field1)
	}
	t.Logf("expect %v got %v", govimNidOptEnumDefault.GetField1(), msg.Field1)
}

func TestDefaultProtoDefNidGoToGoGoOptEnumDefault(t *testing.T) {
	data, err := goproto.Marshal(godefNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.NidOptEnumDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(godefNidOptEnumDefault.GetField1()) != int32(msg.Field1) {
		t.Fatalf("expect %v got %v", godefNidOptEnumDefault.GetField1(), msg.Field1)
	}
	t.Logf("expect %v got %v", godefNidOptEnumDefault.GetField1(), msg.Field1)
}

func TestDefaultProtoVimNidGoGoToGoOptEnumDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogovimNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.NidOptEnumDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(gogovimNidOptEnumDefault.Field1) != int32(msg.GetField1()) {
		t.Fatalf("expect %v got %v", gogovimNidOptEnumDefault.Field1, msg.GetField1())
	}
	t.Logf("expect %v got %v", gogovimNidOptEnumDefault.Field1, msg.GetField1())
}

func TestDefaultProtoDefNidGoGoToGoOptEnumDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogodefNidOptEnumDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.NidOptEnumDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if int32(gogodefNidOptEnumDefault.Field1) != int32(msg.GetField1()) {
		t.Fatalf("expect %v got %v", gogodefNidOptEnumDefault.Field1, msg.GetField1())
	}
	t.Logf("expect %v got %v", gogodefNidOptEnumDefault.Field1, msg.GetField1())
}

func TestDefaultProtoVimNidGoOptNativeDefault(t *testing.T) {
	data, err := goproto.Marshal(govimNidOptNativeDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.NidOptNativeDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if govimNidOptNativeDefault.GetField1() != msg.GetField1() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField1(), msg.GetField1())
	}
	if govimNidOptNativeDefault.GetField2() != msg.GetField2() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField2(), msg.GetField2())
	}
	if govimNidOptNativeDefault.GetField3() != msg.GetField3() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField3(), msg.GetField3())
	}
	if govimNidOptNativeDefault.GetField4() != msg.GetField4() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField4(), msg.GetField4())
	}
	if govimNidOptNativeDefault.GetField5() != msg.GetField5() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField5(), msg.GetField5())
	}
	if govimNidOptNativeDefault.GetField6() != msg.GetField6() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField6(), msg.GetField6())
	}
	if govimNidOptNativeDefault.GetField7() != msg.GetField7() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField7(), msg.GetField7())
	}
	if govimNidOptNativeDefault.GetField8() != msg.GetField8() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField8(), msg.GetField8())
	}
	if govimNidOptNativeDefault.GetField9() != msg.GetField9() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField9(), msg.GetField9())
	}
	if govimNidOptNativeDefault.GetField10() != msg.GetField10() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField10(), msg.GetField10())
	}
	if govimNidOptNativeDefault.GetField11() != msg.GetField11() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField11(), msg.GetField11())
	}
	if govimNidOptNativeDefault.GetField12() != msg.GetField12() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField12(), msg.GetField12())
	}
	if govimNidOptNativeDefault.GetField13() != msg.GetField13() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField13(), msg.GetField13())
	}
	if govimNidOptNativeDefault.GetField14() != msg.GetField14() {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField14(), msg.GetField14())
	}
}

func TestDefaultProtoDefNidGoOptNativeDefault(t *testing.T) {
	data, err := goproto.Marshal(godefNidOptNativeDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.NidOptNativeDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if godefNidOptNativeDefault.GetField1() != msg.GetField1() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField1(), msg.GetField1())
	}
	if godefNidOptNativeDefault.GetField2() != msg.GetField2() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField2(), msg.GetField2())
	}
	if godefNidOptNativeDefault.GetField3() != msg.GetField3() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField3(), msg.GetField3())
	}
	if godefNidOptNativeDefault.GetField4() != msg.GetField4() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField4(), msg.GetField4())
	}
	if godefNidOptNativeDefault.GetField5() != msg.GetField5() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField5(), msg.GetField5())
	}
	if godefNidOptNativeDefault.GetField6() != msg.GetField6() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField6(), msg.GetField6())
	}
	if godefNidOptNativeDefault.GetField7() != msg.GetField7() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField7(), msg.GetField7())
	}
	if godefNidOptNativeDefault.GetField8() != msg.GetField8() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField8(), msg.GetField8())
	}
	if godefNidOptNativeDefault.GetField9() != msg.GetField9() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField9(), msg.GetField9())
	}
	if godefNidOptNativeDefault.GetField10() != msg.GetField10() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField10(), msg.GetField10())
	}
	if godefNidOptNativeDefault.GetField11() != msg.GetField11() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField11(), msg.GetField11())
	}
	if godefNidOptNativeDefault.GetField12() != msg.GetField12() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField12(), msg.GetField12())
	}
	if godefNidOptNativeDefault.GetField13() != msg.GetField13() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField13(), msg.GetField13())
	}
	if godefNidOptNativeDefault.GetField14() != msg.GetField14() {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField14(), msg.GetField14())
	}
}

func TestDefaultProtoVimNidGoGoOptNativeDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogovimNidOptNativeDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.NidOptNativeDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogovimNidOptNativeDefault.Field1 != msg.Field1 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field1, msg.Field1)
	}
	if gogovimNidOptNativeDefault.Field2 != msg.Field2 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field2, msg.Field2)
	}
	if gogovimNidOptNativeDefault.Field3 != msg.Field3 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field3, msg.Field3)
	}
	if gogovimNidOptNativeDefault.Field4 != msg.Field4 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field4, msg.Field4)
	}
	if gogovimNidOptNativeDefault.Field5 != msg.Field5 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field5, msg.Field5)
	}
	if gogovimNidOptNativeDefault.Field6 != msg.Field6 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field6, msg.Field6)
	}
	if gogovimNidOptNativeDefault.Field7 != msg.Field7 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field7, msg.Field7)
	}
	if gogovimNidOptNativeDefault.Field8 != msg.Field8 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field8, msg.Field8)
	}
	if gogovimNidOptNativeDefault.Field9 != msg.Field9 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field9, msg.Field9)
	}
	if gogovimNidOptNativeDefault.Field10 != msg.Field10 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field10, msg.Field10)
	}
	if gogovimNidOptNativeDefault.Field11 != msg.Field11 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field11, msg.Field11)
	}
	if gogovimNidOptNativeDefault.Field12 != msg.Field12 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field12, msg.Field12)
	}
	if gogovimNidOptNativeDefault.Field13 != msg.Field13 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field13, msg.Field13)
	}
	if gogovimNidOptNativeDefault.Field14 != msg.Field14 {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field14, msg.Field14)
	}
}

func TestDefaultProtoDefNidGoGoOptNativeDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogodefNidOptNativeDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.NidOptNativeDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogodefNidOptNativeDefault.Field1 != msg.Field1 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field1, msg.Field1)
	}
	if gogodefNidOptNativeDefault.Field2 != msg.Field2 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field2, msg.Field2)
	}
	if gogodefNidOptNativeDefault.Field3 != msg.Field3 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field3, msg.Field3)
	}
	if gogodefNidOptNativeDefault.Field4 != msg.Field4 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field4, msg.Field4)
	}
	if gogodefNidOptNativeDefault.Field5 != msg.Field5 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field5, msg.Field5)
	}
	if gogodefNidOptNativeDefault.Field6 != msg.Field6 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field6, msg.Field6)
	}
	if gogodefNidOptNativeDefault.Field7 != msg.Field7 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field7, msg.Field7)
	}
	if gogodefNidOptNativeDefault.Field8 != msg.Field8 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field8, msg.Field8)
	}
	if gogodefNidOptNativeDefault.Field9 != msg.Field9 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field9, msg.Field9)
	}
	if gogodefNidOptNativeDefault.Field10 != msg.Field10 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field10, msg.Field10)
	}
	if gogodefNidOptNativeDefault.Field11 != msg.Field11 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field11, msg.Field11)
	}
	if gogodefNidOptNativeDefault.Field12 != msg.Field12 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field12, msg.Field12)
	}
	if gogodefNidOptNativeDefault.Field13 != msg.Field13 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field13, msg.Field13)
	}
	if gogodefNidOptNativeDefault.Field14 != msg.Field14 {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field14, msg.Field14)
	}
}

func TestDefaultProtoVimNidGoToGoGoOptNativeDefault(t *testing.T) {
	data, err := goproto.Marshal(govimNidOptNativeDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.NidOptNativeDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if govimNidOptNativeDefault.GetField1() != msg.Field1 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField1(), msg.Field1)
	}
	if govimNidOptNativeDefault.GetField2() != msg.Field2 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField2(), msg.Field2)
	}
	if govimNidOptNativeDefault.GetField3() != msg.Field3 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField3(), msg.Field3)
	}
	if govimNidOptNativeDefault.GetField4() != msg.Field4 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField4(), msg.Field4)
	}
	if govimNidOptNativeDefault.GetField5() != msg.Field5 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField5(), msg.Field5)
	}
	if govimNidOptNativeDefault.GetField6() != msg.Field6 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField6(), msg.Field6)
	}
	if govimNidOptNativeDefault.GetField7() != msg.Field7 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField7(), msg.Field7)
	}
	if govimNidOptNativeDefault.GetField8() != msg.Field8 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField8(), msg.Field8)
	}
	if govimNidOptNativeDefault.GetField9() != msg.Field9 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField9(), msg.Field9)
	}
	if govimNidOptNativeDefault.GetField10() != msg.Field10 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField10(), msg.Field10)
	}
	if govimNidOptNativeDefault.GetField11() != msg.Field11 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField11(), msg.Field11)
	}
	if govimNidOptNativeDefault.GetField12() != msg.Field12 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField12(), msg.Field12)
	}
	if govimNidOptNativeDefault.GetField13() != msg.Field13 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField1(), msg.Field13)
	}
	if govimNidOptNativeDefault.GetField14() != msg.Field14 {
		t.Fatalf("expect %v got %v", govimNidOptNativeDefault.GetField1(), msg.Field14)
	}
}

func TestDefaultProtoDefNidGoToGoGoOptNativeDefault(t *testing.T) {
	data, err := goproto.Marshal(godefNidOptNativeDefault)
	if err != nil {
		panic(err)
	}
	msg := &gogotest.NidOptNativeDefault{}
	if err := gogoproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if godefNidOptNativeDefault.GetField1() != msg.Field1 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField1(), msg.Field1)
	}
	if godefNidOptNativeDefault.GetField2() != msg.Field2 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField2(), msg.Field2)
	}
	if godefNidOptNativeDefault.GetField3() != msg.Field3 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField3(), msg.Field3)
	}
	if godefNidOptNativeDefault.GetField4() != msg.Field4 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField4(), msg.Field4)
	}
	if godefNidOptNativeDefault.GetField5() != msg.Field5 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField5(), msg.Field5)
	}
	if godefNidOptNativeDefault.GetField6() != msg.Field6 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField6(), msg.Field6)
	}
	if godefNidOptNativeDefault.GetField7() != msg.Field7 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField7(), msg.Field7)
	}
	if godefNidOptNativeDefault.GetField8() != msg.Field8 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField8(), msg.Field8)
	}
	if godefNidOptNativeDefault.GetField9() != msg.Field9 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField9(), msg.Field9)
	}
	if godefNidOptNativeDefault.GetField10() != msg.Field10 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField10(), msg.Field10)
	}
	if godefNidOptNativeDefault.GetField11() != msg.Field11 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField11(), msg.Field11)
	}
	if godefNidOptNativeDefault.GetField12() != msg.Field12 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField12(), msg.Field12)
	}
	if godefNidOptNativeDefault.GetField13() != msg.Field13 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField1(), msg.Field13)
	}
	if godefNidOptNativeDefault.GetField14() != msg.Field14 {
		t.Fatalf("expect %v got %v", godefNidOptNativeDefault.GetField1(), msg.Field14)
	}
}

func TestDefaultProtoVimNidGoGoToGoOptNativeDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogovimNidOptNativeDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.NidOptNativeDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogovimNidOptNativeDefault.Field1 != msg.GetField1() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field1, msg.GetField1())
	}
	if gogovimNidOptNativeDefault.Field2 != msg.GetField2() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field2, msg.GetField2())
	}
	if gogovimNidOptNativeDefault.Field3 != msg.GetField3() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field3, msg.GetField3())
	}
	if gogovimNidOptNativeDefault.Field4 != msg.GetField4() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field4, msg.GetField4())
	}
	if gogovimNidOptNativeDefault.Field5 != msg.GetField5() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field5, msg.GetField5())
	}
	if gogovimNidOptNativeDefault.Field6 != msg.GetField6() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field6, msg.GetField6())
	}
	if gogovimNidOptNativeDefault.Field7 != msg.GetField7() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field7, msg.GetField7())
	}
	if gogovimNidOptNativeDefault.Field8 != msg.GetField8() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field8, msg.GetField8())
	}
	if gogovimNidOptNativeDefault.Field9 != msg.GetField9() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field9, msg.GetField9())
	}
	if gogovimNidOptNativeDefault.Field10 != msg.GetField10() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field10, msg.GetField10())
	}
	if gogovimNidOptNativeDefault.Field11 != msg.GetField11() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field11, msg.GetField11())
	}
	if gogovimNidOptNativeDefault.Field12 != msg.GetField12() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field12, msg.GetField12())
	}
	if gogovimNidOptNativeDefault.Field13 != msg.GetField13() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field13, msg.GetField13())
	}
	if gogovimNidOptNativeDefault.Field14 != msg.GetField14() {
		t.Fatalf("expect %v got %v", gogovimNidOptNativeDefault.Field14, msg.GetField14())
	}
}

func TestDefaultProtoDefNidGoGoToGoOptNativeDefault(t *testing.T) {
	data, err := gogoproto.Marshal(gogodefNidOptNativeDefault)
	if err != nil {
		panic(err)
	}
	msg := &gotest.NidOptNativeDefault{}
	if err := goproto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if gogodefNidOptNativeDefault.Field1 != msg.GetField1() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field1, msg.GetField1())
	}
	if gogodefNidOptNativeDefault.Field2 != msg.GetField2() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field2, msg.GetField2())
	}
	if gogodefNidOptNativeDefault.Field3 != msg.GetField3() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field3, msg.GetField3())
	}
	if gogodefNidOptNativeDefault.Field4 != msg.GetField4() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field4, msg.GetField4())
	}
	if gogodefNidOptNativeDefault.Field5 != msg.GetField5() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field5, msg.GetField5())
	}
	if gogodefNidOptNativeDefault.Field6 != msg.GetField6() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field6, msg.GetField6())
	}
	if gogodefNidOptNativeDefault.Field7 != msg.GetField7() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field7, msg.GetField7())
	}
	if gogodefNidOptNativeDefault.Field8 != msg.GetField8() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field8, msg.GetField8())
	}
	if gogodefNidOptNativeDefault.Field9 != msg.GetField9() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field9, msg.GetField9())
	}
	if gogodefNidOptNativeDefault.Field10 != msg.GetField10() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field10, msg.GetField10())
	}
	if gogodefNidOptNativeDefault.Field11 != msg.GetField11() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field11, msg.GetField11())
	}
	if gogodefNidOptNativeDefault.Field12 != msg.GetField12() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field12, msg.GetField12())
	}
	if gogodefNidOptNativeDefault.Field13 != msg.GetField13() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field13, msg.GetField13())
	}
	if gogodefNidOptNativeDefault.Field14 != msg.GetField14() {
		t.Fatalf("expect %v got %v", gogodefNidOptNativeDefault.Field14, msg.GetField14())
	}
}
