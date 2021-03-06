package population

import (
	"fmt"
	"time"
)

// WorldPopulationRank contains date of birth, sex, country and rank
type WorldPopulationRank struct {
	Dob     string `json:"dob"`
	Sex     string `json:"sex"`
	Country string `json:"country"`
	Rank    int    `json:"rank"`
}

// RankToday contains a world population rank
type RankToday struct {
	WorldPopulationRank
}

// RankToday calculates the world population rank of a person
// with the given date of birth, sex and country of origin as of today.
//
// <http://api.population.io/#!/wp-rank/worldPopulationRankToday>
//
func (c *Client) RankToday(sex, country, dob string) (RankToday, error) {
	path := fmt.Sprintf(`/wp-rank/%s/%s/%s/today/`, dob, sex, country)
	resp := RankToday{}

	if _, err := time.Parse("2006-01-02", dob); err != nil {
		return resp, err
	}

	err := c.Get(path, &resp)

	return resp, err
}

// RankByDate contains a world population rank with a date
type RankByDate struct {
	WorldPopulationRank
	Date string `json:"date"`
}

// RankByDate calculates the world population rank of a person with the
// given date of birth, sex and country of origin on a certain date.
//
// <http://api.population.io/#!/wp-rank/worldPopulationRankByDate>
//
func (c *Client) RankByDate(sex, country, dob, date string) (RankByDate, error) {
	path := fmt.Sprintf(`/wp-rank/%s/%s/%s/on/%s/`, dob, sex, country, date)
	resp := RankByDate{}

	if _, err := time.Parse("2006-01-02", dob); err != nil {
		return resp, err
	}

	if _, err := time.Parse("2006-01-02", date); err != nil {
		return resp, err
	}

	err := c.Get(path, &resp)

	return resp, err
}

// RankByAge contains a world population rank with an age
type RankByAge struct {
	WorldPopulationRank
	Age string `json:"age"`
}

// RankByAge calculates the world population rank of a person with
// the given date of birth, sex and country of origin on a certain
// date as expressed by the person's age.
//
// <http://api.population.io/#!/wp-rank/worldPopulationRankByAge>
//
func (c *Client) RankByAge(sex, country, dob, age string) (RankByAge, error) {
	path := fmt.Sprintf(`/wp-rank/%s/%s/%s/aged/%s/`, dob, sex, country, age)
	resp := RankByAge{}

	if _, err := time.Parse("2006-01-02", dob); err != nil {
		return resp, err
	}

	err := c.Get(path, &resp)

	return resp, err
}

// RankWithOffset contains a world population rank with an offset
type RankWithOffset struct {
	WorldPopulationRank
	Offset string `json:"offset"`
}

// RankInPast calculates the world population rank of a person with
// the given date of birth, sex and country of origin on a certain
// date as expressed by an offset towards the future from today.
//
// <http://api.population.io/#!/wp-rank/worldPopulationRankInPast>
//
func (c *Client) RankInPast(sex, country, dob, ago string) (RankWithOffset, error) {
	path := fmt.Sprintf(`/wp-rank/%s/%s/%s/ago/%s/`, dob, sex, country, ago)
	resp := RankWithOffset{}

	if _, err := time.Parse("2006-01-02", dob); err != nil {
		return resp, err
	}

	err := c.Get(path, &resp)

	return resp, err
}

// RankInFuture calculates the world population rank of a person with
// the given date of birth, sex and country of origin on a certain
// date as expressed by an offset towards the future from today.
//
// <http://api.population.io/#!/wp-rank/worldPopulationRankInFuture>
//
func (c *Client) RankInFuture(sex, country, dob, in string) (RankWithOffset, error) {
	path := fmt.Sprintf(`/wp-rank/%s/%s/%s/in/%s/`, dob, sex, country, in)
	resp := RankWithOffset{}

	if _, err := time.Parse("2006-01-02", dob); err != nil {
		return resp, err
	}

	err := c.Get(path, &resp)

	return resp, err
}

// DateByRank contains a world population rank with a date on rank
type DateByRank struct {
	WorldPopulationRank
	DateOnRank string `json:"date_on_rank"`
}

// DateByRank calculates the day on which a person with the given
// date of birth, sex and country of origin has reached (or will reach)
// a certain world population rank.
//
// <http://api.population.io/#!/wp-rank/dateByWorldPopulationRank>
//
func (c *Client) DateByRank(sex, country, dob string, rank int) (DateByRank, error) {
	path := fmt.Sprintf(`/wp-rank/%s/%s/%s/ranked/%d/`, dob, sex, country, rank)
	resp := DateByRank{}

	if _, err := time.Parse("2006-01-02", dob); err != nil {
		return resp, err
	}

	err := c.Get(path, &resp)

	return resp, err
}
