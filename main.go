package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var romans = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100,
}

var arabians = map[int]string{
	100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
}

func isRoman(s string) bool {
	_, ok := romans[string(s[0])]
	return ok
}

func romanToInt(s string) int {
	end := len(s) - 1
	arr := strings.Split(s, "")
	result := romans[arr[end]]

	for i := end - 1; i >= 0; i-- {
		num := romans[arr[i]]

		if num >= romans[arr[i+1]] {
			result += num
		} else {
			result -= num
		}
	}
	return result
}

func intToRoman(num int) string {
	if num < 1 {
		fmt.Println(errors.New("В римской системе нет отрицательных чисел"))
	}

	var result string
	sortedKeys := make([]int, 0, len(arabians))

	for k := range arabians {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedKeys)))

	for num > 0 {
		for _, key := range sortedKeys {
			if num >= key {
				result += arabians[key]
				num -= key
				break
			}
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение:")
		exp, _ := reader.ReadString('\n')
		exp = strings.TrimSpace(exp)

		arr := strings.Split(exp, " ")
		if len(arr) != 3 {
			fmt.Println(errors.New("введено неккоректное выражение"))
			continue
		}

		operand1 := arr[0]
		operand2 := arr[2]
		operator := arr[1]

		if isRoman(operand1) == isRoman(operand2) {
			var res int
			var a, b int
			isRoman := isRoman(operand1)
			if isRoman {
				a = romanToInt(operand1)
				b = romanToInt(operand2)
			} else {
				a, _ = strconv.Atoi(operand1)

				b, _ = strconv.Atoi(operand2)
			}

			if a > 10 || b > 10 || a < 1 || b < 1 {
				fmt.Println(errors.New("числа должны быть от 1 до 10 включительно"))
				break
			}

			switch operator {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			default:
				fmt.Println(errors.New("неподдерживаемая операция"))
				return
			}

			if isRoman {
				fmt.Println(intToRoman(res))
			} else {
				fmt.Println(res)
			}
		} else {
			fmt.Println(errors.New("числа должны быть в одном формате"))
			break
		}

	}

}
