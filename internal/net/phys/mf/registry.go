package mf

import "github.com/Elfsilon/net_sim/internal/net/phys/shared"

type ManufacturerRegistry struct {
	last int64
	reg  map[int64]shared.ManufacturerInfo
}

func New() *ManufacturerRegistry {
	return &ManufacturerRegistry{
		reg: make(map[int64]shared.ManufacturerInfo, 0),
	}
}

func (r *ManufacturerRegistry) RegisterManufacturer(name string) *Manufacturer {
	r.last++
	currentID := r.last

	info := shared.ManufacturerInfo{
		ID:   currentID,
		Name: name,
	}
	r.reg[currentID] = info

	return &Manufacturer{info: info}
}
