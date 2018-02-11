# Voronoi diagrams in Go

A Implementation of of Steven J. Fortune's algorithm to
efficiently compute Voronoi diagrams in Go language. Based on 
a Raymond Hill's javascript implementation 
(https://raw.github.com/gorhill/Javascript-Voronoi).

## Usage

```
import "github.com/pzsz/voronoi"

func main() {
    // Sites of Voronoi diagram
	sites := []voronoi.Vertex{
		{X: 4.0, Y: 5.0},
		{X: 6.0, Y: 5.0},
		...
	}

	// Create bounding box of [0, 20] in X axis
	// and [0, 10] in Y axis
	bbox := voronoi.NewBBox(0, 20, 0, 10)

	// Compute diagram and close cells (add half edges from bounding box)
	diagram := voronoi.ComputeDiagram(sites, bbox, true)

	// Iterate over cells
	for _, cell := range diagram.Cells {
		for _, hedge := range cell.Halfedges {
		    ...
		}
	}

	// Iterate over all edges
	for _, edge := range diagram.Edges {
		if edge.LeftCell != nil && edge.RightCell != nil {
		    ...
		}
	}
}