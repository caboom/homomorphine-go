package homomorphine

import (
	"testing"
)

func TestBooleanCircuitBackend(t *testing.T) {
	backend := CreateBooleanCircuitHomomorphineBackend("tfhe")
	backend.Init()
}
