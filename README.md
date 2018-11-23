# Hash similarity

This program aims to help to determinates with how many hashed generated keys we can be near/similar of a specific hash, using string distance algorithms.

Algorithms available:

- Leveinshtein (lev)
- Hamming (ham)
- Simhash (sim)
- Jaro Winkler (jar)
- Cosine (cos)
- Longest Common Subsequence (lcs)

## Running

```go
go run main.go -t {RATIO_SIMILARITY_THRESHOLD} -{ALGO}
```