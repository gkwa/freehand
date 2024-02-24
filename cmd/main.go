package main

import (
	"fmt"

	"github.com/taylormonacelli/freehand/functional_options"
	"github.com/taylormonacelli/freehand/map_parameters"
	"github.com/taylormonacelli/freehand/named_return_values"
	"github.com/taylormonacelli/freehand/struct_example"
	"github.com/taylormonacelli/freehand/variadic_functions"
)

func main() {
	struct_example.MyFunction()

	functional_options.MyFunctionWithOptions(
		functional_options.WithParam1(42),
		functional_options.WithParam2("hello"),
		functional_options.WithParam3(true),
	)

	map_parameters.MyFunctionMapParams(
		map[string]interface{}{
			"Param1": 42,
			"Param2": "hello",
			"Param3": true,
		},
	)

	variadic_functions.MyFunctionVariadicParams(
		"Param1", 42,
		"Param2", "hello",
		"Param3", true,
	)

	result1, result2, result3 := named_return_values.MyFunctionNamedReturnValues()
	fmt.Println(result1, result2, result3)
}
