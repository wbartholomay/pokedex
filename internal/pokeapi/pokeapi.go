package pokeapi

import (
	"github.com/wbartholomay/pokedexcli/internal/pokecache"
	"net/http"
	"io"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)


func makeCachedRequest(client *Client, url string, c pokecache.Cache) ([]byte, error){
	data, ok := c.Get(url)
	if ok { return data, nil }

	//only supporting get requests for now
	return makeGetRequest(client, url)
}

func makeGetRequest(client *Client, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}