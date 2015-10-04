package population

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestMortalityDistributionTable(t *testing.T) {
	for i, tt := range []struct {
		path    string
		country string
		sex     string
		age     string
		data    []MortalityDistribution
	}{
		{
			"/mortality-distribution/United Kingdom/male/49y2m/today/",
			"United Kingdom",
			"male",
			"49y2m",
			[]MortalityDistribution{
				{45, 0},
				{50, 0.25046935870135467},
			},
		},
		{
			"/mortality-distribution/Germany/female/30y6m/today/",
			"Germany",
			"female",
			"30y6m",
			[]MortalityDistribution{
				{45, 0},
				{50, 0.25046935870135467},
			},
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]interface{}{"detail": "invalid request"})
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"mortality_distribution": tt.data,
			})
		})

		resp, err := c.MortalityDistributionTable(tt.country, tt.sex, tt.age)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if got, want := len(resp.MortalityDistribution), len(tt.data); got != want {
			t.Fatalf(`len(resp.MortalityDistribution) = %d, want %d`, got, want)
		}
	}
}
