package main

import "simple/pkgForBuilder/pkgForBuilder"

func main() {
	asusCollector := pkgForBuilder.GetCollector("Asus") // Determine equipment
	hpCollector := pkgForBuilder.GetCollector("HP")

	factory := pkgForBuilder.NewFactory(asusCollector) // Determine factory to create asus computer
	AsusComputer := factory.CreateComputer()
	AsusComputer.PrintPC()

	factory.SetCollector(hpCollector) // Determine factory to create hp computer
	HpComputer := factory.CreateComputer()
	HpComputer.PrintPC()
}
