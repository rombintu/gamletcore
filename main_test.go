package main_test

import (
	"fmt"
	"testing"

	"github.com/rombintu/gamletcore/core"
)

func TestNewGamma(t *testing.T) {
	secret := "secret testing"
	fmt.Println("SECRET:", secret)
	key, gams := core.Encode(secret)
	fmt.Printf(
		"KY: %X\nG1: %X\nG2: %X\nG3: %X\n",
		key, gams[0], gams[1], gams[2],
	)
	decoded := core.Decode(key, gams)
	fmt.Println("DECODED SECRET: ", core.Convert(decoded))
}
