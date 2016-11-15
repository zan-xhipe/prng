// Package prng provides a pseudo-random number generator that is an
// implementation of xor shift*
// It implements the math/rand Source interface.
// The main reason to use this generator is that it has a small state
// that can be saved and restored to continue the sequence from where
// it left off.  This is specifically for small games where you want
// deterministic random numbers between from a save file.
// The seed should not be set to 0.
package prng

type Source struct {
	State int64
}

// Seed uses the provided seed value to initialize the Source state.
// You should never give this source a seed of 0 as it will be unable
// to generate values from a zero state.
func (s *Source) Seed(seed int64) {
	s.State = seed
}

// Int63 returns a non-negative pseudo-random number from the prng Source.
func (s *Source) Int63() int64 {
	s.State ^= s.State >> 12 // a
	s.State ^= s.State << 25 // b
	s.State ^= s.State >> 27 // c
	r := s.State * 2685821657736338717

	// get absolute value
	temp := r >> 64
	r ^= temp
	r += temp & 1
	return r
}

// NewSource returns a new pseudo-random Source seeded with the given value.
func NewSource(seed int64) *Source {
	return &Source{State: seed}
}
