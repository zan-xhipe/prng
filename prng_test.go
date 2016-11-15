package prng_test

import (
	"math/rand"
	"testing"

	"github.com/zan-xhipe/prng"
)

func TestSaveRestore(t *testing.T) {
	s0 := prng.NewSource(7531902468)
	p1 := rand.New(s0)

	for i := 0; i < 1000; i++ {
		p1.Int63()
	}

	p2 := rand.New(prng.NewSource(s0.State))

	v1 := p1.Int63()
	v2 := p2.Int63()

	if v1 != v2 {
		t.Error("failed to capture state properly")
	}
}
