package homomorphine

import (
	"testing"
)

func TestHomomorphine(t *testing.T) {
	homomorphine := CreateHomomorphineBackend("seal")
	homomorphine.SetAlgorithm("bfv")
	homomorphine.Init()

	homomorphineEncrypt := CreateHomomorphineBackend("seal")
	homomorphineEncrypt.SetAlgorithm("bfv")
	homomorphineEncrypt.Init()

	homomorphine.GenerateKeys()
	homomorphine.Encrypt(10000)

	homomorphineEncrypt.SetPublicKey(homomorphine.GetPublicKey())
	homomorphineEncrypt.SetCipher(homomorphine.GetCipher())
	homomorphineEncrypt.Add(10)
	homomorphineEncrypt.Multiply(20)

	homomorphine.SetCipher(homomorphineEncrypt.GetCipher())

	if homomorphine.Decrypt() != 200200 {
		t.Errorf("Incorrect results of numerical operations: %d", homomorphine.Decrypt())
	}
}
