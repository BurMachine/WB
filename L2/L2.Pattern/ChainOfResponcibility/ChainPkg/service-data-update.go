package ChainPkg

import "log"

type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		log.Println(upd.Name, ": Data from device is already update")
		upd.Next.Execute(data)
		return
	}
	log.Println("Update data from device", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(service Service) {
	upd.Next = service
}
