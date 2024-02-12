package params

import "testing"

type Ruler struct {
	Name    string `http:"n"`
	Age     int    `http:"a"`
	Country string `http:"c"`
}

func TestPack(t *testing.T) {

	tests := []struct {
		r    Ruler
		want string
	}{
		{Ruler{"Dracula", 593, "Ro"}, "a=593&c=Ro&n=Dracula"},
		{Ruler{"Caesar", 55, "SPQR"}, "a=55&c=SPQR&n=Caesar"},
		{Ruler{"Constantine", 65, "SPQR"}, "a=65&c=SPQR&n=Constantine"},
	}
	for _, test := range tests {
		u, err := Pack(&test.r)
		if err != nil {
			t.Errorf("Pack(%q): %q", test, err)
		}
		got := u.RawQuery
		if got != test.want {
			t.Errorf("Pack(%q), got %s, want %s", test, got, test.want)
		}
	}
}
