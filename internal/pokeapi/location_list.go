package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// List the Locations
func (c *Client) GetLocationArea(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationArea{}, err
	}

	var areas LocationArea
	if err = json.Unmarshal(body, &areas); err != nil {
		fmt.Println(err)
		return LocationArea{}, err
	}

	return areas, nil
}
