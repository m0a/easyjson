# m0a/easyjson
========

go's easy json access library


## Usage

simple example:

```go

package main

import (
	"github.com/m0a/easyjson"
	"net/http"
	"fmt"
)


func main() {

	url := "http://maps.googleapis.com/maps/api/directions/json?origin=Boston,MA&destination=Concord,MA&sensor=false"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}

	json,err := easyjson.NewEasyJson(resp.Body)
	if err!=nil {
		panic("json convert err")
	}

	//easy access!
	json.K("routes",0,"bounds","southwest").PrettyPrint()

	//support method chain
	json.K("routes").K(0).K("bounds").K("southwest").PrettyPrint()

    //if use loop
	for k,v:=range json.K("routes").K(0).RangeObjects() {
		fmt.Printf("%v:%v\n",k,v)
	}


}
```


