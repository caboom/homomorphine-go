package homomorphine

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lhomomorphine
// #include <stdio.h>
// #include <stdlib.h>
// #include <homomorphine/clang_arithmetic_backend_interface.hpp>
import (
	"C"
)
import "unsafe"

type ArithmeticBackend struct {
	backend C.ArithmeticBackendWrapper
}

func CreateArithmeticHomomorphineBackend(name string) ArithmeticBackend {
	var ret ArithmeticBackend
	ret.backend = C.CreateArithmeticBackend(C.CString(name))
	return ret
}

func (b ArithmeticBackend) Free() {
	C.FreeArithmeticBackend(b.backend)
}

func (b ArithmeticBackend) Init() {
	C.InitArithmeticBackend(b.backend)
}

func (b ArithmeticBackend) SetAlgorithm(algorithm string) {
	csAlgorithm := C.CString(algorithm)
	defer C.free(unsafe.Pointer(csAlgorithm))

	C.SetArithmeticBackendAlgorithm(b.backend, csAlgorithm)
}

func (b ArithmeticBackend) GetParam(key string) string {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	csParam := C.GetArithmeticBackendParam(b.backend, csKey)
	defer C.free(unsafe.Pointer(csParam))
	return C.GoString(csParam)
}

func (b ArithmeticBackend) SetParam(key string, value string) {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	C.SetArithmeticBackendParam(b.backend, csKey, csValue)
}

func (b ArithmeticBackend) GenerateKeys() {
	C.GenerateArithmeticBackendKeys(b.backend)
}

func (b ArithmeticBackend) GetPublicKey() string {
	csPublicKey := C.GetArithmeticBackendPublicKey(b.backend)
	defer C.free(unsafe.Pointer(csPublicKey))

	return C.GoString(csPublicKey)
}

func (b ArithmeticBackend) GetSecretKey() string {
	csSecretKey := C.GetArithmeticBackendSecretKey(b.backend)
	defer C.free(unsafe.Pointer(csSecretKey))

	return C.GoString(csSecretKey)
}

func (b ArithmeticBackend) SetPublicKey(publicKey string) {
	csPublicKey := C.CString(publicKey)
	defer C.free(unsafe.Pointer(csPublicKey))

	C.SetArithmeticBackendPublicKey(b.backend, csPublicKey)
}

func (b ArithmeticBackend) SetSecretKey(secretKey string) {
	csSecretKey := C.CString(secretKey)
	defer C.free(unsafe.Pointer(csSecretKey))

	C.SetArithmeticBackendSecretKey(b.backend, csSecretKey)
}

func (b ArithmeticBackend) GetCipher() string {
	csCipher := C.GetArithmeticBackendCipher(b.backend)
	defer C.free(unsafe.Pointer(csCipher))

	return C.GoString(csCipher)
}

func (b ArithmeticBackend) SetCipher(cipher string) {
	csCipher := C.CString(cipher)
	defer C.free(unsafe.Pointer(csCipher))

	C.SetArithmeticBackendCipher(b.backend, csCipher)
}

func (b ArithmeticBackend) Encrypt(value int) {
	C.ArithmeticBackendEncrypt(b.backend, C.long(value))
}

func (b ArithmeticBackend) Decrypt() int {
	return int(C.ArithmeticBackendDecrypt(b.backend))
}

func (b ArithmeticBackend) Add(value int) {
	C.ArithmeticBackendAdd(b.backend, C.long(value))
}

func (b ArithmeticBackend) Negate() {
	C.ArithmeticBackendNegate(b.backend)
}

func (b ArithmeticBackend) Multiply(value int) {
	C.ArithmeticBackendMultiply(b.backend, C.long(value))
}
