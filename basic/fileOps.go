package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/home/junius/code/src/github.com/juniuszhou/GoProject/basic/a.txt")
	check(err)
	fmt.Println(string(dat))

}
