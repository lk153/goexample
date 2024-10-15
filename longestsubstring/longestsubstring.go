package longestsubstring

// func main() {
// result := longestsubstring.Exec("ABCABCDAB")
// fmt.Println(result)
// }

func Exec(input string) (output string) {
	counter := 0
	mapCheck := map[rune]bool{}
	max := 0
	tmpStr := ""
	for _, c := range input {
		if mapCheck[c] {
			counter = 0
			mapCheck = map[rune]bool{}
			tmpStr = ""
		}

		mapCheck[c] = true
		tmpStr += string(c)
		counter++

		if counter > max {
			max = counter
			output = tmpStr
		}
	}

	return
}
