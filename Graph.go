package main

var g ItemGraph

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

	g.AddNode(&northF)
	g.AddNode(&northS)
	g.AddNode(&northA)
	g.AddNode(&northB)
	g.AddNode(&northC)

	g.AddNode(&eastF)
	g.AddNode(&eastS)
	g.AddNode(&eastA)
	g.AddNode(&eastB)
	g.AddNode(&eastC)

	g.AddNode(&southF)
	g.AddNode(&southS)
	g.AddNode(&southA)
	g.AddNode(&southB)
	g.AddNode(&southC)

	g.AddNode(&westF)
	g.AddNode(&westS)
	g.AddNode(&westA)
	g.AddNode(&westB)
	g.AddNode(&westC)

	g.AddNode(&centerN)
	g.AddNode(&centerE)
	g.AddNode(&centerS)
	g.AddNode(&centerW)

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
