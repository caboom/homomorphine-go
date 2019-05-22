package homomorphine

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lhomomorphine
// #include <stdio.h>
// #include <stdlib.h>
// #include <homomorphine/clang_backend_interface.hpp>
import "C"

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
	C.SetBackendAlgorithm(b.backend, C.CString(algorithm))
}

func (b Homomorphine) GetParam(key string) string {
	return C.GoString(C.GetBackendParam(b.backend, C.CString(key)))
}

func (b Homomorphine) SetParam(key string, value string) {
	C.SetBackendParam(b.backend, C.CString(key), C.CString(value))
}

func (b Homomorphine) GenerateKeys() {
	C.GenerateBackendKeys(b.backend)
}

func (b Homomorphine) GetPublicKey() string {
	return C.GoString(C.GetBackendPublicKey(b.backend))
}

func (b Homomorphine) GetSecretKey() string {
	return C.GoString(C.GetBackendSecretKey(b.backend))
}

func (b Homomorphine) SetPublicKey(publicKey string) {
	C.SetBackendPublicKey(b.backend, C.CString(publicKey))
}

func (b Homomorphine) SetSecretKey(secretKey string) {
	C.SetBackendSecretKey(b.backend, C.CString(secretKey))
}

func (b Homomorphine) GetCipher() string {
	return C.GoString(C.GetBackendCipher(b.backend))
}

func (b Homomorphine) SetCipher(cipher string) {
	C.SetBackendCipher(b.backend, C.CString(cipher))
}

func (b Homomorphine) Encrypt(value int) string {
	return C.GoString(C.BackendEncrypt(b.backend, C.long(value)))
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
