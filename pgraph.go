package pgraph

var err error

// Vertex is a node in a properties graph. With the help of Edges,
// it constructs a graph structure that can be interated through for data retrieval.
type Vertex struct {
	ID       int
	OutEdges []Edge
	InEdges  []Edge
	Label    string
	// Collection of k:v pairs
	Properties map[string]string
}

// NewVertex returns an empty Vertex struct
func NewVertex() *Vertex {
	v := &Vertex{InEdges: nil}
	return v
}

// NewRootVertex returns an empty Vertex struct with its
// InEdges set to 0, defining it as the root of the tree
func NewRootVertex() *Vertex {
	v := &Vertex{InEdges: nil}
	return v
}

// NumEdges iterates through the list of edges contained in both
// the InEdges and OutEdges attributes in a Vertex
func (v *Vertex) NumEdges(vertex Vertex) int {
	i := 0
	for i := range vertex.OutEdges {
		i++
	}
	for i := range vertex.InEdges {
		i++
	}
	return i
}

// NextVertex takes the head of the current edge and returns its
// vertex, and panics once it hits a node with no head i.e. the RootNode
func (v *Vertex) NextVertex(e *Edge) *Vertex {
	if e.Head != nil {
		return e.Head
	} else {
		panic(err)
	}
}

func (v *Vertex) GetVertextByID(id int) *Vertex {
	return &Vertex{ID: id}
}

func (v *Vertex) SetVertexProperty(key, val string, vertex *Vertex) *Vertex {
	vertex.Properties[key] = val
	return vertex
}

func (v *Vertex) GetVertexProperty(key string, vertex *Vertex) string {
	p := vertex.Properties[key]
	return p
}

// Edge is an edge between to vertices in a properties graph. It points
// to both a head and tail vertex, and can contain a map of empty interfaces,
// accesable via .Properties
type Edge struct {
	ID    int
	Tail  *Vertex
	Head  *Vertex
	Label string
	// Collection of k:v pairs
	Properties map[string]string
}

func (e *Edge) GetEdgeByID(id int) *Edge {
	return &Edge{ID: id}
}

func (e *Edge) SetEdgeProperty(key, val string, edge *Edge) *Edge {
	edge.Properties[key] = val
	return edge
}

func (e *Edge) GetEdgeProperty(key string, edge *Edge) string {
	p := edge.Properties[key]
	return p
}

// Graph represents a struct implementation of a properties graph. It consists
// of a RootNode and a list of Nodes that exist in the graph.
type Graph struct {
	RootNode *Vertex
	Node     []Vertex
	Edge     []Edge
}

// NewGraph constructs a property graph with a root Vertex
func (g *Graph) NewGraph() *Graph {
	root := NewRootVertex()
	graph := &Graph{RootNode: root}
	return graph
}

func (g *Graph) GetVertexByID(e *Edge) *Vertex {
	return &Vertex{}
}
