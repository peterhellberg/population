package population

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestRankToday(t *testing.T) {
	for i, tt := range []struct {
		path    string
		dob     string
		sex     string
		country string
		rank    int
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/today/",
			"1952-03-11",
			"male",
			"United Kingdom",
			30000000,
		},
		{
			"/wp-rank/1988-03-11/female/Germany/today/",
			"1988-03-11",
			"female",
			"Germany",
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
				"dob":     tt.dob,
				"sex":     tt.sex,
				"country": tt.country,
				"rank":    tt.rank,
			})
		})

		resp, err := c.RankToday(tt.dob, tt.sex, tt.country)
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
		dob     string
		sex     string
		country string
		rank    int
		date    string
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/on/2015-10-04/",
			"1952-03-11",
			"male",
			"United Kingdom",
			30000000,
			"2015-10-04",
		},
		{
			"/wp-rank/1988-03-11/female/Germany/on/2015-10-04/",
			"1988-03-11",
			"female",
			"Germany",
			10000000,
			"2015-10-04",
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"dob":     tt.dob,
				"sex":     tt.sex,
				"country": tt.country,
				"rank":    tt.rank,
				"date":    tt.date,
			})
		})

		resp, err := c.RankByDate(tt.dob, tt.sex, tt.country, tt.date)
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
		dob     string
		sex     string
		country string
		rank    int
		age     string
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/aged/49y2m/",
			"1952-03-11",
			"male",
			"United Kingdom",
			30000000,
			"49y2m",
		},
		{
			"/wp-rank/1988-03-11/female/Germany/aged/30y6m/",
			"1988-03-11",
			"female",
			"Germany",
			10000000,
			"30y6m",
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"dob":     tt.dob,
				"sex":     tt.sex,
				"country": tt.country,
				"rank":    tt.rank,
				"age":     tt.age,
			})
		})

		resp, err := c.RankByAge(tt.dob, tt.sex, tt.country, tt.age)
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
		dob     string
		sex     string
		country string
		rank    int
		offset  string
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/ago/9y2m/",
			"1952-03-11",
			"male",
			"United Kingdom",
			30000000,
			"9y2m",
		},
		{
			"/wp-rank/1988-03-11/female/Germany/ago/1y6m/",
			"1988-03-11",
			"female",
			"Germany",
			10000000,
			"1y6m",
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"dob":     tt.dob,
				"sex":     tt.sex,
				"country": tt.country,
				"rank":    tt.rank,
				"offset":  tt.offset,
			})
		})

		resp, err := c.RankInPast(tt.dob, tt.sex, tt.country, tt.offset)
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
		dob     string
		sex     string
		country string
		rank    int
		offset  string
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/in/9y2m/",
			"1952-03-11",
			"male",
			"United Kingdom",
			30000000,
			"9y2m",
		},
		{
			"/wp-rank/1988-03-11/female/Germany/in/1y6m/",
			"1988-03-11",
			"female",
			"Germany",
			10000000,
			"1y6m",
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"dob":     tt.dob,
				"sex":     tt.sex,
				"country": tt.country,
				"rank":    tt.rank,
				"offset":  tt.offset,
			})
		})

		resp, err := c.RankInFuture(tt.dob, tt.sex, tt.country, tt.offset)
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
		dob        string
		sex        string
		country    string
		rank       int
		dateOnRank string
	}{
		{
			"/wp-rank/1952-03-11/male/United Kingdom/ranked/30000000/",
			"1952-03-11",
			"male",
			"United Kingdom",
			30000000,
			"2025-10-19",
		},
		{
			"/wp-rank/1988-03-11/female/Germany/ranked/10000000/",
			"1988-03-11",
			"female",
			"Germany",
			10000000,
			"2065-03-12",
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]interface{}{
				"dob":          tt.dob,
				"sex":          tt.sex,
				"country":      tt.country,
				"rank":         tt.rank,
				"date_on_rank": tt.dateOnRank,
			})
		})

		resp, err := c.DateByRank(tt.dob, tt.sex, tt.country, tt.rank)
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
