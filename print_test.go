package easyjson_test

import (
	"testing"
	"fmt"
	"github.com/m0a/easyjson"
)


func TestPrintString1(t *testing.T){
	jsonStr:=`{"a":"abc","b":"bcd"}`
	obj:=easyjson.NewEasyJson(jsonStr)
	str:=fmt.Sprintf("%v",obj)
	if  str != jsonStr {
		t.Fatalf("print:%s,correct:%s don't match",str,jsonStr)
	}
}

func TestPrintString2(t *testing.T){
	jsonStr:=`{"a":"abc","b":[true,false,true,"a","b","c"]}`
	obj:=easyjson.NewEasyJson(jsonStr)
	str:=fmt.Sprintf("%v",obj)
	if  str != jsonStr {
		t.Fatalf("print:%s,correct:%s don't match",str,jsonStr)
	}
}

func TestPrintStringNil(t *testing.T){
	jsonStr:=`null`
	obj:=easyjson.NewEasyJson(jsonStr)
	str:=fmt.Sprintf("%v",obj)
	if  str != jsonStr {
		t.Fatalf("print:%s,correct:%s don't match",str,jsonStr)
	}
}
