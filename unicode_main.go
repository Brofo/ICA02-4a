package main

import (
	"fmt"

	"github.com/Brofo/ica02/oppg4/unicode"
)

const norsk = "\x6e\x6f\x72\x64\x6f\x67\x73\xc3\xb8\x72"
const japan = "jp"
const island = "is"

func main() {
	translation := unicode.Translate(norsk, island)
	fmt.Println(translation)
}
