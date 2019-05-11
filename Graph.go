package main

var g ItemGraph

//Graph : initializer for this class
func Graph() {

	northF := Node{mCar: nil, "Nf", "E, S, W", false, false, false, true}

	northS := Node{mCar: nil, "Ns", "all", false, false, false, false}
	northA := Node{mCar: nil, "Na", "all", false, false, false, false}
	northB := Node{mCar: nil, "Nb", "all", false, false, false, false}
	northC := Node{mCar: nil, "Nc", "all", false, false, false, false}

	eastF := Node{mCar: nil, "Ef", "N, S, W", false, false, false, true}

	eastS := Node{mCar: nil, "Es", "all", false, false, false, false}
	eastA := Node{mCar: nil, "Ea", "all", false, false, false, false}
	eastB := Node{mCar: nil, "Eb", "all", false, false, false, false}
	eastC := Node{mCar: nil, "Ec", "all", false, false, false, false}

	southF := Node{mCar: nil, "Sf", "N, E, W", false, false, false, true}

	southS := Node{mCar: nil, "Ss", "all", false, false, false, false}
	southA := Node{mCar: nil, "Sa", "all", false, false, false, false}
	southB := Node{mCar: nil, "Sb", "all", false, false, false, false}
	southC := Node{mCar: nil, "Sc", "all", false, false, false, false}

	westF := Node{mCar: nil, "Wf", "N, E, S", false, false, false, true}

	westS := Node{mCar: nil, "Ws", "all", false, false, false, false}
	westA := Node{mCar: nil, "Wa", "all", false, false, false, false}
	westB := Node{mCar: nil, "Wb", "all", false, false, false, false}
	westC := Node{mCar: nil, "Wc", "all", false, false, false, false}

	centerN := Node{mCar: nil, "N", "E, S, W", false, false, false, false}
	centerE := Node{mCar: nil, "E", "N, S, W", false, false, false, false}
	centerS := Node{mCar: nil, "S", "N, E, W", false, false, false, false}
	centerW := Node{mCar: nil, "W", "N, E, S", false, false, false, false}

	g.AddNode(&northF)
	g.AddNode(&northA)
	g.AddNode(&northB)
	g.AddNode(&northC)

	g.AddNode(&eastF)
	g.AddNode(&eastA)
	g.AddNode(&eastB)
	g.AddNode(&eastC)

	g.AddNode(&southF)
	g.AddNode(&southA)
	g.AddNode(&southB)
	g.AddNode(&southC)

	g.AddNode(&westF)
	g.AddNode(&westA)
	g.AddNode(&westB)
	g.AddNode(&westC)

	g.AddNode(&centerN)
	g.AddNode(&centerE)
	g.AddNode(&centerS)
	g.AddNode(&centerW)

	g.AddEdge(&northC, &northB)
	g.AddEdge(&northC, &northA)
	g.AddEdge(&northA, &centerN)

	g.AddEdge(&eastC, &eastB)
	g.AddEdge(&eastB, &eastA)
	g.AddEdge(&eastA, &centerE)

	g.AddEdge(&southC, &southB)
	g.AddEdge(&southB, &southA)
	g.AddEdge(&southA, &centerS)

	g.AddEdge(&westC, &westB)
	g.AddEdge(&westB, &westA)
	g.AddEdge(&westA, &centerW)

	g.AddEdge(&centerN, &westF)
	g.AddEdge(&centerN, &centerW)

	g.AddEdge(&centerW, &southF)
	g.AddEdge(&centerW, &centerS)

	g.AddEdge(&centerS, &eastF)
	g.AddEdge(&centerS, &centerE)

	g.AddEdge(&centerS, &northF)
	g.AddEdge(&centerS, &centerN)
}
