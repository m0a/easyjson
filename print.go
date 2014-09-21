package easyjson

import (
	"fmt"
	"strings"
)


func (e easyJsonObj) String() string {
	return recursivePrint(e.v)
}


func recursivePrint(i interface {}) string {

	switch i.(type){
	case bool:
		if i.(bool) {
			return "true"
		} else {
			return "false"
		}
	case int:
		return string(i.(int))
	case string:
		return fmt.Sprintf("\"%s\"",i.(string))
	case float64:
		return fmt.Sprintf("%f",i.(float64))
	case []interface {}:
		array:=i.([]interface {})
		str:="["
		list:=make([]string,len(array))
		for i,v:=range array {
			list[i]=recursivePrint(v)
		}
		str+=strings.Join(list,",")
		str+="]"
		return str
	case map[string]interface {}:
		dict:=i.(map[string]interface {})
		str:="{"
		list:=make([]string,len(dict))
		i:=0
		for k,v:=range dict {
			list[i]=fmt.Sprintf("\"%s\":%s",k,recursivePrint(v))
			i++
		}
		str+=strings.Join(list,",")
		str+="}"
		return str
	case nil:
		return "null"
	default:
		return "!panic1!"
	}

	return "!panic2!"
}
