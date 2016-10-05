package main

import (
	"fmt"
	"net/http"

	"github.com/m0a/easyjson"
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

	json, err := easyjson.NewEasyJson(resp.Body)
	if err != nil {
		panic("json convert err")
	}

	//easy access!
	json.K("routes", 0, "bounds", "southwest").PrettyPrint()

	//support method chain
	json.K("routes").K(0).K("bounds").K("southwest").PrettyPrint()

	// if use loop
	for k, v := range json.K("routes").K(0).RangeObjects() {
		fmt.Printf("%v:%v\n", k, v)
	}

	//string value
	copyrights, err := json.K("routes").K(0).K("copyrights").AsString()
	if err != nil {
		panic("AsString err")
	}
	fmt.Printf("copyrights=%s", copyrights)

	//more easy use
	fmt.Printf("copyrights=%s", json.AsStringPanic("routes", 0, "copyrights"))

}
