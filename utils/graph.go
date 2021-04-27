package utils

import (
	"sync"
)

// Node a single node that composes the tree
type Node struct {
	id    string
	value interface{}
}

// Graph the graph itself
type Graph struct {
	sync.RWMutex
	nodes map[string]*Node
	edges map[string]map[string]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		nodes: map[string]*Node{},
		edges: map[string]map[string]struct{}{},
	}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) {
	g.Lock()
	g.nodes[n.id] = n
	g.Unlock()
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.Lock()
	if g.edges[n1.id] == nil {
		g.edges[n1.id] = map[string]struct{}{}
	}
	g.edges[n1.id][n2.id] = struct{}{}
	if g.edges[n2.id] == nil {
		g.edges[n2.id] = map[string]struct{}{}
	}
	g.edges[n2.id][n1.id] = struct{}{}
	g.Unlock()
}

// GetNodes gets slice of nodes
func (g *Graph) GetNodes() []*Node {
	g.Lock()
	var nodes []*Node
	for _, n := range g.nodes {
		nodes = append(nodes, n)
	}
	g.Unlock()
	return nodes
}

// AddNode removes node from the graph and all its edges
func (g *Graph) RemoveNode(n *Node) {
	g.Lock()
	defer g.Unlock()
	if _, ok := g.nodes[n.id]; !ok {
		return
	}
	//remove all edges
	for connected := range g.edges[n.id] {
		delete(g.edges[connected], n.id)
	}
	delete(g.edges, n.id)
	delete(g.nodes, n.id)
}

func (g *Graph) MergeGraph(gg *Graph) {
	gg.Lock()
	defer gg.Unlock()
	for _, n := range gg.nodes {
		g.AddNode(n)
		for connected := range gg.edges[n.id] {
			g.AddEdge(n, gg.nodes[connected])
		}
	}
}

// String converts graph to string representation
func (g *Graph) String() string {
	g.RLock()
	s := ""
	for from := range g.nodes {
		s += from + " -> "
		for to := range g.edges[from] {
			s += to + " "
		}
		s += "\n"
	}
	g.RUnlock()
	return s
}
