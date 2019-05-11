package main

var g ItemGraph
var m map[string]int

//Graph : initializer for this class
func Graph() {

	northF := Node{nil, "Nf", "E, S, W", false, false, false, true}

	northS := Node{nil, "Ns", "all", false, false, false, false}
	northA := Node{nil, "Na", "all", false, false, false, false}
	northB := Node{nil, "Nb", "all", false, false, false, false}
	northC := Node{nil, "Nc", "all", false, false, false, false}

	eastF := Node{nil, "Ef", "N, S, W", false, false, false, true}

	eastS := Node{nil, "Es", "all", false, false, false, false}
	eastA := Node{nil, "Ea", "all", false, false, false, false}
	eastB := Node{nil, "Eb", "all", false, false, false, false}
	eastC := Node{nil, "Ec", "all", false, false, false, false}

	southF := Node{nil, "Sf", "N, E, W", false, false, false, true}

	southS := Node{nil, "Ss", "all", false, false, false, false}
	southA := Node{nil, "Sa", "all", false, false, false, false}
	southB := Node{nil, "Sb", "all", false, false, false, false}
	southC := Node{nil, "Sc", "all", false, false, false, false}

	westF := Node{nil, "Wf", "N, E, S", false, false, false, true}

	westS := Node{nil, "Ws", "all", false, false, false, false}
	westA := Node{nil, "Wa", "all", false, false, false, false}
	westB := Node{nil, "Wb", "all", false, false, false, false}
	westC := Node{nil, "Wc", "all", false, false, false, false}

	centerN := Node{nil, "N", "E, S, W", false, false, false, false}
	centerE := Node{nil, "E", "N, S, W", false, false, false, false}
	centerS := Node{nil, "S", "N, E, W", false, false, false, false}
	centerW := Node{nil, "W", "N, E, S", false, false, false, false}

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

	g.AddEdge(&northC, &northB)
	g.AddEdge(&northC, &northA)
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

	g.AddEdge(&centerS, &northF)
	g.AddEdge(&centerS, &centerN)

}

func getItemGraph() *ItemGraph {
	return &g
}

func getIndex(s string) int {
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

	return m[s]

}
