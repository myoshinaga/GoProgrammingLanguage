package main

import "fmt"

func main() {
	fmt.Printf("%s\n", panicAndRecover())
}

func panicAndRecover() (ret string) {
	defer func() {
		if p := recover(); p != nil {
			ret = fmt.Sprintf("recover:%v\n", p)
		}
	}()
	panic("panic!")
}
