package ieee

import (
	"errors"

	"github.com/Elfsilon/net_sim/internal/net/phys/shared"
)

type CastByteMask byte
type VisibilityByteMask byte

const (
	UnicastMask   CastByteMask = 0x00
	MilticastMask CastByteMask = 0x01

	GloballyUniqueMask       VisibilityByteMask = 0x00 << 1
	LocallyAdministratedMask VisibilityByteMask = 0x01 << 1
)

type IEEE struct {
	lfb, lsb, ltb byte
	reg           map[shared.MFGCode]shared.ManufacturerInfo
}

func New() *IEEE {
	return &IEEE{
		reg: make(map[shared.MFGCode]shared.ManufacturerInfo, 0),
	}
}

// MFG - Manufacturing code
func (i *IEEE) GetMFG(info shared.ManufacturerInfo) (shared.MFGCode, error) {
	i.ltb = +0x01
	if i.ltb == 0 {
		i.lsb += 0x01
		if i.lsb == 0 {
			i.lfb += 0x01
			if i.lfb == 0 {
				return [3]byte{}, errors.New("maximum of released MFG codes reached")
			}
		}
	}
	fb, sb, tb := i.maskFirstByte(i.lfb, UnicastMask, GloballyUniqueMask), i.lsb, i.ltb
	return [3]byte{fb, sb, tb}, nil
}

func (i *IEEE) DecodeMFG(code shared.MFGCode) shared.MFGInfo {
	fb := code[0]

	cast := fb & 0x01
	castInfo := shared.MFGByteInfo{Byte: cast}
	if cast == 0x01 {
		castInfo.Value = "multicast"
		castInfo.Description = ""
	} else {
		castInfo.Value = "unicast"
		castInfo.Description = ""
	}

	vis := fb >> 1 & 0x01
	visInfo := shared.MFGByteInfo{Byte: vis}
	if vis == 0x01 {
		visInfo.Value = "locally administrated"
		visInfo.Description = ""
	} else {
		visInfo.Value = "globally unique (OUI enforced)"
		visInfo.Description = ""
	}

	return shared.MFGInfo{
		Cast:       castInfo,
		Visibility: visInfo,
	}
}

// b - first 'raw' byte of the MFG, raw means that it is a counter before shifting and applying bit masks
func (i *IEEE) maskFirstByte(b byte, cast CastByteMask, visibility VisibilityByteMask) byte {
	return b<<2 | byte(cast) | byte(visibility)
}
