package pkgForBuilder

type AsusCollector struct {
	Core    int
	Brand   string
	Memory  int
	Monitor int
	Graph   int
}

func (a *AsusCollector) SetCore() {
	a.Core = 4
}

func (a *AsusCollector) SetBrand() {
	a.Brand = "Asus"
}

func (a *AsusCollector) SetMemory() {
	a.Memory = 8
}

func (a *AsusCollector) SetGraph() {
	a.Graph = 2
}

func (a *AsusCollector) SetMonitor() {
	a.Monitor = 1
}

func (a *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:    a.Core,
		Brand:   a.Brand,
		Memory:  a.Memory,
		Monitor: a.Monitor,
		Graph:   a.Graph,
	}
}
