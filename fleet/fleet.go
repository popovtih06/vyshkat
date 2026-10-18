package fleet

import "vyshkat/domain"

type Inspector interface {
	Inspect(v domain.Rentable) error
}

type Fleet struct {
	inspector Inspector
	vehicles  []domain.Rentable
	equipment []domain.Inventoriable
}

func NewFleet(inspector Inspector) *Fleet {
	return &Fleet{inspector: inspector}
}

func (f *Fleet) AddVehicle(v domain.Rentable) error {
	if err := f.inspector.Inspect(v); err != nil {
		return err
	}
	f.vehicles = append(f.vehicles, v)
	return nil
}

func (f *Fleet) AddEquipment(e domain.Inventoriable) {
	f.equipment = append(f.equipment, e)
}

func (f *Fleet) Vehicles() []domain.Rentable {
	out := make([]domain.Rentable, len(f.vehicles))
	copy(out, f.vehicles)
	return out
}

func (f *Fleet) Equipment() []domain.Inventoriable {
	out := make([]domain.Inventoriable, len(f.equipment))
	copy(out, f.equipment)
	return out
}

func (f *Fleet) All() []domain.Inventoriable {
	all := make([]domain.Inventoriable, 0, len(f.vehicles)+len(f.equipment))
	for _, v := range f.vehicles {
		all = append(all, v)
	}
	all = append(all, f.equipment...)
	return all
}
