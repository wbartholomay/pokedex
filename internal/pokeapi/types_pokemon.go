package pokeapi

type RespPokemonDetails struct{
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}