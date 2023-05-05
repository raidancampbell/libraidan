package runsafe

import (
	"reflect"
	"unsafe"
)

// GetBytes returns the underlying byte slice of the given string
// mutations to the returned byte slice will affect the string, which breaks string immutability
// source: https://groups.google.com/g/golang-nuts/c/Zsfk-VMd_fU/m/O1ru4fO-BgAJ
func GetBytes(s string) []byte {
	if s == "" {
		return []byte{}
	}
	return (*[0x7fff0000]byte)(unsafe.Pointer(
		(*reflect.StringHeader)(unsafe.Pointer(&s)).Data),
	)[:len(s):len(s)]
}
