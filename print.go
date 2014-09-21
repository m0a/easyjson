package easyjson

import (
	"fmt"
	"strings"
)

func (e easyJsonObj) String() string {
	return recursivePrint(e.v,false,0)
}

func (e easyJsonObj) PretyString() string {
	return recursivePrint(e.v,true,1)
}
func (e easyJsonObj) PretyPrint(){
	fmt.Println(e.PretyString())
}

func recursivePrint(i interface{},pretyPrint bool,step int) string {

	createPadding:=func(s int) string {
		ret:=""
		for i:=0; i< s; i++ {
			ret+="\t"
		}
		return ret
	}
	padding:=createPadding(step)
	old_padding:=createPadding(step-1)


	switch i.(type) {
	case bool:
		if i.(bool) {
			return "true"
		} else {
			return "false"
		}
	case int:
		return string(i.(int))
	case string:
		return fmt.Sprintf("\"%s\"", i.(string))
	case float64:
		return fmt.Sprintf("%f", i.(float64))

	//list case
	case []interface{}:
		array := i.([]interface{})
		str := "["
		if pretyPrint {
			str += "\n"
			str += padding
		}

		list := make([]string, len(array))
		for i, v := range array {
			list[i]=""
			list[i] += recursivePrint(v,pretyPrint,step+1)
		}
		if pretyPrint {
			str += strings.Join(list, ",\n"+padding)
		} else {
			str += strings.Join(list, ",")
		}

		if pretyPrint {
			str+="\n"
			str += old_padding
		}

		str += "]"
		return str
	//dicitionary case
	case map[string]interface{}:
		dict := i.(map[string]interface{})
		str := "{"
		if pretyPrint {
			str += "\n"
		}

		list := make([]string, len(dict))
		i := 0
		for k, v := range dict {
			list[i]=""
			if pretyPrint {
				list[i]+=padding
			}
			list[i] += fmt.Sprintf("\"%s\":%s", k, recursivePrint(v,pretyPrint,step+1))
			i++
		}
		if pretyPrint {
			str += strings.Join(list, ",\n")
		} else {
			str += strings.Join(list, ",")
		}

		if pretyPrint {
			str+="\n"
			str+=old_padding
		}
		str += "}"
		return str
	case nil:
		return "null"
	default:
		return "!panic1!"
	}

	return "!panic2!"
}
