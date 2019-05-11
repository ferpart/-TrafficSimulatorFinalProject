package graph

import (
	"fmt"
	"sync"
)

type mcar string
type id string
type permit string
type isSemaphor bool
type semaphorState bool
type hasCar bool
type isFinal bool

//Node : Base structure for graph node implementation.
type Node struct {
	mcar
	id
	permit
	isSemaphor
	semaphorState
	hasCar
	isFinal
}

func (n *Node) String() string {
	return fmt.Sprintf("%s, %s", n.id, n.permit)

}

//ItemGraph : Node storage (graph).
type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

//AddNode : used to add a node to the graph.
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

//AddEdge : addes an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)

	g.lock.Unlock()
}

func (g *ItemGraph) String() {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}
