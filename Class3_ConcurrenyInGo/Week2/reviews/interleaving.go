package main

import (
	"fmt"
	"time"
)


var i int

func method1(){
	i=10
}

func method2(){
	i=20
}





func main() {
	fmt.Println("program begins")
	for j:=0;j<20;j++ {
		go method1()
		go method2()

		time.Sleep(50)

		fmt.Print("i = ")
		fmt.Println(i)
	}





}



