package population

import "fmt"

type MortalityDistribution struct {
	Age              float64 `json:"age"`
	MortalityPercent float64 `json:"mortality_percent"`
}

type MortalityDistributionTable struct {
	MortalityDistribution []MortalityDistribution `json:"mortality_distribution"`
}

func (c *Client) MortalityDistributionTable(country, sex, age string) (MortalityDistributionTable, error) {
	path := fmt.Sprintf(`/mortality-distribution/%s/%s/%s/today/`, country, sex, age)
	resp := MortalityDistributionTable{}

	err := c.Get(path, &resp)

	return resp, err
}
