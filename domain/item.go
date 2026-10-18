package domain

type InventoryItem struct {
	name            string
	inventoryNumber string
}

func newInventoryItem(name, inventoryNumber string) InventoryItem {
	return InventoryItem{name: name, inventoryNumber: inventoryNumber}
}

func (i InventoryItem) Name() string {
	return i.name
}

func (i InventoryItem) InventoryNumber() string {
	return i.inventoryNumber
}

type Inventoriable interface {
	InventoryNumber() string
	Name() string
	Kind() string
}

type EnergyConsumer interface {
	DailyEnergyKWh() float64
}
