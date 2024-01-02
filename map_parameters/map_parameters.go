package map_parameters

import "fmt"

func MyFunctionMapParams(params map[string]interface{}) {
	fmt.Println(params["Param1"], params["Param2"], params["Param3"])
}
