// Code generated by "stringer -type=advanceCode"; DO NOT EDIT.

package sql

import "fmt"

const _advanceCode_name = "advanceUnknownstayInPlaceadvanceOneskipQueryStrrewind"

var _advanceCode_index = [...]uint8{0, 14, 25, 35, 47, 53}

func (i advanceCode) String() string {
	if i < 0 || i >= advanceCode(len(_advanceCode_index)-1) {
		return fmt.Sprintf("advanceCode(%d)", i)
	}
	return _advanceCode_name[_advanceCode_index[i]:_advanceCode_index[i+1]]
}