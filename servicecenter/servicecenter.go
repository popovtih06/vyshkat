package servicecenter

import (
	"errors"
	"fmt"
	"vyshkat/domain"
)

type ServiceCenter struct{}

func New() ServiceCenter {
	return ServiceCenter{}
}

func (sc ServiceCenter) Inspect(v domain.Rentable) error {
	if v.InventoryNumber() == "" {
		return errors.New("инвентарный номер не указан")
	}
	if v.Easiness() < 1 || v.Easiness() > 10 {
		return fmt.Errorf("простота %d вне диапазона 1–10", v.Easiness())
	}
	if c, ok := v.(domain.EnergyConsumer); ok {
		if c.DailyEnergyKWh() < 0 {
			return fmt.Errorf("отрицательный расход энергии: %.2f", c.DailyEnergyKWh())
		}
	}
	return nil
}
