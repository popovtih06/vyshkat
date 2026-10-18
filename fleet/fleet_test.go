package fleet

import (
	"errors"
	"testing"
	"vyshkat/domain"
)

// техосмотр, который всё отклоняет
type rejectingInspector struct{}

func (rejectingInspector) Inspect(v domain.Rentable) error {
	return errors.New("техосмотр не пройден")
}

// техосмотр, который всё принимает
type acceptingInspector struct{}

func (acceptingInspector) Inspect(v domain.Rentable) error {
	return nil
}

func TestAddVehicleRejectedByInspector(t *testing.T) {
	park := NewFleet(rejectingInspector{})

	err := park.AddVehicle(domain.NewScooter("Самокат-1", "VK-0001", 7))

	if err == nil {
		t.Fatal("ожидали ошибку: техосмотр отклонил технику")
	}
	if len(park.Vehicles()) != 0 {
		t.Errorf("парк должен остаться пустым, а в нём %d единиц", len(park.Vehicles()))
	}
}

func TestAddVehicleAcceptedByInspector(t *testing.T) {
	park := NewFleet(acceptingInspector{})

	err := park.AddVehicle(domain.NewScooter("Самокат-1", "VK-0001", 7))

	if err != nil {
		t.Fatalf("техосмотр должен принять технику, получили ошибку: %v", err)
	}

	if len(park.Vehicles()) != 1 {
		t.Errorf("в парк должна была добавится единица, а в нём %d единиц", len(park.Vehicles()))
	}

}

func TestAddEquipmentSkipsInspection(t *testing.T) {
	park := NewFleet(rejectingInspector{})
	park.AddEquipment(domain.NewHelmet("Шлем-1", "VK-0011"))
	if len(park.Equipment()) != 1 {
		t.Errorf("шлем должен попасть на баланс без техосмотра, в парке %d вещей", len(park.Equipment()))
	}
}

func TestAllCombinesVehiclesAndEquipment(t *testing.T) {
	park := NewFleet(acceptingInspector{})
	err := park.AddVehicle(domain.NewScooter("Самокат-1", "VK-0001", 7))
	if err != nil {
		t.Fatalf("техосмотр должен принять технику, получили ошибку: %v", err)
	}
	park.AddEquipment(domain.NewHelmet("Шлем-1", "VK-0011"))
	if len(park.All()) != 2 {
		t.Errorf("в парк должно было добавиться две вещи, а в нём %d вещей", len(park.All()))
	}

}
