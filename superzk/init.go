package superzk

import "C"
import (
	"github.com/axis-cash/go-axis-import/c_superzk"
	"github.com/axis-cash/go-axis-import/c_type"
)

func ZeroInit(account_dir string, netType c_type.NetType) error {
	c_superzk.InitParams()
	return nil
}

func ZeroInit_NoCircuit() error {
	c_superzk.InitParams_NoCircuit()
	return nil
}

func ZeroInit_OnlyInOuts() error {
	c_superzk.InitParams()
	return nil
}
