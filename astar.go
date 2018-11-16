package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct{
	x int
	y int
	val interface{}
}

type theMap struct{
	AllPoints 	[][]*Point
	BlockPoints map[string]*Point
	MaxX   		int
	MaxY 		int
}

type AsPoint struct {
	*Point
	Father *AsPoint
	gVal   int
	hVal   int
	fVal   int
}

type AStarRoad struct{
	*theMap  							//地图信息
	OpenList 	openList			//开放节点
	CloSet  	map[string]*AsPoint 	//关闭节点
	OpenSet     map[string]*AsPoint
	Start       *AsPoint            	//开始点
	End 		*AsPoint		   		//结束点
	Road         []AsPoint		   		//路径
}


type openList []*AsPoint
func (self openList) Len() int           { return len(self) }
func (self openList) Less(i, j int) bool { return self[i].fVal < self[j].fVal }
func (self openList) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (this *openList) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*this = append(*this, x.(*AsPoint))
}

func (this *openList) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}


/*
格式化key
 */
func formatKey(x,y int) string {
	return strconv.Itoa(x)+ strconv.Itoa(y)
}

func newMap(charMap []string) *theMap {
	mp := new(theMap)
	mp.BlockPoints = make(map[string]*Point,0)
	mp.AllPoints = make([][]*Point,len(charMap))
	for x,val := range charMap {
		cols := strings.Split(val, " ")
		mp.AllPoints[x] = make([]*Point,len(cols))
		for y, v := range cols {
			point := new(Point)
			point.x = x
			point.y = y
			point.val = v
			mp.AllPoints[x][y] = point
			if v == "X" {
				mp.BlockPoints[formatKey(x,y)] = point
			}
		}
	}
	mp.MaxX = len(mp.AllPoints)
	mp.MaxY = len(mp.AllPoints[0])
	return mp
}

//初始化寻路信息
func newAStarRoad(startX,startY,endX,endY int,charMap *theMap) *AStarRoad {
	key := formatKey(startX,startY)
	if _,ok := charMap.BlockPoints[key];ok {
		fmt.Println("起始点在锁定区域")
		return nil
	}
	 key = formatKey(endX,endY)
	 if _,ok := charMap.BlockPoints[key];ok {
	 	fmt.Println("结束点在锁定区域")
	 	return nil
	 }
	road := new(AStarRoad)
	road.theMap = charMap
	road.Start = new(AsPoint)
	road.Start.Point = new(Point)
	road.Start.x = startX
	road.Start.y = startY
	road.Start.val = "S"
	road.End = new(AsPoint)
	road.End.Point = new(Point)
	road.End.x = endX
	road.End.y = endY
	road.End.val ="E"
	road.OpenSet = make(map[string]*AsPoint,0)
	road.CloSet = make(map[string]*AsPoint,0)
	road.Road = make([]AsPoint,0)
	//开始节点加入开放列表
	heap.Init(&road.OpenList)
	heap.Push(&road.OpenList,road.Start)
	//将障碍点放入关闭集合
	for key, val := range road.theMap.BlockPoints {
		road.CloSet[key] = newAsPoint(val,nil,nil)
	}
	return road
}

func newAsPoint(point *Point,father *AsPoint,end *AsPoint) *AsPoint {
	asPoint := new(AsPoint)
	asPoint.Point = point
	asPoint.Father = father
	asPoint.fVal =0
	asPoint.gVal =0
	asPoint.hVal =0
	if end != nil {
		asPoint.getFVal(end)
	}
	return asPoint
}

func (this *AsPoint) getGVal() int {
	if this.Father != nil {
		deltaX := math.Abs(float64(this.Father.x - this.x))
		deltaY := math.Abs(float64(this.Father.y - this.y))
		if deltaX == 1 && deltaY == 0 {
			this.gVal = this.Father.gVal + 10
		} else if deltaX == 0 && deltaY == 1 {
			this.gVal = this.Father.gVal + 10
		} else if deltaX == 1 && deltaY == 1 {
			this.gVal = this.Father.gVal + 14
		} else {
			fmt.Println(this.Father.x,this.Father.y,this.x,this.y)
			// panic("father point is invalid!")
		}
	}
	return this.gVal
}

