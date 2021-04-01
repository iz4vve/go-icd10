package icd10

import (
	"testing"
)

func TestSetup(t *testing.T) {

}

func TestSetupPCS(t *testing.T) {
	_, err := setupPCS()

	if err != nil {
		t.Error(err)
	}
}

func BenchmarkSetup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = setup()
	}
}

func BenchmarkSetupPCS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = setupPCS()
	}
}
