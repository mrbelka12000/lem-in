package code

//All ..
type All struct {
	Ants          int
	StartRoom     string
	EndRoom       string
	Map           string
	Bfsres        string
	DisRes        string
	Coords        []string
	Rooms         []string
	Links         []string
	Bhandari      [][]string
	DisjointPaths [][]string
	BfsPath       [][]string
	StepsBfs      int
	StepsDisjoint int
	Original      Graph
	NewGraph      Graph
}

//Graph ..
type Graph struct {
	Vertices []*Vertex
}

//Vertex ..
type Vertex struct {
	Key      string
	Links    []*Vertex
	Tail     *Vertex
	visit    bool
	flow     int
	Capacity int
}

//AntsQueue ..
type AntsQueue struct {
	Ants []*Ant
	Num  int
}

//Ant ..
type Ant struct {
	Imya  int
	Room  string
	Check bool
	Path  []string
}
