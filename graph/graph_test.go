package graph_test

import (
	"testing"

	"github.com/Carter907/golang-dsa/graph"
)

func TestGraph(t *testing.T) {
	graph := graph.New(
		[][]uint8{
			{0, 1},
			{1, 0},
		},
	)
	t.Log(graph.AjMat)
}
