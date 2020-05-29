package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Olá Mundo."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, quant %q", got, want)
	}
}
