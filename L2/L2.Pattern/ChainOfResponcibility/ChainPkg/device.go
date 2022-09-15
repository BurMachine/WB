package ChainPkg

import "log"

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		log.Println(d.Name, ": Data from device")
		d.Next.Execute(data)
		return
	}
	log.Println("Data from device", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(service Service) {
	d.Next = service
}
