package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

/*
结构体定义
F = G + H   G = 父节点到起点的距离g+当前节点到父节点的距离  H = 曼哈顿距离  a*a + b*b 开根号   约等于 |(x1-x2)|+ |(y1-y2)|
*/
type gNode struct {
	pos         position
	parentGNode *gNode
	f, g, h     float64
}
type OpenList []gNode

func (openList *OpenList) Len() int { return len(*openList) }

func (openList *OpenList) Less(i, j int) bool { return (*openList)[i].f < (*openList)[j].f }

func (openList OpenList) Swap(i, j int) {
	openList[i], openList[j] = openList[j], openList[i]
}

func (openList *OpenList) Push(x interface{}) {
	item := x.(gNode)
	*openList = append(*openList, item)
}

func (openList *OpenList) Pop() interface{} {
	old := *openList
	allLen := len(*openList)
	item := old[allLen-1]
	*openList = old[0 : allLen-1]
	return item

}

type posNode struct {
	Pos    position `json:"pos"`
	IsPass bool     `json:"is_pass"` //是否可以通行
}

type position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type myJson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var AllNode [][]posNode

func main() {
	fmt.Println("hell A Star Web")
	http.HandleFunc("/getMap", getMap)
	http.HandleFunc("/getRoad", getRoad)
	err := http.ListenAndServe("192.168.4.65:8099", nil)
	if err != nil {
		fmt.Println("http_err", err)
	}
}

/*
	http请求处理
*/
//获取地图数据
func getMap(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("hell")
	mapSize, err1 := strconv.Atoi(req.FormValue(`mapSize`))
	limitNum, err2 := strconv.Atoi(req.FormValue(`limitNum`))
	if err1 != nil || err2 != nil {
		fmt.Println("args error", err1, err2)
		return
	}
	fmt.Println("mapSize ====", mapSize, limitNum)
	rand.Seed(time.Now().UnixNano())
	AllNode = makeMap(mapSize, limitNum)
	x1 := rand.Intn(mapSize)
	y1 := rand.Intn(mapSize)
	x2 := rand.Intn(mapSize)
	y2 := rand.Intn(mapSize)
	fmt.Println(x1, y1, x2, y2)
	fmt.Println("AllNode ==", AllNode)
	// 设置响应内容类型为JSON
	rw.Header().Set("Content-Type", "application/json")
	// 使用json.NewEncoder将结构体编码为JSON并写入响应体

	//AStar(&AllNode[x1][y1].Pos, &AllNode[x2][y2].Pos)
	json.NewEncoder(rw).Encode(AllNode)
}

// 获取地图数据
func getRoad(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")
	starXPos, err1 := strconv.Atoi(req.FormValue(`starX`))
	starYPos, err2 := strconv.Atoi(req.FormValue(`starY`))
	endXPos, err3 := strconv.Atoi(req.FormValue(`endX`))
	endXYos, err4 := strconv.Atoi(req.FormValue(`endY`))
	fmt.Println("GetRole Args ===", starXPos, starYPos, endXPos, endXYos)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("args error", err1, err2)
		return
	}
	Line, lErr := AStar(&AllNode[starXPos][starYPos].Pos, &AllNode[endXPos][endXYos].Pos)
	if lErr == nil {
		json.NewEncoder(rw).Encode(Line)
	}
	return
}

func makeMap(size, limit int) [][]posNode {
	allNode := make([][]posNode, size)
	for i := range allNode {
		allNode[i] = make([]posNode, size)
		for j := range allNode[i] {
			isPass := true
			if rand.Float64() < 0.4 && limit > 0 {
				limit--
				isPass = false
			}
			allNode[i][j] = posNode{position{float64(i), float64(j)}, isPass}
		}
	}
	return allNode
}

/*
	posNode方法
*/
//计算两个posNode的曼哈顿距离 H
func getH(thisNode *position, endNode *position) float64 {
	return math.Abs(thisNode.X-endNode.X) + math.Abs(thisNode.Y-endNode.Y)
}
func getG(thisNode *position, endNode *position) float64 {
	if math.Abs(thisNode.X-endNode.X)+math.Abs(thisNode.Y-endNode.Y) == 1 {
		return 1.0
	}
	return 1.4
}

// 获取节点周围八个节点
func (thisNode *position) get8Around() (posArray []position) {
	oX, oY := thisNode.X, thisNode.Y
	posArray1 := [8]position{{oX + 1, oY}, {oX + 1, oY + 1}, {oX, oY + 1}, {oX - 1, oY + 1}, {oX - 1, oY}, {oX - 1, oY - 1}, {oX, oY - 1}, {oX + 1, oY - 1}}
	//节点是否存在
	for _, node := range posArray1 {
		if node.X >= 0 && node.Y >= 0 {
			posArray = append(posArray, node)
		}
	}
	return posArray
}

/*aStar算法*/

func AStar(startPos *position, endPos *position) ([]position, error) {
	//开启列表
	openList := make(OpenList, 0)
	//检查过的节点map
	nodeMap := make(map[position]bool)
	//关闭列表
	closeList := make([]gNode, 0)
	//初始化起点的数据
	g := 0.0
	h := getH(startPos, endPos)
	starGNode := gNode{*startPos, nil, g + h, g, h}
	//将起点的gNode放进开启列表中
	heap.Push(&openList, starGNode)
	//开启列表为空 需要结束
	for openList.Len() > 0 {
		current := heap.Pop(&openList).(gNode)
		closeList = append(closeList, current)
		//取出第一个 (f最小值) 判断是否为终点
		fmt.Println("open is =", openList)
		fmt.Println("close is =", closeList)
		if current.pos == *endPos {
			fmt.Println("找到路径拉")
			fmt.Println("END open is =", openList)
			fmt.Println("End close is =", closeList)
			//构建路线
			Line := makeLine(&closeList)
			fmt.Println("最终路线为:===", Line)
			return Line, nil
		}
		//	遍历周围八点 继续找F最小的节点
		around := current.pos.get8Around()
		for _, tempPos := range around {
			if !nodeMap[tempPos] && checkNode(&tempPos) {
				nodeMap[tempPos] = true
				newGNode := new(gNode)
				newGNode.g = getG(&tempPos, &current.pos) + current.g
				newGNode.h = getH(&tempPos, endPos)
				newGNode.f = newGNode.g + newGNode.h
				newGNode.pos = tempPos
				newGNode.parentGNode = &current
				heap.Push(&openList, *newGNode)
			}
		}

	}
	return nil, fmt.Errorf("error")
}

func checkNode(pos *position) bool {
	p1 := int32(pos.X)
	p2 := int32(pos.Y)
	len1 := int32(len(AllNode))
	if p1 >= len1 || p2 >= len1 {
		return false
	}
	return AllNode[int32(pos.X)][int32(pos.Y)].IsPass
}

// 绘制路线
func makeLine(closeSet *[]gNode) []position {
	GNode := (*closeSet)[len(*closeSet)-1]
	Line := make([]position, 0)
	Line = append(Line, GNode.pos)
	for GNode.parentGNode != nil {
		GNode = *GNode.parentGNode
		Line = append(Line, GNode.pos)
	}
	return Line
}
