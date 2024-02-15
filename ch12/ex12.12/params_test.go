package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"testing"
)

type name struct {
	Name string `http:"name",check:"name"`
}

var mCheck = map[string]Check{
	"name": isName,
}

func isName(v any) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("isName(): %v must be a string", v)
	}
	re := regexp.MustCompile("[a-zA-Z]{1, 20}")
	if !re.MatchString(s) {
		return fmt.Errorf("Name only contains letters")
	}
	return nil
}

func TestCheck(t *testing.T) {
	tests := []struct {
		req  *http.Request
		want name
	}{
		{
			&http.Request{
				Form: url.Values{
					"name": []string{"Gabriel"},
				},
			},
			name{"Gabriel"},
		},
	}
	for _, test := range tests {
		var got name
		err := Unpack(test.req, &got, mCheck)
		if err != nil {
			t.Errorf("Unpack(): %s", err)
		}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("Unpack(): got %v, want %v", got, test.want)
		}
	}
}
