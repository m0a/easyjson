package easyjson

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"errors"
)

type easyJsonObj struct {
	v     interface{}
	err   error
	Debug bool
}

type Keys []interface{}

// startFunction
func NewEasyJson(i interface{}) easyJsonObj {

	var jo easyJsonObj
	if r, ok := i.(io.Reader); ok {
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

	errorlog := func(format string, a ...interface{}) {
		e.err = fmt.Errorf(format, a...)
	}
	debuglog := func(format string, a ...interface{}) {
		_=format
		_=a
	}
	if e.Debug {
		errorlog = func(format string, a ...interface{}) {
			str:=fmt.Sprintf(format,a...)
			e.err = errors.New(str)
			fmt.Print(str)
			panic(str)
		}
		debuglog = func(format string, a ...interface{}) {
			fmt.Printf(format, a...)
		}
	}

	debuglog(">>start K(%v)\n", keys)
	defer debuglog(">>end K(%v)\n", keys)

	ret = e
	for _, key := range keys {
		switch key.(type) {
		case int:
			//array access
			array, ok := ret.v.([]interface{})
			if !ok {
				//				ret.err=fmt.Errorf()
				errorlog("Not array. please use string key.")
				return
			}

			v := array[key.(int)]
			if v == nil {
				return
			}
			ret.v = v
		case string:
			dict, ok := ret.v.(map[string]interface{})
			if !ok {
				errorlog("Not dictionary. please use int key.")
				return
			}

			v := dict[key.(string)]
			if v == nil {
				errorlog("dict access value= nil key=%v dict=%v\n", key, dict)
				return
			}
			ret.v = v
		default:
			errorlog("key is only use string,int")
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
func (e easyJsonObj) AsInt(k ...interface{}) (i int, err error) {
	v, err := e.AsFloat64(k...)
	i = int(v)
	return
}

func (e easyJsonObj) AsFloat64(k ...interface{}) (v float64, err error) {
	ejo := e.K(k...)
	if ejo.err != nil {
		err = ejo.err
		return
	}

	if !ejo.IsNumber() {
		err = fmt.Errorf("value is not a number type = %#v", ejo.v)
		return
	}
	v = ejo.v.(float64)
	return
}

func (e easyJsonObj) AsString(k ...interface{}) (str string, err error) {
	ejo := e.K(k...)
	if ejo.err != nil {
		err = ejo.err
		return
	}

	if !ejo.IsString() {
		err = fmt.Errorf("value is not a string type = %#v", ejo.v)
		return
	}

	str = e.K(k...).v.(string)
	return
}

// Value check
func (e easyJsonObj) IsDict(k ...interface{}) bool {
	_, ok := e.K(k...).v.(map[string]interface{})
	return ok
}

func (e easyJsonObj) IsArray(k ...interface{}) bool {

	_, ok := e.K(k...).v.([]interface{})
	return ok
}

func (e easyJsonObj) IsBool(k ...interface{}) bool {
	_, ok := e.K(k...).v.(bool)
	return ok
}

func (e easyJsonObj) IsNumber(k ...interface{}) bool {
	_, ok := e.K(k...).v.(float64)
	return ok
}

func (e easyJsonObj) IsString(k ...interface{}) bool {
	_, ok := e.K(k...).v.(string)
	return ok
}
