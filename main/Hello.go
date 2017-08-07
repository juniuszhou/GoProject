package main

import (
	"fmt"
	"github.com/juniuszhou/GoProject/grammar"
)


func main(){
	grammar.PrintFmt("programme done.")
	// error printed with red.
	// grammar.LogData("ERROR", "programme down.")
	grammar.DefineVar()
	fmt.Println("game over")

	// class interface usage
	var call grammar.ICall = new(grammar.Implement)
	call.CallFunc()

	// log usage

	fmt.Println("over")

}

