package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldCreateGraphWithThreeNodesConnectedWithEachOtherAndRemoveOneNodeAfterwards(t *testing.T) {
	// Given
	n1 := &Node{id: "n-1"}
	n2 := &Node{id: "n-2"}
	n3 := &Node{id: "n-3"}

	graph := NewGraph()

	// When
	graph.AddNode(n1)
	graph.AddNode(n2)
	graph.AddNode(n3)

	graph.AddEdge(n1, n2)
	graph.AddEdge(n1, n3)
	graph.AddEdge(n2, n3)

	// Then
	fmt.Println(graph)
	assert.Equal(t, []*Node{n1, n2, n3}, graph.GetNodes(), "expected to get slice of 3 nodes")

	// When
	graph.RemoveNode(n3)

	// Then
	fmt.Println(graph)
	assert.Equal(t, []*Node{n1, n2}, graph.GetNodes(), "expected to get slice of 2 nodes after removal")
}

func Test_ShouldMergeTwoGraphs(t *testing.T) {
	// Given
	n1 := &Node{id: "n-1"}
	n2 := &Node{id: "n-2"}
	n3 := &Node{id: "n-3"}

	graph := NewGraph()
	graph.AddNode(n1)
	graph.AddNode(n2)
	graph.AddNode(n3)
	graph.AddEdge(n1, n2)
	graph.AddEdge(n1, n3)
	graph.AddEdge(n2, n3)

	graphToMerge := NewGraph()
	graphToMerge.AddNode(n1)
	graphToMerge.AddNode(n2)
	graphToMerge.AddEdge(n1, n2)

	// When
	graph.MergeGraph(graphToMerge)

	// Then
	fmt.Println(graph)
	assert.Equal(t, []*Node{n1, n2, n3}, graph.GetNodes(), "expected to get slice of 3 nodes")
}
