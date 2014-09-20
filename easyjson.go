package easyjson

import (
	"encoding/json"
	"fmt"
	"strings"
)

type easyJsonObj struct {
	v interface{}
	err error
}

type Keys []interface{}

// NewEasyJsonObj
func NewEasyJson(i interface{}) easyJsonObj {

	var jo easyJsonObj
	if str, ok := i.(string); ok {
		dec := json.NewDecoder(strings.NewReader(str))
		dec.Decode(&jo.v)
	} else {
		jo.v = i.(easyJsonObj)
	}

	return jo
}

// K
func (e easyJsonObj) K(keys ...interface{}) (ejo easyJsonObj) {

//	fmt.Printf(			"###start K(keys = [%#v])\n", keys)
//	defer fmt.Printf(	"###end K(keys = [%#v])\n", keys)
	ejo = e
	for _,key:=range keys {
//		fmt.Printf("key==<%#v>\n",key)
		switch key.(type) {
		case int:
			//array access
			array, ok := ejo.v.([]interface{})
			if !ok {
//				fmt.Printf("#try acs array index= %d but err(%#v)\n",key.(int),ejo.v)
				ejo.err=fmt.Errorf("value is not array please don't use int key.")
				return
			}

			v:=array[key.(int)]
			if v ==nil {
				errstr:=fmt.Sprintf("array access but value= nil key=%v array=%v\n",key,array)
//				fmt.Print(errstr)
				ejo.err=fmt.Errorf(errstr)
				return
			}
			ejo.v =v
		case string:
			dict, ok := ejo.v.(map[string]interface{})
			if !ok {
//				fmt.Printf("#try acs dict key= %d but err(%#v)\n",key.(string),ejo.v)
				ejo.err=fmt.Errorf("error v is not dictionary please don't use string key.")
				return 
			}

			v:= dict[key.(string)]
			if v ==nil {
				errstr:=fmt.Sprintf("dict access but value= nil key=%v dict=%v\n",key,dict)
//				fmt.Print(errstr)
				ejo.err=fmt.Errorf(errstr)
				return
			}

//			fmt.Printf("before dict change:  %v\n",ejo.v)
			ejo.v = v
//			fmt.Printf("after dict change:  %v\n",ejo.v)

		default:
			ejo.err=fmt.Errorf("error key can use int or strings only sory..")
			return
		}
	}
	return
}

func panicf(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	panic(str)
}

//値の取得
func (e easyJsonObj) AsInt(k ...interface {}) (i int,err error) {
	v,err:=e.AsFloat64(k...)
	i=int(v)
	return
}

func (e easyJsonObj) AsFloat64(k ...interface {}) (v float64,err error) {
	ejo:=e.K(k...)
	if ejo.err!=nil {
		err = ejo.err
		return
	}

	if !ejo.IsNumber() {
		err=fmt.Errorf("value is not a number type = %#v",ejo.v)
		return
	}
	v=ejo.v.(float64)
	return
}

func (e easyJsonObj) AsString(k ...interface {}) (str string,err error) {
	ejo:=e.K(k...)
	if ejo.err!=nil {
		err = ejo.err
		return
	}

	if !ejo.IsString() {
		err=fmt.Errorf("value is not a string type = %#v",ejo.v)
		return
	}

	str=e.K(k...).v.(string)
	return
}

// Value check
func (e easyJsonObj) IsDict(k ...interface {}) bool {
	_, ok := e.K(k...).v.(map[string]interface{})
	return ok
}

func (e easyJsonObj) IsArray(k ...interface {}) bool {

	_, ok := e.K(k...).v.([]interface{})
	return ok
}

func (e easyJsonObj) IsNumber(k ...interface {}) bool {
	_, ok := e.K(k...).v.(float64)
	return ok
}

func (e easyJsonObj) IsString(k ...interface {}) bool {
	_, ok := e.K(k...).v.(string)
	return ok
}
