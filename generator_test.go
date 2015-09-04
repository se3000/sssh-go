package generator

import "testing"

func TestInitialization(t *testing.T) {
	shadowCount := 123
	secret := GenerateSecret(shadowCount)
	if shadowCount != len(secret.Shadows) {
		t.Error("Expected", len(secret.Shadows), "to equal", shadowCount)
	}
}
