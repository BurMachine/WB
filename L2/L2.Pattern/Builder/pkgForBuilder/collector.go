package pkgForBuilder

/*
	Collector
	Interface for assembling computers of different brands
*/
type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetGraph()
	SetMonitor()
	GetComputer() Computer
}

/*
	AsusCollectorType, HpCollectorsType const
*/
const (
	AsusCollectorType = "Asus"
	HpCollectorsType  = "HP"
)

func GetCollector(CollectorType string) Collector {
	switch CollectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorsType:
		return &HpCollector{}
	}

}
