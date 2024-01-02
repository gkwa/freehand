package variadic_functions

import "fmt"

func MyFunctionVariadicParams(params ...interface{}) {
	for i := 0; i < len(params); i += 2 {
		key := params[i].(string)
		value := params[i+1]
		fmt.Println(key, value)
	}
}
