package math

import "fmt"

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

// 无向图
type UndirectedGraph struct {
	edges      map[Point]map[Point]bool // map[a][b]有值代表a和b之间有一条边
	allDrawWay []string                 // 所有的画法
}

func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		edges: map[Point]map[Point]bool{},
	}
}

func (u *UndirectedGraph) AddEdge(a, b Point) {
	// a到b有一条边
	if u.edges[a] == nil {
		u.edges[a] = map[Point]bool{}
	}
	u.edges[a][b] = true

	// 无向图b到a也有一条边
	if u.edges[b] == nil {
		u.edges[b] = map[Point]bool{}
	}
	u.edges[b][a] = true
}

func (u *UndirectedGraph) removeEdge(a, b Point) {
	// 删除a到b的边
	delete(u.edges[a], b)
	if len(u.edges[a]) == 0 {
		// a没有边了，删除a点
		delete(u.edges, a)
	}

	// 删除b到a的边
	delete(u.edges[b], a)
	if len(u.edges[b]) == 0 {
		// b没有边了，删除b点
		delete(u.edges, b)
	}
}

func (u *UndirectedGraph) drawFrom(p Point, path string) {
	if len(u.edges) == 0 {
		// 所有的边都用过了，图画完了
		u.allDrawWay = append(u.allDrawWay, path)
		return
	}

	// 找到p能到的所有的点
	points, ok := u.edges[p]
	if !ok || len(points) == 0 {
		// 图还没画完，但是从p点找不到边了，此路不通
		return
	}

	// 记录一下所有能到的点
	allPoints := make([]Point, 0, 100)
	for point, _ := range points {
		allPoints = append(allPoints, point)
	}

	for _, point := range allPoints {
		// 随便找一条边继续往下画
		newPath := fmt.Sprintf("%s->%s", path, point.String())
		u.removeEdge(p, point)
		u.drawFrom(point, newPath)
		u.AddEdge(p, point)
	}
}

func (u *UndirectedGraph) Draw() {
	u.allDrawWay = make([]string, 0, 100)
	// 先记录所有的起始点
	allPoints := make([]Point, 0, 100)
	for point, _ := range u.edges {
		allPoints = append(allPoints, point)
	}
	// 遍历所有的起始点
	for _, point := range allPoints {
		path := point.String()
		u.drawFrom(point, path)
	}

	fmt.Println(len(u.allDrawWay))
	for _, way := range u.allDrawWay {
		fmt.Println(way)
	}
}
