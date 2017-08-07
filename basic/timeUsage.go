package main

import (
	"fmt"
	"time"
)

func main(){
	// time to string
	var t = time.Now()
	fmt.Println(t)
	// must set the layout at this timestamp
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	// string to time. format could be flexible, but timestamp must be the same.
	//layout := "2014-09-12T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	tt , _ := time.Parse(layout , str)
	fmt.Println(tt)

	//
	format := "20060102150405"
	timeStr := "20140102030405"
	ttt , _ := time.Parse(format, timeStr)
	fmt.Println(ttt)

}
