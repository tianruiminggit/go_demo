package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func generateRandomNumbers(min, max, count int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, count)

	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(max-min+1) + min
	}

	return numbers
}

//func main() {
//	fmt.Println("大乐透彩票出奖程序")
//	fmt.Println("====================")
//
//	// 假设大乐透的规则为：前区选择5个号码（1-35），后区选择2个号码（1-12）
//	frontNumbers := generateRandomNumbers(1, 35, 5)
//	backNumbers := generateRandomNumbers(1, 12, 2)
//
//	fmt.Println("前区号码:", frontNumbers)
//	fmt.Println("后区号码:", backNumbers)
//}

func generateDistinctRandomNumbers(min, max, count int, existingNumbers map[int]bool) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 0)

	for len(numbers) < count {
		num := rand.Intn(max-min+1) + min
		if !existingNumbers[num] {
			numbers = append(numbers, num)
			existingNumbers[num] = true
		}
	}

	return numbers
}

func main() {
	fmt.Println("大乐透彩票出奖程序")
	fmt.Println("====================")
	awardNumber(1)
}

func awardNumber(times int) {
	for i := 0; i < times; i++ {
		existingNumbers := make(map[int]bool)
		// 前区选择5个号码（1-35）
		frontNumbers := generateDistinctRandomNumbers(1, 35, 5, existingNumbers)
		sort.Ints(frontNumbers) // 对前区号码进行升序排序

		// 后区选择2个号码（1-12）
		backNumbers := generateDistinctRandomNumbers(1, 12, 2, existingNumbers)
		sort.Ints(backNumbers) // 对后区号码进行升序排序

		fmt.Println("前区号码:", frontNumbers)
		fmt.Println("后区号码:", backNumbers)
	}
}
