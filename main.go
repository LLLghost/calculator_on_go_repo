package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RimToArab(RimNumber string) int {
	var Arab int
	switch {
	case RimNumber == "I":
		Arab = 1
	case RimNumber == "II":
		Arab = 2
	case RimNumber == "III":
		Arab = 3
	case RimNumber == "IV":
		Arab = 4
	case RimNumber == "V":
		Arab = 5
	case RimNumber == "VI":
		Arab = 6
	case RimNumber == "VII":
		Arab = 7
	case RimNumber == "VIII":
		Arab = 8
	case RimNumber == "IX":
		Arab = 9
	case RimNumber == "X":
		Arab = 10
	}
	return Arab
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите два числа от 1 до 10 и операцию (+, -, *, /): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Неверный формат ввода. Введите два числа и операцию в одной строке.")
			continue
		}

		firstNumber, err1 := strconv.Atoi(parts[0])
		secondNumber, err2 := strconv.Atoi(parts[2])
		if err1 != err2 {
			fmt.Println("Числа должны быть одного вида")
			continue
		} else if err1 != nil {
			firstNumber = RimToArab(parts[0])
			secondNumber = RimToArab(parts[2])
		}

		if firstNumber < 1 || firstNumber > 10 || secondNumber < 1 || secondNumber > 10 {
			fmt.Println("Числа должны быть строго от 1 до 10 включительно")
			continue
		}

		operation := parts[1]
		switch operation {
		case "+":
			result := firstNumber + secondNumber
			fmt.Println("Результат:", result)
		case "-":
			result := firstNumber - secondNumber
			fmt.Println("Результат:", result)
		case "*":
			result := firstNumber * secondNumber
			fmt.Println("Результат:", result)
		case "/":
			result := firstNumber / secondNumber
			fmt.Println("Результат:", result)
		default:
			fmt.Println("Неверная операция")
		}
	}
}
