package main

import "fmt"

const prefixHello = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "mundo"
	}
	if language == "spanish" {
		return "Hola, " + name
	}
	if language == "french" {
		return "Bonjour, " + name
	}
	return prefixHello + name
}

func main() {
	fmt.Println(Hello("mundo", ""))

}
