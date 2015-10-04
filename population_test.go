package population

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestTable(t *testing.T) {
	for i, tt := range []struct {
		path    string
		country string
		year    int
		age     int
		data    Table
	}{
		{
			"/population/1980/Brazil/18/",
			"Brazil",
			1980,
			18,
			Table{
				{
					Total:   2719710,
					Age:     18,
					Males:   1365687,
					Females: 1354023,
					Year:    1980,
				},
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

			json.NewEncoder(w).Encode(tt.data)
		})

		resp, err := c.Table(tt.country, tt.year, tt.age)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if got, want := len(resp), len(tt.data); got != want {
			t.Fatalf(`[%d] len(resp) = %d, want %d`, i, got, want)
		}
	}
}

func TestTableAllAges(t *testing.T) {
	for i, tt := range []struct {
		path    string
		country string
		year    int
		data    Table
	}{
		{
			"/population/1980/Brazil/",
			"Brazil",
			1980,
			Table{
				{
					Total:   3643507,
					Age:     0,
					Males:   1844792,
					Females: 1798715,
					Year:    1980,
				},
				{
					Total:   3510106,
					Age:     1,
					Males:   1772216,
					Females: 1737890,
					Year:    1980,
				},
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

			json.NewEncoder(w).Encode(tt.data)
		})

		resp, err := c.TableAllAges(tt.country, tt.year)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if got, want := len(resp), len(tt.data); got != want {
			t.Fatalf(`[%d] len(resp) = %d, want %d`, i, got, want)
		}
	}
}

func TestTableAllYears(t *testing.T) {
	for i, tt := range []struct {
		path    string
		country string
		age     int
		data    Table
	}{
		{
			"/population/Brazil/18/",
			"Brazil",
			18,
			Table{
				{
					Total:   1044742,
					Age:     18,
					Males:   523927,
					Females: 520815,
					Year:    1950,
				},
				{
					Total:   1059439,
					Age:     18,
					Males:   531574,
					Females: 527868,
					Year:    1951,
				},
			},
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"detail": fmt.Sprintf(`%q != %q`, r.URL.Path, tt.path),
				})
				return
			}

			json.NewEncoder(w).Encode(tt.data)
		})

		resp, err := c.TableAllYears(tt.country, tt.age)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if got, want := len(resp), len(tt.data); got != want {
			t.Fatalf(`[%d] len(resp) = %d, want %d`, i, got, want)
		}
	}
}

func TestTotalPopulationByDate(t *testing.T) {
	for i, tt := range []struct {
		path    string
		country string
		date    string
		data    TotalPopulation
	}{
		{
			"/population/Brazil/2015-12-24/",
			"Brazil",
			"2015-12-24",
			TotalPopulation{"2015-12-24", 204413867},
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"detail": fmt.Sprintf(`%q != %q`, r.URL.Path, tt.path),
				})
				return
			}

			json.NewEncoder(w).Encode(map[string]TotalPopulation{
				"total_population": tt.data,
			})
		})

		resp, err := c.TotalPopulationByDate(tt.country, tt.date)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if resp.Date != tt.date {
			t.Fatalf(`[%d] resp.Date = %q, want %q`, i, resp.Date, tt.date)
		}
	}
}

func TestTotalPopulationTodayAndTomorrow(t *testing.T) {
	for i, tt := range []struct {
		path     string
		country  string
		today    TotalPopulation
		tomorrow TotalPopulation
	}{
		{
			"/population/Brazil/today-and-tomorrow/",
			"Brazil",
			TotalPopulation{"2015-10-04", 204063621},
			TotalPopulation{"2015-10-05", 204067960},
		},
	} {
		ts, c := testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path != tt.path {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"detail": fmt.Sprintf(`%q != %q`, r.URL.Path, tt.path),
				})
				return
			}

			json.NewEncoder(w).Encode(map[string][]TotalPopulation{
				"total_population": []TotalPopulation{tt.today, tt.tomorrow},
			})
		})

		today, tomorrow, err := c.TotalPopulationTodayAndTomorrow(tt.country)
		ts.Close()

		if err != nil {
			t.Fatalf(`[%d] unexpected error: %v`, i, err)
		}

		if today.Date != tt.today.Date {
			t.Fatalf(`[%d] today.Date = %q, want %q`, i, today.Date, tt.today.Date)
		}

		if today.Population != tt.today.Population {
			t.Fatalf(`[%d] today.Population = %d, want %d`, i, today.Population, tt.today.Population)
		}

		if tomorrow.Date != tt.tomorrow.Date {
			t.Fatalf(`[%d] tomorrow.Date = %q, want %q`, i, tomorrow.Date, tt.tomorrow.Date)
		}

		if tomorrow.Population != tt.tomorrow.Population {
			t.Fatalf(`[%d] tomorrow.Population = %d, want %d`, i, tomorrow.Population, tt.tomorrow.Population)
		}
	}
}
