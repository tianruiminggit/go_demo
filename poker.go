package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type poker struct {
	pokerType pokerFlower
	pokerNum  int
}
type pokerArr []poker
type pokerHands struct {
	handsType int
	pokerNum1 int
	pokerNum2 int
	pokerNum3 int
	pokerNum4 int
	pokerNum5 int
}

// poker花色
type pokerFlower int

const (
	SPADE         = 1  //黑桃 ♠
	HEART         = 2  //红桃 ♥
	CLUB          = 3  //梅花 ♣
	DIAMOND       = 4  //方片 ♦
	AllPoker      = 5  //所有的
	StraightFlush = 10 //同花顺
	FourOfKind    = 9  //四条
	FullHouse     = 8  //葫芦
	Flush         = 7  //同花
	Straight      = 6  //顺子
	ThreeOfKind   = 5  //三条
	TwoPair       = 4  //两对
	OnePair       = 3  //对子
	HighCard      = 2  //高牌
)

func (poker poker) show() {
	fmt.Println(poker)
}

func (poker poker) getNum() int {
	return poker.pokerNum
}

func (poker *poker) set(num int) {
	poker.pokerNum = num
}

// Len 实现 sort.Interface 的三个方法
func (pokerList *pokerArr) Len() int {
	return len(*pokerList)
}

func (pokerList *pokerArr) Less(i, j int) bool {
	return (*pokerList)[i].pokerNum > (*pokerList)[j].pokerNum
}

func (pokerList *pokerArr) Swap(i, j int) {
	(*pokerList)[i], (*pokerList)[j] = (*pokerList)[j], (*pokerList)[i]
}

func main() {
	poker1 := poker{SPADE, 1}
	poker2 := &poker{HEART, 2}

	poker1.show()
	poker1.set(12)
	poker1.show()

	poker2.show()
	poker2.set(22)
	poker2.show()

	pokerList1 := []poker{{HEART, 1}, {AllPoker, 7}, {HEART, 2}}
	map1 := make(map[int][]poker)
	map1[1] = pokerList1
	map2 := map[int][]poker{
		2: {{HEART, 2}},
	}
	fmt.Println(map2[2])
	pokerList2 := &pokerArr{{HEART, 13}, {AllPoker, 7}, {HEART, 2}, {HEART, 3}, {HEART, 4}, {HEART, 5}}
	fmt.Println(pokerList2)
	sort.Sort(pokerList2)
	fmt.Println(pokerList2)

	bool1, ret := pokerList2.isStraight(false)
	fmt.Println(bool1)
	fmt.Println(ret)
	//judgeNotStraight(*pokerList2)
	pokerCardList := initPokerCard()
	fmt.Println(pokerCardList)
	randomPokers := randomPoker(15, 52, pokerCardList)
	fmt.Println(randomPokers)

	judgePoker(randomPokers)
}

func findMax(intList []int) int {
	if len(intList) > 0 {
		max := intList[0]
		for _, value := range intList {
			if value > max {
				max = value
			}
		}
		return max
	} else {
		return 0
	}
}

// 判断Poker牌型
func judgePoker(arr pokerArr) {
	var myMap = make(map[pokerFlower]pokerArr)
	for _, poker := range arr {
		value, exist := myMap[poker.pokerType]
		if exist {
			myMap[poker.pokerType] = append(value, poker)
		} else {
			newValue := pokerArr{poker}
			myMap[poker.pokerType] = newValue
		}
	}
	flag, handsType := isFlush(myMap)
	if flag {
		fmt.Println(handsType)
		return
	}
	flag, handsType = arr.isStraight(false)
	if flag {
		fmt.Println(handsType)
		return
	}
	judgeNotStraight(arr)
}

// 是否是同花
func isFlush(pokerMap map[pokerFlower]pokerArr) (bool, pokerHands) {
	for key, pokerList := range pokerMap {
		if key != AllPoker {
			if len(pokerList) >= 5 {
				isStraight, handsType := pokerList.isStraight(true)
				if isStraight {
					return true, handsType
				}
				return false, handsType
			}
		}
	}
	return false, pokerHands{}
}

