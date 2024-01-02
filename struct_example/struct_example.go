package struct_example

import "fmt"

type MyParams struct {
	Param1 int
	Param2 string
	Param3 bool
}

func MyFunction() {
	params := MyParams{Param1: 42, Param2: "hello", Param3: true}
	fmt.Println(params.Param1, params.Param2, params.Param3)
}
