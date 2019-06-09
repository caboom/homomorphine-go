package homomorphine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmeticBackend(t *testing.T) {
	backend := CreateArithmeticHomomorphineBackend("seal")
	backend.SetAlgorithm("bfv")
	backend.Init()

	backendEncrypt := CreateArithmeticHomomorphineBackend("seal")
	backendEncrypt.SetAlgorithm("bfv")
	backendEncrypt.Init()

	backend.GenerateKeys()
	backend.Encrypt(10000)

	backendEncrypt.SetPublicKey(backend.GetPublicKey())

	backendEncrypt.SetCipher(backend.GetCipher())
	backendEncrypt.Add(10)
	backendEncrypt.Multiply(20)

	backend.SetCipher(backendEncrypt.GetCipher())

	assert.Equal(t, backend.Decrypt(), 200200, "Incorrect results of numerical operations: %d", backend.Decrypt())
}
