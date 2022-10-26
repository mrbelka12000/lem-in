package code

import (
	"fmt"
	"strings"
)

//ConstructGraph строим граф
func (all *All) ConstructGraph() (Graph, error) {
	graph := Graph{}
	for _, rooms := range all.Rooms {
		if err := graph.addVertex(rooms); err != nil {
			return graph, err
		}
	}
	for _, links := range all.Links {
		link := strings.Split(links, "-")
		if err := graph.addEdges(link[0], link[1]); err != nil {
			return graph, err
		}
		graph.addEdges(link[1], link[0])
	}
	all.Original = graph
	return graph, nil
}

func (g *Graph) addVertex(k string) error {
	if contain(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added it is an existing Key", k)
		return err
	}
	g.Vertices = append(g.Vertices, &Vertex{Key: k})
	return nil
}

func (g *Graph) addEdges(from, to string) error {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge %v --> %v ", from, to)
		return err
	} else if contain(fromVertex.Links, to) {
		err := fmt.Errorf("Existing edge %v --> %v ", from, to)
		return err
	} else {
		fromVertex.Links = append(fromVertex.Links, toVertex)
		return nil
	}
}

func contain(s []*Vertex, k string) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

func (g *Graph) getVertex(k string) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == k {
			return v
		}
	}
	return nil
}

//Print показывается весь граф
func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\n Vertex %v :", v.Key)
		for _, val := range v.Links {
			fmt.Printf(" %v ", val.Key)
		}
	}
	fmt.Println()
}

func (g *Graph) update(i int) {
	for _, v := range g.Vertices {
		if i == 0 {
			v.flow = 9223372036854775807
		} else {
			v.visit = false
		}
	}
}

//BuildNewGraph ..
func (g *Graph) BuildNewGraph(all All) Graph {
	var bhandari Graph
	for _, path := range all.Bhandari {
		for _, v := range path {
			bhandari.addVertex(v)
		}
		for i := 0; i < len(path)-1; i++ {
			bhandari.addEdges(path[i], path[i+1])
		}
	}
	bhandari.removeDeadLinks(all)
	return bhandari
}

func (g *Graph) removeDeadLinks(all All) {
	for _, paths := range all.Bhandari {
		for _, k := range paths {
			node := g.getVertex(k)
			for i, link := range node.Links {
				if link != nil {
					for ind, linkback := range link.Links {
						if linkback == node {
							link.Links[ind] = nil
							node.Links[i] = nil
						}
					}
				}
			}
		}
	}
}
