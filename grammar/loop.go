package grammar

import "fmt"

func Loop() {
	var values [] int
	for value := range values {
		fmt.Println(value)
	}
}
