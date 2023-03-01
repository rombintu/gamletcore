package core

import (
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

type Gams [3][]rune

func randRune() rune {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31()
}

func GenGams(size int) Gams {
	var g1, g2, g3 []rune
	for i := 0; i < size; i++ {
		g1 = append(g1, randRune())
		g2 = append(g2, randRune())
		g3 = append(g3, randRune())
	}
	return Gams{g1, g2, g3}
}

func Encode(secret string) (key []rune, gams Gams) {
	gams = GenGams(len(secret))
	for i, s := range secret {
		key = append(key, (s ^ gams[0][i] ^ gams[1][i] ^ gams[2][i]))
	}
	return
}

func Decode(key []rune, gams Gams) (secret []rune) {
	for i, k := range key {
		secret = append(secret, (k ^ gams[0][i] ^ gams[1][i] ^ gams[2][i]))
	}
	return secret
}

func Convert(payload []rune) (secret string) {
	for _, s := range payload {
		secret += string(s)
	}
	return secret
}

func DecodeFromHex(keyHex string, gamsHex []string) (secret []rune) {
	var key []rune
	var gams Gams
	for _, ks := range strings.Split(keyHex, " ") {
		rh, _ := utf8.DecodeRune([]byte(ks))
		key = append(key, rh)
	}
	for i, gh := range gamsHex {
		for _, ks := range strings.Split(gh, " ") {
			rh, _ := utf8.DecodeRune([]byte(ks))
			gams[i] = append(gams[i], rh)
		}
	}
	return Decode(key, gams)
}
