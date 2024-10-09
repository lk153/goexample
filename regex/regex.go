package regex

import (
	"fmt"
	"regexp"
	"strconv"
)

// func main() {
// 	result := regex.StrExpress("minusonezeropluseight")
// 	fmt.Println("RESULT:", result)
// }

func StrExpress(input string) (output string) {
	re := regexp.MustCompile("zero|one|two|three|four|five|six|seven|eight|nine|minus|plus")
	result := re.FindAll([]byte(input), -1)
	parsed := []string{}
	for _, str := range result {
		parsed = append(parsed, string(str))
	}

	total := calculate(parsed)
	return transform(total)
}

func transform(total int) string {
	fmt.Println("total", total)
	strMap := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
		"5": "five",
		"6": "six",
		"7": "seven",
		"8": "eight",
		"9": "nine",
		"0": "zero",
		"-": "negative",
	}

	str := strconv.Itoa(total)
	result := ""
	for _, c := range str {
		result += strMap[string(c)]
	}

	return result
}

func calculate(parsed []string) int {
	fmt.Println("parsed", parsed)
	numMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
		"minus": "-",
		"plus":  "+",
	}

	st := ""
	com := []string{}
	parsed = append(parsed, "plus")
	for _, word := range parsed {
		switch {
		case word == "plus":
			if st == "" {
				st += numMap[word]
				continue
			}

			com = append(com, st)
			st = ""
		case word == "minus":
			if st == "" {
				st += numMap[word]
				continue
			}

			com = append(com, st)
			st = "-"
		default:
			st += numMap[word]
		}
	}

	total := 0
	for _, st := range com {
		i, err := strconv.Atoi(st)
		if err != nil {
			return 0
		}
		total += i
	}

	return total
}
