package ft_test

import (
	"testing"

	"github.com/spobly/ft"
)

func TestGetCategory(t *testing.T) {
	extension := ".jpg"
	want := "Images"

	got := ft.GetCategory(extension)

	if got != want {
		t.Fatalf("for %s: want %s, got %s", extension, want, got)
	}
}
