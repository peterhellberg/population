package population

import (
	"fmt"
	"time"
)

// RemainingLifeExpectancy contains remaining life expectancy of
// a person with given sex, country, and age at a given point in time
type RemainingLifeExpectancy struct {
	Sex       string  `json:"sex"`
	Country   string  `json:"country"`
	Date      string  `json:"date"`
	Age       string  `json:"age"`
	Remaining float64 `json:"remaining_life_expectancy"`
}

// RemainingLifeExpectancy calculates remaining life expectancy of a
// person with given sex, country, and age at a given point in time
//
// <http://api.population.io/#!/life-expectancy/calculateRemainingLifeExpectancy>
//
func (c *Client) RemainingLifeExpectancy(sex, country, date, age string) (RemainingLifeExpectancy, error) {
	path := fmt.Sprintf(`/life-expectancy/remaining/%s/%s/%s/%s/`, sex, country, date, age)
	resp := RemainingLifeExpectancy{}

	if _, err := time.Parse("2006-01-02", date); err != nil {
		return resp, err
	}

	err := c.Get(path, &resp)

	return resp, err
}

// TotalLifeExpectancy contains the total life expectancy of a person
type TotalLifeExpectancy struct {
	Sex     string  `json:"sex"`
	Country string  `json:"country"`
	Dob     string  `json:"dob"`
	Total   float64 `json:"total_life_expectancy"`
}

// TotalLifeExpectancy calculates total life expectancy of
// a person with given sex, country, and date of birth.
//
// <http://api.population.io/#!/life-expectancy/calculateTotalLifeExpectancy>
//
func (c *Client) TotalLifeExpectancy(sex, country, dob string) (TotalLifeExpectancy, error) {
	path := fmt.Sprintf(`/life-expectancy/total/%s/%s/%s/`, sex, country, dob)
	resp := TotalLifeExpectancy{}

	if _, err := time.Parse("2006-01-02", dob); err != nil {
		return resp, err
	}

	err := c.Get(path, &resp)

	return resp, err
}
