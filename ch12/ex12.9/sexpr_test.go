package sexpr

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

    // Encoding 
    data, err := Marshal(strangelove)
    if err != nil {
        t.Fatalf("Marshal failed: %v", err)
    }
    t.Logf("Marshal() = %s\n", data)

    // Decoding
    var movie Movie
    err = Unmarshal(data, &movie)
    if err != nil {
        t.Fatalf("Unmarshal failed: %v", err)
    }
    t.Logf("Unmarshal() = %+v\n", movie)

    // Check equality
    if !reflect.DeepEqual(strangelove, movie) {
        t.Fatal("Not equal")
    }

	// Decode it
	dec := NewDecoder(strings.NewReader(string(data)))
	for {
		_, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Errorf("Decode failed: %s", err)
			break
		}
	}

}
