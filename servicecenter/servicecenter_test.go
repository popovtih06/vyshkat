package servicecenter

import (
	"testing"
	"vyshkat/domain"
)

func TestInspect(t *testing.T) {
	service := New()
	cases := []struct {
		name    string
		dom     domain.Rentable
		wantErr bool
	}{
		{"исправный самокат (простота 7, номер есть)", domain.NewScooter("Самокат-1", "VK-0001", 7), false},
		{"пустой инвентарный номер", domain.NewScooter("Самокат-2", "", 7), true},
		{"простота 0", domain.NewScooter("Самокат-3", "VK-0003", 0), true},
		{"простота 11", domain.NewScooter("Самокат-4", "VK-0004", 11), true},
		{"электросамокат с расходом −1.5", domain.NewElectricScooter("Самокат-5", "VK-0005", 8, -1.5), true},
		{"велосипед с простотой 1", domain.NewBicycle("Велосипед-1", "VK-0006", 1), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := service.Inspect(c.dom)
			if (err != nil) != c.wantErr {
				t.Errorf("%s, ожидали ли ошибку: %v, получили ли ошибку: %v", c.name, c.wantErr, (err != nil))
			}
		})
	}

}
