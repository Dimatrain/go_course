package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count, shiftPos int
	fmt.Print("1. Вращение массива. \n")
	fmt.Print("Введите количество чисел для генерации массива: ")
	fmt.Scanf("%d\n", &count)
	arrayInt := generateArray(count)
	fmt.Print("Исходный массив: ")
	fmt.Println(arrayInt)
	fmt.Print("Введите количество сдвигов: ")
	fmt.Scanf("%d\n", &shiftPos)
	fmt.Print("Сдвинутый массив: ")
	fmt.Println(shftArray(arrayInt, shiftPos))

	fmt.Print("2. Дан массив целых чисел, необходимо посчитать количество уникальных значений. \n")
	fmt.Print("Введите количество чисел для генерации массива: ")
	fmt.Scanf("%d\n", &count)
	arrayInt = generateArray(count)
	fmt.Print("Исходный массив: ")
	fmt.Println(arrayInt)
	fmt.Print("Количество уникальных значений: ")
	fmt.Println(countUniqValues(arrayInt))
}

func generateArray(n int) []int {
	arrayInt := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := range arrayInt {
		arrayInt[i] = rand.Intn(100)
	}

	return arrayInt
}

func countUniqValues(arr []int) int {
	uniq := make(map[int]bool)
	for _, val := range arr {
		uniq[val] = true
	}
	return len(uniq)
}

func shftArray(arrayInt []int, shiftPos int) []int  {
	maxSize := len(arrayInt)

	if shiftPos == maxSize {
		return arrayInt
	}

	if shiftPos > maxSize {
		shiftPos = shiftPos % maxSize
	}
	rightSlice := arrayInt[0:maxSize - shiftPos]
	leftSlice := arrayInt[maxSize - shiftPos:maxSize]

	return append(leftSlice, rightSlice...)
}
