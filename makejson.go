package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	var name string
	var address string

	fmt.Println("Please provide your name")
	fmt.Scan(&name)
	fmt.Println("Please provide your address")
	fmt.Scan(&address)

	strMap := make(map[string]string)
	strMap["name"] = name
	strMap["address"] = address

	json_type_, err := json.Marshal(strMap)
	if err != nil {
		fmt.Print("Error !")
	} else {
		fmt.Println(string(json_type_))
	}

}