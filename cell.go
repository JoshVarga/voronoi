// Copyright 2013 Przemyslaw Szczepaniak.
// MIT License: See https://github.com/gorhill/Javascript-Voronoi/LICENSE.md

// Author: Przemyslaw Szczepaniak (przeszczep@gmail.com)
// Port of Raymond Hill's (rhill@raymondhill.net) javascript implementation 
// of Steven Forune's algorithm to compute Voronoi diagrams

package voronoi

import "sort"

// Cell of voronoi diagram
type Cell struct {
	// Site of the cell
	Site Vertex
	// Array of halfedges sorted counterclockwise
	Halfedges []*Halfedge
}

func newCell(site Vertex) *Cell {
	return &Cell{Site: site}
}

func (t *Cell) prepare() int {
	halfedges := t.Halfedges
	iHalfedge := len(halfedges) - 1

	for ; iHalfedge >= 0; iHalfedge-- {
		edge := halfedges[iHalfedge].Edge

		if edge.Vb.Vertex == NO_VERTEX || edge.Va.Vertex == NO_VERTEX {
			halfedges[iHalfedge] = halfedges[len(halfedges)-1]
			halfedges = halfedges[:len(halfedges)-1]
		}
	}

	sort.Sort(halfedgesByAngle{halfedges})
	t.Halfedges = halfedges
	return len(halfedges)
}
