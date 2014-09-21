package easyjson_test

import (
	"fmt"
	"github.com/m0a/easyjson"
	"testing"
)

func TestPrintString1(t *testing.T) {
	jsonStr := `{"a":"abc","b":"bcd"}`
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
	jsonStr := `null`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}

	str := fmt.Sprintf("%v", obj)
	if str != jsonStr {
		t.Fatalf("print:%s,correct:%s don't match", str, jsonStr)
	}
}
