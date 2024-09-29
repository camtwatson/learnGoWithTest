// This is the beginning of me learning Go with Learn Go With Test

package main

import (
	"fmt"
)

const (
	helloPrefix   = "Hello, "
	holaPrefix    = "Hola, "
	bonjourPrefix = "Bonjour, "

	spanish = "spanish"
	french  = "french"
)

func hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return greeting(lang) + name
}

func greeting(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = holaPrefix
	case french:
		prefix = bonjourPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(hello("world", " "))
}
