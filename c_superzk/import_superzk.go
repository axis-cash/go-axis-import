package c_superzk

import _ "github.com/axis-cash/go-axis-import/czero"

/*

#cgo CFLAGS: -I ../czero/szk_include

#cgo LDFLAGS: -L ../czero/lib -lsuperzk

*/
import "C"

func Is_czero_debug() bool {
	return false
}
