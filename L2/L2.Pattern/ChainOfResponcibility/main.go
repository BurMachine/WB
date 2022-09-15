package main

import "simple/Chains/ChainPkg"

/*

	Chain of Responsibilities
	Суть паттерна в том, что каждое звено решает какую-то свою задачу
	и передает новую задачу последующему звену по цепочке
*/
func main() {
	device := &ChainPkg.Device{Name: "device1"}
	updateSvc := &ChainPkg.UpdateDataService{Name: "device1"}
	saveSvc := &ChainPkg.ServiceDataSave{}
	device.SetNext(updateSvc)
	updateSvc.SetNext(saveSvc)
	data := &ChainPkg.Data{}
	device.Execute(data)
}
