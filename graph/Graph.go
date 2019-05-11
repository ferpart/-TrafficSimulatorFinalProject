package graph

var g ItemGraph

func Graph() {
	northF := Node{"", "N", "E, S, W", false, false, false, true}
	northS := Node{"", "N", "", true, false, false, false}
	northA := Node{"car", "N", "", false, false, true, false}
	northB := Node{{"car", "N", "", false, false, true, false}
	northC := Node{"car", "N", "", false, false, true, false}

	eastF := Node{"", "E", "N, S, W", false, false, false, true}
	eastS := Node{"", "E", "", true, false, false, false}
	eastA := Node{"car","E", "", false, false, true, false}
	eastB := Node{"car","E", "", false, false, true, false}
	eastC := Node{"car","E", "", false, false, true, false}

	southF := Node{"", "S", "N, E, W", false, false, false, true}
	southS := Node{"", "S", "", true, false, false, false}
	southA := Node{"car","S", "", false, false, true, false}
	southB := Node{"car","S", "", false, false, true, false}
	southC := Node{"car","S", "", false, false, true, false}

	westF := Node{"", "W", "N, E, S", false, false, false, true}
	westS := Node{"", "W", "", true, false, false, false}
	westA := Node{"car","W", "", false, false, true, false}
	westB := Node{"car","W", "", false, false, true, false}
	westC := Node{"car","W", "", false, false, true, false}

	centerN := Node{"","N", "E, S, W",false, false, false, false}
	centerE := Node{"","E", "N, S, W",false, false, false, false}
	centerS := Node{"","S", "N, E, W",false, false, false, false}
	centerW := Node{"","W", "N, E, S",false, false, false, false}

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
	g.AddEdge(&centerW, &centerF)

	g.AddEdge(&centerS, &eastF)
	g.AddEdge(&centerS, &centerE)

	g.AddEdge(&centerS, &northF)
	g.AddEdge(&centerS, &centerN)

}
