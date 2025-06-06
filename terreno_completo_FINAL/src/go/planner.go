// internal/planner/planner.go
package planner

import (
	"encoding/json" // Para unmarshal de spec
	"fmt"
	"loglineos/internal/parser"
	"os"
	"strings"

	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

// Node representa um bloco no grafo do workflow
type Node struct {
	id   int64	// ID interno para o grafo
	Block parser.Block
	DOTID string // ID para o formato DOT (o ID do bloco)
}

// ID retorna o ID interno do nó
func (n Node) ID() int64 {
	return n.id
}

// DOTID retorna o ID para o formato DOT
func (n Node) DOTID() string {
	return n.DOTID
}

// Attributes de nó para Graphviz (cor, label, etc.)
func (n Node) Attributes() []dot.Attribute {
	label := fmt.Sprintf("%s\\n(%s)", n.Block.ID, n.Block.Type)
	color := "black"
	shape := "box"
	fillcolor := "lightgrey" // Default fill color

	switch n.Block.Type {
	case "register":
		color = "blue"
		fillcolor = "lightblue"
	case "affair":
		color = "darkgreen"
		fillcolor = "lightgreen"
	case "observe":
		color = "orange"
		fillcolor = "lightyellow"
	case "commit":
		color = "purple"
		fillcolor = "lavender"
	case "mechanical":
		color = "gray"
		fillcolor = "lightgray"
	case "cognitive":
		color = "brown"
		fillcolor = "peachpuff"
	case "loop": // Loop blocks get a special shape
		shape = "doubleoctagon"
		color = "darkred"
		fillcolor = "pink"
	case "await": // Await blocks get a special shape
		shape = "parallelogram"
		color = "darkblue"
		fillcolor = "lightcyan"
	}

	return []dot.Attribute{
		{Key: "label", Value: dot.Quote(label)},
		{Key: "color", Value: dot.Quote(color)},
		{Key: "shape", Value: dot.Quote(shape)},
		{Key: "style", Value: dot.Quote("filled")},
		{Key: "fillcolor", Value: dot.Quote(fillcolor)},
	}
}

// BuildGraph constrói um grafo Gonum a partir de um workflow LogLine
func BuildGraph(workflow *parser.Workflow) (*simple.DirectedGraph, error) {
	g := simple.NewDirectedGraph()
	blockNodes := make(map[string]*Node)
	var nodeID int64 = 0

	// Função auxiliar para adicionar blocos recursivamente e criar nós no grafo
	var addBlocksRecursively func(blocks []parser.Block)
	addBlocksRecursively = func(blocks []parser.Block) {
		for _, block := range blocks {
			if _, exists := blockNodes[block.ID]; !exists {
				node := &Node{id: nodeID, Block: block, DOTID: block.ID}
				g.AddNode(node)
				blockNodes[block.ID] = node
				nodeID++
			}

			if block.Type == "loop" && block.Tasks != nil && len(block.Tasks) > 0 {
				addBlocksRecursively(block.Tasks) // Adiciona blocos aninhados do loop
			}
		}
	}
	addBlocksRecursively(workflow.Tasks)


	// Função auxiliar para processar dependências recursivamente e adicionar arestas
	var processDependenciesRecursively func(blocks []parser.Block)
	processDependenciesRecursively = func(blocks []parser.Block) {
		for i, block := range blocks {
			currentNode := blockNodes[block.ID]

			// 1. Dependência sequencial implícita
			if i > 0 {
				prevBlock := blocks[i-1]
				prevNode := blockNodes[prevBlock.ID]
				if !g.HasEdgeFromTo(prevNode.ID(), currentNode.ID()) {
					g.SetEdge(g.NewEdge(prevNode, currentNode))
				}
			}

			// 2. Dependências explícitas (when, inputs)
			if block.When != "" {
				// Para 'when', tentamos inferir dependências por TaskID referenciados
				// Esta é uma inferência básica; expressões complexas exigiriam análise de AST da expressão
				for prevID, prevNode := range blockNodes {
					if strings.Contains(block.When, prevID) {
						if !g.HasEdgeFromTo(prevNode.ID(), currentNode.ID()) {
							g.SetEdge(g.NewEdge(prevNode, currentNode))
						}
					}
				}
			}

			// 3. Dependência para 'await' blocks
			if block.Type == "await" {
				var awaitSpec parser.AwaitSpec
				if err := json.Unmarshal(block.Spec, &awaitSpec); err == nil {
					targetID := awaitSpec.TaskID
					if targetID == "" {
						// Se await for por span_id, e não task_id, é mais difícil
						// inferir no planner sem histórico. Pulamos por agora.
					}
					if targetID != "" {
						if targetNode, ok := blockNodes[targetID]; ok {
							if !g.HasEdgeFromTo(targetNode.ID(), currentNode.ID()) {
								g.SetEdge(g.NewEdge(targetNode, currentNode))
							}
						}
					}
				}
			}

			// 4. Para blocos 'commit', eles dependem de seus 'inputs'
			if block.Type == "commit" {
				var commitSpec parser.CommitSpec
				if err := json.Unmarshal(block.Spec, &commitSpec); err == nil {
					for _, inputID := range commitSpec.Inputs {
						if inputNode, ok := blockNodes[inputID]; ok {
							if !g.HasEdgeFromTo(inputNode.ID(), currentNode.ID()) {
								g.SetEdge(g.NewEdge(inputNode, currentNode))
							}
						}
					}
				}
			}

            // 5. Para blocos 'cognitive' com 'action_on_output'
            if block.Type == "cognitive" {
                var cognitiveSpec parser.CognitiveSpec
                if err := json.Unmarshal(block.Spec, &cognitiveSpec); err == nil && cognitiveSpec.ActionOnOutput != nil {
                    targetID := cognitiveSpec.ActionOnOutput.TriggerBlockID
                    if targetNode, ok := blockNodes[targetID]; ok {
                        if !g.HasEdgeFromTo(currentNode.ID(), targetNode.ID()) {
                            g.SetEdge(g.NewEdge(currentNode, targetNode)) // A aresta vai do cognitive para o bloco disparado
                        }
                    }
                }
            }


			// 6. Se for um bloco 'loop', seus blocos aninhados também têm dependências
			if block.Type == "loop" && block.Tasks != nil && len(block.Tasks) > 0 {
				loopNode := blockNodes[block.ID]
				for _, innerBlock := range block.Tasks {
					innerNode := blockNodes[innerBlock.ID]
					// O loop block "contém" os inner blocks, mas também há dependência explícita
					// do loop para o primeiro bloco aninhado para visualização de fluxo.
					if !g.HasEdgeFromTo(loopNode.ID(), innerNode.ID()) {
						g.SetEdge(g.NewEdge(loopNode, innerNode))
					}
				}
				processDependenciesRecursively(block.Tasks) // Chamada recursiva para dependências aninhadas
			}
		}
	}

	// Começa o processamento de dependências da lista de blocos raiz
	processDependenciesRecursively(workflow.Tasks)

	return g, nil
}

// GenerateDOTFile gera um arquivo .dot a partir do grafo
func GenerateDOTFile(graph *simple.DirectedGraph, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("planner: falha ao criar arquivo DOT: %w", err)
	}
	defer file.Close()

	_, err = dot.Marshal(graph, graph.Name(), "", "", file)
	if err != nil {
		return fmt.Errorf("planner: falha ao Marshal grafo para DOT: %w", err)
	}
	return nil
}
