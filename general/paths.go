package code

import (
	"errors"
)

//Paths ..
func (g *Graph) Paths(all *All) error {

	bfspath := g.bfs(all, 1)
	if len(bfspath) == 0 {
		return errors.New("No valid paths")
	}
	for bfspath != nil {
		all.BfsPath = append(all.BfsPath, bfspath)
		bfspath = g.bfs(all, 1)
		if !checkPaths(bfspath, all.BfsPath[0]) {
			break
		}
	}

	bhandaripaths := g.bhandari(all)

	for bhandaripaths != nil {
		for i := len(bhandaripaths) - 1; i > 0; i-- {
			node := g.getVertex(bhandaripaths[i])
			nodelink := g.getVertex(bhandaripaths[i-1])
			for _, j := range node.Links {
				if j == nodelink {
					for i, v := range j.Links {
						if v == node {
							j.Links[i] = nil
						}
					}
				}
			}
		}
		all.Bhandari = append(all.Bhandari, bhandaripaths)
		bhandaripaths = g.bhandari(all)
	}

	newgraph := g.BuildNewGraph(*all)

	disjointpaths := newgraph.bfs(all, 2)

	for disjointpaths != nil {
		all.DisjointPaths = append(all.DisjointPaths, disjointpaths)
		disjointpaths = newgraph.bfs(all, 2)
		if !checkPaths(disjointpaths, all.DisjointPaths[0]) {
			break
		}
	}
	all.NewGraph = newgraph
	return nil
}

func (g *Graph) bfs(all *All, i int) []string {
	queue := Queue{}
	start := g.getVertex(all.StartRoom)
	queue.Enqueu(start)
	g.update(1)

	if i == 1 {
		for _, path := range all.BfsPath {
			for _, k := range path {
				vertex := g.getVertex(k)
				vertex.visit = true
			}
		}
	}

	if i == 2 {
		for _, path := range all.DisjointPaths {
			for _, k := range path {
				vertex := g.getVertex(k)
				vertex.visit = true
			}
		}
	}

	for !queue.IsEmpty() {
		v := queue.Dequeue()
		for _, link := range v.Links {
			if link == nil {
				continue
			}
			if link.Key == all.EndRoom {
				link.Tail = v
				path := constructPath(link, start)
				return path
			}
			if !link.visit {
				link.visit = true
				link.Tail = v
				queue.Enqueu(link)
			}
		}
	}
	return nil
}

func (g *Graph) bhandari(all *All) []string {
	queue := Queue{}
	start := g.getVertex(all.StartRoom)
	queue.Enqueu(start)
	g.update(0)
	start.flow = 0
	for !queue.IsEmpty() {

		v := queue.Dequeue()

		for _, link := range v.Links {
			if link == nil {
				continue
			}
			if link.Key == all.EndRoom {
				link.Tail = v
				path := constructPath(link, start)
				return path
			}

			if v.flow < link.flow {
				link.Tail = v
				link.flow = v.flow + 1
				queue.Enqueu(link)
			}
		}
	}
	return nil
}

func constructPath(k, start *Vertex) []string {
	path := []string{}
	path = append(path, k.Key)

	for k != start {
		path = append(path, k.Tail.Key)
		k = k.Tail
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func checkPaths(ch, wh []string) bool {
	count := 0

	for _, v := range ch {
		for _, val := range wh {
			if v == val {
				count++
			}
		}
	}

	if count == len(ch) {
		return false
	}

	return true
}
