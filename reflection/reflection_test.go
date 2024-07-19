package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Liam"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("got %v want %v", got[0], expected)
	}

}

func TestWalkNonNaive(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"liamc"},
			[]string{"liamc"},
		},
		{
			"struct with two string field",
			struct {
				Name    string
				Partner string
			}{"liamc", "pattyh"},
			[]string{"liamc", "pattyh"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"liamc", 34},
			[]string{"liamc"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
