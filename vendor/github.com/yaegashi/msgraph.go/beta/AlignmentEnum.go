// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Alignment undocumented
type Alignment int

const (
	// AlignmentVLeft undocumented
	AlignmentVLeft Alignment = 0
	// AlignmentVRight undocumented
	AlignmentVRight Alignment = 1
	// AlignmentVCenter undocumented
	AlignmentVCenter Alignment = 2
)

// AlignmentPLeft returns a pointer to AlignmentVLeft
func AlignmentPLeft() *Alignment {
	v := AlignmentVLeft
	return &v
}

// AlignmentPRight returns a pointer to AlignmentVRight
func AlignmentPRight() *Alignment {
	v := AlignmentVRight
	return &v
}

// AlignmentPCenter returns a pointer to AlignmentVCenter
func AlignmentPCenter() *Alignment {
	v := AlignmentVCenter
	return &v
}