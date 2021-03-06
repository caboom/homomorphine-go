package homomorphine

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lhomomorphine
// #include <stdio.h>
// #include <stdlib.h>
// #include <homomorphine/clang_types.hpp>
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

func (b ArithmeticBackend) GetPublicKey() blob {
	var publicKey blob
	csPublicKey := C.GetArithmeticBackendPublicKey(b.backend)
	defer C.free(unsafe.Pointer(csPublicKey.content))

	publicKey.size = int64(csPublicKey.size)
	publicKey.content = C.GoBytes(unsafe.Pointer(csPublicKey.content), C.int(int(csPublicKey.size)))

	return publicKey
}

func (b ArithmeticBackend) GetSecretKey() blob {
	var secretKey blob
	csSecretKey := C.GetArithmeticBackendSecretKey(b.backend)
	defer C.free(unsafe.Pointer(csSecretKey.content))

	secretKey.size = int64(csSecretKey.size)
	secretKey.content = C.GoBytes(unsafe.Pointer(csSecretKey.content), C.int(int(csSecretKey.size)))

	return secretKey
}

func (b ArithmeticBackend) SetPublicKey(publicKey blob) {
	C.SetArithmeticBackendPublicKey(b.backend, *((*C.struct_blob_t)(unsafe.Pointer(&publicKey))))
}

func (b ArithmeticBackend) SetSecretKey(secretKey blob) {
	C.SetArithmeticBackendPublicKey(b.backend, *((*C.struct_blob_t)(unsafe.Pointer(&secretKey))))
}

func (b ArithmeticBackend) GetCipher() blob {
	var cipher blob
	csCipher := C.GetArithmeticBackendCipher(b.backend)
	defer C.free(unsafe.Pointer(csCipher.content))

	cipher.size = int64(csCipher.size)
	cipher.content = C.GoBytes(unsafe.Pointer(csCipher.content), C.int(int(csCipher.size)))

	return cipher
}

func (b ArithmeticBackend) SetCipher(cipher blob) {
	C.SetArithmeticBackendCipher(b.backend, *((*C.struct_blob_t)(unsafe.Pointer(&cipher))))
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
