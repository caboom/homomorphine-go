package homomorphine

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lhomomorphine
// #include <stdio.h>
// #include <stdlib.h>
// #include <homomorphine/clang_types.hpp>
// #include <homomorphine/clang_boolean_circuit_backend_interface.hpp>
import (
	"C"
)
import "unsafe"

type BooleanCircuitBackend struct {
	backend C.BooleanCircuitBackendWrapper
}

func CreateBooleanCircuitHomomorphineBackend(name string) BooleanCircuitBackend {
	var ret BooleanCircuitBackend
	ret.backend = C.CreateBooleanCircuitBackend(C.CString(name))
	return ret
}

func (b BooleanCircuitBackend) Free() {
	C.FreeBooleanCircuitBackend(b.backend)
}

func (b BooleanCircuitBackend) Init() {
	C.InitBooleanCircuitBackend(b.backend)
}

func (b BooleanCircuitBackend) SetAlgorithm(algorithm string) {
	csAlgorithm := C.CString(algorithm)
	defer C.free(unsafe.Pointer(csAlgorithm))

	C.SetBooleanCircuitBackendAlgorithm(b.backend, csAlgorithm)
}

func (b BooleanCircuitBackend) GetParam(key string) string {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	csParam := C.GetBooleanCircuitBackendParam(b.backend, csKey)
	defer C.free(unsafe.Pointer(csParam))
	return C.GoString(csParam)
}

func (b BooleanCircuitBackend) SetParam(key string, value string) {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	C.SetBooleanCircuitBackendParam(b.backend, csKey, csValue)
}
