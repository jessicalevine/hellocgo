package main

import "fmt"

//extern void CallGoFunc();
import "C"

//export GoFunc
func GoFunc() {
	fmt.Println("GoFunc called")
}

func main() {
	C.CallGoFunc()
}
