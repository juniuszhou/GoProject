package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLineByLine(path string) {
	f, err := os.Open(path)
	defer f.Close()
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func readAll(path string) {
	// read all
	dat, err := ioutil.ReadFile(path)
	check(err)
	fmt.Println(string(dat))
}

func readBuffer(path string) {
	// open file , read 10 bytes each time.
	f, err := os.Open(path)
	defer f.Close()
	check(err)
	b1 := make([]byte, 10)
	for {
		n, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s", n, string(b1))
		if n < 10 {
			fmt.Println("File read done.")
			break
		}
	}
}

func locateFile(path string) {
	// open file , read 10 bytes each time.
	f, err := os.Open(path)
	defer f.Close()
	check(err)
	b1 := make([]byte, 10)
	// whence 0 means begin, 1 current, 2 end.
	// offset to whence.
	f.Seek(3, 0)
	for {
		n, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s", n, string(b1))
		if n < 10 {
			fmt.Println("File read done.")
			break
		}
	}
}

func main() {
	path := "/home/junius/code/src/github.com/juniuszhou/GoProject/basic/a.txt"
	readAll(path)
	readBuffer(path)
	readLineByLine(path)

}
