package roper

import "reflect"

// IsDefaultValue returns whether the given interface is the default for its kind
// i.e. IsDefaultValue(1) -> false
// IsDefaultValue("") -> true
func IsDefaultValue(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
