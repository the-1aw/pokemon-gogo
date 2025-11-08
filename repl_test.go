package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"simple case":        {input: "hello world", want: []string{"hello", "world"}},
		"lots of spaces":     {input: "      hello          world      ", want: []string{"hello", "world"}},
		"pokemons with caps": {input: "Charmander Bulbasaur PIKACHU", want: []string{"charmander", "bulbasaur", "pikachu"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := CleanInput(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("expected: %#v, got: %#v", test.want, got)
			}

		})
	}
}
