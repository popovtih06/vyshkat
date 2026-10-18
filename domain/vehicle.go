package domain

type Rentable interface {
	Inventoriable
	Easiness() int
}

type RentalVehicle struct {
	InventoryItem
	easiness int
}

func (rv RentalVehicle) Easiness() int { return rv.easiness }

func newRentalVehicle(name, inventoryNumber string, easiness int) RentalVehicle {
	return RentalVehicle{
		InventoryItem: newInventoryItem(name, inventoryNumber),
		easiness:      easiness}
}

type Scooter struct {
	RentalVehicle
}

func NewScooter(name, inventoryNumber string, easiness int) Scooter {
	return Scooter{
		RentalVehicle: newRentalVehicle(name, inventoryNumber, easiness),
	}
}

type ElectricScooter struct {
	RentalVehicle
	dailyEnergyKWh float64
}

func (es ElectricScooter) DailyEnergyKWh() float64 {
	return es.dailyEnergyKWh
}

func NewElectricScooter(name, inventoryNumber string, easiness int, dailyEnergyKWh float64) ElectricScooter {
	return ElectricScooter{RentalVehicle: newRentalVehicle(name, inventoryNumber, easiness),
		dailyEnergyKWh: dailyEnergyKWh,
	}
}

type Bicycle struct {
	RentalVehicle
}

func NewBicycle(name, inventoryNumber string, easiness int) Bicycle {
	return Bicycle{RentalVehicle: newRentalVehicle(name, inventoryNumber, easiness)}
}

type ElectricBicycle struct {
	RentalVehicle
	dailyEnergyKWh float64
}

func (eb ElectricBicycle) DailyEnergyKWh() float64 {
	return eb.dailyEnergyKWh
}

func NewElectricBicycle(name, inventoryNumber string, easiness int, dailyEnergyKWh float64) ElectricBicycle {
	return ElectricBicycle{
		RentalVehicle:  newRentalVehicle(name, inventoryNumber, easiness),
		dailyEnergyKWh: dailyEnergyKWh,
	}
}

func (Scooter) Kind() string         { return "Самокат" }
func (ElectricScooter) Kind() string { return "Электросамокат" }
func (Bicycle) Kind() string         { return "Велосипед" }
func (ElectricBicycle) Kind() string { return "Электровелосипед" }

var _ Inventoriable = Scooter{}
var _ Inventoriable = ElectricScooter{}
var _ Inventoriable = Bicycle{}
var _ Inventoriable = ElectricBicycle{}

var _ EnergyConsumer = ElectricScooter{}
var _ EnergyConsumer = ElectricBicycle{}

var _ Rentable = Scooter{}
var _ Rentable = ElectricScooter{}
var _ Rentable = ElectricBicycle{}
var _ Rentable = Bicycle{}
