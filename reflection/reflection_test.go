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

func TestWalkNoNaive(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"liamc"},
			[]string{"liamc"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"liamc", "NYC"},
			[]string{"liamc", "NYC"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"liamc", 33},
			[]string{"liamc"},
		},
		{
			"nested fields",
			Person{
				"liamc",
				Profile{33, "NYC"},
			},
			[]string{"liamc", "NYC"},
		},
		{
			"pointers to things",
			&Person{
				"liamc",
				Profile{33, "NYC"},
			},
			[]string{"liamc", "NYC"},
		},
		{
			"slices",
			[]Profile{
				{33, "NYC"},
				{34, "AMS"},
			},
			[]string{"NYC", "AMS"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "NYC"},
				{34, "AMS"},
			},
			[]string{"NYC", "AMS"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Grand Rapids"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Grand Rapids"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Grand Rapids"}
		}

		var got []string
		want := []string{"Berlin", "Grand Rapids"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
