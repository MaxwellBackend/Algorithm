package astar

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const K_COST_1 int = 10 //直移一格消耗
const K_COST_2 int = 14 //斜移一格消耗

// 地图坐标类型
const (
	PROINT_VIEW_TYPE_S string = "S" // 起点
	PROINT_VIEW_TYPE_E string = "E" // 终点
	PROINT_VIEW_TYPE_1 string = "." // 正常可以行走路径
	PROINT_VIEW_TYPE_2 string = "X" // 阻挡物，无法行走
	PROINT_VIEW_TYPE_3 string = "*" // 计算所得的路径

)

// 坐标
type Point struct {
	x    int
	y    int
	view string // 类型 （. 为路径， X 为障碍物）
}

// 地图
type Map struct {
	points [][]Point         // 地图坐标点
	blocks map[string]*Point // 障碍点
	maxX   int               // 最大X的坐标点
	maxY   int               // 最大Y的坐标点
}

// 创建地图
func NewMap(charMap []string) (m Map) {
	m.points = make([][]Point, len(charMap))
	m.blocks = make(map[string]*Point, len(charMap)*2)
	for x, row := range charMap {
		cols := strings.Split(row, " ")
		m.points[x] = make([]Point, len(cols))
		for y, view := range cols {
			m.points[x][y] = Point{x, y, view}
			if view == PROINT_VIEW_TYPE_2 {
				m.blocks[pointAsKey(x, y)] = &m.points[x][y]
			}
		} // end of cols
	} // end of row
	m.maxX = len(m.points)
	m.maxY = len(m.points[0])
	return m
}

