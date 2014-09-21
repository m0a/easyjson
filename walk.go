package easyjson


// parameter walker is function(key interface{},value JsonAccessor)
func (e easyJsonObj)Walk(walker func(key interface{},value JsonAccessor)){
	for k,v:=range e.RangeObjects(){
		go walker(k,v)
	}
}

// example: for k,v:= range XXX.RangeObjects() {
func (e easyJsonObj)RangeObjects()(ret map[interface {}]JsonAccessor){

	switch e.v.(type){
	//list
	case []interface {}:
		list:=e.v.([]interface {})
		ret = make(map[interface {}]JsonAccessor,len(list))
		for i,v:=range list {
			ejo:=easyJsonObj{v:v}
			ret[i]=ejo
		}
	case map[string]interface {}:
		dict:=e.v.(map[string]interface {})
		ret = make(map[interface {}]JsonAccessor,len(dict))
		for k,v:=range dict {
			ejo:=easyJsonObj{v:v}
			ret[k]=ejo
		}
	default:
		panicf("not Dictionary.not Array")
	}
	return
}

