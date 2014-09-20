package easyjson

import (
	"encoding/json"
	"fmt"
	"strings"
	"io"
)

type easyJsonObj struct {
	v interface{}
	err error
}

type Keys []interface{}

// startFunction
func NewEasyJson(i interface{}) easyJsonObj {

	var jo easyJsonObj
	if r,ok:=i.(io.Reader); ok {
		dec := json.NewDecoder(r)
		dec.Decode(&jo.v)
	} else if str, ok := i.(string); ok {
		dec := json.NewDecoder(strings.NewReader(str))
		dec.Decode(&jo.v)
	} else {
		jo.v = i.(easyJsonObj)
	}

	return jo
}

// access function
func (e easyJsonObj) K(keys ...interface{}) (ret easyJsonObj) {

	ret = e
	for _,key:=range keys {
		switch key.(type) {
		case int:
			//array access
			array, ok := ret.v.([]interface{})
			if !ok {
				ret.err=fmt.Errorf("not array please don't use int key.")
				return
			}

			v:=array[key.(int)]
			if v ==nil {
				errStr:=fmt.Sprintf("array access value= nil key=%v array=%v\n",key,array)
				ret.err=fmt.Errorf(errStr)
				return
			}
			ret.v =v
		case string:
			dict, ok := ret.v.(map[string]interface{})
			if !ok {
				ret.err=fmt.Errorf("not dictionary please don't use string key.")
				return
			}

			v:= dict[key.(string)]
			if v ==nil {
				errstr:=fmt.Sprintf("dict access value= nil key=%v dict=%v\n",key,dict)
				ret.err=fmt.Errorf(errstr)
				return
			}
			ret.v = v
		default:
			ret.err=fmt.Errorf("key is only use string,int")
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
