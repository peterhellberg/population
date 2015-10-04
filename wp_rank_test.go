package population

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestRankToday(t *testing.T) {
	for i, tt := range []struct {
		path    string
		sex     string
		country string
		dob     string
		rank    int
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/today/",
			"male",
			"United Kingdom",
			"1952-03-11",
			30000000,
		},
		{
			"/wp-rank/1988-03-11/female/Germany/today/",
			"female",
			"Germany",
			"1988-03-11",
			10000000,
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"sex":     tt.sex,
				"country": tt.country,
				"dob":     tt.dob,
				"rank":    tt.rank,
			})
		})

		resp, err := c.RankToday(tt.sex, tt.country, tt.dob)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Dob != tt.dob {
			t.Fatalf(`[%d] resp.Dob = %q, want %q`, i, resp.Dob, tt.dob)
		}

		if resp.Sex != tt.sex {
			t.Fatalf(`[%d] resp.Sex = %q, want %q`, i, resp.Sex, tt.sex)
		}

		if resp.Country != tt.country {
			t.Fatalf(`[%d] resp.Country = %q, want %q`, i, resp.Country, tt.country)
		}

		if resp.Rank != tt.rank {
			t.Fatalf(`[%d] resp.Rank = %q, want %q`, i, resp.Rank, tt.rank)
		}
	}
}

func TestRankByDate(t *testing.T) {
	for i, tt := range []struct {
		path    string
		sex     string
		country string
		dob     string
		date    string
		rank    int
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/on/2015-10-04/",
			"male",
			"United Kingdom",
			"1952-03-11",
			"2015-10-04",
			30000000,
		},
		{
			"/wp-rank/1988-03-11/female/Germany/on/2015-10-04/",
			"female",
			"Germany",
			"1988-03-11",
			"2015-10-04",
			10000000,
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"sex":     tt.sex,
				"country": tt.country,
				"dob":     tt.dob,
				"date":    tt.date,
				"rank":    tt.rank,
			})
		})

		resp, err := c.RankByDate(tt.sex, tt.country, tt.dob, tt.date)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Dob != tt.dob {
			t.Fatalf(`[%d] resp.Dob = %q, want %q`, i, resp.Dob, tt.dob)
		}

		if resp.Sex != tt.sex {
			t.Fatalf(`[%d] resp.Sex = %q, want %q`, i, resp.Sex, tt.sex)
		}

		if resp.Country != tt.country {
			t.Fatalf(`[%d] resp.Country = %q, want %q`, i, resp.Country, tt.country)
		}

		if resp.Rank != tt.rank {
			t.Fatalf(`[%d] resp.Rank = %q, want %q`, i, resp.Rank, tt.rank)
		}

		if resp.Date != tt.date {
			t.Fatalf(`[%d] resp.Date = %q, want %q`, i, resp.Date, tt.date)
		}
	}
}

func TestRankByAge(t *testing.T) {
	for i, tt := range []struct {
		path    string
		sex     string
		country string
		dob     string
		age     string
		rank    int
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/aged/49y2m/",
			"male",
			"United Kingdom",
			"1952-03-11",
			"49y2m",
			30000000,
		},
		{
			"/wp-rank/1988-03-11/female/Germany/aged/30y6m/",
			"female",
			"Germany",
			"1988-03-11",
			"30y6m",
			10000000,
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"sex":     tt.sex,
				"country": tt.country,
				"dob":     tt.dob,
				"age":     tt.age,
				"rank":    tt.rank,
			})
		})

		resp, err := c.RankByAge(tt.sex, tt.country, tt.dob, tt.age)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Dob != tt.dob {
			t.Fatalf(`[%d] resp.Dob = %q, want %q`, i, resp.Dob, tt.dob)
		}

		if resp.Sex != tt.sex {
			t.Fatalf(`[%d] resp.Sex = %q, want %q`, i, resp.Sex, tt.sex)
		}

		if resp.Country != tt.country {
			t.Fatalf(`[%d] resp.Country = %q, want %q`, i, resp.Country, tt.country)
		}

		if resp.Rank != tt.rank {
			t.Fatalf(`[%d] resp.Rank = %q, want %q`, i, resp.Rank, tt.rank)
		}

		if resp.Age != tt.age {
			t.Fatalf(`[%d] resp.Age = %q, want %q`, i, resp.Age, tt.age)
		}
	}
}

