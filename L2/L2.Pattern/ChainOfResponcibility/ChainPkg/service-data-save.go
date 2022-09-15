package ChainPkg

import "log"

type ServiceDataSave struct {
	Next Service
}

func (s *ServiceDataSave) Execute(data *Data) {
	if !data.UpdateSource {
		log.Println(": Data not update")
		return
	}
	log.Println("Data save")
}

func (s *ServiceDataSave) SetNext(service Service) {
	s.Next = service
}
