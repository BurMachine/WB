package pkgForBuilder

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (f *Factory) SetCollector(collector Collector) {
	f.Collector = collector
}

/*
	CreateComputer - a method that will sequentially perform the appropriate steps
	Separates the construction of a complex object from its representation so that
	the same construction process can result in different representations.
	(Отделяет конструирование сложного объекта от его представления так, что в результате
	одного и того же процесса конструирования могут получаться разные представления.)
*/
func (f *Factory) CreateComputer() Computer {
	f.Collector.SetCore()
	f.Collector.SetBrand()
	f.Collector.SetMemory()
	f.Collector.SetGraph()
	f.Collector.SetMonitor()
	return f.Collector.GetComputer()
}