func (this *AsPoint) getHVal(end *AsPoint)  int {
	this.hVal = int(math.Abs(float64(end.x-this.x)) + math.Abs(float64(end.y-this.y)))
	return this.hVal
}

func (this *AsPoint) getFVal(end *AsPoint) int {
	this.fVal = this.getGVal() + this.getHVal(end)
	return this.fVal
}

/*
查找周围节点
 */
func (this *AStarRoad) getAroundPoint(curPoint *AsPoint)  []*Point{
	points := make([]*Point,0)
	if x, y := curPoint.x, curPoint.y-1; x >= 0 && x < this.MaxX && y >= 0 && y < this.MaxY {
		points = append(points, this.AllPoints[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y-1; x >= 0 && x < this.MaxX && y >= 0 && y < this.MaxY {
		points = append(points, this.AllPoints[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y; x >= 0 && x < this.MaxX && y >= 0 && y < this.MaxY {
		points = append(points, this.AllPoints[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y+1; x >= 0 && x < this.MaxX && y >= 0 && y < this.MaxY {
		points = append(points, this.AllPoints[x][y])
	}
	if x, y := curPoint.x, curPoint.y+1; x >= 0 && x < this.MaxX && y >= 0 && y < this.MaxY {
		points = append(points, this.AllPoints[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y+1; x >= 0 && x < this.MaxX && y >= 0 && y < this.MaxY {
		points = append(points, this.AllPoints[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y; x >= 0 && x < this.MaxX && y >= 0 && y < this.MaxY {
		points = append(points, this.AllPoints[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y-1; x >= 0 && x < this.MaxX && y >= 0 && y < this.MaxY {
		points = append(points, this.AllPoints[x][y])
	}
	return points
}

func  (this *AStarRoad)FindRoad() bool {
	for len(this.OpenList) >0 {
			point := heap.Pop(&this.OpenList)
			curpoint,_ := point.(*AsPoint)
			delete(this.OpenSet,formatKey(curpoint.x,curpoint.y))
			this.CloSet[formatKey(curpoint.x,curpoint.y)]= curpoint
			points := this.getAroundPoint(curpoint)
			for _, p := range points {
					asPint := newAsPoint(p,curpoint,this.End)
					key := formatKey(asPint.x,asPint.y)
					if formatKey(this.End.x,this.End.y) == key{
						for asPint.Father != nil {
							this.Road = append(this.Road,*asPint)
							asPint.val = "$"
							asPint = asPint.Father
						}
						return true
					}
					_, ok := this.CloSet[key]
					if ok {
						continue
					}

					ap,ok := this.OpenSet[key]
					if !ok {
						heap.Push(&this.OpenList,asPint)
						this.OpenSet[key] = asPint
					}else {
						oldGVal,oldFather := ap.gVal,ap.Father
						ap.Father = curpoint
						if  ap.getGVal() > oldGVal {
							ap.Father = oldFather
							ap.gVal = oldGVal
						}

					}
			}
	}
	return false
}

func main()  {
	m := []string{
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",   //0
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //1
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //2
		". . . S . . . . . . . . . . . . . . . . . . . . . . .",  //3
		"X . X X X X X X X X X X X X X X X X X X X X X X X X X",  //4
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //5
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //6
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //7
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //8
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //9
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //10
		". . . . . E . . . . . . . . . . . . . . . . . . . . .",  //11
		"X X X X X X X X X X X X X X X X X X X X X X X X . X X",  //12
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //13
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //14
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //15
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //16
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //17
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //18
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",  //19
	}
	charMap := newMap(m)
	asRoad := newAStarRoad(5,2,11,17,charMap)
	if asRoad == nil {
		return
	}
	bl := asRoad.FindRoad()
	if bl {
		for _, val := range asRoad.Road {
			fmt.Println(val.Point,val.fVal)
		}
	}else {
		fmt.Println("not found")
	}

}