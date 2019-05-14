package main

import (
	"strings"
)

var g ItemGraph
var m map[string]int
var finalPath []string

//Graph : initializer for this class
func Graph() {
	createNodes()
	mapMaker()
}

func createNodes() {
	northF := Node{nil, "northF", false, false, false, true, false}

	northS := Node{nil, "northS", false, false, false, false, false}
	northA := Node{nil, "northA", false, false, false, false, false}
	northB := Node{nil, "northB", false, false, false, false, false}
	northC := Node{nil, "northC", false, false, false, false, false}

	eastF := Node{nil, "eastF", false, false, false, true, false}

	eastS := Node{nil, "eastS", false, false, false, false, false}
	eastA := Node{nil, "eastA", false, false, false, false, false}
	eastB := Node{nil, "eastB", false, false, false, false, false}
	eastC := Node{nil, "eastC", false, false, false, false, false}

	southF := Node{nil, "southF", false, false, false, true, false}

	southS := Node{nil, "southS", false, false, false, false, false}
	southA := Node{nil, "southA", false, false, false, false, false}
	southB := Node{nil, "southB", false, false, false, false, false}
	southC := Node{nil, "southC", false, false, false, false, false}

	westF := Node{nil, "westF", false, false, false, true, false}

	westS := Node{nil, "westS", false, false, false, false, false}
	westA := Node{nil, "westA", false, false, false, false, false}
	westB := Node{nil, "westB", false, false, false, false, false}
	westC := Node{nil, "westC", false, false, false, false, false}

	centerN := Node{nil, "centerN", false, false, false, false, false}
	centerE := Node{nil, "centerE", false, false, false, false, false}
	centerS := Node{nil, "centerS", false, false, false, false, false}
	centerW := Node{nil, "centerW", false, false, false, false, false}

	g.AddNode(&northA) // 0
	g.AddNode(&northB) // 1
	g.AddNode(&northC) // 2
	g.AddNode(&northS) // 3
	g.AddNode(&northF) // 4

	g.AddNode(&eastA) // 5
	g.AddNode(&eastB) // 6
	g.AddNode(&eastC) // 7
	g.AddNode(&eastS) // 8
	g.AddNode(&eastF) // 9

	g.AddNode(&southA) // 10
	g.AddNode(&southB) // 11
	g.AddNode(&southC) // 12
	g.AddNode(&southS) // 13
	g.AddNode(&southF) // 14

	g.AddNode(&westA) // 15
	g.AddNode(&westB) // 16
	g.AddNode(&westC) // 17
	g.AddNode(&westS) // 18
	g.AddNode(&westF) // 19

	g.AddNode(&centerN) // 20
	g.AddNode(&centerE) // 21
	g.AddNode(&centerS) // 22
	g.AddNode(&centerW) // 23

	g.edges = make(map[*Node]([]*Node))

	g.AddEdge(&northC, &northB)
	g.AddEdge(&northB, &northA)
	g.AddEdge(&northA, &northS)
	g.AddEdge(&northS, &centerN)

	g.AddEdge(&eastC, &eastB)
	g.AddEdge(&eastB, &eastA)
	g.AddEdge(&eastA, &eastS)
	g.AddEdge(&eastS, &centerE)

	g.AddEdge(&southC, &southB)
	g.AddEdge(&southB, &southA)
	g.AddEdge(&southA, &southS)
	g.AddEdge(&southS, &centerS)

	g.AddEdge(&westC, &westB)
	g.AddEdge(&westB, &westA)
	g.AddEdge(&westA, &westS)
	g.AddEdge(&westS, &centerW)

	g.AddEdge(&centerN, &westF)
	g.AddEdge(&centerN, &centerW)

	g.AddEdge(&centerW, &southF)
	g.AddEdge(&centerW, &centerS)

	g.AddEdge(&centerS, &eastF)
	g.AddEdge(&centerS, &centerE)

	g.AddEdge(&centerE, &northF)
	g.AddEdge(&centerE, &centerN)
}

func mapMaker() {
	m = make(map[string]int)

	m["northA"] = 0
	m["northB"] = 1
	m["northC"] = 2
	m["northS"] = 3
	m["northF"] = 4
	m["eastA"] = 5
	m["eastB"] = 6
	m["eastC"] = 7
	m["eastS"] = 8
	m["eastF"] = 9
	m["southA"] = 10
	m["southB"] = 11
	m["southC"] = 12
	m["southS"] = 13
	m["southF"] = 14
	m["westA"] = 15
	m["westB"] = 16
	m["westC"] = 17
	m["westS"] = 18
	m["westF"] = 19
	m["centerN"] = 20
	m["centerE"] = 21
	m["centerS"] = 22
	m["centerW"] = 23
}

func getPath(start, end string) []string {
	for _, n := range g.nodes {
		n.setVisited(false)
	}

	path := make([]string, 0)
	finalPath = path
	getPathL(start, start, end, &path)

	return getFinalPath()
}

func getPathL(staticStart, start, end string, path *[]string) {

	g.nodes[getIndex(start)].setVisited(true)
	*path = append(*path, start)

	if strings.Compare(start, end) == 0 {
		finalPath = *path
		return
	}
	for _, v := range g.edges[g.nodes[getIndex(start)]] {
		if v.getVisited() == false && validPath(staticStart, v.getID()) {

			getPathL(staticStart, v.getID(), end, path)

		}
	}

	*path = (*path)[:len(*path)-1]
	g.nodes[getIndex(start)].setVisited(false)

}

func validPath(start, end string) bool {
	startL := strings.SplitN(start, "", 2)
	switch startL[0] {
	case "n":
		if strings.Compare(end, "centerE") == 0 {
			return false
		}
		break
	case "e":
		if strings.Compare(end, "centerS") == 0 {
			return false
		}
		break
	case "s":
		if strings.Compare(end, "centerW") == 0 {
			return false
		}
		break
	case "w":
		if strings.Compare(end, "centerN") == 0 {
			return false
		}
		break
	}
	return true
}

func getFinalPath() []string {
	return finalPath
}

func getItemGraph() *ItemGraph {
	return &g
}

func getIndex(s string) int {
	return m[s]
}
