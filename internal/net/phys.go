package net

import (
	"fmt"
	"log"

	"github.com/Elfsilon/net_sim/internal/net/phys/ieee"
	"github.com/Elfsilon/net_sim/internal/net/phys/mf"
)

func Run() {
	IEEE := ieee.New()

	mf := mf.New()
	cisco := mf.RegisterManufacturer("Cisco")
	if err := cisco.GetMFG(IEEE); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cisco)

	decoded := IEEE.DecodeMFG(cisco.MFG())
	fmt.Println(decoded)
}

// type NICRegistry struct {
// 	macs []byte
// }

// func NewNICRegistry() *NICRegistry {
// 	return &NICRegistry{
// 		adrs: make([6]byte, 0),
// 	}
// }

// func (r *NICRegistry) NewNIC() *NIC {
// 	var mac byte
// 	if len(r.macs) == 0 {
// 		mac = 0
// 	} else {
// 		mac = r.macs[len(r.macs)-1] + 1
// 	}
// 	r.macs = append(r.macs, mac)

// 	return &NIC{
// 		mac: mac,
// 	}
// }

// // Network Interface Card/Controller (NIC)
// type NIC struct {
// 	mac byte
// }
