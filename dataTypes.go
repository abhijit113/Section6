package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	boolFn()
	numFn()
	strFn()
	consFn()
}

func boolFn() {
	a := 30
	b := 42
	fmt.Println("inside function foo")
	fmt.Println(a == b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
}

func numFn() {
	a1 := 40
	a2 := 40.00
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func strFn() {
	s := `Hello Playground`
	fmt.Println(s)

	bs:=[]byte(s)  // convert each character into byte
	fmt.Println("printing bs vale")
	fmt.Println(bs)
	fmt.Println("printing s type")
	fmt.Printf("%T\n",s)
	fmt.Println("printing unicode values")
	fmt.Printf("%c\n",bs)

	for i:=0; i< len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	for i, v:=range s{
		fmt.Printf("at index position %d we hahve the value hex %#x\n", i, v)
	}

}

func consFn(){
	const a=12
	const b=20.03
	const c="Abhi"
	fmt.Println(a,b,c)
}
