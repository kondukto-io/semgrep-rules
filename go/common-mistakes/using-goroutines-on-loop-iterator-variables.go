package main

import (
	"fmt"
)

func bad(values []string) {
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
}

func bad(values []string) {
	for _, val := range values {
		go func() {
			x := val
			fmt.Println(x)
		}()
	}
}

func bad(values []string) {
	for _, val := range values {
		go func() {
			fmt.Println("hop")
		}()
	}
}

func good(values []string) {
	for _, val := range values {
		go func(val interface{}) {
			fmt.Println(val)
		}(val)
	}
}

//
//
//type val struct {}
//
//func foo() {
//	for _, val := range values {
//		go val.MyMethod()
//	}
//}
//
//func (v *val) MyMethod() {
//	fmt.Println(v)
//}
//
//func bar() {
//	for _, val := range values {
//		newVal := val
//		go newVal.MyMethod()
//	}
//}
