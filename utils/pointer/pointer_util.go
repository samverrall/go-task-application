package pointerutil

// ToPointer takes any value and returns a pointer of said value.
func ToPointer[T any](value T) *T {
	return &value
}

// PointerVal returns the deferenced pointer of v, a pointer. If the pointer
// is nil the zerovalue is returned.
func PointerVal[T any](v *T) T {
	if v == nil {
		var zeroVal T
		return zeroVal
	}
	return *v
}
