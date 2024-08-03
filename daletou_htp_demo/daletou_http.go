package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type ReturnData struct {
	FrontNumberCount NumberList `json:"front_number_count"`
	BackNumberCount  NumberList `json:"back_number_count"`
}

type Number struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}
type NumberList []Number

func (list *NumberList) Len() int { return len(*list) }

func (list *NumberList) Less(i, j int) bool { return (*list)[i].Value < (*list)[j].Value }

func (list NumberList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func generateDistinctRandomNumbers(min, max, count int, frontNumbers map[int]int, keyMap map[string]int) {
	i := 0
	existingNumbers := make(map[int]bool)
	var numberArray []int
	for i < count {
		num := rand.Intn(max-min+1) + min
		if !existingNumbers[num] {
			numberArray = append(numberArray, num)
			existingNumbers[num] = true
			frontNumbers[num] = frontNumbers[num] + 1
			i++
		}
	}
	keyStr := intArrayToString(numberArray)
	vv, exist := keyMap[keyStr]
	if exist {
		keyMap[keyStr] = vv + 1
	} else {
		keyMap[keyStr] = 1
	}
}

func main() {

	http.HandleFunc("/getNumbers", getNumbers)
	err := http.ListenAndServe("117.50.205.248:8100", nil)
	if err != nil {

	}
	rand.Seed(time.Now().UnixNano())

	//fmt.Println("大乐透彩票出奖程序")
	//fmt.Println("====================")
	//map1, map2 := awardNumber(2000000)
	//list1 := mapToList(map1)
	//sort.Sort(list1)
	//list2 := mapToList(map2)
	//sort.Sort(list2)
	//fmt.Println("==========================前区数据(号码,次数)==========================")
	//fmt.Println(list1)
	//fmt.Println("==========================后区数据(号码,次数)==========================")
	//fmt.Println(list2)
}

func getNumbers(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	// 设置响应内容类型为JSON
	rw.Header().Set("Content-Type", "application/json")
	num, err1 := strconv.Atoi(req.FormValue(`num`))
	if err1 != nil {
		//fmt.Println("args error", err1, err2)
		return
	}
	map1, map2, map3, map4 := awardNumber(num)
	list1 := mapToList(map1)
	sort.Sort(list1)
	list2 := mapToList(map2)
	sort.Sort(list2)
	list3 := mapToList1(map3)
	sort.Sort(list3)
	list4 := mapToList1(map4)
	sort.Sort(list4)
	fmt.Println(*list1)
	json.NewEncoder(rw).Encode(ReturnData{FrontNumberCount: *list1, BackNumberCount: *list2})
	//json.NewEncoder(rw).Encode(*list2)
}

func awardNumber(times int) (frontNumbers map[int]int, backNumbers map[int]int, front1 map[string]int, back1 map[string]int) {
	frontNumbers = make(map[int]int)
	backNumbers = make(map[int]int)

	front1 = make(map[string]int)
	back1 = make(map[string]int)
	for i := 0; i < times; i++ {

		// 前区选择5个号码（1-35）
		generateDistinctRandomNumbers(1, 35, 5, frontNumbers, front1)
		//sort.Ints(frontNumbers) // 对前区号码进行升序排序
		// 后区选择2个号码（1-12）
		generateDistinctRandomNumbers(1, 12, 2, backNumbers, back1)
		//sort.Ints(backNumbers) // 对后区号码进行升序排序

		//fmt.Println("前区号码:", frontNumbers)
		//fmt.Println("后区号码:", backNumbers)

	}

	return frontNumbers, backNumbers, front1, back1
}

func mapToList(dataMap map[int]int) *NumberList {
	count := 0
	MyNumberList := make(NumberList, 1)
	for k, v := range dataMap {
		count++
		MyNumberList = append(MyNumberList, Number{Key: strconv.Itoa(k), Value: v})
	}
	start := len(MyNumberList) - count
	MyNumberList = MyNumberList[start:]
	return &MyNumberList
}

func mapToList1(dataMap map[string]int) *NumberList {
	count := 0
	MyNumberList := make(NumberList, 1)
	for k, v := range dataMap {
		count++
		MyNumberList = append(MyNumberList, Number{Key: k, Value: v})
	}
	start := len(MyNumberList) - count
	MyNumberList = MyNumberList[start:]
	return &MyNumberList
}

func intArrayToString(intArray []int) string {
	str := ""
	sort.Ints(intArray)
	for _, v := range intArray {
		str = str + " " + strconv.Itoa(v)
	}
	return str
}
