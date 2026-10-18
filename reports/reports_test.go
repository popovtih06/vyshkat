package reports

import (
	"math"
	"testing"
	"vyshkat/domain"
)

func TestTotalEnergyKWh(t *testing.T) {
	report := New(domain.NewBeginnerPolicy(6))

	t.Run("пустой парк", func(t *testing.T) {
		got := report.TotalEnergyKWh(nil)

		if got != 0 {
			t.Errorf("для пустого парка ожидали 0, получили %v", got)
		}
	})

	t.Run("только потребители энергии попадают в сумму", func(t *testing.T) {
		items := []domain.Inventoriable{
			domain.NewElectricScooter("Электросамокат-1", "VK-0001", 8, 1.2),
			domain.NewHelmet("Шлем-1", "VK-0100"),
			domain.NewChargingCabinet("Шкаф-1", "VK-0200", 4.5),
		}

		got := report.TotalEnergyKWh(items)
		want := 5.7

		if math.Abs(got-want) > 1e-9 {
			t.Errorf("суммарная энергия: получили %v, ожидали %v", got, want)
		}
	})
}

func TestCountByKind(t *testing.T) {
	report := New(domain.NewBeginnerPolicy(6))

	items := []domain.Inventoriable{
		domain.NewScooter("Самокат-1", "VK-0001", 7),
		domain.NewScooter("Самокат-2", "VK-0002", 5),
		domain.NewHelmet("Шлем-1", "VK-0100"),
	}

	counts := report.CountByKind(items)

	if len(counts) != 2 {
		t.Errorf("ожидали 2 вида, получили %d: %v", len(counts), counts)
	}
	if counts["Самокат"] != 2 {
		t.Errorf("самокатов должно быть 2, получили %d", counts["Самокат"])
	}
	if counts["Шлем"] != 1 {
		t.Errorf("шлемов должно быть 1, получили %d", counts["Шлем"])
	}
}

func TestBeginnerFriendly(t *testing.T) {
	report := New(domain.NewBeginnerPolicy(6))
	vehicles := []domain.Rentable{
		domain.NewScooter("Самокат-1", "VK-0001", 4),
		domain.NewScooter("Самокат-2", "VK-0002", 6),
		domain.NewScooter("Самокат-3", "VK-0003", 8),
	}
	got := report.BeginnerFriendly(vehicles)
	if len(got) != 2 {
		t.Fatalf("Ожидалась длина 2, получено %d", len(got))
	}
	if got[0].InventoryNumber() != "VK-0002" || got[1].InventoryNumber() != "VK-0003" {
		t.Errorf("ожидали VK-0002 и VK-0003, получили %s и %s", got[0].InventoryNumber(), got[1].InventoryNumber())
	}
}

func TestSortedByInventoryNumber(t *testing.T) {
	report := New(domain.NewBeginnerPolicy(6))

	items := []domain.Inventoriable{
		domain.NewScooter("Самокат-1", "VK-0003", 7),
		domain.NewScooter("Самокат-2", "VK-0001", 5),
		domain.NewHelmet("Шлем-1", "VK-0002"),
	}
	want := []string{"VK-0001", "VK-0002", "VK-0003"}
	got := report.SortedByInventoryNumber(items)

	if len(got) != len(want) {
		t.Fatalf("ожидали %d элементов, получили %d", len(want), len(got))
	}
	for i := range want {
		if got[i].InventoryNumber() != want[i] {
			t.Errorf("позиция %d: ожидали %s, получили %s", i, want[i], got[i].InventoryNumber())
		}
	}
}
