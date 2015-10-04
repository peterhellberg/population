package population

// Countries is a list of countries as strings
type Countries []string

// CountryList contains a list of countries
type CountryList struct {
	Countries Countries `json:"countries"`
}

// ListCountries return a list of countries from the World Population API.
//
// <http://api.population.io/#!/countries/listCountries>
//
func (c *Client) ListCountries() (CountryList, error) {
	path := "/countries"
	resp := CountryList{}

	err := c.Get(path, &resp)

	return resp, err
}
