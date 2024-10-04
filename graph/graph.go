package graph

type Graph struct {
	AjMat [][]uint8
	v     uint8
}

func New(aj [][]uint8) Graph {
	return Graph{
		AjMat: aj,
		v:     uint8(len(aj)),
	}
}
