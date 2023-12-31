package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	verifyCorrectAnswer := func(t *testing.T, expected, result string) {
		t.Helper()
		if result != expected {
			t.Errorf("O esperado '%s' não foi esperado '%s' : ", result, expected)
		}
	}
	t.Run(" expected to say hello with name", func(t *testing.T) {
		result := Hello("Gabriel", "")
		expected := "Hello, Gabriel"

		verifyCorrectAnswer(t, result, expected)
	})

	t.Run("Expected to say Hello World if receives a empty string", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world"

		verifyCorrectAnswer(t, result, expected)

	})

	t.Run("Expected to return the lang spanish", func(t *testing.T) {
		result := Hello("Gabriel", "spanish")
		expected := "Hola, Gabriel"

		verifyCorrectAnswer(t, result, expected)
	})

	t.Run("Expect to return in french", func(t *testing.T) {
		result := Hello("Gabriel", "french")
		expected := "Bonjour, Gabriel"
		verifyCorrectAnswer(t, result, expected)
	})

}
