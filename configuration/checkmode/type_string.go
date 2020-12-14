// Code generated by "stringer -type=Type -linecomment"; DO NOT EDIT.

package checkmode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Strict-0]
	_ = x[Specification-1]
	_ = x[Permissive-2]
	_ = x[LibraryManagerSubmission-3]
	_ = x[LibraryManagerIndexed-4]
	_ = x[Official-5]
	_ = x[Default-6]
}

const _Type_name = "strictspecificationpermissivesubmitupdateARDUINO_LINT_OFFICIALdefault"

var _Type_index = [...]uint8{0, 6, 19, 29, 35, 41, 62, 69}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
