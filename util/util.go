package util

import "time"

//----------------------------------------------------------------------------------
// dur()
//----------------------------------------------------------------------------------
// Parse Duration while ignoring the error to allow inlining
//----------------------------------------------------------------------------------
func Dur(t string) time.Duration {
	d, _ := time.ParseDuration(t)
	return d
}

