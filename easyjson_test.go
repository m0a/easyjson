package easyjson

import (
	"fmt"
	"reflect"
	"testing"
)


func TestNewEasyJson(t *testing.T) {

	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5]
	}`

	testobj := NewEasyJson(jsonString)
	testobj.Debug=true
	if testobj.v == nil {
		t.Fatal("NewEasyJson Error")
	}

}

func TestK00(t *testing.T) {
	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5]
	}`

	jso := NewEasyJson(jsonString)
	jso.Debug=true
	if v, _ := jso.K("key2").K(1).AsInt(); v != 2 {
		t.Fatal("convert int err")
	}

}

func TestK01(t *testing.T) {
	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5]
	}`
	jso := NewEasyJson(jsonString)
	jso = jso.K("key2", 0)
	if jso.v != "a" {
		t.Fatalf("value:%v correct:'a'\n", jso.v)
	}
}

func TestK02(t *testing.T) {
	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,{"a":1,"b":2},5]
	}`

	jso := NewEasyJson(jsonString)
	if v, err := jso.K("key2", 3, "a").AsInt(); err != nil || v != 1 {
		t.Fatalf("value:%v,err:%s correct:1\n", jso.v, err)
	}
}

func TestK03(t *testing.T) {

	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5],
	"key3":{"a":1,"b":2,
		"c":[0,1,2,3,4,5]}
	}`

	jso := NewEasyJson(jsonString)
	keys := Keys{"key3", "c", 4}
	num, err := jso.K(keys...).AsInt()
	if err != nil || num != 4 {
		t.Fatalf("num=%v err:%s correct:4\n", num, err, jso.err)
	}
}

func TestK04(t *testing.T) {

	const jsonString string = `[
		"value",
		["a",2,3,4,5],
		{"a":1,"b":2,"c":[0,1,2,3,4,5]}
	]`

	jso := NewEasyJson(jsonString)
	keys := Keys{2, "c", 4}
	num, err := jso.K(keys...).AsInt()
	if err != nil || num != 4 {
		t.Fatalf("num=%v err:%s correct:4\n", num, err, jso.err)
	}
}

func TestAsInt00(t *testing.T) {
	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,{"a":1,"b":2},5]
	}`
	jso := NewEasyJson(jsonString)
	if v, err := jso.AsInt("key2", 3, "a"); err != nil || v != 1 {
		t.Fatalf("value:%v,err:%s correct:1\n", jso.v, err)
	}
}

func TestKMultiAccess(t *testing.T) {
	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5]
	}`

	jso := NewEasyJson(jsonString)
	if v, _ := jso.K("key2", 1).AsInt(); v != 2 {
		t.Fatal("convert int err")
	}
}

func TestAsString(t *testing.T) {
	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5]
	}`

	jso := NewEasyJson(jsonString)
	str, _ := jso.K("key1").AsString()
	fmt.Println("str =", str)
	if reflect.TypeOf(str).Kind() != reflect.String {
		t.Fatalf("str is not string current type is %s ", reflect.TypeOf(str).Kind().String())
	}
}

func TestAsInt(t *testing.T) {

	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5]
	}`

	jso := NewEasyJson(jsonString)
	v, err := jso.K("key2").K(1).AsInt()
	if err != nil {
		t.Fatalf("err:%s", err)
	}

	if v != 2 {
		t.Fatalf("value is %d correct:2 ", v)
	}
}

const jsonString2 string = `{
	"key1":"value",
	"key2":["a",2,3,4,5],
	"key3":{"a":1,"b":2,
		"c":[0,1,2,3,4,5]}
	}`

func TestUseAsInt2(t *testing.T) {

	jso := NewEasyJson(jsonString2)
	keys := Keys{"key2", 3}
	num, _ := jso.AsInt(keys...)

	if !jso.IsNumber(keys...) {
		t.Fatalf("%v is not number ", num)
	}
	fmt.Printf("jso.asInt(%v)=%d\n", keys, num)
}

func TestUseAsFloat64(t *testing.T) {

	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5],
	"key3":{"a":1,"b":2,
		"c":[0,1,2,3,4,5]}
	}`

	jso := NewEasyJson(jsonString)
	keys := Keys{"key3", "c", 4}
	num, _ := jso.AsFloat64(keys...)
	if num != 4.0 {
		t.Fatalf("num!=4 current num = %v err=%v\n", num, jso.err)
	}
}

func TestIsXXX(t *testing.T) {
	const jsonString string = `{
	"key1":"value",
	"key2":["a",2,3,4,5],
	"key3":{"a":1,"b":2,"c":[0,1,2,3,4,5],"d":10,"e":6},
	"key4":true,
	}`


	jso := NewEasyJson(jsonString)

	checkList := map[string]Keys{
		"string"	: Keys{"key1"},
		"dict"		: Keys{"key3"},
		"array"		: Keys{"key3", "c"},
		"number"	: Keys{"key3", "c", 3},
		"bool"		: Keys{"key4"},

	}

	for v, k := range checkList {
		switch v {
		case "string":
			if !jso.IsString(k...) {
				t.Fatalf("key:%v value is not %s(%v)\n", k,v, jso.K(k...).v)
			}
		case "array":
			if !jso.IsArray(k...) {
				t.Fatalf("key:%v value is not %s(%v)\n", k, v, jso.K(k...).v)
			}
		case "number":
			if !jso.IsNumber(k...) {
				t.Fatalf("key:%v value is not %s(%v)\n", k, v, jso.K(k...).v)
			}
		case "dict":
			if !jso.IsDict(k...) {
				t.Fatalf("key:%v value is not %s(%v)\n", k, v, jso.K(k...).v)
			}
		case "bool":
			if !jso.IsBool(k...) {
				t.Fatalf("key:%v value is not %s(%v)\n", k, v, jso.K(k...).v)
			}
		}
	}

}
