A go psuedo-random number generator for my small games.

# Why?
The built-in go psuedo-random number genrator in math/rand does not provide easy access to the rng state.  This makes it difficult to save and restore that state in game save files and have deterministic saved games.

# Usage

```
// Create source and prng
s := prng.NewSource(time.Now().Unix())
r := rand.New(s)

// Save state
state := s.State

// Restore state
r = rand.New(prng.NewSource(state)
```
