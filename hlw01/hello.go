package main

import "fmt"

const prefixHello = "Olá, "

func Hello(name string) string {
	if name == "" {
		name = "mundo"
	}
	return prefixHello + name + "."
}

func main() {
	fmt.Println(Hello("mundo"))

}
