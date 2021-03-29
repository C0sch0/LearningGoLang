package main

import (
	"fmt"
)

func main()  {
	var idMap map[string]int
	idMap = make(map[string]int)
	idMap["joe"] = 2
	// id: Value, p: bool of existence
	id, p := idMap["joe"]
	fmt.Println(id)
	fmt.Println(p)

	idMap["jane"] = 23
	fmt.Println(idMap)

	delete(idMap, "joe")
	fmt.Println(idMap)
	idMap["joe"] = 2

	for key, val := range idMap {
		fmt.Println(key, val)
	}
}
