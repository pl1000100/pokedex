package pokeapi

type locationAreaResponse struct {
	Count    uint
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}
