package homomorphine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBooleanCircuitBackend(t *testing.T) {
	value := 16

	// create and initialize backends
	backend := CreateBooleanCircuitHomomorphineBackend("tfhe")
	encryptBackend := CreateBooleanCircuitHomomorphineBackend("tfhe")

	backend.Init()
	encryptBackend.Init()

	// generate keys and encrypt the value
	backend.GenerateKeys()
	cipher := backend.Encrypt(value)

	// perform NOT operation on a public side
	encryptBackend.SetPublicKey(backend.GetPublicKey())
	resultCipher := encryptBackend.NOT(cipher)

	assert.Equal(t, backend.Decrypt(resultCipher), 65519,
		"Incorrect results of boolean operations: %d", backend.Decrypt(resultCipher))
}
