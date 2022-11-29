package factory_abstract

func Run() {
	var factory AbstractFactory

	factory = new(HuaWeiFactory)
	factory.CreateTelevision().Watch()
	factory.CreateAirConditioner().SetTemperature(1)

	factory = new(MiFactory)
	factory.CreateTelevision().Watch()
	factory.CreateAirConditioner().SetTemperature(1)
}
