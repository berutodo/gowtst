package main

import "testing"

func TestHello(t *testing.T) {

	verifyCorrectAnswer := func(t *testing.T, expected, result string) {
		t.Helper()
		if result != expected {
			t.Errorf("O esperado '%s' não foi esperado '%s' : ", result, expected)
		}
	}
	t.Run(" expected to say hello with name", func(t *testing.T) {
		result := Hello("Gabriel")
		expected := "Olá, Gabriel."

		verifyCorrectAnswer(t, result, expected)
	})

	t.Run("Expected to say Hello World if receives a empty string", func(t *testing.T) {
		result := Hello("")
		expected := "Olá, mundo."

		verifyCorrectAnswer(t, result, expected)

	})

}
