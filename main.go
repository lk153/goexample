package main

import (
	"fmt"
	"goexample/regex"
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

	// relationships := []string{"a1:a2", "a2:a3", "a4:a3", "a4:a5", "a1:a3", "a1:a1"}
	// n := relationship.GetMinStep(relationships, "a1", "a5")
	// fmt.Println("RESULT:", n)

	result := regex.StrExpress("minusonezeropluseight")
	fmt.Println("RESULT:", result)
}

func IsEqual(actual any, expected any) bool {
	return reflect.DeepEqual(actual, expected)
}
