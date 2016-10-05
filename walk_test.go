package easyjson_test

import (
	"fmt"
	"testing"

	"github.com/m0a/easyjson"
)

func TestRangeObjects(t *testing.T) {
	jsonStr :=
		`[
		"0",
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

	for k, v := range obj.RangeObjects() {
		switch s := k.(int); s {
		case 0:
			if str, _ := v.AsString(); str != "0" {
				t.Fatalf("not \"0\" i=%s\n", str)
			}
		case 5:
			if str, _ := v.AsString(); str != "5" {
				t.Fatalf("not \"5\" i=%s\n", str)
			}

		}
	}
}

func TestListWalk(t *testing.T) {
	jsonStr :=
		`[
		"0",
		"1",
		"2",
		{"3":[1,2,3,4,5,6,7]},
		null,
		"5",
		"6"
	]`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}

	obj.Walk(func(key interface{}, value easyjson.JsonAccessor) {
		fmt.Printf("%v:%v\n", key, value)
		str, _ := value.AsString()
		if key == 0 && str != "0" {
			t.Fatal("key 0 value !=0")
		}
	})
}

func TestDictionaryWalk(t *testing.T) {
	jsonStr :=
		`{
		"0":"a",
		"1":"b",
		"2":"c",
		"3":{"3":[1,2,3,4,5,6,7]},
		"a":"d",
		"5":"e",
		"6":"f"
	}`
	obj, err := easyjson.NewEasyJson(jsonStr)
	if err != nil {
		t.Fatal("json convert err")
	}

	obj.Walk(func(key interface{}, value easyjson.JsonAccessor) {
		fmt.Printf("%v:%v\n", key, value)
		str, _ := value.AsString()
		if key == "0" && str != "a" {
			t.Fatal("key \"0\" value !=\"a\"")
		}
	})
}
