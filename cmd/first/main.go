package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var count, decToBinNum int
	fmt.Print("1. Напишите генератор квадратов натуральных чисел. \n")
	fmt.Print("Введите количество чисел для генерации: ")
	fmt.Scanf("%d\n", &count)
	numbers, squares := generateSequences(count)
	fmt.Println(numbers)
	fmt.Println(squares)

	fmt.Print("2. Найдите самую длинную последовательность нулей в двоичном представлении целого числа. \n")
	fmt.Print("Введите число: ")
	fmt.Scanf("%d\n", &decToBinNum)
	fmt.Println("Двоичное представление: " + decToBin(decToBinNum))
	max:= solution(decToBinNum)
	fmt.Print("Самая длинная последовательность нулей: ", max)
}

func generateSequences(count int) (numbers, squares []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		numbers = append(numbers, rand.Intn(50))
		squares = append(squares, square(numbers[i]))
	}

	return numbers, squares
}

func square(x int) int {
	return x * x
}

func solution(n int) int {
	bin := decToBin(n)
	zerosArray := strings.Split(bin, "1")
	max := 0
	for _, item := range zerosArray {
		if cuurentLen := len(item); cuurentLen > max {
			max = cuurentLen
		}
	}
	return max
}

func decToBin(n int) string {
	return strconv.FormatInt(int64(n), 2)
}