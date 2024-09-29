package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := repeat("c", 8)
	want := "cccccccc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("c", 8)
	}
}
