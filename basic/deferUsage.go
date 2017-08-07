package main

import "os"

func main(){
	srcFile, err := os.Open("a.txt")
	if err != nil {
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create("a.txt")
	if err != nil {
		return
	}

	defer dstFile.Close()
	
}