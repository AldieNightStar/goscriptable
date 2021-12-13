package goscriptable

import "strings"

// DivString - String Divider structure
// Used to append and then divide strings
type DivString struct {
	Strings []string
	Divider string
}

// Create new String Divider
func NewDivString(div string) *DivString {
	return &DivString{
		Strings: make([]string, 0, 8),
		Divider: div,
	}
}

// Add element to String Divider
func (d *DivString) Append(s string) {
	d.Strings = append(d.Strings, s)
}

// Divide strings by div string
func (d *DivString) Divide(div string) string {
	return strings.Join(d.Strings, div)
}

// Set New Divideer string
func (d *DivString) SetDivider(div string) {
	d.Divider = div
}

// Divide strings by dicvider string, specified by constructor
func (d *DivString) String() string {
	return d.Divide(d.Divider)
}
