package testableexample_test

import (
	"fmt"

	"github.com/sky0621/study-go/temp/testableexample"
)

func ExampleHello() {
	fmt.Println(testableexample.Hello())
	// Output:
	// Hello, someone.
}

func Example() {
	fmt.Println(testableexample.Hello())
	f := testableexample.Foo{Name: "Taro"}
	fmt.Println(f.Hello())
	// Output:
	// Hello, someone.
	// Hello, Taro
}
