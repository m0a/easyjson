package easyjson

import (
	"fmt"
	"strings"
)

func (e easyJsonObj) String() string {
	return recursivePrint(e.v,false,0)
}

func (e easyJsonObj) PrettyString() string {
	return recursivePrint(e.v,true,1)
}
func (e easyJsonObj) PrettyPrint(){
	fmt.Println(e.PrettyString())
}

func recursivePrint(i interface{},prettyPrint bool,step int) string {

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
		if prettyPrint {
			str += "\n"
			str += padding
		}

		list := make([]string, len(array))
		for i, v := range array {
			list[i]=""
			list[i] += recursivePrint(v,prettyPrint,step+1)
		}
		if prettyPrint {
			str += strings.Join(list, ",\n"+padding)
		} else {
			str += strings.Join(list, ",")
		}

		if prettyPrint {
			str+="\n"
			str += old_padding
		}

		str += "]"
		return str
	//dicitionary case
	case map[string]interface{}:
		dict := i.(map[string]interface{})
		str := "{"
		if prettyPrint {
			str += "\n"
		}

		list := make([]string, len(dict))
		i := 0
		for k, v := range dict {
			list[i]=""
			if prettyPrint {
				list[i]+=padding
			}
			list[i] += fmt.Sprintf("\"%s\":%s", k, recursivePrint(v,prettyPrint,step+1))
			i++
		}
		if prettyPrint {
			str += strings.Join(list, ",\n")
		} else {
			str += strings.Join(list, ",")
		}

		if prettyPrint {
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
