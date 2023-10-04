package goslices

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// type complex interface {
// 	~complex64 | ~complex128
// }

type float interface {
	~float32 | ~float64
}

type number interface {
	float | signed | unsigned
}

type integer interface {
	signed | unsigned
}

type ordered interface {
	integer | float | ~string
}
