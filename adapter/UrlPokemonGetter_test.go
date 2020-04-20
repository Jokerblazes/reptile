package adapter

import "testing"

func TestUrlPokemonGetter_pokemons(t *testing.T) {
	type fields struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"demoTest", fields{"https://pokedex.org/"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getter := UrlPokemonGetter{}
			pokemons := getter.Pokemons()
			if pokemons == nil {
				t.Fail()
			}
		})
	}
}
