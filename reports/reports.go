package reports

import (
	"sort"
	"vyshkat/domain"
)

type Reports struct {
	policy domain.BeginnerPolicy
}

func New(policy domain.BeginnerPolicy) Reports {
	return Reports{policy: policy}
}

func (r Reports) CountByKind(items []domain.Inventoriable) map[string]int {
	counts := make(map[string]int)
	for _, item := range items {
		counts[item.Kind()]++
	}
	return counts
}

func (r Reports) TotalEnergyKWh(items []domain.Inventoriable) float64 {
	total := 0.0
	for _, item := range items {
		if c, ok := item.(domain.EnergyConsumer); ok {
			total += c.DailyEnergyKWh()
		}
	}
	return total
}

func (r Reports) BeginnerFriendly(vehicles []domain.Rentable) []domain.Rentable {
	out := make([]domain.Rentable, 0)
	for _, v := range vehicles {
		if r.policy.Allows(v) {
			out = append(out, v)
		}
	}
	return out
}

func (r Reports) SortedByInventoryNumber(items []domain.Inventoriable) []domain.Inventoriable {
	out := make([]domain.Inventoriable, len(items))
	copy(out, items)
	sort.Slice(out, func(i, j int) bool {
		return out[i].InventoryNumber() < out[j].InventoryNumber()
	})
	return out
}
