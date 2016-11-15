// Package prng provides a pseudorandom number generator that is an
// implementation of xorshift*
// It implements the math/rand Source interface.
// The main reason to use this generator is that it has a small state
// that can be saved and restored to continue the sequence from where
// it left off.  This is specifically for small games where you want
// deterministic random numbers between from a save file.
package prng

type Source struct {
	State int64
}

func (s *Source) Seed(seed int64) {
	s.State = seed
}

func (s *Source) Int63() int64 {
	s.State ^= s.State >> 12 // a
	s.State ^= s.State << 25 // b
	s.State ^= s.State >> 27 // c
	return s.State * 2685821657736338717
}

func NewSource(seed int64) *Source {
	return &Source{State: seed}
}
