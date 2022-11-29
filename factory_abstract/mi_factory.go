package factory_abstract

type MiFactory struct{}

type MiTelevision struct{}

type MiAirConditioner struct{}

func (*MiTelevision) Watch() {
	println("mi tv")
}

func (*MiAirConditioner) SetTemperature(_ int) {
	println("mi air condition")
}

func (*MiFactory) CreateTelevision() ITelevision {
	return &MiTelevision{}
}

func (*MiFactory) CreateAirConditioner() IAirConditioner {
	return &MiAirConditioner{}
}