// 获取邻近的坐标点
func (this *Map) getAdjacentPoint(curPoint *Point) (adjacents []*Point) {
	if x, y := curPoint.x, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		adjacents = append(adjacents, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		adjacents = append(adjacents, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		adjacents = append(adjacents, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		adjacents = append(adjacents, &this.points[x][y])
	}
	if x, y := curPoint.x, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		adjacents = append(adjacents, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		adjacents = append(adjacents, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		adjacents = append(adjacents, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		adjacents = append(adjacents, &this.points[x][y])
	}
	return adjacents
}

// 打印地图信息
func (this *Map) PrintMap(path *SearchRoad) {
	for x := 0; x < this.maxX; x++ {
		for y := 0; y < this.maxY; y++ {
			if path != nil {
				if x == path.start.x && y == path.start.y {
					fmt.Print(PROINT_VIEW_TYPE_S)
					goto NEXT
				}
				if x == path.end.x && y == path.end.y {
					fmt.Print(PROINT_VIEW_TYPE_E)
					goto NEXT
				}
				for i := 0; i < len(path.TheRoad); i++ {
					if path.TheRoad[i].x == x && path.TheRoad[i].y == y {
						fmt.Print(PROINT_VIEW_TYPE_3)
						goto NEXT
					}
				}
			}
			fmt.Print(this.points[x][y].view)
		NEXT:
		}
		fmt.Println()
	}
}
func pointAsKey(x, y int) (key string) {
	key = strconv.Itoa(x) + "," + strconv.Itoa(y)
	return key
}

// 单个路径点信息
type PassPoint struct {
	Point
	father *PassPoint // 上一个路径点的信息
	gVal   int
	hVal   int
	fVal   int
}

func NewPassPoint(p *Point, father *PassPoint, end *PassPoint) (ap *PassPoint) {
	ap = &PassPoint{*p, father, 0, 0, 0}
	if end != nil {
		ap.calcFVal(end)
	}
	return ap
}

// 记录A*算法的G值
func (this *PassPoint) calcGVal() int {
	if this.father != nil {
		deltaX := math.Abs(float64(this.father.x - this.x))
		deltaY := math.Abs(float64(this.father.y - this.y))
		if deltaX == 1 && deltaY == 0 {
			this.gVal = this.father.gVal + K_COST_1
		} else if deltaX == 0 && deltaY == 1 {
			this.gVal = this.father.gVal + K_COST_1
		} else if deltaX == 1 && deltaY == 1 {
			this.gVal = this.father.gVal + K_COST_2
		} else {
			panic("father point is invalid!")
		}
	}
	return this.gVal
}

// 记录A*算法的H值 (曼哈顿距离)
func (this *PassPoint) calcHVal(end *PassPoint) int {
	this.hVal = int(math.Abs(float64(end.x-this.x)) + math.Abs(float64(end.y-this.y)))
	return this.hVal
}

// 记录A*算法的F值
func (this *PassPoint) calcFVal(end *PassPoint) int {
	this.fVal = this.calcGVal() + this.calcHVal(end)
	return this.fVal
}

//========================================================================================
type OpenList []*PassPoint

func (self OpenList) Len() int {
	return len(self)
}
func (self OpenList) Less(i, j int) bool {
	return self[i].fVal < self[j].fVal
}
func (self OpenList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}
func (this *OpenList) Push(x interface{}) {
	*this = append(*this, x.(*PassPoint))
}
func (this *OpenList) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}

//========================================================================================
type SearchRoad struct {
	theMap  *Map                  // 地图
	start   PassPoint             // 起点
	end     PassPoint             // 终点
	closeLi map[string]*PassPoint // 关闭列表
	openLi  OpenList              // 开放列表
	openSet map[string]*PassPoint // 计算路径记录
	TheRoad []*PassPoint          // 寻找到的路径
}

// 初始化搜查路径
func NewSearchRoad(startx, starty, endx, endy int, m *Map) *SearchRoad {
	sr := &SearchRoad{}
	sr.theMap = m
	sr.start = *NewPassPoint(&Point{startx, starty, PROINT_VIEW_TYPE_S}, nil, nil)
	sr.end = *NewPassPoint(&Point{endx, endy, PROINT_VIEW_TYPE_E}, nil, nil)
	sr.TheRoad = make([]*PassPoint, 0)
	sr.openSet = make(map[string]*PassPoint, m.maxX+m.maxY)
	sr.closeLi = make(map[string]*PassPoint, m.maxX+m.maxY)
	heap.Init(&sr.openLi)
	heap.Push(&sr.openLi, &sr.start) // 首先把起点加入开放列表
	sr.openSet[pointAsKey(sr.start.x, sr.start.y)] = &sr.start
	// 将障碍点放入关闭列表
	for k, v := range m.blocks {
		sr.closeLi[k] = NewPassPoint(v, nil, nil)
	}
	return sr
}

// 查找路径
func (this *SearchRoad) FindRoad() bool {
	for len(this.openLi) > 0 {
		// 将节点从开放列表移到关闭列表当中。
		x := heap.Pop(&this.openLi)
		curPoint := x.(*PassPoint)
		delete(this.openSet, pointAsKey(curPoint.x, curPoint.y))
		this.closeLi[pointAsKey(curPoint.x, curPoint.y)] = curPoint
		adjacs := this.theMap.getAdjacentPoint(&curPoint.Point)
		for _, p := range adjacs {
			theAP := NewPassPoint(p, curPoint, &this.end)
			if pointAsKey(theAP.x, theAP.y) == pointAsKey(this.end.x, this.end.y) {
				// 找出路径了, 标记路径
				for theAP.father != nil {
					this.TheRoad = append(this.TheRoad, theAP)
					theAP.view = PROINT_VIEW_TYPE_3
					theAP = theAP.father
				}
				return true
			}
			_, ok := this.closeLi[pointAsKey(p.x, p.y)]
			if ok {
				continue
			}
			existAP, ok := this.openSet[pointAsKey(p.x, p.y)]
			if !ok {
				heap.Push(&this.openLi, theAP)
				this.openSet[pointAsKey(theAP.x, theAP.y)] = theAP
			} else {
				oldGVal, oldFather := existAP.gVal, existAP.father
				existAP.father = curPoint
				existAP.calcGVal()
				// 如果新的节点的G值还不如老的节点就恢复老的节点
				if existAP.gVal > oldGVal {
					// restore father
					existAP.father = oldFather
					existAP.gVal = oldGVal
				}
			}
		}
	}
	return false
}
