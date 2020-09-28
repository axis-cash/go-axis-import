package superzk

import "C"
import (
	"fmt"

	"github.com/axis-cash/go-axis-import/c_superzk"

	"github.com/axis-cash/go-axis-import/c_type"
	"github.com/axis-cash/go-axis-import/axisparam"
)

func Pk2PKrAndLICr(addr *c_type.Uint512, height uint64) (pkr c_type.PKr, licr c_type.LICr, ret bool) {
	if height >= axisparam.XIP0() {
		if IsPKValid(addr) {
			pkr = Pk2PKr(addr, c_superzk.RandomFr().NewRef())
			ret = true
		}
		return
	} else {
		fmt.Println("Pk2PKrAndLICr error: czero not support after XIP0")
		ret = false
		return
	}
}

func CheckLICr(pkr *c_type.PKr, licr *c_type.LICr, height uint64) bool {
	if height >= axisparam.XIP0() {
		return IsPKrValid(pkr)
	} else {
		fmt.Println("CheckLICr error: czero not support after XIP0")
		return false
	}
}
