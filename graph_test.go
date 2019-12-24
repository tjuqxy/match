package math

import "testing"

func TestUndirectedGraph_Draw(t *testing.T) {
	graph := NewUndirectedGraph()
	graph.AddEdge(Point{0, 0}, Point{0, 1})
	graph.AddEdge(Point{0, 0}, Point{1, 0})
	graph.AddEdge(Point{0, 1}, Point{1, 0})
	graph.AddEdge(Point{0, 1}, Point{1, 1})
	graph.AddEdge(Point{0, 1}, Point{0, 2})
	graph.AddEdge(Point{0, 2}, Point{1, 1})
	graph.AddEdge(Point{0, 2}, Point{1, 2})
	graph.AddEdge(Point{1, 2}, Point{1, 1})
	graph.AddEdge(Point{1, 1}, Point{1, 0})
	graph.Draw()
}
