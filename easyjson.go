package easyjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
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

	errorLog := func(format string, a ...interface{}) {
		e.err = fmt.Errorf(format, a...)
	}
	debugLog := func(format string, a ...interface{}) {
		_ = format
		_ = a
	}
	if e.Debug {
		errorLog = func(format string, a ...interface{}) {
			str := fmt.Sprintf(format, a...)
			e.err = errors.New(str)
			fmt.Print(str)
			panic(str)
		}
		debugLog = func(format string, a ...interface{}) {
			fmt.Printf(format, a...)
		}
	}

	debugLog(">>start K(%v)\n", keys)
	defer debugLog(">>end K(%v)\n", keys)

	ret = e
	for _, key := range keys {
		switch key.(type) {
		case int:
			//array access
			array, ok := ret.v.([]interface{})
			if !ok {
				errorLog("Not array. please use string key.")
				return
			}

			v := array[key.(int)]
			if v == nil {
				errorLog("array access value= nil key=%v array=%v\n", key, array)
				return
			}
			ret.v = v
		case string:
			dict, ok := ret.v.(map[string]interface{})
			if !ok {
				errorLog("Not dictionary. please use int key. currentJson:%v",ret)
				return
			}

			v := dict[key.(string)]
			if v == nil {
				errorLog("dict access value= nil key=%v dict=%v\n", key, dict)
				return
			}
			ret.v = v
		default:
			errorLog("key shoud use string or int current type:%T",key)
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

	str = ejo.v.(string)
	return
}


func (e easyJsonObj) AsBool(k ...interface{}) (b bool, err error) {
	ejo := e.K(k...)
	if ejo.err != nil {
		err = ejo.err
		return
	}

	if !ejo.IsBool() {
		err = fmt.Errorf("value is not a bool type = %#v", ejo.v)
		return
	}

	b = ejo.v.(bool)
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
