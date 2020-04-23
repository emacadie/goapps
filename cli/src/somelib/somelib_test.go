package somelib

import (
	"testing"
)

func Test_BasicChecks(t *testing.T) {
	t.Parallel()
	t.Run("Go can add", func(t *testing.T) {
		if 1+1 != 2 {
			t.Fail()
		} else {
			println("Adding is good")
		}
	})

	t.Run("Go can concat strings", func(t *testing.T) {
		if "Hello, " + "Go" != "Hello, Go" {
			t.Fail()
		}
	})
}

/*
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
*/
func TestAdd(t *testing.T) {
	if 1+1 != 2 {
		t.Fail()
	}
}

