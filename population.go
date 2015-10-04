package population

import (
	"fmt"
	"time"
)

// Table contains a list of table rows
type Table []TableRow

// TableRow contains population table info about a specific age group in the given country.
type TableRow struct {
	Total   int `json:"total"`
	Females int `json:"females"`
	Males   int `json:"males"`
	Year    int `json:"year"`
	Age     int `json:"age"`
}

// Table retrieves population table for a specific
// age group in the given year and country.
//
// <http://api.population.io/#!/population/retrievePopulationTable>
//
func (c *Client) Table(country string, year, age int) (Table, error) {
	path := fmt.Sprintf(`/population/%d/%s/%d/`, year, country, age)
	resp := Table{}

	err := c.Get(path, &resp)

	return resp, err
}

// TableAllAges retrieves population tables for a given year
// and country. Returns tables for all ages from 0 to 100.
//
// <http://api.population.io/#!/population/retrievePopulationTableAllAges>
//
func (c *Client) TableAllAges(country string, year int) (Table, error) {
	path := fmt.Sprintf(`/population/%d/%s/`, year, country)
	resp := Table{}

	err := c.Get(path, &resp)

	return resp, err
}

// TableAllYears retrieves population tables for a specific
// age group in the given country. Returns tables for all
// years from 1950 to 2100.
//
// <http://api.population.io/#!/population/retrievePopulationTableAllYears>
//
func (c *Client) TableAllYears(country string, age int) (Table, error) {
	path := fmt.Sprintf(`/population/%s/%d/`, country, age)
	resp := Table{}

	err := c.Get(path, &resp)

	return resp, err
}

// TotalPopulation contains the total population of a country on a given date.
type TotalPopulation struct {
	Date       string `json:"date"`
	Population int    `json:"population"`
}

// TotalPopulationByDate determines total population for a given country
// on a given date. Valid dates are 2013-01-01 to 2022-12-31.
//
// <http://api.population.io/#!/population/determineTotalPopulationByDate>
//
func (c *Client) TotalPopulationByDate(country, date string) (TotalPopulation, error) {
	path := fmt.Sprintf(`/population/%s/%s/`, country, date)
	resp := TotalPopulation{}

	if _, err := time.Parse("2006-01-02", date); err != nil {
		return resp, err
	}

	var doc map[string]TotalPopulation

	err := c.Get(path, &doc)

	return doc["total_population"], err
}

// TotalPopulationTodayAndTomorrow determines total population for a
// given country with separate results for today and tomorrow.
//
// <http://api.population.io/#!/population/determineTotalPopulationTodayAndTomorrow>
//
func (c *Client) TotalPopulationTodayAndTomorrow(country string) (TotalPopulation, TotalPopulation, error) {
	path := fmt.Sprintf(`/population/%s/today-and-tomorrow/`, country)

	var doc = map[string][]TotalPopulation{
		"total_population": []TotalPopulation{},
	}

	err := c.Get(path, &doc)

	tp := doc["total_population"]

	if len(tp) != 2 {
		return TotalPopulation{}, TotalPopulation{}, fmt.Errorf(`wrong number of elements in response`)
	}

	return tp[0], tp[1], err
}
