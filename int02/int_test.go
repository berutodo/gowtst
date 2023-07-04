package inteiros

import "testing"

func TestAdd(t *testing.T) {
	t.Run("expect to have a sum of two numbers", func(t *testing.T) {
		result := Sum(2, 2)
		expect := 4

		t.Helper()
		if result != expect {
			t.Errorf("Expected %d, result: %d", expect, result)
		}

	})
}
