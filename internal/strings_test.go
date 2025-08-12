package internal

import "testing"

func TestGenerateRandomString(t *testing.T) {
	t.Parallel()

	t.Run("should generate a random string", func(t *testing.T) {
		got := GenerateRandomString(10)
		if len(got) != 10 {
			t.Fatalf("GenerateRandomString() = %v, want len = 10", got)
		}
	})
}
