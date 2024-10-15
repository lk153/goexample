package main

import (
	"fmt"
	"goexample/longestsubstring"
	"reflect"
)

func main() {
	// type test struct {
	// 	input  []int
	// 	output []interface{}
	// }
	// testcases := map[string]test{
	// 	"case-1": {
	// 		input:  []int{5, 5, 2, 2, 3, 4, 2, 2, 2, 2, 3, 2},
	// 		output: []interface{}{2, 2, 2, 2},
	// 	},
	// }

	// for caseName, cases := range testcases {
	// 	actual := bitwiseor.FindLongestSlice(cases.input)
	// 	// actual := ""
	// 	fmt.Println("caseName:", caseName)
	// 	fmt.Println("result:", actual)
	// 	if IsEqual(actual, cases.output) {
	// 		fmt.Println("PASSED")
	// 	} else {
	// 		fmt.Println("FAILED")
	// 	}
	// 	fmt.Println()
	// }

	result := longestsubstring.Exec("ABCABCDAB")
	fmt.Println(result)
}

func IsEqual(actual any, expected any) bool {
	return reflect.DeepEqual(actual, expected)
}
