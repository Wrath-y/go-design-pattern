package factory_abstract

type HuaWeiFactory struct{}

type HuaWeiTelevision struct{}

type HuaWeiAirConditioner struct{}

func (*HuaWeiTelevision) Watch() {
	println("huawei tv")
}

func (*HuaWeiAirConditioner) SetTemperature(_ int) {
	println("huawei air condition")
}

func (*HuaWeiFactory) CreateTelevision() ITelevision {
	return &HuaWeiTelevision{}
}

func (*HuaWeiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaWeiAirConditioner{}
}
