package functional_options

import "fmt"

type MyParams struct {
	Param1 int
	Param2 string
	Param3 bool
}

type MyParamsOption func(*MyParams)

func WithParam1(value int) MyParamsOption {
	return func(params *MyParams) {
		params.Param1 = value
	}
}

func WithParam2(value string) MyParamsOption {
	return func(params *MyParams) {
		params.Param2 = value
	}
}

func WithParam3(value bool) MyParamsOption {
	return func(params *MyParams) {
		params.Param3 = value
	}
}

func MyFunctionWithOptions(options ...MyParamsOption) {
	params := MyParams{}
	for _, option := range options {
		option(&params)
	}
	fmt.Println(params.Param1, params.Param2, params.Param3)
}
