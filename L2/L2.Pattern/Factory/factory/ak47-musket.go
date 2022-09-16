package factory

/*
	Конкретный продукт в конкретном виде
*/

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}

//8888888888

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			Name:  "Musket gun",
			Power: 1,
		},
	}
}
