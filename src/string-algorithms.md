# Algorithms Workshop - String Algorithms

## String Definition in Computer Science

[Strings (Wikipedia Definition)](https://en.wikipedia.org/wiki/String_%28computer_science%29)

> In computer programming, a string is traditionally a sequence of characters, either as a literal constant or as some kind of variable. The latter may allow its elements to be mutated and the length changed, or it may be fixed (after creation). A string is generally considered as a data type and is often implemented as an array data structure of bytes (or words) that stores a sequence of elements, typically characters, using some character encoding. String may also denote more general arrays or other sequence (or list) data types and structures.

> Depending on the programming language and precise data type used, a variable declared to be a string may either cause storage in memory to be statically allocated for a predetermined maximum length or employ dynamic allocation to allow it to hold a variable number of elements.

## Types of String Algorithms

* String Searching
    * Brute Force String Search
    * Deterministic Finite Automaton (DFA) Search
    * [Knuth Morris Pratt Algorithms](https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm)

* String Sorting
    * [American Flag Sort](https://en.wikipedia.org/wiki/American_flag_sort)
    * [Burst Sort](https://en.wikipedia.org/wiki/Burstsort)
    * [Multi-key quicksort](https://en.wikipedia.org/wiki/Multi-key_quicksort)
    * [Radix sort](https://en.wikipedia.org/wiki/Radix_sort)

* String Matching
    * [Bitap Algorithm](https://en.wikipedia.org/wiki/Bitap_algorithm)

* String Collation
    * [Unicode Collation Algorithm](https://en.wikipedia.org/wiki/Unicode_collation_algorithm)

## Brute Force String Searching Algorithm

```go
package brutesearch

func search(text string, pattern string) int {
	var index int
	for i, ch := range text {
		if string(ch) == pattern {
			index = i
		}
	}
	return index
}
```

This is a naive search implementation

#### Rabin Karp Golang Source Code

```go
// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
	n := len(substr)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return IndexByte(s, substr[0])
	case n == len(s):
		if substr == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	case n <= bytealg.MaxLen:
		// Use brute force when s and substr both are small
		if len(s) <= bytealg.MaxBruteForce {
			return bytealg.IndexString(s, substr)
		}
		c0 := substr[0]
		c1 := substr[1]
		i := 0
		t := len(s) - n + 1
		fails := 0
		for i < t {
			if s[i] != c0 {
				// IndexByte is faster than bytealg.IndexString, so use it as long as
				// we're not getting lots of false positives.
				o := IndexByte(s[i:t], c0)
				if o < 0 {
					return -1
				}
				i += o
			}
			if s[i+1] == c1 && s[i:i+n] == substr {
				return i
			}
			fails++
			i++
			// Switch to bytealg.IndexString when IndexByte produces too many false positives.
			if fails > bytealg.Cutover(i) {
				r := bytealg.IndexString(s[i:], substr)
				if r >= 0 {
					return r + i
				}
				return -1
			}
		}
		return -1
	}
	c0 := substr[0]
	c1 := substr[1]
	i := 0
	t := len(s) - n + 1
	fails := 0
	for i < t {
		if s[i] != c0 {
			o := IndexByte(s[i:t], c0)
			if o < 0 {
				return -1
			}
			i += o
		}
		if s[i+1] == c1 && s[i:i+n] == substr {
			return i
		}
		i++
		fails++
		if fails >= 4+i>>4 && i < t {
			// See comment in ../bytes/bytes_generic.go.
			j := indexRabinKarp(s[i:], substr)
			if j < 0 {
				return -1
			}
			return i + j
		}
	}
	return -1
}

func indexRabinKarp(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}
// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
```

#### String Search Test Run 

[String Search Test](https://github.com/jbelmont/algorithms-workshop/blob/master/stringAlgorithms/brutesearch/search_test.go)