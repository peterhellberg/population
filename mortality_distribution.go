package population

import "fmt"

// MortalityDistributionRow is a row in a mortality distribution table
type MortalityDistributionRow struct {
	Age              float64 `json:"age"`
	MortalityPercent float64 `json:"mortality_percent"`
}

// MortalityDistributionTable contains a mortality distribution table
type MortalityDistributionTable struct {
	MortalityDistribution []MortalityDistributionRow `json:"mortality_distribution"`
}

// MortalityDistributionTable retrieves the mortality distribution table for the given country, sex and age.
func (c *Client) MortalityDistributionTable(country, sex, age string) (MortalityDistributionTable, error) {
	path := fmt.Sprintf(`/mortality-distribution/%s/%s/%s/today/`, country, sex, age)
	resp := MortalityDistributionTable{}

	err := c.Get(path, &resp)

	return resp, err
}
