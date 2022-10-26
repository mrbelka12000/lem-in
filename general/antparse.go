package code

import (
	"strconv"
)

const (
	caseBFS      int = iota
	caseDisjoint int = iota
)

//WhatPaths ..
func (all *All) WhatPaths() {
	bfsants := initqueues(all, all.BfsPath)

	all.Original.dispatchants(bfsants, all, 0)

	disants := initqueues(all, all.DisjointPaths)

	all.NewGraph.dispatchants(disants, all, 1)
}

func initqueues(all *All, puti [][]string) PathQueues {
	AntsPathQueues := PathQueues{}
	curr := 0
	for range puti {
		AntsPathQueues = append(AntsPathQueues, &AntsQueue{})
	}
	n := len(puti) - 1
	for i := 1; i <= all.Ants; i++ {
		c := len(puti[curr]) + AntsPathQueues[curr].Num
		next := curr
		if curr < n {
			next++
		} else {
			next = 0
		}
		if c > len(puti[next]) {
			ant := Ant{Imya: i, Room: all.StartRoom}
			AntsPathQueues[next].AntEnqueue(&ant)
			AntsPathQueues[next].Num++
			curr = next
		} else {
			ant := Ant{Imya: i, Room: all.StartRoom}
			AntsPathQueues[curr].AntEnqueue(&ant)
			AntsPathQueues[curr].Num++
		}
	}
	for i, v := range AntsPathQueues {
		for _, ant := range v.Ants {
			ant.Path = puti[i]
		}
	}
	return AntsPathQueues
}

func (g *Graph) dispatchants(antquques PathQueues, all *All, code int) {
	var movement MoveAnts
	node := g.getVertex(all.EndRoom)
	for node.Capacity != all.Ants {
		for _, antpaths := range antquques {
			ant := antpaths.AntDequeue()
			if ant != nil {
				movement.MEnqueu(ant)
			}
		}
		for j := 0; j < len(movement); {
			ant := movement[j]
			movement[j].Path = movement[j].Path[1:]
			if movement[j].Path[0] == all.EndRoom {
				node.Capacity++
				movement.MDeQueue(movement[j])
				j--
			}
			j++
			if code == caseBFS {
				all.Bfsres += "L" + strconv.Itoa(ant.Imya) + "-" + ant.Path[0] + " "
			} else if code == caseDisjoint {
				all.DisRes += "L" + strconv.Itoa(ant.Imya) + "-" + ant.Path[0] + " "
			}
		}
		if code == caseBFS {
			if len(all.BfsPath[0]) != 2 {
				all.Bfsres += "\n"
			}
			all.StepsBfs++
		} else if code == caseDisjoint {
			if len(all.DisjointPaths[0]) != 2 {
				all.DisRes += "\n"
			}
			all.StepsDisjoint++
		}
	}
	if len(all.BfsPath[0]) == 2 {
		all.Bfsres += "\n"
	}
	if len(all.DisjointPaths[0]) == 2 {
		all.DisRes += "\n"
	}
	node.Capacity = 0
}
