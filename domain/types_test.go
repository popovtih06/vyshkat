package domain

import (
	"math"
	"testing"
)

// Все шесть видов должны называться по-разному.
// Если при копипасте Kind() где-то продублировался,
// отчёт по видам склеит два вида в один - а этот тест упадёт.
func TestKindsAreDistinct(t *testing.T) {
	items := []Inventoriable{
		NewScooter("с", "VK-1", 7),
		NewElectricScooter("эс", "VK-2", 7, 1.2),
		NewBicycle("в", "VK-3", 7),
		NewElectricBicycle("эв", "VK-4", 7, 2.1),
		NewHelmet("ш", "VK-5"),
		NewChargingCabinet("шк", "VK-6", 4.5),
	}

	kinds := make(map[string]bool)
	for _, item := range items {
		kinds[item.Kind()] = true
	}

	if len(kinds) != len(items) {
		t.Errorf("ожидали %d разных видов, получили %d: %v", len(items), len(kinds), kinds)
	}
}

// Потребители энергии должны возвращать ровно тот расход,
// который передали в конструктор.
func TestEnergyConsumersReportConsumption(t *testing.T) {
	cases := []struct {
		name string
		item EnergyConsumer
		want float64
	}{
		{"электросамокат", NewElectricScooter("эс", "VK-1", 7, 1.2), 1.2},
		{"электровелосипед", NewElectricBicycle("эв", "VK-2", 7, 2.1), 2.1},
		{"зарядный шкаф", NewChargingCabinet("шк", "VK-3", 4.5), 4.5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.item.DailyEnergyKWh()

			if math.Abs(got-c.want) > 1e-9 {
				t.Errorf("расход: получили %v, ожидали %v", got, c.want)
			}
		})
	}
}

// Название и инвентарный номер должны доезжать до конкретного типа
// через встраивание InventoryItem.
func TestInventoryItemFields(t *testing.T) {
	scooter := NewScooter("Самокат-1", "VK-0001", 7)

	if scooter.Name() != "Самокат-1" {
		t.Errorf("название: получили %q, ожидали %q", scooter.Name(), "Самокат-1")
	}
	if scooter.InventoryNumber() != "VK-0001" {
		t.Errorf("инвентарный номер: получили %q, ожидали %q", scooter.InventoryNumber(), "VK-0001")
	}
	if scooter.Easiness() != 7 {
		t.Errorf("простота: получили %d, ожидали %d", scooter.Easiness(), 7)
	}
}
