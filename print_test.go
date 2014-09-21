package easyjson_test

import (
	"fmt"
	"github.com/m0a/easyjson"
	"testing"
)

func TestPrintString1(t *testing.T) {
	jsonStr := `["1","2","3","4","5"]`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}
	str := fmt.Sprintf("%v", obj)
	if str != jsonStr {
		t.Fatalf("print:%s,correct:%s don't match", str, jsonStr)
	}
}

func TestPrintString2(t *testing.T) {
	jsonStr := `{"a":"abc","b":[true,false,true,"a","b","c"]}`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}

	str := fmt.Sprintf("%v", obj)
	if str != jsonStr {
		t.Fatalf("print:%s,correct:%s don't match", str, jsonStr)
	}
}

func TestPrintStringNil(t *testing.T) {
	jsonStr := `["1","2","3",null,"5","6"]`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}
	str := fmt.Sprintf("%v", obj)
	if str != jsonStr {
		t.Fatalf("print:%s,correct:%s don't match", str, jsonStr)
	}
}

func TestPretyPrint1(t *testing.T) {
	jsonStr :=
`[
	"1",
	"2",
	"3",
	null,
	"5",
	"6"
]`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}

	if obj.PretyString() != jsonStr {
		t.Fatalf("don't match \ncurrent: \n%s \ncorrct:\n%s\n",obj.PretyString(),jsonStr)
	}

}

func TestPretyPrint2(t *testing.T) {
	jsonStr :=
`[
	"1",
	[
		"a",
		"b"
	],
	"3"
]`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}

	if obj.PretyString() != jsonStr {
		t.Fatalf("don't match \ncurrent: \n%s \ncorrct:\n%s\n",obj.PretyString(),jsonStr)
	}

}

func TestPretyPrint3(t *testing.T) {
	jsonStr :=
`{
	"1":"a",
	"b":{
		"a":"1",
		"b":"2"
	},
	"3":"c"
}`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}

	if obj.PretyString() != jsonStr {
		t.Fatalf("don't match \ncurrent: \n%s \ncorrct:\n%s\n",obj.PretyString(),jsonStr)
	}

}
