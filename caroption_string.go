// Code generated by "stringer -type=CarOption"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GPS-1]
	_ = x[AWD-2]
	_ = x[SunRoof-4]
	_ = x[HeatedSeat-8]
	_ = x[DriveAssist-16]
}

const (
	_CarOption_name_0 = "GPSAWD"
	_CarOption_name_1 = "SunRoof"
	_CarOption_name_2 = "HeatedSeat"
	_CarOption_name_3 = "DriveAssist"
)

var (
	_CarOption_index_0 = [...]uint8{0, 3, 6}
)

func (i CarOption) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _CarOption_name_0[_CarOption_index_0[i]:_CarOption_index_0[i+1]]
	case i == 4:
		return _CarOption_name_1
	case i == 8:
		return _CarOption_name_2
	case i == 16:
		return _CarOption_name_3
	default:
		return "CarOption(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
