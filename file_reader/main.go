package main

import (
	"os"
	"fmt"
)

func reader(f *os.File){

}

func main(){
	f, err := os.Open("test.txt")
	if err != nil{
		fmt.Println(err)
	}
	reader(f)
}