package population

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestRemainingLifeExpectancy(t *testing.T) {
	for i, tt := range []struct {
		path      string
		sex       string
		country   string
		date      string
		age       string
		remaining float64
	}{
		{
			"/life-expectancy/remaining/male/United Kingdom/2001-05-11/49y2m/",
			"male",
			"United Kingdom",
			"2001-05-11",
			"49y2m",
			32.63468564999831,
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"detail":"bad request"}`))
				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"sex":     tt.sex,
				"country": tt.country,
				"date":    tt.date,
				"age":     tt.age,
				"remaining_life_expectancy": tt.remaining,
			})
		})

		resp, err := c.RemainingLifeExpectancy(tt.sex, tt.country, tt.date, tt.age)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Sex != tt.sex {
			t.Fatalf(`[%d] resp.Sex = %q, want %q`, i, resp.Sex, tt.sex)
		}

		if resp.Country != tt.country {
			t.Fatalf(`[%d] resp.Country = %q, want %q`, i, resp.Country, tt.country)
		}

		if resp.Date != tt.date {
			t.Fatalf(`[%d] resp.Date = %q, want %q`, i, resp.Date, tt.date)
		}

		if resp.Age != tt.age {
			t.Fatalf(`[%d] resp.Age = %q, want %q`, i, resp.Age, tt.age)
		}

		if resp.Remaining != tt.remaining {
			t.Fatalf(`[%d] resp.Remaining = %v, want %v`, i, resp.Remaining, tt.remaining)
		}
	}
}

func TestTotalLifeExpectancy(t *testing.T) {
	for i, tt := range []struct {
		path    string
		sex     string
		country string
		dob     string
		total   float64
	}{
		{
			"/life-expectancy/total/male/United Kingdom/1952-03-11/",
			"male",
			"United Kingdom",
			"1952-03-11",
			80.66075832186138,
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"detail":"bad request"}`))
				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"sex":     tt.sex,
				"country": tt.country,
				"dob":     tt.dob,
				"total_life_expectancy": tt.total,
			})
		})

		resp, err := c.TotalLifeExpectancy(tt.sex, tt.country, tt.dob)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Sex != tt.sex {
			t.Fatalf(`[%d] resp.Sex = %q, want %q`, i, resp.Sex, tt.sex)
		}

		if resp.Country != tt.country {
			t.Fatalf(`[%d] resp.Country = %q, want %q`, i, resp.Country, tt.country)
		}

		if resp.Dob != tt.dob {
			t.Fatalf(`[%d] resp.Dob = %q, want %q`, i, resp.Dob, tt.dob)
		}

		if resp.Total != tt.total {
			t.Fatalf(`[%d] resp.Total = %v, want %v`, i, resp.Total, tt.total)
		}
	}
}