// 是否是顺子
func (pokerList *pokerArr) isStraight(flushFlag bool) (straightFlag bool, pokerHands pokerHands) {
	//排序  降序
	sort.Sort(pokerList)
	newPokerList := (*pokerList)[:]
	//取第一个  是否是Ace  是A的话 判断顺子需要将点数1加入数组
	if (*pokerList)[0].pokerNum == 13 {
		newPokerList = append(newPokerList, poker{(*pokerList)[0].pokerType, 1})
	}
	//去最后一个 是否是1  是否是Ace
	//lastPoker := (*pokerList)[len(*pokerList)-1]
	//下标变量
	index := 0
	//长度
	listLen := len(newPokerList)
	for i := index; index+5 <= listLen; i = index {
		count := 1
		for j := i + 1; j < index+5; j++ {
			//未成顺子
			if (newPokerList[i].pokerNum - newPokerList[j].pokerNum) != 1 {
				index = j
				count = 1
				break
			}
			if (newPokerList[i].pokerNum - newPokerList[j].pokerNum) == 0 {
				i++
			}
			i++
			count++
		}
		if count == 5 {
			straightFlag = true
			if flushFlag {
				pokerHands.handsType = StraightFlush
			} else {
				pokerHands.handsType = Straight
			}
			pokerHands.pokerNum1 = newPokerList[index].pokerNum
			return
		}
		//到此 表示成顺  因为是倒序 所以是最大顺子 只需要判断是否index的poker是否为K 是否有A
		//若有则是A顺 才是最大
		//if (*pokerList)[index].pokerNum == 13 && lastPoker.pokerNum == 1 {
		//	pokerHands.pokerNum1 = 14
		//	return
		//}
	}
	straightFlag = false
	return
}

func judgeNotStraight(pokerArr pokerArr) {
	tempMap := make(map[int]int)
	for _, poker := range pokerArr {
		value, exists := tempMap[poker.pokerNum]
		if exists {
			tempMap[poker.pokerNum] = value + 1
		} else {
			tempMap[poker.pokerNum] = 1
		}
	}
	var tempHandsType int
	var fourNum int
	var threeNum int
	var twoNum int //最大对
	var oneNum int //第二大对
	var highNum int
	for key, value := range tempMap {
		switch value {
		case 1:
			if tempHandsType > HighCard {

			} else {
				tempHandsType = HighCard
				if highNum < key {
					highNum = key
				}
			}
		case 2:
			switch tempHandsType {
			case HighCard:
				tempHandsType = OnePair
				twoNum = key
			case OnePair: //一对变两对  需要赋值大小对的值
				tempHandsType = TwoPair
				if twoNum > key {
					oneNum = key
				} else {
					twoNum, oneNum = key, twoNum
				}
			case TwoPair: //两对遇对子 更新大小对
				if key > twoNum {
					twoNum, oneNum = key, twoNum
				}
				if key > oneNum {
					oneNum = key
				}
			case ThreeOfKind: //三条遇对变葫芦
				tempHandsType = FullHouse
				twoNum = key
			case FullHouse: //葫芦遇对子 更新大小对
				if key > twoNum {
					twoNum, oneNum = key, twoNum
				}
				if key > oneNum {
					oneNum = key
				}
			case FourOfKind:
			}
		case 3:
			tempThree := key
			switch tempHandsType {
			case OnePair: //葫芦
				tempHandsType = FullHouse
			case TwoPair: //葫芦
				tempHandsType = FullHouse
			case ThreeOfKind: //三条遇三条 变葫芦
				tempHandsType = FullHouse
				if key > threeNum {
					threeNum, twoNum = key, threeNum
				} else {
					twoNum = key
				}
			case FourOfKind:
				tempHandsType = FourOfKind
				fourNum = key
			}
			threeNum = tempThree
		case 4:
			tempHandsType = FourOfKind
			fourNum = key
		default:
			fmt.Println("My Key Is", tempHandsType, key, fourNum, threeNum)
		}
	}
	fmt.Println(tempHandsType, fourNum, threeNum, twoNum, oneNum, highNum)
}

func initPokerCard() *pokerArr {
	pokerArr := make(pokerArr, 52)
	count := 0
	for i := 0; i < 4; i++ {
		for j := 2; j < 15; j++ {
			var temp pokerFlower
			switch i {
			case 0:
				temp = SPADE
			case 1:
				temp = HEART
			case 2:
				temp = CLUB
			case 3:
				temp = DIAMOND

			}
			pokerArr[count] = poker{temp, j}
			count++
		}
	}
	return &pokerArr
}

func randomPoker(need, count int, arr *pokerArr) pokerArr {
	unSameMap := make(map[int]bool, need)
	for i := 0; i < need; {
		rand.Seed(time.Now().UnixNano())
		temp := rand.Intn(count)
		if !unSameMap[temp] {
			unSameMap[temp] = true
			i++
		}
	}
	index := 0
	pokerArr := make(pokerArr, need)
	for k, _ := range unSameMap {
		pokerArr[index] = (*arr)[k]
		index++
	}
	return pokerArr
}
