package main

//**************************

type Builder interface {
	SetHp()
	SetTwistingMoment()

	GetAuto() Auto
}

func getAuto(name string) Builder {
	if name == "supra" {
		return newSupra()
	} else if name == "challenger" {
		return newChalle
	}
}

//***************************

type Auto struct {
	HorsePower     int
	TwistingMoment int
}
