package mapper

import (
	"testing"

	"github.com/jkzjt/mygo/db/model"
)

func TestFindByConds(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1", 3},
		{"0", 2},
	}

	for _, c := range cases {
		got := FindByConds(model.Student{Gender: c.in})
		if c.want != len(got) {
			t.Errorf("in %s, want %v but got %v", c.in, c.want, len(got))
		}
	}
}

func TestFindAll(t *testing.T) {
	want := 5
	got := FindByConds()
	if len(got) != want {
		t.Errorf("want %v, but got %v", want, len(got))
	}
}
