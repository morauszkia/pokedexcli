package pokeapi

type LocationBasic struct {
	Name	string	`json:"name"`
	URL		string	`json:"url"`
}

type LocationAreasList struct {
	Next			*string				`json:"next"`
	Previous		*string				`json:"previous"`
	Results			[]LocationBasic		`json:"results"`
}