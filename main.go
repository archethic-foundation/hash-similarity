package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	crand "crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"flag"
	"io"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/AllenDang/simhash"

	"github.com/dexyk/stringosim"
)

func main() {

	t := flag.Float64("t", 0.5, "Threshold similarity to reach before stoping the address generation")
	lev := flag.Bool("lev", false, "Leveinshtein algorithm")
	ham := flag.Bool("ham", false, "Hamming algorithm")
	sim := flag.Bool("sim", false, "Simhash algorithm")
	jaro := flag.Bool("jar", false, "Jaro Winkler algorithm")
	cos := flag.Bool("cos", false, "Cosine algorithm")
	lcs := flag.Bool("lcs", false, "Longest Common Subsequence algorithm")
	flag.Parse()

	if *lev == false && *ham == false && *sim == false && *jaro == false && *cos == false && *lcs == false {
		flag.Usage()
		return
	}

	rand.Seed(time.Now().UnixNano())

	dataHash := "843a54527fdcedfabc0e8f0234e76f63c1d149d26790531998c43d73e455e0da"

	var s float64

	cycles := 0
	var key string
	for s < *t {
		key = randomKeyHash()
		if *lev {
			dist := stringosim.Levenshtein([]rune(dataHash), []rune(key), stringosim.LevenshteinSimilarityOptions{
				SubstituteCost:  2,
				DeleteCost:      1,
				InsertCost:      1,
				CaseInsensitive: false,
			})
			s = math.Abs(1 - float64(dist)/64)
			log.Print(s)
		}

		if *sim {
			s = simhash.GetLikenessValue(dataHash, key)
			log.Print(s)
		}

		if *ham {
			dist, _ := stringosim.Hamming([]rune(dataHash), []rune(key))
			s = math.Abs(1 - float64(dist)/64)
			log.Print(s)
		}

		if *jaro {
			s = stringosim.JaroWinkler([]rune(dataHash), []rune(key))
			log.Print(s)
		}

		if *cos {
			s = stringosim.Cosine([]rune(dataHash), []rune(key))
			log.Print(s)
		}

		if *lcs {
			dist := stringosim.LCS([]rune(dataHash), []rune(key))
			s = 1 - (float64(dist) / 64)
			log.Print(s)
		}

		cycles++
	}

	log.Printf("data hash to compared: %s\n", dataHash)
	log.Printf("hash of the generated key similar: %s\n", key)
	log.Printf("%2.f percent similarity after %d keys generated", *t*100, cycles)
}

func randomKeyHash() string {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), crand.Reader)
	pubKey, _ := x509.MarshalPKIXPublicKey(key.Public())

	hash := sha256.New()
	hash.Write(pubKey)
	return hex.EncodeToString(hash.Sum(nil))
}

func randomHash() string {
	hash := sha256.New

	r := make([]byte, hash().Size())
	io.ReadFull(crand.Reader, r)
	return hex.EncodeToString(r)
}
