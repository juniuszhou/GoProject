package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main(){
	// encode
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	result, err := json.Marshal(group)
	if err != nil{
		fmt.Println(result)
		fmt.Println("++++++++++++++++++++++++++++++++++")
	} else {
		//print with bytes array
		fmt.Println(result)
		fmt.Println(string(result))
	}

	fmt.Println("++++++++++++++++++++++++++++++++++")
	//print json string
	os.Stdout.Write(result)
	os.Stdout.Write([]byte("\n"))

	//decode
	err = json.Unmarshal(result, &group)
	fmt.Println(group.ID, group.Name)

}