package domain

type Helmet struct {
	InventoryItem
}

func NewHelmet(name, inventoryNumber string) Helmet {
	return Helmet{InventoryItem: newInventoryItem(name, inventoryNumber)}
}

type ChargingCabinet struct {
	InventoryItem
	dailyEnergyKWh float64
}

func NewChargingCabinet(name, inventoryNumber string, dailyEnergyKWh float64) ChargingCabinet {
	return ChargingCabinet{InventoryItem: newInventoryItem(name, inventoryNumber), dailyEnergyKWh: dailyEnergyKWh}
}

func (cc ChargingCabinet) DailyEnergyKWh() float64 {
	return cc.dailyEnergyKWh
}

func (Helmet) Kind() string          { return "Шлем" }
func (ChargingCabinet) Kind() string { return "Зарядный шкаф" }

var _ Inventoriable = Helmet{}
var _ Inventoriable = ChargingCabinet{}

var _ EnergyConsumer = ChargingCabinet{}
