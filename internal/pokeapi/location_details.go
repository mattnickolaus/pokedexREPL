package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetails(locationName string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("Used explore Cache")
		fmt.Println("----------")

		// Use the cache
		var locDetails LocationDetails
		if err := json.Unmarshal(val, &locDetails); err != nil {
			return LocationDetails{}, err
		}
		return locDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationDetails{}, err
	}

	c.cache.Add(url, body)

	var locDetails LocationDetails
	if err := json.Unmarshal(body, &locDetails); err != nil {
		fmt.Println(err)
		return LocationDetails{}, err
	}

	return locDetails, nil
}
