package utils

import "fmt"

func AnySliceToString(s []any) []string {
	var r []string
	for _, v := range s {
		r = append(r, fmt.Sprintf("%s", v))
	}
	return r
}
