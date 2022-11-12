package goscriptable

import "strings"

// Splits strings and removes empty spaces between
//  Returns array of splitted strings
func SplitStringBy(s string, div string) []string {
	arr := strings.Split(s, div)
	// Remove empty string elements
	// And trim the spaces
	{
		narr := make([]string, 0, 8)
		for _, ss := range arr {
			if len(ss) < 1 || ss == " " {
				continue
			}
			narr = append(narr, strings.Trim(ss, " \t"))
		}
		arr = narr
	}
	return arr
}
