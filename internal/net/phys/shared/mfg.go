package shared

import (
	"encoding/hex"
	"fmt"
)

type MFGCode [3]byte

func (mfg MFGCode) String() string {
	return fmt.Sprintf("MFG(0x%v)", hex.EncodeToString(mfg[:]))
}

type MFGByteInfo struct {
	Byte        byte
	Value       string
	Description string
}

func (i MFGByteInfo) String() string {
	return fmt.Sprintf("Oct(%v, '%v', '%v')", i.Byte, i.Value, i.Description)
}

type MFGInfo struct {
	MFG        MFGCode
	Cast       MFGByteInfo
	Visibility MFGByteInfo
}

func (i MFGInfo) String() string {
	return fmt.Sprintf("MFGInfo(%v, %v, %v)", i.MFG, i.Cast, i.Visibility)
}

type MFGProvider interface {
	GetMFG(info ManufacturerInfo) (MFGCode, error)
}
