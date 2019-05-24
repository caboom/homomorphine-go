package homomorphine

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lhomomorphine
// #include <stdio.h>
// #include <stdlib.h>
// #include <homomorphine/clang_backend_interface.hpp>
import (
	"C"
)
import "unsafe"

type Homomorphine struct {
	backend C.BackendWrapper
}

func CreateHomomorphineBackend(name string) Homomorphine {
	var ret Homomorphine
	ret.backend = C.CreateBackend(C.CString(name))
	return ret
}

func (b Homomorphine) Free() {
	C.FreeBackend(b.backend)
}

func (b Homomorphine) Init() {
	C.InitBackend(b.backend)
}

func (b Homomorphine) SetAlgorithm(algorithm string) {
	csAlgorithm := C.CString(algorithm)
	defer C.free(unsafe.Pointer(csAlgorithm))

	C.SetBackendAlgorithm(b.backend, csAlgorithm)
}

func (b Homomorphine) GetParam(key string) string {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	csParam := C.GetBackendParam(b.backend, csKey)
	defer C.free(unsafe.Pointer(csParam))
	return C.GoString(csParam)
}

func (b Homomorphine) SetParam(key string, value string) {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	C.SetBackendParam(b.backend, csKey, csValue)
}

func (b Homomorphine) GenerateKeys() {
	C.GenerateBackendKeys(b.backend)
}

func (b Homomorphine) GetPublicKey() string {
	csPublicKey := C.GetBackendPublicKey(b.backend)
	defer C.free(unsafe.Pointer(csPublicKey))

	return C.GoString(csPublicKey)
}

func (b Homomorphine) GetSecretKey() string {
	csSecretKey := C.GetBackendSecretKey(b.backend)
	defer C.free(unsafe.Pointer(csSecretKey))

	return C.GoString(csSecretKey)
}

func (b Homomorphine) SetPublicKey(publicKey string) {
	csPublicKey := C.CString(publicKey)
	defer C.free(unsafe.Pointer(csPublicKey))

	C.SetBackendPublicKey(b.backend, csPublicKey)
}

func (b Homomorphine) SetSecretKey(secretKey string) {
	csSecretKey := C.CString(secretKey)
	defer C.free(unsafe.Pointer(csSecretKey))

	C.SetBackendSecretKey(b.backend, csSecretKey)
}

func (b Homomorphine) GetCipher() string {
	csCipher := C.GetBackendCipher(b.backend)
	defer C.free(unsafe.Pointer(csCipher))

	return C.GoString(csCipher)
}

func (b Homomorphine) SetCipher(cipher string) {
	csCipher := C.CString(cipher)
	defer C.free(unsafe.Pointer(csCipher))

	C.SetBackendCipher(b.backend, csCipher)
}

func (b Homomorphine) Encrypt(value int) string {
	csCipher := C.BackendEncrypt(b.backend, C.long(value))
	defer C.free(unsafe.Pointer(csCipher))

	return C.GoString(csCipher)
}

func (b Homomorphine) Decrypt() int {
	return int(C.BackendDecrypt(b.backend))
}

func (b Homomorphine) Add(value int) {
	C.BackendAdd(b.backend, C.long(value))
}

func (b Homomorphine) Negate() {
	C.BackendNegate(b.backend)
}

func (b Homomorphine) Multiply(value int) {
	C.BackendMultiply(b.backend, C.long(value))
}
