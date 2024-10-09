package regex

import (
	"fmt"
	"regexp"
	"strconv"
)

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

func transform(total int) (result string) {
	numMap := map[string]string{
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
		"+": "plus",
	}

	str := strconv.Itoa(total)
	for _, c := range str {
		result += numMap[string(c)]
	}

	return
}

func calculate(parsed []string) int {
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

	total := 0
	st := ""
	parsed = append(parsed, "plus")
	for _, word := range parsed {
		if word == "plus" {
			fmt.Println("st", st)
			i, err := strconv.Atoi(st)
			if err != nil {
				return 0
			}
			total += i
			st = ""
		} else {
			st += numMap[word]
			fmt.Println("defaultst", st)
		}
	}

	fmt.Println("total", total)
	return total
}
