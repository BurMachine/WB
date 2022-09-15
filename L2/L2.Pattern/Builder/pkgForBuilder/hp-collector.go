package pkgForBuilder

type HpCollector struct {
	Core    int
	Brand   string
	Memory  int
	Monitor int
	Graph   int
}

func (a *HpCollector) SetCore() {
	a.Core = 4
}

func (a *HpCollector) SetBrand() {
	a.Brand = "HP"
}

func (a *HpCollector) SetMemory() {
	a.Memory = 12
}

func (a *HpCollector) SetGraph() {
	a.Graph = 2
}

func (a *HpCollector) SetMonitor() {
	a.Monitor = 2
}

func (a *HpCollector) GetComputer() Computer {
	return Computer{
		Core:    a.Core,
		Brand:   a.Brand,
		Memory:  a.Memory,
		Monitor: a.Monitor,
		Graph:   a.Graph,
	}
}
