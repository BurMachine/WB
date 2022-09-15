package pkgForBuilder

import "fmt"

type Computer struct {
	Core    int
	Brand   string
	Memory  int
	Monitor int
	Graph   int
}

func (pc *Computer) PrintPC() {
	fmt.Printf("brand:%s; Core: %d; Memory: %d; Graph: %d; Monitor: %d", pc.Brand, pc.Core, pc.Memory, pc.Graph, pc.Monitor)
	print("\n")

}
