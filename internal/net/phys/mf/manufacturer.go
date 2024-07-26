package mf

import (
	"fmt"

	"github.com/Elfsilon/net_sim/internal/net/phys/shared"
)

type Manufacturer struct {
	info shared.ManufacturerInfo
	mfg  shared.MFGCode
}

func (m *Manufacturer) String() string {
	template := "Manufacturer(#%v, '%v', %v)"
	return fmt.Sprintf(template, m.info.ID, m.info.Name, m.mfg)
}

func (m *Manufacturer) MFG() shared.MFGCode {
	return m.mfg
}

func (m *Manufacturer) GetMFG(provider shared.MFGProvider) error {
	code, err := provider.GetMFG(m.info)
	if err != nil {
		return err
	}
	m.mfg = code
	return nil
}
