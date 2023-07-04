package main

import (
	"fmt"
	"strings"
)

const prefixHello = "Hello, "

func Hello(name string, language string) string {

	if name == "" {
		name = "world"
	}
	if strings.ToLower(language) == "spanish" {
		return "Hola, " + name
	}
	if strings.ToLower(language) == "french" {
		return "Bonjour, " + name
	}
	return prefixHello + name
}

func main() {
	fmt.Println(Hello("world", ""))

}