func TestRankInPast(t *testing.T) {
	for i, tt := range []struct {
		path    string
		sex     string
		country string
		dob     string
		offset  string
		rank    int
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/ago/9y2m/",
			"male",
			"United Kingdom",
			"1952-03-11",
			"9y2m",
			30000000,
		},
		{
			"/wp-rank/1988-03-11/female/Germany/ago/1y6m/",
			"female",
			"Germany",
			"1988-03-11",
			"1y6m",
			10000000,
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"sex":     tt.sex,
				"country": tt.country,
				"dob":     tt.dob,
				"offset":  tt.offset,
				"rank":    tt.rank,
			})
		})

		resp, err := c.RankInPast(tt.sex, tt.country, tt.dob, tt.offset)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Dob != tt.dob {
			t.Fatalf(`[%d] resp.Dob = %q, want %q`, i, resp.Dob, tt.dob)
		}

		if resp.Sex != tt.sex {
			t.Fatalf(`[%d] resp.Sex = %q, want %q`, i, resp.Sex, tt.sex)
		}

		if resp.Country != tt.country {
			t.Fatalf(`[%d] resp.Country = %q, want %q`, i, resp.Country, tt.country)
		}

		if resp.Rank != tt.rank {
			t.Fatalf(`[%d] resp.Rank = %q, want %q`, i, resp.Rank, tt.rank)
		}

		if resp.Offset != tt.offset {
			t.Fatalf(`[%d] resp.Offset = %q, want %q`, i, resp.Offset, tt.offset)
		}
	}
}

func TestRankInFuture(t *testing.T) {
	for i, tt := range []struct {
		path    string
		sex     string
		country string
		dob     string
		offset  string
		rank    int
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/in/9y2m/",
			"male",
			"United Kingdom",
			"1952-03-11",
			"9y2m",
			30000000,
		},
		{
			"/wp-rank/1988-03-11/female/Germany/in/1y6m/",
			"female",
			"Germany",
			"1988-03-11",
			"1y6m",
			10000000,
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"sex":     tt.sex,
				"country": tt.country,
				"dob":     tt.dob,
				"offset":  tt.offset,
				"rank":    tt.rank,
			})
		})

		resp, err := c.RankInFuture(tt.sex, tt.country, tt.dob, tt.offset)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Dob != tt.dob {
			t.Fatalf(`[%d] resp.Dob = %q, want %q`, i, resp.Dob, tt.dob)
		}

		if resp.Sex != tt.sex {
			t.Fatalf(`[%d] resp.Sex = %q, want %q`, i, resp.Sex, tt.sex)
		}

		if resp.Country != tt.country {
			t.Fatalf(`[%d] resp.Country = %q, want %q`, i, resp.Country, tt.country)
		}

		if resp.Rank != tt.rank {
			t.Fatalf(`[%d] resp.Rank = %q, want %q`, i, resp.Rank, tt.rank)
		}

		if resp.Offset != tt.offset {
			t.Fatalf(`[%d] resp.Offset = %q, want %q`, i, resp.Offset, tt.offset)
		}
	}
}

func TestDateByRank(t *testing.T) {
	for i, tt := range []struct {
		path       string
		sex        string
		country    string
		dob        string
		dateOnRank string
		rank       int
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/ranked/30000000/",
			"male",
			"United Kingdom",
			"1952-03-11",
			"2025-10-19",
			30000000,
		},
		{
			"/wp-rank/1988-03-11/female/Germany/ranked/10000000/",
			"female",
			"Germany",
			"1988-03-11",
			"2065-03-12",
			10000000,
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"sex":          tt.sex,
				"country":      tt.country,
				"dob":          tt.dob,
				"date_on_rank": tt.dateOnRank,
				"rank":         tt.rank,
			})
		})

		resp, err := c.DateByRank(tt.sex, tt.country, tt.dob, tt.rank)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Dob != tt.dob {
			t.Fatalf(`[%d] resp.Dob = %q, want %q`, i, resp.Dob, tt.dob)
		}

		if resp.Sex != tt.sex {
			t.Fatalf(`[%d] resp.Sex = %q, want %q`, i, resp.Sex, tt.sex)
		}

		if resp.Country != tt.country {
			t.Fatalf(`[%d] resp.Country = %q, want %q`, i, resp.Country, tt.country)
		}

		if resp.Rank != tt.rank {
			t.Fatalf(`[%d] resp.Rank = %q, want %q`, i, resp.Rank, tt.rank)
		}

		if resp.DateOnRank != tt.dateOnRank {
			t.Fatalf(`[%d] resp.DateOnRank = %q, want %q`, i, resp.DateOnRank, tt.dateOnRank)
		}
	}
}
