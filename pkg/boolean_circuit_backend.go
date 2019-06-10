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

func (b BooleanCircuitBackend) GenerateKeys() {
	C.GenerateBooleanCircuitBackendKeys(b.backend)
}

func (b BooleanCircuitBackend) GetPublicKey() blob {
	var publicKey blob
	csPublicKey := C.GetBooleanCircuitBackendPublicKey(b.backend)
	defer C.free(unsafe.Pointer(csPublicKey.content))

	publicKey.size = int64(csPublicKey.size)
	publicKey.content = C.GoBytes(unsafe.Pointer(csPublicKey.content), C.int(int(csPublicKey.size)))

	return publicKey
}

func (b BooleanCircuitBackend) GetSecretKey() blob {
	var secretKey blob
	csSecretKey := C.GetBooleanCircuitBackendSecretKey(b.backend)
	defer C.free(unsafe.Pointer(csSecretKey.content))

	secretKey.size = int64(csSecretKey.size)
	secretKey.content = C.GoBytes(unsafe.Pointer(csSecretKey.content), C.int(int(csSecretKey.size)))

	return secretKey
}

func (b BooleanCircuitBackend) SetPublicKey(publicKey blob) {
	C.SetBooleanCircuitBackendPublicKey(b.backend, *((*C.struct_blob_t)(unsafe.Pointer(&publicKey))))
}

func (b BooleanCircuitBackend) SetSecretKey(secretKey blob) {
	C.SetBooleanCircuitBackendPublicKey(b.backend, *((*C.struct_blob_t)(unsafe.Pointer(&secretKey))))
}

func (b BooleanCircuitBackend) Encrypt(value int) blob {
	var cipher blob
	csCipher := C.BooleanCircuitEncrypt(b.backend, C.int(value))
	defer C.free(unsafe.Pointer(csCipher.content))

	cipher.size = int64(csCipher.size)
	cipher.content = C.GoBytes(unsafe.Pointer(csCipher.content), C.int(int(csCipher.size)))

	return cipher
}

func (b BooleanCircuitBackend) Encode(value int) blob {
	var cipher blob
	csCipher := C.BooleanCircuitEncode(b.backend, C.int(value))
	defer C.free(unsafe.Pointer(csCipher.content))

	cipher.size = int64(csCipher.size)
	cipher.content = C.GoBytes(unsafe.Pointer(csCipher.content), C.int(int(csCipher.size)))

	return cipher
}

func (b BooleanCircuitBackend) Decrypt(cipher blob) int {
	return int(C.BooleanCircuitDecrypt(b.backend, *((*C.struct_blob_t)(unsafe.Pointer(&cipher)))))
}

func (b BooleanCircuitBackend) NOT(cipher blob) blob {
	var result blob
	csResult := C.BooleanCircuitNOT(b.backend, *((*C.struct_blob_t)(unsafe.Pointer(&cipher))))
	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}

func (b BooleanCircuitBackend) COPY(cipher blob) blob {
	var result blob
	csResult := C.BooleanCircuitCOPY(b.backend, *((*C.struct_blob_t)(unsafe.Pointer(&cipher))))
	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}

func (b BooleanCircuitBackend) NAND(cipher_x blob, cipher_y blob) blob {
	var result blob
	csResult := C.BooleanCircuitNAND(b.backend,
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_x))),
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_y))))

	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}

func (b BooleanCircuitBackend) OR(cipher_x blob, cipher_y blob) blob {
	var result blob
	csResult := C.BooleanCircuitOR(b.backend,
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_x))),
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_y))))

	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}

func (b BooleanCircuitBackend) AND(cipher_x blob, cipher_y blob) blob {
	var result blob
	csResult := C.BooleanCircuitAND(b.backend,
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_x))),
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_y))))

	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}

func (b BooleanCircuitBackend) XOR(cipher_x blob, cipher_y blob) blob {
	var result blob
	csResult := C.BooleanCircuitXOR(b.backend,
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_x))),
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_y))))

	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}

func (b BooleanCircuitBackend) XNOR(cipher_x blob, cipher_y blob) blob {
	var result blob
	csResult := C.BooleanCircuitXNOR(b.backend,
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_x))),
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_y))))

	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}

func (b BooleanCircuitBackend) NOR(cipher_x blob, cipher_y blob) blob {
	var result blob
	csResult := C.BooleanCircuitNOR(b.backend,
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_x))),
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_y))))

	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}

func (b BooleanCircuitBackend) MUX(cipher_x blob, cipher_y blob, cipher_z blob) blob {
	var result blob
	csResult := C.BooleanCircuitMUX(b.backend,
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_x))),
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_y))),
		*((*C.struct_blob_t)(unsafe.Pointer(&cipher_z))))

	defer C.free(unsafe.Pointer(csResult.content))

	result.size = int64(csResult.size)
	result.content = C.GoBytes(unsafe.Pointer(csResult.content), C.int(int(csResult.size)))

	return result
}
