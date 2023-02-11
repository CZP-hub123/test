package main

import "fmt"

func citymapprint(mymap map[string]string) {
	for key, value := range mymap {
		fmt.Println(key, "的首都是", value)

	}
}

func main() {
	cityMap := make(map[string]string)
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	citymapprint(cityMap)

	cityMap["USA"] = "dc"

	delete(cityMap, "China")
	citymapprint(cityMap)
}
