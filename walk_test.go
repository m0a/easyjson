package easyjson_test

import (
	"testing"
	"github.com/m0a/easyjson"
//	"fmt"
)


func TestRangeObjects(t *testing.T) {
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

	for k,v:= range obj.RangeObjects() {
//		fmt.Printf("%v:%v\n",k,v)
		switch k.(int) {
		case 0:
			if i,_:=v.AsInt(); i != 0 {
				t.Fatal("not 0")
			}
		}
	}


}
