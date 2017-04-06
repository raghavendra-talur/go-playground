package main

import "fmt"

func innerFunc() (e error) {

	s := "hi"
	e = fmt.Errorf("first error")
	defer func() {
		if e != nil {
			fmt.Println("hello in defer %v", s)
		}
	}()
	e = fmt.Errorf("second error")
	fmt.Println("hello")
	s = "bye"
	return e
}

func main() {

	err := innerFunc()
	if err != nil {
		fmt.Println("got err")
	}
}
